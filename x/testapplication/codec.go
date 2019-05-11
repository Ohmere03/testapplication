package testapplication

import (
	"github.com/cosmos/cosmos-sdk/codec"
)


// RegisterCodec registers concrete types on wire codec
func RegisterCodec(cdc *codec.Codec) {
	cdc.RegisterConcrete(MsgTransmitBol{}, "testapplication/TransmitBol", nil)
	cdc.RegisterConcrete(MsgCreateBol{}, "testapplication/CreateBol", nil)
	cdc.RegisterConcrete(MsgSendMoney{}, "testapplication/SendMoney", nil)

}