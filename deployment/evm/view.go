package evm

import (
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deploy "github.com/smartcontractkit/chainlink/deployment"

	firedrill_entrypoint_v1_5 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_5/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/chains/evm/generated/v1_6/gethwrappers/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

func EVMViewFiredrill(e deployment.Environment) (map[string]*shared.ChainView, error) {
	ab, err := e.ExistingAddresses.Addresses()
	if err != nil {
		return nil, err
	}
	chains := make(map[string]*shared.ChainView)
	for chainSel, addresses := range ab {
		if chain, ok := e.Chains[chainSel]; ok {
			chainView, err := shared.NewChainView(chain.Selector)
			if err != nil {
				return nil, err
			}
			for addressStr, typeAndVersion := range addresses {
				address := common.HexToAddress(addressStr)
				switch typeAndVersion.String() {
				case deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deploy.Version1_5_0).String():
					contract, err := firedrill_entrypoint_v1_5.NewFiredrillEntrypoint(address, chain.Client)
					if err != nil {
						return nil, err
					}
					view, err := contractView(contract, address, typeAndVersion.String())
					if err != nil {
						return nil, err
					}
					chainView.FiredrillEntrypoint[addressStr] = view
				case deployment.NewTypeAndVersion(shared.FiredrillEntrypointType, deploy.Version1_6_0).String():
					contract, err := firedrill_entrypoint.NewFiredrillEntrypoint(address, chain.Client)
					if err != nil {
						return nil, err
					}
					view, err := contractView(contract, address, typeAndVersion.String())
					if err != nil {
						return nil, err
					}
					chainView.FiredrillEntrypoint[addressStr] = view
				}
			}
			if len(chainView.FiredrillEntrypoint) > 0 {
				chains[chain.Name()] = chainView
			}
		}
	}
	return chains, nil
}

func contractView(contract FiredrillEntrypoint, address common.Address, typeAndVersion string) (shared.FiredrillEntrypointView, error) {
	owner, err := contract.Owner(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	token, err := contract.Token(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	onramp, err := contract.OnRamp(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	offramp, err := contract.OffRamp(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	compound, err := contract.Compound(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	receiver, err := contract.Receiver(nil)
	if err != nil {
		return shared.FiredrillEntrypointView{}, err
	}
	return shared.FiredrillEntrypointView{
		ContractMetaData: shared.ContractMetaData{
			TypeAndVersion: typeAndVersion,
			Address:        strings.ToLower(address.String()),
			Owner:          strings.ToLower(owner.String()),
		},
		Token:    strings.ToLower(token.String()),
		OnRamp:   strings.ToLower(onramp.String()),
		OffRamp:  strings.ToLower(offramp.String()),
		Compound: strings.ToLower(compound.String()),
		Receiver: strings.ToLower(receiver.String()),
	}, nil
}
