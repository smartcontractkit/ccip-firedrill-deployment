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

// FiredrillOffRampDynamicConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampDynamicConfig struct {
	FeeQuoter                               common.Address
	PermissionLessExecutionThresholdSeconds uint32
	MessageInterceptor                      common.Address
}

// FiredrillOffRampMerkleRoot is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampMerkleRoot struct {
	SourceChainSelector uint64
	OnRampAddress       []byte
	MinSeqNr            uint64
	MaxSeqNr            uint64
	MerkleRoot          [32]byte
}

// FiredrillOffRampSourceChainConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampSourceChainConfig struct {
	Router                    common.Address
	IsEnabled                 bool
	MinSeqNr                  uint64
	IsRMNVerificationDisabled bool
	OnRamp                    []byte
}

// FiredrillOffRampStaticConfig is an auto generated low-level Go binding around an user-defined struct.
type FiredrillOffRampStaticConfig struct {
	ChainSelector        uint64
	GasForCallExactCheck uint16
	RmnRemote            common.Address
	TokenAdminRegistry   common.Address
	NonceManager         common.Address
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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"ctrl\",\"type\":\"address\",\"internalType\":\"contractFiredrillCompound\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitCommitReportAccepted\",\"inputs\":[{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"emitSourceChainConfigSet\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDynamicConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.DynamicConfig\",\"components\":[{\"name\":\"feeQuoter\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"permissionLessExecutionThresholdSeconds\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"messageInterceptor\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSourceChainConfig\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStaticConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structFiredrillOffRamp.StaticConfig\",\"components\":[{\"name\":\"chainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasForCallExactCheck\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"rmnRemote\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"tokenAdminRegistry\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nonceManager\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"CommitReportAccepted\",\"inputs\":[{\"name\":\"blessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"unblessedMerkleRoots\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.MerkleRoot[]\",\"components\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"onRampAddress\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"merkleRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"priceUpdates\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structInternal.PriceUpdates\",\"components\":[{\"name\":\"tokenPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.TokenPriceUpdate[]\",\"components\":[{\"name\":\"sourceToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"usdPerToken\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]},{\"name\":\"gasPriceUpdates\",\"type\":\"tuple[]\",\"internalType\":\"structInternal.GasPriceUpdate[]\",\"components\":[{\"name\":\"destChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"usdPerUnitGas\",\"type\":\"uint224\",\"internalType\":\"uint224\"}]}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SourceChainConfigSet\",\"inputs\":[{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sourceConfig\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structFiredrillOffRamp.SourceChainConfig\",\"components\":[{\"name\":\"router\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isEnabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"minSeqNr\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"isRMNVerificationDisabled\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"onRamp\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60a060405234801561000f575f80fd5b5060405161106238038061106283398101604081905261002e916100d9565b338061005357604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61005c8161006e565b506001600160a01b0316608052610106565b600180546001600160a01b03191690556100878161008a565b50565b5f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f602082840312156100e9575f80fd5b81516001600160a01b03811681146100ff575f80fd5b9392505050565b608051610ef761016b5f395f8181610177015281816102910152818161032a015281816103ca01528181610428015281816104dd0152818161052c0152818161066b015281816106f9015281816107c00152818161097601526109c50152610ef75ff3fe608060405234801561000f575f80fd5b50600436106100a6575f3560e01c806382c9978c1161006e57806382c9978c146101d75780638da5cb5b146101df5780638e73139f14610203578063e30c397814610216578063e9d68a8e14610227578063f2fde38b14610247575f80fd5b806306285c69146100aa578063181f5a7714610110578063715018a61461013f5780637437ff9f1461014957806379ba5097146101cf575b5f80fd5b6100b261025a565b6040805182516001600160401b0316815260208084015161ffff1690820152828201516001600160a01b03908116928201929092526060808401518316908201526080928301519091169181019190915260a0015b60405180910390f35b604080518082018252600d81526c04f666652616d7020312e362e3609c1b602082015290516101079190610bc5565b610147610363565b005b60408051606080820183525f8083526020808401829052928401819052835180830185526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116808352600a8387019081529287019384528651908152915163ffffffff169482019490945290519092169282019290925201610107565b610147610376565b6101476103bf565b5f546001600160a01b03165b6040516001600160a01b039091168152602001610107565b610147610211366004610bf2565b6105f9565b6001546001600160a01b03166101eb565b61023a610235366004610c29565b61093d565b6040516101079190610c44565b610147610255366004610cb7565b610a83565b6040805160a0810182525f808252602082018190529181018290526060810182905260808101919091526040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156102eb573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061030f9190610cd2565b6001600160401b031681525f60208201526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000166040820181905260608201819052608090910152919050565b61036b610af3565b6103745f610b1f565b565b60015433906001600160a01b031681146103b35760405163118cdaa760e01b81526001600160a01b03821660048201526024015b60405180910390fd5b6103bc81610b1f565b50565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146104265760405162461bcd60e51b815260206004820152600c60248201526b1bdb9b1e4818dbdb9d1c9bdb60a21b60448201526064016103aa565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa158015610482573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104a69190610cd2565b6001600160401b03167fbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c6040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016001151581526020015f6001600160401b031681526020015f151581526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610586573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105aa9190610ced565b6040516020016105d2919060609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f198184030181529181529152516105ef9190610c44565b60405180910390a2565b604080518082018252606080825260208201528151600180825281840190935290915f9190816020015b6040805160a0810182525f808252606060208084018290529383018290528201819052608082015282525f199092019101816106235790505090506040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316634e4bc8476040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106c5573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906106e99190610cd2565b6001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610753573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906107779190610ced565b604080516001600160a01b039092166020830152016040516020818303038152906040528152602001856001600160401b03168152602001846001600160401b031681526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561081a573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061083e9190610ced565b604080516001600160a01b0390921660208301526001600160401b03808916918301919091528616606082015260800160405160208183030381529060405280519060200120815250815f8151811061089957610899610d08565b60209081029190910101527fb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47815f60405190808252806020026020018201604052801561091e57816020015b6040805160a0810182525f808252606060208084018290529383018290528201819052608082015282525f199092019101816108e55790505b508460405161092f93929190610e0f565b60405180910390a150505050565b6040805160a0810182525f8082526020820181905291810182905260608082019290925260808101919091526040518060a001604052807f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031681526020016001151581526020015f6001600160401b031681526020015f151581526020017f00000000000000000000000000000000000000000000000000000000000000006001600160a01b031663c021e73c6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610a1f573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a439190610ced565b604051602001610a6b919060609190911b6bffffffffffffffffffffffff1916815260140190565b60408051601f19818403018152919052905292915050565b610a8b610af3565b600180546001600160a01b0383166001600160a01b03199091168117909155610abb5f546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b5f546001600160a01b031633146103745760405163118cdaa760e01b81523360048201526024016103aa565b600180546001600160a01b03191690556103bc815f80546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b5f81518084525f5b81811015610ba657602081850181015186830182015201610b8a565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610bd76020830184610b82565b9392505050565b6001600160401b03811681146103bc575f80fd5b5f8060408385031215610c03575f80fd5b8235610c0e81610bde565b91506020830135610c1e81610bde565b809150509250929050565b5f60208284031215610c39575f80fd5b8135610bd781610bde565b6020815260018060a01b0382511660208201526020820151151560408201526001600160401b0360408301511660608201526060820151151560808201525f608083015160a080840152610c9b60c0840182610b82565b949350505050565b6001600160a01b03811681146103bc575f80fd5b5f60208284031215610cc7575f80fd5b8135610bd781610ca3565b5f60208284031215610ce2575f80fd5b8151610bd781610bde565b5f60208284031215610cfd575f80fd5b8151610bd781610ca3565b634e487b7160e01b5f52603260045260245ffd5b5f82825180855260208086019550808260051b8401018186015f5b84811015610dab57601f19868403018952815160a06001600160401b03808351168652868301518288880152610d6f83880182610b82565b60408581015184169089015260608086015190931692880192909252506080928301519290950191909152509783019790830190600101610d37565b5090979650505050505050565b5f815180845260208085019450602084015f5b83811015610e0457815180516001600160401b031688528301516001600160e01b03168388015260409096019590820190600101610dcb565b509495945050505050565b606081525f610e216060830186610d1c565b602083820381850152610e348287610d1c565b9150604084830360408601526040830186516040855281815180845260608701915085830193505f92505b80831015610e9a57835180516001600160a01b031683528601516001600160e01b031686830152928501926001929092019190840190610e5f565b5084890151935085810385870152610eb28185610db8565b9b9a505050505050505050505056fea2646970667358221220bc5723f1040735c9f54e3f7f1cfd0a516dc65332668a7d0df4e0fac91943e32964736f6c63430008180033",
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

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,uint32,address))
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
// Solidity: function getDynamicConfig() view returns((address,uint32,address))
func (_FiredrillOffRamp *FiredrillOffRampSession) GetDynamicConfig() (FiredrillOffRampDynamicConfig, error) {
	return _FiredrillOffRamp.Contract.GetDynamicConfig(&_FiredrillOffRamp.CallOpts)
}

// GetDynamicConfig is a free data retrieval call binding the contract method 0x7437ff9f.
//
// Solidity: function getDynamicConfig() view returns((address,uint32,address))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetDynamicConfig() (FiredrillOffRampDynamicConfig, error) {
	return _FiredrillOffRamp.Contract.GetDynamicConfig(&_FiredrillOffRamp.CallOpts)
}

// GetSourceChainConfig is a free data retrieval call binding the contract method 0xe9d68a8e.
//
// Solidity: function getSourceChainConfig(uint64 ) view returns((address,bool,uint64,bool,bytes))
func (_FiredrillOffRamp *FiredrillOffRampCaller) GetSourceChainConfig(opts *bind.CallOpts, arg0 uint64) (FiredrillOffRampSourceChainConfig, error) {
	var out []interface{}
	err := _FiredrillOffRamp.contract.Call(opts, &out, "getSourceChainConfig", arg0)

	if err != nil {
		return *new(FiredrillOffRampSourceChainConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(FiredrillOffRampSourceChainConfig)).(*FiredrillOffRampSourceChainConfig)

	return out0, err

}

// GetSourceChainConfig is a free data retrieval call binding the contract method 0xe9d68a8e.
//
// Solidity: function getSourceChainConfig(uint64 ) view returns((address,bool,uint64,bool,bytes))
func (_FiredrillOffRamp *FiredrillOffRampSession) GetSourceChainConfig(arg0 uint64) (FiredrillOffRampSourceChainConfig, error) {
	return _FiredrillOffRamp.Contract.GetSourceChainConfig(&_FiredrillOffRamp.CallOpts, arg0)
}

// GetSourceChainConfig is a free data retrieval call binding the contract method 0xe9d68a8e.
//
// Solidity: function getSourceChainConfig(uint64 ) view returns((address,bool,uint64,bool,bytes))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetSourceChainConfig(arg0 uint64) (FiredrillOffRampSourceChainConfig, error) {
	return _FiredrillOffRamp.Contract.GetSourceChainConfig(&_FiredrillOffRamp.CallOpts, arg0)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint64,uint16,address,address,address))
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
// Solidity: function getStaticConfig() view returns((uint64,uint16,address,address,address))
func (_FiredrillOffRamp *FiredrillOffRampSession) GetStaticConfig() (FiredrillOffRampStaticConfig, error) {
	return _FiredrillOffRamp.Contract.GetStaticConfig(&_FiredrillOffRamp.CallOpts)
}

// GetStaticConfig is a free data retrieval call binding the contract method 0x06285c69.
//
// Solidity: function getStaticConfig() view returns((uint64,uint16,address,address,address))
func (_FiredrillOffRamp *FiredrillOffRampCallerSession) GetStaticConfig() (FiredrillOffRampStaticConfig, error) {
	return _FiredrillOffRamp.Contract.GetStaticConfig(&_FiredrillOffRamp.CallOpts)
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

// EmitCommitReportAccepted is a paid mutator transaction binding the contract method 0x8e73139f.
//
// Solidity: function emitCommitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitCommitReportAccepted(opts *bind.TransactOpts, minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitCommitReportAccepted", minSeqNr, maxSeqNr)
}

// EmitCommitReportAccepted is a paid mutator transaction binding the contract method 0x8e73139f.
//
// Solidity: function emitCommitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitCommitReportAccepted(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitCommitReportAccepted(&_FiredrillOffRamp.TransactOpts, minSeqNr, maxSeqNr)
}

// EmitCommitReportAccepted is a paid mutator transaction binding the contract method 0x8e73139f.
//
// Solidity: function emitCommitReportAccepted(uint64 minSeqNr, uint64 maxSeqNr) returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitCommitReportAccepted(minSeqNr uint64, maxSeqNr uint64) (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitCommitReportAccepted(&_FiredrillOffRamp.TransactOpts, minSeqNr, maxSeqNr)
}

// EmitSourceChainConfigSet is a paid mutator transaction binding the contract method 0x82c9978c.
//
// Solidity: function emitSourceChainConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactor) EmitSourceChainConfigSet(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillOffRamp.contract.Transact(opts, "emitSourceChainConfigSet")
}

// EmitSourceChainConfigSet is a paid mutator transaction binding the contract method 0x82c9978c.
//
// Solidity: function emitSourceChainConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampSession) EmitSourceChainConfigSet() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitSourceChainConfigSet(&_FiredrillOffRamp.TransactOpts)
}

// EmitSourceChainConfigSet is a paid mutator transaction binding the contract method 0x82c9978c.
//
// Solidity: function emitSourceChainConfigSet() returns()
func (_FiredrillOffRamp *FiredrillOffRampTransactorSession) EmitSourceChainConfigSet() (*types.Transaction, error) {
	return _FiredrillOffRamp.Contract.EmitSourceChainConfigSet(&_FiredrillOffRamp.TransactOpts)
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

// FiredrillOffRampCommitReportAcceptedIterator is returned from FilterCommitReportAccepted and is used to iterate over the raw logs and unpacked data for CommitReportAccepted events raised by the FiredrillOffRamp contract.
type FiredrillOffRampCommitReportAcceptedIterator struct {
	Event *FiredrillOffRampCommitReportAccepted // Event containing the contract specifics and raw log

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
func (it *FiredrillOffRampCommitReportAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampCommitReportAccepted)
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
		it.Event = new(FiredrillOffRampCommitReportAccepted)
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
func (it *FiredrillOffRampCommitReportAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampCommitReportAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampCommitReportAccepted represents a CommitReportAccepted event raised by the FiredrillOffRamp contract.
type FiredrillOffRampCommitReportAccepted struct {
	BlessedMerkleRoots   []FiredrillOffRampMerkleRoot
	UnblessedMerkleRoots []FiredrillOffRampMerkleRoot
	PriceUpdates         InternalPriceUpdates
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterCommitReportAccepted is a free log retrieval operation binding the contract event 0xb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47.
//
// Solidity: event CommitReportAccepted((uint64,bytes,uint64,uint64,bytes32)[] blessedMerkleRoots, (uint64,bytes,uint64,uint64,bytes32)[] unblessedMerkleRoots, ((address,uint224)[],(uint64,uint224)[]) priceUpdates)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterCommitReportAccepted(opts *bind.FilterOpts) (*FiredrillOffRampCommitReportAcceptedIterator, error) {

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampCommitReportAcceptedIterator{contract: _FiredrillOffRamp.contract, event: "CommitReportAccepted", logs: logs, sub: sub}, nil
}

// WatchCommitReportAccepted is a free log subscription operation binding the contract event 0xb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47.
//
// Solidity: event CommitReportAccepted((uint64,bytes,uint64,uint64,bytes32)[] blessedMerkleRoots, (uint64,bytes,uint64,uint64,bytes32)[] unblessedMerkleRoots, ((address,uint224)[],(uint64,uint224)[]) priceUpdates)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchCommitReportAccepted(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampCommitReportAccepted) (event.Subscription, error) {

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "CommitReportAccepted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampCommitReportAccepted)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

// ParseCommitReportAccepted is a log parse operation binding the contract event 0xb967c9b9e1b7af9a61ca71ff00e9f5b89ec6f2e268de8dacf12f0de8e51f3e47.
//
// Solidity: event CommitReportAccepted((uint64,bytes,uint64,uint64,bytes32)[] blessedMerkleRoots, (uint64,bytes,uint64,uint64,bytes32)[] unblessedMerkleRoots, ((address,uint224)[],(uint64,uint224)[]) priceUpdates)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseCommitReportAccepted(log types.Log) (*FiredrillOffRampCommitReportAccepted, error) {
	event := new(FiredrillOffRampCommitReportAccepted)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "CommitReportAccepted", log); err != nil {
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

// FiredrillOffRampSourceChainConfigSetIterator is returned from FilterSourceChainConfigSet and is used to iterate over the raw logs and unpacked data for SourceChainConfigSet events raised by the FiredrillOffRamp contract.
type FiredrillOffRampSourceChainConfigSetIterator struct {
	Event *FiredrillOffRampSourceChainConfigSet // Event containing the contract specifics and raw log

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
func (it *FiredrillOffRampSourceChainConfigSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillOffRampSourceChainConfigSet)
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
		it.Event = new(FiredrillOffRampSourceChainConfigSet)
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
func (it *FiredrillOffRampSourceChainConfigSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillOffRampSourceChainConfigSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillOffRampSourceChainConfigSet represents a SourceChainConfigSet event raised by the FiredrillOffRamp contract.
type FiredrillOffRampSourceChainConfigSet struct {
	SourceChainSelector uint64
	SourceConfig        FiredrillOffRampSourceChainConfig
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterSourceChainConfigSet is a free log retrieval operation binding the contract event 0xbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c.
//
// Solidity: event SourceChainConfigSet(uint64 indexed sourceChainSelector, (address,bool,uint64,bool,bytes) sourceConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) FilterSourceChainConfigSet(opts *bind.FilterOpts, sourceChainSelector []uint64) (*FiredrillOffRampSourceChainConfigSetIterator, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.FilterLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillOffRampSourceChainConfigSetIterator{contract: _FiredrillOffRamp.contract, event: "SourceChainConfigSet", logs: logs, sub: sub}, nil
}

// WatchSourceChainConfigSet is a free log subscription operation binding the contract event 0xbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c.
//
// Solidity: event SourceChainConfigSet(uint64 indexed sourceChainSelector, (address,bool,uint64,bool,bytes) sourceConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) WatchSourceChainConfigSet(opts *bind.WatchOpts, sink chan<- *FiredrillOffRampSourceChainConfigSet, sourceChainSelector []uint64) (event.Subscription, error) {

	var sourceChainSelectorRule []interface{}
	for _, sourceChainSelectorItem := range sourceChainSelector {
		sourceChainSelectorRule = append(sourceChainSelectorRule, sourceChainSelectorItem)
	}

	logs, sub, err := _FiredrillOffRamp.contract.WatchLogs(opts, "SourceChainConfigSet", sourceChainSelectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillOffRampSourceChainConfigSet)
				if err := _FiredrillOffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
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

// ParseSourceChainConfigSet is a log parse operation binding the contract event 0xbd1ab25a0ff0a36a588597ba1af11e30f3f210de8b9e818cc9bbc457c94c8d8c.
//
// Solidity: event SourceChainConfigSet(uint64 indexed sourceChainSelector, (address,bool,uint64,bool,bytes) sourceConfig)
func (_FiredrillOffRamp *FiredrillOffRampFilterer) ParseSourceChainConfigSet(log types.Log) (*FiredrillOffRampSourceChainConfigSet, error) {
	event := new(FiredrillOffRampSourceChainConfigSet)
	if err := _FiredrillOffRamp.contract.UnpackLog(event, "SourceChainConfigSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
