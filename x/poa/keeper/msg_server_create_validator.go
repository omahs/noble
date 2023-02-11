package keeper

import (
	"context"
	"fmt"

	"github.com/strangelove-ventures/noble/x/poa/types"

	sdkerrors "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateValidator(goCtx context.Context, msg *types.MsgCreateValidator) (*types.MsgCreateValidatorResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	valAddr, err := sdk.AccAddressFromBech32(msg.Address)
	if err != nil {
		return nil, fmt.Errorf("failed to decode address as bech32: %w", err)
	}

	if _, found := k.GetValidator(ctx, valAddr); found {
		return nil, sdkerrors.Wrap(types.ErrBadValidatorAddr, fmt.Sprintf("unrecognized %s validator already exists: %T", types.ModuleName, msg))
	}

	validator := &types.Validator{
		Description: msg.Description,
		Address:     valAddr,
		Pubkey:      msg.Pubkey,
	}

	k.SaveValidator(ctx, validator)

	// ctx.EventManager().EmitEvents(sdk.Events{
	// 	sdk.NewEvent(
	// 		stakingtypes.EventTypeCreateValidator,
	// 		sdk.NewAttribute(stakingtypes.AttributeKeyValidator, msg.Address.String()),
	// 		sdk.NewAttribute(sdk.AttributeKeySender, msg.Owner.String()),
	// 	),
	// })

	err = ctx.EventManager().EmitTypedEvent(msg)

	return &types.MsgCreateValidatorResponse{}, err
}