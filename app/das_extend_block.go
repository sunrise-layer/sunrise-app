package app

// import (
// 	"github.com/sunriselayer/sunrise/pkg/appconsts"
// 	"github.com/sunriselayer/sunrise/pkg/da"
// 	"github.com/sunriselayer/sunrise/pkg/shares"
// 	"github.com/sunriselayer/sunrise/pkg/square"

// 	"github.com/celestiaorg/rsmt2d"
// 	coretypes "github.com/cometbft/cometbft/types"
// )

// // ExtendBlock extends the given block data into a data square for a given app
// // version.
// func ExtendBlock(data coretypes.Data, appVersion uint64) (*rsmt2d.ExtendedDataSquare, error) {
// 	// Construct the data square from the block's transactions
// 	dataSquare, err := square.Construct(data.Txs.ToSliceOfBytes(), appVersion, appconsts.SquareSizeUpperBound(appVersion))
// 	if err != nil {
// 		return nil, err
// 	}

// 	return da.ExtendShares(shares.ToBytes(dataSquare))
// }

// // EmptyBlock returns true if the given block data is considered empty by the
// // application at a given version.
// func IsEmptyBlock(data coretypes.Data, _ uint64) bool {
// 	return len(data.Txs) == 0
// }
