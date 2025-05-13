// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package firedrill_on_ramp

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

// FiredrillOnRampDynamicConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampDynamicConfig struct {
	Router                            common.Address
	MaxNumberOfTokensPerMsg           uint16
	DestGasOverhead                   uint32
	DestGasPerPayloadByte             uint16
	DestDataAvailabilityOverheadGas   uint32
	DestGasPerDataAvailabilityByte    uint16
	DestDataAvailabilityMultiplierBps uint16
	PriceRegistry                     common.Address
	MaxDataBytes                      uint32
	MaxPerMsgGasLimit                 uint32
	DefaultTokenFeeUSDCents           uint16
	DefaultTokenDestGasOverhead       uint32
	DefaultTokenDestBytesOverhead     uint32
	EnforceOutOfOrder                 bool
}

// FiredrillOnRampEVM2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampEVM2EVMMessage struct {
	SourceChainSelector uint64
	Sender              common.Address
	Receiver            common.Address
	SequenceNumber      uint64
	GasLimit            *big.Int
	Strict              bool
	Nonce               uint64
	FeeToken            common.Address
	FeeTokenAmount      *big.Int
	Data                []byte
	TokenAmounts        []FiredrillOnRampEVMTokenAmount
	SourceTokenData     [][]byte
	MessageId           [32]byte
}

// FiredrillOnRampEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// FiredrillOnRampStaticConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampStaticConfig struct {
	LinkToken          common.Address
	ChainSelector      uint64
	DestChainSelector  uint64
	DefaultTxGasLimit  uint64
	MaxNopFeesJuels    *big.Int
	PrevOnRamp         common.Address
	RmnProxy           common.Address
	TokenAdminRegistry common.Address
}

// FiredrillOnRampMetaData contains all meta data concerning the FiredrillOnRamp contract.
var FiredrillOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ctrl\",\"type\":\"address\",\"internalType\":\"contractFiredrillCompound\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitCCIPSendRequested\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"dynamicConfig\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOnRamp.DynamicConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxNumberOfTokensPerMsg\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerPayloadByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityOverheadGas\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"destGasPerDataAvailabilityByte\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"destDataAvailabilityMultiplierBps\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"priceRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxDataBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"maxPerMsgGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTokenFeeUSDCents\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"defaultTokenDestGasOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"defaultTokenDestBytesOverhead\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"enforceOutOfOrder\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOnRamp.StaticConfig\",\"components\":[{\"name\":\"linkToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"defaultTxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxNopFeesJuels\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"prevOnRamp\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"rmnProxy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CCIPSendRequested\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOnRamp.EVM2EVMMessage\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"strict\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structFiredrillOnRamp.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"sourceTokenData\",\"type\":\"bytes[]\",\"internalType\":\"bytes[]\"},{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610f85380380610f8583398101604081905261002e916100d9565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161006e565b506001600160a01b0316608052610106565b600180546001600160a01b03191690556100878161008a565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e9575f80fd5b81516001600160a01b03811681146100ff575f80fd5b9392505050565b608051610e3661014f5f395f8181610186015281816102b701528181610345015281816103d401528181610474015281816104b90152818161054901526106270152610e365ff3fe608060405234801561000f575f80fd5b5060043610610090575f3560e01c80637437ff9f116100635780637437ff9f146100fd57806379ba5097146102195780638da5cb5b14610221578063e30c397814610245578063f2fde38b14610256575f80fd5b806306285c6914610094578063181f5a77146100b257806351a2be12146100e0578063715018a6146100f5575b5f80fd5b61009c610269565b6040516100a9919061093c565b60405180910390f35b604080518082018252600c81526b04f6e52616d7020312e352e360a41b602082015290516100a99190610a32565b6100f36100ee366004610a74565b6104ae565b005b6100f36107e6565b61020c604080516101c0810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101829052610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915250604080516101c0810182526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168082525f60208301819052928201839052606082018390526080820183905260a0820183905260c0820183905260e0820152610100810182905261012081018290526101408101829052610160810182905261018081018290526101a081019190915290565b6040516100a99190610aab565b6100f36107f9565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100a9565b6001546001600160a01b031661022d565b6100f3610264366004610bc1565b61083d565b60408051610100810182525f80825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e08101919091526040518061010001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610311573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103359190610bdc565b6001600160a01b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa15801561039f573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103c39190610bf7565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa15801561042e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104529190610bf7565b67ffffffffffffffff1681525f602082018190526040820181905260608201527f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166080820181905260a090910152919050565b336001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000161461051a5760405162461bcd60e51b815260206004820152600c60248201526b1bdb9b1e4818dbdb9d1c9bdb60a21b60448201526064015b60405180910390fd5b7fd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd604051806101a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105a3573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105c79190610bf7565b67ffffffffffffffff168152602001846001600160a01b03168152602001846001600160a01b031681526020018367ffffffffffffffff1681526020016103e881526020015f15158152602001600167ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610681573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106a59190610bdc565b6001600160a01b031681526020015f81526020016040518060400160405280600381526020016231323360e81b81525081526020015f67ffffffffffffffff8111156106f3576106f3610c12565b60405190808252806020026020018201604052801561073757816020015b604080518082019091525f80825260208201528152602001906001900390816107115790505b5081526020015f60405190808252806020026020018201604052801561077157816020015b606081526020019060019003908161075c5790505b50815260200184846040516020016107b492919060609290921b6bffffffffffffffffffffffff1916825260c01b6001600160c01b0319166014820152601c0190565b604051602081830303815290604052805190602001208152506040516107da9190610ccc565b60405180910390a15050565b6107ee6108ad565b6107f75f6108d9565b565b60015433906001600160a01b031681146108315760405163118cdaa760e01b81526001600160a01b0382166004820152602401610511565b61083a816108d9565b50565b6108456108ad565b600180546001600160a01b0383166001600160a01b031990911681179091556108755f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b031633146107f75760405163118cdaa760e01b8152336004820152602401610511565b600180546001600160a01b031916905561083a815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f6101008201905060018060a01b038351168252602083015167ffffffffffffffff808216602085015280604086015116604085015280606086015116606085015250506bffffffffffffffffffffffff608084015116608083015260a08301516109b260a08401826001600160a01b03169052565b5060c08301516109cd60c08401826001600160a01b03169052565b5060e08301516109e860e08401826001600160a01b03169052565b5092915050565b5f81518084525f5b81811015610a13576020818501810151868301820152016109f7565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610a4460208301846109ef565b9392505050565b6001600160a01b038116811461083a575f80fd5b67ffffffffffffffff8116811461083a575f80fd5b5f8060408385031215610a85575f80fd5b8235610a9081610a4b565b91506020830135610aa081610a5f565b809150509250929050565b81516001600160a01b031681526101c081016020830151610ad2602084018261ffff169052565b506040830151610aea604084018263ffffffff169052565b506060830151610b00606084018261ffff169052565b506080830151610b18608084018263ffffffff169052565b5060a0830151610b2e60a084018261ffff169052565b5060c0830151610b4460c084018261ffff169052565b5060e0830151610b5f60e08401826001600160a01b03169052565b506101008381015163ffffffff90811691840191909152610120808501518216908401526101408085015161ffff16908401526101608085015182169084015261018080850151909116908301526101a0928301511515929091019190915290565b5f60208284031215610bd1575f80fd5b8135610a4481610a4b565b5f60208284031215610bec575f80fd5b8151610a4481610a4b565b5f60208284031215610c07575f80fd5b8151610a4481610a5f565b634e487b7160e01b5f52604160045260245ffd5b5f815180845260208085019450602084015f5b83811015610c6957815180516001600160a01b031688528301518388015260409096019590820190600101610c39565b509495945050505050565b5f8282518085526020808601955060208260051b840101602086015f5b84811015610cbf57601f19868403018952610cad8383516109ef565b98840198925090830190600101610c91565b5090979650505050505050565b60208152610ce760208201835167ffffffffffffffff169052565b5f6020830151610d0260408401826001600160a01b03169052565b5060408301516001600160a01b038116606084015250606083015167ffffffffffffffff8116608084015250608083015160a083015260a0830151610d4b60c084018215159052565b5060c083015167ffffffffffffffff811660e08401525060e0830151610100610d7e818501836001600160a01b03169052565b840151610120848101919091528401516101a061014080860182905291925090610dac6101c08601846109ef565b9250808601519050601f19610160818786030181880152610dcd8584610c26565b945080880151925050610180818786030181880152610dec8584610c74565b97015195909201949094525092939250505056fea2646970667358221220f7fd235aa8fa6814c7c4c236f9466661b5737df483a6f3c5030b6bc5001b002c64736f6c63430008180033",
}

// FiredrillOnRampABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillOnRampMetaData.ABI instead.
var FiredrillOnRampABI = FiredrillOnRampMetaData.ABI

// FiredrillOnRampBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiredrillOnRampMetaData.Bin instead.
var FiredrillOnRampBin = FiredrillOnRampMetaData.Bin

// DeployFiredrillOnRamp deploys a new Ethereum contract, binding an instance of FiredrillOnRamp to it.
func DeployFiredrillOnRamp(auth *bind.TransactOpts, backend bind.ContractBackend, ctrl common.Address) (common.Address, *types.Transaction, *FiredrillOnRamp, error) {
	parsed, err := FiredrillOnRampMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiredrillOnRampBin), backend, ctrl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiredrillOnRamp{FiredrillOnRampCaller: FiredrillOnRampCaller{contract: contract}, FiredrillOnRampTransactor: FiredrillOnRampTransactor{contract: contract}, FiredrillOnRampFilterer: FiredrillOnRampFilterer{contract: contract}}, nil
}

// FiredrillOnRamp is an auto generated Go binding around an Ethereum contract.
type FiredrillOnRamp struct {
	FiredrillOnRampCaller     // Read-only binding to the contract
	FiredrillOnRampTransactor // Write-only binding to the contract
	FiredrillOnRampFilterer   // Log filterer for contract events
}

// FiredrillOnRampCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiredrillOnRampCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOnRampTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiredrillOnRampTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOnRampFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiredrillOnRampFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillOnRampSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiredrillOnRampSession struct {
	Contract     *FiredrillOnRamp  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FiredrillOnRampCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiredrillOnRampCallerSession struct {
	Contract *FiredrillOnRampCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// FiredrillOnRampTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiredrillOnRampTransactorSession struct {
	Contract     *FiredrillOnRampTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// FiredrillOnRampRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiredrillOnRampRaw struct {
	Contract *FiredrillOnRamp // Generic contract binding to access the raw methods on
}

// FiredrillOnRampCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiredrillOnRampCallerRaw struct {
	Contract *FiredrillOnRampCaller // Generic read-only contract binding to access the raw methods on
}

// FiredrillOnRampTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiredrillOnRampTransactorRaw struct {
	Contract *FiredrillOnRampTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiredrillOnRamp creates a new instance of FiredrillOnRamp, bound to a specific deployed contract.
func NewFiredrillOnRamp(address common.Address, backend bind.ContractBackend) (*FiredrillOnRamp, error) {
	contract, err := bindFiredrillOnRamp(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRamp{FiredrillOnRampCaller: FiredrillOnRampCaller{contract: contract}, FiredrillOnRampTransactor: FiredrillOnRampTransactor{contract: contract}, FiredrillOnRampFilterer: FiredrillOnRampFilterer{contract: contract}}, nil
}

// NewFiredrillOnRampCaller creates a new read-only instance of FiredrillOnRamp, bound to a specific deployed contract.
func NewFiredrillOnRampCaller(address common.Address, caller bind.ContractCaller) (*FiredrillOnRampCaller, error) {
	contract, err := bindFiredrillOnRamp(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampCaller{contract: contract}, nil
}

// NewFiredrillOnRampTransactor creates a new write-only instance of FiredrillOnRamp, bound to a specific deployed contract.
func NewFiredrillOnRampTransactor(address common.Address, transactor bind.ContractTransactor) (*FiredrillOnRampTransactor, error) {
	contract, err := bindFiredrillOnRamp(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampTransactor{contract: contract}, nil
}

// NewFiredrillOnRampFilterer creates a new log filterer instance of FiredrillOnRamp, bound to a specific deployed contract.
func NewFiredrillOnRampFilterer(address common.Address, filterer bind.ContractFilterer) (*FiredrillOnRampFilterer, error) {
	contract, err := bindFiredrillOnRamp(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampFilterer{contract: contract}, nil
}

// bindFiredrillOnRamp binds a generic wrapper to an already deployed contract.
func bindFiredrillOnRamp(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiredrillOnRampMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillOnRamp *FiredrillOnRampRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillOnRamp.Contract.FiredrillOnRampCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillOnRamp *FiredrillOnRampRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.FiredrillOnRampTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillOnRamp *FiredrillOnRampRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.FiredrillOnRampTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillOnRamp *FiredrillOnRampCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillOnRamp.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillOnRamp *FiredrillOnRampTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillOnRamp *FiredrillOnRampTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.contract.Transact(opts, method, params...)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,uint16,uint32,uint16,uint32,uint16,uint16,address,uint32,uint32,uint16,uint32,uint32,bool) dynamicConfig)
func (_FiredrillOnRamp *FiredrillOnRampCaller) GetDynamicConfig(opts *bind.CallOpts) (FiredrillOnRampDynamicConfig, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "getDynamicConfig")

	if err != nil {
		return *new(FiredrillOnRampDynamicConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOnRampDynamicConfig)).(*FiredrillOnRampDynamicConfig)

	return out0, err

}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,uint16,uint32,uint16,uint32,uint16,uint16,address,uint32,uint32,uint16,uint32,uint32,bool) dynamicConfig)
func (_FiredrillOnRamp *FiredrillOnRampSession) GetDynamicConfig() (FiredrillOnRampDynamicConfig, error) {
	return _FiredrillOnRamp.Contract.GetDynamicConfig(&_FiredrillOnRamp.CallOpts)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,uint16,uint32,uint16,uint32,uint16,uint16,address,uint32,uint32,uint16,uint32,uint32,bool) dynamicConfig)
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) GetDynamicConfig() (FiredrillOnRampDynamicConfig, error) {
	return _FiredrillOnRamp.Contract.GetDynamicConfig(&_FiredrillOnRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,uint64,uint96,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampCaller) GetStaticConfig(opts *bind.CallOpts) (FiredrillOnRampStaticConfig, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "getStaticConfig")

	if err != nil {
		return *new(FiredrillOnRampStaticConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOnRampStaticConfig)).(*FiredrillOnRampStaticConfig)

	return out0, err

}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,uint64,uint96,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampSession) GetStaticConfig() (FiredrillOnRampStaticConfig, error) {
	return _FiredrillOnRamp.Contract.GetStaticConfig(&_FiredrillOnRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((address,uint64,uint64,uint64,uint96,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) GetStaticConfig() (FiredrillOnRampStaticConfig, error) {
	return _FiredrillOnRamp.Contract.GetStaticConfig(&_FiredrillOnRamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampSession) Owner() (common.Address, error) {
	return _FiredrillOnRamp.Contract.Owner(&_FiredrillOnRamp.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) Owner() (common.Address, error) {
	return _FiredrillOnRamp.Contract.Owner(&_FiredrillOnRamp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampSession) PendingOwner() (common.Address, error) {
	return _FiredrillOnRamp.Contract.PendingOwner(&_FiredrillOnRamp.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) PendingOwner() (common.Address, error) {
	return _FiredrillOnRamp.Contract.PendingOwner(&_FiredrillOnRamp.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOnRamp *FiredrillOnRampCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOnRamp *FiredrillOnRampSession) TypeAndVersion() (string, error) {
	return _FiredrillOnRamp.Contract.TypeAndVersion(&_FiredrillOnRamp.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) TypeAndVersion() (string, error) {
	return _FiredrillOnRamp.Contract.TypeAndVersion(&_FiredrillOnRamp.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOnRamp.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.AcceptOwnership(&_FiredrillOnRamp.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.AcceptOwnership(&_FiredrillOnRamp.TransactOpts)
}

// EmitCCIPSendRequested is a paid mutator transaction binding the contract method 0x51a2be12.
//
// Solidity: function emitCCIPSendRequested(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactor) EmitCCIPSendRequested(opts *bind.TransactOpts, sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.contract.Transact(opts, "emitCCIPSendRequested", sender, index)
}

// EmitCCIPSendRequested is a paid mutator transaction binding the contract method 0x51a2be12.
//
// Solidity: function emitCCIPSendRequested(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampSession) EmitCCIPSendRequested(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.EmitCCIPSendRequested(&_FiredrillOnRamp.TransactOpts, sender, index)
}

// EmitCCIPSendRequested is a paid mutator transaction binding the contract method 0x51a2be12.
//
// Solidity: function emitCCIPSendRequested(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactorSession) EmitCCIPSendRequested(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.EmitCCIPSendRequested(&_FiredrillOnRamp.TransactOpts, sender, index)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOnRamp.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.RenounceOwnership(&_FiredrillOnRamp.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.RenounceOwnership(&_FiredrillOnRamp.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOnRamp.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOnRamp *FiredrillOnRampSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.TransferOwnership(&_FiredrillOnRamp.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.TransferOwnership(&_FiredrillOnRamp.TransactOpts, newOwner)
}

// FiredrillOnRampCCIPSendRequestedIterator is returned from FilterCCIPSendRequested and is used to iterate over the raw logs and unpacked data for CCIPSendRequested events raised by the FiredrillOnRamp contract.
type FiredrillOnRampCCIPSendRequestedIterator struct {
	Event *FiredrillOnRampCCIPSendRequested // Event containing the contract specifics and raw log

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
func (it *FiredrillOnRampCCIPSendRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOnRampCCIPSendRequested)
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
		it.Event = new(FiredrillOnRampCCIPSendRequested)
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
func (it *FiredrillOnRampCCIPSendRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOnRampCCIPSendRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOnRampCCIPSendRequested represents a CCIPSendRequested event raised by the FiredrillOnRamp contract.
type FiredrillOnRampCCIPSendRequested struct {
	Message FiredrillOnRampEVM2EVMMessage
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCCIPSendRequested is a free log retrieval operation binding the contract event 0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd.
//
// Solidity: event CCIPSendRequested((uint64,address,address,uint64,uint256,bool,uint64,address,uint256,bytes,(address,uint256)[],bytes[],bytes32) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) FilterCCIPSendRequested(opts *bind.FilterOpts) (*FiredrillOnRampCCIPSendRequestedIterator, error) {

	logs, sub, err := _FiredrillOnRamp.contract.FilterLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampCCIPSendRequestedIterator{contract: _FiredrillOnRamp.contract, event: "CCIPSendRequested", logs: logs, sub: sub}, nil
}

// WatchCCIPSendRequested is a free log subscription operation binding the contract event 0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd.
//
// Solidity: event CCIPSendRequested((uint64,address,address,uint64,uint256,bool,uint64,address,uint256,bytes,(address,uint256)[],bytes[],bytes32) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) WatchCCIPSendRequested(opts *bind.WatchOpts, sink chan<- *FiredrillOnRampCCIPSendRequested) (event.Subscription, error) {

	logs, sub, err := _FiredrillOnRamp.contract.WatchLogs(opts, "CCIPSendRequested")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOnRampCCIPSendRequested)
				if err := _FiredrillOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
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

// ParseCCIPSendRequested is a log parse operation binding the contract event 0xd0c3c799bf9e2639de44391e7f524d229b2b55f5b1ea94b2bf7da42f7243dddd.
//
// Solidity: event CCIPSendRequested((uint64,address,address,uint64,uint256,bool,uint64,address,uint256,bytes,(address,uint256)[],bytes[],bytes32) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) ParseCCIPSendRequested(log types.Log) (*FiredrillOnRampCCIPSendRequested, error) {
	event := new(FiredrillOnRampCCIPSendRequested)
	if err := _FiredrillOnRamp.contract.UnpackLog(event, "CCIPSendRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOnRampOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FiredrillOnRamp contract.
type FiredrillOnRampOwnershipTransferStartedIterator struct {
	Event *FiredrillOnRampOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FiredrillOnRampOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOnRampOwnershipTransferStarted)
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
		it.Event = new(FiredrillOnRampOwnershipTransferStarted)
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
func (it *FiredrillOnRampOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOnRampOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOnRampOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FiredrillOnRamp contract.
type FiredrillOnRampOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillOnRampOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampOwnershipTransferStartedIterator{contract: _FiredrillOnRamp.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FiredrillOnRampOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOnRampOwnershipTransferStarted)
				if err := _FiredrillOnRamp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FiredrillOnRamp *FiredrillOnRampFilterer) ParseOwnershipTransferStarted(log types.Log) (*FiredrillOnRampOwnershipTransferStarted, error) {
	event := new(FiredrillOnRampOwnershipTransferStarted)
	if err := _FiredrillOnRamp.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillOnRampOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FiredrillOnRamp contract.
type FiredrillOnRampOwnershipTransferredIterator struct {
	Event *FiredrillOnRampOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FiredrillOnRampOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOnRampOwnershipTransferred)
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
		it.Event = new(FiredrillOnRampOwnershipTransferred)
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
func (it *FiredrillOnRampOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOnRampOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOnRampOwnershipTransferred represents a OwnershipTransferred event raised by the FiredrillOnRamp contract.
type FiredrillOnRampOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillOnRampOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampOwnershipTransferredIterator{contract: _FiredrillOnRamp.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiredrillOnRampOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOnRampOwnershipTransferred)
				if err := _FiredrillOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FiredrillOnRamp *FiredrillOnRampFilterer) ParseOwnershipTransferred(log types.Log) (*FiredrillOnRampOwnershipTransferred, error) {
	event := new(FiredrillOnRampOwnershipTransferred)
	if err := _FiredrillOnRamp.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
