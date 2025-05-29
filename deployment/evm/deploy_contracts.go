package evm

import (
	"errors"
	"fmt"

	"github.com/Masterminds/semver/v3"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	cldf_evm "github.com/smartcontractkit/chainlink-deployments-framework/chain/evm"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deploy "github.com/smartcontractkit/chainlink/deployment"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_6/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

type FiredrillEntrypoint interface {
	Owner(opts *bind.CallOpts) (common.Address, error)
	Token(opts *bind.CallOpts) (common.Address, error)
	OnRamp(opts *bind.CallOpts) (common.Address, error)
	OffRamp(opts *bind.CallOpts) (common.Address, error)
	Receiver(opts *bind.CallOpts) (common.Address, error)
}

var _ deployment.ChangeSetV2[shared.FiredrillConfig] = FiredrillDeployRegisterChangeSet{}

type FiredrillDeployRegisterChangeSet struct{}

func (c FiredrillDeployRegisterChangeSet) Apply(e deployment.Environment, config shared.FiredrillConfig) (deployment.ChangesetOutput, error) {
	changesetOutput, err := DeployFiredrillContracts(e, config)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	err = e.ExistingAddresses.Merge(changesetOutput.AddressBook)
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	err = FiredrillRegisterContracts(e.Logger, changesetOutput.AddressBook, e.BlockChains.EVMChains()[config.ChainSelector])
	if err != nil {
		return deployment.ChangesetOutput{}, err
	}
	return changesetOutput, nil
}

func (c FiredrillDeployRegisterChangeSet) VerifyPreconditions(e deployment.Environment, config shared.FiredrillConfig) error {
	if (config.Version == semver.Version{}) {
		return errors.New("missing Version")
	}
	if config.Version != deploy.Version1_5_0 && config.Version != deploy.Version1_6_0 {
		return fmt.Errorf("invalid Firedrill version %s. Only 1.5.0 and 1.6.0 is allowed", config.Version)
	}
	if config.ChainSelector == 0 {
		return errors.New("missing ChainSelector")
	}
	_, ok := e.BlockChains.EVMChains()[config.ChainSelector]
	if !ok {
		return fmt.Errorf("missing chain for selector %d. If this is first deployment for this chain, "+
			"you need to specify empty address book, i.e. in addresses.json: `\"%d\": {}`", config.ChainSelector, config.ChainSelector)
	}
	if config.SourceChainSelector == 0 {
		return errors.New("missing SourceChainSelector")
	}
	return nil
}

func DeployFiredrillContracts(e deployment.Environment, config shared.FiredrillConfig) (deployment.ChangesetOutput, error) {
	ab := deployment.NewMemoryAddressBook()
	evmChains := e.BlockChains.EVMChains()
	switch config.Version {
	case deploy.Version1_5_0:
		_, err := deployment.DeployContract(e.Logger, evmChains[config.ChainSelector], ab, deployFiredrillEntrypointV1_5)
		if err != nil {
			return deployment.ChangesetOutput{}, err
		}
	case deploy.Version1_6_0:
		_, err := deployment.DeployContract(e.Logger, evmChains[config.ChainSelector], ab, deployFiredrillEntrypoint)
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

func deployFiredrillEntrypointV1_5(chain cldf_evm.Chain) deployment.ContractDeploy[*firedrill_entrypoint_v1_5.FiredrillEntrypoint] {
	address, tx, inst, err := firedrill_entrypoint_v1_5.DeployFiredrillEntrypoint(chain.DeployerKey, chain.Client, chain.Selector)
	return deployment.ContractDeploy[*firedrill_entrypoint_v1_5.FiredrillEntrypoint]{
		Address:  address,
		Contract: inst,
		Tx:       tx,
		Tv:       deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deploy.Version1_5_0),
		Err:      err,
	}
}

func deployFiredrillEntrypoint(chain cldf_evm.Chain) deployment.ContractDeploy[*firedrill_entrypoint.FiredrillEntrypoint] {
	address, tx, inst, err := firedrill_entrypoint.DeployFiredrillEntrypoint(chain.DeployerKey, chain.Client, chain.Selector)
	return deployment.ContractDeploy[*firedrill_entrypoint.FiredrillEntrypoint]{
		Address:  address,
		Contract: inst,
		Tx:       tx,
		Tv:       deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deploy.Version1_6_0),
		Err:      err,
	}
}

func FiredrillRegisterContracts(lggr logger.Logger, addressBook deployment.AddressBook, chain cldf_evm.Chain) error {
	firedrillEntrypointAddr, err := deployment.SearchAddressBook(addressBook, chain.Selector, shared.FiredrillEntrypointType)
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
	entrypointReceiverAddr, err4 := firedrillEntrypoint.Receiver(nil)
	if err := errors.Join(err1, err2, err3, err4); err != nil {
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
	lggr.Infof("FiredrillEntrypoint register events mined. Address: %s, tx: %s, blockNum: %d", firedrillEntrypointAddr, tx.Hash(), blockNum)
	return nil
}
