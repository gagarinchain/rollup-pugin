// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package gateway

import (
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
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// GatewayPending is an auto generated low-level Go binding around an user-defined struct.
type GatewayPending struct {
	Amount      *big.Int
	BlockNumber *big.Int
}

// GatewayABI is the input ABI used to generate the binding from.
const GatewayABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"delta\",\"type\":\"uint256\"}],\"name\":\"DepositCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositConfirmed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"DepositReturned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"eth\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"gagarin\",\"type\":\"address\"}],\"name\":\"Registered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"eth\",\"type\":\"address\"}],\"name\":\"Unregistered\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"governance\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingTTL\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rollupsManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governance\",\"type\":\"address\"}],\"name\":\"setGovernance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_manager\",\"type\":\"address\"}],\"name\":\"setRollupManager\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"lockTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newTTL\",\"type\":\"uint32\"}],\"name\":\"changePendingTTL\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gagarinOwner\",\"type\":\"address\"}],\"name\":\"confirmDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gagarinOwner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"returnDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"cancelDeposit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gAddress\",\"type\":\"address\"}],\"name\":\"getPending\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structGateway.Pending\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"gAddress\",\"type\":\"address\"}],\"name\":\"register\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unregister\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Gateway is an auto generated Go binding around an Ethereum contract.
type Gateway struct {
	GatewayCaller     // Read-only binding to the contract
	GatewayTransactor // Write-only binding to the contract
	GatewayFilterer   // Log filterer for contract events
}

// GatewayCaller is an auto generated read-only Go binding around an Ethereum contract.
type GatewayCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GatewayTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewayFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GatewayFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GatewaySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GatewaySession struct {
	Contract     *Gateway          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// GatewayCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GatewayCallerSession struct {
	Contract *GatewayCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// GatewayTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GatewayTransactorSession struct {
	Contract     *GatewayTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// GatewayRaw is an auto generated low-level Go binding around an Ethereum contract.
type GatewayRaw struct {
	Contract *Gateway // Generic contract binding to access the raw methods on
}

// GatewayCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GatewayCallerRaw struct {
	Contract *GatewayCaller // Generic read-only contract binding to access the raw methods on
}

// GatewayTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GatewayTransactorRaw struct {
	Contract *GatewayTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGateway creates a new instance of Gateway, bound to a specific deployed contract.
func NewGateway(address common.Address, backend bind.ContractBackend) (*Gateway, error) {
	contract, err := bindGateway(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Gateway{GatewayCaller: GatewayCaller{contract: contract}, GatewayTransactor: GatewayTransactor{contract: contract}, GatewayFilterer: GatewayFilterer{contract: contract}}, nil
}

// NewGatewayCaller creates a new read-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayCaller(address common.Address, caller bind.ContractCaller) (*GatewayCaller, error) {
	contract, err := bindGateway(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayCaller{contract: contract}, nil
}

// NewGatewayTransactor creates a new write-only instance of Gateway, bound to a specific deployed contract.
func NewGatewayTransactor(address common.Address, transactor bind.ContractTransactor) (*GatewayTransactor, error) {
	contract, err := bindGateway(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GatewayTransactor{contract: contract}, nil
}

// NewGatewayFilterer creates a new log filterer instance of Gateway, bound to a specific deployed contract.
func NewGatewayFilterer(address common.Address, filterer bind.ContractFilterer) (*GatewayFilterer, error) {
	contract, err := bindGateway(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GatewayFilterer{contract: contract}, nil
}

// bindGateway binds a generic wrapper to an already deployed contract.
func bindGateway(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GatewayABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.GatewayCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.GatewayTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Gateway *GatewayCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Gateway.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Gateway *GatewayTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Gateway *GatewayTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Gateway.Contract.contract.Transact(opts, method, params...)
}

// GetPending is a free data retrieval call binding the contract method 0xd086c254.
//
// Solidity: function getPending(address gAddress) view returns((uint256,uint256))
func (_Gateway *GatewayCaller) GetPending(opts *bind.CallOpts, gAddress common.Address) (GatewayPending, error) {
	var (
		ret0 = new(GatewayPending)
	)
	out := ret0
	err := _Gateway.contract.Call(opts, out, "getPending", gAddress)
	return *ret0, err
}

// GetPending is a free data retrieval call binding the contract method 0xd086c254.
//
// Solidity: function getPending(address gAddress) view returns((uint256,uint256))
func (_Gateway *GatewaySession) GetPending(gAddress common.Address) (GatewayPending, error) {
	return _Gateway.Contract.GetPending(&_Gateway.CallOpts, gAddress)
}

// GetPending is a free data retrieval call binding the contract method 0xd086c254.
//
// Solidity: function getPending(address gAddress) view returns((uint256,uint256))
func (_Gateway *GatewayCallerSession) GetPending(gAddress common.Address) (GatewayPending, error) {
	return _Gateway.Contract.GetPending(&_Gateway.CallOpts, gAddress)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Gateway *GatewayCaller) Governance(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gateway.contract.Call(opts, out, "governance")
	return *ret0, err
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Gateway *GatewaySession) Governance() (common.Address, error) {
	return _Gateway.Contract.Governance(&_Gateway.CallOpts)
}

// Governance is a free data retrieval call binding the contract method 0x5aa6e675.
//
// Solidity: function governance() view returns(address)
func (_Gateway *GatewayCallerSession) Governance() (common.Address, error) {
	return _Gateway.Contract.Governance(&_Gateway.CallOpts)
}

// PendingTTL is a free data retrieval call binding the contract method 0xd80dbae2.
//
// Solidity: function pendingTTL() view returns(uint32)
func (_Gateway *GatewayCaller) PendingTTL(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Gateway.contract.Call(opts, out, "pendingTTL")
	return *ret0, err
}

// PendingTTL is a free data retrieval call binding the contract method 0xd80dbae2.
//
// Solidity: function pendingTTL() view returns(uint32)
func (_Gateway *GatewaySession) PendingTTL() (uint32, error) {
	return _Gateway.Contract.PendingTTL(&_Gateway.CallOpts)
}

// PendingTTL is a free data retrieval call binding the contract method 0xd80dbae2.
//
// Solidity: function pendingTTL() view returns(uint32)
func (_Gateway *GatewayCallerSession) PendingTTL() (uint32, error) {
	return _Gateway.Contract.PendingTTL(&_Gateway.CallOpts)
}

// RollupsManager is a free data retrieval call binding the contract method 0xe612262f.
//
// Solidity: function rollupsManager() view returns(address)
func (_Gateway *GatewayCaller) RollupsManager(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gateway.contract.Call(opts, out, "rollupsManager")
	return *ret0, err
}

// RollupsManager is a free data retrieval call binding the contract method 0xe612262f.
//
// Solidity: function rollupsManager() view returns(address)
func (_Gateway *GatewaySession) RollupsManager() (common.Address, error) {
	return _Gateway.Contract.RollupsManager(&_Gateway.CallOpts)
}

// RollupsManager is a free data retrieval call binding the contract method 0xe612262f.
//
// Solidity: function rollupsManager() view returns(address)
func (_Gateway *GatewayCallerSession) RollupsManager() (common.Address, error) {
	return _Gateway.Contract.RollupsManager(&_Gateway.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Gateway *GatewayCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Gateway.contract.Call(opts, out, "token")
	return *ret0, err
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Gateway *GatewaySession) Token() (common.Address, error) {
	return _Gateway.Contract.Token(&_Gateway.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Gateway *GatewayCallerSession) Token() (common.Address, error) {
	return _Gateway.Contract.Token(&_Gateway.CallOpts)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x3d561602.
//
// Solidity: function cancelDeposit() returns()
func (_Gateway *GatewayTransactor) CancelDeposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "cancelDeposit")
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x3d561602.
//
// Solidity: function cancelDeposit() returns()
func (_Gateway *GatewaySession) CancelDeposit() (*types.Transaction, error) {
	return _Gateway.Contract.CancelDeposit(&_Gateway.TransactOpts)
}

// CancelDeposit is a paid mutator transaction binding the contract method 0x3d561602.
//
// Solidity: function cancelDeposit() returns()
func (_Gateway *GatewayTransactorSession) CancelDeposit() (*types.Transaction, error) {
	return _Gateway.Contract.CancelDeposit(&_Gateway.TransactOpts)
}

// ChangePendingTTL is a paid mutator transaction binding the contract method 0x0d9e2d04.
//
// Solidity: function changePendingTTL(uint32 newTTL) returns()
func (_Gateway *GatewayTransactor) ChangePendingTTL(opts *bind.TransactOpts, newTTL uint32) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "changePendingTTL", newTTL)
}

// ChangePendingTTL is a paid mutator transaction binding the contract method 0x0d9e2d04.
//
// Solidity: function changePendingTTL(uint32 newTTL) returns()
func (_Gateway *GatewaySession) ChangePendingTTL(newTTL uint32) (*types.Transaction, error) {
	return _Gateway.Contract.ChangePendingTTL(&_Gateway.TransactOpts, newTTL)
}

// ChangePendingTTL is a paid mutator transaction binding the contract method 0x0d9e2d04.
//
// Solidity: function changePendingTTL(uint32 newTTL) returns()
func (_Gateway *GatewayTransactorSession) ChangePendingTTL(newTTL uint32) (*types.Transaction, error) {
	return _Gateway.Contract.ChangePendingTTL(&_Gateway.TransactOpts, newTTL)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x6452fca1.
//
// Solidity: function confirmDeposit(address gagarinOwner) returns()
func (_Gateway *GatewayTransactor) ConfirmDeposit(opts *bind.TransactOpts, gagarinOwner common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "confirmDeposit", gagarinOwner)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x6452fca1.
//
// Solidity: function confirmDeposit(address gagarinOwner) returns()
func (_Gateway *GatewaySession) ConfirmDeposit(gagarinOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.ConfirmDeposit(&_Gateway.TransactOpts, gagarinOwner)
}

// ConfirmDeposit is a paid mutator transaction binding the contract method 0x6452fca1.
//
// Solidity: function confirmDeposit(address gagarinOwner) returns()
func (_Gateway *GatewayTransactorSession) ConfirmDeposit(gagarinOwner common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.ConfirmDeposit(&_Gateway.TransactOpts, gagarinOwner)
}

// LockTokens is a paid mutator transaction binding the contract method 0x6e27d889.
//
// Solidity: function lockTokens(uint256 _amount) returns()
func (_Gateway *GatewayTransactor) LockTokens(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "lockTokens", _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x6e27d889.
//
// Solidity: function lockTokens(uint256 _amount) returns()
func (_Gateway *GatewaySession) LockTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.LockTokens(&_Gateway.TransactOpts, _amount)
}

// LockTokens is a paid mutator transaction binding the contract method 0x6e27d889.
//
// Solidity: function lockTokens(uint256 _amount) returns()
func (_Gateway *GatewayTransactorSession) LockTokens(_amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.LockTokens(&_Gateway.TransactOpts, _amount)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address gAddress) returns()
func (_Gateway *GatewayTransactor) Register(opts *bind.TransactOpts, gAddress common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "register", gAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address gAddress) returns()
func (_Gateway *GatewaySession) Register(gAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Register(&_Gateway.TransactOpts, gAddress)
}

// Register is a paid mutator transaction binding the contract method 0x4420e486.
//
// Solidity: function register(address gAddress) returns()
func (_Gateway *GatewayTransactorSession) Register(gAddress common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.Register(&_Gateway.TransactOpts, gAddress)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0xd29709fb.
//
// Solidity: function returnDeposit(address gagarinOwner, uint256 _amount) returns()
func (_Gateway *GatewayTransactor) ReturnDeposit(opts *bind.TransactOpts, gagarinOwner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "returnDeposit", gagarinOwner, _amount)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0xd29709fb.
//
// Solidity: function returnDeposit(address gagarinOwner, uint256 _amount) returns()
func (_Gateway *GatewaySession) ReturnDeposit(gagarinOwner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.ReturnDeposit(&_Gateway.TransactOpts, gagarinOwner, _amount)
}

// ReturnDeposit is a paid mutator transaction binding the contract method 0xd29709fb.
//
// Solidity: function returnDeposit(address gagarinOwner, uint256 _amount) returns()
func (_Gateway *GatewayTransactorSession) ReturnDeposit(gagarinOwner common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _Gateway.Contract.ReturnDeposit(&_Gateway.TransactOpts, gagarinOwner, _amount)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Gateway *GatewayTransactor) SetGovernance(opts *bind.TransactOpts, _governance common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setGovernance", _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Gateway *GatewaySession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetGovernance(&_Gateway.TransactOpts, _governance)
}

// SetGovernance is a paid mutator transaction binding the contract method 0xab033ea9.
//
// Solidity: function setGovernance(address _governance) returns()
func (_Gateway *GatewayTransactorSession) SetGovernance(_governance common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetGovernance(&_Gateway.TransactOpts, _governance)
}

// SetRollupManager is a paid mutator transaction binding the contract method 0x2ab4a54f.
//
// Solidity: function setRollupManager(address _manager) returns()
func (_Gateway *GatewayTransactor) SetRollupManager(opts *bind.TransactOpts, _manager common.Address) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "setRollupManager", _manager)
}

// SetRollupManager is a paid mutator transaction binding the contract method 0x2ab4a54f.
//
// Solidity: function setRollupManager(address _manager) returns()
func (_Gateway *GatewaySession) SetRollupManager(_manager common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetRollupManager(&_Gateway.TransactOpts, _manager)
}

// SetRollupManager is a paid mutator transaction binding the contract method 0x2ab4a54f.
//
// Solidity: function setRollupManager(address _manager) returns()
func (_Gateway *GatewayTransactorSession) SetRollupManager(_manager common.Address) (*types.Transaction, error) {
	return _Gateway.Contract.SetRollupManager(&_Gateway.TransactOpts, _manager)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Gateway *GatewayTransactor) Unregister(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Gateway.contract.Transact(opts, "unregister")
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Gateway *GatewaySession) Unregister() (*types.Transaction, error) {
	return _Gateway.Contract.Unregister(&_Gateway.TransactOpts)
}

// Unregister is a paid mutator transaction binding the contract method 0xe79a198f.
//
// Solidity: function unregister() returns()
func (_Gateway *GatewayTransactorSession) Unregister() (*types.Transaction, error) {
	return _Gateway.Contract.Unregister(&_Gateway.TransactOpts)
}

// GatewayDepositCancelledIterator is returned from FilterDepositCancelled and is used to iterate over the raw logs and unpacked data for DepositCancelled events raised by the Gateway contract.
type GatewayDepositCancelledIterator struct {
	Event *GatewayDepositCancelled // Event containing the contract specifics and raw log

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
func (it *GatewayDepositCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDepositCancelled)
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
		it.Event = new(GatewayDepositCancelled)
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
func (it *GatewayDepositCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDepositCancelled represents a DepositCancelled event raised by the Gateway contract.
type GatewayDepositCancelled struct {
	Sender common.Address
	Delta  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositCancelled is a free log retrieval operation binding the contract event 0x1578da51d7693d10ae70cf7f1026491f713689b467be8e1a73b5e24656dea729.
//
// Solidity: event DepositCancelled(address sender, uint256 delta)
func (_Gateway *GatewayFilterer) FilterDepositCancelled(opts *bind.FilterOpts) (*GatewayDepositCancelledIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "DepositCancelled")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositCancelledIterator{contract: _Gateway.contract, event: "DepositCancelled", logs: logs, sub: sub}, nil
}

// WatchDepositCancelled is a free log subscription operation binding the contract event 0x1578da51d7693d10ae70cf7f1026491f713689b467be8e1a73b5e24656dea729.
//
// Solidity: event DepositCancelled(address sender, uint256 delta)
func (_Gateway *GatewayFilterer) WatchDepositCancelled(opts *bind.WatchOpts, sink chan<- *GatewayDepositCancelled) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "DepositCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDepositCancelled)
				if err := _Gateway.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
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

// ParseDepositCancelled is a log parse operation binding the contract event 0x1578da51d7693d10ae70cf7f1026491f713689b467be8e1a73b5e24656dea729.
//
// Solidity: event DepositCancelled(address sender, uint256 delta)
func (_Gateway *GatewayFilterer) ParseDepositCancelled(log types.Log) (*GatewayDepositCancelled, error) {
	event := new(GatewayDepositCancelled)
	if err := _Gateway.contract.UnpackLog(event, "DepositCancelled", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GatewayDepositConfirmedIterator is returned from FilterDepositConfirmed and is used to iterate over the raw logs and unpacked data for DepositConfirmed events raised by the Gateway contract.
type GatewayDepositConfirmedIterator struct {
	Event *GatewayDepositConfirmed // Event containing the contract specifics and raw log

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
func (it *GatewayDepositConfirmedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDepositConfirmed)
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
		it.Event = new(GatewayDepositConfirmed)
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
func (it *GatewayDepositConfirmedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositConfirmedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDepositConfirmed represents a DepositConfirmed event raised by the Gateway contract.
type GatewayDepositConfirmed struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositConfirmed is a free log retrieval operation binding the contract event 0xf29f4890dcce1aaf31cc3e0c5f3a777653c6edd5d9cc867a55bc1953e0fb7ee6.
//
// Solidity: event DepositConfirmed(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) FilterDepositConfirmed(opts *bind.FilterOpts) (*GatewayDepositConfirmedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "DepositConfirmed")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositConfirmedIterator{contract: _Gateway.contract, event: "DepositConfirmed", logs: logs, sub: sub}, nil
}

// WatchDepositConfirmed is a free log subscription operation binding the contract event 0xf29f4890dcce1aaf31cc3e0c5f3a777653c6edd5d9cc867a55bc1953e0fb7ee6.
//
// Solidity: event DepositConfirmed(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) WatchDepositConfirmed(opts *bind.WatchOpts, sink chan<- *GatewayDepositConfirmed) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "DepositConfirmed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDepositConfirmed)
				if err := _Gateway.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
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

// ParseDepositConfirmed is a log parse operation binding the contract event 0xf29f4890dcce1aaf31cc3e0c5f3a777653c6edd5d9cc867a55bc1953e0fb7ee6.
//
// Solidity: event DepositConfirmed(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) ParseDepositConfirmed(log types.Log) (*GatewayDepositConfirmed, error) {
	event := new(GatewayDepositConfirmed)
	if err := _Gateway.contract.UnpackLog(event, "DepositConfirmed", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GatewayDepositReturnedIterator is returned from FilterDepositReturned and is used to iterate over the raw logs and unpacked data for DepositReturned events raised by the Gateway contract.
type GatewayDepositReturnedIterator struct {
	Event *GatewayDepositReturned // Event containing the contract specifics and raw log

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
func (it *GatewayDepositReturnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayDepositReturned)
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
		it.Event = new(GatewayDepositReturned)
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
func (it *GatewayDepositReturnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayDepositReturnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayDepositReturned represents a DepositReturned event raised by the Gateway contract.
type GatewayDepositReturned struct {
	Sender common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterDepositReturned is a free log retrieval operation binding the contract event 0xa09244c122a83074363635dc89ce0098166c2f74827960ff64a7c8e6ad1c57a2.
//
// Solidity: event DepositReturned(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) FilterDepositReturned(opts *bind.FilterOpts) (*GatewayDepositReturnedIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "DepositReturned")
	if err != nil {
		return nil, err
	}
	return &GatewayDepositReturnedIterator{contract: _Gateway.contract, event: "DepositReturned", logs: logs, sub: sub}, nil
}

// WatchDepositReturned is a free log subscription operation binding the contract event 0xa09244c122a83074363635dc89ce0098166c2f74827960ff64a7c8e6ad1c57a2.
//
// Solidity: event DepositReturned(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) WatchDepositReturned(opts *bind.WatchOpts, sink chan<- *GatewayDepositReturned) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "DepositReturned")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayDepositReturned)
				if err := _Gateway.contract.UnpackLog(event, "DepositReturned", log); err != nil {
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

// ParseDepositReturned is a log parse operation binding the contract event 0xa09244c122a83074363635dc89ce0098166c2f74827960ff64a7c8e6ad1c57a2.
//
// Solidity: event DepositReturned(address sender, uint256 amount)
func (_Gateway *GatewayFilterer) ParseDepositReturned(log types.Log) (*GatewayDepositReturned, error) {
	event := new(GatewayDepositReturned)
	if err := _Gateway.contract.UnpackLog(event, "DepositReturned", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GatewayRegisteredIterator is returned from FilterRegistered and is used to iterate over the raw logs and unpacked data for Registered events raised by the Gateway contract.
type GatewayRegisteredIterator struct {
	Event *GatewayRegistered // Event containing the contract specifics and raw log

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
func (it *GatewayRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayRegistered)
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
		it.Event = new(GatewayRegistered)
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
func (it *GatewayRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayRegistered represents a Registered event raised by the Gateway contract.
type GatewayRegistered struct {
	Eth     common.Address
	Gagarin common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRegistered is a free log retrieval operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address eth, address gagarin)
func (_Gateway *GatewayFilterer) FilterRegistered(opts *bind.FilterOpts) (*GatewayRegisteredIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return &GatewayRegisteredIterator{contract: _Gateway.contract, event: "Registered", logs: logs, sub: sub}, nil
}

// WatchRegistered is a free log subscription operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address eth, address gagarin)
func (_Gateway *GatewayFilterer) WatchRegistered(opts *bind.WatchOpts, sink chan<- *GatewayRegistered) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Registered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayRegistered)
				if err := _Gateway.contract.UnpackLog(event, "Registered", log); err != nil {
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

// ParseRegistered is a log parse operation binding the contract event 0x0a31ee9d46a828884b81003c8498156ea6aa15b9b54bdd0ef0b533d9eba57e55.
//
// Solidity: event Registered(address eth, address gagarin)
func (_Gateway *GatewayFilterer) ParseRegistered(log types.Log) (*GatewayRegistered, error) {
	event := new(GatewayRegistered)
	if err := _Gateway.contract.UnpackLog(event, "Registered", log); err != nil {
		return nil, err
	}
	return event, nil
}

// GatewayUnregisteredIterator is returned from FilterUnregistered and is used to iterate over the raw logs and unpacked data for Unregistered events raised by the Gateway contract.
type GatewayUnregisteredIterator struct {
	Event *GatewayUnregistered // Event containing the contract specifics and raw log

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
func (it *GatewayUnregisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(GatewayUnregistered)
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
		it.Event = new(GatewayUnregistered)
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
func (it *GatewayUnregisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *GatewayUnregisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// GatewayUnregistered represents a Unregistered event raised by the Gateway contract.
type GatewayUnregistered struct {
	Eth common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnregistered is a free log retrieval operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address eth)
func (_Gateway *GatewayFilterer) FilterUnregistered(opts *bind.FilterOpts) (*GatewayUnregisteredIterator, error) {

	logs, sub, err := _Gateway.contract.FilterLogs(opts, "Unregistered")
	if err != nil {
		return nil, err
	}
	return &GatewayUnregisteredIterator{contract: _Gateway.contract, event: "Unregistered", logs: logs, sub: sub}, nil
}

// WatchUnregistered is a free log subscription operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address eth)
func (_Gateway *GatewayFilterer) WatchUnregistered(opts *bind.WatchOpts, sink chan<- *GatewayUnregistered) (event.Subscription, error) {

	logs, sub, err := _Gateway.contract.WatchLogs(opts, "Unregistered")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(GatewayUnregistered)
				if err := _Gateway.contract.UnpackLog(event, "Unregistered", log); err != nil {
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

// ParseUnregistered is a log parse operation binding the contract event 0x75cd6de711483e11488a1cd9b66172abccb9e5c19572f92015a7880f0c8c0edc.
//
// Solidity: event Unregistered(address eth)
func (_Gateway *GatewayFilterer) ParseUnregistered(log types.Log) (*GatewayUnregistered, error) {
	event := new(GatewayUnregistered)
	if err := _Gateway.contract.UnpackLog(event, "Unregistered", log); err != nil {
		return nil, err
	}
	return event, nil
}
