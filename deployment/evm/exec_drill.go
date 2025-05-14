package evm

import (
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deploy "github.com/smartcontractkit/chainlink/deployment"

	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	firedrill_entrypoint_v1_6 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_6/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func CallDrillPrepareRegister(client deployment.Chain, view shared.FiredrillEntrypointView) (*types.Transaction, error) {
	typeAndVersion := deployment.MustTypeAndVersionFromString(view.TypeAndVersion)
	switch typeAndVersion.Version {
	case deploy.Version1_5_0:
		entrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.PrepareRegister(client.DeployerKey)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	case deploy.Version1_6_0:
		entrypoint, err := firedrill_entrypoint_v1_6.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.PrepareRegister(client.DeployerKey)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	}
	return nil, fmt.Errorf("unsupported version %s", typeAndVersion.Version)
}

func CallDrillPendingCommit(from uint8, to uint8, client deployment.Chain, view shared.FiredrillEntrypointView) (*types.Transaction, error) {
	typeAndVersion := deployment.MustTypeAndVersionFromString(view.TypeAndVersion)
	switch typeAndVersion.Version {
	case deploy.Version1_5_0:
		entrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPendingCommitPendingQueueTxSpike(nil, from, to)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	case deploy.Version1_6_0:
		entrypoint, err := firedrill_entrypoint_v1_6.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPendingCommitPendingQueueTxSpike(nil, from, to)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	}
	return nil, fmt.Errorf("unsupported version %s", typeAndVersion.Version)
}

func CallDrillPendingExec(from uint8, to uint8, client deployment.Chain, view shared.FiredrillEntrypointView) (*types.Transaction, error) {
	typeAndVersion := deployment.MustTypeAndVersionFromString(view.TypeAndVersion)
	switch typeAndVersion.Version {
	case deploy.Version1_5_0:
		entrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPendingExecution(nil, from, to)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	case deploy.Version1_6_0:
		entrypoint, err := firedrill_entrypoint_v1_6.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPendingExecution(nil, from, to)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	}
	return nil, fmt.Errorf("unsupported version %s", typeAndVersion.Version)
}

func CallDrillPendingBless(from uint8, to uint8, client deployment.Chain, view shared.FiredrillEntrypointView) (*types.Transaction, error) {
	typeAndVersion := deployment.MustTypeAndVersionFromString(view.TypeAndVersion)
	switch typeAndVersion.Version {
	case deploy.Version1_5_0:
		entrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPendingExecution(nil, from, to)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	}
	return nil, fmt.Errorf("unsupported version %s", typeAndVersion.Version)
}

func CallDrillPriceRegistry(client deployment.Chain, view shared.FiredrillEntrypointView) (*types.Transaction, error) {
	typeAndVersion := deployment.MustTypeAndVersionFromString(view.TypeAndVersion)
	switch typeAndVersion.Version {
	case deploy.Version1_5_0:
		entrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPriceRegistries(nil)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	case deploy.Version1_6_0:
		entrypoint, err := firedrill_entrypoint_v1_6.NewFiredrillEntrypoint(common.HexToAddress(view.Address), client.Client)
		if err != nil {
			return nil, fmt.Errorf("can't instantiate FiredrillEntrypoint: %w", err)
		}
		tx, err := entrypoint.DrillPriceRegistries(nil)
		_, err = deployment.ConfirmIfNoError(client, tx, err)
		if err != nil {
			return nil, fmt.Errorf("can't confirm PrepareRegister: %w", err)
		}
		return tx, nil
	}
	return nil, fmt.Errorf("unsupported version %s", typeAndVersion.Version)
}
