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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Alastria *AlastriaCaller) IsOwner(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "isOwner")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Alastria *AlastriaSession) IsOwner() (bool, error) {
	return _Alastria.Contract.IsOwner(&_Alastria.CallOpts)
}

// IsOwner is a free data retrieval call binding the contract method 0x8f32d59b.
//
// Solidity: function isOwner() view returns(bool)
func (_Alastria *AlastriaCallerSession) IsOwner() (bool, error) {
	return _Alastria.Contract.IsOwner(&_Alastria.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alastria *AlastriaCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alastria *AlastriaSession) Owner() (common.Address, error) {
	return _Alastria.Contract.Owner(&_Alastria.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Alastria *AlastriaCallerSession) Owner() (common.Address, error) {
	return _Alastria.Contract.Owner(&_Alastria.CallOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Alastria *AlastriaTransactor) Initialize(opts *bind.TransactOpts, sender common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "initialize", sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Alastria *AlastriaSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, sender)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address sender) returns()
func (_Alastria *AlastriaTransactorSession) Initialize(sender common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, sender)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alastria *AlastriaTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alastria *AlastriaSession) RenounceOwnership() (*types.Transaction, error) {
	return _Alastria.Contract.RenounceOwnership(&_Alastria.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Alastria *AlastriaTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Alastria.Contract.RenounceOwnership(&_Alastria.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alastria *AlastriaTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alastria *AlastriaSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.TransferOwnership(&_Alastria.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Alastria *AlastriaTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.TransferOwnership(&_Alastria.TransactOpts, newOwner)
}

// AlastriaOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Alastria contract.
type AlastriaOwnershipTransferredIterator struct {
	Event *AlastriaOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AlastriaOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaOwnershipTransferred)
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
		it.Event = new(AlastriaOwnershipTransferred)
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
func (it *AlastriaOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaOwnershipTransferred represents a OwnershipTransferred event raised by the Alastria contract.
type AlastriaOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Alastria *AlastriaFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AlastriaOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaOwnershipTransferredIterator{contract: _Alastria.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Alastria *AlastriaFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AlastriaOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaOwnershipTransferred)
				if err := _Alastria.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParseOwnershipTransferred(log types.Log) (*AlastriaOwnershipTransferred, error) {
	event := new(AlastriaOwnershipTransferred)
	if err := _Alastria.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
