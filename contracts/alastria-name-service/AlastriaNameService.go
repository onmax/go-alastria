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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_firstIdentity\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"cif\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_AOA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_createAID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_logo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"}],\"name\":\"addEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNameEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"}],\"name\":\"setCifEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"}],\"name\":\"setUrlLogo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"}],\"name\":\"setUrlCreateAID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"}],\"name\":\"setUrlAOA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"}],\"name\":\"getEntity\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"entitiesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) Cif(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "cif")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) Cif() (string, error) {
	return _AlastriaContracts.Contract.Cif(&_AlastriaContracts.CallOpts)
}

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) Cif() (string, error) {
	return _AlastriaContracts.Contract.Cif(&_AlastriaContracts.CallOpts)
}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_AlastriaContracts *AlastriaContractsCaller) EntitiesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "entitiesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_AlastriaContracts *AlastriaContractsSession) EntitiesList() ([]common.Address, error) {
	return _AlastriaContracts.Contract.EntitiesList(&_AlastriaContracts.CallOpts)
}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_AlastriaContracts *AlastriaContractsCallerSession) EntitiesList() ([]common.Address, error) {
	return _AlastriaContracts.Contract.EntitiesList(&_AlastriaContracts.CallOpts)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address _addressEntity) view returns(string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA, bool _active)
func (_AlastriaContracts *AlastriaContractsCaller) GetEntity(opts *bind.CallOpts, _addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "getEntity", _addressEntity)

	outstruct := new(struct {
		Name         string
		Cif          string
		UrlLogo      string
		UrlCreateAID string
		UrlAOA       string
		Active       bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Cif = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.UrlLogo = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.UrlCreateAID = *abi.ConvertType(out[3], new(string)).(*string)
	outstruct.UrlAOA = *abi.ConvertType(out[4], new(string)).(*string)
	outstruct.Active = *abi.ConvertType(out[5], new(bool)).(*bool)

	return *outstruct, err

}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address _addressEntity) view returns(string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA, bool _active)
func (_AlastriaContracts *AlastriaContractsSession) GetEntity(_addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	return _AlastriaContracts.Contract.GetEntity(&_AlastriaContracts.CallOpts, _addressEntity)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address _addressEntity) view returns(string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA, bool _active)
func (_AlastriaContracts *AlastriaContractsCallerSession) GetEntity(_addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	return _AlastriaContracts.Contract.GetEntity(&_AlastriaContracts.CallOpts, _addressEntity)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) Name() (string, error) {
	return _AlastriaContracts.Contract.Name(&_AlastriaContracts.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) Name() (string, error) {
	return _AlastriaContracts.Contract.Name(&_AlastriaContracts.CallOpts)
}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) UrlAOA(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "url_AOA")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) UrlAOA() (string, error) {
	return _AlastriaContracts.Contract.UrlAOA(&_AlastriaContracts.CallOpts)
}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) UrlAOA() (string, error) {
	return _AlastriaContracts.Contract.UrlAOA(&_AlastriaContracts.CallOpts)
}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) UrlCreateAID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "url_createAID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) UrlCreateAID() (string, error) {
	return _AlastriaContracts.Contract.UrlCreateAID(&_AlastriaContracts.CallOpts)
}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) UrlCreateAID() (string, error) {
	return _AlastriaContracts.Contract.UrlCreateAID(&_AlastriaContracts.CallOpts)
}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_AlastriaContracts *AlastriaContractsCaller) UrlLogo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _AlastriaContracts.contract.Call(opts, &out, "url_logo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_AlastriaContracts *AlastriaContractsSession) UrlLogo() (string, error) {
	return _AlastriaContracts.Contract.UrlLogo(&_AlastriaContracts.CallOpts)
}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_AlastriaContracts *AlastriaContractsCallerSession) UrlLogo() (string, error) {
	return _AlastriaContracts.Contract.UrlLogo(&_AlastriaContracts.CallOpts)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) AddEntity(opts *bind.TransactOpts, _addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "addEntity", _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsSession) AddEntity(_addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) AddEntity(_addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.AddEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) SetCifEntity(opts *bind.TransactOpts, _addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "setCifEntity", _addressEntity, _cif)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_AlastriaContracts *AlastriaContractsSession) SetCifEntity(_addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetCifEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _cif)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) SetCifEntity(_addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetCifEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _cif)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) SetNameEntity(opts *bind.TransactOpts, _addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "setNameEntity", _addressEntity, _name)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_AlastriaContracts *AlastriaContractsSession) SetNameEntity(_addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetNameEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _name)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) SetNameEntity(_addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetNameEntity(&_AlastriaContracts.TransactOpts, _addressEntity, _name)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) SetUrlAOA(opts *bind.TransactOpts, _addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "setUrlAOA", _addressEntity, _url_AOA)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsSession) SetUrlAOA(_addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlAOA(&_AlastriaContracts.TransactOpts, _addressEntity, _url_AOA)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) SetUrlAOA(_addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlAOA(&_AlastriaContracts.TransactOpts, _addressEntity, _url_AOA)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) SetUrlCreateAID(opts *bind.TransactOpts, _addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "setUrlCreateAID", _addressEntity, _url_createAID)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_AlastriaContracts *AlastriaContractsSession) SetUrlCreateAID(_addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlCreateAID(&_AlastriaContracts.TransactOpts, _addressEntity, _url_createAID)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) SetUrlCreateAID(_addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlCreateAID(&_AlastriaContracts.TransactOpts, _addressEntity, _url_createAID)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_AlastriaContracts *AlastriaContractsTransactor) SetUrlLogo(opts *bind.TransactOpts, _addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _AlastriaContracts.contract.Transact(opts, "setUrlLogo", _addressEntity, _url_logo)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_AlastriaContracts *AlastriaContractsSession) SetUrlLogo(_addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlLogo(&_AlastriaContracts.TransactOpts, _addressEntity, _url_logo)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_AlastriaContracts *AlastriaContractsTransactorSession) SetUrlLogo(_addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _AlastriaContracts.Contract.SetUrlLogo(&_AlastriaContracts.TransactOpts, _addressEntity, _url_logo)
}
