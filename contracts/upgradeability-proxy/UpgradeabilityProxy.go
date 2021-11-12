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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_logic\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"}]",
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

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Alastria *AlastriaTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Alastria.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Alastria *AlastriaSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Alastria.Contract.Fallback(&_Alastria.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Alastria *AlastriaTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Alastria.Contract.Fallback(&_Alastria.TransactOpts, calldata)
}

// AlastriaUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Alastria contract.
type AlastriaUpgradedIterator struct {
	Event *AlastriaUpgraded // Event containing the contract specifics and raw log

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
func (it *AlastriaUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaUpgraded)
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
		it.Event = new(AlastriaUpgraded)
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
func (it *AlastriaUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaUpgraded represents a Upgraded event raised by the Alastria contract.
type AlastriaUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Alastria *AlastriaFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AlastriaUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaUpgradedIterator{contract: _Alastria.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Alastria *AlastriaFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AlastriaUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaUpgraded)
				if err := _Alastria.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParseUpgraded(log types.Log) (*AlastriaUpgraded, error) {
	event := new(AlastriaUpgraded)
	if err := _Alastria.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
