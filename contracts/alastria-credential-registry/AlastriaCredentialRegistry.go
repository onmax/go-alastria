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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"IssuerCredentialRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"SubjectCredentialDeleted\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"issuerCredentialList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"previousPublishedVersion\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"subjectCredentialList\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"subjectCredentialRegistry\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_previousPublishedVersion\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"string\",\"name\":\"URI\",\"type\":\"string\"}],\"name\":\"addSubjectCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"addIssuerCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"deleteSubjectCredential\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"subjectCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"getSubjectCredentialStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"}],\"name\":\"getSubjectCredentialList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"name\":\"updateCredentialStatus\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"issuerCredentialHash\",\"type\":\"bytes32\"}],\"name\":\"getIssuerCredentialStatus\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"exists\",\"type\":\"bool\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"status\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"subjectStatus\",\"type\":\"uint8\"},{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"issuerStatus\",\"type\":\"uint8\"}],\"name\":\"getCredentialStatus\",\"outputs\":[{\"internalType\":\"enumAlastriaCredentialRegistry.Status\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]",
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

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsCaller) GetCredentialStatus(opts *bind.CallOpts, subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getCredentialStatus", subjectStatus, issuerStatus)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsSession) GetCredentialStatus(subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	return _AlastriaContracts.Contract.GetCredentialStatus(&_AlastriaContracts.CallOpts, subjectStatus, issuerStatus)
}

// GetCredentialStatus is a free data retrieval call binding the contract method 0x5faf7d94.
//
// Solidity: function getCredentialStatus(uint8 subjectStatus, uint8 issuerStatus) pure returns(uint8)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetCredentialStatus(subjectStatus uint8, issuerStatus uint8) (uint8, error) {
	return _AlastriaContracts.Contract.GetCredentialStatus(&_AlastriaContracts.CallOpts, subjectStatus, issuerStatus)
}

// GetIssuerCredentialStatus is a free data retrieval call binding the contract method 0xaeda232a.
//
// Solidity: function getIssuerCredentialStatus(address issuer, bytes32 issuerCredentialHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCaller) GetIssuerCredentialStatus(opts *bind.CallOpts, issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getIssuerCredentialStatus", issuer, issuerCredentialHash)

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
func (_AlastriaContracts *AlastriaContractsSession) GetIssuerCredentialStatus(issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetIssuerCredentialStatus(&_AlastriaContracts.CallOpts, issuer, issuerCredentialHash)
}

// GetIssuerCredentialStatus is a free data retrieval call binding the contract method 0xaeda232a.
//
// Solidity: function getIssuerCredentialStatus(address issuer, bytes32 issuerCredentialHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetIssuerCredentialStatus(issuer common.Address, issuerCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetIssuerCredentialStatus(&_AlastriaContracts.CallOpts, issuer, issuerCredentialHash)
}

// GetSubjectCredentialList is a free data retrieval call binding the contract method 0x4510b1cb.
//
// Solidity: function getSubjectCredentialList(address subject) view returns(uint256, bytes32[])
func (_AlastriaContracts *AlastriaContractsCaller) GetSubjectCredentialList(opts *bind.CallOpts, subject common.Address) (*big.Int, [][32]byte, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getSubjectCredentialList", subject)

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
func (_AlastriaContracts *AlastriaContractsSession) GetSubjectCredentialList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _AlastriaContracts.Contract.GetSubjectCredentialList(&_AlastriaContracts.CallOpts, subject)
}

// GetSubjectCredentialList is a free data retrieval call binding the contract method 0x4510b1cb.
//
// Solidity: function getSubjectCredentialList(address subject) view returns(uint256, bytes32[])
func (_AlastriaContracts *AlastriaContractsCallerSession) GetSubjectCredentialList(subject common.Address) (*big.Int, [][32]byte, error) {
	return _AlastriaContracts.Contract.GetSubjectCredentialList(&_AlastriaContracts.CallOpts, subject)
}

// GetSubjectCredentialStatus is a free data retrieval call binding the contract method 0x55d64732.
//
// Solidity: function getSubjectCredentialStatus(address subject, bytes32 subjectCredentialHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCaller) GetSubjectCredentialStatus(opts *bind.CallOpts, subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getSubjectCredentialStatus", subject, subjectCredentialHash)

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
func (_AlastriaContracts *AlastriaContractsSession) GetSubjectCredentialStatus(subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetSubjectCredentialStatus(&_AlastriaContracts.CallOpts, subject, subjectCredentialHash)
}

// GetSubjectCredentialStatus is a free data retrieval call binding the contract method 0x55d64732.
//
// Solidity: function getSubjectCredentialStatus(address subject, bytes32 subjectCredentialHash) view returns(bool exists, uint8 status)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetSubjectCredentialStatus(subject common.Address, subjectCredentialHash [32]byte) (struct {
	Exists bool
	Status uint8
}, error) {
	return _AlastriaContracts.Contract.GetSubjectCredentialStatus(&_AlastriaContracts.CallOpts, subject, subjectCredentialHash)
}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCaller) IssuerCredentialList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "issuerCredentialList", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsSession) IssuerCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.IssuerCredentialList(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// IssuerCredentialList is a free data retrieval call binding the contract method 0x41e61ff8.
//
// Solidity: function issuerCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCallerSession) IssuerCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.IssuerCredentialList(&_AlastriaContracts.CallOpts, arg0, arg1)
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

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCaller) SubjectCredentialList(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "subjectCredentialList", arg0, arg1)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsSession) SubjectCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.SubjectCredentialList(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectCredentialList is a free data retrieval call binding the contract method 0xc43f6b36.
//
// Solidity: function subjectCredentialList(address , uint256 ) view returns(bytes32)
func (_AlastriaContracts *AlastriaContractsCallerSession) SubjectCredentialList(arg0 common.Address, arg1 *big.Int) ([32]byte, error) {
	return _AlastriaContracts.Contract.SubjectCredentialList(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectCredentialRegistry is a free data retrieval call binding the contract method 0x6792223d.
//
// Solidity: function subjectCredentialRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_AlastriaContracts *AlastriaContractsCaller) SubjectCredentialRegistry(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "subjectCredentialRegistry", arg0, arg1)

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
func (_AlastriaContracts *AlastriaContractsSession) SubjectCredentialRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _AlastriaContracts.Contract.SubjectCredentialRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
}

// SubjectCredentialRegistry is a free data retrieval call binding the contract method 0x6792223d.
//
// Solidity: function subjectCredentialRegistry(address , bytes32 ) view returns(bool exists, uint8 status, string URI)
func (_AlastriaContracts *AlastriaContractsCallerSession) SubjectCredentialRegistry(arg0 common.Address, arg1 [32]byte) (struct {
	Exists bool
	Status uint8
	URI    string
}, error) {
	return _AlastriaContracts.Contract.SubjectCredentialRegistry(&_AlastriaContracts.CallOpts, arg0, arg1)
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

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddIssuerCredential(opts *bind.TransactOpts, issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addIssuerCredential", issuerCredentialHash)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddIssuerCredential(issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIssuerCredential(&_AlastriaContracts.TransactOpts, issuerCredentialHash)
}

// AddIssuerCredential is a paid mutator transaction binding the contract method 0xdc11e39b.
//
// Solidity: function addIssuerCredential(bytes32 issuerCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddIssuerCredential(issuerCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIssuerCredential(&_AlastriaContracts.TransactOpts, issuerCredentialHash)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddSubjectCredential(opts *bind.TransactOpts, subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addSubjectCredential", subjectCredentialHash, URI)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddSubjectCredential(subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddSubjectCredential(&_AlastriaContracts.TransactOpts, subjectCredentialHash, URI)
}

// AddSubjectCredential is a paid mutator transaction binding the contract method 0xe04ce02c.
//
// Solidity: function addSubjectCredential(bytes32 subjectCredentialHash, string URI) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddSubjectCredential(subjectCredentialHash [32]byte, URI string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddSubjectCredential(&_AlastriaContracts.TransactOpts, subjectCredentialHash, URI)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeleteSubjectCredential(opts *bind.TransactOpts, subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deleteSubjectCredential", subjectCredentialHash)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeleteSubjectCredential(subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteSubjectCredential(&_AlastriaContracts.TransactOpts, subjectCredentialHash)
}

// DeleteSubjectCredential is a paid mutator transaction binding the contract method 0xdff9a1a9.
//
// Solidity: function deleteSubjectCredential(bytes32 subjectCredentialHash) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeleteSubjectCredential(subjectCredentialHash [32]byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteSubjectCredential(&_AlastriaContracts.TransactOpts, subjectCredentialHash)
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

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) UpdateCredentialStatus(opts *bind.TransactOpts, issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "updateCredentialStatus", issuerCredentialHash, status)
}

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsSession) UpdateCredentialStatus(issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateCredentialStatus(&_AlastriaContracts.TransactOpts, issuerCredentialHash, status)
}

// UpdateCredentialStatus is a paid mutator transaction binding the contract method 0xdd517e10.
//
// Solidity: function updateCredentialStatus(bytes32 issuerCredentialHash, uint8 status) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) UpdateCredentialStatus(issuerCredentialHash [32]byte, status uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateCredentialStatus(&_AlastriaContracts.TransactOpts, issuerCredentialHash, status)
}

// AlastriaContractsIssuerCredentialRevokedIterator is returned from FilterIssuerCredentialRevoked and is used to iterate over the raw logs and unpacked data for IssuerCredentialRevoked events raised by the AlastriaContracts contract.
type AlastriaContractsIssuerCredentialRevokedIterator struct {
	Event *AlastriaContractsIssuerCredentialRevoked // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsIssuerCredentialRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsIssuerCredentialRevoked)
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
		it.Event = new(AlastriaContractsIssuerCredentialRevoked)
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
func (it *AlastriaContractsIssuerCredentialRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsIssuerCredentialRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsIssuerCredentialRevoked represents a IssuerCredentialRevoked event raised by the AlastriaContracts contract.
type AlastriaContractsIssuerCredentialRevoked struct {
	IssuerCredentialHash [32]byte
	Status               uint8
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterIssuerCredentialRevoked is a free log retrieval operation binding the contract event 0x0bc6eed115cb12ced57d22ce88a4be72a3b052cf01e5a1947a1f901f8ff26b1a.
//
// Solidity: event IssuerCredentialRevoked(bytes32 issuerCredentialHash, uint8 status)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterIssuerCredentialRevoked(opts *bind.FilterOpts) (*AlastriaContractsIssuerCredentialRevokedIterator, error) {

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "IssuerCredentialRevoked")
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsIssuerCredentialRevokedIterator{contract: _AlastriaContracts.contract, event: "IssuerCredentialRevoked", logs: logs, sub: sub}, nil
}

// WatchIssuerCredentialRevoked is a free log subscription operation binding the contract event 0x0bc6eed115cb12ced57d22ce88a4be72a3b052cf01e5a1947a1f901f8ff26b1a.
//
// Solidity: event IssuerCredentialRevoked(bytes32 issuerCredentialHash, uint8 status)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchIssuerCredentialRevoked(opts *bind.WatchOpts, sink chan<- *AlastriaContractsIssuerCredentialRevoked) (event.Subscription, error) {

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "IssuerCredentialRevoked")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsIssuerCredentialRevoked)
				if err := _AlastriaContracts.contract.UnpackLog(event, "IssuerCredentialRevoked", log); err != nil {
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
func (_AlastriaContracts *AlastriaContractsFilterer) ParseIssuerCredentialRevoked(log types.Log) (*AlastriaContractsIssuerCredentialRevoked, error) {
	event := new(AlastriaContractsIssuerCredentialRevoked)
	if err := _AlastriaContracts.contract.UnpackLog(event, "IssuerCredentialRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaContractsSubjectCredentialDeletedIterator is returned from FilterSubjectCredentialDeleted and is used to iterate over the raw logs and unpacked data for SubjectCredentialDeleted events raised by the AlastriaContracts contract.
type AlastriaContractsSubjectCredentialDeletedIterator struct {
	Event *AlastriaContractsSubjectCredentialDeleted // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsSubjectCredentialDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsSubjectCredentialDeleted)
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
		it.Event = new(AlastriaContractsSubjectCredentialDeleted)
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
func (it *AlastriaContractsSubjectCredentialDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsSubjectCredentialDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsSubjectCredentialDeleted represents a SubjectCredentialDeleted event raised by the AlastriaContracts contract.
type AlastriaContractsSubjectCredentialDeleted struct {
	SubjectCredentialHash [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSubjectCredentialDeleted is a free log retrieval operation binding the contract event 0xc83f14fb22d6787113e4b492b1b27daac6284e516a3c61d454359fd426d04c29.
//
// Solidity: event SubjectCredentialDeleted(bytes32 subjectCredentialHash)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterSubjectCredentialDeleted(opts *bind.FilterOpts) (*AlastriaContractsSubjectCredentialDeletedIterator, error) {

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "SubjectCredentialDeleted")
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsSubjectCredentialDeletedIterator{contract: _AlastriaContracts.contract, event: "SubjectCredentialDeleted", logs: logs, sub: sub}, nil
}

// WatchSubjectCredentialDeleted is a free log subscription operation binding the contract event 0xc83f14fb22d6787113e4b492b1b27daac6284e516a3c61d454359fd426d04c29.
//
// Solidity: event SubjectCredentialDeleted(bytes32 subjectCredentialHash)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchSubjectCredentialDeleted(opts *bind.WatchOpts, sink chan<- *AlastriaContractsSubjectCredentialDeleted) (event.Subscription, error) {

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "SubjectCredentialDeleted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsSubjectCredentialDeleted)
				if err := _AlastriaContracts.contract.UnpackLog(event, "SubjectCredentialDeleted", log); err != nil {
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
func (_AlastriaContracts *AlastriaContractsFilterer) ParseSubjectCredentialDeleted(log types.Log) (*AlastriaContractsSubjectCredentialDeleted, error) {
	event := new(AlastriaContractsSubjectCredentialDeleted)
	if err := _AlastriaContracts.contract.UnpackLog(event, "SubjectCredentialDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
