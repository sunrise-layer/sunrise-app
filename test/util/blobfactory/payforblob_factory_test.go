package blobfactory_test

// import (
// 	"testing"

// 	"github.com/sunriselayer/sunrise/app/encoding"
// 	"github.com/sunriselayer/sunrise/pkg/blob"
// 	"github.com/sunriselayer/sunrise/test/util/blobfactory"
// 	testencoding "github.com/sunriselayer/sunrise/test/util/encoding"
// 	"github.com/sunriselayer/sunrise/test/util/testnode"

// 	tmrand "github.com/cometbft/cometbft/libs/rand"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// // TestRandMultiBlobTxsSameSigner_Deterministic tests whether with the same random seed the RandMultiBlobTxsSameSigner function produces the same blob txs.
// func TestRandMultiBlobTxsSameSigner_Deterministic(t *testing.T) {
// 	pfbCount := 10
// 	signer, err := testnode.NewOfflineSigner()
// 	require.NoError(t, err)
// 	encCfg := encoding.MakeConfig(testencoding.ModuleEncodingRegisters...)
// 	decoder := encCfg.TxConfig.TxDecoder()

// 	rand1 := tmrand.NewRand()
// 	rand1.Seed(1)
// 	marshalledBlobTxs1 := blobfactory.RandMultiBlobTxsSameSigner(t, rand1, signer, pfbCount)

// 	signer.ForceSetSequence(0)
// 	rand2 := tmrand.NewRand()
// 	rand2.Seed(1)
// 	marshalledBlobTxs2 := blobfactory.RandMultiBlobTxsSameSigner(t, rand2, signer, pfbCount)

// 	// additional checks for the sake of future debugging
// 	for index := 0; index < pfbCount; index++ {
// 		blobTx1, isBlob := blob.UnmarshalBlobTx(marshalledBlobTxs1[index])
// 		assert.True(t, isBlob)
// 		pfbMsgs1, err := decoder(blobTx1.Tx)
// 		assert.NoError(t, err)

// 		blobTx2, isBlob := blob.UnmarshalBlobTx(marshalledBlobTxs2[index])
// 		assert.True(t, isBlob)
// 		pfbMsgs2, err := decoder(blobTx2.Tx)
// 		assert.NoError(t, err)

// 		assert.Equal(t, blobTx1.Blobs, blobTx2.Blobs)
// 		assert.Equal(t, pfbMsgs1, pfbMsgs2)

// 		msgs2 := pfbMsgs2.GetMsgs()
// 		msgs1 := pfbMsgs1.GetMsgs()
// 		for i, msg := range msgs1 {
// 			assert.Equal(t, msg, msgs2[i])
// 		}
// 	}

// 	assert.Equal(t, marshalledBlobTxs1, marshalledBlobTxs2)
// }
