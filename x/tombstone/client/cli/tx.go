package cli

import (
	"bufio"
	"fmt"
	"time"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/x/auth"
	"github.com/cosmos/cosmos-sdk/x/auth/client/utils"
	"github.com/daoleno/higan/x/tombstone/internal/types"
)

// GetTxCmd returns the transaction commands for this module
func GetTxCmd(cdc *codec.Codec) *cobra.Command {
	tombstoneTxCmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	tombstoneTxCmd.AddCommand(flags.PostCommands(
		GetCmdSetRecord(cdc),
	)...)

	return tombstoneTxCmd
}

// GetCmdSetRecord is the CLI command for doing SetRecord
func GetCmdSetRecord(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:     "record [name] [born] [died] [memo]",
		Example: `higancli tx tombstone record "Arthur Charles Clarke" 12/16/1917 03/19/2008 "He never grew up, but he never stopped growing." --tag writer,futurist --from cosmos1r5ur9cmqtanc3uuayvurcmnx9uzy9kt24u275c`,
		Args:    cobra.ExactArgs(4),
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)
			inBuf := bufio.NewReader(cmd.InOrStdin())
			txBldr := auth.NewTxBuilderFromCLI(inBuf).WithTxEncoder(utils.GetTxEncoder(cdc))

			name := args[0]
			born, err := time.Parse(types.LayoutDate, args[1])
			if err != nil {
				return err
			}
			died, err := time.Parse(types.LayoutDate, args[2])
			if err != nil {
				return err
			}
			memo := args[3]
			if len(memo) > types.MemoMaxLength {
				return fmt.Errorf("Memo should less than 140 charactor")
			}
			tags := viper.GetStringSlice(flagTag)
			recorder := cliCtx.GetFromAddress()

			if !cmd.Flags().Changed(flags.FlagFrom) {
				// TODO: Use Public recorder
				return fmt.Errorf("Flag '%s' must be provided", flags.FlagFrom)
			}

			msg := types.NewMsgSetRecord(name, born, died, memo, tags, recorder)
			err = msg.ValidateBasic()
			if err != nil {
				return err
			}

			return utils.GenerateOrBroadcastMsgs(cliCtx, txBldr, []sdk.Msg{msg})
		},
	}

	cmd.Flags().StringSlice(flagTag, nil, "take comma-separated tag")
	return cmd
}
