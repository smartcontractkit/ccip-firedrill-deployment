package solana

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_offramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_onramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPrepareRegister(client deployment.SolChain, view shared.FiredrillEntrypointView) error {
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	firedrill_offramp.SetProgramID(offRampProgram)
	offRampPDA, _, _ := FindFiredrillOnrampPDA(offRampProgram)
	ix, err := firedrill_offramp.NewEmitSourceChainAddedInstruction(
		offRampPDA,
		client.DeployerKey.PublicKey(),
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

func CallDrillPendingCommit(idx uint64, client deployment.SolChain, view shared.FiredrillEntrypointView) error {
	onRampProgram := solana.MustPublicKeyFromBase58(view.OnRamp)
	onRampPDA, _, _ := FindFiredrillOnrampPDA(onRampProgram)
	firedrill_onramp.SetProgramID(onRampProgram)
	ix, err := firedrill_onramp.NewEmitCcipMessageSentInstruction(
		client.DeployerKey.PublicKey(),
		idx,
		onRampPDA,
		client.DeployerKey.PublicKey(),
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

func CallDrillPendingExec(from uint64, to uint64, client deployment.SolChain, view shared.FiredrillEntrypointView) error {
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	firedrill_offramp.SetProgramID(offRampProgram)
	offRampPDA, _, _ := FindFiredrillOnrampPDA(offRampProgram)
	ix, err := firedrill_offramp.NewEmitCommitReportAcceptedInstruction(
		from, to,
		offRampPDA,
		client.DeployerKey.PublicKey(),
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
