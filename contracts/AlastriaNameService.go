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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_firstIdentity\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"constant\":true,\"inputs\":[],\"name\":\"cif\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_AOA\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_createAID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"url_logo\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"}],\"name\":\"addEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"}],\"name\":\"setNameEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"}],\"name\":\"setCifEntity\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"}],\"name\":\"setUrlLogo\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"}],\"name\":\"setUrlCreateAID\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"}],\"name\":\"setUrlAOA\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_addressEntity\",\"type\":\"address\"}],\"name\":\"getEntity\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"_name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_cif\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_logo\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_createAID\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"_url_AOA\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"_active\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"entitiesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]",
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

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_Alastria *AlastriaCaller) Cif(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "cif")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_Alastria *AlastriaSession) Cif() (string, error) {
	return _Alastria.Contract.Cif(&_Alastria.CallOpts)
}

// Cif is a free data retrieval call binding the contract method 0x7f35d729.
//
// Solidity: function cif() view returns(string)
func (_Alastria *AlastriaCallerSession) Cif() (string, error) {
	return _Alastria.Contract.Cif(&_Alastria.CallOpts)
}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_Alastria *AlastriaCaller) EntitiesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "entitiesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_Alastria *AlastriaSession) EntitiesList() ([]common.Address, error) {
	return _Alastria.Contract.EntitiesList(&_Alastria.CallOpts)
}

// EntitiesList is a free data retrieval call binding the contract method 0x6e7a216b.
//
// Solidity: function entitiesList() view returns(address[])
func (_Alastria *AlastriaCallerSession) EntitiesList() ([]common.Address, error) {
	return _Alastria.Contract.EntitiesList(&_Alastria.CallOpts)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address _addressEntity) view returns(string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA, bool _active)
func (_Alastria *AlastriaCaller) GetEntity(opts *bind.CallOpts, _addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "getEntity", _addressEntity)

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
func (_Alastria *AlastriaSession) GetEntity(_addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	return _Alastria.Contract.GetEntity(&_Alastria.CallOpts, _addressEntity)
}

// GetEntity is a free data retrieval call binding the contract method 0x75894e8c.
//
// Solidity: function getEntity(address _addressEntity) view returns(string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA, bool _active)
func (_Alastria *AlastriaCallerSession) GetEntity(_addressEntity common.Address) (struct {
	Name         string
	Cif          string
	UrlLogo      string
	UrlCreateAID string
	UrlAOA       string
	Active       bool
}, error) {
	return _Alastria.Contract.GetEntity(&_Alastria.CallOpts, _addressEntity)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alastria *AlastriaCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alastria *AlastriaSession) Name() (string, error) {
	return _Alastria.Contract.Name(&_Alastria.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Alastria *AlastriaCallerSession) Name() (string, error) {
	return _Alastria.Contract.Name(&_Alastria.CallOpts)
}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_Alastria *AlastriaCaller) UrlAOA(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "url_AOA")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_Alastria *AlastriaSession) UrlAOA() (string, error) {
	return _Alastria.Contract.UrlAOA(&_Alastria.CallOpts)
}

// UrlAOA is a free data retrieval call binding the contract method 0xbf039d43.
//
// Solidity: function url_AOA() view returns(string)
func (_Alastria *AlastriaCallerSession) UrlAOA() (string, error) {
	return _Alastria.Contract.UrlAOA(&_Alastria.CallOpts)
}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_Alastria *AlastriaCaller) UrlCreateAID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "url_createAID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_Alastria *AlastriaSession) UrlCreateAID() (string, error) {
	return _Alastria.Contract.UrlCreateAID(&_Alastria.CallOpts)
}

// UrlCreateAID is a free data retrieval call binding the contract method 0x3bdf8863.
//
// Solidity: function url_createAID() view returns(string)
func (_Alastria *AlastriaCallerSession) UrlCreateAID() (string, error) {
	return _Alastria.Contract.UrlCreateAID(&_Alastria.CallOpts)
}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_Alastria *AlastriaCaller) UrlLogo(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Alastria.contract.Call(opts, &out, "url_logo")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_Alastria *AlastriaSession) UrlLogo() (string, error) {
	return _Alastria.Contract.UrlLogo(&_Alastria.CallOpts)
}

// UrlLogo is a free data retrieval call binding the contract method 0x500bebf3.
//
// Solidity: function url_logo() view returns(string)
func (_Alastria *AlastriaCallerSession) UrlLogo() (string, error) {
	return _Alastria.Contract.UrlLogo(&_Alastria.CallOpts)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_Alastria *AlastriaTransactor) AddEntity(opts *bind.TransactOpts, _addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "addEntity", _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_Alastria *AlastriaSession) AddEntity(_addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.Contract.AddEntity(&_Alastria.TransactOpts, _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// AddEntity is a paid mutator transaction binding the contract method 0x1252a083.
//
// Solidity: function addEntity(address _addressEntity, string _name, string _cif, string _url_logo, string _url_createAID, string _url_AOA) returns()
func (_Alastria *AlastriaTransactorSession) AddEntity(_addressEntity common.Address, _name string, _cif string, _url_logo string, _url_createAID string, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.Contract.AddEntity(&_Alastria.TransactOpts, _addressEntity, _name, _cif, _url_logo, _url_createAID, _url_AOA)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_Alastria *AlastriaTransactor) SetCifEntity(opts *bind.TransactOpts, _addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "setCifEntity", _addressEntity, _cif)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_Alastria *AlastriaSession) SetCifEntity(_addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _Alastria.Contract.SetCifEntity(&_Alastria.TransactOpts, _addressEntity, _cif)
}

// SetCifEntity is a paid mutator transaction binding the contract method 0x1e5f6f01.
//
// Solidity: function setCifEntity(address _addressEntity, string _cif) returns()
func (_Alastria *AlastriaTransactorSession) SetCifEntity(_addressEntity common.Address, _cif string) (*types.Transaction, error) {
	return _Alastria.Contract.SetCifEntity(&_Alastria.TransactOpts, _addressEntity, _cif)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_Alastria *AlastriaTransactor) SetNameEntity(opts *bind.TransactOpts, _addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "setNameEntity", _addressEntity, _name)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_Alastria *AlastriaSession) SetNameEntity(_addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _Alastria.Contract.SetNameEntity(&_Alastria.TransactOpts, _addressEntity, _name)
}

// SetNameEntity is a paid mutator transaction binding the contract method 0xeb5c426c.
//
// Solidity: function setNameEntity(address _addressEntity, string _name) returns()
func (_Alastria *AlastriaTransactorSession) SetNameEntity(_addressEntity common.Address, _name string) (*types.Transaction, error) {
	return _Alastria.Contract.SetNameEntity(&_Alastria.TransactOpts, _addressEntity, _name)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_Alastria *AlastriaTransactor) SetUrlAOA(opts *bind.TransactOpts, _addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "setUrlAOA", _addressEntity, _url_AOA)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_Alastria *AlastriaSession) SetUrlAOA(_addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlAOA(&_Alastria.TransactOpts, _addressEntity, _url_AOA)
}

// SetUrlAOA is a paid mutator transaction binding the contract method 0x8e672cf7.
//
// Solidity: function setUrlAOA(address _addressEntity, string _url_AOA) returns()
func (_Alastria *AlastriaTransactorSession) SetUrlAOA(_addressEntity common.Address, _url_AOA string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlAOA(&_Alastria.TransactOpts, _addressEntity, _url_AOA)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_Alastria *AlastriaTransactor) SetUrlCreateAID(opts *bind.TransactOpts, _addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "setUrlCreateAID", _addressEntity, _url_createAID)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_Alastria *AlastriaSession) SetUrlCreateAID(_addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlCreateAID(&_Alastria.TransactOpts, _addressEntity, _url_createAID)
}

// SetUrlCreateAID is a paid mutator transaction binding the contract method 0x4b7a1859.
//
// Solidity: function setUrlCreateAID(address _addressEntity, string _url_createAID) returns()
func (_Alastria *AlastriaTransactorSession) SetUrlCreateAID(_addressEntity common.Address, _url_createAID string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlCreateAID(&_Alastria.TransactOpts, _addressEntity, _url_createAID)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_Alastria *AlastriaTransactor) SetUrlLogo(opts *bind.TransactOpts, _addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _Alastria.contract.Transact(opts, "setUrlLogo", _addressEntity, _url_logo)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_Alastria *AlastriaSession) SetUrlLogo(_addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlLogo(&_Alastria.TransactOpts, _addressEntity, _url_logo)
}

// SetUrlLogo is a paid mutator transaction binding the contract method 0x0ff72ef2.
//
// Solidity: function setUrlLogo(address _addressEntity, string _url_logo) returns()
func (_Alastria *AlastriaTransactorSession) SetUrlLogo(_addressEntity common.Address, _url_logo string) (*types.Transaction, error) {
	return _Alastria.Contract.SetUrlLogo(&_Alastria.TransactOpts, _addressEntity, _url_logo)
}
