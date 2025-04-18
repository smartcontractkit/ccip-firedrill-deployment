package solana

import (
	"fmt"

	solana "github.com/gagliardetto/solana-go"
	chain_selectors "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink/deployment"

	firedrill_entrypoint2 "github.com/smartcontractkit/ccip-firedrill-deployment/chains/solana/gobindings/firedrill_entrypoint"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
)

type SolChainView struct {
	ChainSelector       uint64                                    `json:"svmChainSelector,omitempty"`
	ChainID             string                                    `json:"chainID,omitempty"`
	FiredrillEntrypoint map[string]shared.FiredrillEntrypointView `json:"firedrillEntrypoint,omitempty"`
}

func NewSolChainView(chain deployment.SolChain) (*SolChainView, error) {
	chainID, err := chain_selectors.GetChainIDFromSelector(chain.Selector)
	if err != nil {
		return nil, err
	}
	return &SolChainView{
		ChainSelector:       chain.Selector,
		ChainID:             chainID,
		FiredrillEntrypoint: make(map[string]shared.FiredrillEntrypointView),
	}, nil
}

func SolViewFiredrill(e deployment.Environment) (map[string]*SolChainView, error) {
	ab, err := e.ExistingAddresses.Addresses()
	if err != nil {
		return nil, err
	}
	solChains := make(map[string]*SolChainView)
	for chainSel, addresses := range ab {
		if chain, ok := e.SolChains[chainSel]; ok {
			chainView, err := NewSolChainView(chain)
			if err != nil {
				return nil, err
			}
			for programID, typeAndVersion := range addresses {
				if typeAndVersion.Type != shared.FiredrillEntrypointType || typeAndVersion.Version != deployment.Version1_6_0 {
					return nil, fmt.Errorf("only FiredrillEntrypoint 1.6.0 is supported, but was: %s", typeAndVersion.String())
				}
				firedrillEntrypointAddress := solana.MustPublicKeyFromBase58(programID)
				firedrillEntrypointPDA, _, _ := FindFiredrillEntrypointPDA(firedrillEntrypointAddress)
				var firedrillEntrypointAccount firedrill_entrypoint2.FiredrillEntrypoint
				err := chain.GetAccountDataBorshInto(e.GetContext(), firedrillEntrypointPDA, &firedrillEntrypointAccount)
				if err != nil {
					return nil, fmt.Errorf("firedrillEntrypoint %s not found in existing state chain=%s, not initialized", programID, chain.Name())
				}
				view, err := contractView(firedrillEntrypointAccount, firedrillEntrypointAddress, typeAndVersion.String())
				if err != nil {
					return nil, err
				}
				chainView.FiredrillEntrypoint[programID] = view
			}
			if len(chainView.FiredrillEntrypoint) > 0 {
				solChains[chain.Name()] = chainView
			}
		}

	}
	return solChains, nil
}

func contractView(contract firedrill_entrypoint2.FiredrillEntrypoint, address solana.PublicKey, typeAndVersion string) (shared.FiredrillEntrypointView, error) {
	return shared.FiredrillEntrypointView{
		ContractMetaData: shared.ContractMetaData{
			TypeAndVersion: typeAndVersion,
			Address:        address.String(),
			Owner:          contract.Owner.String(),
		},
		Token:    contract.Token.String(),
		OnRamp:   contract.Compound.String(),
		OffRamp:  contract.OffRamp.String(),
		Compound: contract.Compound.String(),
		Receiver: contract.Receiver.String(),
	}, nil
}
