// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package firedrill_off_ramp

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// FiredrillOffRampCommitReport is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampCommitReport struct {
	PriceUpdates InternalPriceUpdates
	Interval     FiredrillOffRampInterval
	MerkleRoot   [32]byte
}

// FiredrillOffRampDynamicConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampDynamicConfig struct {
	PermissionLessExecutionThresholdSeconds uint32
	MaxDataBytes                            uint32
	MaxNumberOfTokensPerMsg                 uint16
	Router                                  common.Address
	PriceRegistry                           common.Address
}

// FiredrillOffRampInterval is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampInterval struct {
	Min uint64
	Max uint64
}

// FiredrillOffRampStaticConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampStaticConfig struct {
	CommitStore         common.Address
	ChainSelector       uint64
	SourceChainSelector uint64
	OnRamp              common.Address
	PrevOffRamp         common.Address
	RmnProxy            common.Address
	TokenAdminRegistry  common.Address
}

// FiredrillOffRampTokenBucket is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

// InternalGasPriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type InternalGasPriceUpdate struct {
	DestChainSelector uint64
	UsdPerUnitGas     *big.Int
}

// InternalPriceUpdates is an auto generated low-level Go binding around an user-defined struct.
type InternalPriceUpdates struct {
	TokenPriceUpdates []InternalTokenPriceUpdate
	GasPriceUpdates   []InternalGasPriceUpdate
}

// InternalTokenPriceUpdate is an auto generated low-level Go binding around an user-defined struct.
type InternalTokenPriceUpdate struct {
	SourceToken common.Address
	UsdPerToken *big.Int
}

// FiredrillOffRampMetaData contains all meta data concerning the FiredrillOffRamp contract.
var FiredrillOffRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ctrl\",\"type\":\"address\",\"internalType\":\"contractFiredrillCompound\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"currentRateLimiterState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emitConfigSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitExecutionStateChanged\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitReportAccepted\",\"inputs\":[{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitTokenConsumed\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.DynamicConfig\",\"components\":[{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.StaticConfig\",\"components\":[{\"name\":\"commitStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRamp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prevOffRamp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"ConfigSet\",\"inputs\":[{\"name\":\"staticConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.StaticConfig\",\"components\":[{\"name\":\"commitStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRamp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"prevOffRamp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.DynamicConfig\",\"components\":[{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"priceRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecutionStateChanged\",\"inputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"state\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumFiredrillOffRamp.MessageExecutionState\"},{\"name\":\"returnData\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ReportAccepted\",\"inputs\":[{\"name\":\"report\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.CommitReport\",\"components\":[{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]},{\"name\":\"interval\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.Interval\",\"components\":[{\"name\":\"min\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"max\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensConsumed\",\"inputs\":[{\"name\":\"tokens\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b506040516112f93803806112f983398101604081905261002e916100d9565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161006e565b506001600160a01b0316608052610106565b600180546001600160a01b03191690556100878161008a565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e9575f80fd5b81516001600160a01b03811681146100ff575f80fd5b9392505050565b6080516111726101875f395f818161027401528181610367015281816103f50152818161048401528181610513015281816105a6015281816105db0152818161073401528181610822015281816108a10152818161090c0152818161099a01528181610a2901528181610ab801528181610b4f0152610bbe01526111725ff3fe608060405234801561000f575f80fd5b50600436106100e5575f3560e01c8063866ec1d611610088578063cdffbd5f11610063578063cdffbd5f146102e3578063cefe12c5146102eb578063e30c3978146102fe578063f2fde38b1461030f575f80fd5b8063866ec1d6146102b85780638da5cb5b146102cb57806392abf95e146102db575f80fd5b8063546719cd116100c3578063546719cd14610156578063715018a6146102155780637437ff9f1461021f57806379ba5097146102b0575f80fd5b806306285c69146100e9578063181f5a771461010757806321df0da714610136575b5f80fd5b6100f1610322565b6040516100fe9190610e04565b60405180910390f35b604080518082018252600d81526c04f666652616d7020312e352e3609c1b602082015290516100fe9190610e5b565b61013e6105d8565b6040516001600160a01b0390911681526020016100fe565b6101b96040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506040805160a0810182525f80825263ffffffff4216602083015260019282019290925260326060820152608081019190915290565b6040516100fe919081516fffffffffffffffffffffffffffffffff908116825260208084015163ffffffff1690830152604080840151151590830152606080840151821690830152608092830151169181019190915260a00190565b61021d61065e565b005b6102a36040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506040805160a081018252600a80825260208201819052918101919091526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001660608201819052608082015290565b6040516100fe9190610ebf565b61021d610671565b61021d6102c6366004610ee2565b6106ba565b5f546001600160a01b031661013e565b61021d610817565b61021d610896565b61021d6102f9366004610f2d565b610bb3565b6001546001600160a01b031661013e565b61021d61031d366004610f49565b610c9f565b6040805160e0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c08101919091526040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166344671a5d6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103e59190610f64565b6001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa15801561044f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104739190610f7f565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156104de573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105029190610f7f565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561056d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105919190610f64565b6001600160a01b0390811682525f60208301527f00000000000000000000000000000000000000000000000000000000000000001660408201819052606090910152919050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610635573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106599190610f64565b905090565b610666610d0f565b61066f5f610d3b565b565b60015433906001600160a01b031681146106ae5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b6106b781610d3b565b50565b60408051808201909152606080825260208201525f60405180604001604052808567ffffffffffffffff1681526020018467ffffffffffffffff1681525090507f291698c01aa71f912280535d88a00d2c59fb63530a3f5d0098560468acb9ebf560405180606001604052808481526020018381526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166344671a5d6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561078e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107b29190610f64565b604080516001600160a01b03909216602083015267ffffffffffffffff808a169183019190915287166060820152608001604051602081830303815290604052805190602001208152506040516108099190610ff2565b60405180910390a150505050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461085f5760405162461bcd60e51b81526004016106a5906110be565b604051606481527f1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a906020015b60405180910390a1565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146108de5760405162461bcd60e51b81526004016106a5906110be565b7f7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d6112606040518060e001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166344671a5d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610966573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061098a9190610f64565b6001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156109f4573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a189190610f7f565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a83573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610aa79190610f7f565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610b12573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b369190610f64565b6001600160a01b0390811682525f6020808401919091527f000000000000000000000000000000000000000000000000000000000000000090911660408084018290526060938401829052805160a081018252600a808252938101849052808201939093529282018190526080820152905161088c9291906110e4565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610bfb5760405162461bcd60e51b81526004016106a5906110be565b6040516bffffffffffffffffffffffff19606084901b1660208201526001600160c01b031960c083901b166034820152603c016040516020818303038152906040528051906020012060017fd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef6560016040518060400160405280600381526020016231323360e81b815250604051610c93929190611100565b60405180910390a35050565b610ca7610d0f565b600180546001600160a01b0383166001600160a01b03199091168117909155610cd75f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b0316331461066f5760405163118cdaa760e01b81523360048201526024016106a5565b600180546001600160a01b03191690556106b7815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b60018060a01b03808251168352602082015167ffffffffffffffff808216602086015280604085015116604086015250508060608301511660608401528060808301511660808401528060a08301511660a08401528060c08301511660c0840152505050565b60e08101610e128284610d9e565b92915050565b5f81518084525f5b81811015610e3c57602081850181015186830182015201610e20565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610e6d6020830184610e18565b9392505050565b805163ffffffff90811683526020808301519091169083015260408082015161ffff16908301526060808201516001600160a01b039081169184019190915260809182015116910152565b60a08101610e128284610e74565b67ffffffffffffffff811681146106b7575f80fd5b5f8060408385031215610ef3575f80fd5b8235610efe81610ecd565b91506020830135610f0e81610ecd565b809150509250929050565b6001600160a01b03811681146106b7575f80fd5b5f8060408385031215610f3e575f80fd5b8235610efe81610f19565b5f60208284031215610f59575f80fd5b8135610e6d81610f19565b5f60208284031215610f74575f80fd5b8151610e6d81610f19565b5f60208284031215610f8f575f80fd5b8151610e6d81610ecd565b5f815180845260208085019450602084015f5b83811015610fe7578151805167ffffffffffffffff1688528301516001600160e01b03168388015260409096019590820190600101610fad565b509495945050505050565b602080825282516080838301528051604060a08501819052815160e086018190525f94939284019185916101008801905b8084101561105e57845180516001600160a01b031683528701516001600160e01b031687830152938601936001939093019290820190611023565b5093850151878503609f190160c08901529361107a8186610f9a565b95890151805167ffffffffffffffff90811660408b015260208201511660608a01529594506110a99350505050565b60408501516080850152809250505092915050565b6020808252600c908201526b1bdb9b1e4818dbdb9d1c9bdb60a21b604082015260600190565b61018081016110f38285610d9e565b610e6d60e0830184610e74565b5f6004841061111d57634e487b7160e01b5f52602160045260245ffd5b838252604060208301526111346040830184610e18565b94935050505056fea2646970667358221220b6182c22c471b7b286375fb14cb8c1ebc8888fd3ea8be8c12d02e25c8d7a8a8d64736f6c63430008180033",
}

// FiredrillOffRampABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillOffRampMetaData.ABI instead.
var FiredrillOffRampABI = FiredrillOffRampMetaData.ABI

// FiredrillOffRampBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiredrillOffRampMetaData.Bin instead.
var FiredrillOffRampBin = FiredrillOffRampMetaData.Bin

// DeployFiredrillOffRamp deploys a new Ethereum contract, binding an instance of FiredrillOffRamp to it.
func DeployFiredrillOffRamp(auth *bind.TransactOpts, backend bind.ContractBackend, ctrl common.Address) (common.Address, *types.Transaction, *FiredrillOffRamp, error) {
	parsed, err := FiredrillOffRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiredrillOffRampBin), backend, ctrl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiredrillOffRamp{FiredrillOffRampCaller: FiredrillOffRampCaller{contract: contract}, FiredrillOffRampTransactor: FiredrillOffRampTransactor{contract: contract}, FiredrillOffRampFilterer: FiredrillOffRampFilterer{contract: contract}}, nil
}

// FiredrillOffRamp is an auto generated Go binding around an Ethereum contract.
type FiredrillOffRamp struct {
	FiredrillOffRampCaller     // Read-only binding to the contract
	FiredrillOffRampTransactor // Write-only binding to the contract
	FiredrillOffRampFilterer   // Log filterer for contract events
}

// FiredrillOffRampCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiredrillOffRampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOffRampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiredrillOffRampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOffRampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiredrillOffRampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOffRampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiredrillOffRampSession struct {
	Contract     *FiredrillOffRamp // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiredrillOffRampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiredrillOffRampCallerSession struct {
	Contract *FiredrillOffRampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// FiredrillOffRampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiredrillOffRampTransactorSession struct {
	Contract     *FiredrillOffRampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// FiredrillOffRampRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiredrillOffRampRaw struct {
	Contract *FiredrillOffRamp // Generic contract binding to access the raw methods on
}

// FiredrillOffRampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiredrillOffRampCallerRaw struct {
	Contract *FiredrillOffRampCaller // Generic read-only contract binding to access the raw methods on
}

// FiredrillOffRampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiredrillOffRampTransactorRaw struct {
	Contract *FiredrillOffRampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiredrillOffRamp creates a new instance of FiredrillOffRamp, bound to a specific deployed contract.
func NewFiredrillOffRamp(address common.Address, backend bind.ContractBackend) (*FiredrillOffRamp, error) {
	contract, err := bindFiredrillOffRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRamp{FiredrillOffRampCaller: FiredrillOffRampCaller{contract: contract}, FiredrillOffRampTransactor: FiredrillOffRampTransactor{contract: contract}, FiredrillOffRampFilterer: FiredrillOffRampFilterer{contract: contract}}, nil
}

// NewFiredrillOffRampCaller creates a new read-only instance of FiredrillOffRamp, bound to a specific deployed contract.
func NewFiredrillOffRampCaller(address common.Address, caller bind.ContractCaller) (*FiredrillOffRampCaller, error) {
	contract, err := bindFiredrillOffRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampCaller{contract: contract}, nil
}

// NewFiredrillOffRampTransactor creates a new write-only instance of FiredrillOffRamp, bound to a specific deployed contract.
func NewFiredrillOffRampTransactor(address common.Address, transactor bind.ContractTransactor) (*FiredrillOffRampTransactor, error) {
	contract, err := bindFiredrillOffRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampTransactor{contract: contract}, nil
}

// NewFiredrillOffRampFilterer creates a new log filterer instance of FiredrillOffRamp, bound to a specific deployed contract.
func NewFiredrillOffRampFilterer(address common.Address, filterer bind.ContractFilterer) (*FiredrillOffRampFilterer, error) {
	contract, err := bindFiredrillOffRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampFilterer{contract: contract}, nil
}

// bindFiredrillOffRamp binds a generic wrapper to an already deployed contract.
func bindFiredrillOffRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiredrillOffRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillOffRamp *FiredrillOffRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillOffRamp.Contract.FiredrillOffRampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillOffRamp *FiredrillOffRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.FiredrillOffRampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillOffRamp *FiredrillOffRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.FiredrillOffRampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillOffRamp *FiredrillOffRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillOffRamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillOffRamp *FiredrillOffRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillOffRamp *FiredrillOffRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.contract.Transact(opts, method, params...)
}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillOffRamp *FiredrillOffRampCaller) CurrentRateLimiterState(opts *bind.CallOpts) (FiredrillOffRampTokenBucket, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(FiredrillOffRampTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOffRampTokenBucket)).(*FiredrillOffRampTokenBucket)

	return out0, err

}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillOffRamp *FiredrillOffRampSession) CurrentRateLimiterState() (FiredrillOffRampTokenBucket, error) {
	return _FiredrillOffRamp.Contract.CurrentRateLimiterState(&_FiredrillOffRamp.CallOpts)
}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) CurrentRateLimiterState() (FiredrillOffRampTokenBucket, error) {
	return _FiredrillOffRamp.Contract.CurrentRateLimiterState(&_FiredrillOffRamp.CallOpts)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((uint32,uint32,uint16,address,address))
func (_FiredrillOffRamp *FiredrillOffRampCaller) GetDynamicConfig(opts *bind.CallOpts) (FiredrillOffRampDynamicConfig, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(FiredrillOffRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOffRampDynamicConfig)).(*FiredrillOffRampDynamicConfig)

	return out0, err

}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((uint32,uint32,uint16,address,address))
func (_FiredrillOffRamp *FiredrillOffRampSession) GetDynamicConfig() (FiredrillOffRampDynamicConfig, error) {
	return _FiredrillOffRamp.Contract.GetDynamicConfig(&_FiredrillOffRamp.CallOpts)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((uint32,uint32,uint16,address,address))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetDynamicConfig() (FiredrillOffRampDynamicConfig, error) {
	return _FiredrillOffRamp.Contract.GetDynamicConfig(&_FiredrillOffRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,address,address,address,address))
func (_FiredrillOffRamp *FiredrillOffRampCaller) GetStaticConfig(opts *bind.CallOpts) (FiredrillOffRampStaticConfig, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FiredrillOffRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOffRampStaticConfig)).(*FiredrillOffRampStaticConfig)

	return out0, err

}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,address,address,address,address))
func (_FiredrillOffRamp *FiredrillOffRampSession) GetStaticConfig() (FiredrillOffRampStaticConfig, error) {
	return _FiredrillOffRamp.Contract.GetStaticConfig(&_FiredrillOffRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,address,address,address,address))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetStaticConfig() (FiredrillOffRampStaticConfig, error) {
	return _FiredrillOffRamp.Contract.GetStaticConfig(&_FiredrillOffRamp.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampSession) GetToken() (common.Address, error) {
	return _FiredrillOffRamp.Contract.GetToken(&_FiredrillOffRamp.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetToken() (common.Address, error) {
	return _FiredrillOffRamp.Contract.GetToken(&_FiredrillOffRamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampSession) Owner() (common.Address, error) {
	return _FiredrillOffRamp.Contract.Owner(&_FiredrillOffRamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) Owner() (common.Address, error) {
	return _FiredrillOffRamp.Contract.Owner(&_FiredrillOffRamp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampSession) PendingOwner() (common.Address, error) {
	return _FiredrillOffRamp.Contract.PendingOwner(&_FiredrillOffRamp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) PendingOwner() (common.Address, error) {
	return _FiredrillOffRamp.Contract.PendingOwner(&_FiredrillOffRamp.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOffRamp *FiredrillOffRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOffRamp *FiredrillOffRampSession) TypeAndVersion() (string, error) {
	return _FiredrillOffRamp.Contract.TypeAndVersion(&_FiredrillOffRamp.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) TypeAndVersion() (string, error) {
	return _FiredrillOffRamp.Contract.TypeAndVersion(&_FiredrillOffRamp.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.AcceptOwnership(&_FiredrillOffRamp.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.AcceptOwnership(&_FiredrillOffRamp.TransactOpts)
}

// EmitConfigSet is a paid mutator transaction binding the contract method 0xcdffbd5f.
//
// Solidity: function emitConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitConfigSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitConfigSet")
}

// EmitConfigSet is a paid mutator transaction binding the contract method 0xcdffbd5f.
//
// Solidity: function emitConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitConfigSet() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitConfigSet(&_FiredrillOffRamp.TransactOpts)
}

// EmitConfigSet is a paid mutator transaction binding the contract method 0xcdffbd5f.
//
// Solidity: function emitConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitConfigSet() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitConfigSet(&_FiredrillOffRamp.TransactOpts)
}

// EmitExecutionStateChanged is a paid mutator transaction binding the contract method 0xcefe12c5.
//
// Solidity: function emitExecutionStateChanged(address sender, uint64 index) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitExecutionStateChanged(opts *bind.TransactOpts, sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitExecutionStateChanged", sender, index)
}

// EmitExecutionStateChanged is a paid mutator transaction binding the contract method 0xcefe12c5.
//
// Solidity: function emitExecutionStateChanged(address sender, uint64 index) returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitExecutionStateChanged(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitExecutionStateChanged(&_FiredrillOffRamp.TransactOpts, sender, index)
}

// EmitExecutionStateChanged is a paid mutator transaction binding the contract method 0xcefe12c5.
//
// Solidity: function emitExecutionStateChanged(address sender, uint64 index) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitExecutionStateChanged(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitExecutionStateChanged(&_FiredrillOffRamp.TransactOpts, sender, index)
}

// EmitReportAccepted is a paid mutator transaction binding the contract method 0x866ec1d6.
//
// Solidity: function emitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitReportAccepted(opts *bind.TransactOpts, minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitReportAccepted", minSeqNr, maxSeqNr)
}

// EmitReportAccepted is a paid mutator transaction binding the contract method 0x866ec1d6.
//
// Solidity: function emitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitReportAccepted(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitReportAccepted(&_FiredrillOffRamp.TransactOpts, minSeqNr, maxSeqNr)
}

// EmitReportAccepted is a paid mutator transaction binding the contract method 0x866ec1d6.
//
// Solidity: function emitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitReportAccepted(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitReportAccepted(&_FiredrillOffRamp.TransactOpts, minSeqNr, maxSeqNr)
}

// EmitTokenConsumed is a paid mutator transaction binding the contract method 0x92abf95e.
//
// Solidity: function emitTokenConsumed() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitTokenConsumed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitTokenConsumed")
}

// EmitTokenConsumed is a paid mutator transaction binding the contract method 0x92abf95e.
//
// Solidity: function emitTokenConsumed() returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitTokenConsumed() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitTokenConsumed(&_FiredrillOffRamp.TransactOpts)
}

// EmitTokenConsumed is a paid mutator transaction binding the contract method 0x92abf95e.
//
// Solidity: function emitTokenConsumed() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitTokenConsumed() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitTokenConsumed(&_FiredrillOffRamp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.RenounceOwnership(&_FiredrillOffRamp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.RenounceOwnership(&_FiredrillOffRamp.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.TransferOwnership(&_FiredrillOffRamp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.TransferOwnership(&_FiredrillOffRamp.TransactOpts, newOwner)
}

// FiredrillOffRampConfigSetIterator is returned from FilterConfigSet and is used to iterate over the raw logs and unpacked data for ConfigSet events raised by the FiredrillOffRamp contract.
type FiredrillOffRampConfigSetIterator struct {
	Event *FiredrillOffRampConfigSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampConfigSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampConfigSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampConfigSet represents a ConfigSet event raised by the FiredrillOffRamp contract.
type FiredrillOffRampConfigSet struct {
	StaticConfig  FiredrillOffRampStaticConfig
	DynamicConfig FiredrillOffRampDynamicConfig
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConfigSet is a free log retrieval operation binding the contract event 0x7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d611260.
//
// Solidity: event ConfigSet((address,uint64,uint64,address,address,address,address) staticConfig, (uint32,uint32,uint16,address,address) dynamicConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterConfigSet(opts *bind.FilterOpts) (*FiredrillOffRampConfigSetIterator, error) {

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampConfigSetIterator{contract: _FiredrillOffRamp.contract, event: "ConfigSet", logs: logs, sub: sub}, nil
}

// WatchConfigSet is a free log subscription operation binding the contract event 0x7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d611260.
//
// Solidity: event ConfigSet((address,uint64,uint64,address,address,address,address) staticConfig, (uint32,uint32,uint16,address,address) dynamicConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchConfigSet(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampConfigSet) (event.Subscription, error) {

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "ConfigSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampConfigSet)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseConfigSet is a log parse operation binding the contract event 0x7879e20bb60a503429de4a2c912b5904f08a39f2af054c10fb46434b5d611260.
//
// Solidity: event ConfigSet((address,uint64,uint64,address,address,address,address) staticConfig, (uint32,uint32,uint16,address,address) dynamicConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseConfigSet(log types.Log) (*FiredrillOffRampConfigSet, error) {
	event := new(FiredrillOffRampConfigSet)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "ConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOffRampExecutionStateChangedIterator is returned from FilterExecutionStateChanged and is used to iterate over the raw logs and unpacked data for ExecutionStateChanged events raised by the FiredrillOffRamp contract.
type FiredrillOffRampExecutionStateChangedIterator struct {
	Event *FiredrillOffRampExecutionStateChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampExecutionStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampExecutionStateChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampExecutionStateChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampExecutionStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampExecutionStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampExecutionStateChanged represents a ExecutionStateChanged event raised by the FiredrillOffRamp contract.
type FiredrillOffRampExecutionStateChanged struct {
	SequenceNumber uint64
	MessageId      [32]byte
	State          uint8
	ReturnData     []byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterExecutionStateChanged is a free log retrieval operation binding the contract event 0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65.
//
// Solidity: event ExecutionStateChanged(uint64 indexed sequenceNumber, bytes32 indexed messageId, uint8 state, bytes returnData)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterExecutionStateChanged(opts *bind.FilterOpts, sequenceNumber []uint64, messageId [][32]byte) (*FiredrillOffRampExecutionStateChangedIterator, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampExecutionStateChangedIterator{contract: _FiredrillOffRamp.contract, event: "ExecutionStateChanged", logs: logs, sub: sub}, nil
}

// WatchExecutionStateChanged is a free log subscription operation binding the contract event 0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65.
//
// Solidity: event ExecutionStateChanged(uint64 indexed sequenceNumber, bytes32 indexed messageId, uint8 state, bytes returnData)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchExecutionStateChanged(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampExecutionStateChanged, sequenceNumber []uint64, messageId [][32]byte) (event.Subscription, error) {

	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}
	var messageIdRule []interface{}
	for _, messageIdItem := range messageId {
		messageIdRule = append(messageIdRule, messageIdItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "ExecutionStateChanged", sequenceNumberRule, messageIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampExecutionStateChanged)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseExecutionStateChanged is a log parse operation binding the contract event 0xd4f851956a5d67c3997d1c9205045fef79bae2947fdee7e9e2641abc7391ef65.
//
// Solidity: event ExecutionStateChanged(uint64 indexed sequenceNumber, bytes32 indexed messageId, uint8 state, bytes returnData)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseExecutionStateChanged(log types.Log) (*FiredrillOffRampExecutionStateChanged, error) {
	event := new(FiredrillOffRampExecutionStateChanged)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "ExecutionStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOffRampOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FiredrillOffRamp contract.
type FiredrillOffRampOwnershipTransferStartedIterator struct {
	Event *FiredrillOffRampOwnershipTransferStarted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampOwnershipTransferStarted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampOwnershipTransferStarted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FiredrillOffRamp contract.
type FiredrillOffRampOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillOffRampOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampOwnershipTransferStartedIterator{contract: _FiredrillOffRamp.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampOwnershipTransferStarted)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferStarted is a log parse operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseOwnershipTransferStarted(log types.Log) (*FiredrillOffRampOwnershipTransferStarted, error) {
	event := new(FiredrillOffRampOwnershipTransferStarted)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOffRampOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FiredrillOffRamp contract.
type FiredrillOffRampOwnershipTransferredIterator struct {
	Event *FiredrillOffRampOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampOwnershipTransferred represents a OwnershipTransferred event raised by the FiredrillOffRamp contract.
type FiredrillOffRampOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillOffRampOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampOwnershipTransferredIterator{contract: _FiredrillOffRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampOwnershipTransferred)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseOwnershipTransferred(log types.Log) (*FiredrillOffRampOwnershipTransferred, error) {
	event := new(FiredrillOffRampOwnershipTransferred)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOffRampReportAcceptedIterator is returned from FilterReportAccepted and is used to iterate over the raw logs and unpacked data for ReportAccepted events raised by the FiredrillOffRamp contract.
type FiredrillOffRampReportAcceptedIterator struct {
	Event *FiredrillOffRampReportAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampReportAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampReportAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampReportAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampReportAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampReportAccepted represents a ReportAccepted event raised by the FiredrillOffRamp contract.
type FiredrillOffRampReportAccepted struct {
	Report FiredrillOffRampCommitReport
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReportAccepted is a free log retrieval operation binding the contract event 0x291698c01aa71f912280535d88a00d2c59fb63530a3f5d0098560468acb9ebf5.
//
// Solidity: event ReportAccepted((((address,uint224)[],(uint64,uint224)[]),(uint64,uint64),bytes32) report)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterReportAccepted(opts *bind.FilterOpts) (*FiredrillOffRampReportAcceptedIterator, error) {

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampReportAcceptedIterator{contract: _FiredrillOffRamp.contract, event: "ReportAccepted", logs: logs, sub: sub}, nil
}

// WatchReportAccepted is a free log subscription operation binding the contract event 0x291698c01aa71f912280535d88a00d2c59fb63530a3f5d0098560468acb9ebf5.
//
// Solidity: event ReportAccepted((((address,uint224)[],(uint64,uint224)[]),(uint64,uint64),bytes32) report)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchReportAccepted(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampReportAccepted) (event.Subscription, error) {

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "ReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampReportAccepted)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseReportAccepted is a log parse operation binding the contract event 0x291698c01aa71f912280535d88a00d2c59fb63530a3f5d0098560468acb9ebf5.
//
// Solidity: event ReportAccepted((((address,uint224)[],(uint64,uint224)[]),(uint64,uint64),bytes32) report)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseReportAccepted(log types.Log) (*FiredrillOffRampReportAccepted, error) {
	event := new(FiredrillOffRampReportAccepted)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "ReportAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOffRampTokensConsumedIterator is returned from FilterTokensConsumed and is used to iterate over the raw logs and unpacked data for TokensConsumed events raised by the FiredrillOffRamp contract.
type FiredrillOffRampTokensConsumedIterator struct {
	Event *FiredrillOffRampTokensConsumed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FiredrillOffRampTokensConsumedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampTokensConsumed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FiredrillOffRampTokensConsumed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FiredrillOffRampTokensConsumedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampTokensConsumedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampTokensConsumed represents a TokensConsumed event raised by the FiredrillOffRamp contract.
type FiredrillOffRampTokensConsumed struct {
	Tokens *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensConsumed is a free log retrieval operation binding the contract event 0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a.
//
// Solidity: event TokensConsumed(uint256 tokens)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterTokensConsumed(opts *bind.FilterOpts) (*FiredrillOffRampTokensConsumedIterator, error) {

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampTokensConsumedIterator{contract: _FiredrillOffRamp.contract, event: "TokensConsumed", logs: logs, sub: sub}, nil
}

// WatchTokensConsumed is a free log subscription operation binding the contract event 0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a.
//
// Solidity: event TokensConsumed(uint256 tokens)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchTokensConsumed(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampTokensConsumed) (event.Subscription, error) {

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "TokensConsumed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampTokensConsumed)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokensConsumed is a log parse operation binding the contract event 0x1871cdf8010e63f2eb8384381a68dfa7416dc571a5517e66e88b2d2d0c0a690a.
//
// Solidity: event TokensConsumed(uint256 tokens)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseTokensConsumed(log types.Log) (*FiredrillOffRampTokensConsumed, error) {
	event := new(FiredrillOffRampTokensConsumed)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "TokensConsumed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
