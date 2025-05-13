// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package firedrill_revert_message_receiver

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

// ClientAny2EVMMessage is an auto generated low-level Go binding around an user-defined struct.
type ClientAny2EVMMessage struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
}

// ClientEVMTokenAmount is an auto generated low-level Go binding around an user-defined struct.
type ClientEVMTokenAmount struct {
	Token  common.Address
	Amount *big.Int
}

// FiredrillRevertMessageReceiverMetaData contains all meta data concerning the FiredrillRevertMessageReceiver contract.
var FiredrillRevertMessageReceiverMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"acceptOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOfToken\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ccipReceive\",\"inputs\":[{\"name\":\"message\",\"type\":\"tuple\",\"internalType\":\"structClient.Any2EVMMessage\",\"components\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pendingOwner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"s_toRevert\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setErr\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setRevert\",\"inputs\":[{\"name\":\"toRevert\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"typeAndVersion\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"withdrawFunds\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdrawTokens\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"MessageReceived\",\"inputs\":[{\"name\":\"messageId\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"sourceChainSelector\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"destTokenAmounts\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structClient.EVMTokenAmount[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"NativeFundsWithdrawn\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferStarted\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokensWithdrawn\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValueReceived\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"CustomError\",\"inputs\":[{\"name\":\"err\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"type\":\"error\",\"name\":\"InsufficientBalance\",\"inputs\":[{\"name\":\"available\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"required\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReceiveRevert\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x60a060405234801561000f575f80fd5b503360808190525f805460ff191660011790558061004657604051631e4fbdf760e01b81525f600482015260240160405180910390fd5b61004f81610055565b506100c2565b600380546001600160a01b031916905561006e81610071565b50565b600280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b608051610ee56100fd5f395f818161034c0152818161043c015281816104da01528181610577015281816105c501526106470152610ee55ff3fe6080604052600436106100dc575f3560e01c806379ba50971161007c5780638fb5f171116100575780638fb5f17114610276578063b99152d0146102a2578063e30c3978146102cf578063f2fde38b146102ec575f80fd5b806379ba50971461021257806385572ffb146102265780638da5cb5b14610245575f80fd5b806324600fc3116100b757806324600fc3146101b35780635100fc21146101c7578063715018a6146101df57806377f5b0e6146101f3575f80fd5b806301ffc9a71461013d57806306b091f914610171578063181f5a7714610192575f80fd5b36610139575f5460ff161561010457604051633085b8db60e01b815260040160405180910390fd5b6040513481527fe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef9060200160405180910390a1005b5f80fd5b348015610148575f80fd5b5061015c610157366004610926565b61030b565b60405190151581526020015b60405180910390f35b34801561017c575f80fd5b5061019061018b366004610968565b610341565b005b34801561019d575f80fd5b506101a661054c565b6040516101689190610990565b3480156101be575f80fd5b5061019061056c565b3480156101d2575f80fd5b505f5461015c9060ff1681565b3480156101ea575f80fd5b506101906106ac565b3480156101fe575f80fd5b5061019061020d3660046109f0565b6106bf565b34801561021d575f80fd5b506101906106cf565b348015610231575f80fd5b50610190610240366004610a9b565b610713565b348015610250575f80fd5b506002546001600160a01b03165b6040516001600160a01b039091168152602001610168565b348015610281575f80fd5b50610190610290366004610adf565b5f805460ff1916911515919091179055565b3480156102ad575f80fd5b506102c16102bc366004610afa565b6107b2565b604051908152602001610168565b3480156102da575f80fd5b506003546001600160a01b031661025e565b3480156102f7575f80fd5b50610190610306366004610afa565b610823565b5f6001600160e01b031982166385572ffb60e01b148061033b57506001600160e01b031982166301ffc9a760e01b145b92915050565b336001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610389576040516282b42960e81b815260040160405180910390fd5b6040516370a0823160e01b815230600482015282905f906001600160a01b038316906370a0823190602401602060405180830381865afa1580156103cf573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906103f39190610b13565b9050828110156104255760405163cf47918160e01b815260048101829052602481018490526044015b60405180910390fd5b60405163a9059cbb60e01b81526001600160a01b037f000000000000000000000000000000000000000000000000000000000000000081166004830152602482018590525f919084169063a9059cbb906044016020604051808303815f875af1158015610494573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104b89190610b2a565b9050806104d8576040516312171d8360e31b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316856001600160a01b03167f6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a48660405161053d91815260200190565b60405180910390a35050505050565b6060604051806060016040528060248152602001610e8c60249139905090565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146105b4576040516282b42960e81b815260040160405180910390fd5b60405147905f906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169083908381818185875af1925050503d805f811461061e576040519150601f19603f3d011682016040523d82523d5f602084013e610623565b606091505b5050905080610645576040516312171d8360e31b815260040160405180910390fd5b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03167fd50b71a2790ecccf5881141fe9ae079e17c66aace5d50ba383d443ecd398ffa5836040516106a091815260200190565b60405180910390a25050565b6106b4610894565b6106bd5f6108c1565b565b60016106cb8282610bc9565b5050565b60035433906001600160a01b031681146107075760405163118cdaa760e01b81526001600160a01b038216600482015260240161041c565b610710816108c1565b50565b5f5460ff1615610739576001604051635a4ff67160e01b815260040161041c9190610c89565b7f707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc813561076c6040840160208501610d13565b6107796040850185610d3a565b6107866060870187610d3a565b6107936080890189610d84565b6040516107a7989796959493929190610df2565b60405180910390a150565b6040516370a0823160e01b81523060048201525f9082906001600160a01b038216906370a0823190602401602060405180830381865afa1580156107f8573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061081c9190610b13565b9392505050565b61082b610894565b600380546001600160a01b0383166001600160a01b0319909116811790915561085c6002546001600160a01b031690565b6001600160a01b03167f38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e2270060405160405180910390a350565b6002546001600160a01b031633146106bd5760405163118cdaa760e01b815233600482015260240161041c565b600380546001600160a01b031916905561071081600280546001600160a01b038381166001600160a01b0319831681179093556040519116919082907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a35050565b5f60208284031215610936575f80fd5b81356001600160e01b03198116811461081c575f80fd5b80356001600160a01b0381168114610963575f80fd5b919050565b5f8060408385031215610979575f80fd5b6109828361094d565b946020939093013593505050565b5f602080835283518060208501525f5b818110156109bc578581018301518582016040015282016109a0565b505f604082860101526040601f19601f8301168501019250505092915050565b634e487b7160e01b5f52604160045260245ffd5b5f60208284031215610a00575f80fd5b813567ffffffffffffffff80821115610a17575f80fd5b818401915084601f830112610a2a575f80fd5b813581811115610a3c57610a3c6109dc565b604051601f8201601f19908116603f01168101908382118183101715610a6457610a646109dc565b81604052828152876020848701011115610a7c575f80fd5b826020860160208301375f928101602001929092525095945050505050565b5f60208284031215610aab575f80fd5b813567ffffffffffffffff811115610ac1575f80fd5b820160a0818503121561081c575f80fd5b8015158114610710575f80fd5b5f60208284031215610aef575f80fd5b813561081c81610ad2565b5f60208284031215610b0a575f80fd5b61081c8261094d565b5f60208284031215610b23575f80fd5b5051919050565b5f60208284031215610b3a575f80fd5b815161081c81610ad2565b600181811c90821680610b5957607f821691505b602082108103610b7757634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115610bc457805f5260205f20601f840160051c81016020851015610ba25750805b601f840160051c820191505b81811015610bc1575f8155600101610bae565b50505b505050565b815167ffffffffffffffff811115610be357610be36109dc565b610bf781610bf18454610b45565b84610b7d565b602080601f831160018114610c2a575f8415610c135750858301515b5f19600386901b1c1916600185901b178555610c81565b5f85815260208120601f198616915b82811015610c5857888601518255948401946001909101908401610c39565b5085821015610c7557878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b5f60208083525f8454610c9b81610b45565b806020870152604060018084165f8114610cbc5760018114610cd857610d05565b60ff19851660408a0152604084151560051b8a01019550610d05565b895f5260205f205f5b85811015610cfc5781548b8201860152908301908801610ce1565b8a016040019650505b509398975050505050505050565b5f60208284031215610d23575f80fd5b813567ffffffffffffffff8116811461081c575f80fd5b5f808335601e19843603018112610d4f575f80fd5b83018035915067ffffffffffffffff821115610d69575f80fd5b602001915036819003821315610d7d575f80fd5b9250929050565b5f808335601e19843603018112610d99575f80fd5b83018035915067ffffffffffffffff821115610db3575f80fd5b6020019150600681901b3603821315610d7d575f80fd5b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b8881525f602067ffffffffffffffff8a166020840152604060a06040850152610e1f60a085018a8c610dca565b8481036060860152610e3281898b610dca565b85810360808701528681528791506020015f5b87811015610e79576001600160a01b03610e5e8461094d565b16825282850135858301529183019190830190600101610e45565b509d9c5050505050505050505050505056fe466972656472696c6c5265766572744d657373616765526563656976657220312e302e30a2646970667358221220d345a01f681061433196d8069e7b6abee675d8f4be98ee7f3561c91486be856164736f6c63430008180033",
}

// FiredrillRevertMessageReceiverABI is the input ABI used to generate the binding from.
// Deprecated: Use FiredrillRevertMessageReceiverMetaData.ABI instead.
var FiredrillRevertMessageReceiverABI = FiredrillRevertMessageReceiverMetaData.ABI

// FiredrillRevertMessageReceiverBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FiredrillRevertMessageReceiverMetaData.Bin instead.
var FiredrillRevertMessageReceiverBin = FiredrillRevertMessageReceiverMetaData.Bin

// DeployFiredrillRevertMessageReceiver deploys a new Ethereum contract, binding an instance of FiredrillRevertMessageReceiver to it.
func DeployFiredrillRevertMessageReceiver(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FiredrillRevertMessageReceiver, error) {
	parsed, err := FiredrillRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FiredrillRevertMessageReceiverBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FiredrillRevertMessageReceiver{FiredrillRevertMessageReceiverCaller: FiredrillRevertMessageReceiverCaller{contract: contract}, FiredrillRevertMessageReceiverTransactor: FiredrillRevertMessageReceiverTransactor{contract: contract}, FiredrillRevertMessageReceiverFilterer: FiredrillRevertMessageReceiverFilterer{contract: contract}}, nil
}

// FiredrillRevertMessageReceiver is an auto generated Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiver struct {
	FiredrillRevertMessageReceiverCaller     // Read-only binding to the contract
	FiredrillRevertMessageReceiverTransactor // Write-only binding to the contract
	FiredrillRevertMessageReceiverFilterer   // Log filterer for contract events
}

// FiredrillRevertMessageReceiverCaller is an auto generated read-only Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiverCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillRevertMessageReceiverTransactor is an auto generated write-only Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiverTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillRevertMessageReceiverFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FiredrillRevertMessageReceiverFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FiredrillRevertMessageReceiverSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FiredrillRevertMessageReceiverSession struct {
	Contract     *FiredrillRevertMessageReceiver // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                   // Call options to use throughout this session
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// FiredrillRevertMessageReceiverCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FiredrillRevertMessageReceiverCallerSession struct {
	Contract *FiredrillRevertMessageReceiverCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                         // Call options to use throughout this session
}

// FiredrillRevertMessageReceiverTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FiredrillRevertMessageReceiverTransactorSession struct {
	Contract     *FiredrillRevertMessageReceiverTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                         // Transaction auth options to use throughout this session
}

// FiredrillRevertMessageReceiverRaw is an auto generated low-level Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiverRaw struct {
	Contract *FiredrillRevertMessageReceiver // Generic contract binding to access the raw methods on
}

// FiredrillRevertMessageReceiverCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiverCallerRaw struct {
	Contract *FiredrillRevertMessageReceiverCaller // Generic read-only contract binding to access the raw methods on
}

// FiredrillRevertMessageReceiverTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FiredrillRevertMessageReceiverTransactorRaw struct {
	Contract *FiredrillRevertMessageReceiverTransactor // Generic write-only contract binding to access the raw methods on
}

// NewFiredrillRevertMessageReceiver creates a new instance of FiredrillRevertMessageReceiver, bound to a specific deployed contract.
func NewFiredrillRevertMessageReceiver(address common.Address, backend bind.ContractBackend) (*FiredrillRevertMessageReceiver, error) {
	contract, err := bindFiredrillRevertMessageReceiver(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiver{FiredrillRevertMessageReceiverCaller: FiredrillRevertMessageReceiverCaller{contract: contract}, FiredrillRevertMessageReceiverTransactor: FiredrillRevertMessageReceiverTransactor{contract: contract}, FiredrillRevertMessageReceiverFilterer: FiredrillRevertMessageReceiverFilterer{contract: contract}}, nil
}

// NewFiredrillRevertMessageReceiverCaller creates a new read-only instance of FiredrillRevertMessageReceiver, bound to a specific deployed contract.
func NewFiredrillRevertMessageReceiverCaller(address common.Address, caller bind.ContractCaller) (*FiredrillRevertMessageReceiverCaller, error) {
	contract, err := bindFiredrillRevertMessageReceiver(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverCaller{contract: contract}, nil
}

// NewFiredrillRevertMessageReceiverTransactor creates a new write-only instance of FiredrillRevertMessageReceiver, bound to a specific deployed contract.
func NewFiredrillRevertMessageReceiverTransactor(address common.Address, transactor bind.ContractTransactor) (*FiredrillRevertMessageReceiverTransactor, error) {
	contract, err := bindFiredrillRevertMessageReceiver(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverTransactor{contract: contract}, nil
}

// NewFiredrillRevertMessageReceiverFilterer creates a new log filterer instance of FiredrillRevertMessageReceiver, bound to a specific deployed contract.
func NewFiredrillRevertMessageReceiverFilterer(address common.Address, filterer bind.ContractFilterer) (*FiredrillRevertMessageReceiverFilterer, error) {
	contract, err := bindFiredrillRevertMessageReceiver(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverFilterer{contract: contract}, nil
}

// bindFiredrillRevertMessageReceiver binds a generic wrapper to an already deployed contract.
func bindFiredrillRevertMessageReceiver(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FiredrillRevertMessageReceiverMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillRevertMessageReceiver.Contract.FiredrillRevertMessageReceiverCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.FiredrillRevertMessageReceiverTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.FiredrillRevertMessageReceiverTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FiredrillRevertMessageReceiver.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.contract.Transact(opts, method, params...)
}

// BalanceOfToken is a free data retrieval call binding the contract method 0xb99152d0.
//
// Solidity: function balanceOfToken(address token) view returns(uint256)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) BalanceOfToken(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "balanceOfToken", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOfToken is a free data retrieval call binding the contract method 0xb99152d0.
//
// Solidity: function balanceOfToken(address token) view returns(uint256)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) BalanceOfToken(token common.Address) (*big.Int, error) {
	return _FiredrillRevertMessageReceiver.Contract.BalanceOfToken(&_FiredrillRevertMessageReceiver.CallOpts, token)
}

// BalanceOfToken is a free data retrieval call binding the contract method 0xb99152d0.
//
// Solidity: function balanceOfToken(address token) view returns(uint256)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) BalanceOfToken(token common.Address) (*big.Int, error) {
	return _FiredrillRevertMessageReceiver.Contract.BalanceOfToken(&_FiredrillRevertMessageReceiver.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) Owner() (common.Address, error) {
	return _FiredrillRevertMessageReceiver.Contract.Owner(&_FiredrillRevertMessageReceiver.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) Owner() (common.Address, error) {
	return _FiredrillRevertMessageReceiver.Contract.Owner(&_FiredrillRevertMessageReceiver.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) PendingOwner() (common.Address, error) {
	return _FiredrillRevertMessageReceiver.Contract.PendingOwner(&_FiredrillRevertMessageReceiver.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) PendingOwner() (common.Address, error) {
	return _FiredrillRevertMessageReceiver.Contract.PendingOwner(&_FiredrillRevertMessageReceiver.CallOpts)
}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) SToRevert(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "s_toRevert")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) SToRevert() (bool, error) {
	return _FiredrillRevertMessageReceiver.Contract.SToRevert(&_FiredrillRevertMessageReceiver.CallOpts)
}

// SToRevert is a free data retrieval call binding the contract method 0x5100fc21.
//
// Solidity: function s_toRevert() view returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) SToRevert() (bool, error) {
	return _FiredrillRevertMessageReceiver.Contract.SToRevert(&_FiredrillRevertMessageReceiver.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FiredrillRevertMessageReceiver.Contract.SupportsInterface(&_FiredrillRevertMessageReceiver.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _FiredrillRevertMessageReceiver.Contract.SupportsInterface(&_FiredrillRevertMessageReceiver.CallOpts, interfaceId)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCaller) TypeAndVersion(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _FiredrillRevertMessageReceiver.contract.Call(opts, &out, "typeAndVersion")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) TypeAndVersion() (string, error) {
	return _FiredrillRevertMessageReceiver.Contract.TypeAndVersion(&_FiredrillRevertMessageReceiver.CallOpts)
}

// TypeAndVersion is a free data retrieval call binding the contract method 0x181f5a77.
//
// Solidity: function typeAndVersion() pure returns(string)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverCallerSession) TypeAndVersion() (string, error) {
	return _FiredrillRevertMessageReceiver.Contract.TypeAndVersion(&_FiredrillRevertMessageReceiver.CallOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.AcceptOwnership(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.AcceptOwnership(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) CcipReceive(opts *bind.TransactOpts, message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "ccipReceive", message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.CcipReceive(&_FiredrillRevertMessageReceiver.TransactOpts, message)
}

// CcipReceive is a paid mutator transaction binding the contract method 0x85572ffb.
//
// Solidity: function ccipReceive((bytes32,uint64,bytes,bytes,(address,uint256)[]) message) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) CcipReceive(message ClientAny2EVMMessage) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.CcipReceive(&_FiredrillRevertMessageReceiver.TransactOpts, message)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.RenounceOwnership(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.RenounceOwnership(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) SetErr(opts *bind.TransactOpts, err []byte) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "setErr", err)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) SetErr(err []byte) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.SetErr(&_FiredrillRevertMessageReceiver.TransactOpts, err)
}

// SetErr is a paid mutator transaction binding the contract method 0x77f5b0e6.
//
// Solidity: function setErr(bytes err) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) SetErr(err []byte) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.SetErr(&_FiredrillRevertMessageReceiver.TransactOpts, err)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) SetRevert(opts *bind.TransactOpts, toRevert bool) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "setRevert", toRevert)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.SetRevert(&_FiredrillRevertMessageReceiver.TransactOpts, toRevert)
}

// SetRevert is a paid mutator transaction binding the contract method 0x8fb5f171.
//
// Solidity: function setRevert(bool toRevert) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) SetRevert(toRevert bool) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.SetRevert(&_FiredrillRevertMessageReceiver.TransactOpts, toRevert)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.TransferOwnership(&_FiredrillRevertMessageReceiver.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.TransferOwnership(&_FiredrillRevertMessageReceiver.TransactOpts, newOwner)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) WithdrawFunds(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "withdrawFunds")
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) WithdrawFunds() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.WithdrawFunds(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// WithdrawFunds is a paid mutator transaction binding the contract method 0x24600fc3.
//
// Solidity: function withdrawFunds() returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) WithdrawFunds() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.WithdrawFunds(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) WithdrawTokens(opts *bind.TransactOpts, token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.Transact(opts, "withdrawTokens", token, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.WithdrawTokens(&_FiredrillRevertMessageReceiver.TransactOpts, token, amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x06b091f9.
//
// Solidity: function withdrawTokens(address token, uint256 amount) returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) WithdrawTokens(token common.Address, amount *big.Int) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.WithdrawTokens(&_FiredrillRevertMessageReceiver.TransactOpts, token, amount)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverSession) Receive() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.Receive(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverTransactorSession) Receive() (*types.Transaction, error) {
	return _FiredrillRevertMessageReceiver.Contract.Receive(&_FiredrillRevertMessageReceiver.TransactOpts)
}

// FiredrillRevertMessageReceiverMessageReceivedIterator is returned from FilterMessageReceived and is used to iterate over the raw logs and unpacked data for MessageReceived events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverMessageReceivedIterator struct {
	Event *FiredrillRevertMessageReceiverMessageReceived // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverMessageReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverMessageReceived)
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
		it.Event = new(FiredrillRevertMessageReceiverMessageReceived)
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
func (it *FiredrillRevertMessageReceiverMessageReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverMessageReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverMessageReceived represents a MessageReceived event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverMessageReceived struct {
	MessageId           [32]byte
	SourceChainSelector uint64
	Sender              []byte
	Data                []byte
	DestTokenAmounts    []ClientEVMTokenAmount
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMessageReceived is a free log retrieval operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterMessageReceived(opts *bind.FilterOpts) (*FiredrillRevertMessageReceiverMessageReceivedIterator, error) {

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverMessageReceivedIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "MessageReceived", logs: logs, sub: sub}, nil
}

// WatchMessageReceived is a free log subscription operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchMessageReceived(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverMessageReceived) (event.Subscription, error) {

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "MessageReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverMessageReceived)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
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

// ParseMessageReceived is a log parse operation binding the contract event 0x707732b700184c0ab3b799f43f03de9b3606a144cfb367f98291044e71972cdc.
//
// Solidity: event MessageReceived(bytes32 messageId, uint64 sourceChainSelector, bytes sender, bytes data, (address,uint256)[] destTokenAmounts)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseMessageReceived(log types.Log) (*FiredrillRevertMessageReceiverMessageReceived, error) {
	event := new(FiredrillRevertMessageReceiverMessageReceived)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "MessageReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator is returned from FilterNativeFundsWithdrawn and is used to iterate over the raw logs and unpacked data for NativeFundsWithdrawn events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator struct {
	Event *FiredrillRevertMessageReceiverNativeFundsWithdrawn // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverNativeFundsWithdrawn)
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
		it.Event = new(FiredrillRevertMessageReceiverNativeFundsWithdrawn)
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
func (it *FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverNativeFundsWithdrawn represents a NativeFundsWithdrawn event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverNativeFundsWithdrawn struct {
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNativeFundsWithdrawn is a free log retrieval operation binding the contract event 0xd50b71a2790ecccf5881141fe9ae079e17c66aace5d50ba383d443ecd398ffa5.
//
// Solidity: event NativeFundsWithdrawn(address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterNativeFundsWithdrawn(opts *bind.FilterOpts, owner []common.Address) (*FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "NativeFundsWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverNativeFundsWithdrawnIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "NativeFundsWithdrawn", logs: logs, sub: sub}, nil
}

// WatchNativeFundsWithdrawn is a free log subscription operation binding the contract event 0xd50b71a2790ecccf5881141fe9ae079e17c66aace5d50ba383d443ecd398ffa5.
//
// Solidity: event NativeFundsWithdrawn(address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchNativeFundsWithdrawn(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverNativeFundsWithdrawn, owner []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "NativeFundsWithdrawn", ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverNativeFundsWithdrawn)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "NativeFundsWithdrawn", log); err != nil {
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

// ParseNativeFundsWithdrawn is a log parse operation binding the contract event 0xd50b71a2790ecccf5881141fe9ae079e17c66aace5d50ba383d443ecd398ffa5.
//
// Solidity: event NativeFundsWithdrawn(address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseNativeFundsWithdrawn(log types.Log) (*FiredrillRevertMessageReceiverNativeFundsWithdrawn, error) {
	event := new(FiredrillRevertMessageReceiverNativeFundsWithdrawn)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "NativeFundsWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillRevertMessageReceiverOwnershipTransferStartedIterator is returned from FilterOwnershipTransferStarted and is used to iterate over the raw logs and unpacked data for OwnershipTransferStarted events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverOwnershipTransferStartedIterator struct {
	Event *FiredrillRevertMessageReceiverOwnershipTransferStarted // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverOwnershipTransferStartedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverOwnershipTransferStarted)
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
		it.Event = new(FiredrillRevertMessageReceiverOwnershipTransferStarted)
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
func (it *FiredrillRevertMessageReceiverOwnershipTransferStartedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverOwnershipTransferStartedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverOwnershipTransferStarted represents a OwnershipTransferStarted event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverOwnershipTransferStarted struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferStarted is a free log retrieval operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterOwnershipTransferStarted(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillRevertMessageReceiverOwnershipTransferStartedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverOwnershipTransferStartedIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "OwnershipTransferStarted", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferStarted is a free log subscription operation binding the contract event 0x38d16b8cac22d99fc7c124b9cd0de2d3fa1faef420bfe791d8c362d765e22700.
//
// Solidity: event OwnershipTransferStarted(address indexed previousOwner, address indexed newOwner)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchOwnershipTransferStarted(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverOwnershipTransferStarted, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "OwnershipTransferStarted", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverOwnershipTransferStarted)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
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
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseOwnershipTransferStarted(log types.Log) (*FiredrillRevertMessageReceiverOwnershipTransferStarted, error) {
	event := new(FiredrillRevertMessageReceiverOwnershipTransferStarted)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "OwnershipTransferStarted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillRevertMessageReceiverOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverOwnershipTransferredIterator struct {
	Event *FiredrillRevertMessageReceiverOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverOwnershipTransferred)
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
		it.Event = new(FiredrillRevertMessageReceiverOwnershipTransferred)
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
func (it *FiredrillRevertMessageReceiverOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverOwnershipTransferred represents a OwnershipTransferred event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FiredrillRevertMessageReceiverOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverOwnershipTransferredIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverOwnershipTransferred)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseOwnershipTransferred(log types.Log) (*FiredrillRevertMessageReceiverOwnershipTransferred, error) {
	event := new(FiredrillRevertMessageReceiverOwnershipTransferred)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillRevertMessageReceiverTokensWithdrawnIterator is returned from FilterTokensWithdrawn and is used to iterate over the raw logs and unpacked data for TokensWithdrawn events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverTokensWithdrawnIterator struct {
	Event *FiredrillRevertMessageReceiverTokensWithdrawn // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverTokensWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverTokensWithdrawn)
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
		it.Event = new(FiredrillRevertMessageReceiverTokensWithdrawn)
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
func (it *FiredrillRevertMessageReceiverTokensWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverTokensWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverTokensWithdrawn represents a TokensWithdrawn event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverTokensWithdrawn struct {
	Token  common.Address
	Owner  common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokensWithdrawn is a free log retrieval operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed token, address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterTokensWithdrawn(opts *bind.FilterOpts, token []common.Address, owner []common.Address) (*FiredrillRevertMessageReceiverTokensWithdrawnIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "TokensWithdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverTokensWithdrawnIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "TokensWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokensWithdrawn is a free log subscription operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed token, address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchTokensWithdrawn(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverTokensWithdrawn, token []common.Address, owner []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "TokensWithdrawn", tokenRule, ownerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverTokensWithdrawn)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
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

// ParseTokensWithdrawn is a log parse operation binding the contract event 0x6337ed398c0e8467698c581374fdce4db14922df487b5a39483079f5f59b60a4.
//
// Solidity: event TokensWithdrawn(address indexed token, address indexed owner, uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseTokensWithdrawn(log types.Log) (*FiredrillRevertMessageReceiverTokensWithdrawn, error) {
	event := new(FiredrillRevertMessageReceiverTokensWithdrawn)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "TokensWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FiredrillRevertMessageReceiverValueReceivedIterator is returned from FilterValueReceived and is used to iterate over the raw logs and unpacked data for ValueReceived events raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverValueReceivedIterator struct {
	Event *FiredrillRevertMessageReceiverValueReceived // Event containing the contract specifics and raw log

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
func (it *FiredrillRevertMessageReceiverValueReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FiredrillRevertMessageReceiverValueReceived)
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
		it.Event = new(FiredrillRevertMessageReceiverValueReceived)
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
func (it *FiredrillRevertMessageReceiverValueReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FiredrillRevertMessageReceiverValueReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FiredrillRevertMessageReceiverValueReceived represents a ValueReceived event raised by the FiredrillRevertMessageReceiver contract.
type FiredrillRevertMessageReceiverValueReceived struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterValueReceived is a free log retrieval operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) FilterValueReceived(opts *bind.FilterOpts) (*FiredrillRevertMessageReceiverValueReceivedIterator, error) {

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.FilterLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return &FiredrillRevertMessageReceiverValueReceivedIterator{contract: _FiredrillRevertMessageReceiver.contract, event: "ValueReceived", logs: logs, sub: sub}, nil
}

// WatchValueReceived is a free log subscription operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) WatchValueReceived(opts *bind.WatchOpts, sink chan<- *FiredrillRevertMessageReceiverValueReceived) (event.Subscription, error) {

	logs, sub, err := _FiredrillRevertMessageReceiver.contract.WatchLogs(opts, "ValueReceived")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FiredrillRevertMessageReceiverValueReceived)
				if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
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

// ParseValueReceived is a log parse operation binding the contract event 0xe12e3b7047ff60a2dd763cf536a43597e5ce7fe7aa7476345bd4cd079912bcef.
//
// Solidity: event ValueReceived(uint256 amount)
func (_FiredrillRevertMessageReceiver *FiredrillRevertMessageReceiverFilterer) ParseValueReceived(log types.Log) (*FiredrillRevertMessageReceiverValueReceived, error) {
	event := new(FiredrillRevertMessageReceiverValueReceived)
	if err := _FiredrillRevertMessageReceiver.contract.UnpackLog(event, "ValueReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
