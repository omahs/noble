package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/cctp module sentinel errors
var (
	ErrUnauthorized             = sdkerrors.Register(ModuleName, 2, "unauthorized")
	ErrUserNotFound             = sdkerrors.Register(ModuleName, 3, "user not found")
	ErrMint                     = sdkerrors.Register(ModuleName, 4, "tokens can not be minted")
	ErrSendCoinsToAccount       = sdkerrors.Register(ModuleName, 5, "can't send tokens to account")
	ErrBurn                     = sdkerrors.Register(ModuleName, 6, "tokens can not be burned")
	ErrPaused                   = sdkerrors.Register(ModuleName, 7, "the chain is paused")
	ErrMintingDenomSet          = sdkerrors.Register(ModuleName, 9, "the minting denom has already been set")
	ErrUserBlacklisted          = sdkerrors.Register(ModuleName, 10, "user is already blacklisted")
	ErrAlreadyPrivileged        = sdkerrors.Register(ModuleName, 11, "address is already assigned to privileged role")
	ErrDenomNotRegistered       = sdkerrors.Register(ModuleName, 12, "denom not registered in bank module")
	ErrAttesterAlreadyFound     = sdkerrors.Register(ModuleName, 13, "attester is already present")
	ErrAttesterNotFound         = sdkerrors.Register(ModuleName, 14, "attester not found")
	ErrAuthorityNotSet          = sdkerrors.Register(ModuleName, 15, "authority not set")
	ErrMalformedField           = sdkerrors.Register(ModuleName, 16, "field cannot be empty or nil")
	ErrReceiveMessage           = sdkerrors.Register(ModuleName, 17, "err in receive message")
	ErrDisableAttester          = sdkerrors.Register(ModuleName, 18, "err in disable attester")
	ErrUpdateSignatureThreshold = sdkerrors.Register(ModuleName, 19, "err in update signature threshold")
	ErrMinterAllowanceNotFound  = sdkerrors.Register(ModuleName, 20, "minter allowance not found")
	ErrTokenPairAlreadyFound    = sdkerrors.Register(ModuleName, 21, "token pair already exists")
	ErrTokenPairNotFound        = sdkerrors.Register(ModuleName, 22, "token pair not found")
	ErrSendMessage              = sdkerrors.Register(ModuleName, 23, "error in send message")
	ErrSendMessageWithCaller    = sdkerrors.Register(ModuleName, 24, "error in send message with caller")
	ErrDepositForBurn           = sdkerrors.Register(ModuleName, 25, "error in deposit for burn")
	ErrInvalidDestinationCaller = sdkerrors.Register(ModuleName, 26, "malformed destination caller")
	ErrSignatureVerification    = sdkerrors.Register(ModuleName, 27, "unable to verify signature")
	ErrReplaceMessage           = sdkerrors.Register(ModuleName, 28, "error in replace message")
	ErrDuringPause              = sdkerrors.Register(ModuleName, 28, "error while trying to pause or unpause")
)