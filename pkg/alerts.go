package pkg

import (
	"fmt"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	chainsel "github.com/smartcontractkit/chain-selectors"
	"github.com/smartcontractkit/chainlink-common/pkg/logger"
	"github.com/smartcontractkit/chainlink-deployments-framework/deployment"
	deploy "github.com/smartcontractkit/chainlink/deployment"
	"github.com/smartcontractkit/chainlink/deployment/ccip/view"

	firedrill "github.com/smartcontractkit/ccip-firedrill-deployment/deployment"
	"github.com/smartcontractkit/ccip-firedrill-deployment/deployment/shared"
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
	Routers              []GlobalInputContract `yaml:"router_addresses"`
	CommitStores         []GlobalInputContract `yaml:"commit_store_addresses,omitempty"`
	OffRamps             []GlobalInputContract `yaml:"off_ramp_addresses"`
	FiredrillTokens      []GlobalInputContract `yaml:"firedrill_token_addresses"`
	FiredrillOnRamps     []GlobalInputContract `yaml:"firedrill_on_ramp_addresses"`
	FiredrillOffRamps    []GlobalInputContract `yaml:"firedrill_off_ramp_addresses"`
	FiredrillEntrypoints []GlobalInputContract `yaml:"firedrill_entrypoint_addresses"`
}

type GlobalInputContract struct {
	Address string `yaml:"address"`
	Network string `yaml:"network"`
}

func AlertsGlobalInput(lggr logger.Logger, destChain chainsel.ChainDetails, version string, ccipFiredrillView firedrill.CCIPFiredrillView, ccipView view.CCIPView) (GlobalInput, error) {
	isTestnet := strings.Contains(destChain.ChainName, "testnet") || strings.Contains(destChain.ChainName, "devnet")
	var srcChain chainsel.ChainDetails
	if isTestnet {
		srcChain = chainsel.ChainDetails{
			ChainSelector: chainsel.ETHEREUM_TESTNET_SEPOLIA.Selector,
			ChainName:     chainsel.ETHEREUM_TESTNET_SEPOLIA.Name,
		}
	} else {
		srcChain = chainsel.ChainDetails{
			ChainSelector: chainsel.ETHEREUM_MAINNET.Selector,
			ChainName:     chainsel.ETHEREUM_MAINNET.Name,
		}
	}
	var entrypoints map[string]shared.FiredrillEntrypointView
	if firedrillChainView, ok := ccipFiredrillView.Chains[destChain.ChainName]; ok {
		entrypoints = firedrillChainView.FiredrillEntrypoint
	} else if firedrillChainView, ok := ccipFiredrillView.SolChains[destChain.ChainName]; ok {
		entrypoints = firedrillChainView.FiredrillEntrypoint
	} else {
		return GlobalInput{}, fmt.Errorf("no firedrills known for chain %s", destChain.ChainName)
	}
	var network GlobalInputNetwork
	for _, entrypointView := range entrypoints {
		entrypointVersion := deployment.MustTypeAndVersionFromString(entrypointView.TypeAndVersion)
		if entrypointVersion.Version.String() != version {
			continue
		}
		lggr.Infof("Adding active firedrill entrypoint %+v", entrypointView)
		network.Contracts.FiredrillTokens = append(
			network.Contracts.FiredrillTokens,
			GlobalInputContract{Address: entrypointView.Token, Network: destChain.ChainName},
		)
		network.Contracts.FiredrillOnRamps = append(
			network.Contracts.FiredrillOnRamps,
			GlobalInputContract{Address: entrypointView.OnRamp, Network: destChain.ChainName},
		)
		network.Contracts.FiredrillOffRamps = append(
			network.Contracts.FiredrillOffRamps,
			GlobalInputContract{Address: entrypointView.OffRamp, Network: destChain.ChainName},
		)
		network.Contracts.FiredrillEntrypoints = append(
			network.Contracts.FiredrillEntrypoints,
			GlobalInputContract{Address: entrypointView.Address, Network: destChain.ChainName},
		)
	}
	if len(network.Contracts.FiredrillEntrypoints) == 0 {
		return GlobalInput{}, fmt.Errorf("no active firedrills %s known for chain %s", version, destChain.ChainName)
	}
	switch version {
	case deploy.Version1_5_0.String():
		contractsSrcDest := lookupCCIPv1_5(ccipView, srcChain, destChain)
		if len(contractsSrcDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s -> %s", srcChain.ChainName, destChain.ChainName)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsSrcDest.Routers...)
		network.Contracts.CommitStores = append(network.Contracts.CommitStores, contractsSrcDest.CommitStores...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsSrcDest.OffRamps...)
		contractsDestSrc := lookupCCIPv1_5(ccipView, destChain, srcChain)
		if len(contractsSrcDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s <- %s", srcChain.ChainName, destChain.ChainName)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsDestSrc.Routers...)
		network.Contracts.CommitStores = append(network.Contracts.CommitStores, contractsDestSrc.CommitStores...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsDestSrc.OffRamps...)
	case deploy.Version1_6_0.String():
		contractsSrc := lookupCCIPv1_6(ccipView, srcChain, destChain)
		if len(contractsSrc.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s -> %s", srcChain.ChainName, destChain.ChainName)
		}
		network.Contracts.Routers = append(network.Contracts.Routers, contractsSrc.Routers...)
		network.Contracts.OffRamps = append(network.Contracts.OffRamps, contractsSrc.OffRamps...)

		contractsDest := lookupCCIPv1_6(ccipView, destChain, srcChain)
		if len(contractsDest.OffRamps) == 0 {
			return GlobalInput{}, fmt.Errorf("no off ramps found for %s <- %s", srcChain.ChainName, destChain.ChainName)
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

func lookupCCIPv1_5(ccipView view.CCIPView, srcChain chainsel.ChainDetails, destChain chainsel.ChainDetails) GlobalInputContracts {
	destChainView := ccipView.Chains[destChain.ChainName]
	for _, offRampView := range destChainView.EVM2EVMOffRamp {
		if offRampView.StaticConfig.SourceChainSelector == srcChain.ChainSelector {
			// There shouldn't be several matching offramps
			return GlobalInputContracts{
				Routers:      []GlobalInputContract{{Address: strings.ToLower(offRampView.DynamicConfig.Router.Hex()), Network: destChain.ChainName}},
				CommitStores: []GlobalInputContract{{Address: strings.ToLower(offRampView.StaticConfig.CommitStore.Hex()), Network: destChain.ChainName}},
				OffRamps:     []GlobalInputContract{{Address: strings.ToLower(offRampView.Address.Hex()), Network: destChain.ChainName}},
			}
		}
	}
	return GlobalInputContracts{}
}

func lookupCCIPv1_6(ccipView view.CCIPView, srcChain chainsel.ChainDetails, destChain chainsel.ChainDetails) GlobalInputContracts {
	contracts := GlobalInputContracts{}
	if chainView, ok := ccipView.Chains[destChain.ChainName]; ok {
		for _, offRampView := range chainView.OffRamp {
			sourceChainConfig, ok := offRampView.SourceChainConfigs[srcChain.ChainSelector]
			if !ok || sourceChainConfig.Router == (common.Address{}) {
				continue
			}
			contracts.OffRamps = append(
				contracts.OffRamps,
				GlobalInputContract{Address: strings.ToLower(offRampView.Address.Hex()), Network: destChain.ChainName},
			)
			// There shouldn't be several matching offramps
			return GlobalInputContracts{
				Routers:  []GlobalInputContract{{Address: strings.ToLower(sourceChainConfig.Router.Hex()), Network: destChain.ChainName}},
				OffRamps: []GlobalInputContract{{Address: strings.ToLower(offRampView.Address.Hex()), Network: destChain.ChainName}},
			}
		}
	}
	if chainView, ok := ccipView.SolChains[destChain.ChainName]; ok {
		for offRampAddress, offRampView := range chainView.OffRamp {
			sourceChainConfig, ok := offRampView.SourceChains[srcChain.ChainSelector]
			if !ok || !sourceChainConfig.IsEnabled {
				continue
			}
			contracts.OffRamps = append(
				contracts.OffRamps,
				GlobalInputContract{Address: offRampAddress, Network: destChain.ChainName},
			)
			// There shouldn't be several matching offramps
			return GlobalInputContracts{
				Routers:  []GlobalInputContract{{Address: offRampView.ReferenceAddresses.Router, Network: destChain.ChainName}},
				OffRamps: []GlobalInputContract{{Address: offRampAddress, Network: destChain.ChainName}},
			}
		}
	}
	return GlobalInputContracts{}
}
