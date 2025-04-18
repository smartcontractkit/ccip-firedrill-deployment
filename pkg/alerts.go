package pkg

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	cldf "github.com/smartcontractkit/chainlink-deployments-framework/deployment"

	firedrill "github.com/smartcontractkit/ccip-firedrill-deployment/deployment"

	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/ccip/view"
)

type GlobalInput struct {
	CCIP GlobalInputProduct `yaml:"ccip"`
}

type GlobalInputProduct struct {
	Mainnet *GlobalInputNetwork `yaml:"mainnet,omitempty"`
	Testnet *GlobalInputNetwork `yaml:"testnet,omitempty"`
}

type GlobalInputNetwork struct {
	Contracts GlobalInputContracts `yaml:"contracts"`
}

type GlobalInputContracts struct {
	Routers            []GlobalInputContract `yaml:"router_addresses"`
	CommitStores       []GlobalInputContract `yaml:"commit_store_addresses,omitempty"`
	OffRamps           []GlobalInputContract `yaml:"off_ramp_addresses"`
	FiredrillTokens    []GlobalInputContract `yaml:"firedrill_token_addresses"`
	FiredrillOnRamps   []GlobalInputContract `yaml:"firedrill_on_ramp_addresses"`
	FiredrillOffRamps  []GlobalInputContract `yaml:"firedrill_off_ramp_addresses"`
	FiredrillCompounds []GlobalInputContract `yaml:"firedrill_compound_addresses"`
}

type GlobalInputContract struct {
	Address common.Address `yaml:"address"`
	Network string         `yaml:"network"`
}

func AlertsGlobalInput(lggr logger.Logger, destChain chainsel.Chain, version string, ccipFiredrillView firedrill.CCIPFiredrillView, ccipView view.CCIPView) (GlobalInput, error) {
	isTestnet := strings.Contains(destChain.Name, "testnet")
	var srcChain chainsel.Chain
	if isTestnet {
		srcChain = chainsel.ETHEREUM_TESTNET_SEPOLIA
	} else {
		srcChain = chainsel.ETHEREUM_MAINNET
	}
	firedrillChainView, ok := ccipFiredrillView.Chains[destChain.Name]
	if !ok {
		return GlobalInput{}, fmt.Errorf("no firedrills known for chain %s", destChain.Name)
	}
	var network GlobalInputNetwork
	for _, entrypointView := range firedrillChainView.FiredrillEntrypoint {
		if !entrypointView.Active {
			continue
		}
		entrypointVersion := cldf.MustTypeAndVersionFromString(entrypointView.TypeAndVersion)
		if entrypointVersion.Version.String() != version {
			continue
		}
		lggr.Infof("Adding active firedrill entrypoint %+v", entrypointView)
		network.Contracts.FiredrillTokens = append(
			network.Contracts.FiredrillTokens,
			GlobalInputContract{Address: entrypointView.Token, Network: destChain.Name},
		)
		network.Contracts.FiredrillOnRamps = append(
			network.Contracts.FiredrillOnRamps,
			GlobalInputContract{Address: entrypointView.OnRamp, Network: destChain.Name},
		)
		network.Contracts.FiredrillOffRamps = append(
			network.Contracts.FiredrillOffRamps,
			GlobalInputContract{Address: entrypointView.OffRamp, Network: destChain.Name},
		)
		network.Contracts.FiredrillCompounds = append(
			network.Contracts.FiredrillCompounds,
			GlobalInputContract{Address: entrypointView.Compound, Network: destChain.Name},
		)
	}
	if len(network.Contracts.FiredrillCompounds) == 0 {
		return GlobalInput{}, fmt.Errorf("no active firedrills %s known for chain %s", version, destChain.Name)
	}
	switch version {
	case deployment.Version1_5_0.String():
		contractsSrcDest := lookupCCIPv1_5(ccipView, srcChain, destChain)
		if len(contractsSrcDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s -> %s", srcChain.Name, destChain.Name)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsSrcDest.Routers...)
		network.Contracts.CommitStores = append(network.Contracts.CommitStores, contractsSrcDest.CommitStores...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsSrcDest.OffRamps...)
		contractsDestSrc := lookupCCIPv1_5(ccipView, destChain, srcChain)
		if len(contractsSrcDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s <- %s", srcChain.Name, destChain.Name)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsDestSrc.Routers...)
		network.Contracts.CommitStores = append(network.Contracts.CommitStores, contractsDestSrc.CommitStores...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsDestSrc.OffRamps...)
	case deployment.Version1_6_0.String():
		contractsSrc := lookupCCIPv1_6(ccipView, srcChain, destChain)
		if len(contractsSrc.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s", srcChain.Name)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsSrc.Routers...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsSrc.OffRamps...)

		contractsDest := lookupCCIPv1_6(ccipView, destChain, srcChain)
		if len(contractsDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s", destChain.Name)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsDest.Routers...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsDest.OffRamps...)
	default:
		return GlobalInput{}, fmt.Errorf("version %s is not supported", version)
	}
	g := GlobalInput{}
	if isTestnet {
		g.CCIP.Testnet = &network
	} else {
		g.CCIP.Mainnet = &network
	}
	return g, nil
}

func lookupCCIPv1_5(ccipView view.CCIPView, srcChain chainsel.Chain, destChain chainsel.Chain) GlobalInputContracts {
	destChainView := ccipView.Chains[destChain.Name]
	for _, offRampView := range destChainView.EVM2EVMOffRamp {
		if offRampView.StaticConfig.SourceChainSelector == srcChain.Selector {
			// There shouldn't be several matching offramps
			return GlobalInputContracts{
				Routers:      []GlobalInputContract{{Address: offRampView.DynamicConfig.Router, Network: destChain.Name}},
				CommitStores: []GlobalInputContract{{Address: offRampView.StaticConfig.CommitStore, Network: destChain.Name}},
				OffRamps:     []GlobalInputContract{{Address: offRampView.Address, Network: destChain.Name}},
			}
		}
	}
	return GlobalInputContracts{}
}

func lookupCCIPv1_6(ccipView view.CCIPView, srcChain chainsel.Chain, destChain chainsel.Chain) GlobalInputContracts {
	contracts := GlobalInputContracts{}
	destChainView := ccipView.Chains[destChain.Name]
	for _, offRampView := range destChainView.OffRamp {
		sourceChainConfig, ok := offRampView.SourceChainConfigs[srcChain.Selector]
		if !ok || sourceChainConfig.Router == (common.Address{}) {
			continue
		}
		contracts.OffRamps = append(
			contracts.OffRamps,
			GlobalInputContract{Address: offRampView.Address, Network: destChain.Name},
		)
		// There shouldn't be several matching offramps
		return GlobalInputContracts{
			Routers:  []GlobalInputContract{{Address: sourceChainConfig.Router, Network: destChain.Name}},
			OffRamps: []GlobalInputContract{{Address: offRampView.Address, Network: destChain.Name}},
		}
	}
	return GlobalInputContracts{}
}
