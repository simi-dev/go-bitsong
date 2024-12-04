package cmd

import (
	"bufio"
	"fmt"
	"os"

	"cosmossdk.io/math"
	tmcfg "github.com/cometbft/cometbft/config"
	tmjson "github.com/cometbft/cometbft/libs/json"
	tmos "github.com/cometbft/cometbft/libs/os"
	tmtypes "github.com/cometbft/cometbft/types"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/crypto/keyring"
	"github.com/cosmos/cosmos-sdk/server"
	sdk "github.com/cosmos/cosmos-sdk/types"
	banktypes "github.com/cosmos/cosmos-sdk/x/bank/types"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	genutiltypes "github.com/cosmos/cosmos-sdk/x/genutil/types"
	minttypes "github.com/cosmos/cosmos-sdk/x/mint/types"
	slashingtypes "github.com/cosmos/cosmos-sdk/x/slashing/types"
	stakingtypes "github.com/cosmos/cosmos-sdk/x/staking/types"
	"github.com/spf13/cobra"
)

const (
	FlagDelegator = "delegator"
)

var ()

// InitFromStateCmd returns a command that initializes all files needed for Tendermint
// and the respective application.
func VerifySlashedDelegatorsV018(defaultNodeHome string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "v018 [block-height] state_exported.json ",
		Short: "verifies those impacted by distribution module bug",
		Long:  `<add-example-cli>`,
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx := client.GetClientContextFromCmd(cmd)
			serverCtx := server.GetServerContextFromCmd(cmd)
			config := serverCtx.Config

			config.SetRoot(clientCtx.HomeDir)

			blockHeight := args[0]

			stateFile := args[1]

			// attempt to lookup address from Keybase if no address was provided
			kb, err := keyring.New(sdk.KeyringServiceName(), "test", clientCtx.HomeDir, bufio.NewReader(cmd.InOrStdin()), clientCtx.Codec)
			if err != nil {
				return fmt.Errorf("failed to open keyring: %w", err)
			}
			clientCtx.Keyring = kb

			genParams := V018StateExportParams{
				BlockHeight: math.NewUintFromString(blockHeight),
				StateFile:   stateFile,
				// DelegatorAddr:   delegatorAddr,
				// SecondStateFile: secondaryStateFile,
			}

			return rewardVerification(config, clientCtx, genParams)
		},
	}

	return cmd
}

func rewardVerification(_ *tmcfg.Config, cliCtx client.Context, genParams V018StateExportParams) error {

	_, err := V018ConvertStateExport(cliCtx, genParams)
	if err != nil {
		return fmt.Errorf("failed to convert state export: %w", err)
	}

	fmt.Println("Veification Complete")

	return nil
}

type V018StateExportParams struct {
	// DelegatorAddr   string
	StateFile   string
	BlockHeight math.Uint
}

func (s *V018StateExportParams) String() string {
	return fmt.Sprintf(`Block Height: %s
  State File: %s
`, s.BlockHeight, s.StateFile)
}

func (s *V018StateExportParams) Validate() error {
	if s.BlockHeight.LTE(math.NewUint(19818775)) { // v18 upgrade -1
		return fmt.Errorf("block height cannot be less than v18 upgrade")
	}

	if s.StateFile == "" {
		return fmt.Errorf("state file cannot be empty")
	}

	return nil
}

func V018ConvertStateExport(clientCtx client.Context, params V018StateExportParams) (*tmtypes.GenesisDoc, error) {
	if !tmos.FileExists(params.StateFile) {
		return nil, fmt.Errorf("%s does not exist", params.StateFile)
	}

	// print state export params
	fmt.Println(params.String())

	err := params.Validate()
	if err != nil {
		return nil, err
	}

	stateBz, err := os.ReadFile(params.StateFile)
	if err != nil {
		return nil, fmt.Errorf("couldn't read state export file: %w", err)
	}
	genDoc := tmtypes.GenesisDoc{}
	err = tmjson.Unmarshal(stateBz, &genDoc)
	if err != nil {
		return nil, fmt.Errorf("couldn't unmarshal state export file: %w", err)
	}

	if err := genDoc.ValidateAndComplete(); err != nil {
		return nil, err
	}

	appState, err := genutiltypes.GenesisStateFromGenDoc(genDoc)
	if err != nil {
		return nil, err
	}

	var staking stakingtypes.GenesisState
	var distribution distrtypes.GenesisState
	var slashing slashingtypes.GenesisState
	var bank banktypes.GenesisState
	var mint minttypes.GenesisState
	clientCtx.Codec.MustUnmarshalJSON(appState[stakingtypes.ModuleName], &staking)
	clientCtx.Codec.MustUnmarshalJSON(appState[distrtypes.ModuleName], &distribution)
	clientCtx.Codec.MustUnmarshalJSON(appState[banktypes.ModuleName], &bank)
	clientCtx.Codec.MustUnmarshalJSON(appState[minttypes.ModuleName], &mint)
	clientCtx.Codec.MustUnmarshalJSON(appState[slashingtypes.ModuleName], &slashing)

	// Map to store delegations with slash actions after the upgrade
	// Key: Validator Address, Value: Map of Delegator Addresses to their Delegation Amounts
	delegations := staking.Delegations
	delegationsWithSlashActions := make(map[string]map[string]math.Int)
	// GOAL: get delegators for each validator with a slashing action in state
	for _, vse := range distribution.ValidatorSlashEvents {

		// Initialize the validator's delegations map if not already done
		if _, ok := delegationsWithSlashActions[vse.ValidatorAddress]; !ok {
			delegationsWithSlashActions[vse.ValidatorAddress] = make(map[string]math.Int)
		}
		// Now, iterate over delegations to find those related to the validator with a slash action
		for _, delegation := range delegations {
			if delegation.ValidatorAddress == vse.ValidatorAddress {
				fmt.Println("VALIDATOR SLASH ACTION:", vse)
				delegationsWithSlashActions[vse.ValidatorAddress][delegation.DelegatorAddress] = math.Int(delegation.Shares)
			}
		}
	}

	return &genDoc, nil
}
