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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"PublicKeyDeleted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"PublicKeyRevoked\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"publicKeyList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"addKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"revokePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"publicKey\",\"type\":\"string\"}],\"name\":\"deletePublicKey\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getCurrentPublicKey\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"getPublicKeyStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPublicKeyRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"startDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"endDate\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_Alastria *AlastriaCaller) GetCurrentPublicKey(opts *bind.CallOpts, subject common.Address) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getCurrentPublicKey", subject)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_Alastria *AlastriaSession) GetCurrentPublicKey(subject common.Address) (string, error) {
	return _Alastria.Contract.GetCurrentPublicKey(&_Alastria.CallOpts, subject)
}

// GetCurrentPublicKey is a free data retrieval call binding the contract method 0x31f7f259.
//
// Solidity: function getCurrentPublicKey(address subject) view returns(string)
func (_Alastria *AlastriaCallerSession) GetCurrentPublicKey(subject common.Address) (string, error) {
	return _Alastria.Contract.GetCurrentPublicKey(&_Alastria.CallOpts, subject)
}

// GetPublicKeyStatus is a free data retrieval call binding the contract method 0x1226f6a5.
//
// Solidity: function getPublicKeyStatus(address subject, bytes32 publicKey) view returns(bool exists, uint8 status, uint256 startDate, uint256 endDate)
func (_Alastria *AlastriaCaller) GetPublicKeyStatus(opts *bind.CallOpts, subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getPublicKeyStatus", subject, publicKey)

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
func (_Alastria *AlastriaSession) GetPublicKeyStatus(subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	return _Alastria.Contract.GetPublicKeyStatus(&_Alastria.CallOpts, subject, publicKey)
}

// GetPublicKeyStatus is a free data retrieval call binding the contract method 0x1226f6a5.
//
// Solidity: function getPublicKeyStatus(address subject, bytes32 publicKey) view returns(bool exists, uint8 status, uint256 startDate, uint256 endDate)
func (_Alastria *AlastriaCallerSession) GetPublicKeyStatus(subject common.Address, publicKey [32]byte) (struct {
	Exists    bool
	Status    uint8
	StartDate *big.Int
	EndDate   *big.Int
}, error) {
	return _Alastria.Contract.GetPublicKeyStatus(&_Alastria.CallOpts, subject, publicKey)
}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_Alastria *AlastriaCaller) PreviousPublishedVersion(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "previousPublishedVersion")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_Alastria *AlastriaSession) PreviousPublishedVersion() (common.Address, error) {
	return _Alastria.Contract.PreviousPublishedVersion(&_Alastria.CallOpts)
}

// PreviousPublishedVersion is a free data retrieval call binding the contract method 0x6104464f.
//
// Solidity: function previousPublishedVersion() view returns(address)
func (_Alastria *AlastriaCallerSession) PreviousPublishedVersion() (common.Address, error) {
	return _Alastria.Contract.PreviousPublishedVersion(&_Alastria.CallOpts)
}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_Alastria *AlastriaCaller) PublicKeyList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "publicKeyList", arg0, arg1)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_Alastria *AlastriaSession) PublicKeyList(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Alastria.Contract.PublicKeyList(&_Alastria.CallOpts, arg0, arg1)
}

// PublicKeyList is a free data retrieval call binding the contract method 0xc000ceee.
//
// Solidity: function publicKeyList(address , uint256 ) view returns(string)
func (_Alastria *AlastriaCallerSession) PublicKeyList(arg0 common.Address, arg1 *big.Int) (string, error) {
	return _Alastria.Contract.PublicKeyList(&_Alastria.CallOpts, arg0, arg1)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_Alastria *AlastriaCaller) Version(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_Alastria *AlastriaSession) Version() (*big.Int, error) {
	return _Alastria.Contract.Version(&_Alastria.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(int256)
func (_Alastria *AlastriaCallerSession) Version() (*big.Int, error) {
	return _Alastria.Contract.Version(&_Alastria.CallOpts)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_Alastria *AlastriaTransactor) AddKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addKey", publicKey)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_Alastria *AlastriaSession) AddKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.AddKey(&_Alastria.TransactOpts, publicKey)
}

// AddKey is a paid mutator transaction binding the contract method 0x50382c1a.
//
// Solidity: function addKey(string publicKey) returns()
func (_Alastria *AlastriaTransactorSession) AddKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.AddKey(&_Alastria.TransactOpts, publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_Alastria *AlastriaTransactor) DeletePublicKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "deletePublicKey", publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_Alastria *AlastriaSession) DeletePublicKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.DeletePublicKey(&_Alastria.TransactOpts, publicKey)
}

// DeletePublicKey is a paid mutator transaction binding the contract method 0x36241962.
//
// Solidity: function deletePublicKey(string publicKey) returns()
func (_Alastria *AlastriaTransactorSession) DeletePublicKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.DeletePublicKey(&_Alastria.TransactOpts, publicKey)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_Alastria *AlastriaTransactor) Initialize(opts *bind.TransactOpts, _previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "initialize", _previousPublishedVersion)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_Alastria *AlastriaSession) Initialize(_previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, _previousPublishedVersion)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _previousPublishedVersion) returns()
func (_Alastria *AlastriaTransactorSession) Initialize(_previousPublishedVersion common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, _previousPublishedVersion)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_Alastria *AlastriaTransactor) RevokePublicKey(opts *bind.TransactOpts, publicKey string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "revokePublicKey", publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_Alastria *AlastriaSession) RevokePublicKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.RevokePublicKey(&_Alastria.TransactOpts, publicKey)
}

// RevokePublicKey is a paid mutator transaction binding the contract method 0xd16df8ec.
//
// Solidity: function revokePublicKey(string publicKey) returns()
func (_Alastria *AlastriaTransactorSession) RevokePublicKey(publicKey string) (*types.Transaction, error) {
	return _Alastria.Contract.RevokePublicKey(&_Alastria.TransactOpts, publicKey)
}

// AlastriaPublicKeyDeletedIterator is returned from FilterPublicKeyDeleted and is used to iterate over the raw logs and unpacked data for PublicKeyDeleted events raised by the Alastria contract.
type AlastriaPublicKeyDeletedIterator struct {
	Event *AlastriaPublicKeyDeleted // Event containing the contract specifics and raw log

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
func (it *AlastriaPublicKeyDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaPublicKeyDeleted)
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
		it.Event = new(AlastriaPublicKeyDeleted)
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
func (it *AlastriaPublicKeyDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaPublicKeyDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaPublicKeyDeleted represents a PublicKeyDeleted event raised by the Alastria contract.
type AlastriaPublicKeyDeleted struct {
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyDeleted is a free log retrieval operation binding the contract event 0x590ada3b231515dd26dca05728f31024554fc0eca9ef33aa1d623bc4d7d5fd50.
//
// Solidity: event PublicKeyDeleted(string publicKey)
func (_Alastria *AlastriaFilterer) FilterPublicKeyDeleted(opts *bind.FilterOpts) (*AlastriaPublicKeyDeletedIterator, error) {

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "PublicKeyDeleted")
	if err != nil {
		return nil, err
	}
	return &AlastriaPublicKeyDeletedIterator{contract: _Alastria.contract, event: "PublicKeyDeleted", logs: logs, sub: sub}, nil
}

// WatchPublicKeyDeleted is a free log subscription operation binding the contract event 0x590ada3b231515dd26dca05728f31024554fc0eca9ef33aa1d623bc4d7d5fd50.
//
// Solidity: event PublicKeyDeleted(string publicKey)
func (_Alastria *AlastriaFilterer) WatchPublicKeyDeleted(opts *bind.WatchOpts, sink chan<- *AlastriaPublicKeyDeleted) (event.Subscription, error) {

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "PublicKeyDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaPublicKeyDeleted)
				if err := _Alastria.contract.UnpackLog(event, "PublicKeyDeleted", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParsePublicKeyDeleted(log types.Log) (*AlastriaPublicKeyDeleted, error) {
	event := new(AlastriaPublicKeyDeleted)
	if err := _Alastria.contract.UnpackLog(event, "PublicKeyDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaPublicKeyRevokedIterator is returned from FilterPublicKeyRevoked and is used to iterate over the raw logs and unpacked data for PublicKeyRevoked events raised by the Alastria contract.
type AlastriaPublicKeyRevokedIterator struct {
	Event *AlastriaPublicKeyRevoked // Event containing the contract specifics and raw log

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
func (it *AlastriaPublicKeyRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaPublicKeyRevoked)
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
		it.Event = new(AlastriaPublicKeyRevoked)
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
func (it *AlastriaPublicKeyRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaPublicKeyRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaPublicKeyRevoked represents a PublicKeyRevoked event raised by the Alastria contract.
type AlastriaPublicKeyRevoked struct {
	PublicKey string
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPublicKeyRevoked is a free log retrieval operation binding the contract event 0x41d9cf228775966745bd16590f046aff534ab47ea15fc80127ddae02abc8bf34.
//
// Solidity: event PublicKeyRevoked(string publicKey)
func (_Alastria *AlastriaFilterer) FilterPublicKeyRevoked(opts *bind.FilterOpts) (*AlastriaPublicKeyRevokedIterator, error) {

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "PublicKeyRevoked")
	if err != nil {
		return nil, err
	}
	return &AlastriaPublicKeyRevokedIterator{contract: _Alastria.contract, event: "PublicKeyRevoked", logs: logs, sub: sub}, nil
}

// WatchPublicKeyRevoked is a free log subscription operation binding the contract event 0x41d9cf228775966745bd16590f046aff534ab47ea15fc80127ddae02abc8bf34.
//
// Solidity: event PublicKeyRevoked(string publicKey)
func (_Alastria *AlastriaFilterer) WatchPublicKeyRevoked(opts *bind.WatchOpts, sink chan<- *AlastriaPublicKeyRevoked) (event.Subscription, error) {

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "PublicKeyRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaPublicKeyRevoked)
				if err := _Alastria.contract.UnpackLog(event, "PublicKeyRevoked", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParsePublicKeyRevoked(log types.Log) (*AlastriaPublicKeyRevoked, error) {
	event := new(AlastriaPublicKeyRevoked)
	if err := _Alastria.contract.UnpackLog(event, "PublicKeyRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
