package solana

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPrepareRegister(client *deployment.SolChain, view shared.FiredrillEntrypointView) error {
	entrypointAddress := solana.MustPublicKeyFromBase58(view.Address)
	firedrill_entrypoint.SetProgramID(entrypointAddress)
	entrypointPDA, _, _ := FindFiredrillEntrypointPDA(entrypointAddress)
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	offRampPDA, _, _ := FindFiredrillOnrampPDA(offRampProgram)
	ix, err := firedrill_entrypoint.NewPrepareRegisterInstruction(
		entrypointPDA,
		offRampPDA,
		client.DeployerKey.PublicKey(),
		offRampProgram,
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("error building prepare register instruction: %w", err)
	}
	err = client.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("error confirming prepare register instruction: %w", err)
	}
	return nil
}

func CallDrillPendingCommit(from uint8, to uint8, client *deployment.SolChain, view shared.FiredrillEntrypointView) error {
	entrypointAddress := solana.MustPublicKeyFromBase58(view.Address)
	firedrill_entrypoint.SetProgramID(entrypointAddress)
	entrypointPDA, _, _ := FindFiredrillEntrypointPDA(entrypointAddress)
	onRampProgram := solana.MustPublicKeyFromBase58(view.OnRamp)
	onRampPDA, _, _ := FindFiredrillOnrampPDA(onRampProgram)
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	offRampPDA, _, _ := FindFiredrillOnrampPDA(offRampProgram)
	compoundProgram := solana.MustPublicKeyFromBase58(view.Compound)
	compoundPDA, _, _ := FindFiredrillCompoundPDA(compoundProgram)
	ix, err := firedrill_entrypoint.NewDrillPendingCommitQueueTxSpikeInstruction(
		from, to,
		entrypointPDA,
		client.DeployerKey.PublicKey(),
		onRampPDA,
		offRampPDA,
		compoundPDA,
		onRampProgram,
		offRampProgram,
		compoundProgram,
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("error building drill pending commit instruction: %w", err)
	}
	err = client.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("error confirming drill pending commit instruction: %w", err)
	}
	return nil
}

func CallDrillPendingExec(from uint8, to uint8, client *deployment.SolChain, view shared.FiredrillEntrypointView) error {
	entrypointAddress := solana.MustPublicKeyFromBase58(view.Address)
	firedrill_entrypoint.SetProgramID(entrypointAddress)
	entrypointPDA, _, _ := FindFiredrillEntrypointPDA(entrypointAddress)
	onRampProgram := solana.MustPublicKeyFromBase58(view.OnRamp)
	onRampPDA, _, _ := FindFiredrillOnrampPDA(onRampProgram)
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	offRampPDA, _, _ := FindFiredrillOnrampPDA(offRampProgram)
	compoundProgram := solana.MustPublicKeyFromBase58(view.Compound)
	compoundPDA, _, _ := FindFiredrillCompoundPDA(compoundProgram)
	ix, err := firedrill_entrypoint.NewDrillPendingExecutionInstruction(
		from, to,
		entrypointPDA,
		client.DeployerKey.PublicKey(),
		onRampPDA,
		offRampPDA,
		compoundPDA,
		onRampProgram,
		offRampProgram,
		compoundProgram,
	).ValidateAndBuild()
	if err != nil {
		return fmt.Errorf("error building drill pending exec instruction: %w", err)
	}
	err = client.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("error confirming drill pending exec instruction: %w", err)
	}
	return nil
}
