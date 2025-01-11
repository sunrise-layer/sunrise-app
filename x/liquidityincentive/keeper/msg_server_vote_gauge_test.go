package keeper_test

import (
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
	"go.uber.org/mock/gomock"

	"github.com/sunriselayer/sunrise/x/liquidityincentive/keeper"
	"github.com/sunriselayer/sunrise/x/liquidityincentive/types"
	liquiditypooltypes "github.com/sunriselayer/sunrise/x/liquiditypool/types"
)

func TestMsgVoteGauge(t *testing.T) {
	sender := "sunrise126ss57ayztn5287spvxq0dpdfarj6rk0v3p06f"
	k, mocks, ms, ctx := setupMsgServer(t)
	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))
	wctx := sdk.UnwrapSDKContext(ctx)

	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(1)).Return(liquiditypooltypes.Pool{}, true).AnyTimes()
	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(2)).Return(liquiditypooltypes.Pool{}, true).AnyTimes()
	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(3)).Return(liquiditypooltypes.Pool{}, false).AnyTimes()

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgVoteGauge
		expErr    bool
		expErrMsg string
	}{
		{
			name: "not available pool",
			input: &types.MsgVoteGauge{
				Sender:      sender,
				PoolWeights: []types.PoolWeight{{PoolId: 3, Weight: "1"}},
			},
			expErr:    true,
			expErrMsg: "pool not found",
		},
		{
			name: "available pools",
			input: &types.MsgVoteGauge{
				Sender:      sender,
				PoolWeights: []types.PoolWeight{{PoolId: 1, Weight: "0.5"}, {PoolId: 1, Weight: "0.5"}},
			},
			expErr: false,
		},
		{
			name: "empty votes",
			input: &types.MsgVoteGauge{
				Sender:      sender,
				PoolWeights: []types.PoolWeight{},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.VoteGauge(wctx, tc.input)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestMsgVoteGaugePartial(t *testing.T) {
	sender := "sunrise126ss57ayztn5287spvxq0dpdfarj6rk0v3p06f"
	f := initFixture(t)
	ctx := sdk.UnwrapSDKContext(f.ctx)
	k := f.keeper
	ms := keeper.NewMsgServerImpl(k)
	mocks := getMocks(t)

	params := types.DefaultParams()
	require.NoError(t, k.Params.Set(ctx, params))
	wctx := sdk.UnwrapSDKContext(ctx)

	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(1)).Return(liquiditypooltypes.Pool{}, true).AnyTimes()
	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(2)).Return(liquiditypooltypes.Pool{}, true).AnyTimes()
	mocks.LiquiditypoolKeeper.EXPECT().GetPool(gomock.Any(), uint64(3)).Return(liquiditypooltypes.Pool{}, true).AnyTimes()

	// default params
	testCases := []struct {
		name      string
		input     *types.MsgVoteGauge
		expErr    bool
		expErrMsg string
	}{

		{
			name: "all available pools",
			input: &types.MsgVoteGauge{
				Sender:      sender,
				PoolWeights: []types.PoolWeight{{PoolId: 1, Weight: "0.5"}, {PoolId: 2, Weight: "0.5"}, {PoolId: 3, Weight: "0.5"}},
			},
			expErr: false,
		},
		{
			name: "partial pool",
			input: &types.MsgVoteGauge{
				Sender:      sender,
				PoolWeights: []types.PoolWeight{{PoolId: 2, Weight: "0.5"}},
			},
			expErr: false,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			_, err := ms.VoteGauge(wctx, tc.input)
			if tc.expErr {
				require.Error(t, err)
				require.Contains(t, err.Error(), tc.expErrMsg)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
