package keeper

import (
	"context"

	"cosmossdk.io/collections"
	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sunriselayer/sunrise/x/da/types"
)

func (k Keeper) GetPublishedData(ctx context.Context, metadataUri string) (data types.PublishedData) {
	has, err := k.PublishedData.Has(ctx, metadataUri)
	if err != nil {
		panic(err)
	}

	if !has {
		return data
	}

	val, err := k.PublishedData.Get(ctx, metadataUri)
	if err != nil {
		panic(err)
	}

	return val
}

// SetParams set the params
func (k Keeper) SetPublishedData(ctx context.Context, data types.PublishedData) error {
	err := k.PublishedData.Set(ctx, data.MetadataUri, data)
	if err != nil {
		return err
	}

	return nil
}

func (k Keeper) DeletePublishedData(ctx sdk.Context, data types.PublishedData) {
	err := k.PublishedData.Remove(ctx, data.MetadataUri)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) GetAllPublishedData(ctx context.Context) (list []types.PublishedData) {
	err := k.PublishedData.Walk(
		ctx,
		nil,
		func(key string, value types.PublishedData) (bool, error) {
			list = append(list, value)
			return false, nil
		},
	)
	if err != nil {
		panic(err)
	}

	return
}

func (k Keeper) GetUnverifiedDataBeforeTime(ctx sdk.Context, timestamp int64) (list []types.PublishedData, err error) {
	err = k.PublishedData.Indexes.StatusTime.Walk(
		ctx,
		collections.NewPrefixedPairRange[collections.Pair[string, int64], string](
			collections.PairPrefix[string, int64]("unverified"),
		),
		func(key collections.Pair[string, int64], metadataUri string) (bool, error) {
			if key.K2() > timestamp {
				return true, nil
			}
			list = append(list, k.GetPublishedData(ctx, metadataUri))
			return false, nil
		},
	)

	return
}
