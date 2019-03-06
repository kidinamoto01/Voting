package ballot

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"encoding/json"
	"google.golang.org/genproto/googleapis/type/date"
)
// MsgSetVote defines a SetVote message
type MsgSetBallot struct {
	Name string
	deadline date.Date
	Info  string
	Owner  sdk.AccAddress
}

// NewMsgSetName is a constructor function for MsgSetBallot
func NewMsgSetBallot(name string, date date.Date,value string, owner sdk.AccAddress) MsgSetBallot {
	return MsgSetBallot{
		Name: name,
		deadline:date,
		Info:  value,
		Owner:  owner,
	}
}
// Type should return the name of the module
func (msg MsgSetBallot) Route() string { return "nameservice" }

// Name should return the action
func (msg MsgSetBallot) Type() string { return "set_name"}

// ValdateBasic Implements Msg.
func (msg MsgSetBallot) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Name) == 0 || len(msg.Info) == 0 {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetBallot) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetBallot) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}


// MsgSetVote defines a SetVote message
type MsgSetVote struct {
	Name string
	Vote  bool
	From  sdk.AccAddress
}


// NewMsgSetName is a constructor function for MsgSetBallot
func NewMsgSetVote(name string, value bool, from sdk.AccAddress) MsgSetVote {
	return MsgSetVote{
		Name: name,
		Vote:  value,
		From:  from,
	}
}
// Type should return the name of the module
func (msg MsgSetVote) Route() string { return "nameservice" }

// Name should return the action
func (msg MsgSetVote) Type() string { return "set_name"}

// ValdateBasic Implements Msg.
func (msg MsgSetVote) ValidateBasic() sdk.Error {
	if msg.From.Empty() {
		return sdk.ErrInvalidAddress(msg.From.String())
	}
	if len(msg.Name) == 0  {
		return sdk.ErrUnknownRequest("Name and/or Value cannot be empty")
	}
	return nil
}

// GetSignBytes Implements Msg.
func (msg MsgSetVote) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners Implements Msg.
func (msg MsgSetVote) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.From}
}