package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"

	"github.com/sentinel-official/hub/x/subscription/types"
)

func (k Keeper) CancelDuration(ctx sdk.Context) (duration time.Duration) {
	k.params.Get(ctx, types.KeyCancelDuration, &duration)
	return
}

func (k Keeper) SetParams(ctx sdk.Context, params types.Params) {
	k.params.SetParamSet(ctx, &params)
}

func (k Keeper) GetParams(ctx sdk.Context) types.Params {
	return types.NewParams(
		k.CancelDuration(ctx),
	)
}
