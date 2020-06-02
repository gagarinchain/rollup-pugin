// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package eth

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

// EthABI is the input ABI used to generate the binding from.
const EthABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"balance\",\"type\":\"int64\"}],\"name\":\"BalanceChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tranIndex\",\"type\":\"uint256\"}],\"name\":\"Transaction\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"blockCount\",\"type\":\"uint256\"}],\"name\":\"Received\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"eight\",\"type\":\"int32\"}],\"name\":\"TopHeightUpdate\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTopHeight\",\"outputs\":[{\"name\":\"\",\"type\":\"int32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"header\",\"type\":\"bytes\"},{\"name\":\"rollup\",\"type\":\"bytes\"}],\"name\":\"addBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"b\",\"type\":\"bytes\"}],\"name\":\"bytesToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Eth is an auto generated Go binding around an Ethereum contract.
type Eth struct {
	EthCaller     // Read-only binding to the contract
	EthTransactor // Write-only binding to the contract
	EthFilterer   // Log filterer for contract events
}

// EthCaller is an auto generated read-only Go binding around an Ethereum contract.
type EthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EthSession struct {
	Contract     *Eth              // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EthCallerSession struct {
	Contract *EthCaller    // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EthTransactorSession struct {
	Contract     *EthTransactor    // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EthRaw is an auto generated low-level Go binding around an Ethereum contract.
type EthRaw struct {
	Contract *Eth // Generic contract binding to access the raw methods on
}

// EthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EthCallerRaw struct {
	Contract *EthCaller // Generic read-only contract binding to access the raw methods on
}

// EthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EthTransactorRaw struct {
	Contract *EthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEth creates a new instance of Eth, bound to a specific deployed contract.
func NewEth(address common.Address, backend bind.ContractBackend) (*Eth, error) {
	contract, err := bindEth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Eth{EthCaller: EthCaller{contract: contract}, EthTransactor: EthTransactor{contract: contract}, EthFilterer: EthFilterer{contract: contract}}, nil
}

// NewEthCaller creates a new read-only instance of Eth, bound to a specific deployed contract.
func NewEthCaller(address common.Address, caller bind.ContractCaller) (*EthCaller, error) {
	contract, err := bindEth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EthCaller{contract: contract}, nil
}

// NewEthTransactor creates a new write-only instance of Eth, bound to a specific deployed contract.
func NewEthTransactor(address common.Address, transactor bind.ContractTransactor) (*EthTransactor, error) {
	contract, err := bindEth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EthTransactor{contract: contract}, nil
}

// NewEthFilterer creates a new log filterer instance of Eth, bound to a specific deployed contract.
func NewEthFilterer(address common.Address, filterer bind.ContractFilterer) (*EthFilterer, error) {
	contract, err := bindEth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EthFilterer{contract: contract}, nil
}

// bindEth binds a generic wrapper to an already deployed contract.
func bindEth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.EthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.EthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Eth *EthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Eth.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Eth *EthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Eth *EthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Eth.Contract.contract.Transact(opts, method, params...)
}

// BytesToAddress is a free data retrieval call binding the contract method 0x42526e4e.
//
// Solidity: function bytesToAddress(bytes b) constant returns(address)
func (_Eth *EthCaller) BytesToAddress(opts *bind.CallOpts, b []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "bytesToAddress", b)
	return *ret0, err
}

// BytesToAddress is a free data retrieval call binding the contract method 0x42526e4e.
//
// Solidity: function bytesToAddress(bytes b) constant returns(address)
func (_Eth *EthSession) BytesToAddress(b []byte) (common.Address, error) {
	return _Eth.Contract.BytesToAddress(&_Eth.CallOpts, b)
}

// BytesToAddress is a free data retrieval call binding the contract method 0x42526e4e.
//
// Solidity: function bytesToAddress(bytes b) constant returns(address)
func (_Eth *EthCallerSession) BytesToAddress(b []byte) (common.Address, error) {
	return _Eth.Contract.BytesToAddress(&_Eth.CallOpts, b)
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() constant returns(int32)
func (_Eth *EthCaller) GetTopHeight(opts *bind.CallOpts) (int32, error) {
	var (
		ret0 = new(int32)
	)
	out := ret0
	err := _Eth.contract.Call(opts, out, "getTopHeight")
	return *ret0, err
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() constant returns(int32)
func (_Eth *EthSession) GetTopHeight() (int32, error) {
	return _Eth.Contract.GetTopHeight(&_Eth.CallOpts)
}

// GetTopHeight is a free data retrieval call binding the contract method 0x221068ae.
//
// Solidity: function getTopHeight() constant returns(int32)
func (_Eth *EthCallerSession) GetTopHeight() (int32, error) {
	return _Eth.Contract.GetTopHeight(&_Eth.CallOpts)
}

// AddBlock is a paid mutator transaction binding the contract method 0xf358de12.
//
// Solidity: function addBlock(bytes header, bytes rollup) returns()
func (_Eth *EthTransactor) AddBlock(opts *bind.TransactOpts, header []byte, rollup []byte) (*types.Transaction, error) {
	return _Eth.contract.Transact(opts, "addBlock", header, rollup)
}

// AddBlock is a paid mutator transaction binding the contract method 0xf358de12.
//
// Solidity: function addBlock(bytes header, bytes rollup) returns()
func (_Eth *EthSession) AddBlock(header []byte, rollup []byte) (*types.Transaction, error) {
	return _Eth.Contract.AddBlock(&_Eth.TransactOpts, header, rollup)
}

// AddBlock is a paid mutator transaction binding the contract method 0xf358de12.
//
// Solidity: function addBlock(bytes header, bytes rollup) returns()
func (_Eth *EthTransactorSession) AddBlock(header []byte, rollup []byte) (*types.Transaction, error) {
	return _Eth.Contract.AddBlock(&_Eth.TransactOpts, header, rollup)
}

// EthBalanceChangedIterator is returned from FilterBalanceChanged and is used to iterate over the raw logs and unpacked data for BalanceChanged events raised by the Eth contract.
type EthBalanceChangedIterator struct {
	Event *EthBalanceChanged // Event containing the contract specifics and raw log

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
func (it *EthBalanceChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthBalanceChanged)
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
		it.Event = new(EthBalanceChanged)
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
func (it *EthBalanceChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthBalanceChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthBalanceChanged represents a BalanceChanged event raised by the Eth contract.
type EthBalanceChanged struct {
	Owner   common.Address
	Balance int64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBalanceChanged is a free log retrieval operation binding the contract event 0x5b1f8b2cdcdb430bc7d8ba2272270311947a8b822b9bee7b284fcdf0243a07f6.
//
// Solidity: event BalanceChanged(address owner, int64 balance)
func (_Eth *EthFilterer) FilterBalanceChanged(opts *bind.FilterOpts) (*EthBalanceChangedIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return &EthBalanceChangedIterator{contract: _Eth.contract, event: "BalanceChanged", logs: logs, sub: sub}, nil
}

// WatchBalanceChanged is a free log subscription operation binding the contract event 0x5b1f8b2cdcdb430bc7d8ba2272270311947a8b822b9bee7b284fcdf0243a07f6.
//
// Solidity: event BalanceChanged(address owner, int64 balance)
func (_Eth *EthFilterer) WatchBalanceChanged(opts *bind.WatchOpts, sink chan<- *EthBalanceChanged) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "BalanceChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthBalanceChanged)
				if err := _Eth.contract.UnpackLog(event, "BalanceChanged", log); err != nil {
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

// EthReceivedIterator is returned from FilterReceived and is used to iterate over the raw logs and unpacked data for Received events raised by the Eth contract.
type EthReceivedIterator struct {
	Event *EthReceived // Event containing the contract specifics and raw log

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
func (it *EthReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthReceived)
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
		it.Event = new(EthReceived)
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
func (it *EthReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthReceived represents a Received event raised by the Eth contract.
type EthReceived struct {
	BlockCount *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterReceived is a free log retrieval operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 blockCount)
func (_Eth *EthFilterer) FilterReceived(opts *bind.FilterOpts) (*EthReceivedIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return &EthReceivedIterator{contract: _Eth.contract, event: "Received", logs: logs, sub: sub}, nil
}

// WatchReceived is a free log subscription operation binding the contract event 0xa8142743f8f70a4c26f3691cf4ed59718381fb2f18070ec52be1f1022d855557.
//
// Solidity: event Received(uint256 blockCount)
func (_Eth *EthFilterer) WatchReceived(opts *bind.WatchOpts, sink chan<- *EthReceived) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "Received")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthReceived)
				if err := _Eth.contract.UnpackLog(event, "Received", log); err != nil {
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

// EthTopHeightUpdateIterator is returned from FilterTopHeightUpdate and is used to iterate over the raw logs and unpacked data for TopHeightUpdate events raised by the Eth contract.
type EthTopHeightUpdateIterator struct {
	Event *EthTopHeightUpdate // Event containing the contract specifics and raw log

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
func (it *EthTopHeightUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthTopHeightUpdate)
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
		it.Event = new(EthTopHeightUpdate)
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
func (it *EthTopHeightUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthTopHeightUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthTopHeightUpdate represents a TopHeightUpdate event raised by the Eth contract.
type EthTopHeightUpdate struct {
	Eight int32
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTopHeightUpdate is a free log retrieval operation binding the contract event 0xa55a668cee7266e37bcda46dd9cc53ecf7f5cfffd94fdcce5272ce40c6e9b37f.
//
// Solidity: event TopHeightUpdate(int32 eight)
func (_Eth *EthFilterer) FilterTopHeightUpdate(opts *bind.FilterOpts) (*EthTopHeightUpdateIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "TopHeightUpdate")
	if err != nil {
		return nil, err
	}
	return &EthTopHeightUpdateIterator{contract: _Eth.contract, event: "TopHeightUpdate", logs: logs, sub: sub}, nil
}

// WatchTopHeightUpdate is a free log subscription operation binding the contract event 0xa55a668cee7266e37bcda46dd9cc53ecf7f5cfffd94fdcce5272ce40c6e9b37f.
//
// Solidity: event TopHeightUpdate(int32 eight)
func (_Eth *EthFilterer) WatchTopHeightUpdate(opts *bind.WatchOpts, sink chan<- *EthTopHeightUpdate) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "TopHeightUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthTopHeightUpdate)
				if err := _Eth.contract.UnpackLog(event, "TopHeightUpdate", log); err != nil {
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

// EthTransactionIterator is returned from FilterTransaction and is used to iterate over the raw logs and unpacked data for Transaction events raised by the Eth contract.
type EthTransactionIterator struct {
	Event *EthTransaction // Event containing the contract specifics and raw log

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
func (it *EthTransactionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EthTransaction)
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
		it.Event = new(EthTransaction)
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
func (it *EthTransactionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EthTransactionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EthTransaction represents a Transaction event raised by the Eth contract.
type EthTransaction struct {
	TranIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterTransaction is a free log retrieval operation binding the contract event 0x677f502ec68c44ae7ddd81ccc39e600212abede2c98a041949f132f22b8e4e21.
//
// Solidity: event Transaction(uint256 tranIndex)
func (_Eth *EthFilterer) FilterTransaction(opts *bind.FilterOpts) (*EthTransactionIterator, error) {

	logs, sub, err := _Eth.contract.FilterLogs(opts, "Transaction")
	if err != nil {
		return nil, err
	}
	return &EthTransactionIterator{contract: _Eth.contract, event: "Transaction", logs: logs, sub: sub}, nil
}

// WatchTransaction is a free log subscription operation binding the contract event 0x677f502ec68c44ae7ddd81ccc39e600212abede2c98a041949f132f22b8e4e21.
//
// Solidity: event Transaction(uint256 tranIndex)
func (_Eth *EthFilterer) WatchTransaction(opts *bind.WatchOpts, sink chan<- *EthTransaction) (event.Subscription, error) {

	logs, sub, err := _Eth.contract.WatchLogs(opts, "Transaction")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EthTransaction)
				if err := _Eth.contract.UnpackLog(event, "Transaction", log); err != nil {
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
