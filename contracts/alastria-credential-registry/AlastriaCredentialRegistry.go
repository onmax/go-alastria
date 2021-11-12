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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"IssuerCredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"SubjectCredentialDeleted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuerCredentialList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subjectCredentialList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subjectCredentialRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"name\":\"addSubjectCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"addIssuerCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"deleteSubjectCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"getSubjectCredentialStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getSubjectCredentialList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateCredentialStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"getIssuerCredentialStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"subjectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"issuerStatus\",\"type\":\"uint8\"}],\"name\":\"getCredentialStatus\",\"outputs\":[{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_Alastria *AlastriaCaller) GetCredentialStatus(opts *bind.CallOpts, subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getCredentialStatus", subjectStatus, issuerStatus)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_Alastria *AlastriaSession) GetCredentialStatus(subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	return _Alastria.Contract.GetCredentialStatus(&_Alastria.CallOpts, subjectStatus, issuerStatus)
}

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_Alastria *AlastriaCallerSession) GetCredentialStatus(subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	return _Alastria.Contract.GetCredentialStatus(&_Alastria.CallOpts, subjectStatus, issuerStatus)
}

// GetIssuerCredentialStatus is a free data retrieval call binding the contract method 0xaeda232a.
//
// Solidity: function getIssuerCredentialStatus(address issuer, bytes32 issuerCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCaller) GetIssuerCredentialStatus(opts *bind.CallOpts, issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getIssuerCredentialStatus", issuer, issuerCredentialHash)

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

// GetIssuerCredentialStatus is a free data retrieval call binding the contract method 0xaeda232a.
//
// Solidity: function getIssuerCredentialStatus(address issuer, bytes32 issuerCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaSession) GetIssuerCredentialStatus(issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetIssuerCredentialStatus(&_Alastria.CallOpts, issuer, issuerCredentialHash)
}

// GetIssuerCredentialStatus is a free data retrieval call binding the contract method 0xaeda232a.
//
// Solidity: function getIssuerCredentialStatus(address issuer, bytes32 issuerCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCallerSession) GetIssuerCredentialStatus(issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetIssuerCredentialStatus(&_Alastria.CallOpts, issuer, issuerCredentialHash)
}

// GetSubjectCredentialList is a free data retrieval call binding the contract method 0x4510b1cb.
//
// Solidity: function getSubjectCredentialList(address subject) view returns(uint256, bytes32[])
func (_Alastria *AlastriaCaller) GetSubjectCredentialList(opts *bind.CallOpts, subject common.Address) (*big.Int, [][32]byte, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getSubjectCredentialList", subject)

	if err != nil {
		return *new(*big.Int), *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new([][32]byte)).(*[][32]byte)

	return out0, out1, err

}

// GetSubjectCredentialList is a free data retrieval call binding the contract method 0x4510b1cb.
//
// Solidity: function getSubjectCredentialList(address subject) view returns(uint256, bytes32[])
func (_Alastria *AlastriaSession) GetSubjectCredentialList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _Alastria.Contract.GetSubjectCredentialList(&_Alastria.CallOpts, subject)
}

// GetSubjectCredentialList is a free data retrieval call binding the contract method 0x4510b1cb.
//
// Solidity: function getSubjectCredentialList(address subject) view returns(uint256, bytes32[])
func (_Alastria *AlastriaCallerSession) GetSubjectCredentialList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _Alastria.Contract.GetSubjectCredentialList(&_Alastria.CallOpts, subject)
}

// GetSubjectCredentialStatus is a free data retrieval call binding the contract method 0x55d64732.
//
// Solidity: function getSubjectCredentialStatus(address subject, bytes32 subjectCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCaller) GetSubjectCredentialStatus(opts *bind.CallOpts, subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getSubjectCredentialStatus", subject, subjectCredentialHash)

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

// GetSubjectCredentialStatus is a free data retrieval call binding the contract method 0x55d64732.
//
// Solidity: function getSubjectCredentialStatus(address subject, bytes32 subjectCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaSession) GetSubjectCredentialStatus(subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetSubjectCredentialStatus(&_Alastria.CallOpts, subject, subjectCredentialHash)
}

// GetSubjectCredentialStatus is a free data retrieval call binding the contract method 0x55d64732.
//
// Solidity: function getSubjectCredentialStatus(address subject, bytes32 subjectCredentialHash) view returns(bool exists, uint8 status)
func (_Alastria *AlastriaCallerSession) GetSubjectCredentialStatus(subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _Alastria.Contract.GetSubjectCredentialStatus(&_Alastria.CallOpts, subject, subjectCredentialHash)
}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCaller) IssuerCredentialList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "issuerCredentialList", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaSession) IssuerCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.IssuerCredentialList(&_Alastria.CallOpts, arg0, arg1)
}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCallerSession) IssuerCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.IssuerCredentialList(&_Alastria.CallOpts, arg0, arg1)
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

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCaller) SubjectCredentialList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "subjectCredentialList", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaSession) SubjectCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.SubjectCredentialList(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_Alastria *AlastriaCallerSession) SubjectCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _Alastria.Contract.SubjectCredentialList(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectCredentialRegistry is a free data retrieval call binding the contract method 0x6792223d.
//
// Solidity: function subjectCredentialRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_Alastria *AlastriaCaller) SubjectCredentialRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "subjectCredentialRegistry", arg0, arg1)

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

// SubjectCredentialRegistry is a free data retrieval call binding the contract method 0x6792223d.
//
// Solidity: function subjectCredentialRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_Alastria *AlastriaSession) SubjectCredentialRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _Alastria.Contract.SubjectCredentialRegistry(&_Alastria.CallOpts, arg0, arg1)
}

// SubjectCredentialRegistry is a free data retrieval call binding the contract method 0x6792223d.
//
// Solidity: function subjectCredentialRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_Alastria *AlastriaCallerSession) SubjectCredentialRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _Alastria.Contract.SubjectCredentialRegistry(&_Alastria.CallOpts, arg0, arg1)
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

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_Alastria *AlastriaTransactor) AddIssuerCredential(opts *bind.TransactOpts, issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addIssuerCredential", issuerCredentialHash)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_Alastria *AlastriaSession) AddIssuerCredential(issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.Contract.AddIssuerCredential(&_Alastria.TransactOpts, issuerCredentialHash)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_Alastria *AlastriaTransactorSession) AddIssuerCredential(issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.Contract.AddIssuerCredential(&_Alastria.TransactOpts, issuerCredentialHash)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_Alastria *AlastriaTransactor) AddSubjectCredential(opts *bind.TransactOpts, subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addSubjectCredential", subjectCredentialHash, URI)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_Alastria *AlastriaSession) AddSubjectCredential(subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.Contract.AddSubjectCredential(&_Alastria.TransactOpts, subjectCredentialHash, URI)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_Alastria *AlastriaTransactorSession) AddSubjectCredential(subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _Alastria.Contract.AddSubjectCredential(&_Alastria.TransactOpts, subjectCredentialHash, URI)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_Alastria *AlastriaTransactor) DeleteSubjectCredential(opts *bind.TransactOpts, subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "deleteSubjectCredential", subjectCredentialHash)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_Alastria *AlastriaSession) DeleteSubjectCredential(subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteSubjectCredential(&_Alastria.TransactOpts, subjectCredentialHash)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_Alastria *AlastriaTransactorSession) DeleteSubjectCredential(subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteSubjectCredential(&_Alastria.TransactOpts, subjectCredentialHash)
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

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_Alastria *AlastriaTransactor) UpdateCredentialStatus(opts *bind.TransactOpts, issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "updateCredentialStatus", issuerCredentialHash, status)
}

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_Alastria *AlastriaSession) UpdateCredentialStatus(issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateCredentialStatus(&_Alastria.TransactOpts, issuerCredentialHash, status)
}

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_Alastria *AlastriaTransactorSession) UpdateCredentialStatus(issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateCredentialStatus(&_Alastria.TransactOpts, issuerCredentialHash, status)
}

// AlastriaIssuerCredentialRevokedIterator is returned from FilterIssuerCredentialRevoked and is used to iterate over the raw logs and unpacked data for IssuerCredentialRevoked events raised by the Alastria contract.
type AlastriaIssuerCredentialRevokedIterator struct {
	Event *AlastriaIssuerCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *AlastriaIssuerCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaIssuerCredentialRevoked)
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
		it.Event = new(AlastriaIssuerCredentialRevoked)
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
func (it *AlastriaIssuerCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaIssuerCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaIssuerCredentialRevoked represents a IssuerCredentialRevoked event raised by the Alastria contract.
type AlastriaIssuerCredentialRevoked struct {
	IssuerCredentialHash [32]byte
	Status               uint8
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIssuerCredentialRevoked is a free log retrieval operation binding the contract event 0x0bc6eed115cb12ced57d22ce88a4be72a3b052cf01e5a1947a1f901f8ff26b1a.
//
// Solidity: event IssuerCredentialRevoked(bytes32 issuerCredentialHash, uint8 status)
func (_Alastria *AlastriaFilterer) FilterIssuerCredentialRevoked(opts *bind.FilterOpts) (*AlastriaIssuerCredentialRevokedIterator, error) {

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "IssuerCredentialRevoked")
	if err != nil {
		return nil, err
	}
	return &AlastriaIssuerCredentialRevokedIterator{contract: _Alastria.contract, event: "IssuerCredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchIssuerCredentialRevoked is a free log subscription operation binding the contract event 0x0bc6eed115cb12ced57d22ce88a4be72a3b052cf01e5a1947a1f901f8ff26b1a.
//
// Solidity: event IssuerCredentialRevoked(bytes32 issuerCredentialHash, uint8 status)
func (_Alastria *AlastriaFilterer) WatchIssuerCredentialRevoked(opts *bind.WatchOpts, sink chan<- *AlastriaIssuerCredentialRevoked) (event.Subscription, error) {

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "IssuerCredentialRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaIssuerCredentialRevoked)
				if err := _Alastria.contract.UnpackLog(event, "IssuerCredentialRevoked", log); err != nil {
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

// ParseIssuerCredentialRevoked is a log parse operation binding the contract event 0x0bc6eed115cb12ced57d22ce88a4be72a3b052cf01e5a1947a1f901f8ff26b1a.
//
// Solidity: event IssuerCredentialRevoked(bytes32 issuerCredentialHash, uint8 status)
func (_Alastria *AlastriaFilterer) ParseIssuerCredentialRevoked(log types.Log) (*AlastriaIssuerCredentialRevoked, error) {
	event := new(AlastriaIssuerCredentialRevoked)
	if err := _Alastria.contract.UnpackLog(event, "IssuerCredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaSubjectCredentialDeletedIterator is returned from FilterSubjectCredentialDeleted and is used to iterate over the raw logs and unpacked data for SubjectCredentialDeleted events raised by the Alastria contract.
type AlastriaSubjectCredentialDeletedIterator struct {
	Event *AlastriaSubjectCredentialDeleted // Event containing the contract specifics and raw log

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
func (it *AlastriaSubjectCredentialDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaSubjectCredentialDeleted)
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
		it.Event = new(AlastriaSubjectCredentialDeleted)
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
func (it *AlastriaSubjectCredentialDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaSubjectCredentialDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaSubjectCredentialDeleted represents a SubjectCredentialDeleted event raised by the Alastria contract.
type AlastriaSubjectCredentialDeleted struct {
	SubjectCredentialHash [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubjectCredentialDeleted is a free log retrieval operation binding the contract event 0xc83f14fb22d6787113e4b492b1b27daac6284e516a3c61d454359fd426d04c29.
//
// Solidity: event SubjectCredentialDeleted(bytes32 subjectCredentialHash)
func (_Alastria *AlastriaFilterer) FilterSubjectCredentialDeleted(opts *bind.FilterOpts) (*AlastriaSubjectCredentialDeletedIterator, error) {

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "SubjectCredentialDeleted")
	if err != nil {
		return nil, err
	}
	return &AlastriaSubjectCredentialDeletedIterator{contract: _Alastria.contract, event: "SubjectCredentialDeleted", logs: logs, sub: sub}, nil
}

// WatchSubjectCredentialDeleted is a free log subscription operation binding the contract event 0xc83f14fb22d6787113e4b492b1b27daac6284e516a3c61d454359fd426d04c29.
//
// Solidity: event SubjectCredentialDeleted(bytes32 subjectCredentialHash)
func (_Alastria *AlastriaFilterer) WatchSubjectCredentialDeleted(opts *bind.WatchOpts, sink chan<- *AlastriaSubjectCredentialDeleted) (event.Subscription, error) {

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "SubjectCredentialDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaSubjectCredentialDeleted)
				if err := _Alastria.contract.UnpackLog(event, "SubjectCredentialDeleted", log); err != nil {
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

// ParseSubjectCredentialDeleted is a log parse operation binding the contract event 0xc83f14fb22d6787113e4b492b1b27daac6284e516a3c61d454359fd426d04c29.
//
// Solidity: event SubjectCredentialDeleted(bytes32 subjectCredentialHash)
func (_Alastria *AlastriaFilterer) ParseSubjectCredentialDeleted(log types.Log) (*AlastriaSubjectCredentialDeleted, error) {
	event := new(AlastriaSubjectCredentialDeleted)
	if err := _Alastria.contract.UnpackLog(event, "SubjectCredentialDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
