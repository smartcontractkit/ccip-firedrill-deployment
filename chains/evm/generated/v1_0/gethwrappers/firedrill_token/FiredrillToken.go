// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package firedrill_token

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

// FiredrillTokenMetaData contains all meta data concerning the FiredrillToken contract.
var FiredrillTokenMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b50338061003557604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61003e81610044565b506100af565b600180546001600160a01b031916905561005d81610060565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b610389806100bc5f395ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c806379ba50971161006357806379ba50971461010f5780638da5cb5b1461011757806395d89b411461013b578063e30c39781461015b578063f2fde38b1461016c575f80fd5b8063181f5a7714610094578063313ce567146100d657806370a08231146100e4578063715018a614610105575b5f80fd5b6040805180820190915260148152730466972656472696c6c546f6b656e20312e302e360641b60208201525b6040516100cd91906102da565b60405180910390f35b6040515f81526020016100cd565b6100f76100f2366004610326565b505f90565b6040519081526020016100cd565b61010d61017f565b005b61010d610192565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100cd565b6040805180820190915260048152634c494e4b60e01b60208201526100c0565b6001546001600160a01b0316610123565b61010d61017a366004610326565b6101db565b61018761024b565b6101905f610277565b565b60015433906001600160a01b031681146101cf5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b6101d881610277565b50565b6101e361024b565b600180546001600160a01b0383166001600160a01b031990911681179091556102135f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b031633146101905760405163118cdaa760e01b81523360048201526024016101c6565b600180546001600160a01b03191690556101d8815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602080835283518060208501525f5b81811015610306578581018301518582016040015282016102ea565b505f604082860101526040601f19601f8301168501019250505092915050565b5f60208284031215610336575f80fd5b81356001600160a01b038116811461034c575f80fd5b939250505056fea264697066735822122073c9e3b0939622ca5302d370413bfe59d14db6fa7a6dc99daafdc9d33c3205dc64736f6c63430008180033",
}

// FiredrillTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillTokenMetaData.ABI instead.
var FiredrillTokenABI = FiredrillTokenMetaData.ABI

// FiredrillTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiredrillTokenMetaData.Bin instead.
var FiredrillTokenBin = FiredrillTokenMetaData.Bin

// DeployFiredrillToken deploys a new Ethereum contract, binding an instance of FiredrillToken to it.
func DeployFiredrillToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FiredrillToken, error) {
	parsed, err := FiredrillTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiredrillTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiredrillToken{FiredrillTokenCaller: FiredrillTokenCaller{contract: contract}, FiredrillTokenTransactor: FiredrillTokenTransactor{contract: contract}, FiredrillTokenFilterer: FiredrillTokenFilterer{contract: contract}}, nil
}

// FiredrillToken is an auto generated Go binding around an Ethereum contract.
type FiredrillToken struct {
	FiredrillTokenCaller     // Read-only binding to the contract
	FiredrillTokenTransactor // Write-only binding to the contract
	FiredrillTokenFilterer   // Log filterer for contract events
}

// FiredrillTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiredrillTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiredrillTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiredrillTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiredrillTokenSession struct {
	Contract     *FiredrillToken   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiredrillTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiredrillTokenCallerSession struct {
	Contract *FiredrillTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// FiredrillTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiredrillTokenTransactorSession struct {
	Contract     *FiredrillTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// FiredrillTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiredrillTokenRaw struct {
	Contract *FiredrillToken // Generic contract binding to access the raw methods on
}

// FiredrillTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiredrillTokenCallerRaw struct {
	Contract *FiredrillTokenCaller // Generic read-only contract binding to access the raw methods on
}

// FiredrillTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiredrillTokenTransactorRaw struct {
	Contract *FiredrillTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiredrillToken creates a new instance of FiredrillToken, bound to a specific deployed contract.
func NewFiredrillToken(address common.Address, backend bind.ContractBackend) (*FiredrillToken, error) {
	contract, err := bindFiredrillToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiredrillToken{FiredrillTokenCaller: FiredrillTokenCaller{contract: contract}, FiredrillTokenTransactor: FiredrillTokenTransactor{contract: contract}, FiredrillTokenFilterer: FiredrillTokenFilterer{contract: contract}}, nil
}

// NewFiredrillTokenCaller creates a new read-only instance of FiredrillToken, bound to a specific deployed contract.
func NewFiredrillTokenCaller(address common.Address, caller bind.ContractCaller) (*FiredrillTokenCaller, error) {
	contract, err := bindFiredrillToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillTokenCaller{contract: contract}, nil
}

// NewFiredrillTokenTransactor creates a new write-only instance of FiredrillToken, bound to a specific deployed contract.
func NewFiredrillTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*FiredrillTokenTransactor, error) {
	contract, err := bindFiredrillToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillTokenTransactor{contract: contract}, nil
}

// NewFiredrillTokenFilterer creates a new log filterer instance of FiredrillToken, bound to a specific deployed contract.
func NewFiredrillTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*FiredrillTokenFilterer, error) {
	contract, err := bindFiredrillToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiredrillTokenFilterer{contract: contract}, nil
}

// bindFiredrillToken binds a generic wrapper to an already deployed contract.
func bindFiredrillToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiredrillTokenMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillToken *FiredrillTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillToken.Contract.FiredrillTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillToken *FiredrillTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillToken.Contract.FiredrillTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillToken *FiredrillTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillToken.Contract.FiredrillTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillToken *FiredrillTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillToken *FiredrillTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillToken *FiredrillTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_FiredrillToken *FiredrillTokenCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_FiredrillToken *FiredrillTokenSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FiredrillToken.Contract.BalanceOf(&_FiredrillToken.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) pure returns(uint256)
func (_FiredrillToken *FiredrillTokenCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _FiredrillToken.Contract.BalanceOf(&_FiredrillToken.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_FiredrillToken *FiredrillTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_FiredrillToken *FiredrillTokenSession) Decimals() (uint8, error) {
	return _FiredrillToken.Contract.Decimals(&_FiredrillToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_FiredrillToken *FiredrillTokenCallerSession) Decimals() (uint8, error) {
	return _FiredrillToken.Contract.Decimals(&_FiredrillToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillToken *FiredrillTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillToken *FiredrillTokenSession) Owner() (common.Address, error) {
	return _FiredrillToken.Contract.Owner(&_FiredrillToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillToken *FiredrillTokenCallerSession) Owner() (common.Address, error) {
	return _FiredrillToken.Contract.Owner(&_FiredrillToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillToken *FiredrillTokenCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillToken *FiredrillTokenSession) PendingOwner() (common.Address, error) {
	return _FiredrillToken.Contract.PendingOwner(&_FiredrillToken.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillToken *FiredrillTokenCallerSession) PendingOwner() (common.Address, error) {
	return _FiredrillToken.Contract.PendingOwner(&_FiredrillToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_FiredrillToken *FiredrillTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_FiredrillToken *FiredrillTokenSession) Symbol() (string, error) {
	return _FiredrillToken.Contract.Symbol(&_FiredrillToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_FiredrillToken *FiredrillTokenCallerSession) Symbol() (string, error) {
	return _FiredrillToken.Contract.Symbol(&_FiredrillToken.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillToken *FiredrillTokenCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillToken.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillToken *FiredrillTokenSession) TypeAndVersion() (string, error) {
	return _FiredrillToken.Contract.TypeAndVersion(&_FiredrillToken.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillToken *FiredrillTokenCallerSession) TypeAndVersion() (string, error) {
	return _FiredrillToken.Contract.TypeAndVersion(&_FiredrillToken.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillToken *FiredrillTokenTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillToken.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillToken *FiredrillTokenSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillToken.Contract.AcceptOwnership(&_FiredrillToken.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillToken *FiredrillTokenTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillToken.Contract.AcceptOwnership(&_FiredrillToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillToken *FiredrillTokenTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillToken.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillToken *FiredrillTokenSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillToken.Contract.RenounceOwnership(&_FiredrillToken.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillToken *FiredrillTokenTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillToken.Contract.RenounceOwnership(&_FiredrillToken.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillToken *FiredrillTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillToken *FiredrillTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillToken.Contract.TransferOwnership(&_FiredrillToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillToken *FiredrillTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillToken.Contract.TransferOwnership(&_FiredrillToken.TransactOpts, newOwner)
}

// FiredrillTokenOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FiredrillToken contract.
type FiredrillTokenOwnershipTransferStartedIterator struct {
	Event *FiredrillTokenOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FiredrillTokenOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillTokenOwnershipTransferStarted)
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
		it.Event = new(FiredrillTokenOwnershipTransferStarted)
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
func (it *FiredrillTokenOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillTokenOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillTokenOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FiredrillToken contract.
type FiredrillTokenOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillToken *FiredrillTokenFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillTokenOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillToken.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillTokenOwnershipTransferStartedIterator{contract: _FiredrillToken.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillToken *FiredrillTokenFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FiredrillTokenOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillToken.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillTokenOwnershipTransferStarted)
				if err := _FiredrillToken.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FiredrillToken *FiredrillTokenFilterer) ParseOwnershipTransferStarted(log types.Log) (*FiredrillTokenOwnershipTransferStarted, error) {
	event := new(FiredrillTokenOwnershipTransferStarted)
	if err := _FiredrillToken.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FiredrillToken contract.
type FiredrillTokenOwnershipTransferredIterator struct {
	Event *FiredrillTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FiredrillTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillTokenOwnershipTransferred)
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
		it.Event = new(FiredrillTokenOwnershipTransferred)
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
func (it *FiredrillTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillTokenOwnershipTransferred represents a OwnershipTransferred event raised by the FiredrillToken contract.
type FiredrillTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillToken *FiredrillTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillTokenOwnershipTransferredIterator{contract: _FiredrillToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillToken *FiredrillTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiredrillTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillTokenOwnershipTransferred)
				if err := _FiredrillToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FiredrillToken *FiredrillTokenFilterer) ParseOwnershipTransferred(log types.Log) (*FiredrillTokenOwnershipTransferred, error) {
	event := new(FiredrillTokenOwnershipTransferred)
	if err := _FiredrillToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
