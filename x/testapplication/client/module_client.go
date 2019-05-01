package client

import (
	"github.com/cosmos/cosmos-sdk/client"
	testapplicationcmd "github.com/Ohmere03/testapplication/x/testapplication/client/cli"
	"github.com/spf13/cobra"
	amino "github.com/tendermint/go-amino"

)

// ModuleClient exports all client functionality from this module
type ModuleClient struct {
	storeKey string
	cdc      *amino.Codec
}

func NewModuleClient(storeKey string, cdc *amino.Codec) ModuleClient {
	return ModuleClient{storeKey, cdc}
}

// GetQueryCmd returns the cli query commands for this module
func (mc ModuleClient) GetQueryCmd() *cobra.Command {
	// Group testapplication queries under a subcommand
	testapplicationQueryCmd := &cobra.Command{
		Use:   "testapplication",
		Short: "Querying commands for the testapplication module",
	}

	testapplicationQueryCmd.AddCommand(client.GetCommands(
		testapplicationcmd.GetCmdResolveHash(mc.storeKey, mc.cdc),
		testapplicationcmd.GetCmdBol(mc.storeKey, mc.cdc),
	)...)

	return testapplicationQueryCmd
}

// GetTxCmd returns the transaction commands for this module
func (mc ModuleClient) GetTxCmd() *cobra.Command {
	testapplicationTxCmd := &cobra.Command{
		Use:   "testapplication",
		Short: "Testapplication transactions subcommands",
	}

	testapplicationTxCmd.AddCommand(client.PostCommands(
		testapplicationcmd.GetCmdCreateBol(mc.cdc),
		testapplicationcmd.GetCmdTransmitBol(mc.cdc),
	)...)

	return testapplicationTxCmd
}