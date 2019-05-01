package testapplication

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

// MsgTransmitBol defines a TransmitBol message

type MsgTransmitBoll struct {
	Hash string
	Owner sdk.AccAddress
	NewOwner sdk.AccAddress

}

// NewMsgTransmitBol is a constructor function for MsgTransmitBol

func NewMsgTransmitBol(hash string, owner sdk.AccAddress, newOwner sdk.AccAddress) MsgTransmitBoll{
	return MsgTransmitBoll{
		Hash:hash,
		Owner:owner,
		NewOwner: newOwner,
	}
}

//Implementation of the Msg interface


//Route should return the name of the module

func (msg MsgTransmitBoll) Route() string {return "testapplication"}


//Type should return the action

func (msg MsgTransmitBoll) Type() string {return "Transmit_bol"}

//ValidateBasic run stateless checks on the message

func (msg MsgTransmitBoll) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Hash) == 0 || len(msg.NewOwner) == 0 {
		return sdk.ErrUnknownRequest("Hash/new Owner and/or Retrieve address cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransmitBoll) GetSignBytes() []byte{
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required

func (msg MsgTransmitBoll) GetSigners() []sdk.AccAddress{
	return []sdk.AccAddress{msg.Owner}
}

//____________________________________________________________________
//Message CreateBol
//type to define the createBol message

type MsgCreateBol struct {
	Hash string
	Owner sdk.AccAddress
	Retrieve string

}

//constructor function for MsgBuyName
func NewMsgCreateBol(hash string, owner sdk.AccAddress, retrieve string) MsgCreateBol{
	return MsgCreateBol{
		Hash:hash,
		Owner:owner,
		Retrieve:retrieve,
	}
}

//Route should return the name of the module

func (msg MsgCreateBol) Route()  string{return "testapplication"}

//Type should return the action

func (msg MsgCreateBol) Type() string {return "create_bol"}

//Validate Basic runs stateless checks on the message

func (msg MsgCreateBol) ValidateBasic() sdk.Error{
	if msg.Owner.Empty(){
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Hash) == 0 {
		return sdk.ErrUnknownRequest("Hash cannot be empty")
	}
	if len(msg.Retrieve) == 0 {
		return sdk.ErrUnknownRequest("Retrieve indication cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgCreateBol) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgCreateBol) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}


