// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package buildtoken

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

// BuildTokenControlledTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type BuildTokenControlledTokenConfig struct {
	Name       string
	Symbol     string
	Decimals   uint8
	Controller common.Address
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractControlledTokenProxyFactory\",\"name\":\"_controlledTokenProxyFactory\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"CreatedControlledToken\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"controlledTokenProxyFactory\",\"outputs\":[{\"internalType\":\"contractControlledTokenProxyFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controller\",\"outputs\":[{\"internalType\":\"contractTokenControllerInterface\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"contractTokenControllerInterface\",\"name\":\"controller\",\"type\":\"address\"}],\"internalType\":\"structBuildToken.ControlledTokenConfig\",\"name\":\"config\",\"type\":\"tuple\"}],\"name\":\"createControlledToken\",\"outputs\":[{\"internalType\":\"contractControlledToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040523480156100115760006000fd5b5060405161092c38038061092c83398181016040528101906100339190610108565b5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156100a6576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161009d9061019a565b60405180910390fd5b80600060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b506102305661022f565b60008151905061010181610214565b5b92915050565b60006020828403121561011b5760006000fd5b6000610129848285016100f2565b9150505b92915050565b60006101406025836101bb565b91507f426c696e64426f78204572723a436f6e74726f6c6c6564546f6b656e204e6f7460008301527f205a65726f00000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600060208201905081810360008301526101b381610133565b90505b919050565b60008282526020820190505b92915050565b60006101d8826101f3565b90505b919050565b60006101eb826101cd565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b61021d816101e0565b8114151561022b5760006000fd5b5b50565b5b6106ed8061023f6000396000f3fe60806040523480156100115760006000fd5b50600436106100465760003560e01c80636a81d8bd1461004c5780638c0cd38d1461006a578063f77c47911461009a57610046565b60006000fd5b6100546100b8565b604051610061919061044d565b60405180910390f35b610084600480360381019061007f9190610370565b6100de565b6040516100919190610469565b60405180910390f35b6100a2610291565b6040516100af9190610485565b60405180910390f35b600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006000600060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663efc81a8c6040518163ffffffff1660e01b8152600401602060405180830381600087803b15801561014d5760006000fd5b505af1158015610162573d600060003e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610186919061031a565b90508073ffffffffffffffffffffffffffffffffffffffff1663de7ea79d8480600001906101b491906104f9565b8680602001906101c491906104f9565b8860400160208101906101d791906103b4565b8960600160208101906101ea9190610345565b6040518763ffffffff1660e01b815260040161020b969594939291906104a1565b600060405180830381600087803b1580156102265760006000fd5b505af115801561023b573d600060003e3d6000fd5b505050508073ffffffffffffffffffffffffffffffffffffffff167fe3d5734f17a493c850907f8a8366a543676afd8eeb9b7cd16e22c998297d8ebd60405160405180910390a28091505061028c56505b919050565b600160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681566106b6565b6000815190506102ca81610665565b5b92915050565b6000813590506102e081610680565b5b92915050565b6000608082840312156102fa5760006000fd5b8190505b92915050565b6000813590506103138161069b565b5b92915050565b60006020828403121561032d5760006000fd5b600061033b848285016102bb565b9150505b92915050565b6000602082840312156103585760006000fd5b6000610366848285016102d1565b9150505b92915050565b6000602082840312156103835760006000fd5b600082013567ffffffffffffffff81111561039e5760006000fd5b6103aa848285016102e7565b9150505b92915050565b6000602082840312156103c75760006000fd5b60006103d584828501610304565b9150505b92915050565b6103e8816105d1565b82525b5050565b6103f8816105f7565b82525b5050565b6104088161061d565b82525b5050565b600061041b8385610557565b9350610428838584610643565b61043183610653565b840190505b9392505050565b610446816105c3565b82525b5050565b600060208201905061046260008301846103df565b5b92915050565b600060208201905061047e60008301846103ef565b5b92915050565b600060208201905061049a60008301846103ff565b5b92915050565b600060808201905081810360008301526104bc81888a61040f565b905081810360208301526104d181868861040f565b90506104e0604083018561043d565b6104ed60608301846103ff565b5b979650505050505050565b600060008335600160200384360303811215156105165760006000fd5b80840192508235915067ffffffffffffffff8211156105355760006000fd5b60208301925060018202360383131561054e5760006000fd5b505b9250929050565b60008282526020820190505b92915050565b6000610574826105a2565b90505b919050565b600061058782610569565b90505b919050565b600061059a82610569565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b600060ff821690505b919050565b60006105dc826105e4565b90505b919050565b60006105ef826105a2565b90505b919050565b60006106028261060a565b90505b919050565b6000610615826105a2565b90505b919050565b600061062882610630565b90505b919050565b600061063b826105a2565b90505b919050565b828183376000838301525b505050565b6000601f19601f83011690505b919050565b61066e8161057c565b8114151561067c5760006000fd5b5b50565b6106898161058f565b811415156106975760006000fd5b5b50565b6106a4816105c3565b811415156106b25760006000fd5b5b50565bfea2646970667358221220bbef4c63e13eda7c14b0c4457919f6b3b3ae0081465cff5d1dd7709eef3057f464736f6c63430006050033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _controlledTokenProxyFactory common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _controlledTokenProxyFactory)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// Store is an auto generated Go binding around an Ethereum contract.
type Store struct {
	StoreCaller     // Read-only binding to the contract
	StoreTransactor // Write-only binding to the contract
	StoreFilterer   // Log filterer for contract events
}

// StoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type StoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StoreSession struct {
	Contract     *Store            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StoreCallerSession struct {
	Contract *StoreCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StoreTransactorSession struct {
	Contract     *StoreTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type StoreRaw struct {
	Contract *Store // Generic contract binding to access the raw methods on
}

// StoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StoreCallerRaw struct {
	Contract *StoreCaller // Generic read-only contract binding to access the raw methods on
}

// StoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StoreTransactorRaw struct {
	Contract *StoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStore creates a new instance of Store, bound to a specific deployed contract.
func NewStore(address common.Address, backend bind.ContractBackend) (*Store, error) {
	contract, err := bindStore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Store{StoreCaller: StoreCaller{contract: contract}, StoreTransactor: StoreTransactor{contract: contract}, StoreFilterer: StoreFilterer{contract: contract}}, nil
}

// NewStoreCaller creates a new read-only instance of Store, bound to a specific deployed contract.
func NewStoreCaller(address common.Address, caller bind.ContractCaller) (*StoreCaller, error) {
	contract, err := bindStore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StoreCaller{contract: contract}, nil
}

// NewStoreTransactor creates a new write-only instance of Store, bound to a specific deployed contract.
func NewStoreTransactor(address common.Address, transactor bind.ContractTransactor) (*StoreTransactor, error) {
	contract, err := bindStore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StoreTransactor{contract: contract}, nil
}

// NewStoreFilterer creates a new log filterer instance of Store, bound to a specific deployed contract.
func NewStoreFilterer(address common.Address, filterer bind.ContractFilterer) (*StoreFilterer, error) {
	contract, err := bindStore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StoreFilterer{contract: contract}, nil
}

// bindStore binds a generic wrapper to an already deployed contract.
func bindStore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.StoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.StoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Store *StoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Store.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Store *StoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Store *StoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Store.Contract.contract.Transact(opts, method, params...)
}

// ControlledTokenProxyFactory is a free data retrieval call binding the contract method 0x6a81d8bd.
//
// Solidity: function controlledTokenProxyFactory() view returns(address)
func (_Store *StoreCaller) ControlledTokenProxyFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "controlledTokenProxyFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ControlledTokenProxyFactory is a free data retrieval call binding the contract method 0x6a81d8bd.
//
// Solidity: function controlledTokenProxyFactory() view returns(address)
func (_Store *StoreSession) ControlledTokenProxyFactory() (common.Address, error) {
	return _Store.Contract.ControlledTokenProxyFactory(&_Store.CallOpts)
}

// ControlledTokenProxyFactory is a free data retrieval call binding the contract method 0x6a81d8bd.
//
// Solidity: function controlledTokenProxyFactory() view returns(address)
func (_Store *StoreCallerSession) ControlledTokenProxyFactory() (common.Address, error) {
	return _Store.Contract.ControlledTokenProxyFactory(&_Store.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Store *StoreCaller) Controller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "controller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Store *StoreSession) Controller() (common.Address, error) {
	return _Store.Contract.Controller(&_Store.CallOpts)
}

// Controller is a free data retrieval call binding the contract method 0xf77c4791.
//
// Solidity: function controller() view returns(address)
func (_Store *StoreCallerSession) Controller() (common.Address, error) {
	return _Store.Contract.Controller(&_Store.CallOpts)
}

// CreateControlledToken is a paid mutator transaction binding the contract method 0x8c0cd38d.
//
// Solidity: function createControlledToken((string,string,uint8,address) config) returns(address)
func (_Store *StoreTransactor) CreateControlledToken(opts *bind.TransactOpts, config BuildTokenControlledTokenConfig) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "createControlledToken", config)
}

// CreateControlledToken is a paid mutator transaction binding the contract method 0x8c0cd38d.
//
// Solidity: function createControlledToken((string,string,uint8,address) config) returns(address)
func (_Store *StoreSession) CreateControlledToken(config BuildTokenControlledTokenConfig) (*types.Transaction, error) {
	return _Store.Contract.CreateControlledToken(&_Store.TransactOpts, config)
}

// CreateControlledToken is a paid mutator transaction binding the contract method 0x8c0cd38d.
//
// Solidity: function createControlledToken((string,string,uint8,address) config) returns(address)
func (_Store *StoreTransactorSession) CreateControlledToken(config BuildTokenControlledTokenConfig) (*types.Transaction, error) {
	return _Store.Contract.CreateControlledToken(&_Store.TransactOpts, config)
}

// StoreCreatedControlledTokenIterator is returned from FilterCreatedControlledToken and is used to iterate over the raw logs and unpacked data for CreatedControlledToken events raised by the Store contract.
type StoreCreatedControlledTokenIterator struct {
	Event *StoreCreatedControlledToken // Event containing the contract specifics and raw log

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
func (it *StoreCreatedControlledTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCreatedControlledToken)
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
		it.Event = new(StoreCreatedControlledToken)
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
func (it *StoreCreatedControlledTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCreatedControlledTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCreatedControlledToken represents a CreatedControlledToken event raised by the Store contract.
type StoreCreatedControlledToken struct {
	Token common.Address
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterCreatedControlledToken is a free log retrieval operation binding the contract event 0xe3d5734f17a493c850907f8a8366a543676afd8eeb9b7cd16e22c998297d8ebd.
//
// Solidity: event CreatedControlledToken(address indexed token)
func (_Store *StoreFilterer) FilterCreatedControlledToken(opts *bind.FilterOpts, token []common.Address) (*StoreCreatedControlledTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "CreatedControlledToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &StoreCreatedControlledTokenIterator{contract: _Store.contract, event: "CreatedControlledToken", logs: logs, sub: sub}, nil
}

// WatchCreatedControlledToken is a free log subscription operation binding the contract event 0xe3d5734f17a493c850907f8a8366a543676afd8eeb9b7cd16e22c998297d8ebd.
//
// Solidity: event CreatedControlledToken(address indexed token)
func (_Store *StoreFilterer) WatchCreatedControlledToken(opts *bind.WatchOpts, sink chan<- *StoreCreatedControlledToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "CreatedControlledToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCreatedControlledToken)
				if err := _Store.contract.UnpackLog(event, "CreatedControlledToken", log); err != nil {
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

// ParseCreatedControlledToken is a log parse operation binding the contract event 0xe3d5734f17a493c850907f8a8366a543676afd8eeb9b7cd16e22c998297d8ebd.
//
// Solidity: event CreatedControlledToken(address indexed token)
func (_Store *StoreFilterer) ParseCreatedControlledToken(log types.Log) (*StoreCreatedControlledToken, error) {
	event := new(StoreCreatedControlledToken)
	if err := _Store.contract.UnpackLog(event, "CreatedControlledToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
