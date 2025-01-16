package keeper

import (
	"context"

	"cosmossdk.io/math"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k Keeper) GetFaultCounter(ctx context.Context, operator sdk.ValAddress) uint64 {
	has, err := k.FaultCounts.Has(ctx, operator)
	if err != nil {
		panic(err)
	}

	if !has {
		return 0
	}

	val, err := k.FaultCounts.Get(ctx, operator)
	if err != nil {
		panic(err)
	}

	return val
}

func (k Keeper) SetFaultCounter(ctx context.Context, operator sdk.ValAddress, faultCounter uint64) {
	err := k.FaultCounts.Set(ctx, operator, faultCounter)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) DeleteFaultCounter(ctx context.Context, operator sdk.ValAddress) {
	err := k.FaultCounts.Remove(ctx, operator)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) IterateFaultCounters(ctx context.Context,
	handler func(operator sdk.ValAddress, faultCount uint64) (stop bool),
) {
	err := k.FaultCounts.Walk(
		ctx,
		nil,
		func(key []byte, value uint64) (bool, error) {
			return handler(key, value), nil
		},
	)
	if err != nil {
		panic(err)
	}
}

func (k Keeper) HandleSlashEpoch(ctx sdk.Context) {
	// TODO: error handling
	params, _ := k.Params.Get(ctx)
	slashFraction := math.LegacyMustNewDecFromStr(params.SlashFraction) // TODO: remove with Dec
	powerReduction := k.StakingKeeper.PowerReduction(ctx)
	k.IterateFaultCounters(ctx, func(operator sdk.ValAddress, faultCount uint64) bool {
		validator, err := k.StakingKeeper.Validator(ctx, operator)
		if err != nil {
			panic(err)
		}

		defer k.DeleteFaultCounter(ctx, operator)
		if validator.IsJailed() || !validator.IsBonded() {
			return false
		}

		if faultCount <= params.EpochMaxFault {
			return false
		}

		consAddr, err := validator.GetConsAddr()
		if err != nil {
			panic(err)
		}

		err = k.SlashingKeeper.Slash(
			ctx, consAddr, slashFraction,
			validator.GetConsensusPower(powerReduction),
			ctx.BlockHeight()-sdk.ValidatorUpdateDelay-1,
		)
		if err != nil {
			panic(err)
		}
		err = k.SlashingKeeper.Jail(ctx, consAddr)
		if err != nil {
			panic(err)
		}
		return false
	})
}
