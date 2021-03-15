// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// DeprecatebleABI is the input ABI used to generate the binding from.
const DeprecatebleABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// DeprecatebleFuncSigs maps the 4-byte function signature to its string representation.
var DeprecatebleFuncSigs = map[string]string{
	"0fcc0c28": "deprecate()",
	"f2fde38b": "transferOwnership(address)",
}

// Deprecateble is an auto generated Go binding around an Ethereum contract.
type Deprecateble struct {
	DeprecatebleCaller     // Read-only binding to the contract
	DeprecatebleTransactor // Write-only binding to the contract
	DeprecatebleFilterer   // Log filterer for contract events
}

// DeprecatebleCaller is an auto generated read-only Go binding around an Ethereum contract.
type DeprecatebleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatebleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DeprecatebleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatebleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DeprecatebleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DeprecatebleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DeprecatebleSession struct {
	Contract     *Deprecateble     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DeprecatebleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DeprecatebleCallerSession struct {
	Contract *DeprecatebleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// DeprecatebleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DeprecatebleTransactorSession struct {
	Contract     *DeprecatebleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// DeprecatebleRaw is an auto generated low-level Go binding around an Ethereum contract.
type DeprecatebleRaw struct {
	Contract *Deprecateble // Generic contract binding to access the raw methods on
}

// DeprecatebleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DeprecatebleCallerRaw struct {
	Contract *DeprecatebleCaller // Generic read-only contract binding to access the raw methods on
}

// DeprecatebleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DeprecatebleTransactorRaw struct {
	Contract *DeprecatebleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDeprecateble creates a new instance of Deprecateble, bound to a specific deployed contract.
func NewDeprecateble(address common.Address, backend bind.ContractBackend) (*Deprecateble, error) {
	contract, err := bindDeprecateble(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Deprecateble{DeprecatebleCaller: DeprecatebleCaller{contract: contract}, DeprecatebleTransactor: DeprecatebleTransactor{contract: contract}, DeprecatebleFilterer: DeprecatebleFilterer{contract: contract}}, nil
}

// NewDeprecatebleCaller creates a new read-only instance of Deprecateble, bound to a specific deployed contract.
func NewDeprecatebleCaller(address common.Address, caller bind.ContractCaller) (*DeprecatebleCaller, error) {
	contract, err := bindDeprecateble(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DeprecatebleCaller{contract: contract}, nil
}

// NewDeprecatebleTransactor creates a new write-only instance of Deprecateble, bound to a specific deployed contract.
func NewDeprecatebleTransactor(address common.Address, transactor bind.ContractTransactor) (*DeprecatebleTransactor, error) {
	contract, err := bindDeprecateble(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DeprecatebleTransactor{contract: contract}, nil
}

// NewDeprecatebleFilterer creates a new log filterer instance of Deprecateble, bound to a specific deployed contract.
func NewDeprecatebleFilterer(address common.Address, filterer bind.ContractFilterer) (*DeprecatebleFilterer, error) {
	contract, err := bindDeprecateble(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DeprecatebleFilterer{contract: contract}, nil
}

// bindDeprecateble binds a generic wrapper to an already deployed contract.
func bindDeprecateble(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DeprecatebleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deprecateble *DeprecatebleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deprecateble.Contract.DeprecatebleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deprecateble *DeprecatebleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deprecateble.Contract.DeprecatebleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deprecateble *DeprecatebleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deprecateble.Contract.DeprecatebleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Deprecateble *DeprecatebleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Deprecateble.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Deprecateble *DeprecatebleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deprecateble.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Deprecateble *DeprecatebleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Deprecateble.Contract.contract.Transact(opts, method, params...)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Deprecateble *DeprecatebleTransactor) Deprecate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Deprecateble.contract.Transact(opts, "deprecate")
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Deprecateble *DeprecatebleSession) Deprecate() (*types.Transaction, error) {
	return _Deprecateble.Contract.Deprecate(&_Deprecateble.TransactOpts)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_Deprecateble *DeprecatebleTransactorSession) Deprecate() (*types.Transaction, error) {
	return _Deprecateble.Contract.Deprecate(&_Deprecateble.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deprecateble *DeprecatebleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Deprecateble.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deprecateble *DeprecatebleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deprecateble.Contract.TransferOwnership(&_Deprecateble.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Deprecateble *DeprecatebleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Deprecateble.Contract.TransferOwnership(&_Deprecateble.TransactOpts, newOwner)
}

// DeprecatebleDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the Deprecateble contract.
type DeprecatebleDeprecateIterator struct {
	Event *DeprecatebleDeprecate // Event containing the contract specifics and raw log

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
func (it *DeprecatebleDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(DeprecatebleDeprecate)
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
		it.Event = new(DeprecatebleDeprecate)
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
func (it *DeprecatebleDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *DeprecatebleDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// DeprecatebleDeprecate represents a Deprecate event raised by the Deprecateble contract.
type DeprecatebleDeprecate struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_Deprecateble *DeprecatebleFilterer) FilterDeprecate(opts *bind.FilterOpts, account []common.Address) (*DeprecatebleDeprecateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Deprecateble.contract.FilterLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return &DeprecatebleDeprecateIterator{contract: _Deprecateble.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_Deprecateble *DeprecatebleFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *DeprecatebleDeprecate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _Deprecateble.contract.WatchLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(DeprecatebleDeprecate)
				if err := _Deprecateble.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_Deprecateble *DeprecatebleFilterer) ParseDeprecate(log types.Log) (*DeprecatebleDeprecate, error) {
	event := new(DeprecatebleDeprecate)
	if err := _Deprecateble.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"39509351": "increaseAllowance(address,uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Transactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20Session) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.DecreaseAllowance(&_ERC20.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Transactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20Session) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_ERC20 *ERC20TransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.IncreaseAllowance(&_ERC20.TransactOpts, spender, addedValue)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, sender, recipient, amount)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"f2fde38b": "transferOwnership(address)",
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// SWOPABI is the input ABI used to generate the binding from.
const SWOPABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// SWOPFuncSigs maps the 4-byte function signature to its string representation.
var SWOPFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"47e7ef24": "deposit(address,uint256)",
	"0fcc0c28": "deprecate()",
	"39509351": "increaseAllowance(address,uint256)",
	"06fdde03": "name()",
	"a694fc3a": "stake(uint256)",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"51cff8d9": "withdraw(address)",
}

// SWOPBin is the compiled bytecode used for deploying new contracts.
var SWOPBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03199081163390811783556001805490921617815560028054918201815590915264e8d4a510007f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910155611a28806100776000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806351cff8d911610097578063a694fc3a11610066578063a694fc3a14610312578063a9059cbb1461032f578063dd62ed3e1461035b578063f2fde38b1461038957610100565b806351cff8d91461029257806370a08231146102b857806395d89b41146102de578063a457c2d7146102e657610100565b806323b872dd116100d357806323b872dd146101e6578063313ce5671461021c578063395093511461023a57806347e7ef241461026657610100565b806306fdde0314610105578063095ea7b3146101825780630fcc0c28146101c257806318160ddd146101cc575b600080fd5b61010d6103af565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014757818101518382015260200161012f565b50505050905090810190601f1680156101745780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101ae6004803603604081101561019857600080fd5b506001600160a01b0381351690602001356103d5565b604080519115158252519081900360200190f35b6101ca6103eb565b005b6101d4610489565b60408051918252519081900360200190f35b6101ae600480360360608110156101fc57600080fd5b506001600160a01b038135811691602081013590911690604001356104da565b610224610562565b6040805160ff9092168252519081900360200190f35b6101ae6004803603604081101561025057600080fd5b506001600160a01b038135169060200135610567565b6101ae6004803603604081101561027c57600080fd5b506001600160a01b0381351690602001356105ed565b6101ae600480360360208110156102a857600080fd5b50356001600160a01b0316610920565b6101d4600480360360208110156102ce57600080fd5b50356001600160a01b0316610bdb565b61010d610de6565b6101ae600480360360408110156102fc57600080fd5b506001600160a01b038135169060200135610e06565b6101ae6004803603602081101561032857600080fd5b5035610e76565b6101ae6004803603604081101561034557600080fd5b506001600160a01b03813516906020013561114d565b6101d46004803603604081101561037157600080fd5b506001600160a01b038135811691602001351661115a565b6101ca6004803603602081101561039f57600080fd5b50356001600160a01b0316611185565b60408051808201909152600c81526b0577261707065642053574f560a41b602082015290565b60006103e233848461124a565b50600192915050565b6000546001600160a01b031633148061040e57506001546001600160a01b031633145b6104495760405162461bcd60e51b81526004018080602001828103825260298152602001806118826029913960400191505060405180910390fd5b6001805460ff60a01b1916600160a01b17905560405133907fcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e90600090a2565b600354600454600091908082018211156104d45760405162461bcd60e51b81526004018080602001828103825260228152602001806118f86022913960400191505060405180910390fd5b01905090565b60006104e784848461137f565b6001600160a01b03841660009081526008602090815260408083203384529091529020548083111561054a5760405162461bcd60e51b81526004018080602001828103825260288152602001806118ab6028913960400191505060405180910390fd5b610557853385840361124a565b506001949350505050565b601290565b3360009081526008602090815260408083206001600160a01b03861684529091528120548281018111156105d6576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b6105e3338585840161124a565b5060019392505050565b600080546001600160a01b031633148061061157506001546001600160a01b031633145b61064c5760405162461bcd60e51b81526004018080602001828103825260298152602001806118826029913960400191505060405180910390fd5b600154600160a01b900460ff16156106955760405162461bcd60e51b81526004018080602001828103825260248152602001806119cf6024913960400191505060405180910390fd5b600082116106e1576040805162461bcd60e51b81526020600482015260146024820152730616d6f756e742073686f756c64206265203e20360641b604482015290519081900360640190fd5b6001600160a01b03831661073c576040805162461bcd60e51b815260206004820152601b60248201527f6465706f73697420746f20746865207a65726f20616464726573730000000000604482015290519081900360640190fd5b600454828101811115610796576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f72206465706f736974000000604482015290519081900360640190fd5b8083016004556001600160a01b03841660009081526006602052604090205480610801576107c385610bdb565b6001600160a01b03861660009081526005602090815260408083209390935560025460078252838320600019909101905560069052208490556108e6565b6001600160a01b038516600090815260076020526040902054600254600019018114156108a15781858301101561087f576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f72206465706f736974000000604482015290519081900360640190fd5b6001600160a01b038616600090815260066020526040902082860190556108e4565b6108aa86610bdb565b6001600160a01b03871660009081526005602090815260408083209390935560025460078252838320600019909101905560069052208590555b505b6040805185815290516001600160a01b038716916000916000805160206119418339815191529181900360200190a3506001949350505050565b600080546001600160a01b031633148061094457506001546001600160a01b031633145b61097f5760405162461bcd60e51b81526004018080602001828103825260298152602001806118826029913960400191505060405180910390fd5b600154600160a01b900460ff16156109c85760405162461bcd60e51b81526004018080602001828103825260248152602001806119cf6024913960400191505060405180910390fd5b6001600160a01b03821660009081526006602090815260408083205460079092529091205460025460001901811415610b22576001600160a01b038416600090815260056020526040902054600354811115610a555760405162461bcd60e51b81526004018080602001828103825260258152602001806118d36025913960400191505060405180910390fd5b600380548290039055600454831115610a9f5760405162461bcd60e51b815260040180806020018281038252602781526020018061191a6027913960400191505060405180910390fd5b8260045403600481905550808382011015610aeb5760405162461bcd60e51b815260040180806020018281038252603081526020018061178b6030913960400191505060405180910390fd5b60408051828501815290516000916001600160a01b038816916000805160206119418339815191529181900360200190a350610baa565b6000610b2d85610bdb565b60035490915080821115610b725760405162461bcd60e51b81526004018080602001828103825260258152602001806118d36025913960400191505060405180910390fd5b8181036003556040805183815290516000916001600160a01b038916916000805160206119418339815191529181900360200190a350505b5050506001600160a01b0381166000908152600560209081526040808320839055600690915281205560015b919050565b6001600160a01b038116600090815260056020908152604080832054600690925282205481158015610c0b575080155b15610c1b57600092505050610bd6565b6001600160a01b03841660009081526007602052604090205460025460001901811415610ca257828284011015610c99576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722062616c616e6365000000604482015290519081900360640190fd5b50019050610bd6565b81610cf85760028054600091906000198101908110610cbd57fe5b9060005260206000200154905060028281548110610cd757fe5b906000526020600020015484820281610cec57fe5b04945050505050610bd6565b600060028281548110610d0757fe5b600091825260209091200154600280546000198101908110610d2557fe5b9060005260206000200154850281610d3957fe5b049050600060028360010181548110610d4e57fe5b600091825260209091200154600280546000198101908110610d6c57fe5b9060005260206000200154850281610d8057fe5b049050818282011015610dda576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722062616c616e6365000000604482015290519081900360640190fd5b019350610bd692505050565b6040805180820190915260068152650677753574f560d41b602082015290565b3360009081526008602090815260408083206001600160a01b038616845290915281205480831115610e695760405162461bcd60e51b81526004018080602001828103825260258152602001806119aa6025913960400191505060405180910390fd5b6105e3338585840361124a565b600080546001600160a01b0316331480610e9a57506001546001600160a01b031633145b610ed55760405162461bcd60e51b81526004018080602001828103825260298152602001806118826029913960400191505060405180910390fd5b600154600160a01b900460ff1615610f1e5760405162461bcd60e51b81526004018080602001828103825260248152602001806119cf6024913960400191505060405180910390fd5b60008211610f6a576040805162461bcd60e51b815260206004820152601460248201527307265776172642073686f756c64206265203e20360641b604482015290519081900360640190fd5b60035460045481610fb5576002805460018101825560009190915264e8d4a510007f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101556110ba565b60028054600091906000198101908110610fcb57fe5b9060005260206000200154905060008364e8d4a51000870281610fea57fe5b0490508064e8d4a5100082011015611049576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722070657263656e74000000604482015290519081900360640190fd5b64e8d4a51000818101906002908483028254600181018455600093845260209093209190049101558685018511156110b25760405162461bcd60e51b815260040180806020018281038252602b815260200180611857602b913960400191505060405180910390fd5b505050908301905b8181830110156110fb5760405162461bcd60e51b81526004018080602001828103825260228152602001806118f86022913960400191505060405180910390fd5b8181016003556000600455600254604080519182526020820186905280517f45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da529281900390910190a15060019392505050565b60006103e233848461137f565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b6000546001600160a01b03163314806111a857506001546001600160a01b031633145b6111e35760405162461bcd60e51b81526004018080602001828103825260298152602001806118826029913960400191505060405180910390fd5b6001600160a01b0381166112285760405162461bcd60e51b81526004018080602001828103825260268152602001806117bb6026913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff16156112935760405162461bcd60e51b81526004018080602001828103825260248152602001806119cf6024913960400191505060405180910390fd5b6001600160a01b0383166112d85760405162461bcd60e51b81526004018080602001828103825260248152602001806119866024913960400191505060405180910390fd5b6001600160a01b03821661131d5760405162461bcd60e51b81526004018080602001828103825260228152602001806117e16022913960400191505060405180910390fd5b6001600160a01b03808416600081815260086020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600154600160a01b900460ff16156113c85760405162461bcd60e51b81526004018080602001828103825260248152602001806119cf6024913960400191505060405180910390fd5b60008111611414576040805162461bcd60e51b81526020600482015260146024820152730616d6f756e742073686f756c64206265203e20360641b604482015290519081900360640190fd5b6001600160a01b0383166114595760405162461bcd60e51b81526004018080602001828103825260258152602001806119616025913960400191505060405180910390fd5b6001600160a01b03821661149e5760405162461bcd60e51b81526004018080602001828103825260238152602001806117686023913960400191505060405180910390fd5b6001600160a01b038316600090815260066020908152604080832054600790925282205490918215806114d75750600254600019018214155b1561156c5760006114e787610bdb565b9050808511156115285760405162461bcd60e51b81526004018080602001828103825260268152602001806118036026913960400191505060405180910390fd5b6001600160a01b0387166000908152600560209081526040808320938890039093556006815282822082905560025460079091529190206000199091019055611620565b82841161159757506001600160a01b0385166000908152600660205260409020838303905582611620565b6001600160a01b0386166000908152600560205260409020548385038110156115f15760405162461bcd60e51b81526004018080602001828103825260268152602001806118036026913960400191505060405180910390fd5b6001600160a01b0387166000908152600560209081526040808320878903909403909355600690529081205550815b6001600160a01b038516600090815260066020908152604080832054600790925290912054909350915082158061165d5750600254600019018214155b156116f657600061166d86610bdb565b905080818387030110156116b25760405162461bcd60e51b815260040180806020018281038252602e815260200180611829602e913960400191505060405180910390fd5b6001600160a01b0386166000908152600560209081526040808320858903949094019093556002546007825283832060001990910190556006905220819055611726565b6001600160a01b038516600090815260056020908152604080832080548589030190556006909152902083820190555b846001600160a01b0316866001600160a01b0316600080516020611941833981519152866040518082815260200191505060405180910390a350505050505056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573736164646974696f6e206f766572666c6f7720666f7220746f74616c2062616c616e6365202b206f6c644465706f7369744f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a206164646974696f6e206f766572666c6f7720666f7220726563697069656e742062616c616e63656164646974696f6e206f766572666c6f7720666f7220746f74616c20737570706c79202b207265776172644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572206f722061646d696e45524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63657375627472616374696f6e206f766572666c6f7720666f7220746f74616c20737570706c796164646974696f6e206f766572666c6f7720666f7220746f74616c20737570706c797375627472616374696f6e206f766572666c6f7720666f72206c6971756964206465706f736974ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f446570726563617465626c653a20636f6e74726163742069732064657072656361746564a26469706673582212205e20de16d0f1a56c0314b152506fc8512bbc91a0f056c4a7b8bd4299d5f72fd464736f6c634300060a0033"

// DeploySWOP deploys a new Ethereum contract, binding an instance of SWOP to it.
func DeploySWOP(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SWOP, error) {
	parsed, err := abi.JSON(strings.NewReader(SWOPABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SWOPBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SWOP{SWOPCaller: SWOPCaller{contract: contract}, SWOPTransactor: SWOPTransactor{contract: contract}, SWOPFilterer: SWOPFilterer{contract: contract}}, nil
}

// SWOP is an auto generated Go binding around an Ethereum contract.
type SWOP struct {
	SWOPCaller     // Read-only binding to the contract
	SWOPTransactor // Write-only binding to the contract
	SWOPFilterer   // Log filterer for contract events
}

// SWOPCaller is an auto generated read-only Go binding around an Ethereum contract.
type SWOPCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWOPTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SWOPTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWOPFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SWOPFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SWOPSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SWOPSession struct {
	Contract     *SWOP             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SWOPCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SWOPCallerSession struct {
	Contract *SWOPCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// SWOPTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SWOPTransactorSession struct {
	Contract     *SWOPTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SWOPRaw is an auto generated low-level Go binding around an Ethereum contract.
type SWOPRaw struct {
	Contract *SWOP // Generic contract binding to access the raw methods on
}

// SWOPCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SWOPCallerRaw struct {
	Contract *SWOPCaller // Generic read-only contract binding to access the raw methods on
}

// SWOPTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SWOPTransactorRaw struct {
	Contract *SWOPTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSWOP creates a new instance of SWOP, bound to a specific deployed contract.
func NewSWOP(address common.Address, backend bind.ContractBackend) (*SWOP, error) {
	contract, err := bindSWOP(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SWOP{SWOPCaller: SWOPCaller{contract: contract}, SWOPTransactor: SWOPTransactor{contract: contract}, SWOPFilterer: SWOPFilterer{contract: contract}}, nil
}

// NewSWOPCaller creates a new read-only instance of SWOP, bound to a specific deployed contract.
func NewSWOPCaller(address common.Address, caller bind.ContractCaller) (*SWOPCaller, error) {
	contract, err := bindSWOP(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SWOPCaller{contract: contract}, nil
}

// NewSWOPTransactor creates a new write-only instance of SWOP, bound to a specific deployed contract.
func NewSWOPTransactor(address common.Address, transactor bind.ContractTransactor) (*SWOPTransactor, error) {
	contract, err := bindSWOP(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SWOPTransactor{contract: contract}, nil
}

// NewSWOPFilterer creates a new log filterer instance of SWOP, bound to a specific deployed contract.
func NewSWOPFilterer(address common.Address, filterer bind.ContractFilterer) (*SWOPFilterer, error) {
	contract, err := bindSWOP(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SWOPFilterer{contract: contract}, nil
}

// bindSWOP binds a generic wrapper to an already deployed contract.
func bindSWOP(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SWOPABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWOP *SWOPRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SWOP.Contract.SWOPCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWOP *SWOPRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWOP.Contract.SWOPTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWOP *SWOPRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWOP.Contract.SWOPTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SWOP *SWOPCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SWOP.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SWOP *SWOPTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWOP.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SWOP *SWOPTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SWOP.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SWOP *SWOPCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SWOP *SWOPSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SWOP.Contract.Allowance(&_SWOP.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_SWOP *SWOPCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _SWOP.Contract.Allowance(&_SWOP.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SWOP *SWOPCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SWOP *SWOPSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SWOP.Contract.BalanceOf(&_SWOP.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_SWOP *SWOPCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _SWOP.Contract.BalanceOf(&_SWOP.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SWOP *SWOPCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SWOP *SWOPSession) Decimals() (uint8, error) {
	return _SWOP.Contract.Decimals(&_SWOP.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_SWOP *SWOPCallerSession) Decimals() (uint8, error) {
	return _SWOP.Contract.Decimals(&_SWOP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SWOP *SWOPCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SWOP *SWOPSession) Name() (string, error) {
	return _SWOP.Contract.Name(&_SWOP.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_SWOP *SWOPCallerSession) Name() (string, error) {
	return _SWOP.Contract.Name(&_SWOP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SWOP *SWOPCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SWOP *SWOPSession) Symbol() (string, error) {
	return _SWOP.Contract.Symbol(&_SWOP.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_SWOP *SWOPCallerSession) Symbol() (string, error) {
	return _SWOP.Contract.Symbol(&_SWOP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SWOP *SWOPCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SWOP.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SWOP *SWOPSession) TotalSupply() (*big.Int, error) {
	return _SWOP.Contract.TotalSupply(&_SWOP.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_SWOP *SWOPCallerSession) TotalSupply() (*big.Int, error) {
	return _SWOP.Contract.TotalSupply(&_SWOP.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SWOP *SWOPSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Approve(&_SWOP.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Approve(&_SWOP.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SWOP *SWOPTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SWOP *SWOPSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.DecreaseAllowance(&_SWOP.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_SWOP *SWOPTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.DecreaseAllowance(&_SWOP.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_SWOP *SWOPSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Deposit(&_SWOP.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Deposit(&_SWOP.TransactOpts, account, amount)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_SWOP *SWOPTransactor) Deprecate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "deprecate")
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_SWOP *SWOPSession) Deprecate() (*types.Transaction, error) {
	return _SWOP.Contract.Deprecate(&_SWOP.TransactOpts)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_SWOP *SWOPTransactorSession) Deprecate() (*types.Transaction, error) {
	return _SWOP.Contract.Deprecate(&_SWOP.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SWOP *SWOPTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SWOP *SWOPSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.IncreaseAllowance(&_SWOP.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_SWOP *SWOPTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.IncreaseAllowance(&_SWOP.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_SWOP *SWOPTransactor) Stake(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "stake", reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_SWOP *SWOPSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Stake(&_SWOP.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_SWOP *SWOPTransactorSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Stake(&_SWOP.TransactOpts, reward)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Transfer(&_SWOP.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.Transfer(&_SWOP.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.TransferFrom(&_SWOP.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_SWOP *SWOPTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _SWOP.Contract.TransferFrom(&_SWOP.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SWOP *SWOPTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SWOP *SWOPSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SWOP.Contract.TransferOwnership(&_SWOP.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SWOP *SWOPTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SWOP.Contract.TransferOwnership(&_SWOP.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_SWOP *SWOPTransactor) Withdraw(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _SWOP.contract.Transact(opts, "withdraw", account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_SWOP *SWOPSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _SWOP.Contract.Withdraw(&_SWOP.TransactOpts, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_SWOP *SWOPTransactorSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _SWOP.Contract.Withdraw(&_SWOP.TransactOpts, account)
}

// SWOPApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the SWOP contract.
type SWOPApprovalIterator struct {
	Event *SWOPApproval // Event containing the contract specifics and raw log

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
func (it *SWOPApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SWOPApproval)
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
		it.Event = new(SWOPApproval)
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
func (it *SWOPApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SWOPApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SWOPApproval represents a Approval event raised by the SWOP contract.
type SWOPApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SWOP *SWOPFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*SWOPApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SWOP.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &SWOPApprovalIterator{contract: _SWOP.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SWOP *SWOPFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *SWOPApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _SWOP.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SWOPApproval)
				if err := _SWOP.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_SWOP *SWOPFilterer) ParseApproval(log types.Log) (*SWOPApproval, error) {
	event := new(SWOPApproval)
	if err := _SWOP.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SWOPDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the SWOP contract.
type SWOPDeprecateIterator struct {
	Event *SWOPDeprecate // Event containing the contract specifics and raw log

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
func (it *SWOPDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SWOPDeprecate)
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
		it.Event = new(SWOPDeprecate)
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
func (it *SWOPDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SWOPDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SWOPDeprecate represents a Deprecate event raised by the SWOP contract.
type SWOPDeprecate struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_SWOP *SWOPFilterer) FilterDeprecate(opts *bind.FilterOpts, account []common.Address) (*SWOPDeprecateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SWOP.contract.FilterLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return &SWOPDeprecateIterator{contract: _SWOP.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_SWOP *SWOPFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *SWOPDeprecate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _SWOP.contract.WatchLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SWOPDeprecate)
				if err := _SWOP.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_SWOP *SWOPFilterer) ParseDeprecate(log types.Log) (*SWOPDeprecate, error) {
	event := new(SWOPDeprecate)
	if err := _SWOP.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SWOPRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the SWOP contract.
type SWOPRewardIterator struct {
	Event *SWOPReward // Event containing the contract specifics and raw log

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
func (it *SWOPRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SWOPReward)
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
		it.Event = new(SWOPReward)
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
func (it *SWOPRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SWOPRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SWOPReward represents a Reward event raised by the SWOP contract.
type SWOPReward struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_SWOP *SWOPFilterer) FilterReward(opts *bind.FilterOpts) (*SWOPRewardIterator, error) {

	logs, sub, err := _SWOP.contract.FilterLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return &SWOPRewardIterator{contract: _SWOP.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_SWOP *SWOPFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *SWOPReward) (event.Subscription, error) {

	logs, sub, err := _SWOP.contract.WatchLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SWOPReward)
				if err := _SWOP.contract.UnpackLog(event, "Reward", log); err != nil {
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

// ParseReward is a log parse operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_SWOP *SWOPFilterer) ParseReward(log types.Log) (*SWOPReward, error) {
	event := new(SWOPReward)
	if err := _SWOP.contract.UnpackLog(event, "Reward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// SWOPTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the SWOP contract.
type SWOPTransferIterator struct {
	Event *SWOPTransfer // Event containing the contract specifics and raw log

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
func (it *SWOPTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SWOPTransfer)
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
		it.Event = new(SWOPTransfer)
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
func (it *SWOPTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SWOPTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SWOPTransfer represents a Transfer event raised by the SWOP contract.
type SWOPTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SWOP *SWOPFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*SWOPTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SWOP.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SWOPTransferIterator{contract: _SWOP.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SWOP *SWOPFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *SWOPTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SWOP.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SWOPTransfer)
				if err := _SWOP.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_SWOP *SWOPFilterer) ParseTransfer(log types.Log) (*SWOPTransfer, error) {
	event := new(SWOPTransfer)
	if err := _SWOP.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StakingABI is the input ABI used to generate the binding from.
const StakingABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StakingFuncSigs maps the 4-byte function signature to its string representation.
var StakingFuncSigs = map[string]string{
	"47e7ef24": "deposit(address,uint256)",
	"a694fc3a": "stake(uint256)",
	"51cff8d9": "withdraw(address)",
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_Staking *StakingTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_Staking *StakingSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_Staking *StakingTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Deposit(&_Staking.TransactOpts, account, amount)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_Staking *StakingTransactor) Stake(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "stake", reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_Staking *StakingSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_Staking *StakingTransactorSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _Staking.Contract.Stake(&_Staking.TransactOpts, reward)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_Staking *StakingTransactor) Withdraw(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "withdraw", account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_Staking *StakingSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_Staking *StakingTransactorSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Withdraw(&_Staking.TransactOpts, account)
}

// StakingRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the Staking contract.
type StakingRewardIterator struct {
	Event *StakingReward // Event containing the contract specifics and raw log

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
func (it *StakingRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingReward)
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
		it.Event = new(StakingReward)
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
func (it *StakingRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingReward represents a Reward event raised by the Staking contract.
type StakingReward struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_Staking *StakingFilterer) FilterReward(opts *bind.FilterOpts) (*StakingRewardIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return &StakingRewardIterator{contract: _Staking.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_Staking *StakingFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *StakingReward) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingReward)
				if err := _Staking.contract.UnpackLog(event, "Reward", log); err != nil {
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

// ParseReward is a log parse operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_Staking *StakingFilterer) ParseReward(log types.Log) (*StakingReward, error) {
	event := new(StakingReward)
	if err := _Staking.contract.UnpackLog(event, "Reward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandartTokenABI is the input ABI used to generate the binding from.
const StandartTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StandartTokenFuncSigs maps the 4-byte function signature to its string representation.
var StandartTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"a457c2d7": "decreaseAllowance(address,uint256)",
	"47e7ef24": "deposit(address,uint256)",
	"0fcc0c28": "deprecate()",
	"39509351": "increaseAllowance(address,uint256)",
	"a694fc3a": "stake(uint256)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"51cff8d9": "withdraw(address)",
}

// StandartToken is an auto generated Go binding around an Ethereum contract.
type StandartToken struct {
	StandartTokenCaller     // Read-only binding to the contract
	StandartTokenTransactor // Write-only binding to the contract
	StandartTokenFilterer   // Log filterer for contract events
}

// StandartTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StandartTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandartTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StandartTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandartTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StandartTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StandartTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StandartTokenSession struct {
	Contract     *StandartToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StandartTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StandartTokenCallerSession struct {
	Contract *StandartTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// StandartTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StandartTokenTransactorSession struct {
	Contract     *StandartTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// StandartTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StandartTokenRaw struct {
	Contract *StandartToken // Generic contract binding to access the raw methods on
}

// StandartTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StandartTokenCallerRaw struct {
	Contract *StandartTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StandartTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StandartTokenTransactorRaw struct {
	Contract *StandartTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStandartToken creates a new instance of StandartToken, bound to a specific deployed contract.
func NewStandartToken(address common.Address, backend bind.ContractBackend) (*StandartToken, error) {
	contract, err := bindStandartToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StandartToken{StandartTokenCaller: StandartTokenCaller{contract: contract}, StandartTokenTransactor: StandartTokenTransactor{contract: contract}, StandartTokenFilterer: StandartTokenFilterer{contract: contract}}, nil
}

// NewStandartTokenCaller creates a new read-only instance of StandartToken, bound to a specific deployed contract.
func NewStandartTokenCaller(address common.Address, caller bind.ContractCaller) (*StandartTokenCaller, error) {
	contract, err := bindStandartToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StandartTokenCaller{contract: contract}, nil
}

// NewStandartTokenTransactor creates a new write-only instance of StandartToken, bound to a specific deployed contract.
func NewStandartTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StandartTokenTransactor, error) {
	contract, err := bindStandartToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StandartTokenTransactor{contract: contract}, nil
}

// NewStandartTokenFilterer creates a new log filterer instance of StandartToken, bound to a specific deployed contract.
func NewStandartTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StandartTokenFilterer, error) {
	contract, err := bindStandartToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StandartTokenFilterer{contract: contract}, nil
}

// bindStandartToken binds a generic wrapper to an already deployed contract.
func bindStandartToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StandartTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandartToken *StandartTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandartToken.Contract.StandartTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandartToken *StandartTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandartToken.Contract.StandartTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandartToken *StandartTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandartToken.Contract.StandartTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StandartToken *StandartTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StandartToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StandartToken *StandartTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandartToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StandartToken *StandartTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StandartToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandartToken *StandartTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandartToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandartToken *StandartTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandartToken.Contract.Allowance(&_StandartToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StandartToken *StandartTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StandartToken.Contract.Allowance(&_StandartToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandartToken *StandartTokenCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StandartToken.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandartToken *StandartTokenSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StandartToken.Contract.BalanceOf(&_StandartToken.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StandartToken *StandartTokenCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StandartToken.Contract.BalanceOf(&_StandartToken.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandartToken *StandartTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StandartToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandartToken *StandartTokenSession) TotalSupply() (*big.Int, error) {
	return _StandartToken.Contract.TotalSupply(&_StandartToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StandartToken *StandartTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StandartToken.Contract.TotalSupply(&_StandartToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Approve(&_StandartToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Approve(&_StandartToken.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandartToken *StandartTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandartToken *StandartTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.DecreaseAllowance(&_StandartToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.DecreaseAllowance(&_StandartToken.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Deposit(&_StandartToken.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Deposit(&_StandartToken.TransactOpts, account, amount)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_StandartToken *StandartTokenTransactor) Deprecate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "deprecate")
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_StandartToken *StandartTokenSession) Deprecate() (*types.Transaction, error) {
	return _StandartToken.Contract.Deprecate(&_StandartToken.TransactOpts)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_StandartToken *StandartTokenTransactorSession) Deprecate() (*types.Transaction, error) {
	return _StandartToken.Contract.Deprecate(&_StandartToken.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandartToken *StandartTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandartToken *StandartTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.IncreaseAllowance(&_StandartToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.IncreaseAllowance(&_StandartToken.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_StandartToken *StandartTokenTransactor) Stake(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "stake", reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_StandartToken *StandartTokenSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Stake(&_StandartToken.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Stake(&_StandartToken.TransactOpts, reward)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Transfer(&_StandartToken.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.Transfer(&_StandartToken.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.TransferFrom(&_StandartToken.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _StandartToken.Contract.TransferFrom(&_StandartToken.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandartToken *StandartTokenTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandartToken *StandartTokenSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandartToken.Contract.TransferOwnership(&_StandartToken.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_StandartToken *StandartTokenTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _StandartToken.Contract.TransferOwnership(&_StandartToken.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_StandartToken *StandartTokenTransactor) Withdraw(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StandartToken.contract.Transact(opts, "withdraw", account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_StandartToken *StandartTokenSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _StandartToken.Contract.Withdraw(&_StandartToken.TransactOpts, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_StandartToken *StandartTokenTransactorSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _StandartToken.Contract.Withdraw(&_StandartToken.TransactOpts, account)
}

// StandartTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StandartToken contract.
type StandartTokenApprovalIterator struct {
	Event *StandartTokenApproval // Event containing the contract specifics and raw log

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
func (it *StandartTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandartTokenApproval)
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
		it.Event = new(StandartTokenApproval)
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
func (it *StandartTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandartTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandartTokenApproval represents a Approval event raised by the StandartToken contract.
type StandartTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandartToken *StandartTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StandartTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandartToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StandartTokenApprovalIterator{contract: _StandartToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandartToken *StandartTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StandartTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StandartToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandartTokenApproval)
				if err := _StandartToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StandartToken *StandartTokenFilterer) ParseApproval(log types.Log) (*StandartTokenApproval, error) {
	event := new(StandartTokenApproval)
	if err := _StandartToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandartTokenDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the StandartToken contract.
type StandartTokenDeprecateIterator struct {
	Event *StandartTokenDeprecate // Event containing the contract specifics and raw log

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
func (it *StandartTokenDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandartTokenDeprecate)
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
		it.Event = new(StandartTokenDeprecate)
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
func (it *StandartTokenDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandartTokenDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandartTokenDeprecate represents a Deprecate event raised by the StandartToken contract.
type StandartTokenDeprecate struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_StandartToken *StandartTokenFilterer) FilterDeprecate(opts *bind.FilterOpts, account []common.Address) (*StandartTokenDeprecateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StandartToken.contract.FilterLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return &StandartTokenDeprecateIterator{contract: _StandartToken.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_StandartToken *StandartTokenFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *StandartTokenDeprecate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _StandartToken.contract.WatchLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandartTokenDeprecate)
				if err := _StandartToken.contract.UnpackLog(event, "Deprecate", log); err != nil {
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

// ParseDeprecate is a log parse operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_StandartToken *StandartTokenFilterer) ParseDeprecate(log types.Log) (*StandartTokenDeprecate, error) {
	event := new(StandartTokenDeprecate)
	if err := _StandartToken.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandartTokenRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the StandartToken contract.
type StandartTokenRewardIterator struct {
	Event *StandartTokenReward // Event containing the contract specifics and raw log

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
func (it *StandartTokenRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandartTokenReward)
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
		it.Event = new(StandartTokenReward)
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
func (it *StandartTokenRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandartTokenRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandartTokenReward represents a Reward event raised by the StandartToken contract.
type StandartTokenReward struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_StandartToken *StandartTokenFilterer) FilterReward(opts *bind.FilterOpts) (*StandartTokenRewardIterator, error) {

	logs, sub, err := _StandartToken.contract.FilterLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return &StandartTokenRewardIterator{contract: _StandartToken.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_StandartToken *StandartTokenFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *StandartTokenReward) (event.Subscription, error) {

	logs, sub, err := _StandartToken.contract.WatchLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandartTokenReward)
				if err := _StandartToken.contract.UnpackLog(event, "Reward", log); err != nil {
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

// ParseReward is a log parse operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_StandartToken *StandartTokenFilterer) ParseReward(log types.Log) (*StandartTokenReward, error) {
	event := new(StandartTokenReward)
	if err := _StandartToken.contract.UnpackLog(event, "Reward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// StandartTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StandartToken contract.
type StandartTokenTransferIterator struct {
	Event *StandartTokenTransfer // Event containing the contract specifics and raw log

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
func (it *StandartTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StandartTokenTransfer)
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
		it.Event = new(StandartTokenTransfer)
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
func (it *StandartTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StandartTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StandartTokenTransfer represents a Transfer event raised by the StandartToken contract.
type StandartTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandartToken *StandartTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StandartTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandartToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StandartTokenTransferIterator{contract: _StandartToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandartToken *StandartTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StandartTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StandartToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StandartTokenTransfer)
				if err := _StandartToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StandartToken *StandartTokenFilterer) ParseTransfer(log types.Log) (*StandartTokenTransfer, error) {
	event := new(StandartTokenTransfer)
	if err := _StandartToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
