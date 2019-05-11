package testapplication

import (
	"encoding/json"
	sdk "github.com/cosmos/cosmos-sdk/types"

)

// MsgBuyName defines the BuyName message
type MsgTransmitBol struct {
	Hash string
	Owner  sdk.AccAddress
	NewOwner sdk.AccAddress
}

// NewMsgBuyName is the constructor function for MsgBuyName
func NewMsgTransmitBol(hash string, owner sdk.AccAddress, newOwner sdk.AccAddress) MsgTransmitBol {
	return MsgTransmitBol{
		Hash: hash,
		Owner:owner,
		NewOwner:newOwner,
	}
}

// Route should return the name of the module
func (msg MsgTransmitBol) Route() string { return "testapplication" }

// Type should return the action
func (msg MsgTransmitBol) Type() string { return "transmit_Bol" }

// ValidateBasic runs stateless checks on the message
func (msg MsgTransmitBol) ValidateBasic() sdk.Error {
	if msg.Owner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if msg.NewOwner.Empty() {
		return sdk.ErrInvalidAddress(msg.Owner.String())
	}
	if len(msg.Hash) == 0  {
		return sdk.ErrUnknownRequest("Hash cannot be empty")
	}
	return nil
}

// GetSignBytes encodes the message for signing
func (msg MsgTransmitBol) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgTransmitBol) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Owner}
}


//____________________________________________________________________
//Message CreateBol
//type to define the createBol message

type MsgCreateBol struct {
	Hash string
	Owner sdk.AccAddress

}

//constructor function for MsgBuyName
func NewMsgCreateBol(hash string, owner sdk.AccAddress) MsgCreateBol{
	return MsgCreateBol{
		Hash:hash,
		Owner:owner,
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


//___________________________________________________________________________________________________

//MessageSendMoney

type MsgSendMoney struct {
	Destination sdk.AccAddress
	Amount      sdk.Coins
	Sender      sdk.AccAddress
}

func NewMsgSendMoney( destination sdk.AccAddress, amount sdk.Coins, sender sdk.AccAddress) MsgSendMoney{
	return MsgSendMoney{
		Destination: destination,
		Amount:      amount,
		Sender:      sender,
	}
}

//Route should return the name of the module
func (msg MsgSendMoney) Route() string {return "testapplication"}

//type should return the action

func (msg MsgSendMoney) Type() string {return "sendMoney"}

//ValidateBasic runs stateless checks on the message

func (msg MsgSendMoney) ValidateBasic() sdk.Error{
	if msg.Sender.Empty(){
		return sdk.ErrInvalidAddress(msg.Sender.String())
	}
	if msg.Destination.Empty(){
		return sdk.ErrInvalidAddress(msg.Destination.String())
	}
	/*if msg.Amount.IsAllPositive(){
		return  sdk.ErrInsufficientCoins("Amount must be positive")

	}*/
	return nil
}


// GetSignBytes encodes the message for signing
func (msg MsgSendMoney) GetSignBytes() []byte {
	b, err := json.Marshal(msg)
	if err != nil {
		panic(err)
	}
	return sdk.MustSortJSON(b)
}

// GetSigners defines whose signature is required
func (msg MsgSendMoney) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}