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
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"identity\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"IdentityCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldAccount\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAccount\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"serviceProvider\",\"type\":\"address\"}],\"name\":\"IdentityRecovered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"string\",\"name\":\"method\",\"type\":\"string\"}],\"name\":\"OperationWasNotSupported\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"signAddress\",\"type\":\"address\"}],\"name\":\"PreparedAlastriaID\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"addIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"addIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaCredentialRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaCredentialRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaPresentationRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaPresentationRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"alastriaPublicKeyRegistry\",\"outputs\":[{\"internalType\":\"contractAlastriaPublicKeyRegistry\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"deleteIdentityIssuer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"deleteIdentityServiceProvider\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"firstIdentityWallet\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"getEidasLevel\",\"outputs\":[{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"identityKeys\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"}],\"name\":\"isIdentityIssuer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityServiceProvider\",\"type\":\"address\"}],\"name\":\"isIdentityServiceProvider\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"isOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingIDs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_identityIssuer\",\"type\":\"address\"},{\"internalType\":\"enumEidas.EidasLevel\",\"name\":\"_level\",\"type\":\"uint8\"}],\"name\":\"updateIdentityIssuerEidasLevel\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_credentialRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_publicKeyRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_presentationRegistry\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_firstIdentityWallet\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_signAddress\",\"type\":\"address\"}],\"name\":\"prepareAlastriaID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"addPublicKeyCallData\",\"type\":\"bytes\"}],\"name\":\"createAlastriaIdentity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_destination\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"delegateCall\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"accountLost\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAccount\",\"type\":\"address\"}],\"name\":\"recoverAccount\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
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

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_Alastria *AlastriaCaller) AlastriaCredentialRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "alastriaCredentialRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_Alastria *AlastriaSession) AlastriaCredentialRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaCredentialRegistry(&_Alastria.CallOpts)
}

// AlastriaCredentialRegistry is a free data retrieval call binding the contract method 0xdcc30e35.
//
// Solidity: function alastriaCredentialRegistry() view returns(address)
func (_Alastria *AlastriaCallerSession) AlastriaCredentialRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaCredentialRegistry(&_Alastria.CallOpts)
}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_Alastria *AlastriaCaller) AlastriaPresentationRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "alastriaPresentationRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_Alastria *AlastriaSession) AlastriaPresentationRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaPresentationRegistry(&_Alastria.CallOpts)
}

// AlastriaPresentationRegistry is a free data retrieval call binding the contract method 0x43eedc70.
//
// Solidity: function alastriaPresentationRegistry() view returns(address)
func (_Alastria *AlastriaCallerSession) AlastriaPresentationRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaPresentationRegistry(&_Alastria.CallOpts)
}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_Alastria *AlastriaCaller) AlastriaPublicKeyRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "alastriaPublicKeyRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_Alastria *AlastriaSession) AlastriaPublicKeyRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaPublicKeyRegistry(&_Alastria.CallOpts)
}

// AlastriaPublicKeyRegistry is a free data retrieval call binding the contract method 0xf014cd41.
//
// Solidity: function alastriaPublicKeyRegistry() view returns(address)
func (_Alastria *AlastriaCallerSession) AlastriaPublicKeyRegistry() (common.Address, error) {
	return _Alastria.Contract.AlastriaPublicKeyRegistry(&_Alastria.CallOpts)
}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_Alastria *AlastriaCaller) FirstIdentityWallet(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "firstIdentityWallet")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_Alastria *AlastriaSession) FirstIdentityWallet() (common.Address, error) {
	return _Alastria.Contract.FirstIdentityWallet(&_Alastria.CallOpts)
}

// FirstIdentityWallet is a free data retrieval call binding the contract method 0x6de41372.
//
// Solidity: function firstIdentityWallet() view returns(address)
func (_Alastria *AlastriaCallerSession) FirstIdentityWallet() (common.Address, error) {
	return _Alastria.Contract.FirstIdentityWallet(&_Alastria.CallOpts)
}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_Alastria *AlastriaCaller) GetEidasLevel(opts *bind.CallOpts, _identityIssuer common.Address) (uint8, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getEidasLevel", _identityIssuer)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_Alastria *AlastriaSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _Alastria.Contract.GetEidasLevel(&_Alastria.CallOpts, _identityIssuer)
}

// GetEidasLevel is a free data retrieval call binding the contract method 0x0e5a4fbb.
//
// Solidity: function getEidasLevel(address _identityIssuer) view returns(uint8)
func (_Alastria *AlastriaCallerSession) GetEidasLevel(_identityIssuer common.Address) (uint8, error) {
	return _Alastria.Contract.GetEidasLevel(&_Alastria.CallOpts, _identityIssuer)
}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_Alastria *AlastriaCaller) IdentityKeys(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "identityKeys", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_Alastria *AlastriaSession) IdentityKeys(arg0 common.Address) (common.Address, error) {
	return _Alastria.Contract.IdentityKeys(&_Alastria.CallOpts, arg0)
}

// IdentityKeys is a free data retrieval call binding the contract method 0x0c91465e.
//
// Solidity: function identityKeys(address ) view returns(address)
func (_Alastria *AlastriaCallerSession) IdentityKeys(arg0 common.Address) (common.Address, error) {
	return _Alastria.Contract.IdentityKeys(&_Alastria.CallOpts, arg0)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_Alastria *AlastriaCaller) IsIdentityIssuer(opts *bind.CallOpts, _identityIssuer common.Address) (bool, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "isIdentityIssuer", _identityIssuer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_Alastria *AlastriaSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityIssuer(&_Alastria.CallOpts, _identityIssuer)
}

// IsIdentityIssuer is a free data retrieval call binding the contract method 0x6554adf0.
//
// Solidity: function isIdentityIssuer(address _identityIssuer) view returns(bool)
func (_Alastria *AlastriaCallerSession) IsIdentityIssuer(_identityIssuer common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityIssuer(&_Alastria.CallOpts, _identityIssuer)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaCaller) IsIdentityServiceProvider(opts *bind.CallOpts, _identityServiceProvider common.Address) (bool, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "isIdentityServiceProvider", _identityServiceProvider)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityServiceProvider(&_Alastria.CallOpts, _identityServiceProvider)
}

// IsIdentityServiceProvider is a free data retrieval call binding the contract method 0xd024d9a4.
//
// Solidity: function isIdentityServiceProvider(address _identityServiceProvider) view returns(bool)
func (_Alastria *AlastriaCallerSession) IsIdentityServiceProvider(_identityServiceProvider common.Address) (bool, error) {
	return _Alastria.Contract.IsIdentityServiceProvider(&_Alastria.CallOpts, _identityServiceProvider)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_Alastria *AlastriaCaller) IsOwner(opts *bind.CallOpts, addr common.Address) (bool, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "isOwner", addr)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_Alastria *AlastriaSession) IsOwner(addr common.Address) (bool, error) {
	return _Alastria.Contract.IsOwner(&_Alastria.CallOpts, addr)
}

// IsOwner is a free data retrieval call binding the contract method 0x2f54bf6e.
//
// Solidity: function isOwner(address addr) view returns(bool)
func (_Alastria *AlastriaCallerSession) IsOwner(addr common.Address) (bool, error) {
	return _Alastria.Contract.IsOwner(&_Alastria.CallOpts, addr)
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

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_Alastria *AlastriaCaller) PendingIDs(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "pendingIDs", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_Alastria *AlastriaSession) PendingIDs(arg0 common.Address) (*big.Int, error) {
	return _Alastria.Contract.PendingIDs(&_Alastria.CallOpts, arg0)
}

// PendingIDs is a free data retrieval call binding the contract method 0x2773d0fa.
//
// Solidity: function pendingIDs(address ) view returns(uint256)
func (_Alastria *AlastriaCallerSession) PendingIDs(arg0 common.Address) (*big.Int, error) {
	return _Alastria.Contract.PendingIDs(&_Alastria.CallOpts, arg0)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
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
// Solidity: function version() view returns(uint256)
func (_Alastria *AlastriaSession) Version() (*big.Int, error) {
	return _Alastria.Contract.Version(&_Alastria.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint256)
func (_Alastria *AlastriaCallerSession) Version() (*big.Int, error) {
	return _Alastria.Contract.Version(&_Alastria.CallOpts)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaTransactor) AddIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addIdentityIssuer", _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityIssuer(&_Alastria.TransactOpts, _identityIssuer, _level)
}

// AddIdentityIssuer is a paid mutator transaction binding the contract method 0x889776a8.
//
// Solidity: function addIdentityIssuer(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaTransactorSession) AddIdentityIssuer(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityIssuer(&_Alastria.TransactOpts, _identityIssuer, _level)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactor) AddIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addIdentityServiceProvider", _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// AddIdentityServiceProvider is a paid mutator transaction binding the contract method 0x0ebbbffc.
//
// Solidity: function addIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactorSession) AddIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.AddIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_Alastria *AlastriaTransactor) CreateAlastriaIdentity(opts *bind.TransactOpts, addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "createAlastriaIdentity", addPublicKeyCallData)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_Alastria *AlastriaSession) CreateAlastriaIdentity(addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _Alastria.Contract.CreateAlastriaIdentity(&_Alastria.TransactOpts, addPublicKeyCallData)
}

// CreateAlastriaIdentity is a paid mutator transaction binding the contract method 0x6d69d99a.
//
// Solidity: function createAlastriaIdentity(bytes addPublicKeyCallData) returns()
func (_Alastria *AlastriaTransactorSession) CreateAlastriaIdentity(addPublicKeyCallData []byte) (*types.Transaction, error) {
	return _Alastria.Contract.CreateAlastriaIdentity(&_Alastria.TransactOpts, addPublicKeyCallData)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Alastria *AlastriaTransactor) DelegateCall(opts *bind.TransactOpts, _destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "delegateCall", _destination, _value, _data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Alastria *AlastriaSession) DelegateCall(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Alastria.Contract.DelegateCall(&_Alastria.TransactOpts, _destination, _value, _data)
}

// DelegateCall is a paid mutator transaction binding the contract method 0x597b2e9b.
//
// Solidity: function delegateCall(address _destination, uint256 _value, bytes _data) returns(bytes)
func (_Alastria *AlastriaTransactorSession) DelegateCall(_destination common.Address, _value *big.Int, _data []byte) (*types.Transaction, error) {
	return _Alastria.Contract.DelegateCall(&_Alastria.TransactOpts, _destination, _value, _data)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_Alastria *AlastriaTransactor) DeleteIdentityIssuer(opts *bind.TransactOpts, _identityIssuer common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "deleteIdentityIssuer", _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_Alastria *AlastriaSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityIssuer(&_Alastria.TransactOpts, _identityIssuer)
}

// DeleteIdentityIssuer is a paid mutator transaction binding the contract method 0xcb691599.
//
// Solidity: function deleteIdentityIssuer(address _identityIssuer) returns()
func (_Alastria *AlastriaTransactorSession) DeleteIdentityIssuer(_identityIssuer common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityIssuer(&_Alastria.TransactOpts, _identityIssuer)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactor) DeleteIdentityServiceProvider(opts *bind.TransactOpts, _identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "deleteIdentityServiceProvider", _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// DeleteIdentityServiceProvider is a paid mutator transaction binding the contract method 0x3bf47215.
//
// Solidity: function deleteIdentityServiceProvider(address _identityServiceProvider) returns()
func (_Alastria *AlastriaTransactorSession) DeleteIdentityServiceProvider(_identityServiceProvider common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.DeleteIdentityServiceProvider(&_Alastria.TransactOpts, _identityServiceProvider)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_Alastria *AlastriaTransactor) Initialize(opts *bind.TransactOpts, _credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "initialize", _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_Alastria *AlastriaSession) Initialize(_credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address _credentialRegistry, address _publicKeyRegistry, address _presentationRegistry, address _firstIdentityWallet) returns()
func (_Alastria *AlastriaTransactorSession) Initialize(_credentialRegistry common.Address, _publicKeyRegistry common.Address, _presentationRegistry common.Address, _firstIdentityWallet common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Initialize(&_Alastria.TransactOpts, _credentialRegistry, _publicKeyRegistry, _presentationRegistry, _firstIdentityWallet)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_Alastria *AlastriaTransactor) PrepareAlastriaID(opts *bind.TransactOpts, _signAddress common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "prepareAlastriaID", _signAddress)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_Alastria *AlastriaSession) PrepareAlastriaID(_signAddress common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.PrepareAlastriaID(&_Alastria.TransactOpts, _signAddress)
}

// PrepareAlastriaID is a paid mutator transaction binding the contract method 0x45748b42.
//
// Solidity: function prepareAlastriaID(address _signAddress) returns()
func (_Alastria *AlastriaTransactorSession) PrepareAlastriaID(_signAddress common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.PrepareAlastriaID(&_Alastria.TransactOpts, _signAddress)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_Alastria *AlastriaTransactor) RecoverAccount(opts *bind.TransactOpts, accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "recoverAccount", accountLost, newAccount)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_Alastria *AlastriaSession) RecoverAccount(accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.RecoverAccount(&_Alastria.TransactOpts, accountLost, newAccount)
}

// RecoverAccount is a paid mutator transaction binding the contract method 0xc71cbcf3.
//
// Solidity: function recoverAccount(address accountLost, address newAccount) returns()
func (_Alastria *AlastriaTransactorSession) RecoverAccount(accountLost common.Address, newAccount common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.RecoverAccount(&_Alastria.TransactOpts, accountLost, newAccount)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_Alastria *AlastriaTransactor) Transfer(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "transfer", newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_Alastria *AlastriaSession) Transfer(newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Transfer(&_Alastria.TransactOpts, newOwner)
}

// Transfer is a paid mutator transaction binding the contract method 0x1a695230.
//
// Solidity: function transfer(address newOwner) returns()
func (_Alastria *AlastriaTransactorSession) Transfer(newOwner common.Address) (*types.Transaction, error) {
	return _Alastria.Contract.Transfer(&_Alastria.TransactOpts, newOwner)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaTransactor) UpdateIdentityIssuerEidasLevel(opts *bind.TransactOpts, _identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "updateIdentityIssuerEidasLevel", _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateIdentityIssuerEidasLevel(&_Alastria.TransactOpts, _identityIssuer, _level)
}

// UpdateIdentityIssuerEidasLevel is a paid mutator transaction binding the contract method 0x44963610.
//
// Solidity: function updateIdentityIssuerEidasLevel(address _identityIssuer, uint8 _level) returns()
func (_Alastria *AlastriaTransactorSession) UpdateIdentityIssuerEidasLevel(_identityIssuer common.Address, _level uint8) (*types.Transaction, error) {
	return _Alastria.Contract.UpdateIdentityIssuerEidasLevel(&_Alastria.TransactOpts, _identityIssuer, _level)
}

// AlastriaIdentityCreatedIterator is returned from FilterIdentityCreated and is used to iterate over the raw logs and unpacked data for IdentityCreated events raised by the Alastria contract.
type AlastriaIdentityCreatedIterator struct {
	Event *AlastriaIdentityCreated // Event containing the contract specifics and raw log

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
func (it *AlastriaIdentityCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaIdentityCreated)
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
		it.Event = new(AlastriaIdentityCreated)
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
func (it *AlastriaIdentityCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaIdentityCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaIdentityCreated represents a IdentityCreated event raised by the Alastria contract.
type AlastriaIdentityCreated struct {
	Identity common.Address
	Creator  common.Address
	Owner    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterIdentityCreated is a free log retrieval operation binding the contract event 0x8f2e597fd6e795e9851eea530f987cceb641315e3c7c9d484b798e3159d42095.
//
// Solidity: event IdentityCreated(address indexed identity, address indexed creator, address owner)
func (_Alastria *AlastriaFilterer) FilterIdentityCreated(opts *bind.FilterOpts, identity []common.Address, creator []common.Address) (*AlastriaIdentityCreatedIterator, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "IdentityCreated", identityRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaIdentityCreatedIterator{contract: _Alastria.contract, event: "IdentityCreated", logs: logs, sub: sub}, nil
}

// WatchIdentityCreated is a free log subscription operation binding the contract event 0x8f2e597fd6e795e9851eea530f987cceb641315e3c7c9d484b798e3159d42095.
//
// Solidity: event IdentityCreated(address indexed identity, address indexed creator, address owner)
func (_Alastria *AlastriaFilterer) WatchIdentityCreated(opts *bind.WatchOpts, sink chan<- *AlastriaIdentityCreated, identity []common.Address, creator []common.Address) (event.Subscription, error) {

	var identityRule []interface{}
	for _, identityItem := range identity {
		identityRule = append(identityRule, identityItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "IdentityCreated", identityRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaIdentityCreated)
				if err := _Alastria.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParseIdentityCreated(log types.Log) (*AlastriaIdentityCreated, error) {
	event := new(AlastriaIdentityCreated)
	if err := _Alastria.contract.UnpackLog(event, "IdentityCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaIdentityRecoveredIterator is returned from FilterIdentityRecovered and is used to iterate over the raw logs and unpacked data for IdentityRecovered events raised by the Alastria contract.
type AlastriaIdentityRecoveredIterator struct {
	Event *AlastriaIdentityRecovered // Event containing the contract specifics and raw log

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
func (it *AlastriaIdentityRecoveredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaIdentityRecovered)
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
		it.Event = new(AlastriaIdentityRecovered)
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
func (it *AlastriaIdentityRecoveredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaIdentityRecoveredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaIdentityRecovered represents a IdentityRecovered event raised by the Alastria contract.
type AlastriaIdentityRecovered struct {
	OldAccount      common.Address
	NewAccount      common.Address
	ServiceProvider common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterIdentityRecovered is a free log retrieval operation binding the contract event 0xf5cf2e915faa705d003410f0f65cd59606d22f8e7fb7361a015d6a00b7723d99.
//
// Solidity: event IdentityRecovered(address indexed oldAccount, address newAccount, address indexed serviceProvider)
func (_Alastria *AlastriaFilterer) FilterIdentityRecovered(opts *bind.FilterOpts, oldAccount []common.Address, serviceProvider []common.Address) (*AlastriaIdentityRecoveredIterator, error) {

	var oldAccountRule []interface{}
	for _, oldAccountItem := range oldAccount {
		oldAccountRule = append(oldAccountRule, oldAccountItem)
	}

	var serviceProviderRule []interface{}
	for _, serviceProviderItem := range serviceProvider {
		serviceProviderRule = append(serviceProviderRule, serviceProviderItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "IdentityRecovered", oldAccountRule, serviceProviderRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaIdentityRecoveredIterator{contract: _Alastria.contract, event: "IdentityRecovered", logs: logs, sub: sub}, nil
}

// WatchIdentityRecovered is a free log subscription operation binding the contract event 0xf5cf2e915faa705d003410f0f65cd59606d22f8e7fb7361a015d6a00b7723d99.
//
// Solidity: event IdentityRecovered(address indexed oldAccount, address newAccount, address indexed serviceProvider)
func (_Alastria *AlastriaFilterer) WatchIdentityRecovered(opts *bind.WatchOpts, sink chan<- *AlastriaIdentityRecovered, oldAccount []common.Address, serviceProvider []common.Address) (event.Subscription, error) {

	var oldAccountRule []interface{}
	for _, oldAccountItem := range oldAccount {
		oldAccountRule = append(oldAccountRule, oldAccountItem)
	}

	var serviceProviderRule []interface{}
	for _, serviceProviderItem := range serviceProvider {
		serviceProviderRule = append(serviceProviderRule, serviceProviderItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "IdentityRecovered", oldAccountRule, serviceProviderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaIdentityRecovered)
				if err := _Alastria.contract.UnpackLog(event, "IdentityRecovered", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParseIdentityRecovered(log types.Log) (*AlastriaIdentityRecovered, error) {
	event := new(AlastriaIdentityRecovered)
	if err := _Alastria.contract.UnpackLog(event, "IdentityRecovered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaOperationWasNotSupportedIterator is returned from FilterOperationWasNotSupported and is used to iterate over the raw logs and unpacked data for OperationWasNotSupported events raised by the Alastria contract.
type AlastriaOperationWasNotSupportedIterator struct {
	Event *AlastriaOperationWasNotSupported // Event containing the contract specifics and raw log

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
func (it *AlastriaOperationWasNotSupportedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaOperationWasNotSupported)
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
		it.Event = new(AlastriaOperationWasNotSupported)
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
func (it *AlastriaOperationWasNotSupportedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaOperationWasNotSupportedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaOperationWasNotSupported represents a OperationWasNotSupported event raised by the Alastria contract.
type AlastriaOperationWasNotSupported struct {
	Method common.Hash
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOperationWasNotSupported is a free log retrieval operation binding the contract event 0x0f759709e9a850d07441d2d9dbef884c5aae908ad08fe5e229ad74bb9bcef0fa.
//
// Solidity: event OperationWasNotSupported(string indexed method)
func (_Alastria *AlastriaFilterer) FilterOperationWasNotSupported(opts *bind.FilterOpts, method []string) (*AlastriaOperationWasNotSupportedIterator, error) {

	var methodRule []interface{}
	for _, methodItem := range method {
		methodRule = append(methodRule, methodItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "OperationWasNotSupported", methodRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaOperationWasNotSupportedIterator{contract: _Alastria.contract, event: "OperationWasNotSupported", logs: logs, sub: sub}, nil
}

// WatchOperationWasNotSupported is a free log subscription operation binding the contract event 0x0f759709e9a850d07441d2d9dbef884c5aae908ad08fe5e229ad74bb9bcef0fa.
//
// Solidity: event OperationWasNotSupported(string indexed method)
func (_Alastria *AlastriaFilterer) WatchOperationWasNotSupported(opts *bind.WatchOpts, sink chan<- *AlastriaOperationWasNotSupported, method []string) (event.Subscription, error) {

	var methodRule []interface{}
	for _, methodItem := range method {
		methodRule = append(methodRule, methodItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "OperationWasNotSupported", methodRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaOperationWasNotSupported)
				if err := _Alastria.contract.UnpackLog(event, "OperationWasNotSupported", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParseOperationWasNotSupported(log types.Log) (*AlastriaOperationWasNotSupported, error) {
	event := new(AlastriaOperationWasNotSupported)
	if err := _Alastria.contract.UnpackLog(event, "OperationWasNotSupported", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AlastriaPreparedAlastriaIDIterator is returned from FilterPreparedAlastriaID and is used to iterate over the raw logs and unpacked data for PreparedAlastriaID events raised by the Alastria contract.
type AlastriaPreparedAlastriaIDIterator struct {
	Event *AlastriaPreparedAlastriaID // Event containing the contract specifics and raw log

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
func (it *AlastriaPreparedAlastriaIDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AlastriaPreparedAlastriaID)
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
		it.Event = new(AlastriaPreparedAlastriaID)
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
func (it *AlastriaPreparedAlastriaIDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AlastriaPreparedAlastriaIDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AlastriaPreparedAlastriaID represents a PreparedAlastriaID event raised by the Alastria contract.
type AlastriaPreparedAlastriaID struct {
	SignAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPreparedAlastriaID is a free log retrieval operation binding the contract event 0xb95ece7391a260096614d2b042c6cca83dcc5d131f72bc57a157ee2502e7ce53.
//
// Solidity: event PreparedAlastriaID(address indexed signAddress)
func (_Alastria *AlastriaFilterer) FilterPreparedAlastriaID(opts *bind.FilterOpts, signAddress []common.Address) (*AlastriaPreparedAlastriaIDIterator, error) {

	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _Alastria.contract.FilterLogs(opts, "PreparedAlastriaID", signAddressRule)
	if err != nil {
		return nil, err
	}
	return &AlastriaPreparedAlastriaIDIterator{contract: _Alastria.contract, event: "PreparedAlastriaID", logs: logs, sub: sub}, nil
}

// WatchPreparedAlastriaID is a free log subscription operation binding the contract event 0xb95ece7391a260096614d2b042c6cca83dcc5d131f72bc57a157ee2502e7ce53.
//
// Solidity: event PreparedAlastriaID(address indexed signAddress)
func (_Alastria *AlastriaFilterer) WatchPreparedAlastriaID(opts *bind.WatchOpts, sink chan<- *AlastriaPreparedAlastriaID, signAddress []common.Address) (event.Subscription, error) {

	var signAddressRule []interface{}
	for _, signAddressItem := range signAddress {
		signAddressRule = append(signAddressRule, signAddressItem)
	}

	logs, sub, err := _Alastria.contract.WatchLogs(opts, "PreparedAlastriaID", signAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AlastriaPreparedAlastriaID)
				if err := _Alastria.contract.UnpackLog(event, "PreparedAlastriaID", log); err != nil {
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
func (_Alastria *AlastriaFilterer) ParsePreparedAlastriaID(log types.Log) (*AlastriaPreparedAlastriaID, error) {
	event := new(AlastriaPreparedAlastriaID)
	if err := _Alastria.contract.UnpackLog(event, "PreparedAlastriaID", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
