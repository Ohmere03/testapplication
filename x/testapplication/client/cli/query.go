package cli


import (
	"fmt"

	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/Ohmere03/testapplication/x/testapplication"
	"github.com/spf13/cobra"

)


// GetCmdResolveHash queries information about a hash

func GetCmdResolveHash(queryRoute string, cdc *codec.Codec) *cobra.Command{
	return &cobra.Command{
		Use: "resolve [hash]",
		Short: "resolve hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hash := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/resolve/%s", queryRoute, hash), nil)
			if err != nil {
				fmt.Printf("could not resolve hash- %s \n", string(hash))
				return nil
			}

			var out testapplication.QueryResResolve
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},

	}

}
// GetCmdBol queries information about a domain
func GetCmdBol(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "bol [hash]",
		Short: "Query bol info of hash",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			hash := args[0]

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/whois/%s", queryRoute, hash), nil)
			if err != nil {
				fmt.Printf("could not resolve whois - %s \n", string(hash))
				return nil
			}

			var out testapplication.Bol
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}
// GetCmdHashes queries a list of all hashes
func GetCmdHashes(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "hashes",
		Short: "hashes",
		// Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, err := cliCtx.QueryWithData(fmt.Sprintf("custom/%s/hashes", queryRoute), nil)
			if err != nil {
				fmt.Printf("could not get query hashes\n")
				return nil
			}

			var out testapplication.QueryResHashes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}
}



