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
	FeeQuoter              common.Address
	ReentrancyGuardEntered bool
	MessageInterceptor     common.Address
	FeeAggregator          common.Address
	AllowlistAdmin         common.Address
}

// FiredrillOnRampEVM2AnyRampMessage is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampEVM2AnyRampMessage struct {
	Header         FiredrillOnRampRampMessageHeader
	Sender         common.Address
	Data           []byte
	Receiver       []byte
	ExtraArgs      []byte
	FeeToken       common.Address
	FeeTokenAmount *big.Int
	FeeValueJuels  *big.Int
	TokenAmounts   []FiredrillOnRampEVM2AnyTokenTransfer
}

// FiredrillOnRampEVM2AnyTokenTransfer is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampEVM2AnyTokenTransfer struct {
	SourcePoolAddress common.Address
	DestTokenAddress  []byte
	ExtraData         []byte
	Amount            *big.Int
	DestExecData      []byte
}

// FiredrillOnRampRampMessageHeader is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampRampMessageHeader struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	DestChainSelector   uint64
	SequenceNumber      uint64
	Nonce               uint64
}

// FiredrillOnRampStaticConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOnRampStaticConfig struct {
	ChainSelector      uint64
	RmnRemote          common.Address
	NonceManager       common.Address
	TokenAdminRegistry common.Address
}

// FiredrillOnRampMetaData contains all meta data concerning the FiredrillOnRamp contract.
var FiredrillOnRampMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ctrl\",\"type\":\"address\",\"internalType\":\"contractFiredrillCompound\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitCCIPMessageSent\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"index\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDestChainConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"allowlistEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOnRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"reentrancyGuardEntered\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeAggregator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowlistAdmin\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOnRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CCIPMessageSent\",\"inputs\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"message\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOnRamp.EVM2AnyRampMessage\",\"components\":[{\"name\":\"header\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOnRamp.RampMessageHeader\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sequenceNumber\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nonce\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"receiver\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraArgs\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"feeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeTokenAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"feeValueJuels\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"tokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structFiredrillOnRamp.EVM2AnyTokenTransfer[]\",\"components\":[{\"name\":\"sourcePoolAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"destTokenAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"extraData\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"destExecData\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b50604051610e26380380610e2683398101604081905261002e916100d9565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161006e565b506001600160a01b0316608052610106565b600180546001600160a01b03191690556100878161008a565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e9575f80fd5b81516001600160a01b03811681146100ff575f80fd5b9392505050565b608051610cd06101565f395f8181610141015281816101de0152818161030101528181610395015281816104a501528181610567015281816105f60152818161070601526108340152610cd05ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c806379ba50971161006357806379ba50971461026e5780638da5cb5b14610276578063e30c39781461029a578063f2fde38b146102ab578063f47d2ea7146102be575f80fd5b806306285c691461009f578063181f5a77146100fe5780636def4ce71461012c578063715018a6146101925780637437ff9f1461019c575b5f80fd5b6100a76102d1565b6040516100f59190815167ffffffffffffffff1681526020808301516001600160a01b0390811691830191909152604080840151821690830152606092830151169181019190915260800190565b60405180910390f35b604080518082018252600c81526b04f6e52616d7020312e362e360a41b602082015290516100f591906109ca565b61016361013a3660046109f8565b505f9081907f000000000000000000000000000000000000000000000000000000000000000090565b6040805167ffffffffffffffff909416845291151560208401526001600160a01b0316908201526060016100f5565b61019a6103ce565b005b61021f6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506040805160a0810182526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000168082525f602083018190529282018390526060820152608081019190915290565b6040516100f5919081516001600160a01b039081168252602080840151151590830152604080840151821690830152606080840151821690830152608092830151169181019190915260a00190565b61019a6103e1565b5f546001600160a01b03165b6040516001600160a01b0390911681526020016100f5565b6001546001600160a01b0316610282565b61019a6102b9366004610a27565b61042a565b61019a6102cc366004610a42565b61049a565b604080516080810182525f80825260208201819052918101829052606081019190915260405180608001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa15801561035b573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061037f9190610a79565b67ffffffffffffffff1681526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166020820181905260408201819052606090910152919050565b6103d66108f8565b6103df5f610924565b565b60015433906001600160a01b0316811461041e5760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b61042781610924565b50565b6104326108f8565b600180546001600160a01b0383166001600160a01b031990911681179091556104625f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105015760405162461bcd60e51b815260206004820152600c60248201526b1bdb9b1e4818dbdb9d1c9bdb60a21b6044820152606401610415565b604080516101c081019091526bffffffffffffffffffffffff19606084901b166101e08201526001600160c01b031960c083901b166101f48201525f90806101208101806101fc83016040516020818303038152906040528051906020012081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156105c1573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105e59190610a79565b67ffffffffffffffff1681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa158015610650573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106749190610a79565b67ffffffffffffffff9081168252861660208083019190915260016040928301529183526001600160a01b03871683830181905281518083018352600381526231323360e81b81850152848301528151928301526060909201910160405160208183030381529060405281526020016040518060400160405280600381526020016231323360e81b81525081526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610760573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107849190610a94565b6001600160a01b031681526020015f81526020015f81526020015f67ffffffffffffffff8111156107b7576107b7610aaf565b60405190808252806020026020018201604052801561082157816020015b61080e6040518060a001604052805f6001600160a01b0316815260200160608152602001606081526020015f8152602001606081525090565b8152602001906001900390816107d55790505b5081525090508167ffffffffffffffff167f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa15801561088e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108b29190610a79565b67ffffffffffffffff167f192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32836040516108eb9190610b7b565b60405180910390a3505050565b5f546001600160a01b031633146103df5760405163118cdaa760e01b8152336004820152602401610415565b600180546001600160a01b0319169055610427815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f81518084525f5b818110156109ab5760208185018101518683018201520161098f565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f6109dc6020830184610987565b9392505050565b67ffffffffffffffff81168114610427575f80fd5b5f60208284031215610a08575f80fd5b81356109dc816109e3565b6001600160a01b0381168114610427575f80fd5b5f60208284031215610a37575f80fd5b81356109dc81610a13565b5f8060408385031215610a53575f80fd5b8235610a5e81610a13565b91506020830135610a6e816109e3565b809150509250929050565b5f60208284031215610a89575f80fd5b81516109dc816109e3565b5f60208284031215610aa4575f80fd5b81516109dc81610a13565b634e487b7160e01b5f52604160045260245ffd5b5f82825180855260208086019550808260051b8401018186015f5b84811015610b6e57858303601f19018952815180516001600160a01b031684528481015160a086860181905290610b1782870182610987565b91505060408083015186830382880152610b318382610987565b92505050606080830151818701525060808083015192508582038187015250610b5a8183610987565b9a86019a9450505090830190600101610ade565b5090979650505050505050565b60208152610bcc60208201835180518252602081015167ffffffffffffffff808216602085015280604084015116604085015280606084015116606085015280608084015116608085015250505050565b5f6020830151610be760c08401826001600160a01b03169052565b5060408301516101a08060e0850152610c046101c0850183610987565b91506060850151601f19610100818786030181880152610c248584610987565b9450608088015192508187860301610120880152610c428584610987565b945060a08801519250610c616101408801846001600160a01b03169052565b60c088015161016088015260e0880151610180880152870151868503909101838701529050610c908382610ac3565b969550505050505056fea26469706673582212206dd0f0689d34d17a2399d4da6e5dfb6f268eb2a51892040a32f942ff08a8714d64736f6c63430008180033",
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

// GetDestChainConfig is a free data retrieval call binding the contract method 0x6def4ce7.
//
// Solidity: function getDestChainConfig(uint64 ) view returns(uint64 sequenceNumber, bool allowlistEnabled, address router)
func (_FiredrillOnRamp *FiredrillOnRampCaller) GetDestChainConfig(opts *bind.CallOpts, arg0 uint64) (struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}, error) {
	var out []interface{}
	err := _FiredrillOnRamp.contract.Call(opts, &out, "getDestChainConfig", arg0)

	outstruct := new(struct {
		SequenceNumber   uint64
		AllowlistEnabled bool
		Router           common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SequenceNumber = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.AllowlistEnabled = *abi.ConvertType(out[1], new(bool)).(*bool)
	outstruct.Router = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetDestChainConfig is a free data retrieval call binding the contract method 0x6def4ce7.
//
// Solidity: function getDestChainConfig(uint64 ) view returns(uint64 sequenceNumber, bool allowlistEnabled, address router)
func (_FiredrillOnRamp *FiredrillOnRampSession) GetDestChainConfig(arg0 uint64) (struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}, error) {
	return _FiredrillOnRamp.Contract.GetDestChainConfig(&_FiredrillOnRamp.CallOpts, arg0)
}

// GetDestChainConfig is a free data retrieval call binding the contract method 0x6def4ce7.
//
// Solidity: function getDestChainConfig(uint64 ) view returns(uint64 sequenceNumber, bool allowlistEnabled, address router)
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) GetDestChainConfig(arg0 uint64) (struct {
	SequenceNumber   uint64
	AllowlistEnabled bool
	Router           common.Address
}, error) {
	return _FiredrillOnRamp.Contract.GetDestChainConfig(&_FiredrillOnRamp.CallOpts, arg0)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,bool,address,address,address))
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
// Solidity: function getDynamicConfig() view returns((address,bool,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampSession) GetDynamicConfig() (FiredrillOnRampDynamicConfig, error) {
	return _FiredrillOnRamp.Contract.GetDynamicConfig(&_FiredrillOnRamp.CallOpts)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,bool,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampCallerSession) GetDynamicConfig() (FiredrillOnRampDynamicConfig, error) {
	return _FiredrillOnRamp.Contract.GetDynamicConfig(&_FiredrillOnRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint64,address,address,address))
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
// Solidity: function getStaticConfig() view returns((uint64,address,address,address))
func (_FiredrillOnRamp *FiredrillOnRampSession) GetStaticConfig() (FiredrillOnRampStaticConfig, error) {
	return _FiredrillOnRamp.Contract.GetStaticConfig(&_FiredrillOnRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint64,address,address,address))
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

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xf47d2ea7.
//
// Solidity: function emitCCIPMessageSent(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactor) EmitCCIPMessageSent(opts *bind.TransactOpts, sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.contract.Transact(opts, "emitCCIPMessageSent", sender, index)
}

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xf47d2ea7.
//
// Solidity: function emitCCIPMessageSent(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampSession) EmitCCIPMessageSent(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.EmitCCIPMessageSent(&_FiredrillOnRamp.TransactOpts, sender, index)
}

// EmitCCIPMessageSent is a paid mutator transaction binding the contract method 0xf47d2ea7.
//
// Solidity: function emitCCIPMessageSent(address sender, uint64 index) returns()
func (_FiredrillOnRamp *FiredrillOnRampTransactorSession) EmitCCIPMessageSent(sender common.Address, index uint64) (*types.Transaction, error) {
	return _FiredrillOnRamp.Contract.EmitCCIPMessageSent(&_FiredrillOnRamp.TransactOpts, sender, index)
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

// FiredrillOnRampCCIPMessageSentIterator is returned from FilterCCIPMessageSent and is used to iterate over the raw logs and unpacked data for CCIPMessageSent events raised by the FiredrillOnRamp contract.
type FiredrillOnRampCCIPMessageSentIterator struct {
	Event *FiredrillOnRampCCIPMessageSent // Event containing the contract specifics and raw log

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
func (it *FiredrillOnRampCCIPMessageSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOnRampCCIPMessageSent)
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
		it.Event = new(FiredrillOnRampCCIPMessageSent)
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
func (it *FiredrillOnRampCCIPMessageSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOnRampCCIPMessageSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOnRampCCIPMessageSent represents a CCIPMessageSent event raised by the FiredrillOnRamp contract.
type FiredrillOnRampCCIPMessageSent struct {
	DestChainSelector uint64
	SequenceNumber    uint64
	Message           FiredrillOnRampEVM2AnyRampMessage
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterCCIPMessageSent is a free log retrieval operation binding the contract event 0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64,uint64),address,bytes,bytes,bytes,address,uint256,uint256,(address,bytes,bytes,uint256,bytes)[]) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) FilterCCIPMessageSent(opts *bind.FilterOpts, destChainSelector []uint64, sequenceNumber []uint64) (*FiredrillOnRampCCIPMessageSentIterator, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.FilterLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOnRampCCIPMessageSentIterator{contract: _FiredrillOnRamp.contract, event: "CCIPMessageSent", logs: logs, sub: sub}, nil
}

// WatchCCIPMessageSent is a free log subscription operation binding the contract event 0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64,uint64),address,bytes,bytes,bytes,address,uint256,uint256,(address,bytes,bytes,uint256,bytes)[]) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) WatchCCIPMessageSent(opts *bind.WatchOpts, sink chan<- *FiredrillOnRampCCIPMessageSent, destChainSelector []uint64, sequenceNumber []uint64) (event.Subscription, error) {

	var destChainSelectorRule []interface{}
	for _, destChainSelectorItem := range destChainSelector {
		destChainSelectorRule = append(destChainSelectorRule, destChainSelectorItem)
	}
	var sequenceNumberRule []interface{}
	for _, sequenceNumberItem := range sequenceNumber {
		sequenceNumberRule = append(sequenceNumberRule, sequenceNumberItem)
	}

	logs, sub, err := _FiredrillOnRamp.contract.WatchLogs(opts, "CCIPMessageSent", destChainSelectorRule, sequenceNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOnRampCCIPMessageSent)
				if err := _FiredrillOnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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

// ParseCCIPMessageSent is a log parse operation binding the contract event 0x192442a2b2adb6a7948f097023cb6b57d29d3a7a5dd33e6666d33c39cc456f32.
//
// Solidity: event CCIPMessageSent(uint64 indexed destChainSelector, uint64 indexed sequenceNumber, ((bytes32,uint64,uint64,uint64,uint64),address,bytes,bytes,bytes,address,uint256,uint256,(address,bytes,bytes,uint256,bytes)[]) message)
func (_FiredrillOnRamp *FiredrillOnRampFilterer) ParseCCIPMessageSent(log types.Log) (*FiredrillOnRampCCIPMessageSent, error) {
	event := new(FiredrillOnRampCCIPMessageSent)
	if err := _FiredrillOnRamp.contract.UnpackLog(event, "CCIPMessageSent", log); err != nil {
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
