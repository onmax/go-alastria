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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"PresentationUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subjectPresentationListRegistry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subjectPresentationRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"name\":\"addSubjectPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateSubjectPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"}],\"name\":\"getSubjectPresentationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getSubjectPresentationList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateReceiverPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiverPresentationHash\",\"type\":\"bytes32\"}],\"name\":\"getReceiverPresentationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"subjectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"receiverStatus\",\"type\":\"uint8\"}],\"name\":\"getPresentationStatus\",\"outputs\":[{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsCaller) GetPresentationStatus(opts *bind.CallOpts, subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getPresentationStatus", subjectStatus, receiverStatus)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsSession) GetPresentationStatus(subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	return _AlastriaContracts.Contract.GetPresentationStatus(&_AlastriaContracts.CallOpts, subjectStatus, receiverStatus)
}

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetPresentationStatus(subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	return _AlastriaContracts.Contract.GetPresentationStatus(&_AlastriaContracts.CallOpts, subjectStatus, receiverStatus)
}

// GetReceiverPresentationStatus is a free data retrieval call binding the contract method 0x9373014a.
//
// Solidity: function getReceiverPresentationStatus(address receiver, bytes32 receiverPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCaller) GetReceiverPresentationStatus(opts *bind.CallOpts, receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getReceiverPresentationStatus", receiver, receiverPresentationHash)

	outstruct := new(struct {
		Exists bool
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetReceiverPresentationStatus is a free data retrieval call binding the contract method 0x9373014a.
//
// Solidity: function getReceiverPresentationStatus(address receiver, bytes32 receiverPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsSession) GetReceiverPresentationStatus(receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetReceiverPresentationStatus(&_AlastriaContracts.CallOpts, receiver, receiverPresentationHash)
}

// GetReceiverPresentationStatus is a free data retrieval call binding the contract method 0x9373014a.
//
// Solidity: function getReceiverPresentationStatus(address receiver, bytes32 receiverPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetReceiverPresentationStatus(receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetReceiverPresentationStatus(&_AlastriaContracts.CallOpts, receiver, receiverPresentationHash)
}

// GetSubjectPresentationList is a free data retrieval call binding the contract method 0xc00dcd0e.
//
// Solidity: function getSubjectPresentationList(address subject) view returns(uint256, bytes32[])
func (_AlastriaContracts *AlastriaContractsCaller) GetSubjectPresentationList(opts *bind.CallOpts, subject common.Address) (*big.Int, [][32]byte, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getSubjectPresentationList", subject)

	if err != nil {
		return *new(*big.Int), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return out0, out1, err

}

// GetSubjectPresentationList is a free data retrieval call binding the contract method 0xc00dcd0e.
//
// Solidity: function getSubjectPresentationList(address subject) view returns(uint256, bytes32[])
func (_AlastriaContracts *AlastriaContractsSession) GetSubjectPresentationList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _AlastriaContracts.Contract.GetSubjectPresentationList(&_AlastriaContracts.CallOpts, subject)
}

// GetSubjectPresentationList is a free data retrieval call binding the contract method 0xc00dcd0e.
//
// Solidity: function getSubjectPresentationList(address subject) view returns(uint256, bytes32[])
func (_AlastriaContracts *AlastriaContractsCallerSession) GetSubjectPresentationList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _AlastriaContracts.Contract.GetSubjectPresentationList(&_AlastriaContracts.CallOpts, subject)
}

// GetSubjectPresentationStatus is a free data retrieval call binding the contract method 0x5205080f.
//
// Solidity: function getSubjectPresentationStatus(address subject, bytes32 subjectPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCaller) GetSubjectPresentationStatus(opts *bind.CallOpts, subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getSubjectPresentationStatus", subject, subjectPresentationHash)

	outstruct := new(struct {
		Exists bool
		Status uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)

	return *outstruct, err

}

// GetSubjectPresentationStatus is a free data retrieval call binding the contract method 0x5205080f.
//
// Solidity: function getSubjectPresentationStatus(address subject, bytes32 subjectPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsSession) GetSubjectPresentationStatus(subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetSubjectPresentationStatus(&_AlastriaContracts.CallOpts, subject, subjectPresentationHash)
}

// GetSubjectPresentationStatus is a free data retrieval call binding the contract method 0x5205080f.
//
// Solidity: function getSubjectPresentationStatus(address subject, bytes32 subjectPresentationHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetSubjectPresentationStatus(subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetSubjectPresentationStatus(&_AlastriaContracts.CallOpts, subject, subjectPresentationHash)
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

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCaller) SubjectPresentationListRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "subjectPresentationListRegistry", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsSession) SubjectPresentationListRegistry(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.SubjectPresentationListRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCallerSession) SubjectPresentationListRegistry(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.SubjectPresentationListRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectPresentationRegistry is a free data retrieval call binding the contract method 0x1e54e325.
//
// Solidity: function subjectPresentationRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_AlastriaContracts *AlastriaContractsCaller) SubjectPresentationRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "subjectPresentationRegistry", arg0, arg1)

	outstruct := new(struct {
		Exists bool
		Status uint8
		URI    string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Exists = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.Status = *abi.ConvertType(out[1], new(uint8)).(*uint8)
	outstruct.URI = *abi.ConvertType(out[2], new(string)).(*string)

	return *outstruct, err

}

// SubjectPresentationRegistry is a free data retrieval call binding the contract method 0x1e54e325.
//
// Solidity: function subjectPresentationRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_AlastriaContracts *AlastriaContractsSession) SubjectPresentationRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _AlastriaContracts.Contract.SubjectPresentationRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectPresentationRegistry is a free data retrieval call binding the contract method 0x1e54e325.
//
// Solidity: function subjectPresentationRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_AlastriaContracts *AlastriaContractsCallerSession) SubjectPresentationRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _AlastriaContracts.Contract.SubjectPresentationRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
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

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddSubjectPresentation(opts *bind.TransactOpts, subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addSubjectPresentation", subjectPresentationHash, URI)
}

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddSubjectPresentation(subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddSubjectPresentation(&_AlastriaContracts.TransactOpts, subjectPresentationHash, URI)
}

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddSubjectPresentation(subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddSubjectPresentation(&_AlastriaContracts.TransactOpts, subjectPresentationHash, URI)
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

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) UpdateReceiverPresentation(opts *bind.TransactOpts, receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "updateReceiverPresentation", receiverPresentationHash, status)
}

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsSession) UpdateReceiverPresentation(receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateReceiverPresentation(&_AlastriaContracts.TransactOpts, receiverPresentationHash, status)
}

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) UpdateReceiverPresentation(receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateReceiverPresentation(&_AlastriaContracts.TransactOpts, receiverPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) UpdateSubjectPresentation(opts *bind.TransactOpts, subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "updateSubjectPresentation", subjectPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsSession) UpdateSubjectPresentation(subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateSubjectPresentation(&_AlastriaContracts.TransactOpts, subjectPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) UpdateSubjectPresentation(subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateSubjectPresentation(&_AlastriaContracts.TransactOpts, subjectPresentationHash, status)
}

// AlastriaContractsPresentationUpdatedIterator is returned from FilterPresentationUpdated and is used to iterate over the raw logs and unpacked data for PresentationUpdated events raised by the AlastriaContracts contract.
type AlastriaContractsPresentationUpdatedIterator struct {
	Event *AlastriaContractsPresentationUpdated // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsPresentationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsPresentationUpdated)
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
		it.Event = new(AlastriaContractsPresentationUpdated)
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
func (it *AlastriaContractsPresentationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsPresentationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsPresentationUpdated represents a PresentationUpdated event raised by the AlastriaContracts contract.
type AlastriaContractsPresentationUpdated struct {
	Hash   [32]byte
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPresentationUpdated is a free log retrieval operation binding the contract event 0x52035ff46a15e119b23cfbf756f83f68ee8740523a40b78f621fdd44f091fb38.
//
// Solidity: event PresentationUpdated(bytes32 hash, uint8 status)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterPresentationUpdated(opts *bind.FilterOpts) (*AlastriaContractsPresentationUpdatedIterator, error) {

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "PresentationUpdated")
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsPresentationUpdatedIterator{contract: _AlastriaContracts.contract, event: "PresentationUpdated", logs: logs, sub: sub}, nil
}

// WatchPresentationUpdated is a free log subscription operation binding the contract event 0x52035ff46a15e119b23cfbf756f83f68ee8740523a40b78f621fdd44f091fb38.
//
// Solidity: event PresentationUpdated(bytes32 hash, uint8 status)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchPresentationUpdated(opts *bind.WatchOpts, sink chan<- *AlastriaContractsPresentationUpdated) (event.Subscription, error) {

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "PresentationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsPresentationUpdated)
				if err := _AlastriaContracts.contract.UnpackLog(event, "PresentationUpdated", log); err != nil {
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

// ParsePresentationUpdated is a log parse operation binding the contract event 0x52035ff46a15e119b23cfbf756f83f68ee8740523a40b78f621fdd44f091fb38.
//
// Solidity: event PresentationUpdated(bytes32 hash, uint8 status)
func (_AlastriaContracts *AlastriaContractsFilterer) ParsePresentationUpdated(log types.Log) (*AlastriaContractsPresentationUpdated, error) {
	event := new(AlastriaContractsPresentationUpdated)
	if err := _AlastriaContracts.contract.UnpackLog(event, "PresentationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
