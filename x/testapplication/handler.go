package testapplication

import (

	"fmt"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


// NewHandler returns a handler for "testapplication" type messages.

func NewHandler(keeper Keeper) sdk.Handler{
	return func(ctx sdk.Context, msg sdk.Msg) sdk.Result{
		switch msg := msg.(type){
		case MsgTransmitBol:
			return  handleMsgTransmitBol(ctx, keeper, msg)

		case MsgCreateBol:
			return handleMsgCreateBol(ctx,keeper,msg)
		case MsgSendMoney:
			return  handleMsgSendMoney(ctx, keeper,msg)
		default:
			errMsg := fmt.Sprintf("Unrecognized testapplication Msg type: %v", msg.Type())
			return sdk.ErrUnknownRequest(errMsg).Result()
		}
	}
}



// func for the MsgTransmitBol

func handleMsgTransmitBol(ctx sdk.Context, keeper Keeper, msg MsgTransmitBol) sdk.Result {
	if !msg.Owner.Equals(keeper.GetOwner(ctx, msg.Hash)) { // Checks if the the msg sender is the same as the current owner
		return sdk.ErrUnauthorized("Incorrect Owner").Result() // If not, throw an error
	}
	keeper.SetOwner(ctx, msg.Hash, msg.NewOwner) // If so, set the owner to the newOwner specified in the msg.
	return sdk.Result{}                      // return
}


func handleMsgCreateBol(ctx sdk.Context, keeper Keeper, msg MsgCreateBol) sdk.Result {
	bol := NewBol(msg.Owner, msg.Hash)
	keeper.SetBol(ctx,msg.Hash,bol)
	return sdk.Result{}

}

//func for MsgSendMoney


func handleMsgSendMoney(ctx sdk.Context, keeper Keeper, msg MsgSendMoney) sdk.Result {
	_, err := keeper.coinKeeper.SendCoins(ctx, msg.Sender, msg.Destination, msg.Amount)
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}

	/*_, err = k.coinKeeper.SubtractCoins(ctx, msg.Sender, msg.Amount)
	if err != nil {
		return sdk.ErrInsufficientCoins("Buyer does not have enough coins").Result()
	}*/
	return sdk.Result{}


}
