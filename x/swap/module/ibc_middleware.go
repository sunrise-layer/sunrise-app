package swap

import (
	"encoding/json"
	"fmt"
	"strings"

	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v8/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v8/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v8/modules/core/05-port/types"
	exported "github.com/cosmos/ibc-go/v8/modules/core/exported"
	"github.com/gogo/protobuf/jsonpb"

	// packetforwardtypes "github.com/cosmos/ibc-apps/middleware/packet-forward-middleware/v8/packetforward/types"

	keeper "github.com/sunriselayer/sunrise/x/swap/keeper"
	types "github.com/sunriselayer/sunrise/x/swap/types"
)

type IBCMiddleware struct {
	porttypes.IBCModule
	keeper *keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddleware given the keeper and underlying application.
func NewIBCMiddleware(
	app porttypes.IBCModule,
	k *keeper.Keeper,
) IBCMiddleware {
	return IBCMiddleware{
		IBCModule: app,
		keeper:    k,
	}
}

// receiveFunds receives funds from the packet into the override receiver
// address and returns an error if the funds cannot be received.
func (im IBCMiddleware) receiveFunds(
	ctx sdk.Context,
	packet channeltypes.Packet,
	data transfertypes.FungibleTokenPacketData,
	overrideReceiver string,
	relayer sdk.AccAddress,
) (exported.Acknowledgement, error) {
	overrideData := transfertypes.FungibleTokenPacketData{
		Denom:    data.Denom,
		Amount:   data.Amount,
		Sender:   data.Sender,
		Receiver: overrideReceiver, // override receiver
		// Memo explicitly zeroed
	}
	overrideDataBz := transfertypes.ModuleCdc.MustMarshalJSON(&overrideData)
	overridePacket := channeltypes.Packet{
		Sequence:           packet.Sequence,
		SourcePort:         packet.SourcePort,
		SourceChannel:      packet.SourceChannel,
		DestinationPort:    packet.DestinationPort,
		DestinationChannel: packet.DestinationChannel,
		Data:               overrideDataBz, // override data
		TimeoutHeight:      packet.TimeoutHeight,
		TimeoutTimestamp:   packet.TimeoutTimestamp,
	}

	ack := im.IBCModule.OnRecvPacket(ctx, overridePacket, relayer)

	if ack == nil {
		return ack, fmt.Errorf("ack is nil")
	}

	if !ack.Success() {
		return ack, fmt.Errorf("ack error: %s", string(ack.Acknowledgement()))
	}

	return ack, nil
}

// OnRecvPacket checks the memo field on this packet and if the metadata inside's root key indicates this packet
// should be handled by the swap middleware it attempts to perform a swap. If the swap is successful
// the underlying application's OnRecvPacket callback is invoked, an ack error is returned otherwise.
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) exported.Acknowledgement {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		// If this happens either a) a user has crafted an invalid packet, b) a
		// software developer has connected the middleware to a stack that does
		// not have a transfer module, or c) the transfer module has been modified
		// to accept other Packets. The best thing we can do here is pass the packet
		// on down the stack.
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	d := make(map[string]interface{})
	err := json.Unmarshal([]byte(data.Memo), &d)
	if err != nil || d["swap"] == nil {
		return im.IBCModule.OnRecvPacket(ctx, packet, relayer)
	}

	m := &types.PacketMetadata{}
	err = jsonpb.Unmarshal(strings.NewReader(data.Memo), m)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(fmt.Errorf("error parsing swap metadata: %w", err))
	}

	metadata := *m.Swap
	fmt.Println("------------------metadata----------------", metadata)

	if err := metadata.Validate(); err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	// Swap
	denomIn := types.GetDenomForThisChain(
		packet.DestinationPort,
		packet.DestinationChannel,
		packet.SourcePort,
		packet.SourceChannel,
		data.Denom,
	)
	if metadata.Route.DenomIn != denomIn {
		return channeltypes.NewErrorAcknowledgement(fmt.Errorf("invalid route: expected %s, got %s", metadata.Route.DenomIn, denomIn))
	}

	// Settle the incoming fund
	swapper := im.keeper.AccountKeeper.GetModuleAddress(types.ModuleName)
	incomingAck, err := im.receiveFunds(ctx, packet, data, swapper.String(), relayer)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	result, interfaceFee, err := im.keeper.SwapIncomingFund(
		ctx,
		packet,
		swapper,
		data,
		metadata,
	)
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	fmt.Println("===================swap result====================", result)

	waitingPacket, err := im.keeper.ProcessSwappedFund(
		ctx,
		packet,
		swapper,
		data,
		metadata,
		result,
		interfaceFee,
		incomingAck,
	)

	fmt.Println("===================waiting packet====================", waitingPacket)

	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	if waitingPacket != nil {
		im.keeper.SetIncomingInFlightPacket(ctx, *waitingPacket)

		// Returning nil ack will prevent WriteAcknowledgement from occurring for forwarded packet.
		// This is intentional so that the acknowledgement will be written later based on the ack/timeout of the forwarded packet.
		return nil
	}

	fullAck := types.SwapAcknowledgement{
		Result:      result,
		IncomingAck: incomingAck.Acknowledgement(),
	}

	bz, err := fullAck.Acknowledgement()
	if err != nil {
		return channeltypes.NewErrorAcknowledgement(err)
	}

	return channeltypes.NewResultAcknowledgement(bz)
}

// OnAcknowledgementPacket implements the IBCModule interface.
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	inflightPacket, found := im.keeper.GetOutgoingInFlightPacket(ctx, packet.SourcePort, packet.SourceChannel, packet.Sequence)
	if !found {
		return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
	}

	err := im.keeper.OnAcknowledgementOutgoingInFlightPacket(ctx, packet, acknowledgement, inflightPacket)
	if err != nil {
		return err
	}

	return im.IBCModule.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface.
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	var data transfertypes.FungibleTokenPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &data); err != nil {
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	inflightPacket, found := im.keeper.GetOutgoingInFlightPacket(ctx, packet.SourcePort, packet.SourceChannel, packet.Sequence)
	if !found {
		return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
	}

	if err := im.keeper.OnTimeoutOutgoingInFlightPacket(ctx, packet, inflightPacket); err != nil {
		return err
	}

	return im.IBCModule.OnTimeoutPacket(ctx, packet, relayer)
}
