package testapplication

import (
	sdk "github.com/cosmos/cosmos-sdk/types"

)

//Bol is a strcut containig the metadata of a hash
type Bol struct {
	Value string 		`json:"value"`
	Owner sdk.AccAddress `json:"owner"`
}

// create a new Bol with the all the values initialized

func NewBol(address sdk.AccAddress, value string) Bol{
	return Bol{
		Value: value,
		Owner:address,

	}

}