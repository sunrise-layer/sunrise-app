package keeper

import (
	"context"

	"cosmossdk.io/store/prefix"
	"github.com/cosmos/cosmos-sdk/runtime"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/sunriselayer/sunrise/x/liquiditypool/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) WrapPoolInfo(ctx context.Context, pool types.Pool) types.PoolInfo {
	tokenBase := k.bankKeeper.GetBalance(ctx, pool.GetAddress(), pool.DenomBase)
	tokenQuote := k.bankKeeper.GetBalance(ctx, pool.GetAddress(), pool.DenomQuote)
	return types.PoolInfo{
		Pool:       pool,
		TokenBase:  tokenBase,
		TokenQuote: tokenQuote,
	}
}

func (q queryServer) Pools(ctx context.Context, req *types.QueryPoolsRequest) (*types.QueryPoolsResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var pools []types.PoolInfo

	store := runtime.KVStoreAdapter(q.k.KVStoreService.OpenKVStore(ctx))
	poolStore := prefix.NewStore(store, types.KeyPrefix(types.PoolKey))

	pageRes, err := query.Paginate(poolStore, req.Pagination, func(key []byte, value []byte) error {
		var pool types.Pool
		if err := q.k.cdc.Unmarshal(value, &pool); err != nil {
			return err
		}

		pools = append(pools, q.k.WrapPoolInfo(ctx, pool))
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryPoolsResponse{Pools: pools, Pagination: pageRes}, nil
}

func (q queryServer) Pool(ctx context.Context, req *types.QueryPoolRequest) (*types.QueryPoolResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	pool, found := q.k.GetPool(ctx, req.Id)
	if !found {
		return nil, sdkerrors.ErrKeyNotFound
	}

	return &types.QueryPoolResponse{Pool: q.k.WrapPoolInfo(ctx, pool)}, nil
}
