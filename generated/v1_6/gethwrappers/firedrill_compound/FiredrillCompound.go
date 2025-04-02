// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package firedrill_compound

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

// FiredrillCompoundOffRamp is an auto generated low-level Go binding around an user-defined struct.
type FiredrillCompoundOffRamp struct {
	SourceChainSelector uint64
	OffRamp             common.Address
}

// FiredrillCompoundStaticConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillCompoundStaticConfig struct {
	MaxFeeJuelsPerMsg            *big.Int
	LinkToken                    common.Address
	TokenPriceStalenessThreshold uint32
}

// FiredrillCompoundMetaData contains all meta data concerning the FiredrillCompound contract.
var FiredrillCompoundMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ctrl\",\"type\":\"address\",\"internalType\":\"contractFiredrillEntrypoint\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitUsdPerTokenUpdated\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getOffRamps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structFiredrillCompound.OffRamp[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOnRamp\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillCompound.StaticConfig\",\"components\":[{\"name\":\"maxFeeJuelsPerMsg\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenPriceStalenessThreshold\",\"type\":\"uint32\",\"internalType\":\"uint32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isActive\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610a5d380380610a5d83398101604081905261002e916100d9565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161006e565b506001600160a01b0316608052610106565b600180546001600160a01b03191690556100878161008a565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e9575f80fd5b81516001600160a01b03811681146100ff575f80fd5b9392505050565b60805161091c6101415f395f8181610215015281816102ac0152818161037b015281816104a90152818161053801526105e7015261091c5ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c806379ba50971161006e57806379ba50971461015f5780638da5cb5b14610167578063a40e69c71461018b578063a8d87a3b146101a0578063e30c3978146101b3578063f2fde38b146101c4575f80fd5b806306285c69146100aa578063181f5a77146100f657806320f938b61461013557806322f3e2d41461013f578063715018a614610157575b5f80fd5b6100b26101d7565b6040805182516bffffffffffffffffffffffff1681526020808401516001600160a01b0316908201529181015163ffffffff16908201526060015b60405180910390f35b604080518082018252601c81527f526f7574657220312e322e302046656551756f74657220312e362e3000000000602082015290516100ed919061076a565b61013d6102aa565b005b610147610378565b60405190151581526020016100ed565b61013d6103fe565b61013d610411565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100ed565b61019361045a565b6040516100ed91906107b6565b6101736101ae36600461082c565b6105e4565b6001546001600160a01b0316610173565b61013d6101d2366004610862565b61066b565b604080516060810182525f8082526020820181905291810191909152604051806060016040528060016bffffffffffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561026f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610293919061087d565b6001600160a01b031681525f602090910152919050565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610306573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061032a919061087d565b6001600160a01b03167f52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a60014260405161036e929190918252602082015260400190565b60405180910390a2565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166322f3e2d46040518163ffffffff1660e01b8152600401602060405180830381865afa1580156103d5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f99190610898565b905090565b6104066106db565b61040f5f610707565b565b60015433906001600160a01b0316811461044e5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b61045781610707565b50565b6040805160018082528183019092526060915f9190816020015b604080518082019091525f808252602082015281526020019060019003908161047457905050905060405180604001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa158015610503573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052791906108b7565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166344671a5d6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610592573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105b6919061087d565b6001600160a01b0316815250815f815181106105d4576105d46108d2565b6020908102919091010152919050565b5f7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610641573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610665919061087d565b92915050565b6106736106db565b600180546001600160a01b0383166001600160a01b031990911681179091556106a35f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b0316331461040f5760405163118cdaa760e01b8152336004820152602401610445565b600180546001600160a01b0319169055610457815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602080835283518060208501525f5b818110156107965785810183015185820160400152820161077a565b505f604082860101526040601f19601f8301168501019250505092915050565b602080825282518282018190525f919060409081850190868401855b8281101561080a578151805167ffffffffffffffff1685528601516001600160a01b03168685015292840192908501906001016107d2565b5091979650505050505050565b67ffffffffffffffff81168114610457575f80fd5b5f6020828403121561083c575f80fd5b813561084781610817565b9392505050565b6001600160a01b0381168114610457575f80fd5b5f60208284031215610872575f80fd5b81356108478161084e565b5f6020828403121561088d575f80fd5b81516108478161084e565b5f602082840312156108a8575f80fd5b81518015158114610847575f80fd5b5f602082840312156108c7575f80fd5b815161084781610817565b634e487b7160e01b5f52603260045260245ffdfea26469706673582212202d65bbb579790612f7e21bb81964a5a8ec77c7f5c059a32cc562cee934050fb064736f6c63430008180033",
}

// FiredrillCompoundABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillCompoundMetaData.ABI instead.
var FiredrillCompoundABI = FiredrillCompoundMetaData.ABI

// FiredrillCompoundBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiredrillCompoundMetaData.Bin instead.
var FiredrillCompoundBin = FiredrillCompoundMetaData.Bin

// DeployFiredrillCompound deploys a new Ethereum contract, binding an instance of FiredrillCompound to it.
func DeployFiredrillCompound(auth *bind.TransactOpts, backend bind.ContractBackend, ctrl common.Address) (common.Address, *types.Transaction, *FiredrillCompound, error) {
	parsed, err := FiredrillCompoundMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiredrillCompoundBin), backend, ctrl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiredrillCompound{FiredrillCompoundCaller: FiredrillCompoundCaller{contract: contract}, FiredrillCompoundTransactor: FiredrillCompoundTransactor{contract: contract}, FiredrillCompoundFilterer: FiredrillCompoundFilterer{contract: contract}}, nil
}

// FiredrillCompound is an auto generated Go binding around an Ethereum contract.
type FiredrillCompound struct {
	FiredrillCompoundCaller     // Read-only binding to the contract
	FiredrillCompoundTransactor // Write-only binding to the contract
	FiredrillCompoundFilterer   // Log filterer for contract events
}

// FiredrillCompoundCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiredrillCompoundCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillCompoundTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiredrillCompoundTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillCompoundFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiredrillCompoundFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillCompoundSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiredrillCompoundSession struct {
	Contract     *FiredrillCompound // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// FiredrillCompoundCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiredrillCompoundCallerSession struct {
	Contract *FiredrillCompoundCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// FiredrillCompoundTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiredrillCompoundTransactorSession struct {
	Contract     *FiredrillCompoundTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// FiredrillCompoundRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiredrillCompoundRaw struct {
	Contract *FiredrillCompound // Generic contract binding to access the raw methods on
}

// FiredrillCompoundCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiredrillCompoundCallerRaw struct {
	Contract *FiredrillCompoundCaller // Generic read-only contract binding to access the raw methods on
}

// FiredrillCompoundTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiredrillCompoundTransactorRaw struct {
	Contract *FiredrillCompoundTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiredrillCompound creates a new instance of FiredrillCompound, bound to a specific deployed contract.
func NewFiredrillCompound(address common.Address, backend bind.ContractBackend) (*FiredrillCompound, error) {
	contract, err := bindFiredrillCompound(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompound{FiredrillCompoundCaller: FiredrillCompoundCaller{contract: contract}, FiredrillCompoundTransactor: FiredrillCompoundTransactor{contract: contract}, FiredrillCompoundFilterer: FiredrillCompoundFilterer{contract: contract}}, nil
}

// NewFiredrillCompoundCaller creates a new read-only instance of FiredrillCompound, bound to a specific deployed contract.
func NewFiredrillCompoundCaller(address common.Address, caller bind.ContractCaller) (*FiredrillCompoundCaller, error) {
	contract, err := bindFiredrillCompound(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundCaller{contract: contract}, nil
}

// NewFiredrillCompoundTransactor creates a new write-only instance of FiredrillCompound, bound to a specific deployed contract.
func NewFiredrillCompoundTransactor(address common.Address, transactor bind.ContractTransactor) (*FiredrillCompoundTransactor, error) {
	contract, err := bindFiredrillCompound(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundTransactor{contract: contract}, nil
}

// NewFiredrillCompoundFilterer creates a new log filterer instance of FiredrillCompound, bound to a specific deployed contract.
func NewFiredrillCompoundFilterer(address common.Address, filterer bind.ContractFilterer) (*FiredrillCompoundFilterer, error) {
	contract, err := bindFiredrillCompound(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundFilterer{contract: contract}, nil
}

// bindFiredrillCompound binds a generic wrapper to an already deployed contract.
func bindFiredrillCompound(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiredrillCompoundMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillCompound *FiredrillCompoundRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillCompound.Contract.FiredrillCompoundCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillCompound *FiredrillCompoundRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.FiredrillCompoundTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillCompound *FiredrillCompoundRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.FiredrillCompoundTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillCompound *FiredrillCompoundCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillCompound.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillCompound *FiredrillCompoundTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillCompound *FiredrillCompoundTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.contract.Transact(opts, method, params...)
}

// GetOffRamps is a free data retrieval call binding the contract method 0xa40e69c7.
//
// Solidity: function getOffRamps() view returns((uint64,address)[])
func (_FiredrillCompound *FiredrillCompoundCaller) GetOffRamps(opts *bind.CallOpts) ([]FiredrillCompoundOffRamp, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getOffRamps")

	if err != nil {
		return *new([]FiredrillCompoundOffRamp), err
	}

	out0 := *abi.ConvertType(out[0], new([]FiredrillCompoundOffRamp)).(*[]FiredrillCompoundOffRamp)

	return out0, err

}

// GetOffRamps is a free data retrieval call binding the contract method 0xa40e69c7.
//
// Solidity: function getOffRamps() view returns((uint64,address)[])
func (_FiredrillCompound *FiredrillCompoundSession) GetOffRamps() ([]FiredrillCompoundOffRamp, error) {
	return _FiredrillCompound.Contract.GetOffRamps(&_FiredrillCompound.CallOpts)
}

// GetOffRamps is a free data retrieval call binding the contract method 0xa40e69c7.
//
// Solidity: function getOffRamps() view returns((uint64,address)[])
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetOffRamps() ([]FiredrillCompoundOffRamp, error) {
	return _FiredrillCompound.Contract.GetOffRamps(&_FiredrillCompound.CallOpts)
}

// GetOnRamp is a free data retrieval call binding the contract method 0xa8d87a3b.
//
// Solidity: function getOnRamp(uint64 ) view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) GetOnRamp(opts *bind.CallOpts, arg0 uint64) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getOnRamp", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOnRamp is a free data retrieval call binding the contract method 0xa8d87a3b.
//
// Solidity: function getOnRamp(uint64 ) view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _FiredrillCompound.Contract.GetOnRamp(&_FiredrillCompound.CallOpts, arg0)
}

// GetOnRamp is a free data retrieval call binding the contract method 0xa8d87a3b.
//
// Solidity: function getOnRamp(uint64 ) view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetOnRamp(arg0 uint64) (common.Address, error) {
	return _FiredrillCompound.Contract.GetOnRamp(&_FiredrillCompound.CallOpts, arg0)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint96,address,uint32))
func (_FiredrillCompound *FiredrillCompoundCaller) GetStaticConfig(opts *bind.CallOpts) (FiredrillCompoundStaticConfig, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FiredrillCompoundStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillCompoundStaticConfig)).(*FiredrillCompoundStaticConfig)

	return out0, err

}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint96,address,uint32))
func (_FiredrillCompound *FiredrillCompoundSession) GetStaticConfig() (FiredrillCompoundStaticConfig, error) {
	return _FiredrillCompound.Contract.GetStaticConfig(&_FiredrillCompound.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint96,address,uint32))
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetStaticConfig() (FiredrillCompoundStaticConfig, error) {
	return _FiredrillCompound.Contract.GetStaticConfig(&_FiredrillCompound.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_FiredrillCompound *FiredrillCompoundCaller) IsActive(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "isActive")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_FiredrillCompound *FiredrillCompoundSession) IsActive() (bool, error) {
	return _FiredrillCompound.Contract.IsActive(&_FiredrillCompound.CallOpts)
}

// IsActive is a free data retrieval call binding the contract method 0x22f3e2d4.
//
// Solidity: function isActive() view returns(bool)
func (_FiredrillCompound *FiredrillCompoundCallerSession) IsActive() (bool, error) {
	return _FiredrillCompound.Contract.IsActive(&_FiredrillCompound.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) Owner() (common.Address, error) {
	return _FiredrillCompound.Contract.Owner(&_FiredrillCompound.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) Owner() (common.Address, error) {
	return _FiredrillCompound.Contract.Owner(&_FiredrillCompound.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) PendingOwner() (common.Address, error) {
	return _FiredrillCompound.Contract.PendingOwner(&_FiredrillCompound.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) PendingOwner() (common.Address, error) {
	return _FiredrillCompound.Contract.PendingOwner(&_FiredrillCompound.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillCompound *FiredrillCompoundCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillCompound *FiredrillCompoundSession) TypeAndVersion() (string, error) {
	return _FiredrillCompound.Contract.TypeAndVersion(&_FiredrillCompound.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillCompound *FiredrillCompoundCallerSession) TypeAndVersion() (string, error) {
	return _FiredrillCompound.Contract.TypeAndVersion(&_FiredrillCompound.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.AcceptOwnership(&_FiredrillCompound.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.AcceptOwnership(&_FiredrillCompound.TransactOpts)
}

// EmitUsdPerTokenUpdated is a paid mutator transaction binding the contract method 0x20f938b6.
//
// Solidity: function emitUsdPerTokenUpdated() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) EmitUsdPerTokenUpdated(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "emitUsdPerTokenUpdated")
}

// EmitUsdPerTokenUpdated is a paid mutator transaction binding the contract method 0x20f938b6.
//
// Solidity: function emitUsdPerTokenUpdated() returns()
func (_FiredrillCompound *FiredrillCompoundSession) EmitUsdPerTokenUpdated() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitUsdPerTokenUpdated(&_FiredrillCompound.TransactOpts)
}

// EmitUsdPerTokenUpdated is a paid mutator transaction binding the contract method 0x20f938b6.
//
// Solidity: function emitUsdPerTokenUpdated() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) EmitUsdPerTokenUpdated() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitUsdPerTokenUpdated(&_FiredrillCompound.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.RenounceOwnership(&_FiredrillCompound.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.RenounceOwnership(&_FiredrillCompound.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillCompound *FiredrillCompoundSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.TransferOwnership(&_FiredrillCompound.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.TransferOwnership(&_FiredrillCompound.TransactOpts, newOwner)
}

// FiredrillCompoundOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FiredrillCompound contract.
type FiredrillCompoundOwnershipTransferStartedIterator struct {
	Event *FiredrillCompoundOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundOwnershipTransferStarted)
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
		it.Event = new(FiredrillCompoundOwnershipTransferStarted)
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
func (it *FiredrillCompoundOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FiredrillCompound contract.
type FiredrillCompoundOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillCompoundOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundOwnershipTransferStartedIterator{contract: _FiredrillCompound.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundOwnershipTransferStarted)
				if err := _FiredrillCompound.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseOwnershipTransferStarted(log types.Log) (*FiredrillCompoundOwnershipTransferStarted, error) {
	event := new(FiredrillCompoundOwnershipTransferStarted)
	if err := _FiredrillCompound.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillCompoundOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FiredrillCompound contract.
type FiredrillCompoundOwnershipTransferredIterator struct {
	Event *FiredrillCompoundOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundOwnershipTransferred)
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
		it.Event = new(FiredrillCompoundOwnershipTransferred)
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
func (it *FiredrillCompoundOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundOwnershipTransferred represents a OwnershipTransferred event raised by the FiredrillCompound contract.
type FiredrillCompoundOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillCompoundOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundOwnershipTransferredIterator{contract: _FiredrillCompound.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundOwnershipTransferred)
				if err := _FiredrillCompound.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseOwnershipTransferred(log types.Log) (*FiredrillCompoundOwnershipTransferred, error) {
	event := new(FiredrillCompoundOwnershipTransferred)
	if err := _FiredrillCompound.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillCompoundUsdPerTokenUpdatedIterator is returned from FilterUsdPerTokenUpdated and is used to iterate over the raw logs and unpacked data for UsdPerTokenUpdated events raised by the FiredrillCompound contract.
type FiredrillCompoundUsdPerTokenUpdatedIterator struct {
	Event *FiredrillCompoundUsdPerTokenUpdated // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundUsdPerTokenUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundUsdPerTokenUpdated)
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
		it.Event = new(FiredrillCompoundUsdPerTokenUpdated)
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
func (it *FiredrillCompoundUsdPerTokenUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundUsdPerTokenUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundUsdPerTokenUpdated represents a UsdPerTokenUpdated event raised by the FiredrillCompound contract.
type FiredrillCompoundUsdPerTokenUpdated struct {
	Token     common.Address
	Value     *big.Int
	Timestamp *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUsdPerTokenUpdated is a free log retrieval operation binding the contract event 0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a.
//
// Solidity: event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterUsdPerTokenUpdated(opts *bind.FilterOpts, token []common.Address) (*FiredrillCompoundUsdPerTokenUpdatedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundUsdPerTokenUpdatedIterator{contract: _FiredrillCompound.contract, event: "UsdPerTokenUpdated", logs: logs, sub: sub}, nil
}

// WatchUsdPerTokenUpdated is a free log subscription operation binding the contract event 0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a.
//
// Solidity: event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchUsdPerTokenUpdated(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundUsdPerTokenUpdated, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "UsdPerTokenUpdated", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundUsdPerTokenUpdated)
				if err := _FiredrillCompound.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
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

// ParseUsdPerTokenUpdated is a log parse operation binding the contract event 0x52f50aa6d1a95a4595361ecf953d095f125d442e4673716dede699e049de148a.
//
// Solidity: event UsdPerTokenUpdated(address indexed token, uint256 value, uint256 timestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseUsdPerTokenUpdated(log types.Log) (*FiredrillCompoundUsdPerTokenUpdated, error) {
	event := new(FiredrillCompoundUsdPerTokenUpdated)
	if err := _FiredrillCompound.contract.UnpackLog(event, "UsdPerTokenUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
