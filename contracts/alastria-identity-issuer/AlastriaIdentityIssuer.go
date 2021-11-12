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
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"addIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"updateIdentityIssuerEidasLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"deleteIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"getEidasLevel\",\"outputs\":[{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"isIdentityIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsCaller) GetEidasLevel(opts *bind.CallOpts, _identityIssuer common.Address) (uint8, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getEidasLevel", _identityIssuer)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _AlastriaContracts.Contract.GetEidasLevel(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _AlastriaContracts.Contract.GetEidasLevel(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCaller) IsIdentityIssuer(opts *bind.CallOpts, _identityIssuer common.Address) (bool, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "isIdentityIssuer", _identityIssuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityIssuer(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCallerSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityIssuer(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addIdentityIssuer", _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeleteIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deleteIdentityIssuer", _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) UpdateIdentityIssuerEidasLevel(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "updateIdentityIssuerEidasLevel", _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateIdentityIssuerEidasLevel(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateIdentityIssuerEidasLevel(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}
