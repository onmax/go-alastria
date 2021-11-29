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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"serviceProvider\",\"type\":\"address\"}],\"name\":\"IdentityRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"OperationWasNotSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"}],\"name\":\"PreparedAlastriaID\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"addIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"addIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaCredentialRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaCredentialRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaPresentationRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaPresentationRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaPublicKeyRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaPublicKeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"deleteIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"deleteIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstIdentityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"getEidasLevel\",\"outputs\":[{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"identityKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"isIdentityIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"isIdentityServiceProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"updateIdentityIssuerEidasLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_credentialRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_publicKeyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_presentationRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_firstIdentityWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signAddress\",\"type\":\"address\"}],\"name\":\"prepareAlastriaID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addPublicKeyCallData\",\"type\":\"bytes\"}],\"name\":\"createAlastriaIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"delegateCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountLost\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAccount\",\"type\":\"address\"}],\"name\":\"recoverAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) AlastriaCredentialRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "alastriaCredentialRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) AlastriaCredentialRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaCredentialRegistry(&_AlastriaContracts.CallOpts)
}

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) AlastriaCredentialRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaCredentialRegistry(&_AlastriaContracts.CallOpts)
}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) AlastriaPresentationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "alastriaPresentationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) AlastriaPresentationRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaPresentationRegistry(&_AlastriaContracts.CallOpts)
}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) AlastriaPresentationRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaPresentationRegistry(&_AlastriaContracts.CallOpts)
}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) AlastriaPublicKeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "alastriaPublicKeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) AlastriaPublicKeyRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaPublicKeyRegistry(&_AlastriaContracts.CallOpts)
}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) AlastriaPublicKeyRegistry() (common.Address, error) {
	return _AlastriaContracts.Contract.AlastriaPublicKeyRegistry(&_AlastriaContracts.CallOpts)
}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) FirstIdentityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "firstIdentityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) FirstIdentityWallet() (common.Address, error) {
	return _AlastriaContracts.Contract.FirstIdentityWallet(&_AlastriaContracts.CallOpts)
}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) FirstIdentityWallet() (common.Address, error) {
	return _AlastriaContracts.Contract.FirstIdentityWallet(&_AlastriaContracts.CallOpts)
}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsCaller) GetEidasLevel(opts *bind.CallOpts, _identityIssuer common.Address) (uint8, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getEidasLevel", _identityIssuer)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _AlastriaContracts.Contract.GetEidasLevel(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _AlastriaContracts.Contract.GetEidasLevel(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) IdentityKeys(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "identityKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) IdentityKeys(arg0 common.Address) (common.Address, error) {
	return _AlastriaContracts.Contract.IdentityKeys(&_AlastriaContracts.CallOpts, arg0)
}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) IdentityKeys(arg0 common.Address) (common.Address, error) {
	return _AlastriaContracts.Contract.IdentityKeys(&_AlastriaContracts.CallOpts, arg0)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCaller) IsIdentityIssuer(opts *bind.CallOpts, _identityIssuer common.Address) (bool, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "isIdentityIssuer", _identityIssuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityIssuer(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCallerSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityIssuer(&_AlastriaContracts.CallOpts, _identityIssuer)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCaller) IsIdentityServiceProvider(opts *bind.CallOpts, _identityServiceProvider common.Address) (bool, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "isIdentityServiceProvider", _identityServiceProvider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityServiceProvider(&_AlastriaContracts.CallOpts, _identityServiceProvider)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCallerSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsIdentityServiceProvider(&_AlastriaContracts.CallOpts, _identityServiceProvider)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCaller) IsOwner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "isOwner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_AlastriaContracts *AlastriaContractsSession) IsOwner(addr common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsOwner(&_AlastriaContracts.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_AlastriaContracts *AlastriaContractsCallerSession) IsOwner(addr common.Address) (bool, error) {
	return _AlastriaContracts.Contract.IsOwner(&_AlastriaContracts.CallOpts, addr)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlastriaContracts *AlastriaContractsCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlastriaContracts *AlastriaContractsSession) Owner() (common.Address, error) {
	return _AlastriaContracts.Contract.Owner(&_AlastriaContracts.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AlastriaContracts *AlastriaContractsCallerSession) Owner() (common.Address, error) {
	return _AlastriaContracts.Contract.Owner(&_AlastriaContracts.CallOpts)
}

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_AlastriaContracts *AlastriaContractsCaller) PendingIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "pendingIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_AlastriaContracts *AlastriaContractsSession) PendingIDs(arg0 common.Address) (*big.Int, error) {
	return _AlastriaContracts.Contract.PendingIDs(&_AlastriaContracts.CallOpts, arg0)
}

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_AlastriaContracts *AlastriaContractsCallerSession) PendingIDs(arg0 common.Address) (*big.Int, error) {
	return _AlastriaContracts.Contract.PendingIDs(&_AlastriaContracts.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
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
// Solidity: function version() view returns(uint256)
func (_AlastriaContracts *AlastriaContractsSession) Version() (*big.Int, error) {
	return _AlastriaContracts.Contract.Version(&_AlastriaContracts.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_AlastriaContracts *AlastriaContractsCallerSession) Version() (*big.Int, error) {
	return _AlastriaContracts.Contract.Version(&_AlastriaContracts.CallOpts)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addIdentityIssuer", _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addIdentityServiceProvider", _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) CreateAlastriaIdentity(opts *bind.TransactOpts, addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "createAlastriaIdentity", addPublicKeyCallData)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_AlastriaContracts *AlastriaContractsSession) CreateAlastriaIdentity(addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.CreateAlastriaIdentity(&_AlastriaContracts.TransactOpts, addPublicKeyCallData)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) CreateAlastriaIdentity(addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.CreateAlastriaIdentity(&_AlastriaContracts.TransactOpts, addPublicKeyCallData)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_AlastriaContracts *AlastriaContractsTransactor) DelegateCall(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "delegateCall", _destination, _value, _data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_AlastriaContracts *AlastriaContractsSession) DelegateCall(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DelegateCall(&_AlastriaContracts.TransactOpts, _destination, _value, _data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_AlastriaContracts *AlastriaContractsTransactorSession) DelegateCall(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DelegateCall(&_AlastriaContracts.TransactOpts, _destination, _value, _data)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeleteIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deleteIdentityIssuer", _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityIssuer(&_AlastriaContracts.TransactOpts, _identityIssuer)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) DeleteIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "deleteIdentityServiceProvider", _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.DeleteIdentityServiceProvider(&_AlastriaContracts.TransactOpts, _identityServiceProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) Initialize(opts *bind.TransactOpts, _credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "initialize", _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_AlastriaContracts *AlastriaContractsSession) Initialize(_credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Initialize(&_AlastriaContracts.TransactOpts, _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) Initialize(_credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Initialize(&_AlastriaContracts.TransactOpts, _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) PrepareAlastriaID(opts *bind.TransactOpts, _signAddress common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "prepareAlastriaID", _signAddress)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_AlastriaContracts *AlastriaContractsSession) PrepareAlastriaID(_signAddress common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.PrepareAlastriaID(&_AlastriaContracts.TransactOpts, _signAddress)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) PrepareAlastriaID(_signAddress common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.PrepareAlastriaID(&_AlastriaContracts.TransactOpts, _signAddress)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) RecoverAccount(opts *bind.TransactOpts, accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "recoverAccount", accountLost, newAccount)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_AlastriaContracts *AlastriaContractsSession) RecoverAccount(accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.RecoverAccount(&_AlastriaContracts.TransactOpts, accountLost, newAccount)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) RecoverAccount(accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.RecoverAccount(&_AlastriaContracts.TransactOpts, accountLost, newAccount)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) Transfer(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "transfer", newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_AlastriaContracts *AlastriaContractsSession) Transfer(newOwner common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Transfer(&_AlastriaContracts.TransactOpts, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) Transfer(newOwner common.Address) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.Transfer(&_AlastriaContracts.TransactOpts, newOwner)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) UpdateIdentityIssuerEidasLevel(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "updateIdentityIssuerEidasLevel", _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateIdentityIssuerEidasLevel(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.UpdateIdentityIssuerEidasLevel(&_AlastriaContracts.TransactOpts, _identityIssuer, _level)
}

// AlastriaContractsIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the AlastriaContracts contract.
type AlastriaContractsIdentityCreatedIterator struct {
	Event *AlastriaContractsIdentityCreated // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsIdentityCreated)
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
		it.Event = new(AlastriaContractsIdentityCreated)
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
func (it *AlastriaContractsIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsIdentityCreated represents a IdentityCreated event raised by the AlastriaContracts contract.
type AlastriaContractsIdentityCreated struct {
	Identity common.Address
	Creator  common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0x8f2e597fd6e795e9851eea530f987cceb641315e3c7c9d484b798e3159d42095.
//
// Solidity: event IdentityCreated(address indexed identity, address indexed creator, address owner)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterIdentityCreated(opts *bind.FilterOpts, identity []common.Address, creator []common.Address) (*AlastriaContractsIdentityCreatedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "IdentityCreated", identityRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsIdentityCreatedIterator{contract: _AlastriaContracts.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0x8f2e597fd6e795e9851eea530f987cceb641315e3c7c9d484b798e3159d42095.
//
// Solidity: event IdentityCreated(address indexed identity, address indexed creator, address owner)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *AlastriaContractsIdentityCreated, identity []common.Address, creator []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "IdentityCreated", identityRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsIdentityCreated)
				if err := _AlastriaContracts.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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

// ParseIdentityCreated is a log parse operation binding the contract event 0x8f2e597fd6e795e9851eea530f987cceb641315e3c7c9d484b798e3159d42095.
//
// Solidity: event IdentityCreated(address indexed identity, address indexed creator, address owner)
func (_AlastriaContracts *AlastriaContractsFilterer) ParseIdentityCreated(log types.Log) (*AlastriaContractsIdentityCreated, error) {
	event := new(AlastriaContractsIdentityCreated)
	if err := _AlastriaContracts.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaContractsIdentityRecoveredIterator is returned from FilterIdentityRecovered and is used to iterate over the raw logs and unpacked data for IdentityRecovered events raised by the AlastriaContracts contract.
type AlastriaContractsIdentityRecoveredIterator struct {
	Event *AlastriaContractsIdentityRecovered // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsIdentityRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsIdentityRecovered)
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
		it.Event = new(AlastriaContractsIdentityRecovered)
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
func (it *AlastriaContractsIdentityRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsIdentityRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsIdentityRecovered represents a IdentityRecovered event raised by the AlastriaContracts contract.
type AlastriaContractsIdentityRecovered struct {
	OldAccount      common.Address
	NewAccount      common.Address
	ServiceProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityRecovered is a free log retrieval operation binding the contract event 0xf5cf2e915faa705d003410f0f65cd59606d22f8e7fb7361a015d6a00b7723d99.
//
// Solidity: event IdentityRecovered(address indexed oldAccount, address newAccount, address indexed serviceProvider)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterIdentityRecovered(opts *bind.FilterOpts, oldAccount []common.Address, serviceProvider []common.Address) (*AlastriaContractsIdentityRecoveredIterator, error) {

	var oldAccountRule []interface{}
	for _, oldAccountItem := range oldAccount {
		oldAccountRule = append(oldAccountRule, oldAccountItem)
	}

	var serviceProviderRule []interface{}
	for _, serviceProviderItem := range serviceProvider {
		serviceProviderRule = append(serviceProviderRule, serviceProviderItem)
	}

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "IdentityRecovered", oldAccountRule, serviceProviderRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsIdentityRecoveredIterator{contract: _AlastriaContracts.contract, event: "IdentityRecovered", logs: logs, sub: sub}, nil
}

// WatchIdentityRecovered is a free log subscription operation binding the contract event 0xf5cf2e915faa705d003410f0f65cd59606d22f8e7fb7361a015d6a00b7723d99.
//
// Solidity: event IdentityRecovered(address indexed oldAccount, address newAccount, address indexed serviceProvider)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchIdentityRecovered(opts *bind.WatchOpts, sink chan<- *AlastriaContractsIdentityRecovered, oldAccount []common.Address, serviceProvider []common.Address) (event.Subscription, error) {

	var oldAccountRule []interface{}
	for _, oldAccountItem := range oldAccount {
		oldAccountRule = append(oldAccountRule, oldAccountItem)
	}

	var serviceProviderRule []interface{}
	for _, serviceProviderItem := range serviceProvider {
		serviceProviderRule = append(serviceProviderRule, serviceProviderItem)
	}

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "IdentityRecovered", oldAccountRule, serviceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsIdentityRecovered)
				if err := _AlastriaContracts.contract.UnpackLog(event, "IdentityRecovered", log); err != nil {
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

// ParseIdentityRecovered is a log parse operation binding the contract event 0xf5cf2e915faa705d003410f0f65cd59606d22f8e7fb7361a015d6a00b7723d99.
//
// Solidity: event IdentityRecovered(address indexed oldAccount, address newAccount, address indexed serviceProvider)
func (_AlastriaContracts *AlastriaContractsFilterer) ParseIdentityRecovered(log types.Log) (*AlastriaContractsIdentityRecovered, error) {
	event := new(AlastriaContractsIdentityRecovered)
	if err := _AlastriaContracts.contract.UnpackLog(event, "IdentityRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaContractsOperationWasNotSupportedIterator is returned from FilterOperationWasNotSupported and is used to iterate over the raw logs and unpacked data for OperationWasNotSupported events raised by the AlastriaContracts contract.
type AlastriaContractsOperationWasNotSupportedIterator struct {
	Event *AlastriaContractsOperationWasNotSupported // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsOperationWasNotSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsOperationWasNotSupported)
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
		it.Event = new(AlastriaContractsOperationWasNotSupported)
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
func (it *AlastriaContractsOperationWasNotSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsOperationWasNotSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsOperationWasNotSupported represents a OperationWasNotSupported event raised by the AlastriaContracts contract.
type AlastriaContractsOperationWasNotSupported struct {
	Method common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOperationWasNotSupported is a free log retrieval operation binding the contract event 0x0f759709e9a850d07441d2d9dbef884c5aae908ad08fe5e229ad74bb9bcef0fa.
//
// Solidity: event OperationWasNotSupported(string indexed method)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterOperationWasNotSupported(opts *bind.FilterOpts, method []string) (*AlastriaContractsOperationWasNotSupportedIterator, error) {

	var methodRule []interface{}
	for _, methodItem := range method {
		methodRule = append(methodRule, methodItem)
	}

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "OperationWasNotSupported", methodRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsOperationWasNotSupportedIterator{contract: _AlastriaContracts.contract, event: "OperationWasNotSupported", logs: logs, sub: sub}, nil
}

// WatchOperationWasNotSupported is a free log subscription operation binding the contract event 0x0f759709e9a850d07441d2d9dbef884c5aae908ad08fe5e229ad74bb9bcef0fa.
//
// Solidity: event OperationWasNotSupported(string indexed method)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchOperationWasNotSupported(opts *bind.WatchOpts, sink chan<- *AlastriaContractsOperationWasNotSupported, method []string) (event.Subscription, error) {

	var methodRule []interface{}
	for _, methodItem := range method {
		methodRule = append(methodRule, methodItem)
	}

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "OperationWasNotSupported", methodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsOperationWasNotSupported)
				if err := _AlastriaContracts.contract.UnpackLog(event, "OperationWasNotSupported", log); err != nil {
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

// ParseOperationWasNotSupported is a log parse operation binding the contract event 0x0f759709e9a850d07441d2d9dbef884c5aae908ad08fe5e229ad74bb9bcef0fa.
//
// Solidity: event OperationWasNotSupported(string indexed method)
func (_AlastriaContracts *AlastriaContractsFilterer) ParseOperationWasNotSupported(log types.Log) (*AlastriaContractsOperationWasNotSupported, error) {
	event := new(AlastriaContractsOperationWasNotSupported)
	if err := _AlastriaContracts.contract.UnpackLog(event, "OperationWasNotSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaContractsPreparedAlastriaIDIterator is returned from FilterPreparedAlastriaID and is used to iterate over the raw logs and unpacked data for PreparedAlastriaID events raised by the AlastriaContracts contract.
type AlastriaContractsPreparedAlastriaIDIterator struct {
	Event *AlastriaContractsPreparedAlastriaID // Event containing the contract specifics and raw log

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
func (it *AlastriaContractsPreparedAlastriaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaContractsPreparedAlastriaID)
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
		it.Event = new(AlastriaContractsPreparedAlastriaID)
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
func (it *AlastriaContractsPreparedAlastriaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaContractsPreparedAlastriaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaContractsPreparedAlastriaID represents a PreparedAlastriaID event raised by the AlastriaContracts contract.
type AlastriaContractsPreparedAlastriaID struct {
	SignAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPreparedAlastriaID is a free log retrieval operation binding the contract event 0xb95ece7391a260096614d2b042c6cca83dcc5d131f72bc57a157ee2502e7ce53.
//
// Solidity: event PreparedAlastriaID(address indexed signAddress)
func (_AlastriaContracts *AlastriaContractsFilterer) FilterPreparedAlastriaID(opts *bind.FilterOpts, signAddress []common.Address) (*AlastriaContractsPreparedAlastriaIDIterator, error) {

	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _AlastriaContracts.contract.FilterLogs(opts, "PreparedAlastriaID", signAddressRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaContractsPreparedAlastriaIDIterator{contract: _AlastriaContracts.contract, event: "PreparedAlastriaID", logs: logs, sub: sub}, nil
}

// WatchPreparedAlastriaID is a free log subscription operation binding the contract event 0xb95ece7391a260096614d2b042c6cca83dcc5d131f72bc57a157ee2502e7ce53.
//
// Solidity: event PreparedAlastriaID(address indexed signAddress)
func (_AlastriaContracts *AlastriaContractsFilterer) WatchPreparedAlastriaID(opts *bind.WatchOpts, sink chan<- *AlastriaContractsPreparedAlastriaID, signAddress []common.Address) (event.Subscription, error) {

	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _AlastriaContracts.contract.WatchLogs(opts, "PreparedAlastriaID", signAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaContractsPreparedAlastriaID)
				if err := _AlastriaContracts.contract.UnpackLog(event, "PreparedAlastriaID", log); err != nil {
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

// ParsePreparedAlastriaID is a log parse operation binding the contract event 0xb95ece7391a260096614d2b042c6cca83dcc5d131f72bc57a157ee2502e7ce53.
//
// Solidity: event PreparedAlastriaID(address indexed signAddress)
func (_AlastriaContracts *AlastriaContractsFilterer) ParsePreparedAlastriaID(log types.Log) (*AlastriaContractsPreparedAlastriaID, error) {
	event := new(AlastriaContractsPreparedAlastriaID)
	if err := _AlastriaContracts.contract.UnpackLog(event, "PreparedAlastriaID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
