package types

import (
	nfttypes "github.com/bitsongofficial/go-bitsong/x/nft/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	authtypes "github.com/cosmos/cosmos-sdk/x/auth/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
)

// BankKeeper defines the expected bank keeper (noalias)
type BankKeeper interface {
	MintCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error
	BurnCoins(ctx sdk.Context, moduleName string, amt sdk.Coins) error

	GetSupply(ctx sdk.Context, denom string) sdk.Coin
	GetBalance(ctx sdk.Context, addr sdk.AccAddress, denom string) sdk.Coin

	SendCoins(ctx sdk.Context, fromAddr sdk.AccAddress, toAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromModuleToAccount(ctx sdk.Context, senderModule string, recipientAddr sdk.AccAddress, amt sdk.Coins) error
	SendCoinsFromAccountToModule(ctx sdk.Context, senderAddr sdk.AccAddress, recipientModule string, amt sdk.Coins) error

	SpendableCoins(ctx sdk.Context, addr sdk.AccAddress) sdk.Coins

	SetDenomMetaData(ctx sdk.Context, denomMetaData banktypes.Metadata)
	GetDenomMetaData(ctx sdk.Context, denom string) (banktypes.Metadata, bool)
}

// AccountKeeper defines the expected account keeper
type AccountKeeper interface {
	GetAccount(ctx sdk.Context, addr sdk.AccAddress) authtypes.AccountI
	GetModuleAddress(name string) sdk.AccAddress
	GetModuleAccount(ctx sdk.Context, name string) authtypes.ModuleAccountI
}

type NftKeeper interface {
	GetNFTById(ctx sdk.Context, id uint64) (nfttypes.NFT, error)
	GetMetadataById(ctx sdk.Context, id uint64) (nfttypes.Metadata, error)
	TransferNFT(ctx sdk.Context, msg *nfttypes.MsgTransferNFT) error
	UpdateMetadataAuthority(ctx sdk.Context, msg *nfttypes.MsgUpdateMetadataAuthority) error
	SetPrimarySaleHappened(ctx sdk.Context, metadataId uint64) error
	PrintEdition(ctx sdk.Context, msg *nfttypes.MsgPrintEdition) (uint64, error)
}
