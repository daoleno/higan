package cli

import (
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"

	"github.com/daoleno/higan/x/tombstone/internal/types"
)

// GetQueryCmd returns the cli query commands for this module
func GetQueryCmd(queryRoute string, cdc *codec.Codec) *cobra.Command {
	// Group tombstone queries under a subcommand
	tombstoneQueryCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tombstoneQueryCmd.AddCommand(
		flags.GetCommands(
		// TODO: Add query Cmds
		)...,
	)

	return tombstoneQueryCmd
}

func getCmdRecord(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:   "record [recorder]",
		Short: "record recorder",
		Args:  cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			recorder := args[0]

			res, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", queryRoute, recorder))
			if err != nil {
				fmt.Printf("could not query recorder - %s \n", recorder)
				fmt.Println(err)
				return nil
			}

			var out types.QueryRecordList
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

}
