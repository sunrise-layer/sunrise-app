package keeper_test

// import (
// 	"bytes"
// 	"fmt"
// 	"testing"

// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	proto "github.com/gogo/protobuf/proto"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// 	"github.com/sunriselayer/sunrise/pkg/appconsts"
// 	"github.com/sunriselayer/sunrise/pkg/blob"
// 	appns "github.com/sunriselayer/sunrise/pkg/namespace"
// 	testkeeper "github.com/sunriselayer/sunrise/testutil/keeper"
// 	"github.com/sunriselayer/sunrise/x/blob/keeper"
// 	"github.com/sunriselayer/sunrise/x/blob/types"
// )

// // TestPayForBlobs verifies the attributes on the emitted event.
// func TestPayForBlobs(t *testing.T) {
// 	k, ctx := testkeeper.BlobKeeper(t)
// 	signer := "sunrise155u042u8wk3al32h3vzxu989jj76k4zcc6d03n"
// 	namespace := appns.MustNewV0(bytes.Repeat([]byte{1}, appns.NamespaceVersionZeroIDSize))
// 	namespaces := [][]byte{namespace.Bytes()}
// 	blobData := []byte("blob")
// 	blobSizes := []uint32{uint32(len(blobData))}

// 	// verify no events exist yet
// 	events := ctx.EventManager().Events().ToABCIEvents()
// 	assert.Len(t, events, 0)

// 	// emit an event by submitting a PayForBlob
// 	msg := createMsgPayForBlob(t, signer, namespace, blobData)
// 	server := keeper.NewMsgServerImpl(k)
// 	_, err := server.PayForBlobs(ctx, msg)
// 	require.NoError(t, err)

// 	// verify that an event was emitted
// 	events = ctx.EventManager().Events().ToABCIEvents()
// 	assert.Len(t, events, 1)
// 	protoEvent, err := sdk.ParseTypedEvent(events[0])
// 	require.NoError(t, err)
// 	event, err := convertToEventPayForBlobs(protoEvent)
// 	require.NoError(t, err)

// 	// verify the attributes of the event
// 	assert.Equal(t, signer, event.Signer)
// 	assert.Equal(t, namespaces, event.Namespaces)
// 	assert.Equal(t, blobSizes, event.BlobSizes)
// }

// func convertToEventPayForBlobs(message proto.Message) (*types.EventPayForBlobs, error) {
// 	if event, ok := message.(*types.EventPayForBlobs); ok {
// 		return event, nil
// 	}
// 	return nil, fmt.Errorf("message is not of type EventPayForBlobs")
// }

// func createMsgPayForBlob(t *testing.T, signer string, namespace appns.Namespace, blobData []byte) *types.MsgPayForBlobs {
// 	blob := blob.New(namespace, blobData, appconsts.ShareVersionZero)
// 	msg, err := types.NewMsgPayForBlobs(signer, blob)
// 	require.NoError(t, err)
// 	return msg
// }
