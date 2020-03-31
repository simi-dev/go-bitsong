package ibc_desmos

// nolint
// autogenerated code using github.com/haasted/alias-generator.
// based on functionality in github.com/rigelrozanski/multitool

import (
	"github.com/bitsongofficial/go-bitsong/x/ibc_desmos/keeper"
	"github.com/bitsongofficial/go-bitsong/x/ibc_desmos/types"
)

const (
	ModuleName            = types.ModuleName
	StoreKey              = types.StoreKey
	RouterKey             = types.RouterKey
	QuerierRoute          = types.QuerierRoute
	DefaultPacketTimeout  = keeper.DefaultPacketTimeout
	DesmosBitsongSubspace = keeper.DesmosBitsongSubspace
)

var (
	// functions aliases
	NewKeeper            = keeper.NewKeeper
	RegisterCodec        = types.RegisterCodec
	NewMsgCreateSongPost = types.NewMsgCreateSongPost

	// variable aliases
	ModuleCdc = types.ModuleCdc
)

type (
	Keeper            = keeper.Keeper
	MsgCreateSongPost = types.MsgCreateSongPost
)
