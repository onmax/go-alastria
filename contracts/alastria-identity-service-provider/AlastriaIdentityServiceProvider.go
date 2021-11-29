// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alastriaContracts

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
)

// AlastriaContractsMetaData contains all meta data concerning the AlastriaContracts contract.
var AlastriaContractsMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"addIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"deleteIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"isIdentityServiceProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AlastriaContractsABI is the input ABI used to generate the binding from.
// Deprecated: Use AlastriaContractsMetaData.ABI instead.
var AlastriaContractsABI = AlastriaContractsMetaData.ABI

// AlastriaContracts is an auto generated Go binding around an Ethereum contract.
type AlastriaContracts struct {
	AlastriaContractsCaller     // Read-only binding to the contract
	AlastriaContractsTransactor // Write-only binding to the contract
	AlastriaContractsFilterer   // Log filterer for contract events
}

// AlastriaContractsCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlastriaContractsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaContractsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlastriaContractsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaContractsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlastriaContractsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaContractsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlastriaContractsSession struct {
	Contract     *AlastriaContracts // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// AlastriaContractsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlastriaContractsCallerSession struct {
	Contract *AlastriaContractsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// AlastriaContractsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlastriaContractsTransactorSession struct {
	Contract     *AlastriaContractsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// AlastriaContractsRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlastriaContractsRaw struct {
	Contract *AlastriaContracts // Generic contract binding to access the raw methods on
}

// AlastriaContractsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlastriaContractsCallerRaw struct {
	Contract *AlastriaContractsCaller // Generic read-only contract binding to access the raw methods on
}

// AlastriaContractsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlastriaContractsTransactorRaw struct {
	Contract *AlastriaContractsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlastriaContracts creates a new instance of AlastriaContracts, bound to a specific deployed contract.
func NewAlastriaContracts(address common.Address, backend bind.ContractBackend) (*AlastriaContracts, error) {
	contract, err := bindAlastriaContracts(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AlastriaContracts{AlastriaContractsCaller: AlastriaContractsCaller{contract: contract}, AlastriaContractsTransactor: AlastriaContractsTransactor{contract: contract}, AlastriaContractsFilterer: AlastriaContractsFilterer{contract: contract}}, nil
}

// NewAlastriaContractsCaller creates a new read-only instance of AlastriaContracts, bound to a specific deployed contract.
func NewAlastriaContractsCaller(address common.Address, caller bind.ContractCaller) (*AlastriaContractsCaller, error) {
	contract, err := bindAlastriaContracts(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsCaller{contract: contract}, nil
}

// NewAlastriaContractsTransactor creates a new write-only instance of AlastriaContracts, bound to a specific deployed contract.
func NewAlastriaContractsTransactor(address common.Address, transactor bind.ContractTransactor) (*AlastriaContractsTransactor, error) {
	contract, err := bindAlastriaContracts(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsTransactor{contract: contract}, nil
}

// NewAlastriaContractsFilterer creates a new log filterer instance of AlastriaContracts, bound to a specific deployed contract.
func NewAlastriaContractsFilterer(address common.Address, filterer bind.ContractFilterer) (*AlastriaContractsFilterer, error) {
	contract, err := bindAlastriaContracts(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsFilterer{contract: contract}, nil
}

// bindAlastriaContracts binds a generic wrapper to an already deployed contract.
func bindAlastriaContracts(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AlastriaContractsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlastriaContracts *AlastriaContractsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlastriaContracts.Contract.AlastriaContractsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlastriaContracts *AlastriaContractsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AlastriaContractsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlastriaContracts *AlastriaContractsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AlastriaContractsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AlastriaContracts *AlastriaContractsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AlastriaContracts.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AlastriaContracts *AlastriaContractsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AlastriaContracts *AlastriaContractsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.contract.Transact(opts, method, params...)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCaller) IsIdentityServiceProvider(opts *bind.CallOpts, _identityServiceProvider common.Address) (bool, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "isIdentityServiceProvider", _identityServiceProvider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityServiceProvider(&_AlastriaContracts.CallOpts, _identityServiceProvider)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCallerSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityServiceProvider(&_AlastriaContracts.CallOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addIdentityServiceProvider", _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeleteIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deleteIdentityServiceProvider", _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}
