package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const (
	TypeMsgCreateNFT                 = "create_nft"
	TypeMsgPrintEdition              = "print_edition"
	TypeMsgTransferNFT               = "transfer_nft"
	TypeMsgSignMetadata              = "sign_metadata"
	TypeMsgUpdateMetadata            = "update_metadata"
	TypeMsgUpdateMetadataAuthority   = "update_metadata_authority"
	TypeMsgCreateCollection          = "create_collection"
	TypeMsgVerifyCollection          = "verify_collection"
	TypeMsgUnverifyCollection        = "unverify_collection"
	TypeMsgUpdateCollectionAuthority = "update_collection_authority"
)

var _ sdk.Msg = &MsgCreateNFT{}

func NewMsgCreateNFT(sender sdk.AccAddress, updateAuthority string, data Data, presaleHappened, isMutable bool, masterEditionMaxSupply uint64) *MsgCreateNFT {
	return &MsgCreateNFT{
		Sender: sender.String(),
		Metadata: Metadata{
			UpdateAuthority:     updateAuthority,
			Mint:                sender.String(),
			Data:                &data,
			PrimarySaleHappened: presaleHappened,
			IsMutable:           isMutable,
			MasterEdition: &MasterEdition{
				MaxSupply: masterEditionMaxSupply,
			},
		},
	}
}

func (msg MsgCreateNFT) Route() string { return RouterKey }

func (msg MsgCreateNFT) Type() string { return TypeMsgCreateNFT }

func (msg MsgCreateNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Metadata.Data.SellerFeeBasisPoints > 100 {
		return ErrInvalidSellerFeeBasisPoints
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateNFT) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCreateNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgPrintEdition{}

func NewMsgPrintEdition(sender sdk.AccAddress, metadataId uint64) *MsgPrintEdition {
	return &MsgPrintEdition{
		Sender:     sender.String(),
		MetadataId: metadataId,
	}
}

func (msg MsgPrintEdition) Route() string { return RouterKey }

func (msg MsgPrintEdition) Type() string { return TypeMsgPrintEdition }

func (msg MsgPrintEdition) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgPrintEdition) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgPrintEdition) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgTransferNFT{}

func NewMsgTransferNFT(sender sdk.AccAddress, nftId uint64, newOwner string) *MsgTransferNFT {
	return &MsgTransferNFT{
		Sender:   sender.String(),
		Id:       nftId,
		NewOwner: newOwner,
	}
}

func (msg MsgTransferNFT) Route() string { return RouterKey }

func (msg MsgTransferNFT) Type() string { return TypeMsgTransferNFT }

func (msg MsgTransferNFT) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgTransferNFT) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgTransferNFT) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgSignMetadata{}

func NewMsgSignMetadata(sender sdk.AccAddress, metadataId uint64) *MsgSignMetadata {
	return &MsgSignMetadata{
		Sender:     sender.String(),
		MetadataId: metadataId,
	}
}

func (msg MsgSignMetadata) Route() string { return RouterKey }

func (msg MsgSignMetadata) Type() string { return TypeMsgSignMetadata }

func (msg MsgSignMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSignMetadata) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSignMetadata) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgUpdateMetadata{}

func NewMsgUpdateMetadata(sender sdk.AccAddress, metadataId uint64, data *Data) *MsgUpdateMetadata {
	return &MsgUpdateMetadata{
		Sender:     sender.String(),
		MetadataId: metadataId,
		Data:       data,
	}
}

func (msg MsgUpdateMetadata) Route() string { return RouterKey }

func (msg MsgUpdateMetadata) Type() string { return TypeMsgUpdateMetadata }

func (msg MsgUpdateMetadata) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if msg.Data.SellerFeeBasisPoints > 100 {
		return ErrInvalidSellerFeeBasisPoints
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateMetadata) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgUpdateMetadata) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgUpdateMetadataAuthority{}

func NewMsgUpdateMetadataAuthority(sender sdk.AccAddress, metadataId uint64, newAuthority string) *MsgUpdateMetadataAuthority {
	return &MsgUpdateMetadataAuthority{
		Sender:       sender.String(),
		MetadataId:   metadataId,
		NewAuthority: newAuthority,
	}
}

func (msg MsgUpdateMetadataAuthority) Route() string { return RouterKey }

func (msg MsgUpdateMetadataAuthority) Type() string { return TypeMsgUpdateMetadataAuthority }

func (msg MsgUpdateMetadataAuthority) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateMetadataAuthority) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgUpdateMetadataAuthority) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgCreateCollection{}

func NewMsgCreateCollection(sender sdk.AccAddress, name, uri, updateAuthority string) *MsgCreateCollection {
	return &MsgCreateCollection{
		Sender:          sender.String(),
		Name:            name,
		Uri:             uri,
		UpdateAuthority: updateAuthority,
	}
}

func (msg MsgCreateCollection) Route() string { return RouterKey }

func (msg MsgCreateCollection) Type() string { return TypeMsgCreateCollection }

func (msg MsgCreateCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgCreateCollection) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgCreateCollection) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgVerifyCollection{}

func NewMsgVerifyCollection(sender sdk.AccAddress, collectionId, nftId uint64) *MsgVerifyCollection {
	return &MsgVerifyCollection{
		Sender:       sender.String(),
		CollectionId: collectionId,
		NftId:        nftId,
	}
}

func (msg MsgVerifyCollection) Route() string { return RouterKey }

func (msg MsgVerifyCollection) Type() string { return TypeMsgVerifyCollection }

func (msg MsgVerifyCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgVerifyCollection) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgVerifyCollection) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgUnverifyCollection{}

func NewMsgUnverifyCollection(sender sdk.AccAddress, collectionId, nftId uint64) *MsgUnverifyCollection {
	return &MsgUnverifyCollection{
		Sender:       sender.String(),
		CollectionId: collectionId,
		NftId:        nftId,
	}
}

func (msg MsgUnverifyCollection) Route() string { return RouterKey }

func (msg MsgUnverifyCollection) Type() string { return TypeMsgUnverifyCollection }

func (msg MsgUnverifyCollection) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUnverifyCollection) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgUnverifyCollection) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

var _ sdk.Msg = &MsgUpdateCollectionAuthority{}

func NewMsgUpdateCollectionAuthority(sender sdk.AccAddress, collectionId uint64, newAuthority string) *MsgUpdateCollectionAuthority {
	return &MsgUpdateCollectionAuthority{
		Sender:       sender.String(),
		CollectionId: collectionId,
		NewAuthority: newAuthority,
	}
}

func (msg MsgUpdateCollectionAuthority) Route() string { return RouterKey }

func (msg MsgUpdateCollectionAuthority) Type() string { return TypeMsgUpdateCollectionAuthority }

func (msg MsgUpdateCollectionAuthority) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgUpdateCollectionAuthority) GetSignBytes() []byte {
	b, err := ModuleCdc.MarshalJSON(&msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgUpdateCollectionAuthority) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}
