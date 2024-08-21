package e2e

// import (
// 	"context"
// 	"errors"
// 	"os"
// 	"testing"
// 	"time"

// 	"github.com/stretchr/testify/require"
// 	"github.com/sunriselayer/sunrise/app/encoding"
// 	v1 "github.com/sunriselayer/sunrise/pkg/appconsts/v1"
// 	"github.com/sunriselayer/sunrise/test/txsim"
// 	testencoding "github.com/sunriselayer/sunrise/test/util/encoding"
// 	"github.com/sunriselayer/sunrise/test/util/testnode"
// )

// const seed = 42

// var latestVersion = "latest"

// // This test runs a simple testnet with 4 validators. It submits both MsgPayForBlobs
// // and MsgSends over 30 seconds and then asserts that at least 10 transactions were
// // committed.
// func TestE2ESimple(t *testing.T) {
// 	if os.Getenv("E2E") != "true" {
// 		t.Skip("skipping e2e test")
// 	}

// 	if os.Getenv("E2E_LATEST_VERSION") != "" {
// 		latestVersion = os.Getenv("E2E_LATEST_VERSION")
// 		_, isSemVer := ParseVersion(latestVersion)
// 		switch {
// 		case isSemVer:
// 		case latestVersion == "latest":
// 		case len(latestVersion) == 7:
// 		case len(latestVersion) == 8:
// 			// assume this is a git commit hash (we need to trim the last digit to match the docker image tag)
// 			latestVersion = latestVersion[:7]
// 		default:
// 			t.Fatalf("unrecognised version: %s", latestVersion)
// 		}
// 	}
// 	t.Log("Running simple e2e test", "version", latestVersion)

// 	testnet, err := New(t.Name(), seed)
// 	require.NoError(t, err)
// 	t.Cleanup(testnet.Cleanup)
// 	require.NoError(t, testnet.CreateGenesisNodes(4, latestVersion, 10000000, 0))

// 	kr, err := testnet.CreateAccount("alice", 1e12)
// 	require.NoError(t, err)

// 	require.NoError(t, testnet.Setup())
// 	require.NoError(t, testnet.Start())

// 	sequences := txsim.NewBlobSequence(txsim.NewRange(200, 4000), txsim.NewRange(1, 3)).Clone(5)
// 	sequences = append(sequences, txsim.NewSendSequence(4, 1000, 100).Clone(5)...)

// 	encCfg := encoding.MakeConfig(testencoding.ModuleEncodingRegisters...)
// 	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
// 	defer cancel()
// 	opts := txsim.DefaultOptions().WithSeed(seed).SuppressLogs()
// 	err = txsim.Run(ctx, testnet.GRPCEndpoints()[0], kr, encCfg, opts, sequences...)
// 	require.True(t, errors.Is(err, context.DeadlineExceeded), err.Error())

// 	blockchain, err := testnode.ReadBlockchain(context.Background(), testnet.Node(0).AddressRPC())
// 	require.NoError(t, err)

// 	totalTxs := 0
// 	for _, block := range blockchain {
// 		require.Equal(t, v1.Version, block.Version.App)
// 		totalTxs += len(block.Data.Txs)
// 	}
// 	require.Greater(t, totalTxs, 10)
// }
