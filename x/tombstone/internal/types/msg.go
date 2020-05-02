package types

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// verify interface at compile time
var _ sdk.Msg = &MsgSetRecord{}

// MsgSetRecord - struct for people be remembered
type MsgSetRecord struct {
	Name string    `json:"name"`
	Born time.Time `json:"born"`
	Died time.Time `json:"died"`
	Memo string    `json:"memo"`
	Tags []string  `json:"tags"`

	Recorder sdk.AccAddress `json:"recorder"`
}

// NewMsgSetRecord creates a new MsgSetRecord instance
func NewMsgSetRecord(name string, born, died time.Time, memo string, tags []string, recorder sdk.AccAddress) MsgSetRecord {
	return MsgSetRecord{
		Name:     name,
		Born:     born,
		Died:     died,
		Memo:     memo,
		Tags:     tags,
		Recorder: recorder,
	}
}

const RecordConst = "Record"

// nolint
func (msg MsgSetRecord) Route() string { return RouterKey }
func (msg MsgSetRecord) Type() string  { return RecordConst }
func (msg MsgSetRecord) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{sdk.AccAddress(msg.Recorder)}
}

// GetSignBytes gets the bytes for the message signer to sign on
func (msg MsgSetRecord) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic validity check for the AnteHandler
func (msg MsgSetRecord) ValidateBasic() error {
	if msg.Recorder.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "missing recorder address")
	}
	return nil
}
