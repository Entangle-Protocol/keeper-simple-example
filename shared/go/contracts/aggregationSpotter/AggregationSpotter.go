// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aggregationSpotter

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

// AggregationSpotterOperationData is an auto generated low-level Go binding around an user-defined struct.
type AggregationSpotterOperationData struct {
	Contr            common.Address
	FunctionSelector [4]byte
	Params           []byte
}

// AggregationSpotterMetaData contains all meta data concerning the AggregationSpotter contract.
var AggregationSpotterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AggregationSpotter__CallFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AggregationSpotter__ContractIsNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AggregationSpotter__KeeperIsAlreadyApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"AggregationSpotter__KeeperIsNotAllowed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AggregationSpotter__OpIsNotApprovedOrExecuted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AggregationSpotter__OperationDoesNotExist\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"AggregationSpotter__OperationIsAlreadyApproved\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"opHash\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cont\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"}],\"name\":\"NewOperation\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"opHash\",\"type\":\"uint256\"}],\"name\":\"ProposalApproved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"opHash\",\"type\":\"uint256\"}],\"name\":\"ProposalExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"EXECUTOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"addAllowedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"addKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"opHash\",\"type\":\"uint256\"}],\"name\":\"executeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[2]\",\"name\":\"initAddr\",\"type\":\"address[2]\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"oracleOpTxId\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"contr\",\"type\":\"address\"},{\"internalType\":\"bytes4\",\"name\":\"functionSelector\",\"type\":\"bytes4\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"}],\"internalType\":\"structAggregationSpotter.OperationData\",\"name\":\"opData\",\"type\":\"tuple\"}],\"name\":\"proposeOperation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"removeAllowedContract\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"keeper\",\"type\":\"address\"}],\"name\":\"removeKeeper\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"setConsensusTargetRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// AggregationSpotterABI is the input ABI used to generate the binding from.
// Deprecated: Use AggregationSpotterMetaData.ABI instead.
var AggregationSpotterABI = AggregationSpotterMetaData.ABI

// AggregationSpotter is an auto generated Go binding around an Ethereum contract.
type AggregationSpotter struct {
	AggregationSpotterCaller     // Read-only binding to the contract
	AggregationSpotterTransactor // Write-only binding to the contract
	AggregationSpotterFilterer   // Log filterer for contract events
}

// AggregationSpotterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AggregationSpotterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationSpotterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AggregationSpotterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationSpotterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AggregationSpotterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AggregationSpotterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AggregationSpotterSession struct {
	Contract     *AggregationSpotter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// AggregationSpotterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AggregationSpotterCallerSession struct {
	Contract *AggregationSpotterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// AggregationSpotterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AggregationSpotterTransactorSession struct {
	Contract     *AggregationSpotterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// AggregationSpotterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AggregationSpotterRaw struct {
	Contract *AggregationSpotter // Generic contract binding to access the raw methods on
}

// AggregationSpotterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AggregationSpotterCallerRaw struct {
	Contract *AggregationSpotterCaller // Generic read-only contract binding to access the raw methods on
}

// AggregationSpotterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AggregationSpotterTransactorRaw struct {
	Contract *AggregationSpotterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAggregationSpotter creates a new instance of AggregationSpotter, bound to a specific deployed contract.
func NewAggregationSpotter(address common.Address, backend bind.ContractBackend) (*AggregationSpotter, error) {
	contract, err := bindAggregationSpotter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotter{AggregationSpotterCaller: AggregationSpotterCaller{contract: contract}, AggregationSpotterTransactor: AggregationSpotterTransactor{contract: contract}, AggregationSpotterFilterer: AggregationSpotterFilterer{contract: contract}}, nil
}

// NewAggregationSpotterCaller creates a new read-only instance of AggregationSpotter, bound to a specific deployed contract.
func NewAggregationSpotterCaller(address common.Address, caller bind.ContractCaller) (*AggregationSpotterCaller, error) {
	contract, err := bindAggregationSpotter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterCaller{contract: contract}, nil
}

// NewAggregationSpotterTransactor creates a new write-only instance of AggregationSpotter, bound to a specific deployed contract.
func NewAggregationSpotterTransactor(address common.Address, transactor bind.ContractTransactor) (*AggregationSpotterTransactor, error) {
	contract, err := bindAggregationSpotter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterTransactor{contract: contract}, nil
}

// NewAggregationSpotterFilterer creates a new log filterer instance of AggregationSpotter, bound to a specific deployed contract.
func NewAggregationSpotterFilterer(address common.Address, filterer bind.ContractFilterer) (*AggregationSpotterFilterer, error) {
	contract, err := bindAggregationSpotter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterFilterer{contract: contract}, nil
}

// bindAggregationSpotter binds a generic wrapper to an already deployed contract.
func bindAggregationSpotter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AggregationSpotterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationSpotter *AggregationSpotterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationSpotter.Contract.AggregationSpotterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationSpotter *AggregationSpotterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AggregationSpotterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationSpotter *AggregationSpotterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AggregationSpotterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AggregationSpotter *AggregationSpotterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AggregationSpotter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AggregationSpotter *AggregationSpotterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AggregationSpotter *AggregationSpotterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.contract.Transact(opts, method, params...)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCaller) ADMIN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "ADMIN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterSession) ADMIN() ([32]byte, error) {
	return _AggregationSpotter.Contract.ADMIN(&_AggregationSpotter.CallOpts)
}

// ADMIN is a free data retrieval call binding the contract method 0x2a0acc6a.
//
// Solidity: function ADMIN() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCallerSession) ADMIN() ([32]byte, error) {
	return _AggregationSpotter.Contract.ADMIN(&_AggregationSpotter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AggregationSpotter.Contract.DEFAULTADMINROLE(&_AggregationSpotter.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _AggregationSpotter.Contract.DEFAULTADMINROLE(&_AggregationSpotter.CallOpts)
}

// EXECUTOR is a free data retrieval call binding the contract method 0x630dc7cb.
//
// Solidity: function EXECUTOR() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCaller) EXECUTOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "EXECUTOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EXECUTOR is a free data retrieval call binding the contract method 0x630dc7cb.
//
// Solidity: function EXECUTOR() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterSession) EXECUTOR() ([32]byte, error) {
	return _AggregationSpotter.Contract.EXECUTOR(&_AggregationSpotter.CallOpts)
}

// EXECUTOR is a free data retrieval call binding the contract method 0x630dc7cb.
//
// Solidity: function EXECUTOR() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCallerSession) EXECUTOR() ([32]byte, error) {
	return _AggregationSpotter.Contract.EXECUTOR(&_AggregationSpotter.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AggregationSpotter.Contract.GetRoleAdmin(&_AggregationSpotter.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _AggregationSpotter.Contract.GetRoleAdmin(&_AggregationSpotter.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AggregationSpotter *AggregationSpotterCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AggregationSpotter *AggregationSpotterSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AggregationSpotter.Contract.HasRole(&_AggregationSpotter.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_AggregationSpotter *AggregationSpotterCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _AggregationSpotter.Contract.HasRole(&_AggregationSpotter.CallOpts, role, account)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationSpotter *AggregationSpotterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationSpotter *AggregationSpotterSession) Owner() (common.Address, error) {
	return _AggregationSpotter.Contract.Owner(&_AggregationSpotter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AggregationSpotter *AggregationSpotterCallerSession) Owner() (common.Address, error) {
	return _AggregationSpotter.Contract.Owner(&_AggregationSpotter.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterSession) ProxiableUUID() ([32]byte, error) {
	return _AggregationSpotter.Contract.ProxiableUUID(&_AggregationSpotter.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_AggregationSpotter *AggregationSpotterCallerSession) ProxiableUUID() ([32]byte, error) {
	return _AggregationSpotter.Contract.ProxiableUUID(&_AggregationSpotter.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AggregationSpotter *AggregationSpotterCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _AggregationSpotter.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AggregationSpotter *AggregationSpotterSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AggregationSpotter.Contract.SupportsInterface(&_AggregationSpotter.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_AggregationSpotter *AggregationSpotterCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _AggregationSpotter.Contract.SupportsInterface(&_AggregationSpotter.CallOpts, interfaceId)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x2c56462f.
//
// Solidity: function addAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) AddAllowedContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "addAllowedContract", _contract)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x2c56462f.
//
// Solidity: function addAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterSession) AddAllowedContract(_contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AddAllowedContract(&_AggregationSpotter.TransactOpts, _contract)
}

// AddAllowedContract is a paid mutator transaction binding the contract method 0x2c56462f.
//
// Solidity: function addAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) AddAllowedContract(_contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AddAllowedContract(&_AggregationSpotter.TransactOpts, _contract)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) AddKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "addKeeper", keeper)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterSession) AddKeeper(keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AddKeeper(&_AggregationSpotter.TransactOpts, keeper)
}

// AddKeeper is a paid mutator transaction binding the contract method 0x4032b72b.
//
// Solidity: function addKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) AddKeeper(keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.AddKeeper(&_AggregationSpotter.TransactOpts, keeper)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x0e70306d.
//
// Solidity: function executeOperation(uint256 opHash) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) ExecuteOperation(opts *bind.TransactOpts, opHash *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "executeOperation", opHash)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x0e70306d.
//
// Solidity: function executeOperation(uint256 opHash) returns()
func (_AggregationSpotter *AggregationSpotterSession) ExecuteOperation(opHash *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.ExecuteOperation(&_AggregationSpotter.TransactOpts, opHash)
}

// ExecuteOperation is a paid mutator transaction binding the contract method 0x0e70306d.
//
// Solidity: function executeOperation(uint256 opHash) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) ExecuteOperation(opHash *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.ExecuteOperation(&_AggregationSpotter.TransactOpts, opHash)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.GrantRole(&_AggregationSpotter.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.GrantRole(&_AggregationSpotter.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8c6d49a.
//
// Solidity: function initialize(address[2] initAddr) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) Initialize(opts *bind.TransactOpts, initAddr [2]common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "initialize", initAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8c6d49a.
//
// Solidity: function initialize(address[2] initAddr) returns()
func (_AggregationSpotter *AggregationSpotterSession) Initialize(initAddr [2]common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.Initialize(&_AggregationSpotter.TransactOpts, initAddr)
}

// Initialize is a paid mutator transaction binding the contract method 0xd8c6d49a.
//
// Solidity: function initialize(address[2] initAddr) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) Initialize(initAddr [2]common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.Initialize(&_AggregationSpotter.TransactOpts, initAddr)
}

// ProposeOperation is a paid mutator transaction binding the contract method 0xa20535eb.
//
// Solidity: function proposeOperation(bytes32 oracleOpTxId, (address,bytes4,bytes) opData) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) ProposeOperation(opts *bind.TransactOpts, oracleOpTxId [32]byte, opData AggregationSpotterOperationData) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "proposeOperation", oracleOpTxId, opData)
}

// ProposeOperation is a paid mutator transaction binding the contract method 0xa20535eb.
//
// Solidity: function proposeOperation(bytes32 oracleOpTxId, (address,bytes4,bytes) opData) returns()
func (_AggregationSpotter *AggregationSpotterSession) ProposeOperation(oracleOpTxId [32]byte, opData AggregationSpotterOperationData) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.ProposeOperation(&_AggregationSpotter.TransactOpts, oracleOpTxId, opData)
}

// ProposeOperation is a paid mutator transaction binding the contract method 0xa20535eb.
//
// Solidity: function proposeOperation(bytes32 oracleOpTxId, (address,bytes4,bytes) opData) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) ProposeOperation(oracleOpTxId [32]byte, opData AggregationSpotterOperationData) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.ProposeOperation(&_AggregationSpotter.TransactOpts, oracleOpTxId, opData)
}

// RemoveAllowedContract is a paid mutator transaction binding the contract method 0x9800fc16.
//
// Solidity: function removeAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) RemoveAllowedContract(opts *bind.TransactOpts, _contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "removeAllowedContract", _contract)
}

// RemoveAllowedContract is a paid mutator transaction binding the contract method 0x9800fc16.
//
// Solidity: function removeAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterSession) RemoveAllowedContract(_contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RemoveAllowedContract(&_AggregationSpotter.TransactOpts, _contract)
}

// RemoveAllowedContract is a paid mutator transaction binding the contract method 0x9800fc16.
//
// Solidity: function removeAllowedContract(address _contract) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) RemoveAllowedContract(_contract common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RemoveAllowedContract(&_AggregationSpotter.TransactOpts, _contract)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) RemoveKeeper(opts *bind.TransactOpts, keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "removeKeeper", keeper)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterSession) RemoveKeeper(keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RemoveKeeper(&_AggregationSpotter.TransactOpts, keeper)
}

// RemoveKeeper is a paid mutator transaction binding the contract method 0x14ae9f2e.
//
// Solidity: function removeKeeper(address keeper) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) RemoveKeeper(keeper common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RemoveKeeper(&_AggregationSpotter.TransactOpts, keeper)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationSpotter *AggregationSpotterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationSpotter *AggregationSpotterSession) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RenounceOwnership(&_AggregationSpotter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RenounceOwnership(&_AggregationSpotter.TransactOpts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RenounceRole(&_AggregationSpotter.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RenounceRole(&_AggregationSpotter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RevokeRole(&_AggregationSpotter.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.RevokeRole(&_AggregationSpotter.TransactOpts, role, account)
}

// SetConsensusTargetRate is a paid mutator transaction binding the contract method 0x4b3dc419.
//
// Solidity: function setConsensusTargetRate(uint256 rate) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) SetConsensusTargetRate(opts *bind.TransactOpts, rate *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "setConsensusTargetRate", rate)
}

// SetConsensusTargetRate is a paid mutator transaction binding the contract method 0x4b3dc419.
//
// Solidity: function setConsensusTargetRate(uint256 rate) returns()
func (_AggregationSpotter *AggregationSpotterSession) SetConsensusTargetRate(rate *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.SetConsensusTargetRate(&_AggregationSpotter.TransactOpts, rate)
}

// SetConsensusTargetRate is a paid mutator transaction binding the contract method 0x4b3dc419.
//
// Solidity: function setConsensusTargetRate(uint256 rate) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) SetConsensusTargetRate(rate *big.Int) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.SetConsensusTargetRate(&_AggregationSpotter.TransactOpts, rate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationSpotter *AggregationSpotterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.TransferOwnership(&_AggregationSpotter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.TransferOwnership(&_AggregationSpotter.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AggregationSpotter *AggregationSpotterTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AggregationSpotter *AggregationSpotterSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.UpgradeTo(&_AggregationSpotter.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.UpgradeTo(&_AggregationSpotter.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AggregationSpotter *AggregationSpotterTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AggregationSpotter.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AggregationSpotter *AggregationSpotterSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.UpgradeToAndCall(&_AggregationSpotter.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_AggregationSpotter *AggregationSpotterTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _AggregationSpotter.Contract.UpgradeToAndCall(&_AggregationSpotter.TransactOpts, newImplementation, data)
}

// AggregationSpotterAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the AggregationSpotter contract.
type AggregationSpotterAdminChangedIterator struct {
	Event *AggregationSpotterAdminChanged // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterAdminChanged)
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
		it.Event = new(AggregationSpotterAdminChanged)
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
func (it *AggregationSpotterAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterAdminChanged represents a AdminChanged event raised by the AggregationSpotter contract.
type AggregationSpotterAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*AggregationSpotterAdminChangedIterator, error) {

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterAdminChangedIterator{contract: _AggregationSpotter.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *AggregationSpotterAdminChanged) (event.Subscription, error) {

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterAdminChanged)
				if err := _AggregationSpotter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseAdminChanged(log types.Log) (*AggregationSpotterAdminChanged, error) {
	event := new(AggregationSpotterAdminChanged)
	if err := _AggregationSpotter.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the AggregationSpotter contract.
type AggregationSpotterBeaconUpgradedIterator struct {
	Event *AggregationSpotterBeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterBeaconUpgraded)
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
		it.Event = new(AggregationSpotterBeaconUpgraded)
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
func (it *AggregationSpotterBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterBeaconUpgraded represents a BeaconUpgraded event raised by the AggregationSpotter contract.
type AggregationSpotterBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*AggregationSpotterBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterBeaconUpgradedIterator{contract: _AggregationSpotter.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *AggregationSpotterBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterBeaconUpgraded)
				if err := _AggregationSpotter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseBeaconUpgraded(log types.Log) (*AggregationSpotterBeaconUpgraded, error) {
	event := new(AggregationSpotterBeaconUpgraded)
	if err := _AggregationSpotter.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the AggregationSpotter contract.
type AggregationSpotterInitializedIterator struct {
	Event *AggregationSpotterInitialized // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterInitialized)
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
		it.Event = new(AggregationSpotterInitialized)
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
func (it *AggregationSpotterInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterInitialized represents a Initialized event raised by the AggregationSpotter contract.
type AggregationSpotterInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterInitialized(opts *bind.FilterOpts) (*AggregationSpotterInitializedIterator, error) {

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterInitializedIterator{contract: _AggregationSpotter.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *AggregationSpotterInitialized) (event.Subscription, error) {

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterInitialized)
				if err := _AggregationSpotter.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseInitialized(log types.Log) (*AggregationSpotterInitialized, error) {
	event := new(AggregationSpotterInitialized)
	if err := _AggregationSpotter.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterNewOperationIterator is returned from FilterNewOperation and is used to iterate over the raw logs and unpacked data for NewOperation events raised by the AggregationSpotter contract.
type AggregationSpotterNewOperationIterator struct {
	Event *AggregationSpotterNewOperation // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterNewOperationIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterNewOperation)
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
		it.Event = new(AggregationSpotterNewOperation)
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
func (it *AggregationSpotterNewOperationIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterNewOperationIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterNewOperation represents a NewOperation event raised by the AggregationSpotter contract.
type AggregationSpotterNewOperation struct {
	OpHash           *big.Int
	Cont             common.Address
	FunctionSelector [4]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterNewOperation is a free log retrieval operation binding the contract event 0xd3aa55602696d0e59bb7b26de99d729ffd669d14f1589053886858dc935b55b9.
//
// Solidity: event NewOperation(uint256 opHash, address cont, bytes4 functionSelector)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterNewOperation(opts *bind.FilterOpts) (*AggregationSpotterNewOperationIterator, error) {

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "NewOperation")
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterNewOperationIterator{contract: _AggregationSpotter.contract, event: "NewOperation", logs: logs, sub: sub}, nil
}

// WatchNewOperation is a free log subscription operation binding the contract event 0xd3aa55602696d0e59bb7b26de99d729ffd669d14f1589053886858dc935b55b9.
//
// Solidity: event NewOperation(uint256 opHash, address cont, bytes4 functionSelector)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchNewOperation(opts *bind.WatchOpts, sink chan<- *AggregationSpotterNewOperation) (event.Subscription, error) {

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "NewOperation")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterNewOperation)
				if err := _AggregationSpotter.contract.UnpackLog(event, "NewOperation", log); err != nil {
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

// ParseNewOperation is a log parse operation binding the contract event 0xd3aa55602696d0e59bb7b26de99d729ffd669d14f1589053886858dc935b55b9.
//
// Solidity: event NewOperation(uint256 opHash, address cont, bytes4 functionSelector)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseNewOperation(log types.Log) (*AggregationSpotterNewOperation, error) {
	event := new(AggregationSpotterNewOperation)
	if err := _AggregationSpotter.contract.UnpackLog(event, "NewOperation", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AggregationSpotter contract.
type AggregationSpotterOwnershipTransferredIterator struct {
	Event *AggregationSpotterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterOwnershipTransferred)
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
		it.Event = new(AggregationSpotterOwnershipTransferred)
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
func (it *AggregationSpotterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterOwnershipTransferred represents a OwnershipTransferred event raised by the AggregationSpotter contract.
type AggregationSpotterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AggregationSpotterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterOwnershipTransferredIterator{contract: _AggregationSpotter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AggregationSpotterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterOwnershipTransferred)
				if err := _AggregationSpotter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseOwnershipTransferred(log types.Log) (*AggregationSpotterOwnershipTransferred, error) {
	event := new(AggregationSpotterOwnershipTransferred)
	if err := _AggregationSpotter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterProposalApprovedIterator is returned from FilterProposalApproved and is used to iterate over the raw logs and unpacked data for ProposalApproved events raised by the AggregationSpotter contract.
type AggregationSpotterProposalApprovedIterator struct {
	Event *AggregationSpotterProposalApproved // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterProposalApprovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterProposalApproved)
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
		it.Event = new(AggregationSpotterProposalApproved)
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
func (it *AggregationSpotterProposalApprovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterProposalApprovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterProposalApproved represents a ProposalApproved event raised by the AggregationSpotter contract.
type AggregationSpotterProposalApproved struct {
	OpHash *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalApproved is a free log retrieval operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterProposalApproved(opts *bind.FilterOpts) (*AggregationSpotterProposalApprovedIterator, error) {

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "ProposalApproved")
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterProposalApprovedIterator{contract: _AggregationSpotter.contract, event: "ProposalApproved", logs: logs, sub: sub}, nil
}

// WatchProposalApproved is a free log subscription operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchProposalApproved(opts *bind.WatchOpts, sink chan<- *AggregationSpotterProposalApproved) (event.Subscription, error) {

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "ProposalApproved")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterProposalApproved)
				if err := _AggregationSpotter.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
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

// ParseProposalApproved is a log parse operation binding the contract event 0x28ec9e38ba73636ceb2f6c1574136f83bd46284a3c74734b711bf45e12f8d929.
//
// Solidity: event ProposalApproved(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseProposalApproved(log types.Log) (*AggregationSpotterProposalApproved, error) {
	event := new(AggregationSpotterProposalApproved)
	if err := _AggregationSpotter.contract.UnpackLog(event, "ProposalApproved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterProposalExecutedIterator is returned from FilterProposalExecuted and is used to iterate over the raw logs and unpacked data for ProposalExecuted events raised by the AggregationSpotter contract.
type AggregationSpotterProposalExecutedIterator struct {
	Event *AggregationSpotterProposalExecuted // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterProposalExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterProposalExecuted)
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
		it.Event = new(AggregationSpotterProposalExecuted)
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
func (it *AggregationSpotterProposalExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterProposalExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterProposalExecuted represents a ProposalExecuted event raised by the AggregationSpotter contract.
type AggregationSpotterProposalExecuted struct {
	OpHash *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProposalExecuted is a free log retrieval operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterProposalExecuted(opts *bind.FilterOpts) (*AggregationSpotterProposalExecutedIterator, error) {

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterProposalExecutedIterator{contract: _AggregationSpotter.contract, event: "ProposalExecuted", logs: logs, sub: sub}, nil
}

// WatchProposalExecuted is a free log subscription operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchProposalExecuted(opts *bind.WatchOpts, sink chan<- *AggregationSpotterProposalExecuted) (event.Subscription, error) {

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "ProposalExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterProposalExecuted)
				if err := _AggregationSpotter.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
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

// ParseProposalExecuted is a log parse operation binding the contract event 0x712ae1383f79ac853f8d882153778e0260ef8f03b504e2866e0593e04d2b291f.
//
// Solidity: event ProposalExecuted(uint256 opHash)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseProposalExecuted(log types.Log) (*AggregationSpotterProposalExecuted, error) {
	event := new(AggregationSpotterProposalExecuted)
	if err := _AggregationSpotter.contract.UnpackLog(event, "ProposalExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the AggregationSpotter contract.
type AggregationSpotterRoleAdminChangedIterator struct {
	Event *AggregationSpotterRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterRoleAdminChanged)
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
		it.Event = new(AggregationSpotterRoleAdminChanged)
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
func (it *AggregationSpotterRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterRoleAdminChanged represents a RoleAdminChanged event raised by the AggregationSpotter contract.
type AggregationSpotterRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*AggregationSpotterRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterRoleAdminChangedIterator{contract: _AggregationSpotter.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *AggregationSpotterRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterRoleAdminChanged)
				if err := _AggregationSpotter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseRoleAdminChanged(log types.Log) (*AggregationSpotterRoleAdminChanged, error) {
	event := new(AggregationSpotterRoleAdminChanged)
	if err := _AggregationSpotter.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the AggregationSpotter contract.
type AggregationSpotterRoleGrantedIterator struct {
	Event *AggregationSpotterRoleGranted // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterRoleGranted)
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
		it.Event = new(AggregationSpotterRoleGranted)
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
func (it *AggregationSpotterRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterRoleGranted represents a RoleGranted event raised by the AggregationSpotter contract.
type AggregationSpotterRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AggregationSpotterRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterRoleGrantedIterator{contract: _AggregationSpotter.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *AggregationSpotterRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterRoleGranted)
				if err := _AggregationSpotter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseRoleGranted(log types.Log) (*AggregationSpotterRoleGranted, error) {
	event := new(AggregationSpotterRoleGranted)
	if err := _AggregationSpotter.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the AggregationSpotter contract.
type AggregationSpotterRoleRevokedIterator struct {
	Event *AggregationSpotterRoleRevoked // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterRoleRevoked)
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
		it.Event = new(AggregationSpotterRoleRevoked)
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
func (it *AggregationSpotterRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterRoleRevoked represents a RoleRevoked event raised by the AggregationSpotter contract.
type AggregationSpotterRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*AggregationSpotterRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterRoleRevokedIterator{contract: _AggregationSpotter.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *AggregationSpotterRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterRoleRevoked)
				if err := _AggregationSpotter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_AggregationSpotter *AggregationSpotterFilterer) ParseRoleRevoked(log types.Log) (*AggregationSpotterRoleRevoked, error) {
	event := new(AggregationSpotterRoleRevoked)
	if err := _AggregationSpotter.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AggregationSpotterUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the AggregationSpotter contract.
type AggregationSpotterUpgradedIterator struct {
	Event *AggregationSpotterUpgraded // Event containing the contract specifics and raw log

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
func (it *AggregationSpotterUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AggregationSpotterUpgraded)
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
		it.Event = new(AggregationSpotterUpgraded)
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
func (it *AggregationSpotterUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AggregationSpotterUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AggregationSpotterUpgraded represents a Upgraded event raised by the AggregationSpotter contract.
type AggregationSpotterUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AggregationSpotter *AggregationSpotterFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*AggregationSpotterUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AggregationSpotter.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &AggregationSpotterUpgradedIterator{contract: _AggregationSpotter.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_AggregationSpotter *AggregationSpotterFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *AggregationSpotterUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _AggregationSpotter.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AggregationSpotterUpgraded)
				if err := _AggregationSpotter.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_AggregationSpotter *AggregationSpotterFilterer) ParseUpgraded(log types.Log) (*AggregationSpotterUpgraded, error) {
	event := new(AggregationSpotterUpgraded)
	if err := _AggregationSpotter.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
