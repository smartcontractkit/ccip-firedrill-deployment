package solana

import (
	"fmt"

	"github.com/gagliardetto/solana-go"
	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPendingCommit(from uint8, to uint8, client *deployment.SolChain, view shared.FiredrillEntrypointView) error {
	entrypointPDA, _, _ := FindFiredrillEntrypointPDA(solana.MustPublicKeyFromBase58(view.Address))
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
		return fmt.Errorf("error building drill entrypoint instruction: %w", err)
	}
	err = client.Confirm([]solana.Instruction{ix})
	if err != nil {
		return fmt.Errorf("error confirming drill entrypoint instruction: %w", err)
	}
	return nil
}
