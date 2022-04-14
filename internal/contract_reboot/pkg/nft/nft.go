// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package nft

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

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenSerialNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenTypeNumber\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenVersion\",\"type\":\"uint256\"}],\"name\":\"DrawCard\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenSerialNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenTypeNumber\",\"type\":\"string\"}],\"name\":\"DrawCardForD\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"}],\"name\":\"cashCheckByTokenIdLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"res\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenSerialNumber\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"tokenTypeNumber\",\"type\":\"string\"}],\"name\":\"gradeComposelog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addrOne\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addrTwo\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"addrThree\",\"type\":\"address\"}],\"name\":\"initLog\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"mint_pro_log\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"IsDCard\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rand\",\"type\":\"uint256\"}],\"name\":\"Draw\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdArray\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DrawPro\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"DrawPros\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"tokenIdArray\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"grade\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_rand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"MintMixPro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rand\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"MintPro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"blindBoxPromotion\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handing\",\"type\":\"address\"}],\"name\":\"addBoxPro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"baseURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_gradeNumbers\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"cashCheckByTokenID\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"checkMix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_grade\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id5\",\"type\":\"uint256\"}],\"name\":\"checkNftGrade\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id_1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"id_5\",\"type\":\"uint256\"}],\"name\":\"checkNftOwner\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"S_check\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"S_res\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"A_check\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"A_res\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"B_check\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"B_res\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"C_check\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"C_res\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"D_check\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"D_res\",\"type\":\"uint128\"}],\"name\":\"checkRet\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id1\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id2\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id3\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id4\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_id5\",\"type\":\"uint256\"}],\"name\":\"checkSeriesId\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lockContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"blindBox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handing\",\"type\":\"address\"},{\"internalType\":\"contractBlindBox\",\"name\":\"blindBox_ct\",\"type\":\"address\"},{\"internalType\":\"contractBlindBoxPromotions\",\"name\":\"blindBox_pt\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"exists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"}],\"name\":\"getAddrAllTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAddrIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_pageSize\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_page\",\"type\":\"uint256\"}],\"name\":\"getAddrSeriesTokenIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"result\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getAddrTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_seriesId\",\"type\":\"uint256\"}],\"name\":\"getAddrTokenNumberForSeries\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getNftInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tSerialNumber\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"tTypeNumber\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"tGrade\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"tGradeId\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOwnerAddr\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seriesId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokenIds\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"_rand\",\"type\":\"uint256\"}],\"name\":\"gradeCompose\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mixTokenId\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"blindBox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lock_contract\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastblockhashused\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastblocknumberused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastresult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"queryConfig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"blindBox\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lockContract\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"reckon\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"res\",\"type\":\"uint128\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenIds\",\"type\":\"uint256[]\"}],\"name\":\"transferArray\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x60806040526001600b6000509090553480156200001c5760006000fd5b506040516200a50d3803806200a50d83398181016040528101906200004291906200046f565b5b81815b5b5b5b620000616301ffc9a760e01b6200026f60201b60201c565b5b81600660005090805190602001906200007d9291906200034e565b508060076000509080519060200190620000999291906200034e565b50620000b26380ac58cd60e01b6200026f60201b60201c565b620000ca635b5e139f60e01b6200026f60201b60201c565b620000e263780e9d6360e01b6200026f60201b60201c565b5b505033600c60005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600c60005060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600c60005060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600c60005060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506000600c60005060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506032601f600050819090905550600160206000508190909055505b5050620005f7565b63ffffffff60e01b817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff191614151515620002dd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401620002d4906200052a565b60405180910390fd5b600160006000506000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060006101000a81548160ff0219169083151502179055505b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200039157805160ff1916838001178555620003c7565b82800160010185558215620003c7579182015b82811115620003c65782518260005090905591602001919060010190620003a4565b5b509050620003d69190620003da565b5090565b620004069190620003e6565b80821115620004025760008181506000905550600101620003e6565b5090565b9056620005f6565b600082601f8301121515620004235760006000fd5b81516200043a62000434826200057d565b6200054d565b91508082526020830160208301858383011115620004585760006000fd5b62000465838284620005be565b5050505b92915050565b6000600060408385031215620004855760006000fd5b600083015167ffffffffffffffff811115620004a15760006000fd5b620004af858286016200040e565b925050602083015167ffffffffffffffff811115620004ce5760006000fd5b620004dc858286016200040e565b9150505b9250929050565b6000620004f6601c83620005ac565b91507f4552433136353a20696e76616c696420696e746572666163652069640000000060008301526020820190505b919050565b600060208201905081810360008301526200054581620004e7565b90505b919050565b6000604051905081810181811067ffffffffffffffff82111715620005725760006000fd5b80604052505b919050565b600067ffffffffffffffff821115620005965760006000fd5b601f19601f83011690506020810190505b919050565b60008282526020820190505b92915050565b60005b83811015620005df5780820151818401525b602081019050620005c1565b83811115620005ef576000848401525b505b505050565b5b619f0680620006076000396000f3fe60806040523480156100115760006000fd5b506004361061029a5760003560e01c80638b13004a11610168578063d5ec2d83116100cf578063e670472311610088578063e6704723146108c7578063e68f909d146108e5578063e985e9c514610905578063ec0f681114610935578063f3c3fceb14610965578063fe49ba32146109955761029a565b8063d5ec2d83146107de578063dc84e07d14610811578063dcc64c841461082d578063dd26c8181461085d578063e223809a1461088d578063e3223bcb146108a95761029a565b8063a0c55e3111610121578063a0c55e311461071e578063a22cb4651461073a578063afe6ddc014610756578063b88d4fde14610774578063bd3491f214610790578063c87b56dd146107ae5761029a565b80638b13004a146106245780638e8aa3d214610654578063917acc8a1461068457806395d89b41146106b45780639c38c5a2146106d25780639dc29fac146107025761029a565b80634e6d4e411161020c5780635c4d06e3116101c55780635c4d06e3146105365780636352211e146105525780636c0360eb1461058257806370a08231146105a057806379502c55146105d05780637c29c516146105f45761029a565b80634e6d4e411461043c5780634f558e791461046c5780634f6ccce71461049c57806354af0187146104cc57806356aebce7146104e857806358c61c02146105185761029a565b8063184b95591161025e578063184b95591461035857806323b872dd146103745780632f745c5914610390578063382e5acd146103c05780633b3a70cd146103f057806342842e0e146104205761029a565b806301ffc9a7146102a057806306fdde03146102d0578063081812fc146102ee578063095ea7b31461031e57806318160ddd1461033a5761029a565b60006000fd5b6102ba60048036038101906102b5919061844b565b6109b1565b6040516102c79190619664565b60405180910390f35b6102d8610a21565b6040516102e5919061969c565b60405180910390f35b61030860048036038101906103039190618583565b610acb565b604051610315919061940a565b60405180910390f35b6103386004803603810190610333919061804c565b610b60565b005b610342610c9b565b60405161034f9190619acb565b60405180910390f35b610372600480360381019061036d9190617e39565b610cbc565b005b61038e60048036038101906103899190617f3b565b6110fd565b005b6103aa60048036038101906103a5919061804c565b611172565b6040516103b79190619acb565b60405180910390f35b6103da60048036038101906103d59190618583565b6111d8565b6040516103e79190619aaf565b60405180910390f35b61040a6004803603810190610405919061804c565b611239565b6040516104179190619acb565b60405180910390f35b61043a60048036038101906104359190617f3b565b6112aa565b005b6104566004803603810190610451919061817f565b6112d4565b6040516104639190619acb565b60405180910390f35b61048660048036038101906104819190618583565b611783565b6040516104939190619664565b60405180910390f35b6104b660048036038101906104b19190618583565b6117a0565b6040516104c39190619acb565b60405180910390f35b6104e660048036038101906104e19190617eca565b6117cd565b005b61050260048036038101906104fd9190618297565b6118c9565b60405161050f9190619641565b60405180910390f35b610520611b0d565b60405161052d9190619acb565b60405180910390f35b610550600480360381019061054b91906182fe565b611d45565b005b61056c60048036038101906105679190618583565b612169565b604051610579919061940a565b60405180910390f35b61058a6121a8565b604051610597919061969c565b60405180910390f35b6105ba60048036038101906105b59190617dcf565b612252565b6040516105c79190619acb565b60405180910390f35b6105d8612327565b6040516105eb97969594939291906194e3565b60405180910390f35b61060e6004803603810190610609919061804c565b61243a565b60405161061b9190619acb565b60405180910390f35b61063e60048036038101906106399190618583565b61254a565b60405161064b9190619acb565b60405180910390f35b61066e60048036038101906106699190618379565b612572565b60405161067b9190619664565b60405180910390f35b61069e600480360381019061069991906185ae565b6126eb565b6040516106ab9190619664565b60405180910390f35b6106bc612742565b6040516106c9919061969c565b60405180910390f35b6106ec60048036038101906106e791906185ae565b6127ec565b6040516106f99190619664565b60405180910390f35b61071c6004803603810190610717919061804c565b612843565b005b6107386004803603810190610733919061808b565b612d7f565b005b610754600480360381019061074f919061800d565b61363a565b005b61075e6137d6565b60405161076b9190619680565b60405180910390f35b61078e60048036038101906107899190617f8d565b6137df565b005b610798613856565b6040516107a5919061969c565b60405180910390f35b6107c860048036038101906107c39190618583565b6138f7565b6040516107d5919061969c565b60405180910390f35b6107f860048036038101906107f39190618583565b613a5d565b6040516108089493929190619ae7565b60405180910390f35b61082b60048036038101906108269190618217565b613c15565b005b610847600480360381019061084291906184a1565b6141c7565b6040516108549190619664565b60405180910390f35b610877600480360381019061087291906182fe565b6142d6565b6040516108849190619641565b60405180910390f35b6108a760048036038101906108a29190618297565b614803565b005b6108b1614c55565b6040516108be9190619acb565b60405180910390f35b6108cf614c5e565b6040516108dc919061940a565b60405180910390f35b6108ed614c93565b6040516108fc939291906194ab565b60405180910390f35b61091f600480360381019061091a9190617e8b565b614d26565b60405161092c9190619664565b60405180910390f35b61094f600480360381019061094a919061812d565b614dc5565b60405161095c9190619641565b60405180910390f35b61097f600480360381019061097a9190618583565b614f68565b60405161098c9190619641565b60405180910390f35b6109af60048036038101906109aa9190617dfa565b615229565b005b600060006000506000837bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19167bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916815260200190815260200160002060009054906101000a900460ff169050610a1c565b919050565b606060066000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610abc5780601f10610a9157610100808354040283529160200191610abc565b820191906000526020600020905b815481529060010190602001808311610a9f57829003601f168201915b50505050509050610ac8565b90565b6000610adc8261546563ffffffff16565b1515610b1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610b1490619944565b60405180910390fd5b6004600050600083815260200190815260200160002060009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050610b5b565b919050565b6000610b718261216963ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515610be4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bdb90619a0a565b60405180910390fd5b8073ffffffffffffffffffffffffffffffffffffffff16610c0961548a63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff161480610c445750610c4381610c3861548a63ffffffff16565b614d2663ffffffff16565b5b1515610c85576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610c7c906198c0565b60405180910390fd5b610c95838361549763ffffffff16565b505b5050565b6000610cb2600260005061555a909063ffffffff16565b9050610cb9565b90565b600c60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d54576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d4b9061987e565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff16600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16148015610e0c5750600073ffffffffffffffffffffffffffffffffffffffff16600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b8015610e6c5750600073ffffffffffffffffffffffffffffffffffffffff16600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1515610ead576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610ea490619734565b60405180910390fd5b81600c60005060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600c60005060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600c60005060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600c60005060050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550600b60005054601360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550600b60005054601360005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050819090905550600b60005054601360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055507f8ab07c0c846b8760fd84f71bc8e5f506d49de3e596c4561fd19f9a99bc42c00e8383836040516110ef93929190619426565b60405180910390a15b505050565b61111a61110e61548a63ffffffff16565b8261557d63ffffffff16565b151561115b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161115290619a2b565b60405180910390fd5b61116c83838361567b63ffffffff16565b5b505050565b60006111cb82600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005061593a90919063ffffffff16565b90506111d2565b92915050565b60006000905080506000600190506000600090505b8381101561123157806fffffffffffffffffffffffffffffffff16826fffffffffffffffffffffffffffffffff16901b8301925082505b80806001019150506111ed565b50505b919050565b6000601a60005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508281548110151561128d57fe5b906000526020600020900160005b505490506112a4565b92915050565b6112ce838383604051806020016040528060008152602001506137df63ffffffff16565b5b505050565b6000600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061138b5750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806113e95750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806114475750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8061149f5750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156114e0576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114d79061987e565b60405180910390fd5b6060600c60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637790f6ba886040518263ffffffff1660e01b81526004016115439190619acb565b60006040518083038186803b15801561155c5760006000fd5b505afa158015611571573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061159a9190618407565b90506060600c60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f229c71c896040518263ffffffff1660e01b81526004016115ff9190619acb565b60006040518083038186803b1580156116185760006000fd5b505afa15801561162d573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906116569190618407565b90506116ab898989898980806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f82011690508083019250505050505050613c1563ffffffff16565b600060206000505490506116ca8a60206000505461596263ffffffff16565b6020600081815054809291906001019190509090555060606116f6828b8b89888861598a63ffffffff16565b905080507fe0eb607747eb5726df347537f089eca562d68096cd2a46fd1dddb1279975e1d2818360146000506000868152602001908152602001600020600050546015600050600087815260200190815260200160002060005060405161176094939291906196bf565b60405180910390a18194505050505061177856505050505b5b9695505050505050565b60006117948261546563ffffffff16565b905061179b565b919050565b600060006117bb836002600050615fcb90919063ffffffff16565b509050809150506117c856505b919050565b600082829050905060008111151561181a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161181190619a8e565b60405180910390fd5b6000600090505b818110156118c05761184c848483818110151561183a57fe5b9050602002013561178363ffffffff16565b151561188d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611884906199e9565b60405180910390fd5b6118b2868686868581811015156118a057fe5b905060200201356110fd63ffffffff16565b5b8080600101915050611821565b50505b50505050565b606060008590506000601a60005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005080549050905080600185038602101515611966576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161195d90619965565b60405180910390fd5b60008590506000859050601f6000505482111561198857601f60005054915081505b6000826001830302905060008382860310156119a85784905080506119b0565b838201905080505b8367ffffffffffffffff811180156119c85760006000fd5b506040519080825280602002602001820160405280156119f75781602001602082028036833780820191505090505b509650865060006000905060008390505b82811015611aee578b60146000506000601a60005060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005084815481101515611a6a57fe5b906000526020600020900160005b50548152602001908152602001600020600050541415611ace57611aa2888261123963ffffffff16565b8983815181101515611ab057fe5b60200260200101909081815260200150508180600101925050611ae0565b86831015611adf5782806001019350505b5b5b8080600101915050611a08565b5087975050505050505050611b0556505050505050505b949350505050565b6000600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611bc45750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611c225750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611c805750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611cd85750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611d19576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d109061987e565b60405180910390fd5b6020600050549050805060206000818150548092919060010191905090905550809050611d41565b5b90565b600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480611dfa5750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611e585750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611eb65750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80611f0e5750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515611f4f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f469061987e565b60405180910390fd5b6060600c60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637790f6ba866040518263ffffffff1660e01b8152600401611fb29190619acb565b60006040518083038186803b158015611fcb5760006000fd5b505afa158015611fe0573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906120099190618407565b90506060600c60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f229c71c876040518263ffffffff1660e01b815260040161206e9190619acb565b60006040518083038186803b1580156120875760006000fd5b505afa15801561209c573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906120c59190618407565b90506120d7878461596263ffffffff16565b60606120ed84888888878761598a63ffffffff16565b90507fe0eb607747eb5726df347537f089eca562d68096cd2a46fd1dddb1279975e1d2818560146000506000888152602001908152602001600020600050546015600050600089815260200190815260200160002060005060405161215594939291906196bf565b60405180910390a15050505b5b5050505050565b600061219c82604051806060016040528060298152602001619ea860299139600260005061600c9092919063ffffffff16565b90506121a3565b919050565b606060096000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122435780601f1061221857610100808354040283529160200191612243565b820191906000526020600020905b81548152906001019060200180831161222657829003601f168201915b5050505050905061224f565b90565b6000600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156122c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122bd906198e1565b60405180910390fd5b61231b600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050616039909063ffffffff16565b9050612322565b919050565b600c6000508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b60006000601a60005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054905090506000600090506000600090505b82811015612537578460146000506000601a60005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050848154811015156124f857fe5b906000526020600020900160005b505481526020019081526020016000206000505414156125295781806001019250505b5b8080600101915050612496565b5080925050506125445650505b92915050565b6000601b600050600083815260200190815260200160002060005054905061256d565b919050565b60008673ffffffffffffffffffffffffffffffffffffffff1661259a8761216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff161480156125f657508673ffffffffffffffffffffffffffffffffffffffff166125de8661216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b801561263b57508673ffffffffffffffffffffffffffffffffffffffff166126238561216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b801561268057508673ffffffffffffffffffffffffffffffffffffffff166126688461216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b80156126c557508673ffffffffffffffffffffffffffffffffffffffff166126ad8361216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b156126d757600190506126e1566126e0565b600090506126e1565b5b9695505050505050565b600086861480156126fb57508685145b801561270657508684145b801561271157508683145b801561271c57508682145b1561272e576001905061273856612737565b60009050612738565b5b9695505050505050565b606060076000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156127dd5780601f106127b2576101008083540402835291602001916127dd565b820191906000526020600020905b8154815290600101906020018083116127c057829003601f168201915b505050505090506127e9565b90565b600086861480156127fc57508685145b801561280757508684145b801561281257508683145b801561281d57508682145b1561282f576001905061283956612838565b60009050612839565b5b9695505050505050565b600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806128f85750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806129565750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806129b45750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80612a0c5750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515612a4d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612a449061987e565b60405180910390fd5b612a5c8161216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515612acb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612ac290619a4c565b60405180910390fd5b6014600050600082815260200190815260200160002060005060009055601660005060008281526020019081526020016000206000612b0a9190617a67565b6018600050600082815260200190815260200160002060005060009055601560005060008281526020019081526020016000206000612b499190617a67565b601760005060008281526020019081526020016000206000506000905560196000506000828152602001908152602001600020600050600090556000601a60005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054905090506000601b60005060008481526020019081526020016000206000505490506000601a60005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060018403815481101515612c4557fe5b906000526020600020900160005b5054905080601a60005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005083815481101515612caa57fe5b906000526020600020900160005b5081909090555081601b6000506000838152602001908152602001600020600050819090905550601a60005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054801515612d2f57fe5b600190038181906000526020600020900160005b50600090559055601b600050600085815260200190815260200160002060005060009055612d768461605c63ffffffff16565b5050505b5b5050565b600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480612e345750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80612e925750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80612ef05750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80612f485750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515612f89576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612f809061987e565b60405180910390fd5b81819050848460048181101515612f9c57fe5b90506020020135858560038181101515612fb257fe5b90506020020135868660028181101515612fc857fe5b90506020020135878760018181101515612fde57fe5b90506020020135888860008181101515612ff457fe5b9050602002013501010101141515613041576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613038906197b8565b60405180910390fd5b60006000905060006000905060006000905060006000905060006000905060006000600090505b8888905081101561349e5760606130df61309b8b8b85818110151561308957fe5b905060200201356161bd63ffffffff16565b6040518060400160405280601781526020017f20546f6b656e696420646f6573206e6f7420657869737400000000000000000081526020015061631463ffffffff16565b90506131048a8a8481811015156130f257fe5b9050602002013561178363ffffffff16565b81901515613148576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161313f919061969c565b60405180910390fd5b508c601460005060008c8c86818110151561315f57fe5b905060200201358152602001908152602001600020600050541415156131ba576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131b190619a6d565b60405180910390fd5b6001601760005060008c8c8681811015156131d157fe5b9050602002013581526020019081526020016000206000505414156131fa5787925082506132c3565b6002601760005060008c8c86818110151561321157fe5b90506020020135815260200190815260200160002060005054141561323a5786925082506132c2565b6003601760005060008c8c86818110151561325157fe5b90506020020135815260200190815260200160002060005054141561327a5785925082506132c1565b6004601760005060008c8c86818110151561329157fe5b9050602002013581526020019081526020016000206000505414156132ba5784925082506132c0565b83925082505b5b5b5b60006001601860005060008d8d8781811015156132dc57fe5b90506020020135815260200190815260200160002060005054036001901b846fffffffffffffffffffffffffffffffff1616141561348f576001601860005060008c8c86818110151561332b57fe5b90506020020135815260200190815260200160002060005054036001901b836fffffffffffffffffffffffffffffffff1617925082506001601760005060008c8c86818110151561337857fe5b9050602002013581526020019081526020016000206000505414156133a157829750875061346a565b6002601760005060008c8c8681811015156133b857fe5b9050602002013581526020019081526020016000206000505414156133e1578296508650613469565b6003601760005060008c8c8681811015156133f857fe5b905060200201358152602001908152602001600020600050541415613421578295508550613468565b6004601760005060008c8c86818110151561343857fe5b905060200201358152602001908152602001600020600050541415613461578294508450613467565b82935083505b5b5b5b61348e8e8b8b85818110151561347c57fe5b9050602002013561284363ffffffff16565b5b505b8080600101915050613068565b5060608a8a80806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505090506135b161350d8260008151811015156134fa57fe5b60200260200101516111d863ffffffff16565b8861353384600181518110151561352057fe5b60200260200101516111d863ffffffff16565b8961355986600281518110151561354657fe5b60200260200101516111d863ffffffff16565b8a61357f88600381518110151561356c57fe5b60200260200101516111d863ffffffff16565b8b6135a58a600481518110151561359257fe5b60200260200101516111d863ffffffff16565b8c6141c763ffffffff16565b15156135f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135e9906199c8565b60405180910390fd5b7f07d6c7fbbec58f3fafbf4721d64886946128f0518351d7db11dfae01ed9dc48d8c6040516136219190619acb565b60405180910390a1505050505050505b5b505050505050565b61364861548a63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156136b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136af906197fa565b60405180910390fd5b80600560005060006136ce61548a63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060006101000a81548160ff0219169083151502179055508173ffffffffffffffffffffffffffffffffffffffff1661378461548a63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff167f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31836040516137c99190619664565b60405180910390a35b5050565b601e6000505481565b6137fc6137f061548a63ffffffff16565b8361557d63ffffffff16565b151561383d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161383490619a2b565b60405180910390fd5b61384f8484848461648d63ffffffff16565b5b50505050565b601c6000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156138ef5780601f106138c4576101008083540402835291602001916138ef565b820191906000526020600020905b8154815290600101906020018083116138d257829003601f168201915b505050505081565b60606139088261546563ffffffff16565b1515613949576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613940906199a7565b60405180910390fd5b6060600860005060008481526020019081526020016000206000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156139f85780601f106139cd576101008083540402835291602001916139f8565b820191906000526020600020905b8154815290600101906020018083116139db57829003601f168201915b50505050509050600081511415613a285760405180602001604052806000815260200150915050613a5856613a56565b600960005081604051602001613a3f9291906193b8565b604051602081830303815290604052915050613a58565b505b919050565b6000606060606000601460005060008681526020019081526020016000206000505493508350601560005060008681526020019081526020016000206000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015613b305780601f10613b0557610100808354040283529160200191613b30565b820191906000526020600020905b815481529060010190602001808311613b1357829003601f168201915b505050505092508250601660005060008681526020019081526020016000206000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015613be65780601f10613bbb57610100808354040283529160200191613be6565b820191906000526020600020905b815481529060010190602001808311613bc957829003601f168201915b5050505050915081506018600050600086815260200190815260200160002060005054905080505b9193509193565b600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480613cca5750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80613d285750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80613d865750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80613dde5750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515613e1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e169061987e565b60405180910390fd5b60058151141515613e65576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e5c9061989f565b60405180910390fd5b613f648360146000506000846000815181101515613e7f57fe5b602002602001015181526020019081526020016000206000505460146000506000856001815181101515613eaf57fe5b602002602001015181526020019081526020016000206000505460146000506000866002815181101515613edf57fe5b602002602001015181526020019081526020016000206000505460146000506000876003815181101515613f0f57fe5b602002602001015181526020019081526020016000206000505460146000506000886004815181101515613f3f57fe5b60200260200101518152602001908152602001600020600050546127ec63ffffffff16565b1515613fa5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f9c90619755565b60405180910390fd5b6140a48260176000506000846000815181101515613fbf57fe5b602002602001015181526020019081526020016000206000505460176000506000856001815181101515613fef57fe5b60200260200101518152602001908152602001600020600050546017600050600086600281518110151561401f57fe5b60200260200101518152602001908152602001600020600050546017600050600087600381518110151561404f57fe5b60200260200101518152602001908152602001600020600050546017600050600088600481518110151561407f57fe5b60200260200101518152602001908152602001600020600050546126eb63ffffffff16565b15156140e5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140dc9061981b565b60405180910390fd5b61413a81601360005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546164f863ffffffff16565b151561417b576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141729061985d565b60405180910390fd5b6000600090505b81518110156141be576141b085838381518110151561419d57fe5b602002602001015161284363ffffffff16565b5b8080600101915050614182565b505b5b50505050565b6000896fffffffffffffffffffffffffffffffff168b6fffffffffffffffffffffffffffffffff1614801561421f5750876fffffffffffffffffffffffffffffffff16896fffffffffffffffffffffffffffffffff16145b801561424e5750856fffffffffffffffffffffffffffffffff16876fffffffffffffffffffffffffffffffff16145b801561427d5750836fffffffffffffffffffffffffffffffff16856fffffffffffffffffffffffffffffffff16145b80156142ac5750816fffffffffffffffffffffffffffffffff16836fffffffffffffffffffffffffffffffff16145b156142be57600190506142c8566142c7565b600090506142c8565b5b9a9950505050505050505050565b6060600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061438d5750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806143eb5750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806144495750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806144a15750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b15156144e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016144d99061987e565b60405180910390fd5b60606060600c60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d54cf71866040518263ffffffff1660e01b81526004016145479190619acb565b60006040518083038186803b1580156145605760006000fd5b505afa158015614575573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061459e9190618407565b91508150600c60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637790f6ba866040518263ffffffff1660e01b81526004016146039190619acb565b60006040518083038186803b15801561461c5760006000fd5b505afa158015614631573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061465a9190618407565b9050805060608767ffffffffffffffff811180156146785760006000fd5b506040519080825280602002602001820160405280156146a75781602001602082028036833780820191505090505b50905060008360038151811015156146bb57fe5b60200260200101518460028151811015156146d257fe5b60200260200101518560018151811015156146e957fe5b602002602001015186600081518110151561470057fe5b602002602001015101010190506000600090505b898110156147e8576000876020600050546040516020016147369291906193dd565b6040516020818303038152906040528051906020012060001c90506000602060005054905080858481518110151561476a57fe5b602002602001019090818152602001505061478b8d8261596263ffffffff16565b6020600081815054809291906001019190509090555060008b14156147c2576147bd8d828c898661656963ffffffff16565b6147d8565b6147d78d828c878b8b8861664263ffffffff16565b5b50505b8080600101915050614714565b50819450505050506147f956505050505b5b95945050505050565b600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614806148b85750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806149165750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806149745750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806149cc5750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515614a0d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614a049061987e565b60405180910390fd5b6060600c60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637d54cf71856040518263ffffffff1660e01b8152600401614a709190619acb565b60006040518083038186803b158015614a895760006000fd5b505afa158015614a9e573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614ac79190618407565b90506060600c60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16637790f6ba866040518263ffffffff1660e01b8152600401614b2c9190619acb565b60006040518083038186803b158015614b455760006000fd5b505afa158015614b5a573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614b839190618407565b90506000826003815181101515614b9657fe5b6020026020010151836002815181101515614bad57fe5b6020026020010151846001815181101515614bc457fe5b6020026020010151856000815181101515614bdb57fe5b60200260200101510101019050614bf8878561596263ffffffff16565b614c0d8785888487878b61664263ffffffff16565b7f0a3b86b6b1f8dde0ef4dc4dfa4963ff7e5e83e81e12095b12708055a04ea4d1e87878787604051614c4294939291906195fb565b60405180910390a15050505b5b50505050565b601d6000505481565b6000600c60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050614c90565b90565b600060006000600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1692508250600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1691508150600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905080505b909192565b6000600560005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060009054906101000a900460ff169050614dbf565b92915050565b60606000601a60005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005080549050905080600184038502101515614e5d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614e5490619965565b60405180910390fd5b60008490506000849050601f60005054821115614e7f57601f60005054915081505b600082600183030290506000838286031015614e9f578490508050614ea7565b838201905080505b8367ffffffffffffffff81118015614ebf5760006000fd5b50604051908082528060200260200182016040528015614eee5781602001602082028036833780820191505090505b509550855060006000905060008390505b82811015614f4c57614f178b8261123963ffffffff16565b8883815181101515614f2557fe5b602002602001019090818152602001505081806001019250505b8080600101915050614eff565b50869650505050505050614f61565050505050505b9392505050565b6060600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16148061501f5750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b8061507d5750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806150db5750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b806151335750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515615174576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161516b9061987e565b60405180910390fd5b60608267ffffffffffffffff8111801561518e5760006000fd5b506040519080825280602002602001820160405280156151bd5781602001602082028036833780820191505090505b5090506000600090505b838110156152185760206000505482828151811015156151e357fe5b6020026020010190908181526020015050602060008181505480929190600101919050909055505b80806001019150506151c7565b508091505061522356505b5b919050565b600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156152c1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016152b89061987e565b60405180910390fd5b81600c60005060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600a60006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600c60005060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506001600b600050540160136000506000600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055506001600b6000505401601360005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055505b5050565b600061547e82600260005061695590919063ffffffff16565b9050615485565b919050565b6000339050615494565b90565b816004600050600083815260200190815260200160002060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550808273ffffffffffffffffffffffffffffffffffffffff166155138361216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45b5050565b60006155718260000160005061697d63ffffffff16565b9050615578565b919050565b600061558e8261546563ffffffff16565b15156155cf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016155c69061983c565b60405180910390fd5b60006155e08361216963ffffffff16565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16148061565557508373ffffffffffffffffffffffffffffffffffffffff1661563d84610acb63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16145b8061566c575061566b8185614d2663ffffffff16565b5b91505061567556505b92915050565b61568c83838361699663ffffffff16565b6000601a60005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054905090506000601b60005060008481526020019081526020016000206000505490506000601a60005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506001840381548110151561574e57fe5b906000526020600020900160005b5054905080601a60005060008873ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050838154811015156157b357fe5b906000526020600020900160005b5081909090555081601b6000506000838152602001908152602001600020600050819090905550601a60005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050805480151561583857fe5b600190038181906000526020600020900160005b506000905590556000601a60005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050805490509050601a60005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005085908060018154018082558091505060019003906000526020600020900160005b9091909190915090905580601b6000506000878152602001908152602001600020600050819090905550505050505b505050565b60006159528360000160005083616bd363ffffffff16565b60001c905061595c565b92915050565b615985828260405180602001604052806000815260200150616c5463ffffffff16565b5b5050565b6060600c60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161480615a415750600c60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80615a9f5750600c60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80615afd5750600c60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b80615b555750600a60009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16145b1515615b96576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615b8d9061987e565b60405180910390fd5b6000620186a085811515615ba657fe5b06905060608360028803815181101515615bbc57fe5b6020026020010151821015615dc6576040518060400160405280600381526020017f77696e0000000000000000000000000000000000000000000000000000000000815260200150905080506002871415615c7957615c7489896040518060400160405280600181526020017f53000000000000000000000000000000000000000000000000000000000000008152602001506001896000815181101515615c6057fe5b60200260200101518b616cbe63ffffffff16565b615dc1565b6003871415615cea57615ce589896040518060400160405280600181526020017f41000000000000000000000000000000000000000000000000000000000000008152602001506002896001815181101515615cd157fe5b60200260200101518b616cbe63ffffffff16565b615dc0565b6004871415615d5b57615d5689896040518060400160405280600181526020017f42000000000000000000000000000000000000000000000000000000000000008152602001506003896002815181101515615d4257fe5b60200260200101518b616cbe63ffffffff16565b615dbf565b615dbe89896040518060400160405280600181526020017f43000000000000000000000000000000000000000000000000000000000000008152602001506004896003815181101515615daa57fe5b60200260200101518b616cbe63ffffffff16565b5b5b5b615fbd565b6040518060400160405280600481526020017f6c6f737300000000000000000000000000000000000000000000000000000000815260200150905080506002871415615e7457615e6f89896040518060400160405280600181526020017f41000000000000000000000000000000000000000000000000000000000000008152602001506002896001815181101515615e5b57fe5b60200260200101518b616cbe63ffffffff16565b615fbc565b6003871415615ee557615ee089896040518060400160405280600181526020017f42000000000000000000000000000000000000000000000000000000000000008152602001506003896002815181101515615ecc57fe5b60200260200101518b616cbe63ffffffff16565b615fbb565b6004871415615f5657615f5189896040518060400160405280600181526020017f43000000000000000000000000000000000000000000000000000000000000008152602001506004896003815181101515615f3d57fe5b60200260200101518b616cbe63ffffffff16565b615fba565b615fb989896040518060400160405280600181526020017f44000000000000000000000000000000000000000000000000000000000000008152602001506005896004815181101515615fa557fe5b60200260200101518b616cbe63ffffffff16565b5b5b5b5b50505b5b9695505050505050565b6000600060006000615fe98660000160005086616e1963ffffffff16565b915091508160001c8160001c8090509350935050506160055650505b9250929050565b6000616028846000016000508460001b84616eb763ffffffff16565b60001c9050616032565b9392505050565b600061605082600001600050616f6b63ffffffff16565b9050616057565b919050565b600061606d8261216963ffffffff16565b905061608181600084616f8463ffffffff16565b61609260008361549763ffffffff16565b6000600860005060008481526020019081526020016000206000508054600181600116156101000203166002900490501415156160ec576008600050600083815260200190815260200160002060006160eb9190617a67565b5b61614382600160005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050616f8a90919063ffffffff16565b5061615b826002600050616fb290919063ffffffff16565b5081600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a4505b50565b60606000821415616208576040518060400160405280600181526020017f3000000000000000000000000000000000000000000000000000000000000000815260200150905061630f565b600082905060005b600082141515616238578080600101915050600a8281151561622e57fe5b0491508150616210565b60608167ffffffffffffffff811180156162525760006000fd5b506040519080825280601f01601f1916602001820160405280156162855781602001600182028036833780820191505090505b5090505b600085141515616301576001820391508150600a858115156162a757fe5b0660300160f81b81838151811015156162bc57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600a858115156162f757fe5b0494508450616289565b80935050505061630f565050505b919050565b6060606083905060608390506060815183510167ffffffffffffffff8111801561633e5760006000fd5b506040519080825280601f01601f1916602001820160405280156163715781602001600182028036833780820191505090505b50905060608190506000600090506000600090505b85518110156163f957858181518110151561639d57fe5b602001015160f81c60f81b83838060010194508151811015156163bc57fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b8080600101915050616386565b506000600090505b845181101561647457848181518110151561641857fe5b602001015160f81c60f81b838380600101945081518110151561643757fe5b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053505b8080600101915050616401565b5082955050505050506164875650505050505b92915050565b61649e84848461567b63ffffffff16565b6164b084848484616fda63ffffffff16565b15156164f1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016164e890619776565b60405180910390fd5b5b50505050565b60006000600090505b8351811015616559578260196000506000868481518110151561652057fe5b602002602001015181526020019081526020016000206000505414151561654b576000915050616563565b5b8080600101915050616501565b5060019050616563565b92915050565b6165cc84846040518060400160405280600181526020017f440000000000000000000000000000000000000000000000000000000000000081526020015060058660048151811015156165b857fe5b602002602001015186616cbe63ffffffff16565b7fdaef7511bc11119dbe5492e8578fe54789a6c99529142ee05a97ba56d543e08e85856014600050600088815260200190815260200160002060005054601560005060008981526020019081526020016000206000506040516166329493929190619553565b60405180910390a15b5050505050565b6000848281151561664f57fe5b06905083600081518110151561666157fe5b6020026020010151811115156166d9576166d487876040518060400160405280600181526020017f530000000000000000000000000000000000000000000000000000000000000081526020015060018760008151811015156166c057fe5b602002602001015187616cbe63ffffffff16565b616895565b8360018151811015156166e857fe5b60200260200101518460008151811015156166ff57fe5b602002602001015101811115156167785761677387876040518060400160405280600181526020017f4100000000000000000000000000000000000000000000000000000000000000815260200150600287600181518110151561675f57fe5b602002602001015187616cbe63ffffffff16565b616894565b83600281518110151561678757fe5b602002602001015184600181518110151561679e57fe5b60200260200101518560008151811015156167b557fe5b602002602001015101018111151561682f5761682a87876040518060400160405280600181526020017f4200000000000000000000000000000000000000000000000000000000000000815260200150600387600281518110151561681657fe5b602002602001015187616cbe63ffffffff16565b616893565b61689287876040518060400160405280600181526020017f4300000000000000000000000000000000000000000000000000000000000000815260200150600487600381518110151561687e57fe5b602002602001015187616cbe63ffffffff16565b5b5b5b7f68835e3b708e8bad6ff90bf6e795d6176ce6ad21f963f6c304614a155173dbd48888601460005060008b815260200190815260200160002060005054601560005060008c8152602001908152602001600020600050601360005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546040516169429594939291906195a0565b60405180910390a1505b50505050505050565b6000616970836000016000508360001b6171df63ffffffff16565b9050616977565b92915050565b600081600001600050805490509050616991565b919050565b8273ffffffffffffffffffffffffffffffffffffffff166169bc8261216963ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff16141515616a14576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616a0b90619986565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515616a86576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616a7d906197d9565b60405180910390fd5b616a97838383616f8463ffffffff16565b616aa860008261549763ffffffff16565b616aff81600160005060008673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050616f8a90919063ffffffff16565b50616b5781600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005061721690919063ffffffff16565b50616b718183600260005061723e9092919063ffffffff16565b50808273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45b505050565b6000818360000160005080549050111515616c23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616c1a90619713565b60405180910390fd5b8260000160005082815481101515616c3757fe5b906000526020600020900160005b50549050616c4e565b92915050565b616c64838361728163ffffffff16565b616c776000848484616fda63ffffffff16565b1515616cb8576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616caf90619776565b60405180910390fd5b5b505050565b846014600050600088815260200190815260200160002060005081909090555083601660005060008881526020019081526020016000206000509080519060200190616d0b929190617aaf565b508260176000506000888152602001908152602001600020600050819090905550601360005060003373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505460196000506000888152602001908152602001600020600050819090905550600060018383811515616da057fe5b060190508060186000506000898152602001908152602001600020600050819090905550616de285616dd7836161bd63ffffffff16565b61631463ffffffff16565b601560005060008981526020019081526020016000206000509080519060200190616e0e929190617aaf565b50505b505050505050565b60006000828460000160005080549050111515616e6b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616e6290619902565b60405180910390fd5b60008460000160005084815481101515616e8157fe5b906000526020600020906002020160005b509050806000016000505481600101600050549250925050616eb056505b9250929050565b600060008460010160005060008560001916600019168152602001908152602001600020600050549050600081141583901515616f2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401616f21919061969c565b60405180910390fd5b508460000160005060018203815481101515616f4257fe5b906000526020600020906002020160005b5060010160005054915050616f6456505b9392505050565b600081600001600050805490509050616f7f565b919050565b5b505050565b6000616fa5836000016000508360001b61737163ffffffff16565b9050616fac565b92915050565b6000616fcd836000016000508360001b6174bc63ffffffff16565b9050616fd4565b92915050565b60006170038473ffffffffffffffffffffffffffffffffffffffff1661764e909063ffffffff16565b151561701257600190506171d7565b600060608573ffffffffffffffffffffffffffffffffffffffff1663150b7a0260e01b61704361548a63ffffffff16565b898888604051602401617059949392919061945e565b604051602081830303815290604052907bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516170c391906193a0565b6000604051808303816000865af19150503d8060008114617100576040519150601f19603f3d011682016040523d82523d6000602084013e617105565b606091505b5091509150811515617169576000815111156171295780518082602001fd50617164565b6040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161715b90619776565b60405180910390fd5b6171d4565b60008180602001905181019061717f9190618476565b905063150b7a0260e01b7bffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916817bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19161493505050506171d756505b50505b949350505050565b6000600083600101600050600084600019166000191681526020019081526020016000206000505414159050617210565b92915050565b6000617231836000016000508360001b6176ae63ffffffff16565b9050617238565b92915050565b6000617273846000016000508460001b8473ffffffffffffffffffffffffffffffffffffffff1660001b61774d63ffffffff16565b905061727a565b9392505050565b617291828261788863ffffffff16565b601a60005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081908060018154018082558091505060019003906000526020600020900160005b909190919091509090556001601a60005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508054905003601b60005060008381526020019081526020016000206000508190909055505b5050565b6000600083600101600050600084600019166000191681526020019081526020016000206000505490506000811415156174aa576000600182039050600060018660000160005080549050039050600086600001600050828154811015156173d557fe5b906000526020600020900160005b505490508087600001600050848154811015156173fc57fe5b906000526020600020900160005b508190906000191690555060018301876001016000506000836000191660001916815260200190815260200160002060005081909090555086600001600050805480151561745457fe5b600190038181906000526020600020900160005b5060009055905586600101600050600087600019166000191681526020019081526020016000206000506000905560019450505050506174b6565050506174b4565b60009150506174b6565b505b92915050565b60006000836001016000506000846000191660001916815260200190815260200160002060005054905060008114151561763c5760006001820390506000600186600001600050805490500390506000866000016000508281548110151561752057fe5b906000526020600020906002020160005b50905080876000016000508481548110151561754957fe5b906000526020600020906002020160005b506000820160005054816000016000509060001916905560018201600050548160010160005090600019169055905050600183018760010160005060008360000160005054600019166000191681526020019081526020016000206000508190909055508660000160005080548015156175d057fe5b600190038181906000526020600020906002020160005b6000820160005060009055600182016000506000905550509055866001016000506000876000191660001916815260200190815260200160002060005060009055600194505050505061764856505050617646565b6000915050617648565b505b92915050565b6000600060007fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a47060001b9050833f9150806000191682600019161415801561769e5750600060001b826000191614155b925050506176a95650505b919050565b60006176c08383617a3063ffffffff16565b151561773d578260000160005082908060018154018082558091505060019003906000526020600020900160005b9091909190915090600019169055826000016000508054905083600101600050600084600019166000191681526020019081526020016000206000508190909055506001905061774756617746565b60009050617747565b5b92915050565b600060008460010160005060008560001916600019168152602001908152602001600020600050549050600081141561783b57846000016000506040518060400160405280866000191681526020018560001916815260200150908060018154018082558091505060019003906000526020600020906002020160005b90919091909150600082015181600001600050906000191690556020820151816001016000509060001916905550508460000160005080549050856001016000506000866000191660001916815260200190815260200160002060005081909090555060019150506178815661787f565b82856000016000506001830381548110151561785357fe5b906000526020600020906002020160005b50600101600050819090600019169055506000915050617881565b505b9392505050565b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156178fa576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016178f190619923565b60405180910390fd5b6179098161546563ffffffff16565b15151561794b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161794290619797565b60405180910390fd5b61795d60008383616f8463ffffffff16565b6179b481600160005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005061721690919063ffffffff16565b506179ce8183600260005061723e9092919063ffffffff16565b50808273ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef60405160405180910390a45b5050565b6000600083600101600050600084600019166000191681526020019081526020016000206000505414159050617a61565b92915050565b50805460018160011615610100020316600290046000825580601f10617a8d5750617aac565b601f016020900490600052602060002090810190617aab9190617b34565b5b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10617af057805160ff1916838001178555617b23565b82800160010185558215617b23579182015b82811115617b225782518260005090905591602001919060010190617b02565b5b509050617b309190617b34565b5090565b617b5c9190617b3e565b80821115617b585760008181506000905550600101617b3e565b5090565b9056619ea6565b600081359050617b7281619e04565b5b92915050565b600081359050617b8881619e1f565b5b92915050565b6000600083601f8401121515617ba55760006000fd5b8235905067ffffffffffffffff811115617bbf5760006000fd5b602083019150836020820283011115617bd85760006000fd5b5b9250929050565b600082601f8301121515617bf45760006000fd5b8135617c07617c0282619b6a565b619b3b565b91508181835260208401935060208101905083856020840282011115617c2d5760006000fd5b60005b83811015617c5e5781617c438882617da3565b8452602084019350602083019250505b600181019050617c30565b505050505b92915050565b600082601f8301121515617c7d5760006000fd5b8151617c90617c8b82619b6a565b619b3b565b91508181835260208401935060208101905083856020840282011115617cb65760006000fd5b60005b83811015617ce75781617ccc8882617db9565b8452602084019350602083019250505b600181019050617cb9565b505050505b92915050565b600081359050617d0181619e3a565b5b92915050565b600081359050617d1781619e55565b5b92915050565b600081519050617d2d81619e55565b5b92915050565b600082601f8301121515617d485760006000fd5b8135617d5b617d5682619b94565b619b3b565b91508082526020830160208301858383011115617d785760006000fd5b617d83838284619da2565b5050505b92915050565b600081359050617d9c81619e70565b5b92915050565b600081359050617db281619e8b565b5b92915050565b600081519050617dc881619e8b565b5b92915050565b600060208284031215617de25760006000fd5b6000617df084828501617b63565b9150505b92915050565b6000600060408385031215617e0f5760006000fd5b6000617e1d85828601617b79565b9250506020617e2e85828601617b63565b9150505b9250929050565b60006000600060608486031215617e505760006000fd5b6000617e5e86828701617b79565b9350506020617e6f86828701617b63565b9250506040617e8086828701617b63565b9150505b9250925092565b6000600060408385031215617ea05760006000fd5b6000617eae85828601617b63565b9250506020617ebf85828601617b63565b9150505b9250929050565b600060006000600060608587031215617ee35760006000fd5b6000617ef187828801617b63565b9450506020617f0287828801617b63565b935050604085013567ffffffffffffffff811115617f205760006000fd5b617f2c87828801617b8f565b92509250505b92959194509250565b60006000600060608486031215617f525760006000fd5b6000617f6086828701617b63565b9350506020617f7186828701617b63565b9250506040617f8286828701617da3565b9150505b9250925092565b600060006000600060808587031215617fa65760006000fd5b6000617fb487828801617b63565b9450506020617fc587828801617b63565b9350506040617fd687828801617da3565b925050606085013567ffffffffffffffff811115617ff45760006000fd5b61800087828801617d34565b9150505b92959194509250565b60006000604083850312156180225760006000fd5b600061803085828601617b63565b925050602061804185828601617cf2565b9150505b9250929050565b60006000604083850312156180615760006000fd5b600061806f85828601617b63565b925050602061808085828601617da3565b9150505b9250929050565b600060006000600060006000608087890312156180a85760006000fd5b60006180b689828a01617b63565b96505060206180c789828a01617da3565b955050604087013567ffffffffffffffff8111156180e55760006000fd5b6180f189828a01617b8f565b9450945050606087013567ffffffffffffffff8111156181115760006000fd5b61811d89828a01617b8f565b92509250505b9295509295509295565b600060006000606084860312156181445760006000fd5b600061815286828701617b63565b935050602061816386828701617da3565b925050604061817486828701617da3565b9150505b9250925092565b60006000600060006000600060a0878903121561819c5760006000fd5b60006181aa89828a01617b63565b96505060206181bb89828a01617da3565b95505060406181cc89828a01617da3565b945050606087013567ffffffffffffffff8111156181ea5760006000fd5b6181f689828a01617b8f565b9350935050608061820989828a01617da3565b9150505b9295509295509295565b6000600060006000608085870312156182305760006000fd5b600061823e87828801617b63565b945050602061824f87828801617da3565b935050604061826087828801617da3565b925050606085013567ffffffffffffffff81111561827e5760006000fd5b61828a87828801617be0565b9150505b92959194509250565b6000600060006000608085870312156182b05760006000fd5b60006182be87828801617b63565b94505060206182cf87828801617da3565b93505060406182e087828801617da3565b92505060606182f187828801617da3565b9150505b92959194509250565b6000600060006000600060a086880312156183195760006000fd5b600061832788828901617b63565b955050602061833888828901617da3565b945050604061834988828901617da3565b935050606061835a88828901617da3565b925050608061836b88828901617da3565b9150505b9295509295909350565b60006000600060006000600060c087890312156183965760006000fd5b60006183a489828a01617b63565b96505060206183b589828a01617da3565b95505060406183c689828a01617da3565b94505060606183d789828a01617da3565b93505060806183e889828a01617da3565b92505060a06183f989828a01617da3565b9150505b9295509295509295565b60006020828403121561841a5760006000fd5b600082015167ffffffffffffffff8111156184355760006000fd5b61844184828501617c69565b9150505b92915050565b60006020828403121561845e5760006000fd5b600061846c84828501617d08565b9150505b92915050565b6000602082840312156184895760006000fd5b600061849784828501617d1e565b9150505b92915050565b60006000600060006000600060006000600060006101408b8d0312156184c75760006000fd5b60006184d58d828e01617d8d565b9a505060206184e68d828e01617d8d565b99505060406184f78d828e01617d8d565b98505060606185088d828e01617d8d565b97505060806185198d828e01617d8d565b96505060a061852a8d828e01617d8d565b95505060c061853b8d828e01617d8d565b94505060e061854c8d828e01617d8d565b93505061010061855e8d828e01617d8d565b9250506101206185708d828e01617d8d565b9150505b9295989b9194979a5092959850565b6000602082840312156185965760006000fd5b60006185a484828501617da3565b9150505b92915050565b60006000600060006000600060c087890312156185cb5760006000fd5b60006185d989828a01617da3565b96505060206185ea89828a01617da3565b95505060406185fb89828a01617da3565b945050606061860c89828a01617da3565b935050608061861d89828a01617da3565b92505060a061862e89828a01617da3565b9150505b9295509295509295565b60006186488383619368565b6020830190505b92915050565b61865e81619d1d565b82525b5050565b61866e81619c7c565b82525b5050565b61867e81619c69565b82525b5050565b600061869082619be9565b61869a8185619c1b565b93506186a583619bc2565b8060005b838110156186d75781516186bd888261863c565b97506186c883619c0d565b9250505b6001810190506186a9565b508593505050505b92915050565b6186ee81619c8f565b82525b5050565b6186fe81619c9c565b82525b5050565b600061871082619bf5565b61871a8185619c2d565b935061872a818560208601619db2565b61873381619df2565b84019150505b92915050565b600061874a82619bf5565b6187548185619c3f565b9350618764818560208601619db2565b8084019150505b92915050565b61877a81619d30565b82525b5050565b61878a81619d56565b82525b5050565b600061879c82619c01565b6187a68185619c4b565b93506187b6818560208601619db2565b6187bf81619df2565b84019150505b92915050565b60006187d682619c01565b6187e08185619c5d565b93506187f0818560208601619db2565b8084019150505b92915050565b60008154600181166000811461881a576001811461884057618885565b607f600283041661882b8187619c4b565b955060ff198316865260208601935050618885565b6002820461884e8187619c4b565b955061885985619bd3565b60005b8281101561887c578154818901526001820191505b60208101905061885c565b80880195505050505b50505b92915050565b6000815460018116600081146188ab57600181146188d057618915565b607f60028304166188bc8187619c5d565b955060ff1983168652808601935050618915565b600282046188de8187619c5d565b95506188e985619bd3565b60005b8281101561890c578154818901526001820191505b6020810190506188ec565b82880195505050505b50505b92915050565b600061892b602283619c4b565b91507f456e756d657261626c655365743a20696e646578206f7574206f6620626f756e60008301527f647300000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618992602183619c4b565b91507f4e6674204572723a43616e206e6f742062652072652d696e697469616c697a6560008301527f640000000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006189f9603a83619c4b565b91507f4e6674204572723a204e6f7420616c6c207468652066697665206e667420626560008301527f6c6f6e67696e6720746f207468652073616d652073657269657300000000000060208301526040820190505b919050565b6000618a60603283619c4b565b91507f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560008301527f63656976657220696d706c656d656e746572000000000000000000000000000060208301526040820190505b919050565b6000618ac7601c83619c4b565b91507f4552433732313a20746f6b656e20616c7265616479206d696e7465640000000060008301526020820190505b919050565b6000618b08602483619c4b565b91507f4e6674204572723a20496e73756666696369656e74206e756d626572206f662060008301527f4e4654730000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618b6f602483619c4b565b91507f4552433732313a207472616e7366657220746f20746865207a65726f2061646460008301527f726573730000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618bd6601983619c4b565b91507f4552433732313a20617070726f766520746f2063616c6c65720000000000000060008301526020820190505b919050565b6000618c17603183619c4b565b91507f4e6674204572723a2054686520677261646573206f662074686520666976652060008301527f6e66742061726520646966666572656e7400000000000000000000000000000060208301526040820190505b919050565b6000618c7e602c83619c4b565b91507f4552433732313a206f70657261746f7220717565727920666f72206e6f6e657860008301527f697374656e7420746f6b656e000000000000000000000000000000000000000060208301526040820190505b919050565b6000618ce5603183619c4b565b91507f4e6674204572723a5468652063757272656e63792075736564206973206e6f7460008301527f207468652073616d652076657273696f6e00000000000000000000000000000060208301526040820190505b919050565b6000618d4c601583619c4b565b91507f4e6674204572723a20556e617574686f72757a6564000000000000000000000060008301526020820190505b919050565b6000618d8d602583619c4b565b91507f4e6674204572723a205b5d746f6b656e496473207175616e74697479206e6f7460008301527f206669766500000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618df4603883619c4b565b91507f4552433732313a20617070726f76652063616c6c6572206973206e6f74206f7760008301527f6e6572206e6f7220617070726f76656420666f7220616c6c000000000000000060208301526040820190505b919050565b6000618e5b602a83619c4b565b91507f4552433732313a2062616c616e636520717565727920666f7220746865207a6560008301527f726f20616464726573730000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618ec2602283619c4b565b91507f456e756d657261626c654d61703a20696e646578206f7574206f6620626f756e60008301527f647300000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000618f29602083619c4b565b91507f4552433732313a206d696e7420746f20746865207a65726f206164647265737360008301526020820190505b919050565b6000618f6a602c83619c4b565b91507f4552433732313a20617070726f76656420717565727920666f72206e6f6e657860008301527f697374656e7420746f6b656e000000000000000000000000000000000000000060208301526040820190505b919050565b6000618fd1601483619c4b565b91507f4e6674204572723a204e6f206d6f7265204e465400000000000000000000000060008301526020820190505b919050565b6000619012602983619c4b565b91507f4552433732313a207472616e73666572206f6620746f6b656e2074686174206960008301527f73206e6f74206f776e000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000619079602f83619c4b565b91507f4552433732314d657461646174613a2055524920717565727920666f72206e6f60008301527f6e6578697374656e7420746f6b656e000000000000000000000000000000000060208301526040820190505b919050565b60006190e0603483619c4b565b91507f4572726f723a206661696c65642c20506c6561736520636f6e6669726d20746860008301527f617420616c6c2061726520636f6c6c656374656400000000000000000000000060208301526040820190505b919050565b6000619147601f83619c4b565b91507f4e6674204572723a20546f6b656e696420646f6573206e6f742065786973740060008301526020820190505b919050565b6000619188602183619c4b565b91507f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560008301527f720000000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006191ef603183619c4b565b91507f4552433732313a207472616e736665722063616c6c6572206973206e6f74206f60008301527f776e6572206e6f7220617070726f76656400000000000000000000000000000060208301526040820190505b919050565b6000619256602283619c4b565b91507f4e6674204572723a20746f6b656e4964206e6f2062656c6f6e6720746f20757360008301527f657200000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006192bd603183619c4b565b91507f4e6674204572723a204e6f7420616c6c206e66742062656c6f6e67696e67207460008301527f6f207468652073616d652073657269657300000000000000000000000000000060208301526040820190505b919050565b6000619324601883619c4b565b91507f4e6674204572723a20546f6b656e4944206973204e756c6c000000000000000060008301526020820190505b919050565b61936181619cd4565b82525b5050565b61937181619d12565b82525b5050565b61938181619d12565b82525b5050565b61939961939482619d12565b619de7565b82525b5050565b60006193ac828461873f565b91508190505b92915050565b60006193c4828561888e565b91506193d082846187cb565b91508190505b9392505050565b60006193e98285619388565b6020820191506193f98284619388565b6020820191508190505b9392505050565b600060208201905061941f6000830184618675565b5b92915050565b600060608201905061943b6000830186618655565b6194486020830185618675565b6194556040830184618675565b5b949350505050565b60006080820190506194736000830187618665565b6194806020830186618675565b61948d6040830185619378565b818103606083015261949f8184618705565b90505b95945050505050565b60006060820190506194c06000830186618675565b6194cd6020830185618675565b6194da6040830184618675565b5b949350505050565b600060e0820190506194f8600083018a618675565b6195056020830189618675565b6195126040830188618675565b61951f6060830187618675565b61952c6080830186618675565b61953960a0830185618781565b61954660c0830184618771565b5b98975050505050505050565b60006080820190506195686000830187618675565b6195756020830186619378565b6195826040830185619378565b818103606083015261959481846187fd565b90505b95945050505050565b600060a0820190506195b56000830188618675565b6195c26020830187619378565b6195cf6040830186619378565b81810360608301526195e181856187fd565b90506195f06080830184619378565b5b9695505050505050565b60006080820190506196106000830187618675565b61961d6020830186619378565b61962a6040830185619378565b6196376060830184619378565b5b95945050505050565b6000602082019050818103600083015261965b8184618685565b90505b92915050565b600060208201905061967960008301846186e5565b5b92915050565b600060208201905061969560008301846186f5565b5b92915050565b600060208201905081810360008301526196b68184618791565b90505b92915050565b600060808201905081810360008301526196d98187618791565b90506196e86020830186619378565b6196f56040830185619378565b818103606083015261970781846187fd565b90505b95945050505050565b6000602082019050818103600083015261972c8161891e565b90505b919050565b6000602082019050818103600083015261974d81618985565b90505b919050565b6000602082019050818103600083015261976e816189ec565b90505b919050565b6000602082019050818103600083015261978f81618a53565b90505b919050565b600060208201905081810360008301526197b081618aba565b90505b919050565b600060208201905081810360008301526197d181618afb565b90505b919050565b600060208201905081810360008301526197f281618b62565b90505b919050565b6000602082019050818103600083015261981381618bc9565b90505b919050565b6000602082019050818103600083015261983481618c0a565b90505b919050565b6000602082019050818103600083015261985581618c71565b90505b919050565b6000602082019050818103600083015261987681618cd8565b90505b919050565b6000602082019050818103600083015261989781618d3f565b90505b919050565b600060208201905081810360008301526198b881618d80565b90505b919050565b600060208201905081810360008301526198d981618de7565b90505b919050565b600060208201905081810360008301526198fa81618e4e565b90505b919050565b6000602082019050818103600083015261991b81618eb5565b90505b919050565b6000602082019050818103600083015261993c81618f1c565b90505b919050565b6000602082019050818103600083015261995d81618f5d565b90505b919050565b6000602082019050818103600083015261997e81618fc4565b90505b919050565b6000602082019050818103600083015261999f81619005565b90505b919050565b600060208201905081810360008301526199c08161906c565b90505b919050565b600060208201905081810360008301526199e1816190d3565b90505b919050565b60006020820190508181036000830152619a028161913a565b90505b919050565b60006020820190508181036000830152619a238161917b565b90505b919050565b60006020820190508181036000830152619a44816191e2565b90505b919050565b60006020820190508181036000830152619a6581619249565b90505b919050565b60006020820190508181036000830152619a86816192b0565b90505b919050565b60006020820190508181036000830152619aa781619317565b90505b919050565b6000602082019050619ac46000830184619358565b5b92915050565b6000602082019050619ae06000830184619378565b5b92915050565b6000608082019050619afc6000830187619378565b8181036020830152619b0e8186618791565b90508181036040830152619b228185618791565b9050619b316060830184619378565b5b95945050505050565b6000604051905081810181811067ffffffffffffffff82111715619b5f5760006000fd5b80604052505b919050565b600067ffffffffffffffff821115619b825760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff821115619bac5760006000fd5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b600081905081600052602060002090505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b60006020820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b6000619c7482619cf1565b90505b919050565b6000619c8782619cf1565b90505b919050565b600081151590505b919050565b60008190505b919050565b60007fffffffff00000000000000000000000000000000000000000000000000000000821690505b919050565b60006fffffffffffffffffffffffffffffffff821690505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6000619d2882619d7c565b90505b919050565b6000619d3b82619d43565b90505b919050565b6000619d4e82619cf1565b90505b919050565b6000619d6182619d69565b90505b919050565b6000619d7482619cf1565b90505b919050565b6000619d8782619d8f565b90505b919050565b6000619d9a82619cf1565b90505b919050565b828183376000838301525b505050565b60005b83811015619dd15780820151818401525b602081019050619db5565b83811115619de0576000848401525b505b505050565b60008190505b919050565b6000601f19601f83011690505b919050565b619e0d81619c69565b81141515619e1b5760006000fd5b5b50565b619e2881619c7c565b81141515619e365760006000fd5b5b50565b619e4381619c8f565b81141515619e515760006000fd5b5b50565b619e5e81619ca7565b81141515619e6c5760006000fd5b5b50565b619e7981619cd4565b81141515619e875760006000fd5b5b50565b619e9481619d12565b81141515619ea25760006000fd5b5b50565bfe4552433732313a206f776e657220717565727920666f72206e6f6e6578697374656e7420746f6b656ea2646970667358221220deb8dbd8b5476a9ad91701fa6d5c2608a62ee440740ee6fce24f48b46f0a7e4b64736f6c63430006050033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, name_ string, symbol_ string) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, name_, symbol_)
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

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_Store *StoreCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Store.Contract.BalanceOf(&_Store.CallOpts, owner)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreCaller) BaseURI(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "baseURI")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreSession) BaseURI() (string, error) {
	return _Store.Contract.BaseURI(&_Store.CallOpts)
}

// BaseURI is a free data retrieval call binding the contract method 0x6c0360eb.
//
// Solidity: function baseURI() view returns(string)
func (_Store *StoreCallerSession) BaseURI() (string, error) {
	return _Store.Contract.BaseURI(&_Store.CallOpts)
}

// CheckNftGrade is a free data retrieval call binding the contract method 0x917acc8a.
//
// Solidity: function checkNftGrade(uint256 _grade, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreCaller) CheckNftGrade(opts *bind.CallOpts, _grade *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkNftGrade", _grade, _id1, _id2, _id3, _id4, _id5)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckNftGrade is a free data retrieval call binding the contract method 0x917acc8a.
//
// Solidity: function checkNftGrade(uint256 _grade, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreSession) CheckNftGrade(_grade *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	return _Store.Contract.CheckNftGrade(&_Store.CallOpts, _grade, _id1, _id2, _id3, _id4, _id5)
}

// CheckNftGrade is a free data retrieval call binding the contract method 0x917acc8a.
//
// Solidity: function checkNftGrade(uint256 _grade, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreCallerSession) CheckNftGrade(_grade *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	return _Store.Contract.CheckNftGrade(&_Store.CallOpts, _grade, _id1, _id2, _id3, _id4, _id5)
}

// CheckNftOwner is a free data retrieval call binding the contract method 0x8e8aa3d2.
//
// Solidity: function checkNftOwner(address user, uint256 id_1, uint256 id_2, uint256 id_3, uint256 id_4, uint256 id_5) view returns(bool)
func (_Store *StoreCaller) CheckNftOwner(opts *bind.CallOpts, user common.Address, id_1 *big.Int, id_2 *big.Int, id_3 *big.Int, id_4 *big.Int, id_5 *big.Int) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkNftOwner", user, id_1, id_2, id_3, id_4, id_5)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckNftOwner is a free data retrieval call binding the contract method 0x8e8aa3d2.
//
// Solidity: function checkNftOwner(address user, uint256 id_1, uint256 id_2, uint256 id_3, uint256 id_4, uint256 id_5) view returns(bool)
func (_Store *StoreSession) CheckNftOwner(user common.Address, id_1 *big.Int, id_2 *big.Int, id_3 *big.Int, id_4 *big.Int, id_5 *big.Int) (bool, error) {
	return _Store.Contract.CheckNftOwner(&_Store.CallOpts, user, id_1, id_2, id_3, id_4, id_5)
}

// CheckNftOwner is a free data retrieval call binding the contract method 0x8e8aa3d2.
//
// Solidity: function checkNftOwner(address user, uint256 id_1, uint256 id_2, uint256 id_3, uint256 id_4, uint256 id_5) view returns(bool)
func (_Store *StoreCallerSession) CheckNftOwner(user common.Address, id_1 *big.Int, id_2 *big.Int, id_3 *big.Int, id_4 *big.Int, id_5 *big.Int) (bool, error) {
	return _Store.Contract.CheckNftOwner(&_Store.CallOpts, user, id_1, id_2, id_3, id_4, id_5)
}

// CheckRet is a free data retrieval call binding the contract method 0xdcc64c84.
//
// Solidity: function checkRet(uint128 S_check, uint128 S_res, uint128 A_check, uint128 A_res, uint128 B_check, uint128 B_res, uint128 C_check, uint128 C_res, uint128 D_check, uint128 D_res) pure returns(bool)
func (_Store *StoreCaller) CheckRet(opts *bind.CallOpts, S_check *big.Int, S_res *big.Int, A_check *big.Int, A_res *big.Int, B_check *big.Int, B_res *big.Int, C_check *big.Int, C_res *big.Int, D_check *big.Int, D_res *big.Int) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkRet", S_check, S_res, A_check, A_res, B_check, B_res, C_check, C_res, D_check, D_res)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckRet is a free data retrieval call binding the contract method 0xdcc64c84.
//
// Solidity: function checkRet(uint128 S_check, uint128 S_res, uint128 A_check, uint128 A_res, uint128 B_check, uint128 B_res, uint128 C_check, uint128 C_res, uint128 D_check, uint128 D_res) pure returns(bool)
func (_Store *StoreSession) CheckRet(S_check *big.Int, S_res *big.Int, A_check *big.Int, A_res *big.Int, B_check *big.Int, B_res *big.Int, C_check *big.Int, C_res *big.Int, D_check *big.Int, D_res *big.Int) (bool, error) {
	return _Store.Contract.CheckRet(&_Store.CallOpts, S_check, S_res, A_check, A_res, B_check, B_res, C_check, C_res, D_check, D_res)
}

// CheckRet is a free data retrieval call binding the contract method 0xdcc64c84.
//
// Solidity: function checkRet(uint128 S_check, uint128 S_res, uint128 A_check, uint128 A_res, uint128 B_check, uint128 B_res, uint128 C_check, uint128 C_res, uint128 D_check, uint128 D_res) pure returns(bool)
func (_Store *StoreCallerSession) CheckRet(S_check *big.Int, S_res *big.Int, A_check *big.Int, A_res *big.Int, B_check *big.Int, B_res *big.Int, C_check *big.Int, C_res *big.Int, D_check *big.Int, D_res *big.Int) (bool, error) {
	return _Store.Contract.CheckRet(&_Store.CallOpts, S_check, S_res, A_check, A_res, B_check, B_res, C_check, C_res, D_check, D_res)
}

// CheckSeriesId is a free data retrieval call binding the contract method 0x9c38c5a2.
//
// Solidity: function checkSeriesId(uint256 _seriesId, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreCaller) CheckSeriesId(opts *bind.CallOpts, _seriesId *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "checkSeriesId", _seriesId, _id1, _id2, _id3, _id4, _id5)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckSeriesId is a free data retrieval call binding the contract method 0x9c38c5a2.
//
// Solidity: function checkSeriesId(uint256 _seriesId, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreSession) CheckSeriesId(_seriesId *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	return _Store.Contract.CheckSeriesId(&_Store.CallOpts, _seriesId, _id1, _id2, _id3, _id4, _id5)
}

// CheckSeriesId is a free data retrieval call binding the contract method 0x9c38c5a2.
//
// Solidity: function checkSeriesId(uint256 _seriesId, uint256 _id1, uint256 _id2, uint256 _id3, uint256 _id4, uint256 _id5) pure returns(bool)
func (_Store *StoreCallerSession) CheckSeriesId(_seriesId *big.Int, _id1 *big.Int, _id2 *big.Int, _id3 *big.Int, _id4 *big.Int, _id5 *big.Int) (bool, error) {
	return _Store.Contract.CheckSeriesId(&_Store.CallOpts, _seriesId, _id1, _id2, _id3, _id4, _id5)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lockContract, address blindBox, address flip, address handing, address blindBox_ct, address blindBox_pt)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Owner        common.Address
	LockContract common.Address
	BlindBox     common.Address
	Flip         common.Address
	Handing      common.Address
	BlindBoxCt   common.Address
	BlindBoxPt   common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner        common.Address
		LockContract common.Address
		BlindBox     common.Address
		Flip         common.Address
		Handing      common.Address
		BlindBoxCt   common.Address
		BlindBoxPt   common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LockContract = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.BlindBox = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Flip = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Handing = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.BlindBoxCt = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.BlindBoxPt = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lockContract, address blindBox, address flip, address handing, address blindBox_ct, address blindBox_pt)
func (_Store *StoreSession) Config() (struct {
	Owner        common.Address
	LockContract common.Address
	BlindBox     common.Address
	Flip         common.Address
	Handing      common.Address
	BlindBoxCt   common.Address
	BlindBoxPt   common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lockContract, address blindBox, address flip, address handing, address blindBox_ct, address blindBox_pt)
func (_Store *StoreCallerSession) Config() (struct {
	Owner        common.Address
	LockContract common.Address
	BlindBox     common.Address
	Flip         common.Address
	Handing      common.Address
	BlindBoxCt   common.Address
	BlindBoxPt   common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Store *StoreCaller) Exists(opts *bind.CallOpts, tokenId *big.Int) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "exists", tokenId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Store *StoreSession) Exists(tokenId *big.Int) (bool, error) {
	return _Store.Contract.Exists(&_Store.CallOpts, tokenId)
}

// Exists is a free data retrieval call binding the contract method 0x4f558e79.
//
// Solidity: function exists(uint256 tokenId) view returns(bool)
func (_Store *StoreCallerSession) Exists(tokenId *big.Int) (bool, error) {
	return _Store.Contract.Exists(&_Store.CallOpts, tokenId)
}

// GetAddrAllTokenIds is a free data retrieval call binding the contract method 0xec0f6811.
//
// Solidity: function getAddrAllTokenIds(address _user, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreCaller) GetAddrAllTokenIds(opts *bind.CallOpts, _user common.Address, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAddrAllTokenIds", _user, _pageSize, _page)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAddrAllTokenIds is a free data retrieval call binding the contract method 0xec0f6811.
//
// Solidity: function getAddrAllTokenIds(address _user, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreSession) GetAddrAllTokenIds(_user common.Address, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	return _Store.Contract.GetAddrAllTokenIds(&_Store.CallOpts, _user, _pageSize, _page)
}

// GetAddrAllTokenIds is a free data retrieval call binding the contract method 0xec0f6811.
//
// Solidity: function getAddrAllTokenIds(address _user, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreCallerSession) GetAddrAllTokenIds(_user common.Address, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	return _Store.Contract.GetAddrAllTokenIds(&_Store.CallOpts, _user, _pageSize, _page)
}

// GetAddrIndex is a free data retrieval call binding the contract method 0x8b13004a.
//
// Solidity: function getAddrIndex(uint256 _tokenId) view returns(uint256)
func (_Store *StoreCaller) GetAddrIndex(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAddrIndex", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddrIndex is a free data retrieval call binding the contract method 0x8b13004a.
//
// Solidity: function getAddrIndex(uint256 _tokenId) view returns(uint256)
func (_Store *StoreSession) GetAddrIndex(_tokenId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrIndex(&_Store.CallOpts, _tokenId)
}

// GetAddrIndex is a free data retrieval call binding the contract method 0x8b13004a.
//
// Solidity: function getAddrIndex(uint256 _tokenId) view returns(uint256)
func (_Store *StoreCallerSession) GetAddrIndex(_tokenId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrIndex(&_Store.CallOpts, _tokenId)
}

// GetAddrSeriesTokenIds is a free data retrieval call binding the contract method 0x56aebce7.
//
// Solidity: function getAddrSeriesTokenIds(address _user, uint256 _seriesId, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreCaller) GetAddrSeriesTokenIds(opts *bind.CallOpts, _user common.Address, _seriesId *big.Int, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAddrSeriesTokenIds", _user, _seriesId, _pageSize, _page)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAddrSeriesTokenIds is a free data retrieval call binding the contract method 0x56aebce7.
//
// Solidity: function getAddrSeriesTokenIds(address _user, uint256 _seriesId, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreSession) GetAddrSeriesTokenIds(_user common.Address, _seriesId *big.Int, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	return _Store.Contract.GetAddrSeriesTokenIds(&_Store.CallOpts, _user, _seriesId, _pageSize, _page)
}

// GetAddrSeriesTokenIds is a free data retrieval call binding the contract method 0x56aebce7.
//
// Solidity: function getAddrSeriesTokenIds(address _user, uint256 _seriesId, uint256 _pageSize, uint256 _page) view returns(uint256[] result)
func (_Store *StoreCallerSession) GetAddrSeriesTokenIds(_user common.Address, _seriesId *big.Int, _pageSize *big.Int, _page *big.Int) ([]*big.Int, error) {
	return _Store.Contract.GetAddrSeriesTokenIds(&_Store.CallOpts, _user, _seriesId, _pageSize, _page)
}

// GetAddrTokenId is a free data retrieval call binding the contract method 0x3b3a70cd.
//
// Solidity: function getAddrTokenId(address _user, uint256 index) view returns(uint256)
func (_Store *StoreCaller) GetAddrTokenId(opts *bind.CallOpts, _user common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAddrTokenId", _user, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddrTokenId is a free data retrieval call binding the contract method 0x3b3a70cd.
//
// Solidity: function getAddrTokenId(address _user, uint256 index) view returns(uint256)
func (_Store *StoreSession) GetAddrTokenId(_user common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrTokenId(&_Store.CallOpts, _user, index)
}

// GetAddrTokenId is a free data retrieval call binding the contract method 0x3b3a70cd.
//
// Solidity: function getAddrTokenId(address _user, uint256 index) view returns(uint256)
func (_Store *StoreCallerSession) GetAddrTokenId(_user common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrTokenId(&_Store.CallOpts, _user, index)
}

// GetAddrTokenNumberForSeries is a free data retrieval call binding the contract method 0x7c29c516.
//
// Solidity: function getAddrTokenNumberForSeries(address _user, uint256 _seriesId) view returns(uint256)
func (_Store *StoreCaller) GetAddrTokenNumberForSeries(opts *bind.CallOpts, _user common.Address, _seriesId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getAddrTokenNumberForSeries", _user, _seriesId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAddrTokenNumberForSeries is a free data retrieval call binding the contract method 0x7c29c516.
//
// Solidity: function getAddrTokenNumberForSeries(address _user, uint256 _seriesId) view returns(uint256)
func (_Store *StoreSession) GetAddrTokenNumberForSeries(_user common.Address, _seriesId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrTokenNumberForSeries(&_Store.CallOpts, _user, _seriesId)
}

// GetAddrTokenNumberForSeries is a free data retrieval call binding the contract method 0x7c29c516.
//
// Solidity: function getAddrTokenNumberForSeries(address _user, uint256 _seriesId) view returns(uint256)
func (_Store *StoreCallerSession) GetAddrTokenNumberForSeries(_user common.Address, _seriesId *big.Int) (*big.Int, error) {
	return _Store.Contract.GetAddrTokenNumberForSeries(&_Store.CallOpts, _user, _seriesId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.GetApproved(&_Store.CallOpts, tokenId)
}

// GetNftInfo is a free data retrieval call binding the contract method 0xd5ec2d83.
//
// Solidity: function getNftInfo(uint256 tokenId) view returns(uint256 tSerialNumber, string tTypeNumber, string tGrade, uint256 tGradeId)
func (_Store *StoreCaller) GetNftInfo(opts *bind.CallOpts, tokenId *big.Int) (struct {
	TSerialNumber *big.Int
	TTypeNumber   string
	TGrade        string
	TGradeId      *big.Int
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getNftInfo", tokenId)

	outstruct := new(struct {
		TSerialNumber *big.Int
		TTypeNumber   string
		TGrade        string
		TGradeId      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TSerialNumber = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TTypeNumber = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.TGrade = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.TGradeId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetNftInfo is a free data retrieval call binding the contract method 0xd5ec2d83.
//
// Solidity: function getNftInfo(uint256 tokenId) view returns(uint256 tSerialNumber, string tTypeNumber, string tGrade, uint256 tGradeId)
func (_Store *StoreSession) GetNftInfo(tokenId *big.Int) (struct {
	TSerialNumber *big.Int
	TTypeNumber   string
	TGrade        string
	TGradeId      *big.Int
}, error) {
	return _Store.Contract.GetNftInfo(&_Store.CallOpts, tokenId)
}

// GetNftInfo is a free data retrieval call binding the contract method 0xd5ec2d83.
//
// Solidity: function getNftInfo(uint256 tokenId) view returns(uint256 tSerialNumber, string tTypeNumber, string tGrade, uint256 tGradeId)
func (_Store *StoreCallerSession) GetNftInfo(tokenId *big.Int) (struct {
	TSerialNumber *big.Int
	TTypeNumber   string
	TGrade        string
	TGradeId      *big.Int
}, error) {
	return _Store.Contract.GetNftInfo(&_Store.CallOpts, tokenId)
}

// GetOwnerAddr is a free data retrieval call binding the contract method 0xe6704723.
//
// Solidity: function getOwnerAddr() view returns(address)
func (_Store *StoreCaller) GetOwnerAddr(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getOwnerAddr")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetOwnerAddr is a free data retrieval call binding the contract method 0xe6704723.
//
// Solidity: function getOwnerAddr() view returns(address)
func (_Store *StoreSession) GetOwnerAddr() (common.Address, error) {
	return _Store.Contract.GetOwnerAddr(&_Store.CallOpts)
}

// GetOwnerAddr is a free data retrieval call binding the contract method 0xe6704723.
//
// Solidity: function getOwnerAddr() view returns(address)
func (_Store *StoreCallerSession) GetOwnerAddr() (common.Address, error) {
	return _Store.Contract.GetOwnerAddr(&_Store.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_Store *StoreCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _Store.Contract.IsApprovedForAll(&_Store.CallOpts, owner, operator)
}

// Lastblockhashused is a free data retrieval call binding the contract method 0xafe6ddc0.
//
// Solidity: function lastblockhashused() view returns(bytes32)
func (_Store *StoreCaller) Lastblockhashused(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "lastblockhashused")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Lastblockhashused is a free data retrieval call binding the contract method 0xafe6ddc0.
//
// Solidity: function lastblockhashused() view returns(bytes32)
func (_Store *StoreSession) Lastblockhashused() ([32]byte, error) {
	return _Store.Contract.Lastblockhashused(&_Store.CallOpts)
}

// Lastblockhashused is a free data retrieval call binding the contract method 0xafe6ddc0.
//
// Solidity: function lastblockhashused() view returns(bytes32)
func (_Store *StoreCallerSession) Lastblockhashused() ([32]byte, error) {
	return _Store.Contract.Lastblockhashused(&_Store.CallOpts)
}

// Lastblocknumberused is a free data retrieval call binding the contract method 0xe3223bcb.
//
// Solidity: function lastblocknumberused() view returns(uint256)
func (_Store *StoreCaller) Lastblocknumberused(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "lastblocknumberused")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Lastblocknumberused is a free data retrieval call binding the contract method 0xe3223bcb.
//
// Solidity: function lastblocknumberused() view returns(uint256)
func (_Store *StoreSession) Lastblocknumberused() (*big.Int, error) {
	return _Store.Contract.Lastblocknumberused(&_Store.CallOpts)
}

// Lastblocknumberused is a free data retrieval call binding the contract method 0xe3223bcb.
//
// Solidity: function lastblocknumberused() view returns(uint256)
func (_Store *StoreCallerSession) Lastblocknumberused() (*big.Int, error) {
	return _Store.Contract.Lastblocknumberused(&_Store.CallOpts)
}

// Lastresult is a free data retrieval call binding the contract method 0xbd3491f2.
//
// Solidity: function lastresult() view returns(string)
func (_Store *StoreCaller) Lastresult(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "lastresult")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Lastresult is a free data retrieval call binding the contract method 0xbd3491f2.
//
// Solidity: function lastresult() view returns(string)
func (_Store *StoreSession) Lastresult() (string, error) {
	return _Store.Contract.Lastresult(&_Store.CallOpts)
}

// Lastresult is a free data retrieval call binding the contract method 0xbd3491f2.
//
// Solidity: function lastresult() view returns(string)
func (_Store *StoreCallerSession) Lastresult() (string, error) {
	return _Store.Contract.Lastresult(&_Store.CallOpts)
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

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_Store *StoreCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _Store.Contract.OwnerOf(&_Store.CallOpts, tokenId)
}

// QueryConfig is a free data retrieval call binding the contract method 0xe68f909d.
//
// Solidity: function queryConfig() view returns(address blindBox, address flip, address lockContract)
func (_Store *StoreCaller) QueryConfig(opts *bind.CallOpts) (struct {
	BlindBox     common.Address
	Flip         common.Address
	LockContract common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "queryConfig")

	outstruct := new(struct {
		BlindBox     common.Address
		Flip         common.Address
		LockContract common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BlindBox = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Flip = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LockContract = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// QueryConfig is a free data retrieval call binding the contract method 0xe68f909d.
//
// Solidity: function queryConfig() view returns(address blindBox, address flip, address lockContract)
func (_Store *StoreSession) QueryConfig() (struct {
	BlindBox     common.Address
	Flip         common.Address
	LockContract common.Address
}, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// QueryConfig is a free data retrieval call binding the contract method 0xe68f909d.
//
// Solidity: function queryConfig() view returns(address blindBox, address flip, address lockContract)
func (_Store *StoreCallerSession) QueryConfig() (struct {
	BlindBox     common.Address
	Flip         common.Address
	LockContract common.Address
}, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// Reckon is a free data retrieval call binding the contract method 0x382e5acd.
//
// Solidity: function reckon(uint256 num) pure returns(uint128 res)
func (_Store *StoreCaller) Reckon(opts *bind.CallOpts, num *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "reckon", num)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Reckon is a free data retrieval call binding the contract method 0x382e5acd.
//
// Solidity: function reckon(uint256 num) pure returns(uint128 res)
func (_Store *StoreSession) Reckon(num *big.Int) (*big.Int, error) {
	return _Store.Contract.Reckon(&_Store.CallOpts, num)
}

// Reckon is a free data retrieval call binding the contract method 0x382e5acd.
//
// Solidity: function reckon(uint256 num) pure returns(uint128 res)
func (_Store *StoreCallerSession) Reckon(num *big.Int) (*big.Int, error) {
	return _Store.Contract.Reckon(&_Store.CallOpts, num)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Store *StoreCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Store.Contract.SupportsInterface(&_Store.CallOpts, interfaceId)
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

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenByIndex(&_Store.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_Store *StoreCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenByIndex(&_Store.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenOfOwnerByIndex(&_Store.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_Store *StoreCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _Store.Contract.TokenOfOwnerByIndex(&_Store.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_Store *StoreCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _Store.Contract.TokenURI(&_Store.CallOpts, tokenId)
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

// Draw is a paid mutator transaction binding the contract method 0xdd26c818.
//
// Solidity: function Draw(address to, uint256 number, uint256 IsDCard, uint256 _seriesId, uint256 rand) returns(uint256[] tokenIdArray)
func (_Store *StoreTransactor) Draw(opts *bind.TransactOpts, to common.Address, number *big.Int, IsDCard *big.Int, _seriesId *big.Int, rand *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Draw", to, number, IsDCard, _seriesId, rand)
}

// Draw is a paid mutator transaction binding the contract method 0xdd26c818.
//
// Solidity: function Draw(address to, uint256 number, uint256 IsDCard, uint256 _seriesId, uint256 rand) returns(uint256[] tokenIdArray)
func (_Store *StoreSession) Draw(to common.Address, number *big.Int, IsDCard *big.Int, _seriesId *big.Int, rand *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Draw(&_Store.TransactOpts, to, number, IsDCard, _seriesId, rand)
}

// Draw is a paid mutator transaction binding the contract method 0xdd26c818.
//
// Solidity: function Draw(address to, uint256 number, uint256 IsDCard, uint256 _seriesId, uint256 rand) returns(uint256[] tokenIdArray)
func (_Store *StoreTransactorSession) Draw(to common.Address, number *big.Int, IsDCard *big.Int, _seriesId *big.Int, rand *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Draw(&_Store.TransactOpts, to, number, IsDCard, _seriesId, rand)
}

// DrawPro is a paid mutator transaction binding the contract method 0x58c61c02.
//
// Solidity: function DrawPro() returns(uint256 tokenId)
func (_Store *StoreTransactor) DrawPro(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DrawPro")
}

// DrawPro is a paid mutator transaction binding the contract method 0x58c61c02.
//
// Solidity: function DrawPro() returns(uint256 tokenId)
func (_Store *StoreSession) DrawPro() (*types.Transaction, error) {
	return _Store.Contract.DrawPro(&_Store.TransactOpts)
}

// DrawPro is a paid mutator transaction binding the contract method 0x58c61c02.
//
// Solidity: function DrawPro() returns(uint256 tokenId)
func (_Store *StoreTransactorSession) DrawPro() (*types.Transaction, error) {
	return _Store.Contract.DrawPro(&_Store.TransactOpts)
}

// DrawPros is a paid mutator transaction binding the contract method 0xf3c3fceb.
//
// Solidity: function DrawPros(uint256 number) returns(uint256[] tokenIdArray)
func (_Store *StoreTransactor) DrawPros(opts *bind.TransactOpts, number *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DrawPros", number)
}

// DrawPros is a paid mutator transaction binding the contract method 0xf3c3fceb.
//
// Solidity: function DrawPros(uint256 number) returns(uint256[] tokenIdArray)
func (_Store *StoreSession) DrawPros(number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DrawPros(&_Store.TransactOpts, number)
}

// DrawPros is a paid mutator transaction binding the contract method 0xf3c3fceb.
//
// Solidity: function DrawPros(uint256 number) returns(uint256[] tokenIdArray)
func (_Store *StoreTransactorSession) DrawPros(number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DrawPros(&_Store.TransactOpts, number)
}

// MintMixPro is a paid mutator transaction binding the contract method 0x5c4d06e3.
//
// Solidity: function MintMixPro(address user, uint256 seriesId, uint256 grade, uint256 _rand, uint256 _tokenId) returns()
func (_Store *StoreTransactor) MintMixPro(opts *bind.TransactOpts, user common.Address, seriesId *big.Int, grade *big.Int, _rand *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "MintMixPro", user, seriesId, grade, _rand, _tokenId)
}

// MintMixPro is a paid mutator transaction binding the contract method 0x5c4d06e3.
//
// Solidity: function MintMixPro(address user, uint256 seriesId, uint256 grade, uint256 _rand, uint256 _tokenId) returns()
func (_Store *StoreSession) MintMixPro(user common.Address, seriesId *big.Int, grade *big.Int, _rand *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintMixPro(&_Store.TransactOpts, user, seriesId, grade, _rand, _tokenId)
}

// MintMixPro is a paid mutator transaction binding the contract method 0x5c4d06e3.
//
// Solidity: function MintMixPro(address user, uint256 seriesId, uint256 grade, uint256 _rand, uint256 _tokenId) returns()
func (_Store *StoreTransactorSession) MintMixPro(user common.Address, seriesId *big.Int, grade *big.Int, _rand *big.Int, _tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintMixPro(&_Store.TransactOpts, user, seriesId, grade, _rand, _tokenId)
}

// MintPro is a paid mutator transaction binding the contract method 0xe223809a.
//
// Solidity: function MintPro(address to, uint256 _seriesId, uint256 rand, uint256 tokenId) returns()
func (_Store *StoreTransactor) MintPro(opts *bind.TransactOpts, to common.Address, _seriesId *big.Int, rand *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "MintPro", to, _seriesId, rand, tokenId)
}

// MintPro is a paid mutator transaction binding the contract method 0xe223809a.
//
// Solidity: function MintPro(address to, uint256 _seriesId, uint256 rand, uint256 tokenId) returns()
func (_Store *StoreSession) MintPro(to common.Address, _seriesId *big.Int, rand *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintPro(&_Store.TransactOpts, to, _seriesId, rand, tokenId)
}

// MintPro is a paid mutator transaction binding the contract method 0xe223809a.
//
// Solidity: function MintPro(address to, uint256 _seriesId, uint256 rand, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) MintPro(to common.Address, _seriesId *big.Int, rand *big.Int, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintPro(&_Store.TransactOpts, to, _seriesId, rand, tokenId)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address blindBoxPromotion, address handing) returns()
func (_Store *StoreTransactor) AddBoxPro(opts *bind.TransactOpts, blindBoxPromotion common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addBoxPro", blindBoxPromotion, handing)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address blindBoxPromotion, address handing) returns()
func (_Store *StoreSession) AddBoxPro(blindBoxPromotion common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBoxPro(&_Store.TransactOpts, blindBoxPromotion, handing)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address blindBoxPromotion, address handing) returns()
func (_Store *StoreTransactorSession) AddBoxPro(blindBoxPromotion common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBoxPro(&_Store.TransactOpts, blindBoxPromotion, handing)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Approve(&_Store.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Store *StoreTransactor) Burn(opts *bind.TransactOpts, user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burn", user, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Store *StoreSession) Burn(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, user, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) Burn(user common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.Burn(&_Store.TransactOpts, user, tokenId)
}

// CashCheckByTokenID is a paid mutator transaction binding the contract method 0xa0c55e31.
//
// Solidity: function cashCheckByTokenID(address _user, uint256 _seriesId, uint256[] _gradeNumbers, uint256[] _tokenIds) returns()
func (_Store *StoreTransactor) CashCheckByTokenID(opts *bind.TransactOpts, _user common.Address, _seriesId *big.Int, _gradeNumbers []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "cashCheckByTokenID", _user, _seriesId, _gradeNumbers, _tokenIds)
}

// CashCheckByTokenID is a paid mutator transaction binding the contract method 0xa0c55e31.
//
// Solidity: function cashCheckByTokenID(address _user, uint256 _seriesId, uint256[] _gradeNumbers, uint256[] _tokenIds) returns()
func (_Store *StoreSession) CashCheckByTokenID(_user common.Address, _seriesId *big.Int, _gradeNumbers []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.CashCheckByTokenID(&_Store.TransactOpts, _user, _seriesId, _gradeNumbers, _tokenIds)
}

// CashCheckByTokenID is a paid mutator transaction binding the contract method 0xa0c55e31.
//
// Solidity: function cashCheckByTokenID(address _user, uint256 _seriesId, uint256[] _gradeNumbers, uint256[] _tokenIds) returns()
func (_Store *StoreTransactorSession) CashCheckByTokenID(_user common.Address, _seriesId *big.Int, _gradeNumbers []*big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.CashCheckByTokenID(&_Store.TransactOpts, _user, _seriesId, _gradeNumbers, _tokenIds)
}

// CheckMix is a paid mutator transaction binding the contract method 0xdc84e07d.
//
// Solidity: function checkMix(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds) returns()
func (_Store *StoreTransactor) CheckMix(opts *bind.TransactOpts, user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "checkMix", user, seriesId, _grade, _tokenIds)
}

// CheckMix is a paid mutator transaction binding the contract method 0xdc84e07d.
//
// Solidity: function checkMix(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds) returns()
func (_Store *StoreSession) CheckMix(user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.CheckMix(&_Store.TransactOpts, user, seriesId, _grade, _tokenIds)
}

// CheckMix is a paid mutator transaction binding the contract method 0xdc84e07d.
//
// Solidity: function checkMix(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds) returns()
func (_Store *StoreTransactorSession) CheckMix(user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.CheckMix(&_Store.TransactOpts, user, seriesId, _grade, _tokenIds)
}

// GradeCompose is a paid mutator transaction binding the contract method 0x4e6d4e41.
//
// Solidity: function gradeCompose(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds, uint256 _rand) returns(uint256 mixTokenId)
func (_Store *StoreTransactor) GradeCompose(opts *bind.TransactOpts, user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int, _rand *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "gradeCompose", user, seriesId, _grade, _tokenIds, _rand)
}

// GradeCompose is a paid mutator transaction binding the contract method 0x4e6d4e41.
//
// Solidity: function gradeCompose(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds, uint256 _rand) returns(uint256 mixTokenId)
func (_Store *StoreSession) GradeCompose(user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int, _rand *big.Int) (*types.Transaction, error) {
	return _Store.Contract.GradeCompose(&_Store.TransactOpts, user, seriesId, _grade, _tokenIds, _rand)
}

// GradeCompose is a paid mutator transaction binding the contract method 0x4e6d4e41.
//
// Solidity: function gradeCompose(address user, uint256 seriesId, uint256 _grade, uint256[] _tokenIds, uint256 _rand) returns(uint256 mixTokenId)
func (_Store *StoreTransactorSession) GradeCompose(user common.Address, seriesId *big.Int, _grade *big.Int, _tokenIds []*big.Int, _rand *big.Int) (*types.Transaction, error) {
	return _Store.Contract.GradeCompose(&_Store.TransactOpts, user, seriesId, _grade, _tokenIds, _rand)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address blindBox, address flip, address lock_contract) returns()
func (_Store *StoreTransactor) Init(opts *bind.TransactOpts, blindBox common.Address, flip common.Address, lock_contract common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "init", blindBox, flip, lock_contract)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address blindBox, address flip, address lock_contract) returns()
func (_Store *StoreSession) Init(blindBox common.Address, flip common.Address, lock_contract common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, blindBox, flip, lock_contract)
}

// Init is a paid mutator transaction binding the contract method 0x184b9559.
//
// Solidity: function init(address blindBox, address flip, address lock_contract) returns()
func (_Store *StoreTransactorSession) Init(blindBox common.Address, flip common.Address, lock_contract common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, blindBox, flip, lock_contract)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_Store *StoreTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _Store.Contract.SafeTransferFrom0(&_Store.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_Store *StoreTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Store.Contract.SetApprovalForAll(&_Store.TransactOpts, operator, approved)
}

// TransferArray is a paid mutator transaction binding the contract method 0x54af0187.
//
// Solidity: function transferArray(address from, address to, uint256[] tokenIds) returns()
func (_Store *StoreTransactor) TransferArray(opts *bind.TransactOpts, from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferArray", from, to, tokenIds)
}

// TransferArray is a paid mutator transaction binding the contract method 0x54af0187.
//
// Solidity: function transferArray(address from, address to, uint256[] tokenIds) returns()
func (_Store *StoreSession) TransferArray(from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferArray(&_Store.TransactOpts, from, to, tokenIds)
}

// TransferArray is a paid mutator transaction binding the contract method 0x54af0187.
//
// Solidity: function transferArray(address from, address to, uint256[] tokenIds) returns()
func (_Store *StoreTransactorSession) TransferArray(from common.Address, to common.Address, tokenIds []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferArray(&_Store.TransactOpts, from, to, tokenIds)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_Store *StoreTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _Store.Contract.TransferFrom(&_Store.TransactOpts, from, to, tokenId)
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
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StoreApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalIterator{contract: _Store.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StoreApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
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
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseApproval(log types.Log) (*StoreApproval, error) {
	event := new(StoreApproval)
	if err := _Store.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the Store contract.
type StoreApprovalForAllIterator struct {
	Event *StoreApprovalForAll // Event containing the contract specifics and raw log

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
func (it *StoreApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreApprovalForAll)
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
		it.Event = new(StoreApprovalForAll)
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
func (it *StoreApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreApprovalForAll represents a ApprovalForAll event raised by the Store contract.
type StoreApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StoreApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StoreApprovalForAllIterator{contract: _Store.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StoreApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreApprovalForAll)
				if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_Store *StoreFilterer) ParseApprovalForAll(log types.Log) (*StoreApprovalForAll, error) {
	event := new(StoreApprovalForAll)
	if err := _Store.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDrawCardIterator is returned from FilterDrawCard and is used to iterate over the raw logs and unpacked data for DrawCard events raised by the Store contract.
type StoreDrawCardIterator struct {
	Event *StoreDrawCard // Event containing the contract specifics and raw log

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
func (it *StoreDrawCardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDrawCard)
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
		it.Event = new(StoreDrawCard)
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
func (it *StoreDrawCardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDrawCardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDrawCard represents a DrawCard event raised by the Store contract.
type StoreDrawCard struct {
	User              common.Address
	TokenId           *big.Int
	TokenSerialNumber *big.Int
	TokenTypeNumber   string
	TokenVersion      *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDrawCard is a free log retrieval operation binding the contract event 0x68835e3b708e8bad6ff90bf6e795d6176ce6ad21f963f6c304614a155173dbd4.
//
// Solidity: event DrawCard(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber, uint256 tokenVersion)
func (_Store *StoreFilterer) FilterDrawCard(opts *bind.FilterOpts) (*StoreDrawCardIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DrawCard")
	if err != nil {
		return nil, err
	}
	return &StoreDrawCardIterator{contract: _Store.contract, event: "DrawCard", logs: logs, sub: sub}, nil
}

// WatchDrawCard is a free log subscription operation binding the contract event 0x68835e3b708e8bad6ff90bf6e795d6176ce6ad21f963f6c304614a155173dbd4.
//
// Solidity: event DrawCard(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber, uint256 tokenVersion)
func (_Store *StoreFilterer) WatchDrawCard(opts *bind.WatchOpts, sink chan<- *StoreDrawCard) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DrawCard")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDrawCard)
				if err := _Store.contract.UnpackLog(event, "DrawCard", log); err != nil {
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

// ParseDrawCard is a log parse operation binding the contract event 0x68835e3b708e8bad6ff90bf6e795d6176ce6ad21f963f6c304614a155173dbd4.
//
// Solidity: event DrawCard(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber, uint256 tokenVersion)
func (_Store *StoreFilterer) ParseDrawCard(log types.Log) (*StoreDrawCard, error) {
	event := new(StoreDrawCard)
	if err := _Store.contract.UnpackLog(event, "DrawCard", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDrawCardForDIterator is returned from FilterDrawCardForD and is used to iterate over the raw logs and unpacked data for DrawCardForD events raised by the Store contract.
type StoreDrawCardForDIterator struct {
	Event *StoreDrawCardForD // Event containing the contract specifics and raw log

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
func (it *StoreDrawCardForDIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDrawCardForD)
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
		it.Event = new(StoreDrawCardForD)
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
func (it *StoreDrawCardForDIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDrawCardForDIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDrawCardForD represents a DrawCardForD event raised by the Store contract.
type StoreDrawCardForD struct {
	User              common.Address
	TokenId           *big.Int
	TokenSerialNumber *big.Int
	TokenTypeNumber   string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterDrawCardForD is a free log retrieval operation binding the contract event 0xdaef7511bc11119dbe5492e8578fe54789a6c99529142ee05a97ba56d543e08e.
//
// Solidity: event DrawCardForD(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) FilterDrawCardForD(opts *bind.FilterOpts) (*StoreDrawCardForDIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "DrawCardForD")
	if err != nil {
		return nil, err
	}
	return &StoreDrawCardForDIterator{contract: _Store.contract, event: "DrawCardForD", logs: logs, sub: sub}, nil
}

// WatchDrawCardForD is a free log subscription operation binding the contract event 0xdaef7511bc11119dbe5492e8578fe54789a6c99529142ee05a97ba56d543e08e.
//
// Solidity: event DrawCardForD(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) WatchDrawCardForD(opts *bind.WatchOpts, sink chan<- *StoreDrawCardForD) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "DrawCardForD")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDrawCardForD)
				if err := _Store.contract.UnpackLog(event, "DrawCardForD", log); err != nil {
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

// ParseDrawCardForD is a log parse operation binding the contract event 0xdaef7511bc11119dbe5492e8578fe54789a6c99529142ee05a97ba56d543e08e.
//
// Solidity: event DrawCardForD(address user, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) ParseDrawCardForD(log types.Log) (*StoreDrawCardForD, error) {
	event := new(StoreDrawCardForD)
	if err := _Store.contract.UnpackLog(event, "DrawCardForD", log); err != nil {
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
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StoreTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StoreTransferIterator{contract: _Store.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StoreTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Store.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_Store *StoreFilterer) ParseTransfer(log types.Log) (*StoreTransfer, error) {
	event := new(StoreTransfer)
	if err := _Store.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreCashCheckByTokenIdLogIterator is returned from FilterCashCheckByTokenIdLog and is used to iterate over the raw logs and unpacked data for CashCheckByTokenIdLog events raised by the Store contract.
type StoreCashCheckByTokenIdLogIterator struct {
	Event *StoreCashCheckByTokenIdLog // Event containing the contract specifics and raw log

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
func (it *StoreCashCheckByTokenIdLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreCashCheckByTokenIdLog)
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
		it.Event = new(StoreCashCheckByTokenIdLog)
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
func (it *StoreCashCheckByTokenIdLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreCashCheckByTokenIdLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreCashCheckByTokenIdLog represents a CashCheckByTokenIdLog event raised by the Store contract.
type StoreCashCheckByTokenIdLog struct {
	SeriesId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterCashCheckByTokenIdLog is a free log retrieval operation binding the contract event 0x07d6c7fbbec58f3fafbf4721d64886946128f0518351d7db11dfae01ed9dc48d.
//
// Solidity: event cashCheckByTokenIdLog(uint256 _seriesId)
func (_Store *StoreFilterer) FilterCashCheckByTokenIdLog(opts *bind.FilterOpts) (*StoreCashCheckByTokenIdLogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "cashCheckByTokenIdLog")
	if err != nil {
		return nil, err
	}
	return &StoreCashCheckByTokenIdLogIterator{contract: _Store.contract, event: "cashCheckByTokenIdLog", logs: logs, sub: sub}, nil
}

// WatchCashCheckByTokenIdLog is a free log subscription operation binding the contract event 0x07d6c7fbbec58f3fafbf4721d64886946128f0518351d7db11dfae01ed9dc48d.
//
// Solidity: event cashCheckByTokenIdLog(uint256 _seriesId)
func (_Store *StoreFilterer) WatchCashCheckByTokenIdLog(opts *bind.WatchOpts, sink chan<- *StoreCashCheckByTokenIdLog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "cashCheckByTokenIdLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreCashCheckByTokenIdLog)
				if err := _Store.contract.UnpackLog(event, "cashCheckByTokenIdLog", log); err != nil {
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

// ParseCashCheckByTokenIdLog is a log parse operation binding the contract event 0x07d6c7fbbec58f3fafbf4721d64886946128f0518351d7db11dfae01ed9dc48d.
//
// Solidity: event cashCheckByTokenIdLog(uint256 _seriesId)
func (_Store *StoreFilterer) ParseCashCheckByTokenIdLog(log types.Log) (*StoreCashCheckByTokenIdLog, error) {
	event := new(StoreCashCheckByTokenIdLog)
	if err := _Store.contract.UnpackLog(event, "cashCheckByTokenIdLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreGradeComposelogIterator is returned from FilterGradeComposelog and is used to iterate over the raw logs and unpacked data for GradeComposelog events raised by the Store contract.
type StoreGradeComposelogIterator struct {
	Event *StoreGradeComposelog // Event containing the contract specifics and raw log

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
func (it *StoreGradeComposelogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreGradeComposelog)
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
		it.Event = new(StoreGradeComposelog)
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
func (it *StoreGradeComposelogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreGradeComposelogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreGradeComposelog represents a GradeComposelog event raised by the Store contract.
type StoreGradeComposelog struct {
	Res               string
	TokenId           *big.Int
	TokenSerialNumber *big.Int
	TokenTypeNumber   string
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterGradeComposelog is a free log retrieval operation binding the contract event 0xe0eb607747eb5726df347537f089eca562d68096cd2a46fd1dddb1279975e1d2.
//
// Solidity: event gradeComposelog(string res, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) FilterGradeComposelog(opts *bind.FilterOpts) (*StoreGradeComposelogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "gradeComposelog")
	if err != nil {
		return nil, err
	}
	return &StoreGradeComposelogIterator{contract: _Store.contract, event: "gradeComposelog", logs: logs, sub: sub}, nil
}

// WatchGradeComposelog is a free log subscription operation binding the contract event 0xe0eb607747eb5726df347537f089eca562d68096cd2a46fd1dddb1279975e1d2.
//
// Solidity: event gradeComposelog(string res, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) WatchGradeComposelog(opts *bind.WatchOpts, sink chan<- *StoreGradeComposelog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "gradeComposelog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreGradeComposelog)
				if err := _Store.contract.UnpackLog(event, "gradeComposelog", log); err != nil {
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

// ParseGradeComposelog is a log parse operation binding the contract event 0xe0eb607747eb5726df347537f089eca562d68096cd2a46fd1dddb1279975e1d2.
//
// Solidity: event gradeComposelog(string res, uint256 tokenId, uint256 tokenSerialNumber, string tokenTypeNumber)
func (_Store *StoreFilterer) ParseGradeComposelog(log types.Log) (*StoreGradeComposelog, error) {
	event := new(StoreGradeComposelog)
	if err := _Store.contract.UnpackLog(event, "gradeComposelog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreInitLogIterator is returned from FilterInitLog and is used to iterate over the raw logs and unpacked data for InitLog events raised by the Store contract.
type StoreInitLogIterator struct {
	Event *StoreInitLog // Event containing the contract specifics and raw log

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
func (it *StoreInitLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreInitLog)
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
		it.Event = new(StoreInitLog)
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
func (it *StoreInitLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreInitLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreInitLog represents a InitLog event raised by the Store contract.
type StoreInitLog struct {
	AddrOne   common.Address
	AddrTwo   common.Address
	AddrThree common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterInitLog is a free log retrieval operation binding the contract event 0x8ab07c0c846b8760fd84f71bc8e5f506d49de3e596c4561fd19f9a99bc42c00e.
//
// Solidity: event initLog(address addrOne, address addrTwo, address addrThree)
func (_Store *StoreFilterer) FilterInitLog(opts *bind.FilterOpts) (*StoreInitLogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "initLog")
	if err != nil {
		return nil, err
	}
	return &StoreInitLogIterator{contract: _Store.contract, event: "initLog", logs: logs, sub: sub}, nil
}

// WatchInitLog is a free log subscription operation binding the contract event 0x8ab07c0c846b8760fd84f71bc8e5f506d49de3e596c4561fd19f9a99bc42c00e.
//
// Solidity: event initLog(address addrOne, address addrTwo, address addrThree)
func (_Store *StoreFilterer) WatchInitLog(opts *bind.WatchOpts, sink chan<- *StoreInitLog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "initLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreInitLog)
				if err := _Store.contract.UnpackLog(event, "initLog", log); err != nil {
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

// ParseInitLog is a log parse operation binding the contract event 0x8ab07c0c846b8760fd84f71bc8e5f506d49de3e596c4561fd19f9a99bc42c00e.
//
// Solidity: event initLog(address addrOne, address addrTwo, address addrThree)
func (_Store *StoreFilterer) ParseInitLog(log types.Log) (*StoreInitLog, error) {
	event := new(StoreInitLog)
	if err := _Store.contract.UnpackLog(event, "initLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreMintProLogIterator is returned from FilterMintProLog and is used to iterate over the raw logs and unpacked data for MintProLog events raised by the Store contract.
type StoreMintProLogIterator struct {
	Event *StoreMintProLog // Event containing the contract specifics and raw log

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
func (it *StoreMintProLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMintProLog)
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
		it.Event = new(StoreMintProLog)
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
func (it *StoreMintProLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMintProLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMintProLog represents a MintProLog event raised by the Store contract.
type StoreMintProLog struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMintProLog is a free log retrieval operation binding the contract event 0x0a3b86b6b1f8dde0ef4dc4dfa4963ff7e5e83e81e12095b12708055a04ea4d1e.
//
// Solidity: event mint_pro_log(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Store *StoreFilterer) FilterMintProLog(opts *bind.FilterOpts) (*StoreMintProLogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "mint_pro_log")
	if err != nil {
		return nil, err
	}
	return &StoreMintProLogIterator{contract: _Store.contract, event: "mint_pro_log", logs: logs, sub: sub}, nil
}

// WatchMintProLog is a free log subscription operation binding the contract event 0x0a3b86b6b1f8dde0ef4dc4dfa4963ff7e5e83e81e12095b12708055a04ea4d1e.
//
// Solidity: event mint_pro_log(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Store *StoreFilterer) WatchMintProLog(opts *bind.WatchOpts, sink chan<- *StoreMintProLog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "mint_pro_log")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMintProLog)
				if err := _Store.contract.UnpackLog(event, "mint_pro_log", log); err != nil {
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

// ParseMintProLog is a log parse operation binding the contract event 0x0a3b86b6b1f8dde0ef4dc4dfa4963ff7e5e83e81e12095b12708055a04ea4d1e.
//
// Solidity: event mint_pro_log(address arg0, uint256 arg1, uint256 arg2, uint256 arg3)
func (_Store *StoreFilterer) ParseMintProLog(log types.Log) (*StoreMintProLog, error) {
	event := new(StoreMintProLog)
	if err := _Store.contract.UnpackLog(event, "mint_pro_log", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
