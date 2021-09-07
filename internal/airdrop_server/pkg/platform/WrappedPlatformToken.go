// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package platform

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mint_address\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_total\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"release_log\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"duration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"beneficiary\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"total\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mint_address\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"release\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"released\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"seeRelease\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x6080604052348015620000125760006000fd5b5060405162002a7238038062002a72833981810160405260a0811015620000395760006000fd5b8101908080519060200190929190805190602001909291908051906020019092919080519060200190929190805190602001909291905050505b846040518060400160405280600481526020017f527a7279000000000000000000000000000000000000000000000000000000008152602001506040518060400160405280600481526020017f527a7279000000000000000000000000000000000000000000000000000000008152602001508385305b5b84845b5b5b81600360005090805190602001906200010b92919062000851565b5080600460005090805190602001906200012792919062000851565b506012600560006101000a81548160ff021916908360ff1602179055505b50506000620001596200034f60201b60201c565b905080600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508073ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3505b62000224866200021884866200035d60201b62001d831790919060201c565b620003b860201b60201c565b620002368183620003b860201b60201c565b5b505050505050620002578483620005b460201b62001e6b1790919060201c565b600d6000508190909055506040518060c00160405280428152602001600081526020018581526020018473ffffffffffffffffffffffffffffffffffffffff16815260200183815260200182815260200150600760005060008201518160000160005090905560208201518160010160005090905560408201518160020160005090905560608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160005090905560a0820151816005016000509090559050505b50505050506200090c565b60003390506200035a565b90565b6000620003aa83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152602001506200060f60201b60201c565b9050620003b2565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151562000461576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f45524332303a206d696e7420746f20746865207a65726f20616464726573730081526020015060200191505060405180910390fd5b6200047560008383620006dd60201b60201c565b6200049481600260005054620006e360201b620019011790919060201c565b6002600050819090905550620004fe81600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054620006e360201b620019011790919060201c565b600060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b5050565b60006200060183836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f0000000000008152602001506200077960201b60201c565b905062000609565b92915050565b60008383111582901515620006c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620006875780820151818401525b60208101905062000669565b50505050905090810190601f168015620006b55780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b506000838503905080915050620006d656505b9392505050565b5b505050565b60006000828401905083811015151562000768576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081526020015060200191505060405180910390fd5b809150506200077356505b92915050565b600060008311829015156200082d576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015620007f15780820151818401525b602081019050620007d3565b50505050905090810190601f1680156200081f5780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083858115156200083c57fe5b049050809150506200084a56505b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200089457805160ff1916838001178555620008ca565b82800160010185558215620008ca579182015b82811115620008c95782518260005090905591602001919060010190620008a7565b5b509050620008d99190620008dd565b5090565b620009099190620008e9565b80821115620009055760008181506000905550600101620008e9565b5090565b90565b612156806200091c6000396000f3fe60806040523480156100115760006000fd5b50600436106101315760003560e01c806386d1a69f116100ae578063a9059cbb11610072578063a9059cbb146105a4578063dd62ed3e1461060b578063e47c81dd14610684578063f2fde38b146106ce578063fcd3533c1461071357610131565b806386d1a69f146104475780638da5cb5b1461045157806395d89b411461049b578063961325211461051f578063a457c2d71461053d57610131565b8063313ce567116100f5578063313ce567146102ec578063395093511461031057806370a0823114610377578063715018a6146103d057806379502c55146103da57610131565b806306fdde0314610137578063095ea7b3146101bb57806318160ddd1461022257806323b872dd14610240578063258674d6146102c757610131565b60006000fd5b61013f610762565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156101805780820151818401525b602081019050610164565b50505050905090810190601f1680156101ad5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610208600480360360408110156101d25760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061080c565b604051808215151515815260200191505060405180910390f35b61022a61083b565b6040518082815260200191505060405180910390f35b6102ad600480360360608110156102575760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff1690602001909291908035906020019092919050505061084d565b604051808215151515815260200191505060405180910390f35b6102cf61094c565b604051808381526020018281526020019250505060405180910390f35b6102f461096b565b604051808260ff1660ff16815260200191505060405180910390f35b61035d600480360360408110156103275760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610987565b604051808215151515815260200191505060405180910390f35b6103ba6004803603602081101561038e5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610a5a565b6040518082815260200191505060405180910390f35b6103d8610aae565b005b6103e2610c46565b604051808781526020018681526020018581526020018473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001838152602001828152602001965050505050505060405180910390f35b61044f610ca2565b005b610459610d4b565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b6104a3610d7a565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156104e45780820151818401525b6020810190506104c8565b50505050905090810190601f1680156105115780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610527610e24565b6040518082815260200191505060405180910390f35b61058a600480360360408110156105545760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610e2d565b604051808215151515815260200191505060405180910390f35b6105f1600480360360408110156105bb5760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff16906020019092919080359060200190929190505050610f1a565b604051808215151515815260200191505060405180910390f35b61066e600480360360408110156106225760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050610f49565b6040518082815260200191505060405180910390f35b61068c610fde565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b610711600480360360208110156106e55760006000fd5b81019080803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611004565b005b6107606004803603604081101561072a5760006000fd5b8101908080359060200190929190803573ffffffffffffffffffffffffffffffffffffffff169060200190929190505050611223565b005b606060036000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156107fd5780601f106107d2576101008083540402835291602001916107fd565b820191906000526020600020905b8154815290600101906020018083116107e057829003601f168201915b50505050509050610809565b90565b600061082c61081f6112b763ffffffff16565b84846112c463ffffffff16565b60019050610835565b92915050565b6000600260005054905061084a565b90565b60006108608484846114cb63ffffffff16565b61093c846108726112b763ffffffff16565b6109318560405180606001604052806028815260200161206a60289139600160005060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060006108e46112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546117b79092919063ffffffff16565b6112c463ffffffff16565b60019050610945565b9392505050565b6000600061095e61188063ffffffff16565b91509150610967565b9091565b6000600560009054906101000a900460ff169050610984565b90565b6000610a4b61099a6112b763ffffffff16565b84610a4085600160005060006109b46112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461190190919063ffffffff16565b6112c463ffffffff16565b60019050610a54565b92915050565b6000600060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050610aa9565b919050565b610abc6112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610b83576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a36000600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b565b60076000508060000160005054908060010160005054908060020160005054908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160005054908060050160005054905086565b60006000610cb461188063ffffffff16565b8092508193505050610cf530600760005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff168461199563ffffffff16565b8060076000506001016000508190909055507f017dd8653616ddcb225b0ab4964b9a9dfac2138c43f371b5f24a96215979e9f88282604051808381526020018281526020019250505060405180910390a150505b565b6000600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610d77565b90565b606060046000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e155780601f10610dea57610100808354040283529160200191610e15565b820191906000526020600020905b815481529060010190602001808311610df857829003601f168201915b50505050509050610e21565b90565b600d6000505481565b6000610f0b610e406112b763ffffffff16565b84610f00856040518060600160405280602581526020016120fc6025913960016000506000610e736112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546117b79092919063ffffffff16565b6112c463ffffffff16565b60019050610f14565b92915050565b6000610f3a610f2d6112b763ffffffff16565b84846114cb63ffffffff16565b60019050610f43565b92915050565b6000600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050610fd8565b92915050565b600660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6110126112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff161415156110d9576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260208152602001807f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657281526020015060200191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515611161576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526026815260200180611fdb6026913960400191505060405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16600560019054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a380600560016101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b5b50565b6112406112346112b763ffffffff16565b83611b9863ffffffff16565b8073ffffffffffffffffffffffffffffffffffffffff166112656112b763ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff167fbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453846040518082815260200191505060405180910390a35b5050565b60003390506112c1565b90565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff161415151561134c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260248152602001806120d86024913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156113d4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260228152602001806120016022913960400191505060405180910390fd5b80600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040518082815260200191505060405180910390a35b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515611553576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260258152602001806120b36025913960400191505060405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156115db576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401808060200182810382526023815260200180611f966023913960400191505060405180910390fd5b6115ec838383611d7d63ffffffff16565b61165e8160405180606001604052806026815260200161202360269139600060005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546117b79092919063ffffffff16565b600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555061170181600060005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461190190919063ffffffff16565b600060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b505050565b60008383111582901515611867576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b8381101561182c5780820151818401525b602081019050611810565b50505050905090810190601f1680156118595780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b50600083850390508091505061187956505b9392505050565b6000600060006007600050600101600050549050600060076000506001016000505414156118b957600760005060000160005054905080505b600042905060006118d38383611d8390919063ffffffff16565b90506118ed600d6000505482611ddb90919063ffffffff16565b82945094505050506118fd565050505b9091565b600060008284019050838110151515611985576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601b8152602001807f536166654d6174683a206164646974696f6e206f766572666c6f77000000000081526020015060200191505060405180910390fd5b8091505061198f56505b92915050565b600060608473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001828152602001925050506040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040518082805190602001908083835b602083101515611a7157805182525b602082019150602081019050602083039250611a4b565b6001836020036101000a0380198251168184511680821785525050505050509050019150506000604051808303816000865af19150503d8060008114611ad3576040519150601f19603f3d011682016040523d82523d6000602084013e611ad8565b606091505b5091509150818015611b195750600081511480611b185750808060200190516020811015611b065760006000fd5b81019080805190602001909291905050505b5b1515611b90576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040180806020018281038252601f8152602001807f5472616e7366657248656c7065723a205452414e534645525f4641494c45440081526020015060200191505060405180910390fd5b50505b505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515611c20576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806120926021913960400191505060405180910390fd5b611c3282600083611d7d63ffffffff16565b611ca481604051806060016040528060228152602001611fb960229139600060005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546117b79092919063ffffffff16565b600060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550611d0781600260005054611d8390919063ffffffff16565b6002600050819090905550600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040518082815260200191505060405180910390a35b5050565b5b505050565b6000611dce83836040518060400160405280601e81526020017f536166654d6174683a207375627472616374696f6e206f766572666c6f7700008152602001506117b763ffffffff16565b9050611dd5565b92915050565b60006000831415611def5760009050611e65565b60008284029050828482811515611e0257fe5b04141515611e5b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825260218152602001806120496021913960400191505060405180910390fd5b80915050611e6556505b92915050565b6000611eb683836040518060400160405280601a81526020017f536166654d6174683a206469766973696f6e206279207a65726f000000000000815260200150611ec363ffffffff16565b9050611ebd565b92915050565b60006000831182901515611f73576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004018080602001828103825283818151815260200191508051906020019080838360005b83811015611f385780820151818401525b602081019050611f1c565b50505050905090810190601f168015611f655780820380516001836020036101000a031916815260200191505b509250505060405180910390fd5b5060008385811515611f8157fe5b04905080915050611f8e56505b939250505056fe45524332303a207472616e7366657220746f20746865207a65726f206164647265737345524332303a206275726e20616d6f756e7420657863656564732062616c616e63654f776e61626c653a206e6577206f776e657220697320746865207a65726f206164647265737345524332303a20617070726f766520746f20746865207a65726f206164647265737345524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e6365536166654d6174683a206d756c7469706c69636174696f6e206f766572666c6f7745524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a206275726e2066726f6d20746865207a65726f206164647265737345524332303a207472616e736665722066726f6d20746865207a65726f206164647265737345524332303a20617070726f76652066726f6d20746865207a65726f206164647265737345524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa26469706673582212208e23c5bfc8df6288f12f09e57870511515ce38a313a909cee218cb0a6810122764736f6c63430006000033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _mint_address common.Address, _duration *big.Int, _beneficiary common.Address, _quantity *big.Int, _total *big.Int) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _mint_address, _duration, _beneficiary, _quantity, _total)
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

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Store *StoreCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Store.Contract.Allowance(&_Store.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, account)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 start, uint256 end, uint256 duration, address beneficiary, uint256 quantity, uint256 total)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Start       *big.Int
	End         *big.Int
	Duration    *big.Int
	Beneficiary common.Address
	Quantity    *big.Int
	Total       *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Start       *big.Int
		End         *big.Int
		Duration    *big.Int
		Beneficiary common.Address
		Quantity    *big.Int
		Total       *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Start = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.End = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Duration = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Beneficiary = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Quantity = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Total = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 start, uint256 end, uint256 duration, address beneficiary, uint256 quantity, uint256 total)
func (_Store *StoreSession) Config() (struct {
	Start       *big.Int
	End         *big.Int
	Duration    *big.Int
	Beneficiary common.Address
	Quantity    *big.Int
	Total       *big.Int
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(uint256 start, uint256 end, uint256 duration, address beneficiary, uint256 quantity, uint256 total)
func (_Store *StoreCallerSession) Config() (struct {
	Start       *big.Int
	End         *big.Int
	Duration    *big.Int
	Beneficiary common.Address
	Quantity    *big.Int
	Total       *big.Int
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Store *StoreCallerSession) Decimals() (uint8, error) {
	return _Store.Contract.Decimals(&_Store.CallOpts)
}

// MintAddress is a free data retrieval call binding the contract method 0xe47c81dd.
//
// Solidity: function mint_address() view returns(address)
func (_Store *StoreCaller) MintAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "mint_address")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// MintAddress is a free data retrieval call binding the contract method 0xe47c81dd.
//
// Solidity: function mint_address() view returns(address)
func (_Store *StoreSession) MintAddress() (common.Address, error) {
	return _Store.Contract.MintAddress(&_Store.CallOpts)
}

// MintAddress is a free data retrieval call binding the contract method 0xe47c81dd.
//
// Solidity: function mint_address() view returns(address)
func (_Store *StoreCallerSession) MintAddress() (common.Address, error) {
	return _Store.Contract.MintAddress(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Store *StoreCallerSession) Name() (string, error) {
	return _Store.Contract.Name(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Store *StoreCallerSession) Owner() (common.Address, error) {
	return _Store.Contract.Owner(&_Store.CallOpts)
}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Store *StoreCaller) Released(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "released")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Store *StoreSession) Released() (*big.Int, error) {
	return _Store.Contract.Released(&_Store.CallOpts)
}

// Released is a free data retrieval call binding the contract method 0x96132521.
//
// Solidity: function released() view returns(uint256)
func (_Store *StoreCallerSession) Released() (*big.Int, error) {
	return _Store.Contract.Released(&_Store.CallOpts)
}

// SeeRelease is a free data retrieval call binding the contract method 0x258674d6.
//
// Solidity: function seeRelease() view returns(uint256, uint256)
func (_Store *StoreCaller) SeeRelease(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "seeRelease")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// SeeRelease is a free data retrieval call binding the contract method 0x258674d6.
//
// Solidity: function seeRelease() view returns(uint256, uint256)
func (_Store *StoreSession) SeeRelease() (*big.Int, *big.Int, error) {
	return _Store.Contract.SeeRelease(&_Store.CallOpts)
}

// SeeRelease is a free data retrieval call binding the contract method 0x258674d6.
//
// Solidity: function seeRelease() view returns(uint256, uint256)
func (_Store *StoreCallerSession) SeeRelease() (*big.Int, *big.Int, error) {
	return _Store.Contract.SeeRelease(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Store *StoreCallerSession) Symbol() (string, error) {
	return _Store.Contract.Symbol(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Store *StoreCallerSession) TotalSupply() (*big.Int, error) {
	return _Store.Contract.TotalSupply(&_Store.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, spender, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 amount, address to) returns()
func (_Store *StoreTransactor) Burn(opts *bind.TransactOpts, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burn", amount, to)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 amount, address to) returns()
func (_Store *StoreSession) Burn(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, amount, to)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 amount, address to) returns()
func (_Store *StoreTransactorSession) Burn(amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, amount, to)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Store *StoreTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DecreaseAllowance(&_Store.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Store *StoreTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Store.Contract.IncreaseAllowance(&_Store.TransactOpts, spender, addedValue)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Store *StoreTransactor) Release(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "release")
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Store *StoreSession) Release() (*types.Transaction, error) {
	return _Store.Contract.Release(&_Store.TransactOpts)
}

// Release is a paid mutator transaction binding the contract method 0x86d1a69f.
//
// Solidity: function release() returns()
func (_Store *StoreTransactorSession) Release() (*types.Transaction, error) {
	return _Store.Contract.Release(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Store *StoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Store.Contract.RenounceOwnership(&_Store.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Transfer(&_Store.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_Store *StoreTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, sender, recipient, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Store *StoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Store.Contract.TransferOwnership(&_Store.TransactOpts, newOwner)
}

// StoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Store contract.
type StoreApprovalIterator struct {
	Event *StoreApproval // Event containing the contract specifics and raw log

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
func (it *StoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApproval)
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
		it.Event = new(StoreApproval)
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
func (it *StoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApproval represents a Approval event raised by the Store contract.
type StoreApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApproval)
				if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the Store contract.
type StoreBurnIterator struct {
	Event *StoreBurn // Event containing the contract specifics and raw log

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
func (it *StoreBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBurn)
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
		it.Event = new(StoreBurn)
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
func (it *StoreBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBurn represents a Burn event raised by the Store contract.
type StoreBurn struct {
	Sender common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed _sender, address indexed _to, uint256 amount)
func (_Store *StoreFilterer) FilterBurn(opts *bind.FilterOpts, _sender []common.Address, _to []common.Address) (*StoreBurnIterator, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Burn", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return &StoreBurnIterator{contract: _Store.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed _sender, address indexed _to, uint256 amount)
func (_Store *StoreFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *StoreBurn, _sender []common.Address, _to []common.Address) (event.Subscription, error) {

	var _senderRule []interface{}
	for _, _senderItem := range _sender {
		_senderRule = append(_senderRule, _senderItem)
	}
	var _toRule []interface{}
	for _, _toItem := range _to {
		_toRule = append(_toRule, _toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Burn", _senderRule, _toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBurn)
				if err := _Store.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0xbac40739b0d4ca32fa2d82fc91630465ba3eddd1598da6fca393b26fb63b9453.
//
// Solidity: event Burn(address indexed _sender, address indexed _to, uint256 amount)
func (_Store *StoreFilterer) ParseBurn(log types.Log) (*StoreBurn, error) {
	event := new(StoreBurn)
	if err := _Store.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Store contract.
type StoreOwnershipTransferredIterator struct {
	Event *StoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreOwnershipTransferred)
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
		it.Event = new(StoreOwnershipTransferred)
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
func (it *StoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreOwnershipTransferred represents a OwnershipTransferred event raised by the Store contract.
type StoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StoreOwnershipTransferredIterator{contract: _Store.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Store *StoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreOwnershipTransferred)
				if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Store *StoreFilterer) ParseOwnershipTransferred(log types.Log) (*StoreOwnershipTransferred, error) {
	event := new(StoreOwnershipTransferred)
	if err := _Store.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Store contract.
type StoreTransferIterator struct {
	Event *StoreTransfer // Event containing the contract specifics and raw log

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
func (it *StoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreTransfer)
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
		it.Event = new(StoreTransfer)
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
func (it *StoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreTransfer represents a Transfer event raised by the Store contract.
type StoreTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreTransfer)
				if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreReleaseLogIterator is returned from FilterReleaseLog and is used to iterate over the raw logs and unpacked data for ReleaseLog events raised by the Store contract.
type StoreReleaseLogIterator struct {
	Event *StoreReleaseLog // Event containing the contract specifics and raw log

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
func (it *StoreReleaseLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreReleaseLog)
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
		it.Event = new(StoreReleaseLog)
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
func (it *StoreReleaseLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreReleaseLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreReleaseLog represents a ReleaseLog event raised by the Store contract.
type StoreReleaseLog struct {
	Arg0 *big.Int
	Arg1 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterReleaseLog is a free log retrieval operation binding the contract event 0x017dd8653616ddcb225b0ab4964b9a9dfac2138c43f371b5f24a96215979e9f8.
//
// Solidity: event release_log(uint256 arg0, uint256 arg1)
func (_Store *StoreFilterer) FilterReleaseLog(opts *bind.FilterOpts) (*StoreReleaseLogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "release_log")
	if err != nil {
		return nil, err
	}
	return &StoreReleaseLogIterator{contract: _Store.contract, event: "release_log", logs: logs, sub: sub}, nil
}

// WatchReleaseLog is a free log subscription operation binding the contract event 0x017dd8653616ddcb225b0ab4964b9a9dfac2138c43f371b5f24a96215979e9f8.
//
// Solidity: event release_log(uint256 arg0, uint256 arg1)
func (_Store *StoreFilterer) WatchReleaseLog(opts *bind.WatchOpts, sink chan<- *StoreReleaseLog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "release_log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreReleaseLog)
				if err := _Store.contract.UnpackLog(event, "release_log", log); err != nil {
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

// ParseReleaseLog is a log parse operation binding the contract event 0x017dd8653616ddcb225b0ab4964b9a9dfac2138c43f371b5f24a96215979e9f8.
//
// Solidity: event release_log(uint256 arg0, uint256 arg1)
func (_Store *StoreFilterer) ParseReleaseLog(log types.Log) (*StoreReleaseLog, error) {
	event := new(StoreReleaseLog)
	if err := _Store.contract.UnpackLog(event, "release_log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
