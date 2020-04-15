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
			getCmdRecord(queryRoute, cdc),
			getCmdList(queryRoute, cdc),
		)...,
	)

	return tombstoneQueryCmd
}

func getCmdRecord(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "record [recorder]",
		Short:   "query record through recorder",
		Example: "higancli query tombstone record cosmos1lxmp6c3229lqy8xuv6tfzjd8fwd8q8yqp443hh",
		Args:    cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			recorder := args[0]

			res, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s/%s", queryRoute, types.QueryRecord, recorder))
			if err != nil {
				fmt.Printf("could not query recorder - %s \n", recorder)
				fmt.Println(err)
				return nil
			}

			var out types.QueryRecordRes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

}

func getCmdList(queryRoute string, cdc *codec.Codec) *cobra.Command {
	listCommand := &cobra.Command{
		Use:                        "list",
		Short:                      "list recorder or records",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	listCommand.AddCommand(
		flags.GetCommands(
			getCmdListRecorder(queryRoute, cdc),
			getCmdListRecord(queryRoute, cdc),
		)...,
	)

	return listCommand
}

func getCmdListRecorder(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "recorder",
		Short:   "list all recorder",
		Example: "higancli query tombstone list recorder",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryAllRecorder))
			if err != nil {
				fmt.Println(err)
				return nil
			}

			var out types.QueryRecorderRes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

}

func getCmdListRecord(queryRoute string, cdc *codec.Codec) *cobra.Command {
	return &cobra.Command{
		Use:     "record",
		Short:   "list all records",
		Example: "higancli query tombstone list record",
		Args:    cobra.ExactArgs(0),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			res, _, err := cliCtx.Query(fmt.Sprintf("custom/%s/%s", queryRoute, types.QueryAllRecord))
			if err != nil {
				fmt.Println(err)
				return nil
			}

			var out types.QueryAllNoteRes
			cdc.MustUnmarshalJSON(res, &out)
			return cliCtx.PrintOutput(out)
		},
	}

}
