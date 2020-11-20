// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ibport

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

// IBPortABI is the input ABI used to generate the binding from.
const IBPortABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_nebula\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAdress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"RequestCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"attachValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"receiver\",\"type\":\"bytes32\"}],\"name\":\"createTransferUnwrapRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nebula\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rqId\",\"type\":\"uint256\"}],\"name\":\"nextRq\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rqId\",\"type\":\"uint256\"}],\"name\":\"prevRq\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"requestsQueue\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"first\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"last\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"swapStatus\",\"outputs\":[{\"internalType\":\"enumIBPort.RequestStatus\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"internalType\":\"contractUSDN\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferTokenOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"unwrapRequests\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"homeAddress\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"foreignAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IBPortFuncSigs maps the 4-byte function signature to its string representation.
var IBPortFuncSigs = map[string]string{
	"cc32a151": "attachValue(bytes)",
	"ab9a9916": "createTransferUnwrapRequest(uint256,bytes32)",
	"4ecde849": "nebula()",
	"29db9a47": "nextRq(uint256)",
	"a6f3ade9": "prevRq(uint256)",
	"56dcda94": "requestsQueue()",
	"0872512b": "swapStatus(uint256)",
	"fc0c546a": "token()",
	"f2fde38b": "transferOwnership(address)",
	"21e6b53d": "transferTokenOwnership(address)",
	"d99c2a72": "unwrapRequests(uint256)",
}

// IBPortBin is the compiled bytecode used for deploying new contracts.
var IBPortBin = "0x608060405234801561001057600080fd5b50604051610fd9380380610fd98339818101604052604081101561003357600080fd5b50805160209091015160008054336001600160a01b03199182168117835560018054831690911790556002805482166001600160a01b039586161790556003805490911693909216929092179055610f4890819061009190396000f3fe608060405234801561001057600080fd5b50600436106100a95760003560e01c8063a6f3ade911610071578063a6f3ade91461018b578063ab9a9916146101a8578063cc32a151146101cb578063d99c2a721461023b578063f2fde38b14610280578063fc0c546a146102a6576100a9565b80630872512b146100ae57806321e6b53d146100ef57806329db9a47146101175780634ecde8491461014657806356dcda941461016a575b600080fd5b6100cb600480360360208110156100c457600080fd5b50356102ae565b604051808260048111156100db57fe5b60ff16815260200191505060405180910390f35b6101156004803603602081101561010557600080fd5b50356001600160a01b03166102c3565b005b6101346004803603602081101561012d57600080fd5b503561038a565b60408051918252519081900360200190f35b61014e61039c565b604080516001600160a01b039092168252519081900360200190f35b6101726103ab565b6040805192835260208301919091528051918290030190f35b610134600480360360208110156101a157600080fd5b50356103b4565b610115600480360360408110156101be57600080fd5b50803590602001356103c6565b610115600480360360208110156101e157600080fd5b8101906020810181356401000000008111156101fc57600080fd5b82018360208201111561020e57600080fd5b8035906020019184600183028401116401000000008311171561023057600080fd5b5090925090506106a8565b6102586004803603602081101561025157600080fd5b5035610937565b604080516001600160a01b039094168452602084019290925282820152519081900360600190f35b6101156004803603602081101561029657600080fd5b50356001600160a01b0316610962565b61014e610a27565b60046020526000908152604090205460ff1681565b6000546001600160a01b03163314806102e657506001546001600160a01b031633145b6103215760405162461bcd60e51b8152600401808060200182810382526029815260200180610eea6029913960400191505060405180910390fd5b6003546040805163f2fde38b60e01b81526001600160a01b0384811660048301529151919092169163f2fde38b91602480830192600092919082900301818387803b15801561036f57600080fd5b505af1158015610383573d6000803e3d6000fd5b5050505050565b60009081526008602052604090205490565b6002546001600160a01b031681565b60065460075482565b60009081526009602052604090205490565b6040805133606081901b602080840191909152603483018590524360548401526074808401879052845180850390910181526094840180865281519183019190912060f48501865283825260b4850187815260d4909501888152600082815260058552878120935184546001600160a01b0319166001600160a01b03918216178555965160018086019190915591516002909401939093556004808552878420805460ff191690921790915560035487516370a0823160e01b8152918201959095529551909591946104e79416926370a082319260248082019391829003018186803b1580156104b557600080fd5b505afa1580156104c9573d6000803e3d6000fd5b505050506040513d60208110156104df57600080fd5b505185610a36565b600354604080516351cff8d960e01b815233600482015290519293506001600160a01b03909116916351cff8d9916024808201926020929091908290030181600087803b15801561053757600080fd5b505af115801561054b573d6000803e3d6000fd5b505050506040513d602081101561056157600080fd5b505080156105ea57600354604080516311f9fbc960e21b81523360048201526024810184905290516001600160a01b03909216916347e7ef24916044808201926020929091908290030181600087803b1580156105bd57600080fd5b505af11580156105d1573d6000803e3d6000fd5b505050506040513d60208110156105e757600080fd5b50505b60408051632941b65560e21b81526006600482015260248101849052905173__$81d2f02810665af78277ac0ae2c955d74f$__9163a506d954916044808301926000929190829003018186803b15801561064357600080fd5b505af4158015610657573d6000803e3d6000fd5b5050604080518581523360208201528082018790526060810188905290517f78e1c38f7bce169c7cf026c9115bab62243678331df819e47ba8f2cd48ba259b9350908190036080019150a150505050565b6002546001600160a01b031633146106f7576040805162461bcd60e51b815260206004820152600d60248201526c1858d8d95cdcc819195b9a5959609a1b604482015290519081900360640190fd5b60005b8181101561093257600083838381811061071057fe5b6001909401936001600160f81b031992013591909116915050606d60f81b81141561083057600061077b85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525087925060209150610a819050565b905060208301925060006107c986868080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525088925060209150610a819050565b9050602084019350600061081487878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250899250610ac2915050565b9050601485019450610827838383610ad0565b5050505061092d565b6001600160f81b03198116606360f81b14156108f157600061088c85858080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525087925060209150610a819050565b905060208301925060006108d786868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250889250610c6a915050565b90506001840193506108e98282610d1c565b50505061092d565b6040805162461bcd60e51b815260206004820152600c60248201526b696e76616c6964206461746160a01b604482015290519081900360640190fd5b6106fa565b505050565b6005602052600090815260409020805460018201546002909201546001600160a01b03909116919083565b6000546001600160a01b031633148061098557506001546001600160a01b031633145b6109c05760405162461bcd60e51b8152600401808060200182810382526029815260200180610eea6029913960400191505060405180910390fd5b6001600160a01b038116610a055760405162461bcd60e51b8152600401808060200182810382526026815260200180610ec46026913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b6003546001600160a01b031681565b6000610a7883836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f770000815250610e2c565b90505b92915050565b600080835b838501811015610ab957858181518110610a9c57fe5b60209101015160f81c610100929092029190910190600101610a86565b50949350505050565b6000610a7883836014610a81565b600083815260046020819052604082205460ff1690811115610aee57fe5b14610b39576040805162461bcd60e51b8152602060048201526016602482015275696e76616c696420726571756573742073746174757360501b604482015290519081900360640190fd5b6001600160a01b038116301415610bca576003546040805163534a7e1d60e11b81526004810185905290516001600160a01b039092169163a694fc3a916024808201926020929091908290030181600087803b158015610b9857600080fd5b505af1158015610bac573d6000803e3d6000fd5b505050506040513d6020811015610bc257600080fd5b50610c4d9050565b600354604080516311f9fbc960e21b81526001600160a01b03848116600483015260248201869052915191909216916347e7ef249160448083019260209291908290030181600087803b158015610c2057600080fd5b505af1158015610c34573d6000803e3d6000fd5b505050506040513d6020811015610c4a57600080fd5b50505b50506000908152600460205260409020805460ff19166003179055565b600080838381518110610c7957fe5b016020015160f81c905080610c92576000915050610a7b565b8060011415610ca5576001915050610a7b565b8060021415610cb8576002915050610a7b565b8060031415610ccb576003915050610a7b565b8060041415610cde576004915050610a7b565b6040805162461bcd60e51b815260206004820152600e60248201526d696e76616c69642073746174757360901b604482015290519081900360640190fd5b600160008381526004602081905260409091205460ff1690811115610d3d57fe5b14610d88576040805162461bcd60e51b8152602060048201526016602482015275696e76616c696420726571756573742073746174757360501b604482015290519081900360640190fd5b60008281526004602081905260409091208054839260ff19909116906001908490811115610db257fe5b021790555060408051639d6ad84b60e01b81526006600482015260248101849052905173__$81d2f02810665af78277ac0ae2c955d74f$__91639d6ad84b916044808301926000929190829003018186803b158015610e1057600080fd5b505af4158015610e24573d6000803e3d6000fd5b505050505050565b60008184841115610ebb5760405162461bcd60e51b81526004018080602001828103825283818151815260200191508051906020019080838360005b83811015610e80578181015183820152602001610e68565b50505050905090810190601f168015610ead5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50505090039056fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f20616464726573734f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572206f722061646d696ea2646970667358221220a4e8f2017a57da8c564e76d8f0f4cbd7915738c7f9f2fe04b4dc94027691740364736f6c634300060b0033"

// DeployIBPort deploys a new Ethereum contract, binding an instance of IBPort to it.
func DeployIBPort(auth *bind.TransactOpts, backend bind.ContractBackend, _nebula common.Address, _tokenAdress common.Address) (common.Address, *types.Transaction, *IBPort, error) {
	parsed, err := abi.JSON(strings.NewReader(IBPortABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	queueLibAddr, _, _, _ := DeployQueueLib(auth, backend)
	IBPortBin = strings.Replace(IBPortBin, "__$81d2f02810665af78277ac0ae2c955d74f$__", queueLibAddr.String()[2:], -1)

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(IBPortBin), backend, _nebula, _tokenAdress)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &IBPort{IBPortCaller: IBPortCaller{contract: contract}, IBPortTransactor: IBPortTransactor{contract: contract}, IBPortFilterer: IBPortFilterer{contract: contract}}, nil
}

// IBPort is an auto generated Go binding around an Ethereum contract.
type IBPort struct {
	IBPortCaller     // Read-only binding to the contract
	IBPortTransactor // Write-only binding to the contract
	IBPortFilterer   // Log filterer for contract events
}

// IBPortCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBPortCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBPortTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBPortTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBPortFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBPortFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBPortSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBPortSession struct {
	Contract     *IBPort           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBPortCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBPortCallerSession struct {
	Contract *IBPortCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBPortTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBPortTransactorSession struct {
	Contract     *IBPortTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBPortRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBPortRaw struct {
	Contract *IBPort // Generic contract binding to access the raw methods on
}

// IBPortCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBPortCallerRaw struct {
	Contract *IBPortCaller // Generic read-only contract binding to access the raw methods on
}

// IBPortTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBPortTransactorRaw struct {
	Contract *IBPortTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBPort creates a new instance of IBPort, bound to a specific deployed contract.
func NewIBPort(address common.Address, backend bind.ContractBackend) (*IBPort, error) {
	contract, err := bindIBPort(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBPort{IBPortCaller: IBPortCaller{contract: contract}, IBPortTransactor: IBPortTransactor{contract: contract}, IBPortFilterer: IBPortFilterer{contract: contract}}, nil
}

// NewIBPortCaller creates a new read-only instance of IBPort, bound to a specific deployed contract.
func NewIBPortCaller(address common.Address, caller bind.ContractCaller) (*IBPortCaller, error) {
	contract, err := bindIBPort(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBPortCaller{contract: contract}, nil
}

// NewIBPortTransactor creates a new write-only instance of IBPort, bound to a specific deployed contract.
func NewIBPortTransactor(address common.Address, transactor bind.ContractTransactor) (*IBPortTransactor, error) {
	contract, err := bindIBPort(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBPortTransactor{contract: contract}, nil
}

// NewIBPortFilterer creates a new log filterer instance of IBPort, bound to a specific deployed contract.
func NewIBPortFilterer(address common.Address, filterer bind.ContractFilterer) (*IBPortFilterer, error) {
	contract, err := bindIBPort(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBPortFilterer{contract: contract}, nil
}

// bindIBPort binds a generic wrapper to an already deployed contract.
func bindIBPort(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBPortABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBPort *IBPortRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBPort.Contract.IBPortCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBPort *IBPortRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBPort.Contract.IBPortTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBPort *IBPortRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBPort.Contract.IBPortTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBPort *IBPortCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IBPort.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBPort *IBPortTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBPort.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBPort *IBPortTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBPort.Contract.contract.Transact(opts, method, params...)
}

// Nebula is a free data retrieval call binding the contract method 0x4ecde849.
//
// Solidity: function nebula() view returns(address)
func (_IBPort *IBPortCaller) Nebula(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "nebula")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nebula is a free data retrieval call binding the contract method 0x4ecde849.
//
// Solidity: function nebula() view returns(address)
func (_IBPort *IBPortSession) Nebula() (common.Address, error) {
	return _IBPort.Contract.Nebula(&_IBPort.CallOpts)
}

// Nebula is a free data retrieval call binding the contract method 0x4ecde849.
//
// Solidity: function nebula() view returns(address)
func (_IBPort *IBPortCallerSession) Nebula() (common.Address, error) {
	return _IBPort.Contract.Nebula(&_IBPort.CallOpts)
}

// NextRq is a free data retrieval call binding the contract method 0x29db9a47.
//
// Solidity: function nextRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortCaller) NextRq(opts *bind.CallOpts, rqId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "nextRq", rqId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextRq is a free data retrieval call binding the contract method 0x29db9a47.
//
// Solidity: function nextRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortSession) NextRq(rqId *big.Int) (*big.Int, error) {
	return _IBPort.Contract.NextRq(&_IBPort.CallOpts, rqId)
}

// NextRq is a free data retrieval call binding the contract method 0x29db9a47.
//
// Solidity: function nextRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortCallerSession) NextRq(rqId *big.Int) (*big.Int, error) {
	return _IBPort.Contract.NextRq(&_IBPort.CallOpts, rqId)
}

// PrevRq is a free data retrieval call binding the contract method 0xa6f3ade9.
//
// Solidity: function prevRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortCaller) PrevRq(opts *bind.CallOpts, rqId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "prevRq", rqId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrevRq is a free data retrieval call binding the contract method 0xa6f3ade9.
//
// Solidity: function prevRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortSession) PrevRq(rqId *big.Int) (*big.Int, error) {
	return _IBPort.Contract.PrevRq(&_IBPort.CallOpts, rqId)
}

// PrevRq is a free data retrieval call binding the contract method 0xa6f3ade9.
//
// Solidity: function prevRq(uint256 rqId) view returns(uint256)
func (_IBPort *IBPortCallerSession) PrevRq(rqId *big.Int) (*big.Int, error) {
	return _IBPort.Contract.PrevRq(&_IBPort.CallOpts, rqId)
}

// RequestsQueue is a free data retrieval call binding the contract method 0x56dcda94.
//
// Solidity: function requestsQueue() view returns(bytes32 first, bytes32 last)
func (_IBPort *IBPortCaller) RequestsQueue(opts *bind.CallOpts) (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "requestsQueue")

	outstruct := new(struct {
		First [32]byte
		Last  [32]byte
	})

	outstruct.First = out[0].([32]byte)
	outstruct.Last = out[1].([32]byte)

	return *outstruct, err

}

// RequestsQueue is a free data retrieval call binding the contract method 0x56dcda94.
//
// Solidity: function requestsQueue() view returns(bytes32 first, bytes32 last)
func (_IBPort *IBPortSession) RequestsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _IBPort.Contract.RequestsQueue(&_IBPort.CallOpts)
}

// RequestsQueue is a free data retrieval call binding the contract method 0x56dcda94.
//
// Solidity: function requestsQueue() view returns(bytes32 first, bytes32 last)
func (_IBPort *IBPortCallerSession) RequestsQueue() (struct {
	First [32]byte
	Last  [32]byte
}, error) {
	return _IBPort.Contract.RequestsQueue(&_IBPort.CallOpts)
}

// SwapStatus is a free data retrieval call binding the contract method 0x0872512b.
//
// Solidity: function swapStatus(uint256 ) view returns(uint8)
func (_IBPort *IBPortCaller) SwapStatus(opts *bind.CallOpts, arg0 *big.Int) (uint8, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "swapStatus", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// SwapStatus is a free data retrieval call binding the contract method 0x0872512b.
//
// Solidity: function swapStatus(uint256 ) view returns(uint8)
func (_IBPort *IBPortSession) SwapStatus(arg0 *big.Int) (uint8, error) {
	return _IBPort.Contract.SwapStatus(&_IBPort.CallOpts, arg0)
}

// SwapStatus is a free data retrieval call binding the contract method 0x0872512b.
//
// Solidity: function swapStatus(uint256 ) view returns(uint8)
func (_IBPort *IBPortCallerSession) SwapStatus(arg0 *big.Int) (uint8, error) {
	return _IBPort.Contract.SwapStatus(&_IBPort.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IBPort *IBPortCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IBPort *IBPortSession) Token() (common.Address, error) {
	return _IBPort.Contract.Token(&_IBPort.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_IBPort *IBPortCallerSession) Token() (common.Address, error) {
	return _IBPort.Contract.Token(&_IBPort.CallOpts)
}

// UnwrapRequests is a free data retrieval call binding the contract method 0xd99c2a72.
//
// Solidity: function unwrapRequests(uint256 ) view returns(address homeAddress, bytes32 foreignAddress, uint256 amount)
func (_IBPort *IBPortCaller) UnwrapRequests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HomeAddress    common.Address
	ForeignAddress [32]byte
	Amount         *big.Int
}, error) {
	var out []interface{}
	err := _IBPort.contract.Call(opts, &out, "unwrapRequests", arg0)

	outstruct := new(struct {
		HomeAddress    common.Address
		ForeignAddress [32]byte
		Amount         *big.Int
	})

	outstruct.HomeAddress = out[0].(common.Address)
	outstruct.ForeignAddress = out[1].([32]byte)
	outstruct.Amount = out[2].(*big.Int)

	return *outstruct, err

}

// UnwrapRequests is a free data retrieval call binding the contract method 0xd99c2a72.
//
// Solidity: function unwrapRequests(uint256 ) view returns(address homeAddress, bytes32 foreignAddress, uint256 amount)
func (_IBPort *IBPortSession) UnwrapRequests(arg0 *big.Int) (struct {
	HomeAddress    common.Address
	ForeignAddress [32]byte
	Amount         *big.Int
}, error) {
	return _IBPort.Contract.UnwrapRequests(&_IBPort.CallOpts, arg0)
}

// UnwrapRequests is a free data retrieval call binding the contract method 0xd99c2a72.
//
// Solidity: function unwrapRequests(uint256 ) view returns(address homeAddress, bytes32 foreignAddress, uint256 amount)
func (_IBPort *IBPortCallerSession) UnwrapRequests(arg0 *big.Int) (struct {
	HomeAddress    common.Address
	ForeignAddress [32]byte
	Amount         *big.Int
}, error) {
	return _IBPort.Contract.UnwrapRequests(&_IBPort.CallOpts, arg0)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_IBPort *IBPortTransactor) AttachValue(opts *bind.TransactOpts, value []byte) (*types.Transaction, error) {
	return _IBPort.contract.Transact(opts, "attachValue", value)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_IBPort *IBPortSession) AttachValue(value []byte) (*types.Transaction, error) {
	return _IBPort.Contract.AttachValue(&_IBPort.TransactOpts, value)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_IBPort *IBPortTransactorSession) AttachValue(value []byte) (*types.Transaction, error) {
	return _IBPort.Contract.AttachValue(&_IBPort.TransactOpts, value)
}

// CreateTransferUnwrapRequest is a paid mutator transaction binding the contract method 0xab9a9916.
//
// Solidity: function createTransferUnwrapRequest(uint256 amount, bytes32 receiver) returns()
func (_IBPort *IBPortTransactor) CreateTransferUnwrapRequest(opts *bind.TransactOpts, amount *big.Int, receiver [32]byte) (*types.Transaction, error) {
	return _IBPort.contract.Transact(opts, "createTransferUnwrapRequest", amount, receiver)
}

// CreateTransferUnwrapRequest is a paid mutator transaction binding the contract method 0xab9a9916.
//
// Solidity: function createTransferUnwrapRequest(uint256 amount, bytes32 receiver) returns()
func (_IBPort *IBPortSession) CreateTransferUnwrapRequest(amount *big.Int, receiver [32]byte) (*types.Transaction, error) {
	return _IBPort.Contract.CreateTransferUnwrapRequest(&_IBPort.TransactOpts, amount, receiver)
}

// CreateTransferUnwrapRequest is a paid mutator transaction binding the contract method 0xab9a9916.
//
// Solidity: function createTransferUnwrapRequest(uint256 amount, bytes32 receiver) returns()
func (_IBPort *IBPortTransactorSession) CreateTransferUnwrapRequest(amount *big.Int, receiver [32]byte) (*types.Transaction, error) {
	return _IBPort.Contract.CreateTransferUnwrapRequest(&_IBPort.TransactOpts, amount, receiver)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IBPort *IBPortTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IBPort *IBPortSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.Contract.TransferOwnership(&_IBPort.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_IBPort *IBPortTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.Contract.TransferOwnership(&_IBPort.TransactOpts, newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address newOwner) returns()
func (_IBPort *IBPortTransactor) TransferTokenOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.contract.Transact(opts, "transferTokenOwnership", newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address newOwner) returns()
func (_IBPort *IBPortSession) TransferTokenOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.Contract.TransferTokenOwnership(&_IBPort.TransactOpts, newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address newOwner) returns()
func (_IBPort *IBPortTransactorSession) TransferTokenOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _IBPort.Contract.TransferTokenOwnership(&_IBPort.TransactOpts, newOwner)
}

// IBPortRequestCreatedIterator is returned from FilterRequestCreated and is used to iterate over the raw logs and unpacked data for RequestCreated events raised by the IBPort contract.
type IBPortRequestCreatedIterator struct {
	Event *IBPortRequestCreated // Event containing the contract specifics and raw log

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
func (it *IBPortRequestCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IBPortRequestCreated)
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
		it.Event = new(IBPortRequestCreated)
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
func (it *IBPortRequestCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IBPortRequestCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IBPortRequestCreated represents a RequestCreated event raised by the IBPort contract.
type IBPortRequestCreated struct {
	Arg0 *big.Int
	Arg1 common.Address
	Arg2 [32]byte
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterRequestCreated is a free log retrieval operation binding the contract event 0x78e1c38f7bce169c7cf026c9115bab62243678331df819e47ba8f2cd48ba259b.
//
// Solidity: event RequestCreated(uint256 arg0, address arg1, bytes32 arg2, uint256 arg3)
func (_IBPort *IBPortFilterer) FilterRequestCreated(opts *bind.FilterOpts) (*IBPortRequestCreatedIterator, error) {

	logs, sub, err := _IBPort.contract.FilterLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return &IBPortRequestCreatedIterator{contract: _IBPort.contract, event: "RequestCreated", logs: logs, sub: sub}, nil
}

// WatchRequestCreated is a free log subscription operation binding the contract event 0x78e1c38f7bce169c7cf026c9115bab62243678331df819e47ba8f2cd48ba259b.
//
// Solidity: event RequestCreated(uint256 arg0, address arg1, bytes32 arg2, uint256 arg3)
func (_IBPort *IBPortFilterer) WatchRequestCreated(opts *bind.WatchOpts, sink chan<- *IBPortRequestCreated) (event.Subscription, error) {

	logs, sub, err := _IBPort.contract.WatchLogs(opts, "RequestCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IBPortRequestCreated)
				if err := _IBPort.contract.UnpackLog(event, "RequestCreated", log); err != nil {
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

// ParseRequestCreated is a log parse operation binding the contract event 0x78e1c38f7bce169c7cf026c9115bab62243678331df819e47ba8f2cd48ba259b.
//
// Solidity: event RequestCreated(uint256 arg0, address arg1, bytes32 arg2, uint256 arg3)
func (_IBPort *IBPortFilterer) ParseRequestCreated(log types.Log) (*IBPortRequestCreated, error) {
	event := new(IBPortRequestCreated)
	if err := _IBPort.contract.UnpackLog(event, "RequestCreated", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ISubscriberBytesABI is the input ABI used to generate the binding from.
const ISubscriberBytesABI = "[{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"value\",\"type\":\"bytes\"}],\"name\":\"attachValue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ISubscriberBytesFuncSigs maps the 4-byte function signature to its string representation.
var ISubscriberBytesFuncSigs = map[string]string{
	"cc32a151": "attachValue(bytes)",
}

// ISubscriberBytes is an auto generated Go binding around an Ethereum contract.
type ISubscriberBytes struct {
	ISubscriberBytesCaller     // Read-only binding to the contract
	ISubscriberBytesTransactor // Write-only binding to the contract
	ISubscriberBytesFilterer   // Log filterer for contract events
}

// ISubscriberBytesCaller is an auto generated read-only Go binding around an Ethereum contract.
type ISubscriberBytesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriberBytesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ISubscriberBytesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriberBytesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ISubscriberBytesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ISubscriberBytesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ISubscriberBytesSession struct {
	Contract     *ISubscriberBytes // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ISubscriberBytesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ISubscriberBytesCallerSession struct {
	Contract *ISubscriberBytesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ISubscriberBytesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ISubscriberBytesTransactorSession struct {
	Contract     *ISubscriberBytesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ISubscriberBytesRaw is an auto generated low-level Go binding around an Ethereum contract.
type ISubscriberBytesRaw struct {
	Contract *ISubscriberBytes // Generic contract binding to access the raw methods on
}

// ISubscriberBytesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ISubscriberBytesCallerRaw struct {
	Contract *ISubscriberBytesCaller // Generic read-only contract binding to access the raw methods on
}

// ISubscriberBytesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ISubscriberBytesTransactorRaw struct {
	Contract *ISubscriberBytesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewISubscriberBytes creates a new instance of ISubscriberBytes, bound to a specific deployed contract.
func NewISubscriberBytes(address common.Address, backend bind.ContractBackend) (*ISubscriberBytes, error) {
	contract, err := bindISubscriberBytes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ISubscriberBytes{ISubscriberBytesCaller: ISubscriberBytesCaller{contract: contract}, ISubscriberBytesTransactor: ISubscriberBytesTransactor{contract: contract}, ISubscriberBytesFilterer: ISubscriberBytesFilterer{contract: contract}}, nil
}

// NewISubscriberBytesCaller creates a new read-only instance of ISubscriberBytes, bound to a specific deployed contract.
func NewISubscriberBytesCaller(address common.Address, caller bind.ContractCaller) (*ISubscriberBytesCaller, error) {
	contract, err := bindISubscriberBytes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriberBytesCaller{contract: contract}, nil
}

// NewISubscriberBytesTransactor creates a new write-only instance of ISubscriberBytes, bound to a specific deployed contract.
func NewISubscriberBytesTransactor(address common.Address, transactor bind.ContractTransactor) (*ISubscriberBytesTransactor, error) {
	contract, err := bindISubscriberBytes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ISubscriberBytesTransactor{contract: contract}, nil
}

// NewISubscriberBytesFilterer creates a new log filterer instance of ISubscriberBytes, bound to a specific deployed contract.
func NewISubscriberBytesFilterer(address common.Address, filterer bind.ContractFilterer) (*ISubscriberBytesFilterer, error) {
	contract, err := bindISubscriberBytes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ISubscriberBytesFilterer{contract: contract}, nil
}

// bindISubscriberBytes binds a generic wrapper to an already deployed contract.
func bindISubscriberBytes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ISubscriberBytesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscriberBytes *ISubscriberBytesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISubscriberBytes.Contract.ISubscriberBytesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscriberBytes *ISubscriberBytesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.ISubscriberBytesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscriberBytes *ISubscriberBytesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.ISubscriberBytesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ISubscriberBytes *ISubscriberBytesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ISubscriberBytes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ISubscriberBytes *ISubscriberBytesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ISubscriberBytes *ISubscriberBytesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.contract.Transact(opts, method, params...)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_ISubscriberBytes *ISubscriberBytesTransactor) AttachValue(opts *bind.TransactOpts, value []byte) (*types.Transaction, error) {
	return _ISubscriberBytes.contract.Transact(opts, "attachValue", value)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_ISubscriberBytes *ISubscriberBytesSession) AttachValue(value []byte) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.AttachValue(&_ISubscriberBytes.TransactOpts, value)
}

// AttachValue is a paid mutator transaction binding the contract method 0xcc32a151.
//
// Solidity: function attachValue(bytes value) returns()
func (_ISubscriberBytes *ISubscriberBytesTransactorSession) AttachValue(value []byte) (*types.Transaction, error) {
	return _ISubscriberBytes.Contract.AttachValue(&_ISubscriberBytes.TransactOpts, value)
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

// QueueLibABI is the input ABI used to generate the binding from.
const QueueLibABI = "[]"

// QueueLibFuncSigs maps the 4-byte function signature to its string representation.
var QueueLibFuncSigs = map[string]string{
	"9d6ad84b": "drop(QueueLib.Queue storage,bytes32)",
	"a506d954": "push(QueueLib.Queue storage,bytes32)",
}

// QueueLibBin is the compiled bytecode used for deploying new contracts.
var QueueLibBin = "0x610198610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600436106100405760003560e01c80639d6ad84b14610045578063a506d95414610077575b600080fd5b81801561005157600080fd5b506100756004803603604081101561006857600080fd5b50803590602001356100a7565b005b81801561008357600080fd5b506100756004803603604081101561009a57600080fd5b5080359060200135610114565b6000818152600383016020908152604080832054600286019092529091205481156100e457600082815260028501602052604090208190556100e8565b8084555b8015610106576000818152600385016020526040902082905561010e565b600184018290555b50505050565b8154610129578082556001820181905561015e565b600182018054600090815260028401602081815260408084208690558454868552600388018352818520559190528120558190555b505056fea26469706673582212200952320fe013b2753ed02f679ed5335b6ae6330b07a2c6d29803ddb39cfdb8dc64736f6c634300060b0033"

// DeployQueueLib deploys a new Ethereum contract, binding an instance of QueueLib to it.
func DeployQueueLib(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *QueueLib, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(QueueLibBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// QueueLib is an auto generated Go binding around an Ethereum contract.
type QueueLib struct {
	QueueLibCaller     // Read-only binding to the contract
	QueueLibTransactor // Write-only binding to the contract
	QueueLibFilterer   // Log filterer for contract events
}

// QueueLibCaller is an auto generated read-only Go binding around an Ethereum contract.
type QueueLibCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibTransactor is an auto generated write-only Go binding around an Ethereum contract.
type QueueLibTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type QueueLibFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// QueueLibSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type QueueLibSession struct {
	Contract     *QueueLib         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// QueueLibCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type QueueLibCallerSession struct {
	Contract *QueueLibCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// QueueLibTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type QueueLibTransactorSession struct {
	Contract     *QueueLibTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// QueueLibRaw is an auto generated low-level Go binding around an Ethereum contract.
type QueueLibRaw struct {
	Contract *QueueLib // Generic contract binding to access the raw methods on
}

// QueueLibCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type QueueLibCallerRaw struct {
	Contract *QueueLibCaller // Generic read-only contract binding to access the raw methods on
}

// QueueLibTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type QueueLibTransactorRaw struct {
	Contract *QueueLibTransactor // Generic write-only contract binding to access the raw methods on
}

// NewQueueLib creates a new instance of QueueLib, bound to a specific deployed contract.
func NewQueueLib(address common.Address, backend bind.ContractBackend) (*QueueLib, error) {
	contract, err := bindQueueLib(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &QueueLib{QueueLibCaller: QueueLibCaller{contract: contract}, QueueLibTransactor: QueueLibTransactor{contract: contract}, QueueLibFilterer: QueueLibFilterer{contract: contract}}, nil
}

// NewQueueLibCaller creates a new read-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibCaller(address common.Address, caller bind.ContractCaller) (*QueueLibCaller, error) {
	contract, err := bindQueueLib(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibCaller{contract: contract}, nil
}

// NewQueueLibTransactor creates a new write-only instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibTransactor(address common.Address, transactor bind.ContractTransactor) (*QueueLibTransactor, error) {
	contract, err := bindQueueLib(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &QueueLibTransactor{contract: contract}, nil
}

// NewQueueLibFilterer creates a new log filterer instance of QueueLib, bound to a specific deployed contract.
func NewQueueLibFilterer(address common.Address, filterer bind.ContractFilterer) (*QueueLibFilterer, error) {
	contract, err := bindQueueLib(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &QueueLibFilterer{contract: contract}, nil
}

// bindQueueLib binds a generic wrapper to an already deployed contract.
func bindQueueLib(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(QueueLibABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.QueueLibCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.QueueLibTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_QueueLib *QueueLibCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _QueueLib.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_QueueLib *QueueLibTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_QueueLib *QueueLibTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _QueueLib.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220c26a9415a71cef479bf2ffa9c25fc875777128f4e62bca62badf4091bf22dab364736f6c634300060b0033"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
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

// USDNABI is the input ABI used to generate the binding from.
const USDNABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"Deprecate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Reward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"deposit\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"deprecate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reward\",\"type\":\"uint256\"}],\"name\":\"stake\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// USDNFuncSigs maps the 4-byte function signature to its string representation.
var USDNFuncSigs = map[string]string{
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

// USDNBin is the compiled bytecode used for deploying new contracts.
var USDNBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b03199081163390811783556001805490921617815560028054918201815590915264e8d4a510007f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace90910155611a30806100776000396000f3fe608060405234801561001057600080fd5b50600436106101005760003560e01c806351cff8d911610097578063a694fc3a11610066578063a694fc3a14610312578063a9059cbb1461032f578063dd62ed3e1461035b578063f2fde38b1461038957610100565b806351cff8d91461029257806370a08231146102b857806395d89b41146102de578063a457c2d7146102e657610100565b806323b872dd116100d357806323b872dd146101e6578063313ce5671461021c578063395093511461023a57806347e7ef241461026657610100565b806306fdde0314610105578063095ea7b3146101825780630fcc0c28146101c257806318160ddd146101cc575b600080fd5b61010d6103af565b6040805160208082528351818301528351919283929083019185019080838360005b8381101561014757818101518382015260200161012f565b50505050905090810190601f1680156101745780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6101ae6004803603604081101561019857600080fd5b506001600160a01b0381351690602001356103dd565b604080519115158252519081900360200190f35b6101ca6103f3565b005b6101d4610491565b60408051918252519081900360200190f35b6101ae600480360360608110156101fc57600080fd5b506001600160a01b038135811691602081013590911690604001356104e2565b61022461056a565b6040805160ff9092168252519081900360200190f35b6101ae6004803603604081101561025057600080fd5b506001600160a01b03813516906020013561056f565b6101ae6004803603604081101561027c57600080fd5b506001600160a01b0381351690602001356105f5565b6101ae600480360360208110156102a857600080fd5b50356001600160a01b0316610928565b6101d4600480360360208110156102ce57600080fd5b50356001600160a01b0316610be3565b61010d610dee565b6101ae600480360360408110156102fc57600080fd5b506001600160a01b038135169060200135610e0e565b6101ae6004803603602081101561032857600080fd5b5035610e7e565b6101ae6004803603604081101561034557600080fd5b506001600160a01b038135169060200135611155565b6101d46004803603604081101561037157600080fd5b506001600160a01b0381358116916020013516611162565b6101ca6004803603602081101561039f57600080fd5b50356001600160a01b031661118d565b60408051808201909152601481527315dc985c1c19590813995d5d1c9a5b9bc81554d160621b602082015290565b60006103ea338484611252565b50600192915050565b6000546001600160a01b031633148061041657506001546001600160a01b031633145b6104515760405162461bcd60e51b815260040180806020018281038252602981526020018061188a6029913960400191505060405180910390fd5b6001805460ff60a01b1916600160a01b17905560405133907fcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e90600090a2565b600354600454600091908082018211156104dc5760405162461bcd60e51b81526004018080602001828103825260228152602001806119006022913960400191505060405180910390fd5b01905090565b60006104ef848484611387565b6001600160a01b0384166000908152600860209081526040808320338452909152902054808311156105525760405162461bcd60e51b81526004018080602001828103825260288152602001806118b36028913960400191505060405180910390fd5b61055f8533858403611252565b506001949350505050565b601290565b3360009081526008602090815260408083206001600160a01b03861684529091528120548281018111156105de576040805162461bcd60e51b81526020600482015260116024820152706164646974696f6e206f766572666c6f7760781b604482015290519081900360640190fd5b6105eb3385858401611252565b5060019392505050565b600080546001600160a01b031633148061061957506001546001600160a01b031633145b6106545760405162461bcd60e51b815260040180806020018281038252602981526020018061188a6029913960400191505060405180910390fd5b600154600160a01b900460ff161561069d5760405162461bcd60e51b81526004018080602001828103825260248152602001806119d76024913960400191505060405180910390fd5b600082116106e9576040805162461bcd60e51b81526020600482015260146024820152730616d6f756e742073686f756c64206265203e20360641b604482015290519081900360640190fd5b6001600160a01b038316610744576040805162461bcd60e51b815260206004820152601b60248201527f6465706f73697420746f20746865207a65726f20616464726573730000000000604482015290519081900360640190fd5b60045482810181111561079e576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f72206465706f736974000000604482015290519081900360640190fd5b8083016004556001600160a01b03841660009081526006602052604090205480610809576107cb85610be3565b6001600160a01b03861660009081526005602090815260408083209390935560025460078252838320600019909101905560069052208490556108ee565b6001600160a01b038516600090815260076020526040902054600254600019018114156108a957818583011015610887576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f72206465706f736974000000604482015290519081900360640190fd5b6001600160a01b038616600090815260066020526040902082860190556108ec565b6108b286610be3565b6001600160a01b03871660009081526005602090815260408083209390935560025460078252838320600019909101905560069052208590555b505b6040805185815290516001600160a01b038716916000916000805160206119498339815191529181900360200190a3506001949350505050565b600080546001600160a01b031633148061094c57506001546001600160a01b031633145b6109875760405162461bcd60e51b815260040180806020018281038252602981526020018061188a6029913960400191505060405180910390fd5b600154600160a01b900460ff16156109d05760405162461bcd60e51b81526004018080602001828103825260248152602001806119d76024913960400191505060405180910390fd5b6001600160a01b03821660009081526006602090815260408083205460079092529091205460025460001901811415610b2a576001600160a01b038416600090815260056020526040902054600354811115610a5d5760405162461bcd60e51b81526004018080602001828103825260258152602001806118db6025913960400191505060405180910390fd5b600380548290039055600454831115610aa75760405162461bcd60e51b81526004018080602001828103825260278152602001806119226027913960400191505060405180910390fd5b8260045403600481905550808382011015610af35760405162461bcd60e51b81526004018080602001828103825260308152602001806117936030913960400191505060405180910390fd5b60408051828501815290516000916001600160a01b038816916000805160206119498339815191529181900360200190a350610bb2565b6000610b3585610be3565b60035490915080821115610b7a5760405162461bcd60e51b81526004018080602001828103825260258152602001806118db6025913960400191505060405180910390fd5b8181036003556040805183815290516000916001600160a01b038916916000805160206119498339815191529181900360200190a350505b5050506001600160a01b0381166000908152600560209081526040808320839055600690915281205560015b919050565b6001600160a01b038116600090815260056020908152604080832054600690925282205481158015610c13575080155b15610c2357600092505050610bde565b6001600160a01b03841660009081526007602052604090205460025460001901811415610caa57828284011015610ca1576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722062616c616e6365000000604482015290519081900360640190fd5b50019050610bde565b81610d005760028054600091906000198101908110610cc557fe5b9060005260206000200154905060028281548110610cdf57fe5b906000526020600020015484820281610cf457fe5b04945050505050610bde565b600060028281548110610d0f57fe5b600091825260209091200154600280546000198101908110610d2d57fe5b9060005260206000200154850281610d4157fe5b049050600060028360010181548110610d5657fe5b600091825260209091200154600280546000198101908110610d7457fe5b9060005260206000200154850281610d8857fe5b049050818282011015610de2576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722062616c616e6365000000604482015290519081900360640190fd5b019350610bde92505050565b60408051808201909152600681526533bbaaa9a22760d11b602082015290565b3360009081526008602090815260408083206001600160a01b038616845290915281205480831115610e715760405162461bcd60e51b81526004018080602001828103825260258152602001806119b26025913960400191505060405180910390fd5b6105eb3385858403611252565b600080546001600160a01b0316331480610ea257506001546001600160a01b031633145b610edd5760405162461bcd60e51b815260040180806020018281038252602981526020018061188a6029913960400191505060405180910390fd5b600154600160a01b900460ff1615610f265760405162461bcd60e51b81526004018080602001828103825260248152602001806119d76024913960400191505060405180910390fd5b60008211610f72576040805162461bcd60e51b815260206004820152601460248201527307265776172642073686f756c64206265203e20360641b604482015290519081900360640190fd5b60035460045481610fbd576002805460018101825560009190915264e8d4a510007f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace909101556110c2565b60028054600091906000198101908110610fd357fe5b9060005260206000200154905060008364e8d4a51000870281610ff257fe5b0490508064e8d4a5100082011015611051576040805162461bcd60e51b815260206004820152601d60248201527f6164646974696f6e206f766572666c6f7720666f722070657263656e74000000604482015290519081900360640190fd5b64e8d4a51000818101906002908483028254600181018455600093845260209093209190049101558685018511156110ba5760405162461bcd60e51b815260040180806020018281038252602b81526020018061185f602b913960400191505060405180910390fd5b505050908301905b8181830110156111035760405162461bcd60e51b81526004018080602001828103825260228152602001806119006022913960400191505060405180910390fd5b8181016003556000600455600254604080519182526020820186905280517f45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da529281900390910190a15060019392505050565b60006103ea338484611387565b6001600160a01b03918216600090815260086020908152604080832093909416825291909152205490565b6000546001600160a01b03163314806111b057506001546001600160a01b031633145b6111eb5760405162461bcd60e51b815260040180806020018281038252602981526020018061188a6029913960400191505060405180910390fd5b6001600160a01b0381166112305760405162461bcd60e51b81526004018080602001828103825260268152602001806117c36026913960400191505060405180910390fd5b600080546001600160a01b0319166001600160a01b0392909216919091179055565b600154600160a01b900460ff161561129b5760405162461bcd60e51b81526004018080602001828103825260248152602001806119d76024913960400191505060405180910390fd5b6001600160a01b0383166112e05760405162461bcd60e51b815260040180806020018281038252602481526020018061198e6024913960400191505060405180910390fd5b6001600160a01b0382166113255760405162461bcd60e51b81526004018080602001828103825260228152602001806117e96022913960400191505060405180910390fd5b6001600160a01b03808416600081815260086020908152604080832094871680845294825291829020859055815185815291517f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259281900390910190a3505050565b600154600160a01b900460ff16156113d05760405162461bcd60e51b81526004018080602001828103825260248152602001806119d76024913960400191505060405180910390fd5b6000811161141c576040805162461bcd60e51b81526020600482015260146024820152730616d6f756e742073686f756c64206265203e20360641b604482015290519081900360640190fd5b6001600160a01b0383166114615760405162461bcd60e51b81526004018080602001828103825260258152602001806119696025913960400191505060405180910390fd5b6001600160a01b0382166114a65760405162461bcd60e51b81526004018080602001828103825260238152602001806117706023913960400191505060405180910390fd5b6001600160a01b038316600090815260066020908152604080832054600790925282205490918215806114df5750600254600019018214155b156115745760006114ef87610be3565b9050808511156115305760405162461bcd60e51b815260040180806020018281038252602681526020018061180b6026913960400191505060405180910390fd5b6001600160a01b0387166000908152600560209081526040808320938890039093556006815282822082905560025460079091529190206000199091019055611628565b82841161159f57506001600160a01b0385166000908152600660205260409020838303905582611628565b6001600160a01b0386166000908152600560205260409020548385038110156115f95760405162461bcd60e51b815260040180806020018281038252602681526020018061180b6026913960400191505060405180910390fd5b6001600160a01b0387166000908152600560209081526040808320878903909403909355600690529081205550815b6001600160a01b03851660009081526006602090815260408083205460079092529091205490935091508215806116655750600254600019018214155b156116fe57600061167586610be3565b905080818387030110156116ba5760405162461bcd60e51b815260040180806020018281038252602e815260200180611831602e913960400191505060405180910390fd5b6001600160a01b038616600090815260056020908152604080832085890394909401909355600254600782528383206000199091019055600690522081905561172e565b6001600160a01b038516600090815260056020908152604080832080548589030190556006909152902083820190555b846001600160a01b0316866001600160a01b0316600080516020611949833981519152866040518082815260200191505060405180910390a350505050505056fe45524332303a207472616e7366657220746f20746865207a65726f20616464726573736164646974696f6e206f766572666c6f7720666f7220746f74616c2062616c616e6365202b206f6c644465706f7369744f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a206164646974696f6e206f766572666c6f7720666f7220726563697069656e742062616c616e63656164646974696f6e206f766572666c6f7720666f7220746f74616c20737570706c79202b207265776172644f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572206f722061646d696e45524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e63657375627472616374696f6e206f766572666c6f7720666f7220746f74616c20737570706c796164646974696f6e206f766572666c6f7720666f7220746f74616c20737570706c797375627472616374696f6e206f766572666c6f7720666f72206c6971756964206465706f736974ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef45524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726f446570726563617465626c653a20636f6e74726163742069732064657072656361746564a2646970667358221220a9088ac48dce3ba78fcd0eb10cc26289a8f21287ddc1822c1ef92d21e91e196564736f6c634300060b0033"

// DeployUSDN deploys a new Ethereum contract, binding an instance of USDN to it.
func DeployUSDN(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *USDN, error) {
	parsed, err := abi.JSON(strings.NewReader(USDNABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(USDNBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &USDN{USDNCaller: USDNCaller{contract: contract}, USDNTransactor: USDNTransactor{contract: contract}, USDNFilterer: USDNFilterer{contract: contract}}, nil
}

// USDN is an auto generated Go binding around an Ethereum contract.
type USDN struct {
	USDNCaller     // Read-only binding to the contract
	USDNTransactor // Write-only binding to the contract
	USDNFilterer   // Log filterer for contract events
}

// USDNCaller is an auto generated read-only Go binding around an Ethereum contract.
type USDNCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDNTransactor is an auto generated write-only Go binding around an Ethereum contract.
type USDNTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDNFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type USDNFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// USDNSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type USDNSession struct {
	Contract     *USDN             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDNCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type USDNCallerSession struct {
	Contract *USDNCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// USDNTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type USDNTransactorSession struct {
	Contract     *USDNTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// USDNRaw is an auto generated low-level Go binding around an Ethereum contract.
type USDNRaw struct {
	Contract *USDN // Generic contract binding to access the raw methods on
}

// USDNCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type USDNCallerRaw struct {
	Contract *USDNCaller // Generic read-only contract binding to access the raw methods on
}

// USDNTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type USDNTransactorRaw struct {
	Contract *USDNTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUSDN creates a new instance of USDN, bound to a specific deployed contract.
func NewUSDN(address common.Address, backend bind.ContractBackend) (*USDN, error) {
	contract, err := bindUSDN(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &USDN{USDNCaller: USDNCaller{contract: contract}, USDNTransactor: USDNTransactor{contract: contract}, USDNFilterer: USDNFilterer{contract: contract}}, nil
}

// NewUSDNCaller creates a new read-only instance of USDN, bound to a specific deployed contract.
func NewUSDNCaller(address common.Address, caller bind.ContractCaller) (*USDNCaller, error) {
	contract, err := bindUSDN(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &USDNCaller{contract: contract}, nil
}

// NewUSDNTransactor creates a new write-only instance of USDN, bound to a specific deployed contract.
func NewUSDNTransactor(address common.Address, transactor bind.ContractTransactor) (*USDNTransactor, error) {
	contract, err := bindUSDN(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &USDNTransactor{contract: contract}, nil
}

// NewUSDNFilterer creates a new log filterer instance of USDN, bound to a specific deployed contract.
func NewUSDNFilterer(address common.Address, filterer bind.ContractFilterer) (*USDNFilterer, error) {
	contract, err := bindUSDN(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &USDNFilterer{contract: contract}, nil
}

// bindUSDN binds a generic wrapper to an already deployed contract.
func bindUSDN(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(USDNABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDN *USDNRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDN.Contract.USDNCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDN *USDNRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDN.Contract.USDNTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDN *USDNRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDN.Contract.USDNTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_USDN *USDNCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _USDN.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_USDN *USDNTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDN.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_USDN *USDNTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _USDN.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDN *USDNCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDN *USDNSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDN.Contract.Allowance(&_USDN.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_USDN *USDNCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _USDN.Contract.Allowance(&_USDN.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDN *USDNCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDN *USDNSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDN.Contract.BalanceOf(&_USDN.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_USDN *USDNCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _USDN.Contract.BalanceOf(&_USDN.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_USDN *USDNCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_USDN *USDNSession) Decimals() (uint8, error) {
	return _USDN.Contract.Decimals(&_USDN.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() pure returns(uint8)
func (_USDN *USDNCallerSession) Decimals() (uint8, error) {
	return _USDN.Contract.Decimals(&_USDN.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_USDN *USDNCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_USDN *USDNSession) Name() (string, error) {
	return _USDN.Contract.Name(&_USDN.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() pure returns(string)
func (_USDN *USDNCallerSession) Name() (string, error) {
	return _USDN.Contract.Name(&_USDN.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_USDN *USDNCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_USDN *USDNSession) Symbol() (string, error) {
	return _USDN.Contract.Symbol(&_USDN.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() pure returns(string)
func (_USDN *USDNCallerSession) Symbol() (string, error) {
	return _USDN.Contract.Symbol(&_USDN.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDN *USDNCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _USDN.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDN *USDNSession) TotalSupply() (*big.Int, error) {
	return _USDN.Contract.TotalSupply(&_USDN.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_USDN *USDNCallerSession) TotalSupply() (*big.Int, error) {
	return _USDN.Contract.TotalSupply(&_USDN.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDN *USDNTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDN *USDNSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Approve(&_USDN.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_USDN *USDNTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Approve(&_USDN.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDN *USDNTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDN *USDNSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.DecreaseAllowance(&_USDN.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_USDN *USDNTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.DecreaseAllowance(&_USDN.TransactOpts, spender, subtractedValue)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_USDN *USDNTransactor) Deposit(opts *bind.TransactOpts, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "deposit", account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_USDN *USDNSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Deposit(&_USDN.TransactOpts, account, amount)
}

// Deposit is a paid mutator transaction binding the contract method 0x47e7ef24.
//
// Solidity: function deposit(address account, uint256 amount) returns(bool)
func (_USDN *USDNTransactorSession) Deposit(account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Deposit(&_USDN.TransactOpts, account, amount)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_USDN *USDNTransactor) Deprecate(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "deprecate")
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_USDN *USDNSession) Deprecate() (*types.Transaction, error) {
	return _USDN.Contract.Deprecate(&_USDN.TransactOpts)
}

// Deprecate is a paid mutator transaction binding the contract method 0x0fcc0c28.
//
// Solidity: function deprecate() returns()
func (_USDN *USDNTransactorSession) Deprecate() (*types.Transaction, error) {
	return _USDN.Contract.Deprecate(&_USDN.TransactOpts)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDN *USDNTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDN *USDNSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.IncreaseAllowance(&_USDN.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_USDN *USDNTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.IncreaseAllowance(&_USDN.TransactOpts, spender, addedValue)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_USDN *USDNTransactor) Stake(opts *bind.TransactOpts, reward *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "stake", reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_USDN *USDNSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Stake(&_USDN.TransactOpts, reward)
}

// Stake is a paid mutator transaction binding the contract method 0xa694fc3a.
//
// Solidity: function stake(uint256 reward) returns(bool)
func (_USDN *USDNTransactorSession) Stake(reward *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Stake(&_USDN.TransactOpts, reward)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDN *USDNTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDN *USDNSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Transfer(&_USDN.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_USDN *USDNTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.Transfer(&_USDN.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDN *USDNTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDN *USDNSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.TransferFrom(&_USDN.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_USDN *USDNTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _USDN.Contract.TransferFrom(&_USDN.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDN *USDNTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDN *USDNSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDN.Contract.TransferOwnership(&_USDN.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_USDN *USDNTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _USDN.Contract.TransferOwnership(&_USDN.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_USDN *USDNTransactor) Withdraw(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _USDN.contract.Transact(opts, "withdraw", account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_USDN *USDNSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _USDN.Contract.Withdraw(&_USDN.TransactOpts, account)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address account) returns(bool)
func (_USDN *USDNTransactorSession) Withdraw(account common.Address) (*types.Transaction, error) {
	return _USDN.Contract.Withdraw(&_USDN.TransactOpts, account)
}

// USDNApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the USDN contract.
type USDNApprovalIterator struct {
	Event *USDNApproval // Event containing the contract specifics and raw log

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
func (it *USDNApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDNApproval)
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
		it.Event = new(USDNApproval)
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
func (it *USDNApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDNApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDNApproval represents a Approval event raised by the USDN contract.
type USDNApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDN *USDNFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*USDNApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDN.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &USDNApprovalIterator{contract: _USDN.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_USDN *USDNFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *USDNApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _USDN.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDNApproval)
				if err := _USDN.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_USDN *USDNFilterer) ParseApproval(log types.Log) (*USDNApproval, error) {
	event := new(USDNApproval)
	if err := _USDN.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDNDeprecateIterator is returned from FilterDeprecate and is used to iterate over the raw logs and unpacked data for Deprecate events raised by the USDN contract.
type USDNDeprecateIterator struct {
	Event *USDNDeprecate // Event containing the contract specifics and raw log

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
func (it *USDNDeprecateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDNDeprecate)
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
		it.Event = new(USDNDeprecate)
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
func (it *USDNDeprecateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDNDeprecateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDNDeprecate represents a Deprecate event raised by the USDN contract.
type USDNDeprecate struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDeprecate is a free log retrieval operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_USDN *USDNFilterer) FilterDeprecate(opts *bind.FilterOpts, account []common.Address) (*USDNDeprecateIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDN.contract.FilterLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return &USDNDeprecateIterator{contract: _USDN.contract, event: "Deprecate", logs: logs, sub: sub}, nil
}

// WatchDeprecate is a free log subscription operation binding the contract event 0xcc358699805e9a8b7f77b522628c7cb9abd07d9efb86b6fb616af1609036a99e.
//
// Solidity: event Deprecate(address indexed account)
func (_USDN *USDNFilterer) WatchDeprecate(opts *bind.WatchOpts, sink chan<- *USDNDeprecate, account []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}

	logs, sub, err := _USDN.contract.WatchLogs(opts, "Deprecate", accountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDNDeprecate)
				if err := _USDN.contract.UnpackLog(event, "Deprecate", log); err != nil {
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
func (_USDN *USDNFilterer) ParseDeprecate(log types.Log) (*USDNDeprecate, error) {
	event := new(USDNDeprecate)
	if err := _USDN.contract.UnpackLog(event, "Deprecate", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDNRewardIterator is returned from FilterReward and is used to iterate over the raw logs and unpacked data for Reward events raised by the USDN contract.
type USDNRewardIterator struct {
	Event *USDNReward // Event containing the contract specifics and raw log

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
func (it *USDNRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDNReward)
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
		it.Event = new(USDNReward)
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
func (it *USDNRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDNRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDNReward represents a Reward event raised by the USDN contract.
type USDNReward struct {
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterReward is a free log retrieval operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_USDN *USDNFilterer) FilterReward(opts *bind.FilterOpts) (*USDNRewardIterator, error) {

	logs, sub, err := _USDN.contract.FilterLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return &USDNRewardIterator{contract: _USDN.contract, event: "Reward", logs: logs, sub: sub}, nil
}

// WatchReward is a free log subscription operation binding the contract event 0x45cad8c10023de80f4c0672ff6c283b671e11aa93c92b9380cdf060d2790da52.
//
// Solidity: event Reward(uint256 id, uint256 amount)
func (_USDN *USDNFilterer) WatchReward(opts *bind.WatchOpts, sink chan<- *USDNReward) (event.Subscription, error) {

	logs, sub, err := _USDN.contract.WatchLogs(opts, "Reward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDNReward)
				if err := _USDN.contract.UnpackLog(event, "Reward", log); err != nil {
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
func (_USDN *USDNFilterer) ParseReward(log types.Log) (*USDNReward, error) {
	event := new(USDNReward)
	if err := _USDN.contract.UnpackLog(event, "Reward", log); err != nil {
		return nil, err
	}
	return event, nil
}

// USDNTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the USDN contract.
type USDNTransferIterator struct {
	Event *USDNTransfer // Event containing the contract specifics and raw log

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
func (it *USDNTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(USDNTransfer)
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
		it.Event = new(USDNTransfer)
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
func (it *USDNTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *USDNTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// USDNTransfer represents a Transfer event raised by the USDN contract.
type USDNTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDN *USDNFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*USDNTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDN.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &USDNTransferIterator{contract: _USDN.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_USDN *USDNFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *USDNTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _USDN.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(USDNTransfer)
				if err := _USDN.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_USDN *USDNFilterer) ParseTransfer(log types.Log) (*USDNTransfer, error) {
	event := new(USDNTransfer)
	if err := _USDN.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
