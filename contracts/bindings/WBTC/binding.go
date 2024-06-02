// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractWBTC

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

// ContractWBTCMetaData contains all meta data concerning the ContractWBTC contract.
var ContractWBTCMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_taskManagerAddress\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false}]",
	Bin: "0x60a060405234801561001057600080fd5b5060405161071438038061071483398101604081905261002f91610040565b6001600160a01b0316608052610070565b60006020828403121561005257600080fd5b81516001600160a01b038116811461006957600080fd5b9392505050565b60805161068961008b60003960006101aa01526106896000f3fe608060405234801561001057600080fd5b50600436106100885760003560e01c806370a082311161005b57806370a08231146100f25780639dc29fac1461011b578063a9059cbb1461012e578063dd62ed3e1461014157600080fd5b8063095ea7b31461008d57806318160ddd146100b857806323b872dd146100ca57806340c10f19146100dd575b600080fd5b6100a361009b366004610572565b600192915050565b60405190151581526020015b60405180910390f35b6002545b6040519081526020016100af565b6100a36100d836600461059c565b61017a565b6100f06100eb366004610572565b610191565b005b6100bc6101003660046105d8565b6001600160a01b031660009081526020819052604090205490565b6100f0610129366004610572565b61019f565b6100a361013c366004610572565b6101de565b6100bc61014f3660046105fa565b6001600160a01b03918216600090815260016020908152604080832093909416825291909152205490565b60006101878484846101e8565b5060019392505050565b61019b8282610393565b5050565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146101d457600080fd5b61019b8282610452565b6000336101878185855b6001600160a01b0383166102515760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b60648201526084015b60405180910390fd5b6001600160a01b0382166102b35760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610248565b6001600160a01b03831660009081526020819052604090205481111561032a5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610248565b6001600160a01b038381166000818152602081815260408083208054879003905593861680835291849020805486019055925184815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35b505050565b6001600160a01b0382166103e95760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f2061646472657373006044820152606401610248565b80600260008282546103fb919061062d565b90915550506001600160a01b038216600081815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b6001600160a01b0382166104b25760405162461bcd60e51b815260206004820152602160248201527f45524332303a206275726e2066726f6d20746865207a65726f206164647265736044820152607360f81b6064820152608401610248565b6001600160a01b038216600090815260208190526040902054818110156105265760405162461bcd60e51b815260206004820152602260248201527f45524332303a206275726e20616d6f756e7420657863656564732062616c616e604482015261636560f01b6064820152608401610248565b6001600160a01b0383166000908152602081905260408120838303905560028054849003905561038e9084908483565b80356001600160a01b038116811461056d57600080fd5b919050565b6000806040838503121561058557600080fd5b61058e83610556565b946020939093013593505050565b6000806000606084860312156105b157600080fd5b6105ba84610556565b92506105c860208501610556565b9150604084013590509250925092565b6000602082840312156105ea57600080fd5b6105f382610556565b9392505050565b6000806040838503121561060d57600080fd5b61061683610556565b915061062460208401610556565b90509250929050565b6000821982111561064e57634e487b7160e01b600052601160045260246000fd5b50019056fea2646970667358221220071092b1ce5a9a86de8585f8a9fca29434d10da12f1315d73d8821fd20ab7dfe64736f6c634300080c0033",
}

// ContractWBTCABI is the input ABI used to generate the binding from.
// Deprecated: Use ContractWBTCMetaData.ABI instead.
var ContractWBTCABI = ContractWBTCMetaData.ABI

// ContractWBTCBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ContractWBTCMetaData.Bin instead.
var ContractWBTCBin = ContractWBTCMetaData.Bin

// DeployContractWBTC deploys a new Ethereum contract, binding an instance of ContractWBTC to it.
func DeployContractWBTC(auth *bind.TransactOpts, backend bind.ContractBackend, _taskManagerAddress common.Address) (common.Address, *types.Transaction, *ContractWBTC, error) {
	parsed, err := ContractWBTCMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ContractWBTCBin), backend, _taskManagerAddress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ContractWBTC{ContractWBTCCaller: ContractWBTCCaller{contract: contract}, ContractWBTCTransactor: ContractWBTCTransactor{contract: contract}, ContractWBTCFilterer: ContractWBTCFilterer{contract: contract}}, nil
}

// ContractWBTC is an auto generated Go binding around an Ethereum contract.
type ContractWBTC struct {
	ContractWBTCCaller     // Read-only binding to the contract
	ContractWBTCTransactor // Write-only binding to the contract
	ContractWBTCFilterer   // Log filterer for contract events
}

// ContractWBTCCaller is an auto generated read-only Go binding around an Ethereum contract.
type ContractWBTCCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWBTCTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ContractWBTCTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWBTCFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ContractWBTCFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ContractWBTCSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ContractWBTCSession struct {
	Contract     *ContractWBTC     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ContractWBTCCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ContractWBTCCallerSession struct {
	Contract *ContractWBTCCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ContractWBTCTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ContractWBTCTransactorSession struct {
	Contract     *ContractWBTCTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ContractWBTCRaw is an auto generated low-level Go binding around an Ethereum contract.
type ContractWBTCRaw struct {
	Contract *ContractWBTC // Generic contract binding to access the raw methods on
}

// ContractWBTCCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ContractWBTCCallerRaw struct {
	Contract *ContractWBTCCaller // Generic read-only contract binding to access the raw methods on
}

// ContractWBTCTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ContractWBTCTransactorRaw struct {
	Contract *ContractWBTCTransactor // Generic write-only contract binding to access the raw methods on
}

// NewContractWBTC creates a new instance of ContractWBTC, bound to a specific deployed contract.
func NewContractWBTC(address common.Address, backend bind.ContractBackend) (*ContractWBTC, error) {
	contract, err := bindContractWBTC(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ContractWBTC{ContractWBTCCaller: ContractWBTCCaller{contract: contract}, ContractWBTCTransactor: ContractWBTCTransactor{contract: contract}, ContractWBTCFilterer: ContractWBTCFilterer{contract: contract}}, nil
}

// NewContractWBTCCaller creates a new read-only instance of ContractWBTC, bound to a specific deployed contract.
func NewContractWBTCCaller(address common.Address, caller bind.ContractCaller) (*ContractWBTCCaller, error) {
	contract, err := bindContractWBTC(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWBTCCaller{contract: contract}, nil
}

// NewContractWBTCTransactor creates a new write-only instance of ContractWBTC, bound to a specific deployed contract.
func NewContractWBTCTransactor(address common.Address, transactor bind.ContractTransactor) (*ContractWBTCTransactor, error) {
	contract, err := bindContractWBTC(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ContractWBTCTransactor{contract: contract}, nil
}

// NewContractWBTCFilterer creates a new log filterer instance of ContractWBTC, bound to a specific deployed contract.
func NewContractWBTCFilterer(address common.Address, filterer bind.ContractFilterer) (*ContractWBTCFilterer, error) {
	contract, err := bindContractWBTC(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ContractWBTCFilterer{contract: contract}, nil
}

// bindContractWBTC binds a generic wrapper to an already deployed contract.
func bindContractWBTC(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ContractWBTCMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWBTC *ContractWBTCRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWBTC.Contract.ContractWBTCCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWBTC *ContractWBTCRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWBTC.Contract.ContractWBTCTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWBTC *ContractWBTCRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWBTC.Contract.ContractWBTCTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ContractWBTC *ContractWBTCCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ContractWBTC.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ContractWBTC *ContractWBTCTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ContractWBTC.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ContractWBTC *ContractWBTCTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ContractWBTC.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractWBTC *ContractWBTCCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractWBTC.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractWBTC *ContractWBTCSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractWBTC.Contract.Allowance(&_ContractWBTC.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ContractWBTC *ContractWBTCCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ContractWBTC.Contract.Allowance(&_ContractWBTC.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractWBTC *ContractWBTCCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ContractWBTC.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractWBTC *ContractWBTCSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractWBTC.Contract.BalanceOf(&_ContractWBTC.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ContractWBTC *ContractWBTCCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ContractWBTC.Contract.BalanceOf(&_ContractWBTC.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractWBTC *ContractWBTCCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ContractWBTC.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractWBTC *ContractWBTCSession) TotalSupply() (*big.Int, error) {
	return _ContractWBTC.Contract.TotalSupply(&_ContractWBTC.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ContractWBTC *ContractWBTCCallerSession) TotalSupply() (*big.Int, error) {
	return _ContractWBTC.Contract.TotalSupply(&_ContractWBTC.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractWBTC *ContractWBTCTransactor) Approve(opts *bind.TransactOpts, arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.contract.Transact(opts, "approve", arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractWBTC *ContractWBTCSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Approve(&_ContractWBTC.TransactOpts, arg0, arg1)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address , uint256 ) returns(bool)
func (_ContractWBTC *ContractWBTCTransactorSession) Approve(arg0 common.Address, arg1 *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Approve(&_ContractWBTC.TransactOpts, arg0, arg1)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCTransactor) Burn(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.contract.Transact(opts, "burn", account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Burn(&_ContractWBTC.TransactOpts, account, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCTransactorSession) Burn(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Burn(&_ContractWBTC.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCTransactor) Mint(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.contract.Transact(opts, "mint", account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Mint(&_ContractWBTC.TransactOpts, account, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address account, uint256 amount) returns()
func (_ContractWBTC *ContractWBTCTransactorSession) Mint(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Mint(&_ContractWBTC.TransactOpts, account, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Transfer(&_ContractWBTC.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.Transfer(&_ContractWBTC.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.TransferFrom(&_ContractWBTC.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_ContractWBTC *ContractWBTCTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ContractWBTC.Contract.TransferFrom(&_ContractWBTC.TransactOpts, from, to, amount)
}

// ContractWBTCApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ContractWBTC contract.
type ContractWBTCApprovalIterator struct {
	Event *ContractWBTCApproval // Event containing the contract specifics and raw log

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
func (it *ContractWBTCApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWBTCApproval)
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
		it.Event = new(ContractWBTCApproval)
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
func (it *ContractWBTCApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWBTCApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWBTCApproval represents a Approval event raised by the ContractWBTC contract.
type ContractWBTCApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ContractWBTCApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractWBTC.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ContractWBTCApprovalIterator{contract: _ContractWBTC.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ContractWBTCApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ContractWBTC.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWBTCApproval)
				if err := _ContractWBTC.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) ParseApproval(log types.Log) (*ContractWBTCApproval, error) {
	event := new(ContractWBTCApproval)
	if err := _ContractWBTC.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ContractWBTCTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ContractWBTC contract.
type ContractWBTCTransferIterator struct {
	Event *ContractWBTCTransfer // Event containing the contract specifics and raw log

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
func (it *ContractWBTCTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ContractWBTCTransfer)
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
		it.Event = new(ContractWBTCTransfer)
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
func (it *ContractWBTCTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ContractWBTCTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ContractWBTCTransfer represents a Transfer event raised by the ContractWBTC contract.
type ContractWBTCTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ContractWBTCTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractWBTC.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ContractWBTCTransferIterator{contract: _ContractWBTC.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ContractWBTCTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ContractWBTC.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ContractWBTCTransfer)
				if err := _ContractWBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ContractWBTC *ContractWBTCFilterer) ParseTransfer(log types.Log) (*ContractWBTCTransfer, error) {
	event := new(ContractWBTCTransfer)
	if err := _ContractWBTC.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
