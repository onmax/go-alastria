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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"PublicKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"PublicKeyRevoked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeyList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"revokePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"deletePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getCurrentPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"getPublicKeyStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPublicKeyRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) GetCurrentPublicKey(opts *bind.CallOpts, subject common.Address) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getCurrentPublicKey", subject)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) GetCurrentPublicKey(subject common.Address) (string, error) {
	return _AlastriaContracts.Contract.GetCurrentPublicKey(&_AlastriaContracts.CallOpts, subject)
}

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetCurrentPublicKey(subject common.Address) (string, error) {
	return _AlastriaContracts.Contract.GetCurrentPublicKey(&_AlastriaContracts.CallOpts, subject)
}

// GetPublicKeyStatus is a free data retrieval call binding the contract method 0x1226f6a5.
//
// Solidity: function getPublicKeyStatus(address subject, bytes32 publicKey) view returns(bool exists, uint8 status, uint256 startDate, uint256 endDate)
func (_AlastriaContracts *AlastriaContractsCaller) GetPublicKeyStatus(opts *bind.CallOpts, subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getPublicKeyStatus", subject, publicKey)

	outstruct := new(struct {
		Exists    bool
		Status    uint8
		StartDate *big.Int
		EndDate   *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.StartDate = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.EndDate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetPublicKeyStatus is a free data retrieval call binding the contract method 0x1226f6a5.
//
// Solidity: function getPublicKeyStatus(address subject, bytes32 publicKey) view returns(bool exists, uint8 status, uint256 startDate, uint256 endDate)
func (_AlastriaContracts *AlastriaContractsSession) GetPublicKeyStatus(subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	return _AlastriaContracts.Contract.GetPublicKeyStatus(&_AlastriaContracts.CallOpts, subject, publicKey)
}

// GetPublicKeyStatus is a free data retrieval call binding the contract method 0x1226f6a5.
//
// Solidity: function getPublicKeyStatus(address subject, bytes32 publicKey) view returns(bool exists, uint8 status, uint256 startDate, uint256 endDate)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetPublicKeyStatus(subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	return _AlastriaContracts.Contract.GetPublicKeyStatus(&_AlastriaContracts.CallOpts, subject, publicKey)
}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) PreviousPublishedVersion(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "previousPublishedVersion")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) PreviousPublishedVersion() (common.Address, error) {
	return _AlastriaContracts.Contract.PreviousPublishedVersion(&_AlastriaContracts.CallOpts)
}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) PreviousPublishedVersion() (common.Address, error) {
	return _AlastriaContracts.Contract.PreviousPublishedVersion(&_AlastriaContracts.CallOpts)
}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) PublicKeyList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "publicKeyList", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) PublicKeyList(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _AlastriaContracts.Contract.PublicKeyList(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) PublicKeyList(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _AlastriaContracts.Contract.PublicKeyList(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_AlastriaContracts *AlastriaContractsCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_AlastriaContracts *AlastriaContractsSession) Version() (*big.Int, error) {
	return _AlastriaContracts.Contract.Version(&_AlastriaContracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_AlastriaContracts *AlastriaContractsCallerSession) Version() (*big.Int, error) {
	return _AlastriaContracts.Contract.Version(&_AlastriaContracts.CallOpts)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addKey", publicKey)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeletePublicKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deletePublicKey", publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeletePublicKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeletePublicKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeletePublicKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeletePublicKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) Initialize(opts *bind.TransactOpts, _previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "initialize", _previousPublishedVersion)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_AlastriaContracts *AlastriaContractsSession) Initialize(_previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Initialize(&_AlastriaContracts.TransactOpts, _previousPublishedVersion)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) Initialize(_previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Initialize(&_AlastriaContracts.TransactOpts, _previousPublishedVersion)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) RevokePublicKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "revokePublicKey", publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsSession) RevokePublicKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.RevokePublicKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) RevokePublicKey(publicKey string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.RevokePublicKey(&_AlastriaContracts.TransactOpts, publicKey)
}

// AlastriaContractsPublicKeyDeletedIterator is returned from FilterPublicKeyDeleted and is used to iterate over the raw logs and unpacked data for PublicKeyDeleted events raised by the AlastriaContracts contract.
type AlastriaContractsPublicKeyDeletedIterator struct {
	Event *AlastriaContractsPublicKeyDeleted // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsPublicKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsPublicKeyDeleted)
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
		it.Event = new(AlastriaContractsPublicKeyDeleted)
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
func (it *AlastriaContractsPublicKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsPublicKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsPublicKeyDeleted represents a PublicKeyDeleted event raised by the AlastriaContracts contract.
type AlastriaContractsPublicKeyDeleted struct {
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyDeleted is a free log retrieval operation binding the contract event 0x590ada3b231515dd26dca05728f31024554fc0eca9ef33aa1d623bc4d7d5fd50.
//
// Solidity: event PublicKeyDeleted(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterPublicKeyDeleted(opts *bind.FilterOpts) (*AlastriaContractsPublicKeyDeletedIterator, error) {

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "PublicKeyDeleted")
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsPublicKeyDeletedIterator{contract: _AlastriaContracts.contract, event: "PublicKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchPublicKeyDeleted is a free log subscription operation binding the contract event 0x590ada3b231515dd26dca05728f31024554fc0eca9ef33aa1d623bc4d7d5fd50.
//
// Solidity: event PublicKeyDeleted(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchPublicKeyDeleted(opts *bind.WatchOpts, sink chan<- *AlastriaContractsPublicKeyDeleted) (event.Subscription, error) {

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "PublicKeyDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsPublicKeyDeleted)
				if err := _AlastriaContracts.contract.UnpackLog(event, "PublicKeyDeleted", log); err != nil {
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

// ParsePublicKeyDeleted is a log parse operation binding the contract event 0x590ada3b231515dd26dca05728f31024554fc0eca9ef33aa1d623bc4d7d5fd50.
//
// Solidity: event PublicKeyDeleted(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) ParsePublicKeyDeleted(log types.Log) (*AlastriaContractsPublicKeyDeleted, error) {
	event := new(AlastriaContractsPublicKeyDeleted)
	if err := _AlastriaContracts.contract.UnpackLog(event, "PublicKeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaContractsPublicKeyRevokedIterator is returned from FilterPublicKeyRevoked and is used to iterate over the raw logs and unpacked data for PublicKeyRevoked events raised by the AlastriaContracts contract.
type AlastriaContractsPublicKeyRevokedIterator struct {
	Event *AlastriaContractsPublicKeyRevoked // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsPublicKeyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsPublicKeyRevoked)
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
		it.Event = new(AlastriaContractsPublicKeyRevoked)
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
func (it *AlastriaContractsPublicKeyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsPublicKeyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsPublicKeyRevoked represents a PublicKeyRevoked event raised by the AlastriaContracts contract.
type AlastriaContractsPublicKeyRevoked struct {
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRevoked is a free log retrieval operation binding the contract event 0x41d9cf228775966745bd16590f046aff534ab47ea15fc80127ddae02abc8bf34.
//
// Solidity: event PublicKeyRevoked(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterPublicKeyRevoked(opts *bind.FilterOpts) (*AlastriaContractsPublicKeyRevokedIterator, error) {

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "PublicKeyRevoked")
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsPublicKeyRevokedIterator{contract: _AlastriaContracts.contract, event: "PublicKeyRevoked", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRevoked is a free log subscription operation binding the contract event 0x41d9cf228775966745bd16590f046aff534ab47ea15fc80127ddae02abc8bf34.
//
// Solidity: event PublicKeyRevoked(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchPublicKeyRevoked(opts *bind.WatchOpts, sink chan<- *AlastriaContractsPublicKeyRevoked) (event.Subscription, error) {

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "PublicKeyRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsPublicKeyRevoked)
				if err := _AlastriaContracts.contract.UnpackLog(event, "PublicKeyRevoked", log); err != nil {
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

// ParsePublicKeyRevoked is a log parse operation binding the contract event 0x41d9cf228775966745bd16590f046aff534ab47ea15fc80127ddae02abc8bf34.
//
// Solidity: event PublicKeyRevoked(string publicKey)
func (_AlastriaContracts *AlastriaContractsFilterer) ParsePublicKeyRevoked(log types.Log) (*AlastriaContractsPublicKeyRevoked, error) {
	event := new(AlastriaContractsPublicKeyRevoked)
	if err := _AlastriaContracts.contract.UnpackLog(event, "PublicKeyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
