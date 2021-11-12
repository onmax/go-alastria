// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package alastria

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

// AlastriaMetaData contains all meta data concerning the Alastria contract.
var AlastriaMetaData = &bind.MetaData{
	ABI: "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"addIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"deleteIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"isIdentityServiceProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// AlastriaABI is the input ABI used to generate the binding from.
// Deprecated: Use AlastriaMetaData.ABI instead.
var AlastriaABI = AlastriaMetaData.ABI

// Alastria is an auto generated Go binding around an Ethereum contract.
type Alastria struct {
	AlastriaCaller     // Read-only binding to the contract
	AlastriaTransactor // Write-only binding to the contract
	AlastriaFilterer   // Log filterer for contract events
}

// AlastriaCaller is an auto generated read-only Go binding around an Ethereum contract.
type AlastriaCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AlastriaTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AlastriaFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AlastriaSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AlastriaSession struct {
	Contract     *Alastria         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AlastriaCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AlastriaCallerSession struct {
	Contract *AlastriaCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// AlastriaTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AlastriaTransactorSession struct {
	Contract     *AlastriaTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AlastriaRaw is an auto generated low-level Go binding around an Ethereum contract.
type AlastriaRaw struct {
	Contract *Alastria // Generic contract binding to access the raw methods on
}

// AlastriaCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AlastriaCallerRaw struct {
	Contract *AlastriaCaller // Generic read-only contract binding to access the raw methods on
}

// AlastriaTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AlastriaTransactorRaw struct {
	Contract *AlastriaTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAlastria creates a new instance of Alastria, bound to a specific deployed contract.
func NewAlastria(address common.Address, backend bind.ContractBackend) (*Alastria, error) {
	contract, err := bindAlastria(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Alastria{AlastriaCaller: AlastriaCaller{contract: contract}, AlastriaTransactor: AlastriaTransactor{contract: contract}, AlastriaFilterer: AlastriaFilterer{contract: contract}}, nil
}

// NewAlastriaCaller creates a new read-only instance of Alastria, bound to a specific deployed contract.
func NewAlastriaCaller(address common.Address, caller bind.ContractCaller) (*AlastriaCaller, error) {
	contract, err := bindAlastria(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AlastriaCaller{contract: contract}, nil
}

// NewAlastriaTransactor creates a new write-only instance of Alastria, bound to a specific deployed contract.
func NewAlastriaTransactor(address common.Address, transactor bind.ContractTransactor) (*AlastriaTransactor, error) {
	contract, err := bindAlastria(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AlastriaTransactor{contract: contract}, nil
}

// NewAlastriaFilterer creates a new log filterer instance of Alastria, bound to a specific deployed contract.
func NewAlastriaFilterer(address common.Address, filterer bind.ContractFilterer) (*AlastriaFilterer, error) {
	contract, err := bindAlastria(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AlastriaFilterer{contract: contract}, nil
}

// bindAlastria binds a generic wrapper to an already deployed contract.
func bindAlastria(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AlastriaABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alastria *AlastriaRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alastria.Contract.AlastriaCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alastria *AlastriaRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alastria.Contract.AlastriaTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alastria *AlastriaRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alastria.Contract.AlastriaTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Alastria *AlastriaCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Alastria.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Alastria *AlastriaTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alastria.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Alastria *AlastriaTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Alastria.Contract.contract.Transact(opts, method, params...)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaCaller) IsIdentityServiceProvider(opts *bind.CallOpts, _identityServiceProvider common.Address) (bool, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "isIdentityServiceProvider", _identityServiceProvider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityServiceProvider(&_Alastria.CallOpts, _identityServiceProvider)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaCallerSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityServiceProvider(&_Alastria.CallOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactor) AddIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addIdentityServiceProvider", _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactorSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactor) DeleteIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "deleteIdentityServiceProvider", _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactorSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}
