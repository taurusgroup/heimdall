package cli

import (
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client/context"
	"github.com/cosmos/cosmos-sdk/codec"
	"github.com/maticnetwork/heimdall/helper"
	"github.com/maticnetwork/heimdall/staking"
	"github.com/maticnetwork/heimdall/types"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// send validator join transaction
func GetValidatorJoinTx(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-join",
		Short: "Join Heimdall as a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			validatorID := viper.GetInt64(FlagValidatorID)
			if validatorID == 0 {
				return fmt.Errorf("Validator ID cannot be zero")
			}

			pubkeyStr := viper.GetString(FlagSignerPubkey)
			if pubkeyStr == "" {
				return fmt.Errorf("Pubkey has to be supplied")
			}

			startEpoch := viper.GetInt64(FlagStartEpoch)

			endEpoch := viper.GetInt64(FlagEndEpoch)

			amountStr := viper.GetString(FlagAmount)


			pubkeyBytes, err := hex.DecodeString(pubkeyStr)
			if err != nil {
				return err
			}
			pubkey := types.NewPubKey(pubkeyBytes)

			msg := staking.NewMsgValidatorJoin(uint64(validatorID), pubkey, uint64(startEpoch), uint64(endEpoch), json.Number(amountStr))

			return helper.CreateAndSendTx(msg, cliCtx)
		},
	}

	cmd.Flags().Int(FlagValidatorID, 0, "--id=<validator ID here>")
	cmd.Flags().String(FlagSignerPubkey, "", "--signer-pubkey=<signer pubkey here>")
	cmd.Flags().String(FlagStartEpoch, "0", "--start-epoch=<start epoch of validator here>")
	cmd.Flags().String(FlagEndEpoch, "0", "--end-epoch=<end epoch of validator here>")
	cmd.Flags().String(FlagAmount, "", "--staked-amount=<staked amount>")

	cmd.MarkFlagRequired(FlagSignerPubkey)
	cmd.MarkFlagRequired(FlagStartEpoch)
	cmd.MarkFlagRequired(FlagEndEpoch)
	cmd.MarkFlagRequired(FlagAmount)
	return cmd
}

// send validator exit transaction
func GetValidatorExitTx(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "validator-exit",
		Short: "Exit heimdall as a validator ",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			validator := viper.GetInt64(FlagValidatorID)
			if validator == 0 {
				return fmt.Errorf("validator ID cannot be 0")
			}
			msg := staking.NewMsgValidatorExit(uint64(validator))

			return helper.CreateAndSendTx(msg, cliCtx)
		},
	}

	cmd.Flags().Int(FlagValidatorID, 0, "--id=<validator ID here>")
	return cmd
}

// send validator update transaction
func GetValidatorUpdateTx(cdc *codec.Codec) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "signer-update",
		Short: "Update signer for a validator",
		RunE: func(cmd *cobra.Command, args []string) error {
			cliCtx := context.NewCLIContext().WithCodec(cdc)

			validator := viper.GetInt64(FlagValidatorID)
			if validator == 0 {
				return fmt.Errorf("validator ID cannot be 0")
			}

			pubkeyStr := viper.GetString(FlagNewSignerPubkey)
			if pubkeyStr == "" {
				return fmt.Errorf("Pubkey has to be supplied")
			}

			amountStr := viper.GetString(FlagAmount)

			pubkeyBytes, err := hex.DecodeString(pubkeyStr)
			if err != nil {
				return err
			}
			pubkey := types.NewPubKey(pubkeyBytes)

			msg := staking.NewMsgValidatorUpdate(uint64(validator), pubkey, json.Number(amountStr))

			return helper.CreateAndSendTx(msg, cliCtx)
		},
	}
	cmd.Flags().Int(FlagValidatorID, 0, "--id=<validator ID here>")
	cmd.Flags().String(FlagNewSignerPubkey, "", "--new-pubkey=< new signer pubkey here>")
	cmd.Flags().String(FlagAmount, "", "--staked-amount=<staked amount>")

	cmd.MarkFlagRequired(FlagNewSignerPubkey)
	cmd.MarkFlagRequired(FlagAmount)

	return cmd
}