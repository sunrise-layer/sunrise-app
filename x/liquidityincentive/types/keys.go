package types

import (
	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

const (
	// ModuleName defines the module name
	ModuleName = "liquidityincentive"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// GovModuleName duplicates the gov module's name to avoid a dependency with x/gov.
	// It should be synced with the gov module's name if it is ever changed.
	// See: https://github.com/cosmos/cosmos-sdk/blob/v0.52.0-beta.2/x/gov/types/keys.go#L9
	GovModuleName = "gov"
)

var (
	// ParamsKey is the prefix to retrieve all Params
	ParamsKey = collections.NewPrefix("params/")

	EpochsKeyPrefix = collections.NewPrefix("epochs/")
	EpochIdKey      = collections.NewPrefix("epoch_id/")
	GaugesKeyPrefix = collections.NewPrefix("gauges/")
	VotesKeyPrefix  = collections.NewPrefix("votes/")
)

var (
	EpochsKeyCodec = collections.Uint64Key
	GaugesKeyCodec = collections.PairKeyCodec(collections.Uint64Key, collections.Uint64Key)
	VotesKeyCodec  = sdk.AccAddressKey
)

func GaugeKey(previousEpochId uint64, poolId uint64) collections.Pair[uint64, uint64] {
	return collections.Join(previousEpochId, poolId)
}
