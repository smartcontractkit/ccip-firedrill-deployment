package deployment

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/deployment"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_6/gethwrappers/firedrill_entrypoint"
)

const FiredrillEntrypointType deployment.ContractType = "FiredrillEntrypoint"

type FiredrillEntrypoint interface {
	Owner(opts *bind.CallOpts) (common.Address, error)
	IsActive(opts *bind.CallOpts) (bool, error)
	Token(opts *bind.CallOpts) (common.Address, error)
	OnRamp(opts *bind.CallOpts) (common.Address, error)
	OffRamp(opts *bind.CallOpts) (common.Address, error)
	Receiver(opts *bind.CallOpts) (common.Address, error)
	Compound(opts *bind.CallOpts) (common.Address, error)
}

var _ deployment.ChangeSetV2[FiredrillConfig] = FiredrillDeployRegisterChangeSet{}

type FiredrillConfig struct {
	Version             semver.Version
	ChainSelector       uint64
	SourceChainSelector uint64
}

type FiredrillDeployRegisterChangeSet struct{}

func (c FiredrillDeployRegisterChangeSet) Apply(e deployment.Environment, config FiredrillConfig) (deployment.ChangesetOutput, error) {
	changesetOutput, err := DeployFiredrillContracts(e, config)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	err = e.ExistingAddresses.Merge(changesetOutput.AddressBook)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	err = FiredrillRegisterContracts(e.Logger, changesetOutput.AddressBook, e.Chains[config.ChainSelector])
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	return changesetOutput, nil
}

func (c FiredrillDeployRegisterChangeSet) VerifyPreconditions(e deployment.Environment, config FiredrillConfig) error {
	if (config.Version == semver.Version{}) {
		return errors.New("missing Version")
	}
	if config.Version != deployment.Version1_5_0 && config.Version != deployment.Version1_6_0 {
		return fmt.Errorf("invalid Firedrill version %s. Only 1.5.0 and 1.6.0 is allowed", config.Version)
	}
	if config.ChainSelector == 0 {
		return errors.New("missing ChainSelector")
	}
	if config.SourceChainSelector == 0 {
		return errors.New("missing SourceChainSelector")
	}
	return nil
}

func DeployFiredrillContracts(e deployment.Environment, config FiredrillConfig) (deployment.ChangesetOutput, error) {
	ab := deployment.NewMemoryAddressBook()
	switch config.Version {
	case deployment.Version1_5_0:
		_, err := deployment.DeployContract(e.Logger, e.Chains[config.ChainSelector], ab, deployFiredrillEntrypointV1_5)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
	case deployment.Version1_6_0:
		_, err := deployment.DeployContract(e.Logger, e.Chains[config.ChainSelector], ab, deployFiredrillEntrypoint)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
	default:
		return deployment.ChangesetOutput{}, fmt.Errorf("unknown version %s", config.Version.String())
	}
	return deployment.ChangesetOutput{
		AddressBook: ab,
	}, nil
}

func deployFiredrillEntrypointV1_5(chain deployment.Chain) deployment.ContractDeploy[*firedrill_entrypoint_v1_5.FiredrillEntrypoint] {
	address, tx, inst, err := firedrill_entrypoint_v1_5.DeployFiredrillEntrypoint(chain.DeployerKey, chain.Client, chain.Selector)
	return deployment.ContractDeploy[*firedrill_entrypoint_v1_5.FiredrillEntrypoint]{
		Address:  address,
		Contract: inst,
		Tx:       tx,
		Tv:       deployment.NewTypeAndVersion(FiredrillEntrypointType, deployment.Version1_5_0),
		Err:      err,
	}
}

func deployFiredrillEntrypoint(chain deployment.Chain) deployment.ContractDeploy[*firedrill_entrypoint.FiredrillEntrypoint] {
	address, tx, inst, err := firedrill_entrypoint.DeployFiredrillEntrypoint(chain.DeployerKey, chain.Client, chain.Selector)
	return deployment.ContractDeploy[*firedrill_entrypoint.FiredrillEntrypoint]{
		Address:  address,
		Contract: inst,
		Tx:       tx,
		Tv:       deployment.NewTypeAndVersion(FiredrillEntrypointType, deployment.Version1_6_0),
		Err:      err,
	}
}

func FiredrillRegisterContracts(lggr logger.Logger, addressBook deployment.AddressBook, chain deployment.Chain) error {
	firedrillEntrypointAddr, err := deployment.SearchAddressBook(addressBook, chain.Selector, FiredrillEntrypointType)
	if err != nil {
		return err
	}
	firedrillEntrypoint, err := firedrill_entrypoint.NewFiredrillEntrypoint(common.HexToAddress(firedrillEntrypointAddr), chain.Client)
	if err != nil {
		return err
	}
	entrypointTokenAddr, err1 := firedrillEntrypoint.Token(nil)
	entrypointOnRampAddr, err2 := firedrillEntrypoint.OnRamp(nil)
	entrypointOffRampAddr, err3 := firedrillEntrypoint.OffRamp(nil)
	entrypointCompoundAddr, err4 := firedrillEntrypoint.Compound(nil)
	entrypointReceiverAddr, err5 := firedrillEntrypoint.Receiver(nil)
	if err := errors.Join(err1, err2, err3, err4, err5); err != nil {
		return err
	}
	if (entrypointTokenAddr == common.Address{}) {
		err = errors.New("FiredrillToken address from address book does not match FiredrillEntrypoint.Token()")
	}
	if (entrypointOnRampAddr == common.Address{}) {
		err = errors.Join(err, errors.New("firedrill onramp on entrypoint can't be 0-address"))
	}
	if (entrypointOffRampAddr == common.Address{}) {
		err = errors.Join(err, errors.New("firedrill offramp on entrypoint can't be 0-address"))
	}
	if (entrypointCompoundAddr == common.Address{}) {
		err = errors.Join(err, errors.New("firedrill compound on entrypoint can't be 0-address"))
	}
	if (entrypointReceiverAddr == common.Address{}) {
		err = errors.Join(err, errors.New("firedrill receiver on entrypoint can't be 0-address"))
	}
	if err != nil {
		return err
	}
	tx, err := firedrillEntrypoint.PrepareRegister(chain.DeployerKey)
	if err != nil {
		return err
	}
	blockNum, err := chain.Confirm(tx)
	if err != nil {
		return err
	}
	lggr.Info("FiredrillEntrypoint register events mined", "address", firedrillEntrypointAddr, "tx", tx.Hash(), "blockNum", blockNum)
	return nil
}
