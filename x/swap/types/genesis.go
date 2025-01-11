package types

import (
	"fmt"
)

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	// Check for duplicated index in incomingPacket
	incomingPacketIndexMap := make(map[string]struct{})

	for _, elem := range gs.IncomingInFlightPackets {
		index := string(IncomingInFlightPacketKey(elem.Index))
		if _, ok := incomingPacketIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for incomingInFlightPacket")
		}
		incomingPacketIndexMap[index] = struct{}{}
	}
	// Check for duplicated index in outgoingInFlightPacket
	outgoingInFlightPacketIndexMap := make(map[string]struct{})

	for _, elem := range gs.OutgoingInFlightPackets {
		index := string(OutgoingInFlightPacketKey(elem.Index))
		if _, ok := outgoingInFlightPacketIndexMap[index]; ok {
			return fmt.Errorf("duplicated index for outgoingInFlightPacket")
		}
		outgoingInFlightPacketIndexMap[index] = struct{}{}
	}

	return gs.Params.Validate()
}
