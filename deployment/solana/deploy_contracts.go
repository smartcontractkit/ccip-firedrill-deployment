package solana

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	solBinary "github.com/gagliardetto/binary"
	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/token"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/failing_receiver"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_compound"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_feequoter"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_offramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_token"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

var _ deployment.ChangeSetV2[shared.FiredrillConfig] = FiredrillDeployRegisterChangeSet{}

type FiredrillDeployRegisterChangeSet struct{}

func (c FiredrillDeployRegisterChangeSet) Apply(e deployment.Environment, config shared.FiredrillConfig) (deployment.ChangesetOutput, error) {
	changesetOutput, err := DeployAndInitializeFiredrillContracts(e, config)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	err = e.ExistingAddresses.Merge(changesetOutput.AddressBook)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	return changesetOutput, nil
}

func (c FiredrillDeployRegisterChangeSet) VerifyPreconditions(e deployment.Environment, config shared.FiredrillConfig) error {
	if (config.Version == semver.Version{}) {
		return errors.New("missing Version")
	}
	if config.Version != deployment.Version1_6_0 {
		return fmt.Errorf("invalid Version %s. Only 1.6.0 is allowed", config.Version)
	}
	if config.ChainSelector == 0 {
		return errors.New("missing ChainSelector")
	}
	_, ok := e.SolChains[config.ChainSelector]
	if !ok {
		return fmt.Errorf("missing chain for selector %d. If this is first deployment for this chain, "+
			"you need to specify empty address book, i.e. in addresses.json: `\"%d\": {}`", config.ChainSelector, config.ChainSelector)
	}
	if config.SourceChainSelector == 0 {
		return errors.New("missing SourceChainSelector")
	}
	return nil
}

func DeployAndInitializeFiredrillContracts(env deployment.Environment, config shared.FiredrillConfig) (deployment.ChangesetOutput, error) {
	ab := deployment.NewMemoryAddressBook()
	chain := env.SolChains[config.ChainSelector]

	firedrillCompoundProgramID, err := chain.DeployProgram(env.Logger, "firedrill_compound", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillCompoundAddress := solana.MustPublicKeyFromBase58(firedrillCompoundProgramID)
	firedrill_compound.SetProgramID(firedrillCompoundAddress)

	firedrillEntrypointProgramID, err := chain.DeployProgram(env.Logger, "firedrill_entrypoint", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillEntrypointAddress := solana.MustPublicKeyFromBase58(firedrillEntrypointProgramID)
	firedrill_entrypoint.SetProgramID(firedrillEntrypointAddress)

	firedrillOfframpProgramID, err := chain.DeployProgram(env.Logger, "firedrill_offramp", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillOfframpAddress := solana.MustPublicKeyFromBase58(firedrillOfframpProgramID)
	firedrill_offramp.SetProgramID(firedrillOfframpAddress)

	firedrillFeeQuoterProgramID, err := chain.DeployProgram(env.Logger, "firedrill_feequoter", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillFeeQuoterAddress := solana.MustPublicKeyFromBase58(firedrillFeeQuoterProgramID)
	firedrill_feequoter.SetProgramID(firedrillFeeQuoterAddress)

	firedrillTokenProgramID, err := chain.DeployProgram(env.Logger, "firedrill_token", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillTokenAddress := solana.MustPublicKeyFromBase58(firedrillTokenProgramID)
	firedrill_token.SetProgramID(firedrillTokenAddress)

	firedrillReceiverProgramID, err := chain.DeployProgram(env.Logger, "failing_receiver", false)
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to deploy program: %w", err)
	}
	firedrillReceiverAddress := solana.MustPublicKeyFromBase58(firedrillReceiverProgramID)
	failing_receiver.SetProgramID(firedrillReceiverAddress)

	var firedrillTokenAccount token.Mint
	firedrillTokenPDA, _, _ := FindFiredrillTokenPDA(firedrillTokenAddress)
	err = chain.GetAccountDataBorshInto(env.GetContext(), firedrillTokenPDA, &firedrillTokenAccount)
	if err != nil {
		initTx, err2 := firedrill_token.NewInitializeInstruction(
			firedrillTokenPDA,
			chain.DeployerKey.PublicKey(),
			solana.TokenProgramID,
			solana.SystemProgramID,
			solana.SysVarRentPubkey,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build instruction: %w", err2)
		}
		if err2 := chain.Confirm([]solana.Instruction{initTx}); err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm token initialization: %w", err2)
		}
		env.Logger.Infow("FiredrillToken initialized", "chain", chain.String())
	} else {
		env.Logger.Infow("FiredrillToken already initialized, skipping initialization", "chain", chain.String())
	}

	var firedrillCompoundAccount firedrill_compound.FiredrillCompound
	firedrillCompoundPDA, _, _ := FindFiredrillCompoundPDA(firedrillCompoundAddress)
	firedrillCompoundConfigPDA, _, _ := FindFiredrillCompoundConfigPDA(firedrillCompoundAddress)
	firedrillCompoundDestChainPDA, _, _ := FindFiredrillCompoundDestChainPDA(firedrillCompoundAddress, config.ChainSelector)
	err = chain.GetAccountDataBorshInto(env.GetContext(), firedrillCompoundPDA, &firedrillCompoundAccount)
	if err != nil {
		initTx, err2 := firedrill_compound.NewInitializeInstruction(
			config.ChainSelector,
			firedrillFeeQuoterAddress,
			firedrillTokenAddress,
			firedrillCompoundPDA,
			firedrillCompoundConfigPDA,
			firedrillCompoundDestChainPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build instruction: %w", err2)
		}
		if err2 := chain.Confirm([]solana.Instruction{initTx}); err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm compound initialization: %w", err2)
		}
		env.Logger.Infow("FiredrillCompound initialized", "chain", chain.String())
	} else {
		env.Logger.Infow("FiredrillCompound already initialized, skipping initialization", "chain", chain.String())
	}

	var firedrillEntrypointAccount firedrill_entrypoint.FiredrillEntrypoint
	firedrillEntrypointPDA, _, _ := FindFiredrillEntrypointPDA(firedrillEntrypointAddress)
	err = chain.GetAccountDataBorshInto(env.GetContext(), firedrillEntrypointPDA, &firedrillEntrypointAccount)
	if err != nil {
		initTx, err2 := firedrill_entrypoint.NewInitializeInstruction(
			chain.Selector,
			firedrillTokenAddress,
			firedrillOfframpAddress,
			firedrillOfframpAddress,
			firedrillCompoundAddress,
			firedrillReceiverAddress,
			firedrillEntrypointPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build instruction: %w", err2)
		}
		if err2 := chain.Confirm([]solana.Instruction{initTx}); err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm entrypoint initialization: %w", err2)
		}
		env.Logger.Infow("FiredrillEntrypoint initialized", "chain", chain.String())
	} else {
		env.Logger.Infow("FiredrillEntrypoint already initialized, skipping initialization", "chain", chain.String())
	}

	var firedrillOfframpConfigAccount firedrill_offramp.Config
	firedrillOfframpPDA, _, _ := FindFiredrillOfframpPDA(firedrillOfframpAddress)
	firedrillOfframpReferenceAddressesPDA, _, _ := FindFiredrillOfframpReferenceAddressesPDA(firedrillOfframpAddress)
	firedrillOfframpSourceChainPDA, _, _ := FindFiredrillOfframpSourceChainPDA(firedrillOfframpAddress, config.ChainSelector)
	firedrillOfframpConfigPDA, _, _ := FindFiredrillOfframpConfigPDA(firedrillOfframpAddress)
	err = chain.GetAccountDataBorshInto(env.GetContext(), firedrillOfframpConfigPDA, &firedrillOfframpConfigAccount)
	if err != nil {
		programData, err2 := solProgramData(env, chain, firedrillOfframpAddress)
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to get firedrill offramp program data: %w", err)
		}
		initTx, err2 := firedrill_offramp.NewInitializeInstruction(
			chain.Selector,
			firedrillTokenAddress,
			firedrillFeeQuoterAddress,
			firedrillCompoundAddress,
			firedrillOfframpSourceChainPDA,
			firedrillOfframpReferenceAddressesPDA,
			firedrillOfframpPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build init instruction: %w", err2)
		}
		initConfigTx, err2 := firedrill_offramp.NewInitializeConfigInstruction(
			chain.Selector,
			firedrillOfframpConfigPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
			firedrillOfframpAddress,
			programData.Address,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build init config instruction: %w", err2)
		}
		if err2 := chain.Confirm([]solana.Instruction{initTx, initConfigTx}); err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm offramp initialization: %w", err2)
		}
		env.Logger.Infow("FiredrillOfframp initialized", "chain", chain.String())
	} else {
		env.Logger.Infow("FiredrillOfframp already initialized, skipping initialization", "chain", chain.String())
	}
	updateIx, err := firedrill_offramp.NewUpdateOnRampInstruction(
		firedrillOfframpConfigPDA,
		firedrillOfframpSourceChainPDA,
		chain.DeployerKey.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to build update offramp: %w", err)
	}
	if err2 := chain.Confirm([]solana.Instruction{updateIx}); err2 != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm update offramp: %w", err2)
	}

	var firedrillFeeQuoterConfigAccount firedrill_feequoter.Config
	firedrillFeeQuoterPDA, _, _ := FindFiredrillFeeQuoterPDA(firedrillFeeQuoterAddress)
	firedrillFeeQuoterDestChainPDA, _, _ := FindFiredrillFeeQuoterDestChainPDA(firedrillFeeQuoterAddress, config.ChainSelector)
	firedrillFeeQuoterConfigPDA, _, _ := FindFiredrillFeeQuoterConfigPDA(firedrillFeeQuoterAddress)
	err = chain.GetAccountDataBorshInto(env.GetContext(), firedrillFeeQuoterConfigPDA, &firedrillFeeQuoterConfigAccount)
	if err != nil {
		initTx, err2 := firedrill_feequoter.NewInitializeInstruction(
			chain.Selector,
			firedrillCompoundAddress,
			firedrillFeeQuoterPDA,
			firedrillFeeQuoterConfigPDA,
			firedrillFeeQuoterDestChainPDA,
			firedrillTokenPDA,
			chain.DeployerKey.PublicKey(),
			solana.SystemProgramID,
		).ValidateAndBuild()
		if err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to build init instruction: %w", err2)
		}
		if err2 := chain.Confirm([]solana.Instruction{initTx}); err2 != nil {
			return deployment.ChangesetOutput{}, fmt.Errorf("failed to confirm FeeQuoter initialization: %w", err2)
		}
		env.Logger.Infow("FiredrillFeeQuoter initialized", "chain", chain.String())
	} else {
		env.Logger.Infow("FiredrillFeeQuoter already initialized, skipping initialization", "chain", chain.String())
	}

	err = ab.Save(chain.Selector, firedrillEntrypointProgramID, deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deployment.Version1_6_0))
	if err != nil {
		return deployment.ChangesetOutput{}, fmt.Errorf("failed to save program to address book: %w", err)
	}
	return deployment.ChangesetOutput{
		AddressBook: ab,
	}, nil

}

func solProgramData(e deployment.Environment, chain deployment.SolChain, programID solana.PublicKey) (struct {
	DataType uint32
	Address  solana.PublicKey
}, error) {
	var programData struct {
		DataType uint32
		Address  solana.PublicKey
	}
	data, err := chain.Client.GetAccountInfoWithOpts(e.GetContext(), programID, &solRpc.GetAccountInfoOpts{
		Commitment: solRpc.CommitmentConfirmed,
	})
	if err != nil {
		return programData, fmt.Errorf("failed to deploy program: %w", err)
	}

	err = solBinary.UnmarshalBorsh(&programData, data.Bytes())
	if err != nil {
		return programData, fmt.Errorf("failed to unmarshal program data: %w", err)
	}
	return programData, nil
}
