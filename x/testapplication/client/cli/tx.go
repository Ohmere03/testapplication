package cli


import (
	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/utils"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/Ohmere03/testapplication/x/testapplication"

	sdk "github.com/cosmos/cosmos-sdk/types"
	authtxb "github.com/cosmos/cosmos-sdk/x/auth/client/txbuilder"

)

// GetCmdCreateBol is the CLI command for sending a CreateBol transaction
func GetCmdCreateBol(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "create-bol [hash] [newowner]",
		Short: "create new bol",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}


			msg := testapplication.NewMsgCreateBol(args[0], cliCtx.GetFromAddress())

			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}

// GetCmdSTransmitBol is the CLI command for sending a TransmitBol transaction
func GetCmdTransmitBol(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "transmit-bol [bol] [newowner]",
		Short: "Set a new owner of the bill of lading",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			address, _ := sdk.AccAddressFromBech32(args[1])
			msg := testapplication.NewMsgTransmitBol(args[0],cliCtx.GetFromAddress(),address)
			err := msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}


// GetCmdSendMoney moves coins from one account to another
func GetCmdSendMoney(cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "send-money [destination] [amount]",
		Short: "Send money from one account to another",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc).WithAccountDecoder(cdc)

			txBldr := authtxb.NewTxBuilderFromCLI().WithTxEncoder(utils.GetTxEncoder(cdc))

			if err := cliCtx.EnsureAccountExists(); err != nil {
				return err
			}
			address, _ := sdk.AccAddressFromBech32(args[0])
			coins, err := sdk.ParseCoins(args[1])
			if err != nil {
				return err
			}
			msg := testapplication.NewMsgSendMoney(address, coins, cliCtx.GetFromAddress())
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			cliCtx.PrintResponse = true

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg}, false)
		},
	}
}
