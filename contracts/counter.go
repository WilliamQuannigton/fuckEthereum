// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
	_ = abi.ConvertType
)

// CounterMetaData contains all meta data concerning the Counter contract.
var CounterMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountDecremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newCount\",\"type\":\"uint256\"}],\"name\":\"CountReset\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"decrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"increment\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015600e575f5ffd5b505f5f819055506103aa806100225f395ff3fe608060405234801561000f575f5ffd5b5060043610610055575f3560e01c80632baeceb714610059578063a87d942c14610077578063d09de08a14610095578063d826f88f146100b3578063dd51babf146100d1575b5f5ffd5b6100616100ef565b60405161006e9190610250565b60405180910390f35b61007f61018a565b60405161008c9190610250565b60405180910390f35b61009d610192565b6040516100aa9190610250565b60405180910390f35b6100bb6101ea565b6040516100c89190610250565b60405180910390f35b6100d9610230565b6040516100e69190610250565b60405180910390f35b5f5f5f5411610133576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161012a906102c3565b60405180910390fd5b60015f5f828254610144919061030e565b925050819055507f36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af95065f5460405161017b9190610250565b60405180910390a15f54905090565b5f5f54905090565b5f60015f5f8282546101a49190610341565b925050819055507f420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb5f546040516101db9190610250565b60405180910390a15f54905090565b5f5f5f819055507f5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f02239775f546040516102219190610250565b60405180910390a15f54905090565b5f5f54905090565b5f819050919050565b61024a81610238565b82525050565b5f6020820190506102635f830184610241565b92915050565b5f82825260208201905092915050565b7f436f756e7465722063616e6e6f74206265206e656761746976650000000000005f82015250565b5f6102ad601a83610269565b91506102b882610279565b602082019050919050565b5f6020820190508181035f8301526102da816102a1565b9050919050565b7f4e487b71000000000000000000000000000000000000000000000000000000005f52601160045260245ffd5b5f61031882610238565b915061032383610238565b925082820390508181111561033b5761033a6102e1565b5b92915050565b5f61034b82610238565b915061035683610238565b925082820190508082111561036e5761036d6102e1565b5b9291505056fea264697066735822122001ca5e3be7468c52122f50cd3a54557f712b6e6823c587ed28a4466a3e73860a64736f6c634300081e0033",
}

// CounterABI is the input ABI used to generate the binding from.
// Deprecated: Use CounterMetaData.ABI instead.
var CounterABI = CounterMetaData.ABI

// CounterBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CounterMetaData.Bin instead.
var CounterBin = CounterMetaData.Bin

// DeployCounter deploys a new Ethereum contract, binding an instance of Counter to it.
func DeployCounter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Counter, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CounterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// Counter is an auto generated Go binding around an Ethereum contract.
type Counter struct {
	CounterCaller     // Read-only binding to the contract
	CounterTransactor // Write-only binding to the contract
	CounterFilterer   // Log filterer for contract events
}

// CounterCaller is an auto generated read-only Go binding around an Ethereum contract.
type CounterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CounterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CounterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CounterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CounterSession struct {
	Contract     *Counter          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CounterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CounterCallerSession struct {
	Contract *CounterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// CounterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CounterTransactorSession struct {
	Contract     *CounterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CounterRaw is an auto generated low-level Go binding around an Ethereum contract.
type CounterRaw struct {
	Contract *Counter // Generic contract binding to access the raw methods on
}

// CounterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CounterCallerRaw struct {
	Contract *CounterCaller // Generic read-only contract binding to access the raw methods on
}

// CounterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CounterTransactorRaw struct {
	Contract *CounterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCounter creates a new instance of Counter, bound to a specific deployed contract.
func NewCounter(address common.Address, backend bind.ContractBackend) (*Counter, error) {
	contract, err := bindCounter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Counter{CounterCaller: CounterCaller{contract: contract}, CounterTransactor: CounterTransactor{contract: contract}, CounterFilterer: CounterFilterer{contract: contract}}, nil
}

// NewCounterCaller creates a new read-only instance of Counter, bound to a specific deployed contract.
func NewCounterCaller(address common.Address, caller bind.ContractCaller) (*CounterCaller, error) {
	contract, err := bindCounter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CounterCaller{contract: contract}, nil
}

// NewCounterTransactor creates a new write-only instance of Counter, bound to a specific deployed contract.
func NewCounterTransactor(address common.Address, transactor bind.ContractTransactor) (*CounterTransactor, error) {
	contract, err := bindCounter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CounterTransactor{contract: contract}, nil
}

// NewCounterFilterer creates a new log filterer instance of Counter, bound to a specific deployed contract.
func NewCounterFilterer(address common.Address, filterer bind.ContractFilterer) (*CounterFilterer, error) {
	contract, err := bindCounter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CounterFilterer{contract: contract}, nil
}

// bindCounter binds a generic wrapper to an already deployed contract.
func bindCounter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := CounterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.CounterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.CounterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Counter *CounterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Counter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Counter *CounterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Counter *CounterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Counter.Contract.contract.Transact(opts, method, params...)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCaller) GetCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCount is a free data retrieval call binding the contract method 0xa87d942c.
//
// Solidity: function getCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCount() (*big.Int, error) {
	return _Counter.Contract.GetCount(&_Counter.CallOpts)
}

// GetCurrentCount is a free data retrieval call binding the contract method 0xdd51babf.
//
// Solidity: function getCurrentCount() view returns(uint256)
func (_Counter *CounterCaller) GetCurrentCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Counter.contract.Call(opts, &out, "getCurrentCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentCount is a free data retrieval call binding the contract method 0xdd51babf.
//
// Solidity: function getCurrentCount() view returns(uint256)
func (_Counter *CounterSession) GetCurrentCount() (*big.Int, error) {
	return _Counter.Contract.GetCurrentCount(&_Counter.CallOpts)
}

// GetCurrentCount is a free data retrieval call binding the contract method 0xdd51babf.
//
// Solidity: function getCurrentCount() view returns(uint256)
func (_Counter *CounterCallerSession) GetCurrentCount() (*big.Int, error) {
	return _Counter.Contract.GetCurrentCount(&_Counter.CallOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterTransactor) Decrement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "decrement")
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Decrement is a paid mutator transaction binding the contract method 0x2baeceb7.
//
// Solidity: function decrement() returns(uint256)
func (_Counter *CounterTransactorSession) Decrement() (*types.Transaction, error) {
	return _Counter.Contract.Decrement(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns(uint256)
func (_Counter *CounterTransactorSession) Increment() (*types.Transaction, error) {
	return _Counter.Contract.Increment(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterTransactor) Reset(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Counter.contract.Transact(opts, "reset")
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// Reset is a paid mutator transaction binding the contract method 0xd826f88f.
//
// Solidity: function reset() returns(uint256)
func (_Counter *CounterTransactorSession) Reset() (*types.Transaction, error) {
	return _Counter.Contract.Reset(&_Counter.TransactOpts)
}

// CounterCountDecrementedIterator is returned from FilterCountDecremented and is used to iterate over the raw logs and unpacked data for CountDecremented events raised by the Counter contract.
type CounterCountDecrementedIterator struct {
	Event *CounterCountDecremented // Event containing the contract specifics and raw log

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
func (it *CounterCountDecrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountDecremented)
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
		it.Event = new(CounterCountDecremented)
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
func (it *CounterCountDecrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountDecrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountDecremented represents a CountDecremented event raised by the Counter contract.
type CounterCountDecremented struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountDecremented is a free log retrieval operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Counter *CounterFilterer) FilterCountDecremented(opts *bind.FilterOpts) (*CounterCountDecrementedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountDecremented")
	if err != nil {
		return nil, err
	}
	return &CounterCountDecrementedIterator{contract: _Counter.contract, event: "CountDecremented", logs: logs, sub: sub}, nil
}

// WatchCountDecremented is a free log subscription operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Counter *CounterFilterer) WatchCountDecremented(opts *bind.WatchOpts, sink chan<- *CounterCountDecremented) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountDecremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountDecremented)
				if err := _Counter.contract.UnpackLog(event, "CountDecremented", log); err != nil {
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

// ParseCountDecremented is a log parse operation binding the contract event 0x36bd77efe73a0782b8356dfffe895475b0a548122d84fdd60264949e18af9506.
//
// Solidity: event CountDecremented(uint256 newCount)
func (_Counter *CounterFilterer) ParseCountDecremented(log types.Log) (*CounterCountDecremented, error) {
	event := new(CounterCountDecremented)
	if err := _Counter.contract.UnpackLog(event, "CountDecremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterCountIncrementedIterator is returned from FilterCountIncremented and is used to iterate over the raw logs and unpacked data for CountIncremented events raised by the Counter contract.
type CounterCountIncrementedIterator struct {
	Event *CounterCountIncremented // Event containing the contract specifics and raw log

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
func (it *CounterCountIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountIncremented)
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
		it.Event = new(CounterCountIncremented)
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
func (it *CounterCountIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountIncremented represents a CountIncremented event raised by the Counter contract.
type CounterCountIncremented struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountIncremented is a free log retrieval operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Counter *CounterFilterer) FilterCountIncremented(opts *bind.FilterOpts) (*CounterCountIncrementedIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountIncremented")
	if err != nil {
		return nil, err
	}
	return &CounterCountIncrementedIterator{contract: _Counter.contract, event: "CountIncremented", logs: logs, sub: sub}, nil
}

// WatchCountIncremented is a free log subscription operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Counter *CounterFilterer) WatchCountIncremented(opts *bind.WatchOpts, sink chan<- *CounterCountIncremented) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountIncremented")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountIncremented)
				if err := _Counter.contract.UnpackLog(event, "CountIncremented", log); err != nil {
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

// ParseCountIncremented is a log parse operation binding the contract event 0x420680a649b45cbb7e97b24365d8ed81598dce543f2a2014d48fe328aa47e8bb.
//
// Solidity: event CountIncremented(uint256 newCount)
func (_Counter *CounterFilterer) ParseCountIncremented(log types.Log) (*CounterCountIncremented, error) {
	event := new(CounterCountIncremented)
	if err := _Counter.contract.UnpackLog(event, "CountIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CounterCountResetIterator is returned from FilterCountReset and is used to iterate over the raw logs and unpacked data for CountReset events raised by the Counter contract.
type CounterCountResetIterator struct {
	Event *CounterCountReset // Event containing the contract specifics and raw log

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
func (it *CounterCountResetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CounterCountReset)
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
		it.Event = new(CounterCountReset)
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
func (it *CounterCountResetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CounterCountResetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CounterCountReset represents a CountReset event raised by the Counter contract.
type CounterCountReset struct {
	NewCount *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCountReset is a free log retrieval operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Counter *CounterFilterer) FilterCountReset(opts *bind.FilterOpts) (*CounterCountResetIterator, error) {

	logs, sub, err := _Counter.contract.FilterLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return &CounterCountResetIterator{contract: _Counter.contract, event: "CountReset", logs: logs, sub: sub}, nil
}

// WatchCountReset is a free log subscription operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Counter *CounterFilterer) WatchCountReset(opts *bind.WatchOpts, sink chan<- *CounterCountReset) (event.Subscription, error) {

	logs, sub, err := _Counter.contract.WatchLogs(opts, "CountReset")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CounterCountReset)
				if err := _Counter.contract.UnpackLog(event, "CountReset", log); err != nil {
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

// ParseCountReset is a log parse operation binding the contract event 0x5b9d10f4ee515225030c54621fd1c542cbe568fa14df6e019aab2bc3f0223977.
//
// Solidity: event CountReset(uint256 newCount)
func (_Counter *CounterFilterer) ParseCountReset(log types.Log) (*CounterCountReset, error) {
	event := new(CounterCountReset)
	if err := _Counter.contract.UnpackLog(event, "CountReset", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
