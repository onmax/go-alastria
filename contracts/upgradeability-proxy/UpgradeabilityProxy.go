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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AlastriaContracts *AlastriaContractsTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _AlastriaContracts.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AlastriaContracts *AlastriaContractsSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Fallback(&_AlastriaContracts.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Fallback(&_AlastriaContracts.TransactOpts, calldata)
}

// AlastriaContractsUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AlastriaContracts contract.
type AlastriaContractsUpgradedIterator struct {
	Event *AlastriaContractsUpgraded // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsUpgraded)
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
		it.Event = new(AlastriaContractsUpgraded)
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
func (it *AlastriaContractsUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsUpgraded represents a Upgraded event raised by the AlastriaContracts contract.
type AlastriaContractsUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AlastriaContractsUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsUpgradedIterator{contract: _AlastriaContracts.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AlastriaContractsUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsUpgraded)
				if err := _AlastriaContracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AlastriaContracts *AlastriaContractsFilterer) ParseUpgraded(log types.Log) (*AlastriaContractsUpgraded, error) {
	event := new(AlastriaContractsUpgraded)
	if err := _AlastriaContracts.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
