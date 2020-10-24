// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rollup

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

// RollupABI is the input ABI used to generate the binding from.
const RollupABI = "[{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"c\",\"type\":\"address[]\"},{\"internalType\":\"address\",\"name\":\"_gateway\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"name\":\"BalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockCount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"height\",\"type\":\"uint32\"}],\"name\":\"TopHeightUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tranIndex\",\"type\":\"uint256\"}],\"name\":\"Transaction\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"committee\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopHeight\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"header\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"rollup\",\"type\":\"bytes\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"}],\"name\":\"addBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Rollup is an auto generated Go binding around an Ethereum contract.
type Rollup struct {
	RollupCaller     // Read-only binding to the contract
	RollupTransactor // Write-only binding to the contract
	RollupFilterer   // Log filterer for contract events
}

// RollupCaller is an auto generated read-only Go binding around an Ethereum contract.
type RollupCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RollupTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RollupFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RollupSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RollupSession struct {
	Contract     *Rollup           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RollupCallerSession struct {
	Contract *RollupCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// RollupTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RollupTransactorSession struct {
	Contract     *RollupTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RollupRaw is an auto generated low-level Go binding around an Ethereum contract.
type RollupRaw struct {
	Contract *Rollup // Generic contract binding to access the raw methods on
}

// RollupCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RollupCallerRaw struct {
	Contract *RollupCaller // Generic read-only contract binding to access the raw methods on
}

// RollupTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RollupTransactorRaw struct {
	Contract *RollupTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRollup creates a new instance of Rollup, bound to a specific deployed contract.
func NewRollup(address common.Address, backend bind.ContractBackend) (*Rollup, error) {
	contract, err := bindRollup(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Rollup{RollupCaller: RollupCaller{contract: contract}, RollupTransactor: RollupTransactor{contract: contract}, RollupFilterer: RollupFilterer{contract: contract}}, nil
}

// NewRollupCaller creates a new read-only instance of Rollup, bound to a specific deployed contract.
func NewRollupCaller(address common.Address, caller bind.ContractCaller) (*RollupCaller, error) {
	contract, err := bindRollup(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RollupCaller{contract: contract}, nil
}

// NewRollupTransactor creates a new write-only instance of Rollup, bound to a specific deployed contract.
func NewRollupTransactor(address common.Address, transactor bind.ContractTransactor) (*RollupTransactor, error) {
	contract, err := bindRollup(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RollupTransactor{contract: contract}, nil
}

// NewRollupFilterer creates a new log filterer instance of Rollup, bound to a specific deployed contract.
func NewRollupFilterer(address common.Address, filterer bind.ContractFilterer) (*RollupFilterer, error) {
	contract, err := bindRollup(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RollupFilterer{contract: contract}, nil
}

// bindRollup binds a generic wrapper to an already deployed contract.
func bindRollup(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RollupABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.RollupCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.RollupTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Rollup *RollupCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Rollup.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Rollup *RollupTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Rollup *RollupTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Rollup.Contract.contract.Transact(opts, method, params...)
}

// Committee is a free data retrieval call binding the contract method 0xafe7fcf4.
//
// Solidity: function committee(uint256 ) view returns(address)
func (_Rollup *RollupCaller) Committee(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "committee", arg0)
	return *ret0, err
}

// Committee is a free data retrieval call binding the contract method 0xafe7fcf4.
//
// Solidity: function committee(uint256 ) view returns(address)
func (_Rollup *RollupSession) Committee(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Committee(&_Rollup.CallOpts, arg0)
}

// Committee is a free data retrieval call binding the contract method 0xafe7fcf4.
//
// Solidity: function committee(uint256 ) view returns(address)
func (_Rollup *RollupCallerSession) Committee(arg0 *big.Int) (common.Address, error) {
	return _Rollup.Contract.Committee(&_Rollup.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address owner) view returns(uint256)
func (_Rollup *RollupCaller) GetBalance(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getBalance", owner)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address owner) view returns(uint256)
func (_Rollup *RollupSession) GetBalance(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.GetBalance(&_Rollup.CallOpts, owner)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address owner) view returns(uint256)
func (_Rollup *RollupCallerSession) GetBalance(owner common.Address) (*big.Int, error) {
	return _Rollup.Contract.GetBalance(&_Rollup.CallOpts, owner)
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() view returns(uint32)
func (_Rollup *RollupCaller) GetTopHeight(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _Rollup.contract.Call(opts, out, "getTopHeight")
	return *ret0, err
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() view returns(uint32)
func (_Rollup *RollupSession) GetTopHeight() (uint32, error) {
	return _Rollup.Contract.GetTopHeight(&_Rollup.CallOpts)
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() view returns(uint32)
func (_Rollup *RollupCallerSession) GetTopHeight() (uint32, error) {
	return _Rollup.Contract.GetTopHeight(&_Rollup.CallOpts)
}

// AddBlock is a paid mutator transaction binding the contract method 0x8b39ed25.
//
// Solidity: function addBlock(bytes header, bytes rollup, bytes signature) returns()
func (_Rollup *RollupTransactor) AddBlock(opts *bind.TransactOpts, header []byte, rollup []byte, signature []byte) (*types.Transaction, error) {
	return _Rollup.contract.Transact(opts, "addBlock", header, rollup, signature)
}

// AddBlock is a paid mutator transaction binding the contract method 0x8b39ed25.
//
// Solidity: function addBlock(bytes header, bytes rollup, bytes signature) returns()
func (_Rollup *RollupSession) AddBlock(header []byte, rollup []byte, signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.AddBlock(&_Rollup.TransactOpts, header, rollup, signature)
}

// AddBlock is a paid mutator transaction binding the contract method 0x8b39ed25.
//
// Solidity: function addBlock(bytes header, bytes rollup, bytes signature) returns()
func (_Rollup *RollupTransactorSession) AddBlock(header []byte, rollup []byte, signature []byte) (*types.Transaction, error) {
	return _Rollup.Contract.AddBlock(&_Rollup.TransactOpts, header, rollup, signature)
}

// RollupBalanceChangedIterator is returned from FilterBalanceChanged and is used to iterate over the raw logs and unpacked data for BalanceChanged events raised by the Rollup contract.
type RollupBalanceChangedIterator struct {
	Event *RollupBalanceChanged // Event containing the contract specifics and raw log

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
func (it *RollupBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupBalanceChanged)
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
		it.Event = new(RollupBalanceChanged)
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
func (it *RollupBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupBalanceChanged represents a BalanceChanged event raised by the Rollup contract.
type RollupBalanceChanged struct {
	Owner   common.Address
	Balance *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceChanged is a free log retrieval operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address owner, uint256 balance)
func (_Rollup *RollupFilterer) FilterBalanceChanged(opts *bind.FilterOpts) (*RollupBalanceChangedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return &RollupBalanceChangedIterator{contract: _Rollup.contract, event: "BalanceChanged", logs: logs, sub: sub}, nil
}

// WatchBalanceChanged is a free log subscription operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address owner, uint256 balance)
func (_Rollup *RollupFilterer) WatchBalanceChanged(opts *bind.WatchOpts, sink chan<- *RollupBalanceChanged) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupBalanceChanged)
				if err := _Rollup.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
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

// ParseBalanceChanged is a log parse operation binding the contract event 0xa448afda7ea1e3a7a10fcab0c29fe9a9dd85791503bf0171f281521551c7ec05.
//
// Solidity: event BalanceChanged(address owner, uint256 balance)
func (_Rollup *RollupFilterer) ParseBalanceChanged(log types.Log) (*RollupBalanceChanged, error) {
	event := new(RollupBalanceChanged)
	if err := _Rollup.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Rollup contract.
type RollupReceivedIterator struct {
	Event *RollupReceived // Event containing the contract specifics and raw log

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
func (it *RollupReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupReceived)
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
		it.Event = new(RollupReceived)
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
func (it *RollupReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupReceived represents a Received event raised by the Rollup contract.
type RollupReceived struct {
	BlockCount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 blockCount)
func (_Rollup *RollupFilterer) FilterReceived(opts *bind.FilterOpts) (*RollupReceivedIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &RollupReceivedIterator{contract: _Rollup.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 blockCount)
func (_Rollup *RollupFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *RollupReceived) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupReceived)
				if err := _Rollup.contract.UnpackLog(event, "Received", log); err != nil {
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

// ParseReceived is a log parse operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 blockCount)
func (_Rollup *RollupFilterer) ParseReceived(log types.Log) (*RollupReceived, error) {
	event := new(RollupReceived)
	if err := _Rollup.contract.UnpackLog(event, "Received", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupTopHeightUpdateIterator is returned from FilterTopHeightUpdate and is used to iterate over the raw logs and unpacked data for TopHeightUpdate events raised by the Rollup contract.
type RollupTopHeightUpdateIterator struct {
	Event *RollupTopHeightUpdate // Event containing the contract specifics and raw log

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
func (it *RollupTopHeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupTopHeightUpdate)
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
		it.Event = new(RollupTopHeightUpdate)
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
func (it *RollupTopHeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupTopHeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupTopHeightUpdate represents a TopHeightUpdate event raised by the Rollup contract.
type RollupTopHeightUpdate struct {
	Height uint32
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTopHeightUpdate is a free log retrieval operation binding the contract event 0xd81514714f4a8739a3eb69a75a0e137ef8b794bbc36e312d096f74984ecd164c.
//
// Solidity: event TopHeightUpdate(uint32 height)
func (_Rollup *RollupFilterer) FilterTopHeightUpdate(opts *bind.FilterOpts) (*RollupTopHeightUpdateIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "TopHeightUpdate")
	if err != nil {
		return nil, err
	}
	return &RollupTopHeightUpdateIterator{contract: _Rollup.contract, event: "TopHeightUpdate", logs: logs, sub: sub}, nil
}

// WatchTopHeightUpdate is a free log subscription operation binding the contract event 0xd81514714f4a8739a3eb69a75a0e137ef8b794bbc36e312d096f74984ecd164c.
//
// Solidity: event TopHeightUpdate(uint32 height)
func (_Rollup *RollupFilterer) WatchTopHeightUpdate(opts *bind.WatchOpts, sink chan<- *RollupTopHeightUpdate) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "TopHeightUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupTopHeightUpdate)
				if err := _Rollup.contract.UnpackLog(event, "TopHeightUpdate", log); err != nil {
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

// ParseTopHeightUpdate is a log parse operation binding the contract event 0xd81514714f4a8739a3eb69a75a0e137ef8b794bbc36e312d096f74984ecd164c.
//
// Solidity: event TopHeightUpdate(uint32 height)
func (_Rollup *RollupFilterer) ParseTopHeightUpdate(log types.Log) (*RollupTopHeightUpdate, error) {
	event := new(RollupTopHeightUpdate)
	if err := _Rollup.contract.UnpackLog(event, "TopHeightUpdate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// RollupTransactionIterator is returned from FilterTransaction and is used to iterate over the raw logs and unpacked data for Transaction events raised by the Rollup contract.
type RollupTransactionIterator struct {
	Event *RollupTransaction // Event containing the contract specifics and raw log

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
func (it *RollupTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RollupTransaction)
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
		it.Event = new(RollupTransaction)
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
func (it *RollupTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RollupTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RollupTransaction represents a Transaction event raised by the Rollup contract.
type RollupTransaction struct {
	TranIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransaction is a free log retrieval operation binding the contract event 0x677f502ec68c44ae7ddd81ccc39e600212abede2c98a041949f132f22b8e4e21.
//
// Solidity: event Transaction(uint256 tranIndex)
func (_Rollup *RollupFilterer) FilterTransaction(opts *bind.FilterOpts) (*RollupTransactionIterator, error) {

	logs, sub, err := _Rollup.contract.FilterLogs(opts, "Transaction")
	if err != nil {
		return nil, err
	}
	return &RollupTransactionIterator{contract: _Rollup.contract, event: "Transaction", logs: logs, sub: sub}, nil
}

// WatchTransaction is a free log subscription operation binding the contract event 0x677f502ec68c44ae7ddd81ccc39e600212abede2c98a041949f132f22b8e4e21.
//
// Solidity: event Transaction(uint256 tranIndex)
func (_Rollup *RollupFilterer) WatchTransaction(opts *bind.WatchOpts, sink chan<- *RollupTransaction) (event.Subscription, error) {

	logs, sub, err := _Rollup.contract.WatchLogs(opts, "Transaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RollupTransaction)
				if err := _Rollup.contract.UnpackLog(event, "Transaction", log); err != nil {
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

// ParseTransaction is a log parse operation binding the contract event 0x677f502ec68c44ae7ddd81ccc39e600212abede2c98a041949f132f22b8e4e21.
//
// Solidity: event Transaction(uint256 tranIndex)
func (_Rollup *RollupFilterer) ParseTransaction(log types.Log) (*RollupTransaction, error) {
	event := new(RollupTransaction)
	if err := _Rollup.contract.UnpackLog(event, "Transaction", log); err != nil {
		return nil, err
	}
	return event, nil
}
