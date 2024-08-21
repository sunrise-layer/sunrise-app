package blobfactory_test

// import (
// 	"testing"

// 	"github.com/sunriselayer/sunrise/app/encoding"
// 	"github.com/sunriselayer/sunrise/pkg/user"
// 	"github.com/sunriselayer/sunrise/test/util/blobfactory"
// 	testencoding "github.com/sunriselayer/sunrise/test/util/encoding"
// 	"github.com/sunriselayer/sunrise/test/util/testfactory"
// 	"github.com/sunriselayer/sunrise/test/util/testnode"

// 	tmrand "github.com/cometbft/cometbft/libs/rand"
// 	"github.com/stretchr/testify/assert"
// 	"github.com/stretchr/testify/require"
// )

// // TestGenerateManyRandomRawSendTxsSameSigner_Deterministic tests whether with the same random seed the GenerateManyRandomRawSendTxsSameSigner function produces the same send transactions.
// func TestGenerateManyRandomRawSendTxsSameSigner_Deterministic(t *testing.T) {
// 	normalTxCount := 10
// 	encCfg := encoding.MakeConfig(testencoding.ModuleBasics)
// 	TxDecoder := encCfg.TxConfig.TxDecoder()

// 	kr, addr := testnode.NewKeyring(testfactory.TestAccName)
// 	signer, err := user.NewSigner(kr, nil, addr[0], encCfg.TxConfig, testfactory.ChainID, 1, 0)
// 	require.NoError(t, err)

// 	rand := tmrand.NewRand()
// 	rand.Seed(1)
// 	encodedTxs1 := blobfactory.GenerateManyRandomRawSendTxsSameSigner(rand, signer, normalTxCount)

// 	signer.ForceSetSequence(0)
// 	rand2 := tmrand.NewRand()
// 	rand2.Seed(1)
// 	encodedTxs2 := blobfactory.GenerateManyRandomRawSendTxsSameSigner(rand2, signer, normalTxCount)

// 	// additional check for the sake of future debugging
// 	for i := 0; i < normalTxCount; i++ {
// 		tx1, err := TxDecoder(encodedTxs1[i])
// 		assert.NoError(t, err)
// 		assert.NotNil(t, tx1)
// 		msgs1 := tx1.GetMsgs()

// 		tx2, err2 := TxDecoder(encodedTxs2[i])
// 		assert.NoError(t, err2)
// 		assert.NotNil(t, tx2)
// 		msgs2 := tx2.GetMsgs()

// 		assert.Equal(t, msgs1, msgs2)
// 		assert.Equal(t, tx1, tx2)
// 	}

// 	assert.Equal(t, encodedTxs1, encodedTxs2)
// }
