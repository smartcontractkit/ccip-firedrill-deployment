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

// FiredrillCompoundTaggedRoot is an auto generated low-level Go binding around an user-defined struct.
type FiredrillCompoundTaggedRoot struct {
	CommitStore common.Address
	Root        [32]byte
}

// FiredrillCompoundTokenBucket is an auto generated low-level Go binding around an user-defined struct.
type FiredrillCompoundTokenBucket struct {
	Tokens      *big.Int
	LastUpdated uint32
	IsEnabled   bool
	Capacity    *big.Int
	Rate        *big.Int
}

// FiredrillCompoundMetaData contains all meta data concerning the FiredrillCompound contract.
var FiredrillCompoundMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainSelector\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"currentRateLimiterState\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillCompound.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"emitCursed\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitReleased\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitTagRootBlessed\",\"inputs\":[{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitUsdPerTokenUpdated\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitVotedToCurse\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getARM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentInboundRateLimiterState\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillCompound.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getCurrentOutboundRateLimiterState\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillCompound.TokenBucket\",\"components\":[{\"name\":\"tokens\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"lastUpdated\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"capacity\",\"type\":\"uint128\",\"internalType\":\"uint128\"},{\"name\":\"rate\",\"type\":\"uint128\",\"internalType\":\"uint128\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOffRamps\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structFiredrillCompound.OffRamp[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offRamp\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getOnRamp\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStalenessThreshold\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint128\",\"internalType\":\"uint128\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"getToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"offRamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"onRamp\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiver\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"Cursed\",\"inputs\":[{\"name\":\"configVersion\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"subject\",\"type\":\"bytes16\",\"indexed\":false,\"internalType\":\"bytes16\"},{\"name\":\"blockTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Released\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TaggedRootBlessed\",\"inputs\":[{\"name\":\"configVersion\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"taggedRoot\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillCompound.TaggedRoot\",\"components\":[{\"name\":\"commitStore\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"root\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"accumulatedWeight\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"UsdPerTokenUpdated\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"timestamp\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"VotedToCurse\",\"inputs\":[{\"name\":\"configVersion\",\"type\":\"uint32\",\"indexed\":true,\"internalType\":\"uint32\"},{\"name\":\"voter\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"subject\",\"type\":\"bytes16\",\"indexed\":false,\"internalType\":\"bytes16\"},{\"name\":\"curseId\",\"type\":\"bytes16\",\"indexed\":false,\"internalType\":\"bytes16\"},{\"name\":\"weight\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"blockTimestamp\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"cursesHash\",\"type\":\"bytes28\",\"indexed\":false,\"internalType\":\"bytes28\"},{\"name\":\"accumulatedWeight\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// FiredrillCompoundABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillCompoundMetaData.ABI instead.
var FiredrillCompoundABI = FiredrillCompoundMetaData.ABI

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

// ChainSelector is a free data retrieval call binding the contract method 0x4e4bc847.
//
// Solidity: function chainSelector() view returns(uint64)
func (_FiredrillCompound *FiredrillCompoundCaller) ChainSelector(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "chainSelector")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainSelector is a free data retrieval call binding the contract method 0x4e4bc847.
//
// Solidity: function chainSelector() view returns(uint64)
func (_FiredrillCompound *FiredrillCompoundSession) ChainSelector() (uint64, error) {
	return _FiredrillCompound.Contract.ChainSelector(&_FiredrillCompound.CallOpts)
}

// ChainSelector is a free data retrieval call binding the contract method 0x4e4bc847.
//
// Solidity: function chainSelector() view returns(uint64)
func (_FiredrillCompound *FiredrillCompoundCallerSession) ChainSelector() (uint64, error) {
	return _FiredrillCompound.Contract.ChainSelector(&_FiredrillCompound.CallOpts)
}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCaller) CurrentRateLimiterState(opts *bind.CallOpts) (FiredrillCompoundTokenBucket, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "currentRateLimiterState")

	if err != nil {
		return *new(FiredrillCompoundTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillCompoundTokenBucket)).(*FiredrillCompoundTokenBucket)

	return out0, err

}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundSession) CurrentRateLimiterState() (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.CurrentRateLimiterState(&_FiredrillCompound.CallOpts)
}

// CurrentRateLimiterState is a free data retrieval call binding the contract method 0x546719cd.
//
// Solidity: function currentRateLimiterState() view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCallerSession) CurrentRateLimiterState() (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.CurrentRateLimiterState(&_FiredrillCompound.CallOpts)
}

// GetARM is a free data retrieval call binding the contract method 0x2e90aa21.
//
// Solidity: function getARM() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) GetARM(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getARM")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetARM is a free data retrieval call binding the contract method 0x2e90aa21.
//
// Solidity: function getARM() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) GetARM() (common.Address, error) {
	return _FiredrillCompound.Contract.GetARM(&_FiredrillCompound.CallOpts)
}

// GetARM is a free data retrieval call binding the contract method 0x2e90aa21.
//
// Solidity: function getARM() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetARM() (common.Address, error) {
	return _FiredrillCompound.Contract.GetARM(&_FiredrillCompound.CallOpts)
}

// GetCurrentInboundRateLimiterState is a free data retrieval call binding the contract method 0xaf58d59f.
//
// Solidity: function getCurrentInboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCaller) GetCurrentInboundRateLimiterState(opts *bind.CallOpts, arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getCurrentInboundRateLimiterState", arg0)

	if err != nil {
		return *new(FiredrillCompoundTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillCompoundTokenBucket)).(*FiredrillCompoundTokenBucket)

	return out0, err

}

// GetCurrentInboundRateLimiterState is a free data retrieval call binding the contract method 0xaf58d59f.
//
// Solidity: function getCurrentInboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundSession) GetCurrentInboundRateLimiterState(arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.GetCurrentInboundRateLimiterState(&_FiredrillCompound.CallOpts, arg0)
}

// GetCurrentInboundRateLimiterState is a free data retrieval call binding the contract method 0xaf58d59f.
//
// Solidity: function getCurrentInboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetCurrentInboundRateLimiterState(arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.GetCurrentInboundRateLimiterState(&_FiredrillCompound.CallOpts, arg0)
}

// GetCurrentOutboundRateLimiterState is a free data retrieval call binding the contract method 0xc75eea9c.
//
// Solidity: function getCurrentOutboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCaller) GetCurrentOutboundRateLimiterState(opts *bind.CallOpts, arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getCurrentOutboundRateLimiterState", arg0)

	if err != nil {
		return *new(FiredrillCompoundTokenBucket), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillCompoundTokenBucket)).(*FiredrillCompoundTokenBucket)

	return out0, err

}

// GetCurrentOutboundRateLimiterState is a free data retrieval call binding the contract method 0xc75eea9c.
//
// Solidity: function getCurrentOutboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundSession) GetCurrentOutboundRateLimiterState(arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.GetCurrentOutboundRateLimiterState(&_FiredrillCompound.CallOpts, arg0)
}

// GetCurrentOutboundRateLimiterState is a free data retrieval call binding the contract method 0xc75eea9c.
//
// Solidity: function getCurrentOutboundRateLimiterState(uint64 ) view returns((uint128,uint32,bool,uint128,uint128))
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetCurrentOutboundRateLimiterState(arg0 uint64) (FiredrillCompoundTokenBucket, error) {
	return _FiredrillCompound.Contract.GetCurrentOutboundRateLimiterState(&_FiredrillCompound.CallOpts, arg0)
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

// GetStalenessThreshold is a free data retrieval call binding the contract method 0xa6c94a73.
//
// Solidity: function getStalenessThreshold() pure returns(uint128)
func (_FiredrillCompound *FiredrillCompoundCaller) GetStalenessThreshold(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getStalenessThreshold")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStalenessThreshold is a free data retrieval call binding the contract method 0xa6c94a73.
//
// Solidity: function getStalenessThreshold() pure returns(uint128)
func (_FiredrillCompound *FiredrillCompoundSession) GetStalenessThreshold() (*big.Int, error) {
	return _FiredrillCompound.Contract.GetStalenessThreshold(&_FiredrillCompound.CallOpts)
}

// GetStalenessThreshold is a free data retrieval call binding the contract method 0xa6c94a73.
//
// Solidity: function getStalenessThreshold() pure returns(uint128)
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetStalenessThreshold() (*big.Int, error) {
	return _FiredrillCompound.Contract.GetStalenessThreshold(&_FiredrillCompound.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) GetToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "getToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) GetToken() (common.Address, error) {
	return _FiredrillCompound.Contract.GetToken(&_FiredrillCompound.CallOpts)
}

// GetToken is a free data retrieval call binding the contract method 0x21df0da7.
//
// Solidity: function getToken() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) GetToken() (common.Address, error) {
	return _FiredrillCompound.Contract.GetToken(&_FiredrillCompound.CallOpts)
}

// OffRamp is a free data retrieval call binding the contract method 0x44671a5d.
//
// Solidity: function offRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) OffRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "offRamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OffRamp is a free data retrieval call binding the contract method 0x44671a5d.
//
// Solidity: function offRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) OffRamp() (common.Address, error) {
	return _FiredrillCompound.Contract.OffRamp(&_FiredrillCompound.CallOpts)
}

// OffRamp is a free data retrieval call binding the contract method 0x44671a5d.
//
// Solidity: function offRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) OffRamp() (common.Address, error) {
	return _FiredrillCompound.Contract.OffRamp(&_FiredrillCompound.CallOpts)
}

// OnRamp is a free data retrieval call binding the contract method 0xc021e73c.
//
// Solidity: function onRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) OnRamp(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "onRamp")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OnRamp is a free data retrieval call binding the contract method 0xc021e73c.
//
// Solidity: function onRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) OnRamp() (common.Address, error) {
	return _FiredrillCompound.Contract.OnRamp(&_FiredrillCompound.CallOpts)
}

// OnRamp is a free data retrieval call binding the contract method 0xc021e73c.
//
// Solidity: function onRamp() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) OnRamp() (common.Address, error) {
	return _FiredrillCompound.Contract.OnRamp(&_FiredrillCompound.CallOpts)
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

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) Receiver(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "receiver")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) Receiver() (common.Address, error) {
	return _FiredrillCompound.Contract.Receiver(&_FiredrillCompound.CallOpts)
}

// Receiver is a free data retrieval call binding the contract method 0xf7260d3e.
//
// Solidity: function receiver() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) Receiver() (common.Address, error) {
	return _FiredrillCompound.Contract.Receiver(&_FiredrillCompound.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillCompound.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FiredrillCompound *FiredrillCompoundSession) Token() (common.Address, error) {
	return _FiredrillCompound.Contract.Token(&_FiredrillCompound.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_FiredrillCompound *FiredrillCompoundCallerSession) Token() (common.Address, error) {
	return _FiredrillCompound.Contract.Token(&_FiredrillCompound.CallOpts)
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

// EmitCursed is a paid mutator transaction binding the contract method 0xd7080f09.
//
// Solidity: function emitCursed() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) EmitCursed(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "emitCursed")
}

// EmitCursed is a paid mutator transaction binding the contract method 0xd7080f09.
//
// Solidity: function emitCursed() returns()
func (_FiredrillCompound *FiredrillCompoundSession) EmitCursed() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitCursed(&_FiredrillCompound.TransactOpts)
}

// EmitCursed is a paid mutator transaction binding the contract method 0xd7080f09.
//
// Solidity: function emitCursed() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) EmitCursed() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitCursed(&_FiredrillCompound.TransactOpts)
}

// EmitReleased is a paid mutator transaction binding the contract method 0xf5de8527.
//
// Solidity: function emitReleased() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) EmitReleased(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "emitReleased")
}

// EmitReleased is a paid mutator transaction binding the contract method 0xf5de8527.
//
// Solidity: function emitReleased() returns()
func (_FiredrillCompound *FiredrillCompoundSession) EmitReleased() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitReleased(&_FiredrillCompound.TransactOpts)
}

// EmitReleased is a paid mutator transaction binding the contract method 0xf5de8527.
//
// Solidity: function emitReleased() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) EmitReleased() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitReleased(&_FiredrillCompound.TransactOpts)
}

// EmitTagRootBlessed is a paid mutator transaction binding the contract method 0x43d24152.
//
// Solidity: function emitTagRootBlessed(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) EmitTagRootBlessed(opts *bind.TransactOpts, minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "emitTagRootBlessed", minSeqNr, maxSeqNr)
}

// EmitTagRootBlessed is a paid mutator transaction binding the contract method 0x43d24152.
//
// Solidity: function emitTagRootBlessed(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillCompound *FiredrillCompoundSession) EmitTagRootBlessed(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitTagRootBlessed(&_FiredrillCompound.TransactOpts, minSeqNr, maxSeqNr)
}

// EmitTagRootBlessed is a paid mutator transaction binding the contract method 0x43d24152.
//
// Solidity: function emitTagRootBlessed(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) EmitTagRootBlessed(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitTagRootBlessed(&_FiredrillCompound.TransactOpts, minSeqNr, maxSeqNr)
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

// EmitVotedToCurse is a paid mutator transaction binding the contract method 0x384b09e8.
//
// Solidity: function emitVotedToCurse() returns()
func (_FiredrillCompound *FiredrillCompoundTransactor) EmitVotedToCurse(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillCompound.contract.Transact(opts, "emitVotedToCurse")
}

// EmitVotedToCurse is a paid mutator transaction binding the contract method 0x384b09e8.
//
// Solidity: function emitVotedToCurse() returns()
func (_FiredrillCompound *FiredrillCompoundSession) EmitVotedToCurse() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitVotedToCurse(&_FiredrillCompound.TransactOpts)
}

// EmitVotedToCurse is a paid mutator transaction binding the contract method 0x384b09e8.
//
// Solidity: function emitVotedToCurse() returns()
func (_FiredrillCompound *FiredrillCompoundTransactorSession) EmitVotedToCurse() (*types.Transaction, error) {
	return _FiredrillCompound.Contract.EmitVotedToCurse(&_FiredrillCompound.TransactOpts)
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

// FiredrillCompoundCursedIterator is returned from FilterCursed and is used to iterate over the raw logs and unpacked data for Cursed events raised by the FiredrillCompound contract.
type FiredrillCompoundCursedIterator struct {
	Event *FiredrillCompoundCursed // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundCursedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundCursed)
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
		it.Event = new(FiredrillCompoundCursed)
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
func (it *FiredrillCompoundCursedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundCursedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundCursed represents a Cursed event raised by the FiredrillCompound contract.
type FiredrillCompoundCursed struct {
	ConfigVersion  uint32
	Subject        [16]byte
	BlockTimestamp uint64
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterCursed is a free log retrieval operation binding the contract event 0xcfdbfd8ce9a56b5f7c202c0e102184d24f47ca87121dc165063fc4c290957bde.
//
// Solidity: event Cursed(uint32 indexed configVersion, bytes16 subject, uint64 blockTimestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterCursed(opts *bind.FilterOpts, configVersion []uint32) (*FiredrillCompoundCursedIterator, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "Cursed", configVersionRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundCursedIterator{contract: _FiredrillCompound.contract, event: "Cursed", logs: logs, sub: sub}, nil
}

// WatchCursed is a free log subscription operation binding the contract event 0xcfdbfd8ce9a56b5f7c202c0e102184d24f47ca87121dc165063fc4c290957bde.
//
// Solidity: event Cursed(uint32 indexed configVersion, bytes16 subject, uint64 blockTimestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchCursed(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundCursed, configVersion []uint32) (event.Subscription, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "Cursed", configVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundCursed)
				if err := _FiredrillCompound.contract.UnpackLog(event, "Cursed", log); err != nil {
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

// ParseCursed is a log parse operation binding the contract event 0xcfdbfd8ce9a56b5f7c202c0e102184d24f47ca87121dc165063fc4c290957bde.
//
// Solidity: event Cursed(uint32 indexed configVersion, bytes16 subject, uint64 blockTimestamp)
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseCursed(log types.Log) (*FiredrillCompoundCursed, error) {
	event := new(FiredrillCompoundCursed)
	if err := _FiredrillCompound.contract.UnpackLog(event, "Cursed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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

// FiredrillCompoundReleasedIterator is returned from FilterReleased and is used to iterate over the raw logs and unpacked data for Released events raised by the FiredrillCompound contract.
type FiredrillCompoundReleasedIterator struct {
	Event *FiredrillCompoundReleased // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundReleasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundReleased)
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
		it.Event = new(FiredrillCompoundReleased)
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
func (it *FiredrillCompoundReleasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundReleasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundReleased represents a Released event raised by the FiredrillCompound contract.
type FiredrillCompoundReleased struct {
	Sender    common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterReleased is a free log retrieval operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed sender, address indexed recipient, uint256 amount)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterReleased(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address) (*FiredrillCompoundReleasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundReleasedIterator{contract: _FiredrillCompound.contract, event: "Released", logs: logs, sub: sub}, nil
}

// WatchReleased is a free log subscription operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed sender, address indexed recipient, uint256 amount)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchReleased(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundReleased, sender []common.Address, recipient []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "Released", senderRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundReleased)
				if err := _FiredrillCompound.contract.UnpackLog(event, "Released", log); err != nil {
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

// ParseReleased is a log parse operation binding the contract event 0x2d87480f50083e2b2759522a8fdda59802650a8055e609a7772cf70c07748f52.
//
// Solidity: event Released(address indexed sender, address indexed recipient, uint256 amount)
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseReleased(log types.Log) (*FiredrillCompoundReleased, error) {
	event := new(FiredrillCompoundReleased)
	if err := _FiredrillCompound.contract.UnpackLog(event, "Released", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillCompoundTaggedRootBlessedIterator is returned from FilterTaggedRootBlessed and is used to iterate over the raw logs and unpacked data for TaggedRootBlessed events raised by the FiredrillCompound contract.
type FiredrillCompoundTaggedRootBlessedIterator struct {
	Event *FiredrillCompoundTaggedRootBlessed // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundTaggedRootBlessedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundTaggedRootBlessed)
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
		it.Event = new(FiredrillCompoundTaggedRootBlessed)
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
func (it *FiredrillCompoundTaggedRootBlessedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundTaggedRootBlessedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundTaggedRootBlessed represents a TaggedRootBlessed event raised by the FiredrillCompound contract.
type FiredrillCompoundTaggedRootBlessed struct {
	ConfigVersion     uint32
	TaggedRoot        FiredrillCompoundTaggedRoot
	AccumulatedWeight uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterTaggedRootBlessed is a free log retrieval operation binding the contract event 0x8257378aa73bf8e4ada848713526584a3dcee0fd3db3beed7397f7a7f5067cc9.
//
// Solidity: event TaggedRootBlessed(uint32 indexed configVersion, (address,bytes32) taggedRoot, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterTaggedRootBlessed(opts *bind.FilterOpts, configVersion []uint32) (*FiredrillCompoundTaggedRootBlessedIterator, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "TaggedRootBlessed", configVersionRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundTaggedRootBlessedIterator{contract: _FiredrillCompound.contract, event: "TaggedRootBlessed", logs: logs, sub: sub}, nil
}

// WatchTaggedRootBlessed is a free log subscription operation binding the contract event 0x8257378aa73bf8e4ada848713526584a3dcee0fd3db3beed7397f7a7f5067cc9.
//
// Solidity: event TaggedRootBlessed(uint32 indexed configVersion, (address,bytes32) taggedRoot, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchTaggedRootBlessed(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundTaggedRootBlessed, configVersion []uint32) (event.Subscription, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "TaggedRootBlessed", configVersionRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundTaggedRootBlessed)
				if err := _FiredrillCompound.contract.UnpackLog(event, "TaggedRootBlessed", log); err != nil {
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

// ParseTaggedRootBlessed is a log parse operation binding the contract event 0x8257378aa73bf8e4ada848713526584a3dcee0fd3db3beed7397f7a7f5067cc9.
//
// Solidity: event TaggedRootBlessed(uint32 indexed configVersion, (address,bytes32) taggedRoot, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseTaggedRootBlessed(log types.Log) (*FiredrillCompoundTaggedRootBlessed, error) {
	event := new(FiredrillCompoundTaggedRootBlessed)
	if err := _FiredrillCompound.contract.UnpackLog(event, "TaggedRootBlessed", log); err != nil {
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

// FiredrillCompoundVotedToCurseIterator is returned from FilterVotedToCurse and is used to iterate over the raw logs and unpacked data for VotedToCurse events raised by the FiredrillCompound contract.
type FiredrillCompoundVotedToCurseIterator struct {
	Event *FiredrillCompoundVotedToCurse // Event containing the contract specifics and raw log

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
func (it *FiredrillCompoundVotedToCurseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillCompoundVotedToCurse)
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
		it.Event = new(FiredrillCompoundVotedToCurse)
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
func (it *FiredrillCompoundVotedToCurseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillCompoundVotedToCurseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillCompoundVotedToCurse represents a VotedToCurse event raised by the FiredrillCompound contract.
type FiredrillCompoundVotedToCurse struct {
	ConfigVersion     uint32
	Voter             common.Address
	Subject           [16]byte
	CurseId           [16]byte
	Weight            uint8
	BlockTimestamp    uint64
	CursesHash        [28]byte
	AccumulatedWeight uint16
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterVotedToCurse is a free log retrieval operation binding the contract event 0x8137bc8a8d712aaa27bfc6506d5566ac405618bd53f9831b8ca6b6fe5442ee7a.
//
// Solidity: event VotedToCurse(uint32 indexed configVersion, address indexed voter, bytes16 subject, bytes16 curseId, uint8 weight, uint64 blockTimestamp, bytes28 cursesHash, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) FilterVotedToCurse(opts *bind.FilterOpts, configVersion []uint32, voter []common.Address) (*FiredrillCompoundVotedToCurseIterator, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FiredrillCompound.contract.FilterLogs(opts, "VotedToCurse", configVersionRule, voterRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillCompoundVotedToCurseIterator{contract: _FiredrillCompound.contract, event: "VotedToCurse", logs: logs, sub: sub}, nil
}

// WatchVotedToCurse is a free log subscription operation binding the contract event 0x8137bc8a8d712aaa27bfc6506d5566ac405618bd53f9831b8ca6b6fe5442ee7a.
//
// Solidity: event VotedToCurse(uint32 indexed configVersion, address indexed voter, bytes16 subject, bytes16 curseId, uint8 weight, uint64 blockTimestamp, bytes28 cursesHash, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) WatchVotedToCurse(opts *bind.WatchOpts, sink chan<- *FiredrillCompoundVotedToCurse, configVersion []uint32, voter []common.Address) (event.Subscription, error) {

	var configVersionRule []interface{}
	for _, configVersionItem := range configVersion {
		configVersionRule = append(configVersionRule, configVersionItem)
	}
	var voterRule []interface{}
	for _, voterItem := range voter {
		voterRule = append(voterRule, voterItem)
	}

	logs, sub, err := _FiredrillCompound.contract.WatchLogs(opts, "VotedToCurse", configVersionRule, voterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillCompoundVotedToCurse)
				if err := _FiredrillCompound.contract.UnpackLog(event, "VotedToCurse", log); err != nil {
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

// ParseVotedToCurse is a log parse operation binding the contract event 0x8137bc8a8d712aaa27bfc6506d5566ac405618bd53f9831b8ca6b6fe5442ee7a.
//
// Solidity: event VotedToCurse(uint32 indexed configVersion, address indexed voter, bytes16 subject, bytes16 curseId, uint8 weight, uint64 blockTimestamp, bytes28 cursesHash, uint16 accumulatedWeight)
func (_FiredrillCompound *FiredrillCompoundFilterer) ParseVotedToCurse(log types.Log) (*FiredrillCompoundVotedToCurse, error) {
	event := new(FiredrillCompoundVotedToCurse)
	if err := _FiredrillCompound.contract.UnpackLog(event, "VotedToCurse", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
