package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	solCommonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	"github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_compound"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_offramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPrepareRegister(client deployment.SolChain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	firedrill_offramp.SetProgramID(offRampProgram)
	offRampPDA, _, _ := FindFiredrillOfframpPDA(offRampProgram)
	ix, err := firedrill_offramp.NewEmitSourceChainAddedInstruction(
		offRampPDA,
		client.DeployerKey.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("error building prepare-register instruction: %w", err)
	}
	tx, err := Confirm(client, ix)
	if err != nil {
		return nil, fmt.Errorf("error confirming prepare-register instruction: %w", err)
	}
	return tx, nil
}

func CallDrillPendingCommit(idx uint64, client deployment.SolChain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
	compoundProgram := solana.MustPublicKeyFromBase58(view.OnRamp)
	compoundPDA, _, _ := FindFiredrillCompoundPDA(compoundProgram)
	firedrill_compound.SetProgramID(compoundProgram)
	ix, err := firedrill_compound.NewEmitCcipMessageSentInstruction(
		client.DeployerKey.PublicKey(),
		idx,
		compoundPDA,
		client.DeployerKey.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("error building pending-commit instruction: %w", err)
	}
	tx, err := Confirm(client, ix)
	if err != nil {
		return nil, fmt.Errorf("error confirming pending-commit instruction: %w", err)
	}
	return tx, nil
}

func CallDrillPendingExec(from uint64, to uint64, client deployment.SolChain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	firedrill_offramp.SetProgramID(offRampProgram)
	offRampPDA, _, _ := FindFiredrillOfframpPDA(offRampProgram)
	ix, err := firedrill_offramp.NewEmitCommitReportAcceptedInstruction(
		from, to,
		offRampPDA,
		client.DeployerKey.PublicKey(),
	).ValidateAndBuild()
	if err != nil {
		return nil, fmt.Errorf("error building pending-exec instruction: %w", err)
	}
	tx, err := Confirm(client, ix)
	if err != nil {
		return nil, fmt.Errorf("error confirming pending-exec instruction: %w", err)
	}
	return tx, nil
}

func Confirm(chain deployment.SolChain, ix solana.Instruction) (*solana.Transaction, error) {
	txReceipt, err := solCommonUtil.SendAndConfirm(
		context.Background(), chain.Client, []solana.Instruction{ix}, *chain.DeployerKey, solRpc.CommitmentConfirmed,
	)
	if err != nil {
		return nil, fmt.Errorf("error confirming instruction: %w", err)
	}
	tx, err := txReceipt.Transaction.GetTransaction()
	if err != nil {
		return nil, fmt.Errorf("error getting transaction: %w", err)
	}
	return tx, nil
}
