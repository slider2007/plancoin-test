package plancoin

import (
	"fmt"

	"time"

	"github.com/Shushsa/plan/x/bank"
	"github.com/Shushsa/plan/x/emission"
	"github.com/Shushsa/plan/x/plancoin/types"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"github.com/cosmos/cosmos-sdk/x/params"
)

// NewHandler creates an sdk.Handler for all the plancoin type messages
func NewHandler(k Keeper, paramsKeeper params.Keeper, bankKeeper bank.Keeper, emissionKeeper emission.Keeper) sdk.Handler {
	return func(ctx sdk.Context, msg sdk.Msg) (*sdk.Result, error) {
		ctx = ctx.WithEventManager(sdk.NewEventManager())
		switch msg := msg.(type) {
		case types.MsgChangeParams:
			return handleChangeParams(ctx, k, msg, paramsKeeper)
		default:
			errMsg := fmt.Sprintf("unrecognized %s message type: %T", ModuleName, msg)
			return nil, sdkerrors.Wrap(sdkerrors.ErrUnknownRequest, errMsg)
		}
	}
}

func handleChangeParams(ctx sdk.Context, k Keeper, msg types.MsgChangeParams, paramsKeeper params.Keeper) (*sdk.Result, error) {
	if !msg.Owner.Equals(types.GetGenesisWallet()) {
		return nil, sdkerrors.Wrapf(params.ErrSettingParameter, "only genesis can call this method")
	}

	ss, ok := paramsKeeper.GetSubspace("staking")

	if !ok {
		return nil, sdkerrors.Wrap(params.ErrUnknownSubspace, "staking")
	}

	var NewValue time.Duration = time.Hour * 24 * 3

	bin, _ := codec.Cdc.MarshalJSON(NewValue)

	if err := ss.Update(ctx, []byte("UnbondingTime"), bin); err != nil {
		fmt.Println(err)

		return nil, sdkerrors.Wrapf(params.ErrSettingParameter, "key: %s, value: %s, err: %s", "unbonding_time", "", err.Error())
	}

	return &sdk.Result{Events: ctx.EventManager().Events()}, nil
}
