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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"PresentationUpdated\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subjectPresentationListRegistry\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subjectPresentationRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"name\":\"addSubjectPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateSubjectPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subjectPresentationHash\",\"type\":\"bytes32\"}],\"name\":\"getSubjectPresentationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getSubjectPresentationList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"receiverPresentationHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateReceiverPresentation\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"receiverPresentationHash\",\"type\":\"bytes32\"}],\"name\":\"getReceiverPresentationStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"subjectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"receiverStatus\",\"type\":\"uint8\"}],\"name\":\"getPresentationStatus\",\"outputs\":[{\"internalType\":\"enumAlastriaPresentationRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_Alastria *AlastriaCaller) GetPresentationStatus(opts *bind.CallOpts, subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getPresentationStatus", subjectStatus, receiverStatus)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_Alastria *AlastriaSession) GetPresentationStatus(subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	return _Alastria.Contract.GetPresentationStatus(&_Alastria.CallOpts, subjectStatus, receiverStatus)
}

// GetPresentationStatus is a free data retrieval call binding the contract method 0x034f8408.
//
// Solidity: function getPresentationStatus(uint8 subjectStatus, uint8 receiverStatus) pure returns(uint8)
func (_Alastria *AlastriaCallerSession) GetPresentationStatus(subjectStatus uint8, receiverStatus uint8) (uint8, error) {
	return _Alastria.Contract.GetPresentationStatus(&_Alastria.CallOpts, subjectStatus, receiverStatus)
}

// GetReceiverPresentationStatus is a free data retrieval call binding the contract method 0x9373014a.
//
// Solidity: function getReceiverPresentationStatus(address receiver, bytes32 receiverPresentationHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCaller) GetReceiverPresentationStatus(opts *bind.CallOpts, receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getReceiverPresentationStatus", receiver, receiverPresentationHash)

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
func (_Alastria *AlastriaSession) GetReceiverPresentationStatus(receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetReceiverPresentationStatus(&_Alastria.CallOpts, receiver, receiverPresentationHash)
}

// GetReceiverPresentationStatus is a free data retrieval call binding the contract method 0x9373014a.
//
// Solidity: function getReceiverPresentationStatus(address receiver, bytes32 receiverPresentationHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCallerSession) GetReceiverPresentationStatus(receiver common.Address, receiverPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetReceiverPresentationStatus(&_Alastria.CallOpts, receiver, receiverPresentationHash)
}

// GetSubjectPresentationList is a free data retrieval call binding the contract method 0xc00dcd0e.
//
// Solidity: function getSubjectPresentationList(address subject) view returns(uint256, bytes32[])
func (_Alastria *AlastriaCaller) GetSubjectPresentationList(opts *bind.CallOpts, subject common.Address) (*big.Int, [][32]byte, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getSubjectPresentationList", subject)

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
func (_Alastria *AlastriaSession) GetSubjectPresentationList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _Alastria.Contract.GetSubjectPresentationList(&_Alastria.CallOpts, subject)
}

// GetSubjectPresentationList is a free data retrieval call binding the contract method 0xc00dcd0e.
//
// Solidity: function getSubjectPresentationList(address subject) view returns(uint256, bytes32[])
func (_Alastria *AlastriaCallerSession) GetSubjectPresentationList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _Alastria.Contract.GetSubjectPresentationList(&_Alastria.CallOpts, subject)
}

// GetSubjectPresentationStatus is a free data retrieval call binding the contract method 0x5205080f.
//
// Solidity: function getSubjectPresentationStatus(address subject, bytes32 subjectPresentationHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCaller) GetSubjectPresentationStatus(opts *bind.CallOpts, subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getSubjectPresentationStatus", subject, subjectPresentationHash)

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
func (_Alastria *AlastriaSession) GetSubjectPresentationStatus(subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetSubjectPresentationStatus(&_Alastria.CallOpts, subject, subjectPresentationHash)
}

// GetSubjectPresentationStatus is a free data retrieval call binding the contract method 0x5205080f.
//
// Solidity: function getSubjectPresentationStatus(address subject, bytes32 subjectPresentationHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCallerSession) GetSubjectPresentationStatus(subject common.Address, subjectPresentationHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetSubjectPresentationStatus(&_Alastria.CallOpts, subject, subjectPresentationHash)
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

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCaller) SubjectPresentationListRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "subjectPresentationListRegistry", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaSession) SubjectPresentationListRegistry(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.SubjectPresentationListRegistry(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectPresentationListRegistry is a free data retrieval call binding the contract method 0xe4734593.
//
// Solidity: function subjectPresentationListRegistry(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCallerSession) SubjectPresentationListRegistry(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.SubjectPresentationListRegistry(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectPresentationRegistry is a free data retrieval call binding the contract method 0x1e54e325.
//
// Solidity: function subjectPresentationRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_Alastria *AlastriaCaller) SubjectPresentationRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "subjectPresentationRegistry", arg0, arg1)

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
func (_Alastria *AlastriaSession) SubjectPresentationRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _Alastria.Contract.SubjectPresentationRegistry(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectPresentationRegistry is a free data retrieval call binding the contract method 0x1e54e325.
//
// Solidity: function subjectPresentationRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_Alastria *AlastriaCallerSession) SubjectPresentationRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _Alastria.Contract.SubjectPresentationRegistry(&_Alastria.CallOpts, arg0, arg1)
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

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_Alastria *AlastriaTransactor) AddSubjectPresentation(opts *bind.TransactOpts, subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addSubjectPresentation", subjectPresentationHash, URI)
}

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_Alastria *AlastriaSession) AddSubjectPresentation(subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.Contract.AddSubjectPresentation(&_Alastria.TransactOpts, subjectPresentationHash, URI)
}

// AddSubjectPresentation is a paid mutator transaction binding the contract method 0x4e3a5de5.
//
// Solidity: function addSubjectPresentation(bytes32 subjectPresentationHash, string URI) returns()
func (_Alastria *AlastriaTransactorSession) AddSubjectPresentation(subjectPresentationHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.Contract.AddSubjectPresentation(&_Alastria.TransactOpts, subjectPresentationHash, URI)
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

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaTransactor) UpdateReceiverPresentation(opts *bind.TransactOpts, receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "updateReceiverPresentation", receiverPresentationHash, status)
}

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaSession) UpdateReceiverPresentation(receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateReceiverPresentation(&_Alastria.TransactOpts, receiverPresentationHash, status)
}

// UpdateReceiverPresentation is a paid mutator transaction binding the contract method 0x3000dc39.
//
// Solidity: function updateReceiverPresentation(bytes32 receiverPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaTransactorSession) UpdateReceiverPresentation(receiverPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateReceiverPresentation(&_Alastria.TransactOpts, receiverPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaTransactor) UpdateSubjectPresentation(opts *bind.TransactOpts, subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "updateSubjectPresentation", subjectPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaSession) UpdateSubjectPresentation(subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateSubjectPresentation(&_Alastria.TransactOpts, subjectPresentationHash, status)
}

// UpdateSubjectPresentation is a paid mutator transaction binding the contract method 0xe64af938.
//
// Solidity: function updateSubjectPresentation(bytes32 subjectPresentationHash, uint8 status) returns()
func (_Alastria *AlastriaTransactorSession) UpdateSubjectPresentation(subjectPresentationHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateSubjectPresentation(&_Alastria.TransactOpts, subjectPresentationHash, status)
}

// AlastriaPresentationUpdatedIterator is returned from FilterPresentationUpdated and is used to iterate over the raw logs and unpacked data for PresentationUpdated events raised by the Alastria contract.
type AlastriaPresentationUpdatedIterator struct {
	Event *AlastriaPresentationUpdated // Event containing the contract specifics and raw log

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
func (it *AlastriaPresentationUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaPresentationUpdated)
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
		it.Event = new(AlastriaPresentationUpdated)
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
func (it *AlastriaPresentationUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaPresentationUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaPresentationUpdated represents a PresentationUpdated event raised by the Alastria contract.
type AlastriaPresentationUpdated struct {
	Hash   [32]byte
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPresentationUpdated is a free log retrieval operation binding the contract event 0x52035ff46a15e119b23cfbf756f83f68ee8740523a40b78f621fdd44f091fb38.
//
// Solidity: event PresentationUpdated(bytes32 hash, uint8 status)
func (_Alastria *AlastriaFilterer) FilterPresentationUpdated(opts *bind.FilterOpts) (*AlastriaPresentationUpdatedIterator, error) {

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "PresentationUpdated")
	if err != nil {
		return nil, err
	}
	return &AlastriaPresentationUpdatedIterator{contract: _Alastria.contract, event: "PresentationUpdated", logs: logs, sub: sub}, nil
}

// WatchPresentationUpdated is a free log subscription operation binding the contract event 0x52035ff46a15e119b23cfbf756f83f68ee8740523a40b78f621fdd44f091fb38.
//
// Solidity: event PresentationUpdated(bytes32 hash, uint8 status)
func (_Alastria *AlastriaFilterer) WatchPresentationUpdated(opts *bind.WatchOpts, sink chan<- *AlastriaPresentationUpdated) (event.Subscription, error) {

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "PresentationUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaPresentationUpdated)
				if err := _Alastria.contract.UnpackLog(event, "PresentationUpdated", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParsePresentationUpdated(log types.Log) (*AlastriaPresentationUpdated, error) {
	event := new(AlastriaPresentationUpdated)
	if err := _Alastria.contract.UnpackLog(event, "PresentationUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
