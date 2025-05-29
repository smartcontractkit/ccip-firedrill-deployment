package solana

import (
	"context"
	"fmt"

	"github.com/gagliardetto/solana-go"
	solRpc "github.com/gagliardetto/solana-go/rpc"
	solCommonUtil "github.com/smartcontractkit/chainlink-ccip/chains/solana/utils/common"
	cldf_solana "github.com/smartcontractkit/chainlink-deployments-framework/chain/solana"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_offramp"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPrepareRegister(client cldf_solana.Chain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
	offRampProgram := solana.MustPublicKeyFromBase58(view.OffRamp)
	firedrill_offramp.SetProgramID(offRampProgram)
	offRampPDA, _, _ := FindFiredrillOfframpPDA(offRampProgram)
	offRampSourceChainPDA, _, _ := FindFiredrillOfframpSourceChainPDA(offRampProgram, client.Selector)
	ix, err := firedrill_offramp.NewEmitSourceChainAddedInstruction(
		offRampPDA,
		client.DeployerKey.PublicKey(),
		offRampSourceChainPDA,
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

func CallDrillPendingCommit(idx uint64, client cldf_solana.Chain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
	entrypointProgram := solana.MustPublicKeyFromBase58(view.Address)
	entrypointPDA, _, _ := FindFiredrillEntrypointPDA(entrypointProgram)
	firedrill_entrypoint.SetProgramID(entrypointProgram)
	ix, err := firedrill_entrypoint.NewEmitCcipMessageSentInstruction(
		client.DeployerKey.PublicKey(),
		idx,
		entrypointPDA,
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

func CallDrillPendingExec(from uint64, to uint64, client cldf_solana.Chain, view shared.FiredrillEntrypointView) (*solana.Transaction, error) {
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

func Confirm(chain cldf_solana.Chain, ix solana.Instruction) (*solana.Transaction, error) {
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
