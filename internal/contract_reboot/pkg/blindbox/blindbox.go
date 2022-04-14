// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blindbox

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

// BlindBoxBox is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxBox struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Level    []*big.Int
	Draw     []*big.Int
	Mix      []*big.Int
	Reward   BlindBoxReward
}

// BlindBoxConfig is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxConfig struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
	ITokenId      common.Address
	BoxPro        common.Address
}

// BlindBoxReward is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxReward struct {
	Token  []common.Address
	Amount []*big.Int
}

// BuildTokenControlledTokenConfig is an auto generated low-level Go binding around an user-defined struct.
type BuildTokenControlledTokenConfig struct {
	Name       string
	Symbol     string
	Decimals   uint8
	Controller common.Address
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lableAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"contractBuildToken\",\"name\":\"_controlledTokenBuilder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_itokenid\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"add_box_pro\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"draw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"draw_out\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"mint_box\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"mix_true\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetDraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetLevel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetMix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"resetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structBlindBox.Reward\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"resetReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reset_ratio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_token_ids\",\"type\":\"uint256[]\"}],\"name\":\"Convert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAUL_DECIMAL_PLACES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"Draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DrawNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"DrawOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIX_TRUE_LOW_LEVEL_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box\",\"name\":\"_box\",\"type\":\"tuple\"}],\"name\":\"MintBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokens_id\",\"type\":\"uint256[]\"}],\"name\":\"MixTrue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryBox\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"QueryBoxs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QueryConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"iTokenId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"box_pro\",\"type\":\"address\"}],\"internalType\":\"structBlindBox.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryDraws\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryImage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryLevels\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryMix\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QueryRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QuerySeriesIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QuerySeriesIdsNotNull\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_draw\",\"type\":\"uint256[]\"}],\"name\":\"ResetDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_drawNumber\",\"type\":\"uint256\"}],\"name\":\"ResetDrawNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_mix\",\"type\":\"uint256[]\"}],\"name\":\"ResetMix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ResetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"ResetRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"_reward\",\"type\":\"tuple\"}],\"name\":\"ResetReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_box_pro\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handing\",\"type\":\"address\"}],\"name\":\"addBoxPro\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"box_info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"burnKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"iTokenId\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"box_pro\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controlledTokenBuilder\",\"outputs\":[{\"internalType\":\"contractBuildToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"contractTokenControllerInterface\",\"name\":\"controller\",\"type\":\"address\"}],\"internalType\":\"structBuildToken.ControlledTokenConfig\",\"name\":\"_config\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prize_pool\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"mintKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600360cc6000509090556103e860cd600050909055348015620000275760006000fd5b50604051620097393803806200973983398181016040528101906200004d919062000466565b5b8260d760006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518061012001604052808773ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020015060ce60005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506101008201518160080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505b505050505050620005895662000588565b60008151905062000447816200054c565b5b92915050565b6000815190506200045f816200056a565b5b92915050565b60006000600060006000600060c08789031215620004845760006000fd5b60006200049489828a0162000436565b9650506020620004a789828a0162000436565b9550506040620004ba89828a0162000436565b9450506060620004cd89828a016200044e565b9350506080620004e089828a0162000436565b92505060a0620004f389828a0162000436565b9150505b9295509295509295565b60006200050e826200052b565b90505b919050565b6000620005238262000501565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b620005578162000501565b81141515620005665760006000fd5b5b50565b620005758162000516565b81141515620005845760006000fd5b5b50565b5b6191a080620005996000396000f3fe6080604052600436106102975760003560e01c806368c8ca0d1161015a578063a9059cbb116100c1578063dd62ed3e1161007a578063dd62ed3e14610abc578063e9292c4714610afa578063f05a020f14610b24578063f209089314610b4e578063f229c71c14610b7a578063fe49ba3214610bb8576102a0565b8063a9059cbb14610992578063af7a19ed146109d0578063c418161114610a11578063ca8585e014610a3b578063d4ecd04014610a68578063d505accf14610a92576102a0565b8063868a686d11610113578063868a686d1461086a57806387ab1642146108a857806395d89b41146108d25780639c724fa8146108fe5780639fcaf49a14610928578063a457c2d714610954576102a0565b806368c8ca0d1461070057806370a082311461073e5780637790f6ba1461077c57806379502c55146107ba5780637d54cf71146107ee5780637ecebe001461082c576102a0565b80633644e515116101fe5780634ef969ed116101b75780634ef969ed146105fc5780634f920f3e146106265780634fe5d7901461065057806350cfeddd1461067c5780635793cb6b146106a85780636067f59f146106d4576102a0565b80633644e515146104e857806339509351146105145780633b667870146105525780633eafcaf21461057c57806340b6ddbf146105a85780634521e181146105d2576102a0565b806318160ddd1161025057806318160ddd146103d157806323b872dd146103fd5780632f575b291461043b57806330157fb914610465578063313ce56714610492578063332ef5b9146104be576102a0565b806306fdde03146102a6578063076ac323146102d2578063095ea7b3146102fc57806312e7a53d1461033a57806312eb1fbf1461037957806314d9dcd0146103a5576102a0565b366102a0575b5b005b60006000fd5b3480156102b35760006000fd5b506102bc610be2565b6040516102c991906185d2565b60405180910390f35b3480156102df5760006000fd5b506102fa60048036038101906102f59190616997565b610c8c565b005b3480156103095760006000fd5b50610324600480360381019061031f919061684f565b611011565b604051610331919061844f565b60405180910390f35b3480156103475760006000fd5b50610362600480360381019061035d9190616b4b565b611040565b6040516103709291906183ca565b60405180910390f35b3480156103865760006000fd5b5061038f61151b565b60405161039c9190618af4565b60405180910390f35b3480156103b25760006000fd5b506103bb611520565b6040516103c891906185b6565b60405180910390f35b3480156103de5760006000fd5b506103e7611546565b6040516103f49190618af4565b60405180910390f35b34801561040a5760006000fd5b506104256004803603810190610420919061675a565b611558565b604051610432919061844f565b60405180910390f35b3480156104485760006000fd5b50610463600480360381019061045e9190616a97565b611657565b005b3480156104725760006000fd5b5061047b611961565b60405161048992919061841e565b60405180910390f35b34801561049f5760006000fd5b506104a8611a96565b6040516104b59190618bcf565b60405180910390f35b3480156104cb5760006000fd5b506104e660048036038101906104e19190616af3565b611ab2565b005b3480156104f55760006000fd5b506104fe611ca8565b60405161050b919061846b565b60405180910390f35b3480156105215760006000fd5b5061053c6004803603810190610537919061684f565b611cc2565b604051610549919061844f565b60405180910390f35b34801561055f5760006000fd5b5061057a60048036038101906105759190616a02565b611d95565b005b3480156105895760006000fd5b50610592611e75565b60405161059f9190618af4565b60405180910390f35b3480156105b55760006000fd5b506105d060048036038101906105cb9190616a97565b611e7a565b005b3480156105df5760006000fd5b506105fa60048036038101906105f59190616b8a565b612232565b005b3480156106095760006000fd5b50610624600480360381019061061f919061684f565b612563565b005b3480156106335760006000fd5b5061064e6004803603810190610649919061684f565b612644565b005b34801561065d5760006000fd5b50610666612779565b60405161067391906183fb565b60405180910390f35b3480156106895760006000fd5b506106926127dc565b60405161069f9190618af4565b60405180910390f35b3480156106b55760006000fd5b506106be6127e1565b6040516106cb9190618af4565b60405180910390f35b3480156106e15760006000fd5b506106ea6127ea565b6040516106f79190618af4565b60405180910390f35b34801561070d5760006000fd5b5061072860048036038101906107239190616a02565b6127ef565b60405161073591906185d2565b60405180910390f35b34801561074b5760006000fd5b50610766600480360381019061076191906166f0565b6128b5565b6040516107739190618af4565b60405180910390f35b3480156107895760006000fd5b506107a4600480360381019061079f9190616a02565b612909565b6040516107b191906183fb565b60405180910390f35b3480156107c75760006000fd5b506107d0612988565b6040516107e5999897969594939291906182a2565b60405180910390f35b3480156107fb5760006000fd5b5061081660048036038101906108119190616a02565b612ae7565b60405161082391906183fb565b60405180910390f35b3480156108395760006000fd5b50610854600480360381019061084f91906166f0565b612b66565b6040516108619190618af4565b60405180910390f35b3480156108775760006000fd5b50610892600480360381019061088d9190616a02565b612bc9565b60405161089f9190618a91565b60405180910390f35b3480156108b55760006000fd5b506108d060048036038101906108cb9190616a97565b612f7e565b005b3480156108df5760006000fd5b506108e86130bc565b6040516108f591906185d2565b60405180910390f35b34801561090b5760006000fd5b50610926600480360381019061092191906166f0565b613166565b005b3480156109355760006000fd5b5061093e613281565b60405161094b9190618af4565b60405180910390f35b3480156109615760006000fd5b5061097c6004803603810190610977919061684f565b613286565b604051610989919061844f565b60405180910390f35b34801561099f5760006000fd5b506109ba60048036038101906109b5919061684f565b613373565b6040516109c7919061844f565b60405180910390f35b3480156109dd5760006000fd5b506109f860048036038101906109f39190616a02565b6133a2565b604051610a0894939291906185f5565b60405180910390f35b348015610a1e5760006000fd5b50610a396004803603810190610a349190616a58565b61360d565b005b348015610a485760006000fd5b50610a516139b9565b604051610a5f929190618ba5565b60405180910390f35b348015610a755760006000fd5b50610a906004803603810190610a8b9190616a02565b613a02565b005b348015610a9f5760006000fd5b50610aba6004803603810190610ab591906167ac565b613aab565b005b348015610ac95760006000fd5b50610ae46004803603810190610adf919061671b565b613c8a565b604051610af19190618af4565b60405180910390f35b348015610b075760006000fd5b50610b226004803603810190610b1d9190616953565b613d1f565b005b348015610b315760006000fd5b50610b4c6004803603810190610b479190616b4b565b6145cd565b005b348015610b5b5760006000fd5b50610b64614a6f565b604051610b719190618ab4565b60405180910390f35b348015610b875760006000fd5b50610ba26004803603810190610b9d9190616a02565b614d9b565b604051610baf91906183fb565b60405180910390f35b348015610bc55760006000fd5b50610be06004803603810190610bdb919061671b565b614e1a565b005b606060366000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610c7d5780601f10610c5257610100808354040283529160200191610c7d565b820191906000526020600020905b815481529060010190602001808311610c6057829003601f168201915b50505050509050610c89565b90565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610d24576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610d1b906189cb565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610db55750600073ffffffffffffffffffffffffffffffffffffffff1660ce60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b8015610e155750600073ffffffffffffffffffffffffffffffffffffffff1660ce60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1515610e56576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610e4d906186f5565b60405180910390fd5b6000610f32848060000190610e6b9190618beb565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050858060200190610ebe9190618beb565b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050866040016020810190610f149190616bfb565b876060016020810190610f279190616928565b61506263ffffffff16565b90508060ce60005060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260ce60005060050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160ce60005060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b5b505050565b600061103161102461515e63ffffffff16565b848461516b63ffffffff16565b6001905061103a565b92915050565b60606000600060d960005080549050905060008111158061106057508385115b8061106a57508085115b1561107f576060808293509350505061151456505b6000849050818511156110925781905080505b601e86820311156110a657601e8601905080505b606086820367ffffffffffffffff811180156110c25760006000fd5b506040519080825280602002602001820160405280156110fc57816020015b6110e9615e75565b8152602001906001900390816110e15790505b509050600060008890505b838110156115005760d8600050600060d96000508381548110151561112857fe5b906000526020600020900160005b505481526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156111f05780601f106111c5576101008083540402835291602001916111f0565b820191906000526020600020905b8154815290600101906020018083116111d357829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156112a25780601f10611277576101008083540402835291602001916112a2565b820191906000526020600020905b81548152906001019060200180831161128557829003601f168201915b505050505081526020016003820160005080548060200260200160405190810160405280929190818152602001828054801561130057602002820191906000526020600020905b8160005054815260200190600101908083116112e9575b505050505081526020016004820160005080548060200260200160405190810160405280929190818152602001828054801561135e57602002820191906000526020600020905b816000505481526020019060010190808311611347575b50505050508152602001600582016000508054806020026020016040519081016040528092919081815260200182805480156113bc57602002820191906000526020600020905b8160005054815260200190600101908083116113a5575b50505050508152602001600682016000506040518060400160405290816000820160005080548060200260200160405190810160405280929190818152602001828054801561146057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311611416575b50505050508152602001600182016000508054806020026020016040519081016040528092919081815260200182805480156114be57602002820191906000526020600020905b8160005054815260200190600101908083116114a7575b5050505050815260200150508152602001505083838151811015156114df57fe5b602002602001018190525081806001019250505b8080600101915050611107565b508184955095505050505061151456505050505b9250929050565b608081565b60d760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b60006035600050549050611555565b90565b600061156b84848461534663ffffffff16565b6116478461157d61515e63ffffffff16565b61163c8560405180606001604052806028815260200161911e60289139603460005060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060006115ef61515e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546156069092919063ffffffff16565b61516b63ffffffff16565b60019050611650565b9392505050565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156116ef576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116e6906189cb565b60405180910390fd5b82600060d860005060008381526020019081526020016000206000509050600081600101600050541415151561175a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611751906187fd565b60405180910390fd5b6004848490501415156117a2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611799906187dc565b60405180910390fd5b600061182360d8600050600088815260200190815260200160002060005060030160005080548060200260200160405190810160405280929190818152602001828054801561181357602002820191906000526020600020905b8160005054815260200190600101908083116117fc575b505050505061566263ffffffff16565b9050600181038585905014151561186f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611866906188a2565b60405180910390fd5b848460d86000506000898152602001908152602001600020600050600501600050919061189d929190615ebb565b506000600090505b8585905081101561191a57620f424086868381811015156118c257fe5b905060200201351115151561190c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161190390618989565b60405180910390fd5b5b80806001019150506118a5565b507f7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a3886868660405161194e93929190618b10565b60405180910390a1505b505b505b505050565b60606000600060009050606060d96000508054905067ffffffffffffffff8111801561198d5760006000fd5b506040519080825280602002602001820160405280156119bc5781602001602082028036833780820191505090505b5090506000600090505b60d960005080549050811015611a82576060611a0860d9600050838154811015156119ed57fe5b906000526020600020900160005b505461290963ffffffff16565b90506000816004815181101515611a1b57fe5b60200260200101511115611a735760d960005082815481101515611a3b57fe5b906000526020600020900160005b50548385815181101515611a5957fe5b602002602001019090818152602001505083806001019450505b505b80806001019150506119c6565b508082935093505050611a925650505b9091565b6000603860009054906101000a900460ff169050611aaf565b90565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611b4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b41906189cb565b60405180910390fd5b81600060d8600050600083815260200190815260200160002060005090506000816001016000505414151515611bb5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bac906187fd565b60405180910390fd5b826020015151836000015151141515611c03576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bfa90618905565b60405180910390fd5b8260d860005060008681526020019081526020016000206000506006016000506000820151816000016000509080519060200190611c42929190615f0d565b506020820151816001016000509080519060200190611c62929190615f97565b509050507f13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d8484604051611c97929190618b74565b60405180910390a15b505b505b5050565b6000611cb86156bf63ffffffff16565b9050611cbf565b90565b6000611d86611cd561515e63ffffffff16565b84611d7b8560346000506000611cef61515e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461570990919063ffffffff16565b61516b63ffffffff16565b60019050611d8f565b92915050565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611e2d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e24906189cb565b60405180910390fd5b8060cd6000508190909055507f9ef7e80651fd9de0992e3b6fca0d2faee38eaf176212ee59054d2654958836f681604051611e689190618af4565b60405180910390a15b5b50565b606481565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611f12576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f09906189cb565b60405180910390fd5b82600060d8600050600083815260200190815260200160002060005090506000816001016000505414151515611f7d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611f74906187fd565b60405180910390fd5b600484849050141515611fc5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fbc906188a2565b60405180910390fd5b600061204660d8600050600088815260200190815260200160002060005060030160005080548060200260200160405190810160405280929190818152602001828054801561203657602002820191906000526020600020905b81600050548152602001906001019080831161201f575b505050505061566263ffffffff16565b9050600061209a868680806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f8201169050808301925050505050505061566263ffffffff16565b905060058214156120f15760018203811415156120ec576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120e3906186d4565b60405180910390fd5b612136565b8181141515612135576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161212c906186d4565b60405180910390fd5b5b60006000600090505b8787905081101561217357878782818110151561215857fe5b905060200201358201915081505b808060010191505061213f565b50620f4240811415156121bb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121b290618716565b60405180910390fd5b868660d860005060008b815260200190815260200160002060005060040160005091906121e9929190615ebb565b507f64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd4807688888860405161221d93929190618b10565b60405180910390a15050505b505b505b505050565b83600060d860005060008381526020019081526020016000206000509050600081600101600050541415151561229d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612294906187fd565b60405180910390fd5b84600581111515156122e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016122db906188c3565b60405180910390fd5b848480806020026020016040519081016040528093929190818152602001838360200280828437600081840152601f19601f820116905080830192505050505050506005815114151561236c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161236390618881565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156123dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016123d390618860565b60405180910390fd5b60ce60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634e6d4e41338a8a8a8a60ce60005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630926e497600036336040518463ffffffff1660e01b815260040161248b93929190618583565b602060405180830381600087803b1580156124a65760006000fd5b505af11580156124bb573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906124df9190616a2d565b6040518763ffffffff1660e01b815260040161250096959493929190618245565b602060405180830381600087803b15801561251b5760006000fd5b505af1158015612530573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906125549190616a2d565b505b5b505b50505b5050505050565b60ce60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156125fb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125f2906189cb565b60405180910390fd5b60018114156126235761261e82670de0b6b3a7640000600161576763ffffffff16565b61263e565b61263d82678ac7230489e80000600161576763ffffffff16565b5b5b5b5050565b60ce60005060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156126dc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016126d3906189cb565b60405180910390fd5b60ce60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166390596dd183836040518363ffffffff1660e01b815260040161273f929190618368565b600060405180830381600087803b15801561275a5760006000fd5b505af115801561276f573d600060003e3d6000fd5b505050505b5b5050565b606060d96000508054806020026020016040519081016040528092919081815260200182805480156127cd57602002820191906000526020600020905b8160005054815260200190600101908083116127b6575b505050505090506127d9565b90565b600481565b60cc6000505481565b600581565b606060d860005060008381526020019081526020016000206000506002016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156128a45780601f10612879576101008083540402835291602001916128a4565b820191906000526020600020905b81548152906001019060200180831161288757829003601f168201915b505050505090506128b0565b919050565b6000603360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050612904565b919050565b606060d8600050600083815260200190815260200160002060005060030160005080548060200260200160405190810160405280929190818152602001828054801561297757602002820191906000526020600020905b816000505481526020019060010190808311612960575b50505050509050612983565b919050565b60ce6000508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060080160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905089565b606060d86000506000838152602001908152602001600020600050600401600050805480602002602001604051908101604052809291908181526020018280548015612b5557602002820191906000526020600020905b816000505481526020019060010190808311612b3e575b50505050509050612b61565b919050565b6000612bbd609960005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050615806909063ffffffff16565b9050612bc4565b919050565b612bd1615e75565b60d860005060008381526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612c915780601f10612c6657610100808354040283529160200191612c91565b820191906000526020600020905b815481529060010190602001808311612c7457829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612d435780601f10612d1857610100808354040283529160200191612d43565b820191906000526020600020905b815481529060010190602001808311612d2657829003601f168201915b5050505050815260200160038201600050805480602002602001604051908101604052809291908181526020018280548015612da157602002820191906000526020600020905b816000505481526020019060010190808311612d8a575b5050505050815260200160048201600050805480602002602001604051908101604052809291908181526020018280548015612dff57602002820191906000526020600020905b816000505481526020019060010190808311612de8575b5050505050815260200160058201600050805480602002602001604051908101604052809291908181526020018280548015612e5d57602002820191906000526020600020905b816000505481526020019060010190808311612e46575b505050505081526020016006820160005060405180604001604052908160008201600050805480602002602001604051908101604052809291908181526020018280548015612f0157602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612eb7575b5050505050815260200160018201600050805480602002602001604051908101604052809291908181526020018280548015612f5f57602002820191906000526020600020905b816000505481526020019060010190808311612f48575b505050505081526020015050815260200150509050612f79565b919050565b82600060d8600050600083815260200190815260200160002060005090506000816001016000505414151515612fe9576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612fe0906187fd565b60405180910390fd5b60ce60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0c55e31338760d860005060008a815260200190815260200160002060005060030160005088886040518663ffffffff1660e01b8152600401613070959493929190618163565b600060405180830381600087803b15801561308b5760006000fd5b505af11580156130a0573d600060003e3d6000fd5b505050506130b38561581c63ffffffff16565b5b505b50505050565b606060376000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156131575780601f1061312c57610100808354040283529160200191613157565b820191906000526020600020905b81548152906001019060200180831161313a57829003601f168201915b50505050509050613163565b90565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156131fe576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131f5906189cb565b60405180910390fd5b8060ce60005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a81604051613274919061804b565b60405180910390a15b5b50565b600881565b600061336461329961515e63ffffffff16565b846133598560405180606001604052806025815260200161914660259139603460005060006132cc61515e63ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546156069092919063ffffffff16565b61516b63ffffffff16565b6001905061336d565b92915050565b600061339361338661515e63ffffffff16565b848461534663ffffffff16565b6001905061339c565b92915050565b60d8600050602052806000526040600020600091509050806000016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156134545780601f1061342957610100808354040283529160200191613454565b820191906000526020600020905b81548152906001019060200180831161343757829003601f168201915b505050505090806001016000505490806002016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156134fe5780601f106134d3576101008083540402835291602001916134fe565b820191906000526020600020905b8154815290600101906020018083116134e157829003601f168201915b505050505090806006016000506040518060400160405290816000820160005080548060200260200160405190810160405280929190818152602001828054801561359e57602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311613554575b50505050508152602001600182016000508054806020026020016040519081016040528092919081815260200182805480156135fc57602002820191906000526020600020905b8160005054815260200190600101908083116135e5575b505050505081526020015050905084565b81600181148061361d5750600a81145b151561365e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161365590618a4f565b60405180910390fd5b600060cd600050546064670de0b6b3a7640000860281151561367c57fe5b04029050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156136f2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136e990618779565b60405180910390fd5b600060ce60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161375c929190618091565b60206040518083038186803b1580156137755760006000fd5b505afa15801561378a573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906137ae9190616a2d565b90508281101515156137f5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137ec90618947565b60405180910390fd5b61382f60ce60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633308661598663ffffffff16565b61389f60ce60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660ce60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166009600a8781151561389257fe5b0402615ac363ffffffff16565b61390f60ce60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660ce60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16600560648781151561390257fe5b0402615ac363ffffffff16565b61395760ce60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1686600560648781151561394a57fe5b0402615ac363ffffffff16565b61397233670de0b6b3a764000088028861576763ffffffff16565b7f7abdf250d54b73efb224a0fca882e6b50f7a946e7262834728106d5ec4cdf9ef338787866040516139a7949392919061811d565b60405180910390a15050505b5b505050565b6000600060cd600050546064670de0b6b3a76400008115156139d757fe5b040260cd600050546064678ac7230489e800008115156139f357fe5b0402915091506139fe565b9091565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613a9a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613a91906189cb565b60405180910390fd5b8060cc6000508190909055505b5b50565b834211151515613af0576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ae7906187bb565b60405180910390fd5b6000609a60005054888888613b50609960005060008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050615806909063ffffffff16565b89604051602001613b6696959493929190618487565b6040516020818303038152906040528051906020012090506000613b8f82615bfd63ffffffff16565b90506000613ba582878787615c4163ffffffff16565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515613c17576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613c0e90618926565b60405180910390fd5b613c6c609960005060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050615dd9909063ffffffff16565b613c7d8a8a8a61516b63ffffffff16565b5050505b50505050505050565b6000603460005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050613d19565b92915050565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613db7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613dae906189cb565b60405180910390fd5b806000816000015151905060048110151515613e08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613dff90618a0d565b60405180910390fd5b600060d8600050600084602001518152602001908152602001600020600050905060008160010160005054141515613e75576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e6c90618758565b60405180910390fd5b6000836040015151905060088110151515613ec5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ebc90618a70565b60405180910390fd5b60808111151515613f0b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f0290618968565b60405180910390fd5b8360c0015160200151518460c001516000015151141515613f61576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613f5890618905565b60405180910390fd5b6000600090505b846060015151811015613fe1576000856060015182815181101515613f8957fe5b6020026020010151111515613fd3576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613fca90618989565b60405180910390fd5b5b8080600101915050613f68565b5060006000600090505b85608001515181101561408457600086608001518281518110151561400c57fe5b6020026020010151111515614056576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161404d90618989565b60405180910390fd5b85608001518181518110151561406857fe5b60200260200101518201915081505b8080600101915050613feb565b50620f4240811415156140cc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016140c390618716565b60405180910390fd5b6002856060015151101580156140e85750600585606001515111155b1515614129576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161412090618716565b60405180910390fd5b60058560600151511415614197576004856080015151148015614151575060048560a0015151145b1515614192576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141899061881e565b60405180910390fd5b6141fe565b8460600151518560800151511480156141bc57506001856060015151038560a0015151145b15156141fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016141f49061881e565b60405180910390fd5b5b60008660600151516005039050600081111561442e576060600567ffffffffffffffff8111801561422f5760006000fd5b5060405190808252806020026020018201604052801561425e5781602001602082028036833780820191505090505b5090506060600467ffffffffffffffff8111801561427c5760006000fd5b506040519080825280602002602001820160405280156142ab5781602001602082028036833780820191505090505b5090506060600467ffffffffffffffff811180156142c95760006000fd5b506040519080825280602002602001820160405280156142f85781602001602082028036833780820191505090505b5090506000600090505b8a6060015151811015614356578a606001518181518110151561432157fe5b6020026020010151848281518110151561433757fe5b60200260200101909081815260200150505b8080600101915050614302565b506000600090505b8a60800151518110156143b2578a608001518181518110151561437d57fe5b6020026020010151838281518110151561439357fe5b60200260200101909081815260200150505b808060010191505061435e565b506000600090505b8a60a001515181101561440e578a60a00151818151811015156143d957fe5b602002602001015182828151811015156143ef57fe5b60200260200101909081815260200150505b80806001019150506143ba565b50808a60a00181905250818a60800181905250828a606001819052505050505b8660d8600050600089602001518152602001908152602001600020600050600082015181600001600050908051906020019061446b929190615fe9565b50602082015181600101600050909055604082015181600201600050908051906020019061449a929190615fe9565b5060608201518160030160005090805190602001906144ba929190615f97565b5060808201518160040160005090805190602001906144da929190615f97565b5060a08201518160050160005090805190602001906144fa929190615f97565b5060c0820151816006016000506000820151816000016000509080519060200190614526929190615f0d565b506020820151816001016000509080519060200190614546929190615f97565b50505090505060d96000508760200151908060018154018082558091505060019003906000526020600020900160005b909190919091509090557fc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb876020015188600001516040516145b9929190618b43565b60405180910390a1505b505050505b505b50565b81600060d8600050600083815260200190815260200160002060005090506000816001016000505414151515614638576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161462f906187fd565b60405180910390fd5b82600181148061464c575060cc6000505481145b151561468d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614684906186b3565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156146fd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146f490618860565b60405180910390fd5b600060ce60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401614767929190618091565b60206040518083038186803b1580156147805760006000fd5b505afa158015614795573d600060003e3d6000fd5b505050506040513d601f19601f820116820180604052508101906147b99190616a2d565b9050670de0b6b3a76400008602811015151561480a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161480190618947565b60405180910390fd5b60ce60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd26c818338860018b60ce60005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630926e497600036336040518463ffffffff1660e01b81526004016148b993929190618583565b602060405180830381600087803b1580156148d45760006000fd5b505af11580156148e9573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061490d9190616a2d565b6040518663ffffffff1660e01b815260040161492d9594939291906181b9565b600060405180830381600087803b1580156149485760006000fd5b505af115801561495d573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f82011682018060405250810190614986919061688e565b5060ce60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166390596dd133670de0b6b3a764000089026040518363ffffffff1660e01b81526004016149f49291906180f3565b600060405180830381600087803b158015614a0f5760006000fd5b505af1158015614a24573d600060003e3d6000fd5b505050507fd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b023040338888604051614a5b9392919061820d565b60405180910390a150505b5b50505b505050565b614a7761606e565b60ce600050604051806101200160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016008820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200150509050614d98565b90565b606060d86000506000838152602001908152602001600020600050600501600050805480602002602001604051908101604052809291908181526020018280548015614e0957602002820191906000526020600020905b816000505481526020019060010190808311614df2575b50505050509050614e15565b919050565b60ce60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515614eb2576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614ea9906189cb565b60405180910390fd5b8160ce60005060080160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060ce60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fe49ba3283836040518363ffffffff1660e01b8152600401614f5c929190618067565b600060405180830381600087803b158015614f775760006000fd5b505af1158015614f8c573d600060003e3d6000fd5b5050505060ce60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166302eca6c0836040518263ffffffff1660e01b8152600401614ff1919061804b565b600060405180830381600087803b15801561500c5760006000fd5b505af1158015615021573d600060003e3d6000fd5b505050507f484ee7be515f50ca021b1d2f6bda39adb20563e2cf9041a996cf3d125346271b82604051615054919061804b565b60405180910390a15b5b5050565b600060d760009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c0cd38d60405180608001604052808881526020018781526020018660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001506040518263ffffffff1660e01b81526004016150fb9190618ad1565b602060405180830381600087803b1580156151165760006000fd5b505af115801561512b573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061514f91906168fd565b9050615156565b949350505050565b6000339050615168565b90565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156151dd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016151d4906189ec565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561524f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161524690618737565b60405180910390fd5b80603460005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516153389190618af4565b60405180910390a35b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff16141515156153b8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016153af906189aa565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561542a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161542190618671565b60405180910390fd5b61543b838383615df463ffffffff16565b6154ad816040518060600160405280602681526020016190f860269139603360005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546156069092919063ffffffff16565b603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555061555081603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461570990919063ffffffff16565b603360005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516155f89190618af4565b60405180910390a35b505050565b60008383111582901515615650576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161564791906185d2565b60405180910390fd5b50828403905061565b565b9392505050565b60006000600090506000600090505b83518110156156af576000848281518110151561568a57fe5b602002602001015111156156a15781806001019250505b5b8080600101915050615671565b50809150506156ba56505b919050565b60006156ff6040516156d090618035565b60405180910390206156e6615dfa63ffffffff16565b6156f4615e0c63ffffffff16565b615e1e63ffffffff16565b9050615706565b90565b600060008284019050838110151515615757576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161574e9061879a565b60405180910390fd5b8091505061576156505b92915050565b60ce60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a8d885f8484846040518463ffffffff1660e01b81526004016157cc93929190618392565b600060405180830381600087803b1580156157e75760006000fd5b505af11580156157fc573d600060003e3d6000fd5b505050505b505050565b600081600001600050549050615817565b919050565b600060d8600050600083815260200190815260200160002060005090506000816006016000506000016000508054905090506000600090505b8181101561597f576000836006016000506000016000508281548110151561587957fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600084600601600050600101600050838154811015156158c457fe5b906000526020600020900160005b5054905060ce60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166387227ab13384846040518463ffffffff1660e01b815260040161593b939291906180bb565b600060405180830381600087803b1580156159565760006000fd5b505af115801561596b573d600060003e3d6000fd5b5050505050505b8080600101915050615855565b5050505b50565b600060608573ffffffffffffffffffffffffffffffffffffffff166323b872dd8686866040516024016159bb93929190618330565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051615a099190617fe5565b6000604051808303816000865af19150503d8060008114615a46576040519150601f19603f3d011682016040523d82523d6000602084013e615a4b565b606091505b5091509150818015615a795750600081511480615a78575080806020019051810190615a7791906168d2565b5b5b1515615aba576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615ab190618a2e565b60405180910390fd5b50505b50505050565b600060608473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401615af6929190618368565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051615b449190617fe5565b6000604051808303816000865af19150503d8060008114615b81576040519150601f19603f3d011682016040523d82523d6000602084013e615b86565b606091505b5091509150818015615bb45750600081511480615bb3575080806020019051810190615bb291906168d2565b5b5b1515615bf5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615bec90618692565b60405180910390fd5b50505b505050565b6000615c0d6156bf63ffffffff16565b82604051602001615c1f929190617ffd565b604051602081830303815290604052805190602001209050615c3c565b919050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11151515615cab576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615ca29061883f565b60405180910390fd5b601b8460ff161480615cc05750601c8460ff16145b1515615d01576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615cf8906188e4565b60405180910390fd5b600060018686868660405160008152602001604052604051615d26949392919061853d565b6020604051602081039080840390855afa158015615d49573d600060003e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515615dc7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401615dbe90618650565b60405180910390fd5b80915050615dd156505b949350505050565b60018160000160008282825054019250508190909055505b50565b5b505050565b60006065600050549050615e09565b90565b60006066600050549050615e1b565b90565b6000838383615e31615e6a63ffffffff16565b30604051602001615e469594939291906184e9565b604051602081830303815290604052805190602001209050615e63565b9392505050565b600030504690505b90565b6040518060e00160405280606081526020016000815260200160608152602001606081526020016060815260200160608152602001615eb2616183565b81526020015090565b828054828255906000526020600020908101928215615efc579160200282015b82811115615efb5782358260005090905591602001919060010190615edb565b5b509050615f0991906161a0565b5090565b828054828255906000526020600020908101928215615f86579160200282015b82811115615f855782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190615f2d565b5b509050615f9391906161cb565b5090565b828054828255906000526020600020908101928215615fd8579160200282015b82811115615fd75782518260005090905591602001919060010190615fb7565b5b509050615fe591906161a0565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061602a57805160ff191683800117855561605d565b8280016001018555821561605d579182015b8281111561605c578251826000509090559160200191906001019061603c565b5b50905061606a91906161a0565b5090565b604051806101200160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020015090565b604051806040016040528060608152602001606081526020015090565b6161c891906161aa565b808211156161c457600081815060009055506001016161aa565b5090565b90565b61620f91906161d5565b8082111561620b57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016161d5565b5090565b90566190f6565b60008135905061622581619039565b5b92915050565b600082601f83011215156162405760006000fd5b813561625361624e82618c78565b618c49565b915081818352602084019350602081019050838560208402820111156162795760006000fd5b60005b838110156162aa578161628f8882616216565b8452602084019350602083019250505b60018101905061627c565b505050505b92915050565b6000600083601f84011215156162cb5760006000fd5b8235905067ffffffffffffffff8111156162e55760006000fd5b6020830191508360208202830111156162fe5760006000fd5b5b9250929050565b600082601f830112151561631a5760006000fd5b813561632d61632882618ca2565b618c49565b915081818352602084019350602081019050838560208402820111156163535760006000fd5b60005b83811015616384578161636988826166ae565b8452602084019350602083019250505b600181019050616356565b505050505b92915050565b600082601f83011215156163a35760006000fd5b81516163b66163b182618ca2565b618c49565b915081818352602084019350602081019050838560208402820111156163dc5760006000fd5b60005b8381101561640d57816163f288826166c4565b8452602084019350602083019250505b6001810190506163df565b505050505b92915050565b60008151905061642781619054565b5b92915050565b60008135905061643d8161906f565b5b92915050565b6000815190506164538161908a565b5b92915050565b600081359050616469816190a5565b5b92915050565b600082601f83011215156164845760006000fd5b813561649761649282618ccc565b618c49565b915080825260208301602083018583830111156164b45760006000fd5b6164bf838284618f9a565b5050505b92915050565b600060e082840312156164dc5760006000fd5b6164e660e0618c49565b9050600082013567ffffffffffffffff8111156165035760006000fd5b61650f84828501616470565b6000830152506020616523848285016166ae565b602083015250604082013567ffffffffffffffff8111156165445760006000fd5b61655084828501616470565b604083015250606082013567ffffffffffffffff8111156165715760006000fd5b61657d84828501616306565b606083015250608082013567ffffffffffffffff81111561659e5760006000fd5b6165aa84828501616306565b60808301525060a082013567ffffffffffffffff8111156165cb5760006000fd5b6165d784828501616306565b60a08301525060c082013567ffffffffffffffff8111156165f85760006000fd5b6166048482850161662e565b60c0830152505b92915050565b6000608082840312156166245760006000fd5b8190505b92915050565b6000604082840312156166415760006000fd5b61664b6040618c49565b9050600082013567ffffffffffffffff8111156166685760006000fd5b6166748482850161622c565b600083015250602082013567ffffffffffffffff8111156166955760006000fd5b6166a184828501616306565b6020830152505b92915050565b6000813590506166bd816190c0565b5b92915050565b6000815190506166d3816190c0565b5b92915050565b6000813590506166e9816190db565b5b92915050565b6000602082840312156167035760006000fd5b600061671184828501616216565b9150505b92915050565b60006000604083850312156167305760006000fd5b600061673e85828601616216565b925050602061674f85828601616216565b9150505b9250929050565b600060006000606084860312156167715760006000fd5b600061677f86828701616216565b935050602061679086828701616216565b92505060406167a1868287016166ae565b9150505b9250925092565b600060006000600060006000600060e0888a0312156167cb5760006000fd5b60006167d98a828b01616216565b97505060206167ea8a828b01616216565b96505060406167fb8a828b016166ae565b955050606061680c8a828b016166ae565b945050608061681d8a828b016166da565b93505060a061682e8a828b0161642e565b92505060c061683f8a828b0161642e565b9150505b92959891949750929550565b60006000604083850312156168645760006000fd5b600061687285828601616216565b9250506020616883858286016166ae565b9150505b9250929050565b6000602082840312156168a15760006000fd5b600082015167ffffffffffffffff8111156168bc5760006000fd5b6168c88482850161638f565b9150505b92915050565b6000602082840312156168e55760006000fd5b60006168f384828501616418565b9150505b92915050565b6000602082840312156169105760006000fd5b600061691e84828501616444565b9150505b92915050565b60006020828403121561693b5760006000fd5b60006169498482850161645a565b9150505b92915050565b6000602082840312156169665760006000fd5b600082013567ffffffffffffffff8111156169815760006000fd5b61698d848285016164c9565b9150505b92915050565b600060006000606084860312156169ae5760006000fd5b600084013567ffffffffffffffff8111156169c95760006000fd5b6169d586828701616611565b93505060206169e686828701616216565b92505060406169f786828701616216565b9150505b9250925092565b600060208284031215616a155760006000fd5b6000616a23848285016166ae565b9150505b92915050565b600060208284031215616a405760006000fd5b6000616a4e848285016166c4565b9150505b92915050565b6000600060408385031215616a6d5760006000fd5b6000616a7b858286016166ae565b9250506020616a8c85828601616216565b9150505b9250929050565b60006000600060408486031215616aae5760006000fd5b6000616abc868287016166ae565b935050602084013567ffffffffffffffff811115616ada5760006000fd5b616ae6868287016162b5565b92509250505b9250925092565b6000600060408385031215616b085760006000fd5b6000616b16858286016166ae565b925050602083013567ffffffffffffffff811115616b345760006000fd5b616b408582860161662e565b9150505b9250929050565b6000600060408385031215616b605760006000fd5b6000616b6e858286016166ae565b9250506020616b7f858286016166ae565b9150505b9250929050565b600060006000600060608587031215616ba35760006000fd5b6000616bb1878288016166ae565b9450506020616bc2878288016166ae565b935050604085013567ffffffffffffffff811115616be05760006000fd5b616bec878288016162b5565b92509250505b92959194509250565b600060208284031215616c0e5760006000fd5b6000616c1c848285016166da565b9150505b92915050565b6000616c328383616c9d565b6020830190505b92915050565b6000616c4b8383617c78565b90505b92915050565b6000616c608383617fa5565b6020830190505b92915050565b616c7681618f02565b82525b5050565b616c8681618e77565b82525b5050565b616c9681618e77565b82525b5050565b616ca681618e64565b82525b5050565b616cb681618e64565b82525b5050565b6000616cc882618d43565b616cd28185618dc3565b9350616cdd83618cfa565b8060005b83811015616d0f578151616cf58882616c26565b9750616d0083618d8b565b9250505b600181019050616ce1565b508593505050505b92915050565b6000616d2882618d4f565b616d328185618dd5565b935083602082028501616d4485618d0b565b8060005b85811015616d815784840389528151616d618582616c3f565b9450616d6c83618d99565b925060208a019950505b600181019050616d48565b5082975087955050505050505b92915050565b6000616da08385618df9565b93507f07ffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff831115616dd05760006000fd5b602083029250616de1838584618f9a565b82840190505b9392505050565b6000616df982618d5b565b616e038185618de7565b9350616e0e83618d1c565b8060005b83811015616e40578151616e268882616c54565b9750616e3183618da7565b9250505b600181019050616e12565b508593505050505b92915050565b6000616e5982618d5b565b616e638185618df9565b9350616e6e83618d1c565b8060005b83811015616ea0578151616e868882616c54565b9750616e9183618da7565b9250505b600181019050616e72565b508593505050505b92915050565b6000616eb982618d67565b616ec38185618df9565b9350616ece83618d2d565b8060005b83811015616f0757616ee382619005565b616eed8882616c54565b9750616ef883618db5565b9250505b600181019050616ed2565b508593505050505b92915050565b616f1e81618e8a565b82525b5050565b616f2e81618e97565b82525b5050565b616f46616f4182618e97565b618ffa565b82525b5050565b6000616f598385618e0b565b9350616f66838584618f9a565b616f6f83619019565b840190505b9392505050565b6000616f8682618d73565b616f908185618e1d565b9350616fa0818560208601618faa565b8084019150505b92915050565b616fb681618f15565b82525b5050565b616fc681618f3b565b82525b5050565b616fd681618f61565b82525b5050565b6000616fe882618d7f565b616ff28185618e29565b9350617002818560208601618faa565b61700b81619019565b84019150505b92915050565b600061702282618d7f565b61702c8185618e3b565b935061703c818560208601618faa565b61704581619019565b84019150505b92915050565b600061705e601883618e3b565b91507f45434453413a20696e76616c6964207369676e6174757265000000000000000060008301526020820190505b919050565b600061709f602383618e3b565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f657373000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617106601f83618e3b565b91507f5472616e7366657248656c7065723a205452414e534645525f4641494c45440060008301526020820190505b919050565b6000617147602b83618e3b565b91507f426c696e64426f78204572723a4f6e6c7920647261772074686520737065636960008301527f66696564206e756d62657200000000000000000000000000000000000000000060208301526040820190505b919050565b60006171ae601583618e3b565b91507f426c696e64626f78204572723a6472617720657272000000000000000000000060008301526020820190505b919050565b60006171ef602683618e3b565b91507f426c696e64426f78204572723a43616e206e6f742062652072652d696e69746960008301527f616c697a6564000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617256601883618e3b565b91507f426c696e64426f78204572723a7265717565737420657272000000000000000060008301526020820190505b919050565b6000617297602283618e3b565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f737300000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006172fe602083618e3b565b91507f426c696e64426f78204572723a20426f7820616c72656164792065786973747360008301526020820190505b919050565b600061733f602c83618e3b565b91507f426c696e64426f78204572723a696e76697465722063616e6e6f74206571756160008301527f6c2061646472657373283029000000000000000000000000000000000000000060208301526040820190505b919050565b60006173a6600283618e4d565b91507f190100000000000000000000000000000000000000000000000000000000000060008301526002820190505b919050565b60006173e7601b83618e3b565b91507f536166654d6174683a206164646974696f6e206f766572666c6f77000000000060008301526020820190505b919050565b6000617428601d83618e3b565b91507f45524332305065726d69743a206578706972656420646561646c696e6500000060008301526020820190505b919050565b6000617469601783618e3b565b91507f426c696e64626f78204572723a2072657365742065727200000000000000000060008301526020820190505b919050565b60006174aa601e83618e3b565b91507f426c696e64426f78204572723a20736572696573206e6f7420666f756e64000060008301526020820190505b919050565b60006174eb601983618e3b565b91507f426c696e64426f78204572723a2052657175657374206572720000000000000060008301526020820190505b919050565b600061752c602283618e3b565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617593601b83618e3b565b91507f426c696e64426f78204572723a72657175657374206661696c6564000000000060008301526020820190505b919050565b60006175d4602983618e3b565b91507f426c696e64426f78204572723a204f6e6c7920726563656976652035206e667460008301527f20746f6b656e206964000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061763b601683618e3b565b91507f426c696e64626f78204572723a7265736574206572720000000000000000000060008301526020820190505b919050565b600061767c602283618e3b565b91507f426c696e64426f7820204572723a477261646520646f6573206e6f742065786960008301527f737400000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006176e3602283618e3b565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061774a603283618e3b565b91507f426c696e64426f78204572723a2072657761726420746f6b656e206e6f74206560008301527f7175616c2072657761726420616d6f756e74000000000000000000000000000060208301526040820190505b919050565b60006177b1605283618e4d565b91507f454950373132446f6d61696e28737472696e67206e616d652c737472696e672060008301527f76657273696f6e2c75696e7432353620636861696e49642c616464726573732060208301527f766572696679696e67436f6e747261637429000000000000000000000000000060408301526052820190505b919050565b600061783e601e83618e3b565b91507f45524332305065726d69743a20696e76616c6964207369676e6174757265000060008301526020820190505b919050565b600061787f602983618e3b565b91507f426c696e64426f78204572723a616d6f756e742063616e6e6f74207468616e2060008301527f616c6c6f77616e6365000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006178e6604783618e3b565b91507f426c696e64426f78204572723a20496d6167654c696e6b206c656e677468206d60008301527f75737420626520736d616c6c207468616e204d41585f494d4147455f4c494e4b60208301527f5f4c454e4754480000000000000000000000000000000000000000000000000060408301526060820190505b919050565b6000617973601983618e3b565b91507f426c696e64426f78204572723a2072657175657374206572720000000000000060008301526020820190505b919050565b60006179b4602583618e3b565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f647265737300000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617a1b601a83618e3b565b91507f426c696e64426f78204572723a20556e617574686f72757a656400000000000060008301526020820190505b919050565b6000617a5c602483618e3b565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f726573730000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617ac3603983618e3b565b91507f426c696e64426f78204572723a206e616d65206c656e677468206d757374206260008301527f65206c657373207468616e204d494e5f4e414d455f4e414d450000000000000060208301526040820190505b919050565b6000617b2a602483618e3b565b91507f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f464160008301527f494c45440000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000617b91603583618e3b565b91507f426c696e64426f78204572723a64726177206e756d6265722063616e206f6e6c60008301527f7920626520657175616c20746f2031206f72203130000000000000000000000060208301526040820190505b919050565b6000617bf8604683618e3b565b91507f426c696e64426f78204572723a20496d6167654c696e6b206c656e677468206d60008301527f757374206265206c657373207468616e204d494e5f494d4147455f4c494e4b5f60208301527f4c454e475448000000000000000000000000000000000000000000000000000060408301526060820190505b919050565b600060e0830160008301518482036000860152617c958282616fdd565b9150506020830151617caa6020860182617fa5565b5060408301518482036040860152617cc28282616fdd565b91505060608301518482036060860152617cdc8282616dee565b91505060808301518482036080860152617cf68282616dee565b91505060a083015184820360a0860152617d108282616dee565b91505060c083015184820360c0860152617d2a8282617f1b565b915050809150505b92915050565b600060e0830160008301518482036000860152617d558282616fdd565b9150506020830151617d6a6020860182617fa5565b5060408301518482036040860152617d828282616fdd565b91505060608301518482036060860152617d9c8282616dee565b91505060808301518482036080860152617db68282616dee565b91505060a083015184820360a0860152617dd08282616dee565b91505060c083015184820360c0860152617dea8282617f1b565b915050809150505b92915050565b61012082016000820151617e0f6000850182616c9d565b506020820151617e226020850182616c9d565b506040820151617e356040850182616c9d565b506060820151617e486060850182616c9d565b506080820151617e5b6080850182616c7d565b5060a0820151617e6e60a0850182616c9d565b5060c0820151617e8160c0850182616c9d565b5060e0820151617e9460e0850182616c9d565b50610100820151617ea9610100850182616c9d565b50505b5050565b60006080830160008301518482036000860152617ecd8282616fdd565b91505060208301518482036020860152617ee78282616fdd565b9150506040830151617efc6040860182617fc5565b506060830151617f0f6060860182616fbd565b50809150505b92915050565b60006040830160008301518482036000860152617f388282616cbd565b91505060208301518482036020860152617f528282616dee565b915050809150505b92915050565b60006040830160008301518482036000860152617f7d8282616cbd565b91505060208301518482036020860152617f978282616dee565b915050809150505b92915050565b617fae81618ee9565b82525b5050565b617fbe81618ee9565b82525b5050565b617fce81618ef4565b82525b5050565b617fde81618ef4565b82525b5050565b6000617ff18284616f7b565b91508190505b92915050565b600061800882617399565b91506180148285616f35565b6020820191506180248284616f35565b6020820191508190505b9392505050565b6000618040826177a4565b91508190505b919050565b60006020820190506180606000830184616cad565b5b92915050565b600060408201905061807c6000830185616c8d565b6180896020830184616cad565b5b9392505050565b60006040820190506180a66000830185616c6d565b6180b36020830184616c6d565b5b9392505050565b60006060820190506180d06000830186616c8d565b6180dd6020830185616cad565b6180ea6040830184617fb5565b5b949350505050565b60006040820190506181086000830185616c6d565b6181156020830184617fb5565b5b9392505050565b60006080820190506181326000830187616c6d565b61813f6020830186617fb5565b61814c6040830185616cad565b6181596060830184617fb5565b5b95945050505050565b60006080820190506181786000830188616c6d565b6181856020830187617fb5565b81810360408301526181978186616eae565b905081810360608301526181ac818486616d94565b90505b9695505050505050565b600060a0820190506181ce6000830188616c6d565b6181db6020830187617fb5565b6181e86040830186616fcd565b6181f56060830185617fb5565b6182026080830184617fb5565b5b9695505050505050565b60006060820190506182226000830186616c6d565b61822f6020830185617fb5565b61823c6040830184617fb5565b5b949350505050565b600060a08201905061825a6000830189616c6d565b6182676020830188617fb5565b6182746040830187617fb5565b8181036060830152618287818587616d94565b90506182966080830184617fb5565b5b979650505050505050565b6000610120820190506182b8600083018c616cad565b6182c5602083018b616cad565b6182d2604083018a616cad565b6182df6060830189616cad565b6182ec6080830188616c8d565b6182f960a0830187616cad565b61830660c0830186616cad565b61831360e0830185616cad565b618321610100830184616cad565b5b9a9950505050505050505050565b60006060820190506183456000830186616cad565b6183526020830185616cad565b61835f6040830184617fb5565b5b949350505050565b600060408201905061837d6000830185616cad565b61838a6020830184617fb5565b5b9392505050565b60006060820190506183a76000830186616cad565b6183b46020830185617fb5565b6183c16040830184617fb5565b5b949350505050565b600060408201905081810360008301526183e48185616d1d565b90506183f36020830184617fb5565b5b9392505050565b600060208201905081810360008301526184158184616e4e565b90505b92915050565b600060408201905081810360008301526184388185616e4e565b90506184476020830184617fb5565b5b9392505050565b60006020820190506184646000830184616f15565b5b92915050565b60006020820190506184806000830184616f25565b5b92915050565b600060c08201905061849c6000830189616f25565b6184a96020830188616cad565b6184b66040830187616cad565b6184c36060830186617fb5565b6184d06080830185617fb5565b6184dd60a0830184617fb5565b5b979650505050505050565b600060a0820190506184fe6000830188616f25565b61850b6020830187616f25565b6185186040830186616f25565b6185256060830185617fb5565b6185326080830184616cad565b5b9695505050505050565b60006080820190506185526000830187616f25565b61855f6020830186617fd5565b61856c6040830185616f25565b6185796060830184616f25565b5b95945050505050565b6000604082019050818103600083015261859e818587616f4d565b90506185ad6020830184616c6d565b5b949350505050565b60006020820190506185cb6000830184616fad565b5b92915050565b600060208201905081810360008301526185ec8184617017565b90505b92915050565b6000608082019050818103600083015261860f8187617017565b905061861e6020830186617fb5565b81810360408301526186308185617017565b905081810360608301526186448184617f60565b90505b95945050505050565b6000602082019050818103600083015261866981617051565b90505b919050565b6000602082019050818103600083015261868a81617092565b90505b919050565b600060208201905081810360008301526186ab816170f9565b90505b919050565b600060208201905081810360008301526186cc8161713a565b90505b919050565b600060208201905081810360008301526186ed816171a1565b90505b919050565b6000602082019050818103600083015261870e816171e2565b90505b919050565b6000602082019050818103600083015261872f81617249565b90505b919050565b600060208201905081810360008301526187508161728a565b90505b919050565b60006020820190508181036000830152618771816172f1565b90505b919050565b6000602082019050818103600083015261879281617332565b90505b919050565b600060208201905081810360008301526187b3816173da565b90505b919050565b600060208201905081810360008301526187d48161741b565b90505b919050565b600060208201905081810360008301526187f58161745c565b90505b919050565b600060208201905081810360008301526188168161749d565b90505b919050565b60006020820190508181036000830152618837816174de565b90505b919050565b600060208201905081810360008301526188588161751f565b90505b919050565b6000602082019050818103600083015261887981617586565b90505b919050565b6000602082019050818103600083015261889a816175c7565b90505b919050565b600060208201905081810360008301526188bb8161762e565b90505b919050565b600060208201905081810360008301526188dc8161766f565b90505b919050565b600060208201905081810360008301526188fd816176d6565b90505b919050565b6000602082019050818103600083015261891e8161773d565b90505b919050565b6000602082019050818103600083015261893f81617831565b90505b919050565b6000602082019050818103600083015261896081617872565b90505b919050565b60006020820190508181036000830152618981816178d9565b90505b919050565b600060208201905081810360008301526189a281617966565b90505b919050565b600060208201905081810360008301526189c3816179a7565b90505b919050565b600060208201905081810360008301526189e481617a0e565b90505b919050565b60006020820190508181036000830152618a0581617a4f565b90505b919050565b60006020820190508181036000830152618a2681617ab6565b90505b919050565b60006020820190508181036000830152618a4781617b1d565b90505b919050565b60006020820190508181036000830152618a6881617b84565b90505b919050565b60006020820190508181036000830152618a8981617beb565b90505b919050565b60006020820190508181036000830152618aab8184617d38565b90505b92915050565b600061012082019050618aca6000830184617df8565b5b92915050565b60006020820190508181036000830152618aeb8184617eb0565b90505b92915050565b6000602082019050618b096000830184617fb5565b5b92915050565b6000604082019050618b256000830186617fb5565b8181036020830152618b38818486616d94565b90505b949350505050565b6000604082019050618b586000830185617fb5565b8181036020830152618b6a8184617017565b90505b9392505050565b6000604082019050618b896000830185617fb5565b8181036020830152618b9b8184617f60565b90505b9392505050565b6000604082019050618bba6000830185617fb5565b618bc76020830184617fb5565b5b9392505050565b6000602082019050618be46000830184617fd5565b5b92915050565b60006000833560016020038436030381121515618c085760006000fd5b80840192508235915067ffffffffffffffff821115618c275760006000fd5b602083019250600182023603831315618c405760006000fd5b505b9250929050565b6000604051905081810181811067ffffffffffffffff82111715618c6d5760006000fd5b80604052505b919050565b600067ffffffffffffffff821115618c905760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff821115618cba5760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff821115618ce45760006000fd5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b600081905081600052602060002090505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b6000815490505b919050565b6000815190505b919050565b6000815190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006001820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008190505b919050565b6000618e6f82618ec8565b90505b919050565b6000618e8282618ec8565b90505b919050565b600081151590505b919050565b60008190505b919050565b6000618ead82618e64565b90505b919050565b6000618ec082618e64565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600060ff821690505b919050565b6000618f0d82618f74565b90505b919050565b6000618f2082618f28565b90505b919050565b6000618f3382618ec8565b90505b919050565b6000618f4682618f4e565b90505b919050565b6000618f5982618ec8565b90505b919050565b6000618f6c82618ee9565b90505b919050565b6000618f7f82618f87565b90505b919050565b6000618f9282618ec8565b90505b919050565b828183376000838301525b505050565b60005b83811015618fc95780820151818401525b602081019050618fad565b83811115618fd8576000848401525b505b505050565b6000618ff2618fed8361902b565b618e59565b90505b919050565b60008190505b919050565b60006190118254618fdf565b90505b919050565b6000601f19601f83011690505b919050565b60008160001c90505b919050565b61904281618e64565b811415156190505760006000fd5b5b50565b61905d81618e8a565b8114151561906b5760006000fd5b5b50565b61907881618e97565b811415156190865760006000fd5b5b50565b61909381618ea2565b811415156190a15760006000fd5b5b50565b6190ae81618eb5565b811415156190bc5760006000fd5b5b50565b6190c981618ee9565b811415156190d75760006000fd5b5b50565b6190e481618ef4565b811415156190f25760006000fd5b5b50565bfe45524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220e252211643abc87eb0d0e835a3b3e51e34917240abebbd607b6470a09614585f64736f6c63430006050033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _lableAddress common.Address, platform_token common.Address, _controlledTokenBuilder common.Address, _nft common.Address, _itokenid common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _owner, _lableAddress, platform_token, _controlledTokenBuilder, _nft, _itokenid)
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

// DEFAULDECIMALPLACES is a free data retrieval call binding the contract method 0x3eafcaf2.
//
// Solidity: function DEFAUL_DECIMAL_PLACES() view returns(uint256)
func (_Store *StoreCaller) DEFAULDECIMALPLACES(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DEFAUL_DECIMAL_PLACES")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DEFAULDECIMALPLACES is a free data retrieval call binding the contract method 0x3eafcaf2.
//
// Solidity: function DEFAUL_DECIMAL_PLACES() view returns(uint256)
func (_Store *StoreSession) DEFAULDECIMALPLACES() (*big.Int, error) {
	return _Store.Contract.DEFAULDECIMALPLACES(&_Store.CallOpts)
}

// DEFAULDECIMALPLACES is a free data retrieval call binding the contract method 0x3eafcaf2.
//
// Solidity: function DEFAUL_DECIMAL_PLACES() view returns(uint256)
func (_Store *StoreCallerSession) DEFAULDECIMALPLACES() (*big.Int, error) {
	return _Store.Contract.DEFAULDECIMALPLACES(&_Store.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Store *StoreCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Store *StoreSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Store.Contract.DOMAINSEPARATOR(&_Store.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Store *StoreCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Store.Contract.DOMAINSEPARATOR(&_Store.CallOpts)
}

// DrawNumber is a free data retrieval call binding the contract method 0x5793cb6b.
//
// Solidity: function DrawNumber() view returns(uint256)
func (_Store *StoreCaller) DrawNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "DrawNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DrawNumber is a free data retrieval call binding the contract method 0x5793cb6b.
//
// Solidity: function DrawNumber() view returns(uint256)
func (_Store *StoreSession) DrawNumber() (*big.Int, error) {
	return _Store.Contract.DrawNumber(&_Store.CallOpts)
}

// DrawNumber is a free data retrieval call binding the contract method 0x5793cb6b.
//
// Solidity: function DrawNumber() view returns(uint256)
func (_Store *StoreCallerSession) DrawNumber() (*big.Int, error) {
	return _Store.Contract.DrawNumber(&_Store.CallOpts)
}

// MAXIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x12eb1fbf.
//
// Solidity: function MAX_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreCaller) MAXIMAGELINKLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MAX_IMAGE_LINK_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x12eb1fbf.
//
// Solidity: function MAX_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreSession) MAXIMAGELINKLENGTH() (*big.Int, error) {
	return _Store.Contract.MAXIMAGELINKLENGTH(&_Store.CallOpts)
}

// MAXIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x12eb1fbf.
//
// Solidity: function MAX_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreCallerSession) MAXIMAGELINKLENGTH() (*big.Int, error) {
	return _Store.Contract.MAXIMAGELINKLENGTH(&_Store.CallOpts)
}

// MINIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x9fcaf49a.
//
// Solidity: function MIN_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreCaller) MINIMAGELINKLENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MIN_IMAGE_LINK_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x9fcaf49a.
//
// Solidity: function MIN_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreSession) MINIMAGELINKLENGTH() (*big.Int, error) {
	return _Store.Contract.MINIMAGELINKLENGTH(&_Store.CallOpts)
}

// MINIMAGELINKLENGTH is a free data retrieval call binding the contract method 0x9fcaf49a.
//
// Solidity: function MIN_IMAGE_LINK_LENGTH() view returns(uint256)
func (_Store *StoreCallerSession) MINIMAGELINKLENGTH() (*big.Int, error) {
	return _Store.Contract.MINIMAGELINKLENGTH(&_Store.CallOpts)
}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Store *StoreCaller) MINNAMELENGTH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MIN_NAME_LENGTH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Store *StoreSession) MINNAMELENGTH() (*big.Int, error) {
	return _Store.Contract.MINNAMELENGTH(&_Store.CallOpts)
}

// MINNAMELENGTH is a free data retrieval call binding the contract method 0x50cfeddd.
//
// Solidity: function MIN_NAME_LENGTH() view returns(uint256)
func (_Store *StoreCallerSession) MINNAMELENGTH() (*big.Int, error) {
	return _Store.Contract.MINNAMELENGTH(&_Store.CallOpts)
}

// MIXTRUELOWLEVELNUMBER is a free data retrieval call binding the contract method 0x6067f59f.
//
// Solidity: function MIX_TRUE_LOW_LEVEL_NUMBER() view returns(uint256)
func (_Store *StoreCaller) MIXTRUELOWLEVELNUMBER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "MIX_TRUE_LOW_LEVEL_NUMBER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MIXTRUELOWLEVELNUMBER is a free data retrieval call binding the contract method 0x6067f59f.
//
// Solidity: function MIX_TRUE_LOW_LEVEL_NUMBER() view returns(uint256)
func (_Store *StoreSession) MIXTRUELOWLEVELNUMBER() (*big.Int, error) {
	return _Store.Contract.MIXTRUELOWLEVELNUMBER(&_Store.CallOpts)
}

// MIXTRUELOWLEVELNUMBER is a free data retrieval call binding the contract method 0x6067f59f.
//
// Solidity: function MIX_TRUE_LOW_LEVEL_NUMBER() view returns(uint256)
func (_Store *StoreCallerSession) MIXTRUELOWLEVELNUMBER() (*big.Int, error) {
	return _Store.Contract.MIXTRUELOWLEVELNUMBER(&_Store.CallOpts)
}

// QueryBox is a free data retrieval call binding the contract method 0x868a686d.
//
// Solidity: function QueryBox(uint256 _series_id) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])))
func (_Store *StoreCaller) QueryBox(opts *bind.CallOpts, _series_id *big.Int) (BlindBoxBox, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryBox", _series_id)

	if err != nil {
		return *new(BlindBoxBox), err
	}

	out0 := *abi.ConvertType(out[0], new(BlindBoxBox)).(*BlindBoxBox)

	return out0, err

}

// QueryBox is a free data retrieval call binding the contract method 0x868a686d.
//
// Solidity: function QueryBox(uint256 _series_id) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])))
func (_Store *StoreSession) QueryBox(_series_id *big.Int) (BlindBoxBox, error) {
	return _Store.Contract.QueryBox(&_Store.CallOpts, _series_id)
}

// QueryBox is a free data retrieval call binding the contract method 0x868a686d.
//
// Solidity: function QueryBox(uint256 _series_id) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])))
func (_Store *StoreCallerSession) QueryBox(_series_id *big.Int) (BlindBoxBox, error) {
	return _Store.Contract.QueryBox(&_Store.CallOpts, _series_id)
}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreCaller) QueryBoxs(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]BlindBoxBox, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryBoxs", start, end)

	if err != nil {
		return *new([]BlindBoxBox), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxBox)).(*[]BlindBoxBox)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreSession) QueryBoxs(start *big.Int, end *big.Int) ([]BlindBoxBox, *big.Int, error) {
	return _Store.Contract.QueryBoxs(&_Store.CallOpts, start, end)
}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreCallerSession) QueryBoxs(start *big.Int, end *big.Int) ([]BlindBoxBox, *big.Int, error) {
	return _Store.Contract.QueryBoxs(&_Store.CallOpts, start, end)
}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address,address))
func (_Store *StoreCaller) QueryConfig(opts *bind.CallOpts) (BlindBoxConfig, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryConfig")

	if err != nil {
		return *new(BlindBoxConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BlindBoxConfig)).(*BlindBoxConfig)

	return out0, err

}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address,address))
func (_Store *StoreSession) QueryConfig() (BlindBoxConfig, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address,address))
func (_Store *StoreCallerSession) QueryConfig() (BlindBoxConfig, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// QueryDraws is a free data retrieval call binding the contract method 0x7d54cf71.
//
// Solidity: function QueryDraws(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCaller) QueryDraws(opts *bind.CallOpts, _series_id *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryDraws", _series_id)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// QueryDraws is a free data retrieval call binding the contract method 0x7d54cf71.
//
// Solidity: function QueryDraws(uint256 _series_id) view returns(uint256[])
func (_Store *StoreSession) QueryDraws(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryDraws(&_Store.CallOpts, _series_id)
}

// QueryDraws is a free data retrieval call binding the contract method 0x7d54cf71.
//
// Solidity: function QueryDraws(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCallerSession) QueryDraws(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryDraws(&_Store.CallOpts, _series_id)
}

// QueryImage is a free data retrieval call binding the contract method 0x68c8ca0d.
//
// Solidity: function QueryImage(uint256 _series_id) view returns(string)
func (_Store *StoreCaller) QueryImage(opts *bind.CallOpts, _series_id *big.Int) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryImage", _series_id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// QueryImage is a free data retrieval call binding the contract method 0x68c8ca0d.
//
// Solidity: function QueryImage(uint256 _series_id) view returns(string)
func (_Store *StoreSession) QueryImage(_series_id *big.Int) (string, error) {
	return _Store.Contract.QueryImage(&_Store.CallOpts, _series_id)
}

// QueryImage is a free data retrieval call binding the contract method 0x68c8ca0d.
//
// Solidity: function QueryImage(uint256 _series_id) view returns(string)
func (_Store *StoreCallerSession) QueryImage(_series_id *big.Int) (string, error) {
	return _Store.Contract.QueryImage(&_Store.CallOpts, _series_id)
}

// QueryLevels is a free data retrieval call binding the contract method 0x7790f6ba.
//
// Solidity: function QueryLevels(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCaller) QueryLevels(opts *bind.CallOpts, _series_id *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryLevels", _series_id)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// QueryLevels is a free data retrieval call binding the contract method 0x7790f6ba.
//
// Solidity: function QueryLevels(uint256 _series_id) view returns(uint256[])
func (_Store *StoreSession) QueryLevels(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryLevels(&_Store.CallOpts, _series_id)
}

// QueryLevels is a free data retrieval call binding the contract method 0x7790f6ba.
//
// Solidity: function QueryLevels(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCallerSession) QueryLevels(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryLevels(&_Store.CallOpts, _series_id)
}

// QueryMix is a free data retrieval call binding the contract method 0xf229c71c.
//
// Solidity: function QueryMix(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCaller) QueryMix(opts *bind.CallOpts, _series_id *big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryMix", _series_id)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// QueryMix is a free data retrieval call binding the contract method 0xf229c71c.
//
// Solidity: function QueryMix(uint256 _series_id) view returns(uint256[])
func (_Store *StoreSession) QueryMix(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryMix(&_Store.CallOpts, _series_id)
}

// QueryMix is a free data retrieval call binding the contract method 0xf229c71c.
//
// Solidity: function QueryMix(uint256 _series_id) view returns(uint256[])
func (_Store *StoreCallerSession) QueryMix(_series_id *big.Int) ([]*big.Int, error) {
	return _Store.Contract.QueryMix(&_Store.CallOpts, _series_id)
}

// QueryRatio is a free data retrieval call binding the contract method 0xca8585e0.
//
// Solidity: function QueryRatio() view returns(uint256, uint256)
func (_Store *StoreCaller) QueryRatio(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryRatio")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QueryRatio is a free data retrieval call binding the contract method 0xca8585e0.
//
// Solidity: function QueryRatio() view returns(uint256, uint256)
func (_Store *StoreSession) QueryRatio() (*big.Int, *big.Int, error) {
	return _Store.Contract.QueryRatio(&_Store.CallOpts)
}

// QueryRatio is a free data retrieval call binding the contract method 0xca8585e0.
//
// Solidity: function QueryRatio() view returns(uint256, uint256)
func (_Store *StoreCallerSession) QueryRatio() (*big.Int, *big.Int, error) {
	return _Store.Contract.QueryRatio(&_Store.CallOpts)
}

// QuerySeriesIds is a free data retrieval call binding the contract method 0x4fe5d790.
//
// Solidity: function QuerySeriesIds() view returns(uint256[])
func (_Store *StoreCaller) QuerySeriesIds(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QuerySeriesIds")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// QuerySeriesIds is a free data retrieval call binding the contract method 0x4fe5d790.
//
// Solidity: function QuerySeriesIds() view returns(uint256[])
func (_Store *StoreSession) QuerySeriesIds() ([]*big.Int, error) {
	return _Store.Contract.QuerySeriesIds(&_Store.CallOpts)
}

// QuerySeriesIds is a free data retrieval call binding the contract method 0x4fe5d790.
//
// Solidity: function QuerySeriesIds() view returns(uint256[])
func (_Store *StoreCallerSession) QuerySeriesIds() ([]*big.Int, error) {
	return _Store.Contract.QuerySeriesIds(&_Store.CallOpts)
}

// QuerySeriesIdsNotNull is a free data retrieval call binding the contract method 0x30157fb9.
//
// Solidity: function QuerySeriesIdsNotNull() view returns(uint256[], uint256)
func (_Store *StoreCaller) QuerySeriesIdsNotNull(opts *bind.CallOpts) ([]*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QuerySeriesIdsNotNull")

	if err != nil {
		return *new([]*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QuerySeriesIdsNotNull is a free data retrieval call binding the contract method 0x30157fb9.
//
// Solidity: function QuerySeriesIdsNotNull() view returns(uint256[], uint256)
func (_Store *StoreSession) QuerySeriesIdsNotNull() ([]*big.Int, *big.Int, error) {
	return _Store.Contract.QuerySeriesIdsNotNull(&_Store.CallOpts)
}

// QuerySeriesIdsNotNull is a free data retrieval call binding the contract method 0x30157fb9.
//
// Solidity: function QuerySeriesIdsNotNull() view returns(uint256[], uint256)
func (_Store *StoreCallerSession) QuerySeriesIdsNotNull() ([]*big.Int, *big.Int, error) {
	return _Store.Contract.QuerySeriesIdsNotNull(&_Store.CallOpts)
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

// BoxInfo is a free data retrieval call binding the contract method 0xaf7a19ed.
//
// Solidity: function box_info(uint256 ) view returns(string name, uint256 series_id, string image, (address[],uint256[]) reward)
func (_Store *StoreCaller) BoxInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Reward   BlindBoxReward
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "box_info", arg0)

	outstruct := new(struct {
		Name     string
		SeriesId *big.Int
		Image    string
		Reward   BlindBoxReward
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.SeriesId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Image = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Reward = *abi.ConvertType(out[3], new(BlindBoxReward)).(*BlindBoxReward)

	return *outstruct, err

}

// BoxInfo is a free data retrieval call binding the contract method 0xaf7a19ed.
//
// Solidity: function box_info(uint256 ) view returns(string name, uint256 series_id, string image, (address[],uint256[]) reward)
func (_Store *StoreSession) BoxInfo(arg0 *big.Int) (struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Reward   BlindBoxReward
}, error) {
	return _Store.Contract.BoxInfo(&_Store.CallOpts, arg0)
}

// BoxInfo is a free data retrieval call binding the contract method 0xaf7a19ed.
//
// Solidity: function box_info(uint256 ) view returns(string name, uint256 series_id, string image, (address[],uint256[]) reward)
func (_Store *StoreCallerSession) BoxInfo(arg0 *big.Int) (struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Reward   BlindBoxReward
}, error) {
	return _Store.Contract.BoxInfo(&_Store.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft, address iTokenId, address box_pro)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
	ITokenId      common.Address
	BoxPro        common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner         common.Address
		LableAddress  common.Address
		PlatformToken common.Address
		KeyToken      common.Address
		PrizePool     common.Address
		Flip          common.Address
		Nft           common.Address
		ITokenId      common.Address
		BoxPro        common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LableAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PlatformToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.KeyToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.PrizePool = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Flip = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.ITokenId = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)
	outstruct.BoxPro = *abi.ConvertType(out[8], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft, address iTokenId, address box_pro)
func (_Store *StoreSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
	ITokenId      common.Address
	BoxPro        common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft, address iTokenId, address box_pro)
func (_Store *StoreCallerSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
	ITokenId      common.Address
	BoxPro        common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// ControlledTokenBuilder is a free data retrieval call binding the contract method 0x14d9dcd0.
//
// Solidity: function controlledTokenBuilder() view returns(address)
func (_Store *StoreCaller) ControlledTokenBuilder(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "controlledTokenBuilder")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ControlledTokenBuilder is a free data retrieval call binding the contract method 0x14d9dcd0.
//
// Solidity: function controlledTokenBuilder() view returns(address)
func (_Store *StoreSession) ControlledTokenBuilder() (common.Address, error) {
	return _Store.Contract.ControlledTokenBuilder(&_Store.CallOpts)
}

// ControlledTokenBuilder is a free data retrieval call binding the contract method 0x14d9dcd0.
//
// Solidity: function controlledTokenBuilder() view returns(address)
func (_Store *StoreCallerSession) ControlledTokenBuilder() (common.Address, error) {
	return _Store.Contract.ControlledTokenBuilder(&_Store.CallOpts)
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

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Store *StoreCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Store *StoreSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Store.Contract.Nonces(&_Store.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Store *StoreCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Store.Contract.Nonces(&_Store.CallOpts, owner)
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

// Convert is a paid mutator transaction binding the contract method 0x87ab1642.
//
// Solidity: function Convert(uint256 _series_id, uint256[] _token_ids) returns()
func (_Store *StoreTransactor) Convert(opts *bind.TransactOpts, _series_id *big.Int, _token_ids []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Convert", _series_id, _token_ids)
}

// Convert is a paid mutator transaction binding the contract method 0x87ab1642.
//
// Solidity: function Convert(uint256 _series_id, uint256[] _token_ids) returns()
func (_Store *StoreSession) Convert(_series_id *big.Int, _token_ids []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.Convert(&_Store.TransactOpts, _series_id, _token_ids)
}

// Convert is a paid mutator transaction binding the contract method 0x87ab1642.
//
// Solidity: function Convert(uint256 _series_id, uint256[] _token_ids) returns()
func (_Store *StoreTransactorSession) Convert(_series_id *big.Int, _token_ids []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.Convert(&_Store.TransactOpts, _series_id, _token_ids)
}

// Draw is a paid mutator transaction binding the contract method 0xc4181611.
//
// Solidity: function Draw(uint256 _number, address _inviter) returns()
func (_Store *StoreTransactor) Draw(opts *bind.TransactOpts, _number *big.Int, _inviter common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "Draw", _number, _inviter)
}

// Draw is a paid mutator transaction binding the contract method 0xc4181611.
//
// Solidity: function Draw(uint256 _number, address _inviter) returns()
func (_Store *StoreSession) Draw(_number *big.Int, _inviter common.Address) (*types.Transaction, error) {
	return _Store.Contract.Draw(&_Store.TransactOpts, _number, _inviter)
}

// Draw is a paid mutator transaction binding the contract method 0xc4181611.
//
// Solidity: function Draw(uint256 _number, address _inviter) returns()
func (_Store *StoreTransactorSession) Draw(_number *big.Int, _inviter common.Address) (*types.Transaction, error) {
	return _Store.Contract.Draw(&_Store.TransactOpts, _number, _inviter)
}

// DrawOut is a paid mutator transaction binding the contract method 0xf05a020f.
//
// Solidity: function DrawOut(uint256 _series_id, uint256 _number) returns()
func (_Store *StoreTransactor) DrawOut(opts *bind.TransactOpts, _series_id *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "DrawOut", _series_id, _number)
}

// DrawOut is a paid mutator transaction binding the contract method 0xf05a020f.
//
// Solidity: function DrawOut(uint256 _series_id, uint256 _number) returns()
func (_Store *StoreSession) DrawOut(_series_id *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DrawOut(&_Store.TransactOpts, _series_id, _number)
}

// DrawOut is a paid mutator transaction binding the contract method 0xf05a020f.
//
// Solidity: function DrawOut(uint256 _series_id, uint256 _number) returns()
func (_Store *StoreTransactorSession) DrawOut(_series_id *big.Int, _number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.DrawOut(&_Store.TransactOpts, _series_id, _number)
}

// MintBox is a paid mutator transaction binding the contract method 0xe9292c47.
//
// Solidity: function MintBox((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])) _box) returns()
func (_Store *StoreTransactor) MintBox(opts *bind.TransactOpts, _box BlindBoxBox) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "MintBox", _box)
}

// MintBox is a paid mutator transaction binding the contract method 0xe9292c47.
//
// Solidity: function MintBox((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])) _box) returns()
func (_Store *StoreSession) MintBox(_box BlindBoxBox) (*types.Transaction, error) {
	return _Store.Contract.MintBox(&_Store.TransactOpts, _box)
}

// MintBox is a paid mutator transaction binding the contract method 0xe9292c47.
//
// Solidity: function MintBox((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])) _box) returns()
func (_Store *StoreTransactorSession) MintBox(_box BlindBoxBox) (*types.Transaction, error) {
	return _Store.Contract.MintBox(&_Store.TransactOpts, _box)
}

// MixTrue is a paid mutator transaction binding the contract method 0x4521e181.
//
// Solidity: function MixTrue(uint256 _series_id, uint256 _grade_id, uint256[] _tokens_id) returns()
func (_Store *StoreTransactor) MixTrue(opts *bind.TransactOpts, _series_id *big.Int, _grade_id *big.Int, _tokens_id []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "MixTrue", _series_id, _grade_id, _tokens_id)
}

// MixTrue is a paid mutator transaction binding the contract method 0x4521e181.
//
// Solidity: function MixTrue(uint256 _series_id, uint256 _grade_id, uint256[] _tokens_id) returns()
func (_Store *StoreSession) MixTrue(_series_id *big.Int, _grade_id *big.Int, _tokens_id []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.MixTrue(&_Store.TransactOpts, _series_id, _grade_id, _tokens_id)
}

// MixTrue is a paid mutator transaction binding the contract method 0x4521e181.
//
// Solidity: function MixTrue(uint256 _series_id, uint256 _grade_id, uint256[] _tokens_id) returns()
func (_Store *StoreTransactorSession) MixTrue(_series_id *big.Int, _grade_id *big.Int, _tokens_id []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.MixTrue(&_Store.TransactOpts, _series_id, _grade_id, _tokens_id)
}

// ResetDraw is a paid mutator transaction binding the contract method 0x40b6ddbf.
//
// Solidity: function ResetDraw(uint256 _series_id, uint256[] _draw) returns()
func (_Store *StoreTransactor) ResetDraw(opts *bind.TransactOpts, _series_id *big.Int, _draw []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetDraw", _series_id, _draw)
}

// ResetDraw is a paid mutator transaction binding the contract method 0x40b6ddbf.
//
// Solidity: function ResetDraw(uint256 _series_id, uint256[] _draw) returns()
func (_Store *StoreSession) ResetDraw(_series_id *big.Int, _draw []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetDraw(&_Store.TransactOpts, _series_id, _draw)
}

// ResetDraw is a paid mutator transaction binding the contract method 0x40b6ddbf.
//
// Solidity: function ResetDraw(uint256 _series_id, uint256[] _draw) returns()
func (_Store *StoreTransactorSession) ResetDraw(_series_id *big.Int, _draw []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetDraw(&_Store.TransactOpts, _series_id, _draw)
}

// ResetDrawNumber is a paid mutator transaction binding the contract method 0xd4ecd040.
//
// Solidity: function ResetDrawNumber(uint256 _drawNumber) returns()
func (_Store *StoreTransactor) ResetDrawNumber(opts *bind.TransactOpts, _drawNumber *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetDrawNumber", _drawNumber)
}

// ResetDrawNumber is a paid mutator transaction binding the contract method 0xd4ecd040.
//
// Solidity: function ResetDrawNumber(uint256 _drawNumber) returns()
func (_Store *StoreSession) ResetDrawNumber(_drawNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetDrawNumber(&_Store.TransactOpts, _drawNumber)
}

// ResetDrawNumber is a paid mutator transaction binding the contract method 0xd4ecd040.
//
// Solidity: function ResetDrawNumber(uint256 _drawNumber) returns()
func (_Store *StoreTransactorSession) ResetDrawNumber(_drawNumber *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetDrawNumber(&_Store.TransactOpts, _drawNumber)
}

// ResetMix is a paid mutator transaction binding the contract method 0x2f575b29.
//
// Solidity: function ResetMix(uint256 _series_id, uint256[] _mix) returns()
func (_Store *StoreTransactor) ResetMix(opts *bind.TransactOpts, _series_id *big.Int, _mix []*big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetMix", _series_id, _mix)
}

// ResetMix is a paid mutator transaction binding the contract method 0x2f575b29.
//
// Solidity: function ResetMix(uint256 _series_id, uint256[] _mix) returns()
func (_Store *StoreSession) ResetMix(_series_id *big.Int, _mix []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetMix(&_Store.TransactOpts, _series_id, _mix)
}

// ResetMix is a paid mutator transaction binding the contract method 0x2f575b29.
//
// Solidity: function ResetMix(uint256 _series_id, uint256[] _mix) returns()
func (_Store *StoreTransactorSession) ResetMix(_series_id *big.Int, _mix []*big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetMix(&_Store.TransactOpts, _series_id, _mix)
}

// ResetOwner is a paid mutator transaction binding the contract method 0x9c724fa8.
//
// Solidity: function ResetOwner(address _owner) returns()
func (_Store *StoreTransactor) ResetOwner(opts *bind.TransactOpts, _owner common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetOwner", _owner)
}

// ResetOwner is a paid mutator transaction binding the contract method 0x9c724fa8.
//
// Solidity: function ResetOwner(address _owner) returns()
func (_Store *StoreSession) ResetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.ResetOwner(&_Store.TransactOpts, _owner)
}

// ResetOwner is a paid mutator transaction binding the contract method 0x9c724fa8.
//
// Solidity: function ResetOwner(address _owner) returns()
func (_Store *StoreTransactorSession) ResetOwner(_owner common.Address) (*types.Transaction, error) {
	return _Store.Contract.ResetOwner(&_Store.TransactOpts, _owner)
}

// ResetRatio is a paid mutator transaction binding the contract method 0x3b667870.
//
// Solidity: function ResetRatio(uint256 _ratio) returns()
func (_Store *StoreTransactor) ResetRatio(opts *bind.TransactOpts, _ratio *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetRatio", _ratio)
}

// ResetRatio is a paid mutator transaction binding the contract method 0x3b667870.
//
// Solidity: function ResetRatio(uint256 _ratio) returns()
func (_Store *StoreSession) ResetRatio(_ratio *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetRatio(&_Store.TransactOpts, _ratio)
}

// ResetRatio is a paid mutator transaction binding the contract method 0x3b667870.
//
// Solidity: function ResetRatio(uint256 _ratio) returns()
func (_Store *StoreTransactorSession) ResetRatio(_ratio *big.Int) (*types.Transaction, error) {
	return _Store.Contract.ResetRatio(&_Store.TransactOpts, _ratio)
}

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreTransactor) ResetReward(opts *bind.TransactOpts, _series_id *big.Int, _reward BlindBoxReward) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetReward", _series_id, _reward)
}

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreSession) ResetReward(_series_id *big.Int, _reward BlindBoxReward) (*types.Transaction, error) {
	return _Store.Contract.ResetReward(&_Store.TransactOpts, _series_id, _reward)
}

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreTransactorSession) ResetReward(_series_id *big.Int, _reward BlindBoxReward) (*types.Transaction, error) {
	return _Store.Contract.ResetReward(&_Store.TransactOpts, _series_id, _reward)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address _box_pro, address handing) returns()
func (_Store *StoreTransactor) AddBoxPro(opts *bind.TransactOpts, _box_pro common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "addBoxPro", _box_pro, handing)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address _box_pro, address handing) returns()
func (_Store *StoreSession) AddBoxPro(_box_pro common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBoxPro(&_Store.TransactOpts, _box_pro, handing)
}

// AddBoxPro is a paid mutator transaction binding the contract method 0xfe49ba32.
//
// Solidity: function addBoxPro(address _box_pro, address handing) returns()
func (_Store *StoreTransactorSession) AddBoxPro(_box_pro common.Address, handing common.Address) (*types.Transaction, error) {
	return _Store.Contract.AddBoxPro(&_Store.TransactOpts, _box_pro, handing)
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

// BurnKey is a paid mutator transaction binding the contract method 0x4f920f3e.
//
// Solidity: function burnKey(address user, uint256 number) returns()
func (_Store *StoreTransactor) BurnKey(opts *bind.TransactOpts, user common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "burnKey", user, number)
}

// BurnKey is a paid mutator transaction binding the contract method 0x4f920f3e.
//
// Solidity: function burnKey(address user, uint256 number) returns()
func (_Store *StoreSession) BurnKey(user common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnKey(&_Store.TransactOpts, user, number)
}

// BurnKey is a paid mutator transaction binding the contract method 0x4f920f3e.
//
// Solidity: function burnKey(address user, uint256 number) returns()
func (_Store *StoreTransactorSession) BurnKey(user common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BurnKey(&_Store.TransactOpts, user, number)
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

// Init is a paid mutator transaction binding the contract method 0x076ac323.
//
// Solidity: function init((string,string,uint8,address) _config, address _flip, address _prize_pool) returns()
func (_Store *StoreTransactor) Init(opts *bind.TransactOpts, _config BuildTokenControlledTokenConfig, _flip common.Address, _prize_pool common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "init", _config, _flip, _prize_pool)
}

// Init is a paid mutator transaction binding the contract method 0x076ac323.
//
// Solidity: function init((string,string,uint8,address) _config, address _flip, address _prize_pool) returns()
func (_Store *StoreSession) Init(_config BuildTokenControlledTokenConfig, _flip common.Address, _prize_pool common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, _config, _flip, _prize_pool)
}

// Init is a paid mutator transaction binding the contract method 0x076ac323.
//
// Solidity: function init((string,string,uint8,address) _config, address _flip, address _prize_pool) returns()
func (_Store *StoreTransactorSession) Init(_config BuildTokenControlledTokenConfig, _flip common.Address, _prize_pool common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, _config, _flip, _prize_pool)
}

// MintKey is a paid mutator transaction binding the contract method 0x4ef969ed.
//
// Solidity: function mintKey(address sender, uint256 number) returns()
func (_Store *StoreTransactor) MintKey(opts *bind.TransactOpts, sender common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "mintKey", sender, number)
}

// MintKey is a paid mutator transaction binding the contract method 0x4ef969ed.
//
// Solidity: function mintKey(address sender, uint256 number) returns()
func (_Store *StoreSession) MintKey(sender common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintKey(&_Store.TransactOpts, sender, number)
}

// MintKey is a paid mutator transaction binding the contract method 0x4ef969ed.
//
// Solidity: function mintKey(address sender, uint256 number) returns()
func (_Store *StoreTransactorSession) MintKey(sender common.Address, number *big.Int) (*types.Transaction, error) {
	return _Store.Contract.MintKey(&_Store.TransactOpts, sender, number)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Store *StoreTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Store *StoreSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Store.Contract.Permit(&_Store.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Store *StoreTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Store.Contract.Permit(&_Store.TransactOpts, owner, spender, value, deadline, v, r, s)
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

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Store.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreSession) Receive() (*types.Transaction, error) {
	return _Store.Contract.Receive(&_Store.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Store *StoreTransactorSession) Receive() (*types.Transaction, error) {
	return _Store.Contract.Receive(&_Store.TransactOpts)
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

// StoreAddBoxProIterator is returned from FilterAddBoxPro and is used to iterate over the raw logs and unpacked data for AddBoxPro events raised by the Store contract.
type StoreAddBoxProIterator struct {
	Event *StoreAddBoxPro // Event containing the contract specifics and raw log

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
func (it *StoreAddBoxProIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreAddBoxPro)
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
		it.Event = new(StoreAddBoxPro)
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
func (it *StoreAddBoxProIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreAddBoxProIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreAddBoxPro represents a AddBoxPro event raised by the Store contract.
type StoreAddBoxPro struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterAddBoxPro is a free log retrieval operation binding the contract event 0x484ee7be515f50ca021b1d2f6bda39adb20563e2cf9041a996cf3d125346271b.
//
// Solidity: event add_box_pro(address arg0)
func (_Store *StoreFilterer) FilterAddBoxPro(opts *bind.FilterOpts) (*StoreAddBoxProIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "add_box_pro")
	if err != nil {
		return nil, err
	}
	return &StoreAddBoxProIterator{contract: _Store.contract, event: "add_box_pro", logs: logs, sub: sub}, nil
}

// WatchAddBoxPro is a free log subscription operation binding the contract event 0x484ee7be515f50ca021b1d2f6bda39adb20563e2cf9041a996cf3d125346271b.
//
// Solidity: event add_box_pro(address arg0)
func (_Store *StoreFilterer) WatchAddBoxPro(opts *bind.WatchOpts, sink chan<- *StoreAddBoxPro) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "add_box_pro")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreAddBoxPro)
				if err := _Store.contract.UnpackLog(event, "add_box_pro", log); err != nil {
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

// ParseAddBoxPro is a log parse operation binding the contract event 0x484ee7be515f50ca021b1d2f6bda39adb20563e2cf9041a996cf3d125346271b.
//
// Solidity: event add_box_pro(address arg0)
func (_Store *StoreFilterer) ParseAddBoxPro(log types.Log) (*StoreAddBoxPro, error) {
	event := new(StoreAddBoxPro)
	if err := _Store.contract.UnpackLog(event, "add_box_pro", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDrawIterator is returned from FilterDraw and is used to iterate over the raw logs and unpacked data for Draw events raised by the Store contract.
type StoreDrawIterator struct {
	Event *StoreDraw // Event containing the contract specifics and raw log

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
func (it *StoreDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDraw)
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
		it.Event = new(StoreDraw)
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
func (it *StoreDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDraw represents a Draw event raised by the Store contract.
type StoreDraw struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 common.Address
	Arg3 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDraw is a free log retrieval operation binding the contract event 0x7abdf250d54b73efb224a0fca882e6b50f7a946e7262834728106d5ec4cdf9ef.
//
// Solidity: event draw(address arg0, uint256 arg1, address arg2, uint256 arg3)
func (_Store *StoreFilterer) FilterDraw(opts *bind.FilterOpts) (*StoreDrawIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "draw")
	if err != nil {
		return nil, err
	}
	return &StoreDrawIterator{contract: _Store.contract, event: "draw", logs: logs, sub: sub}, nil
}

// WatchDraw is a free log subscription operation binding the contract event 0x7abdf250d54b73efb224a0fca882e6b50f7a946e7262834728106d5ec4cdf9ef.
//
// Solidity: event draw(address arg0, uint256 arg1, address arg2, uint256 arg3)
func (_Store *StoreFilterer) WatchDraw(opts *bind.WatchOpts, sink chan<- *StoreDraw) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "draw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDraw)
				if err := _Store.contract.UnpackLog(event, "draw", log); err != nil {
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

// ParseDraw is a log parse operation binding the contract event 0x7abdf250d54b73efb224a0fca882e6b50f7a946e7262834728106d5ec4cdf9ef.
//
// Solidity: event draw(address arg0, uint256 arg1, address arg2, uint256 arg3)
func (_Store *StoreFilterer) ParseDraw(log types.Log) (*StoreDraw, error) {
	event := new(StoreDraw)
	if err := _Store.contract.UnpackLog(event, "draw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreDrawOutIterator is returned from FilterDrawOut and is used to iterate over the raw logs and unpacked data for DrawOut events raised by the Store contract.
type StoreDrawOutIterator struct {
	Event *StoreDrawOut // Event containing the contract specifics and raw log

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
func (it *StoreDrawOutIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreDrawOut)
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
		it.Event = new(StoreDrawOut)
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
func (it *StoreDrawOutIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreDrawOutIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreDrawOut represents a DrawOut event raised by the Store contract.
type StoreDrawOut struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDrawOut is a free log retrieval operation binding the contract event 0xd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b023040.
//
// Solidity: event draw_out(address arg0, uint256 arg1, uint256 arg2)
func (_Store *StoreFilterer) FilterDrawOut(opts *bind.FilterOpts) (*StoreDrawOutIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "draw_out")
	if err != nil {
		return nil, err
	}
	return &StoreDrawOutIterator{contract: _Store.contract, event: "draw_out", logs: logs, sub: sub}, nil
}

// WatchDrawOut is a free log subscription operation binding the contract event 0xd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b023040.
//
// Solidity: event draw_out(address arg0, uint256 arg1, uint256 arg2)
func (_Store *StoreFilterer) WatchDrawOut(opts *bind.WatchOpts, sink chan<- *StoreDrawOut) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "draw_out")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreDrawOut)
				if err := _Store.contract.UnpackLog(event, "draw_out", log); err != nil {
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

// ParseDrawOut is a log parse operation binding the contract event 0xd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b023040.
//
// Solidity: event draw_out(address arg0, uint256 arg1, uint256 arg2)
func (_Store *StoreFilterer) ParseDrawOut(log types.Log) (*StoreDrawOut, error) {
	event := new(StoreDrawOut)
	if err := _Store.contract.UnpackLog(event, "draw_out", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreMintBoxIterator is returned from FilterMintBox and is used to iterate over the raw logs and unpacked data for MintBox events raised by the Store contract.
type StoreMintBoxIterator struct {
	Event *StoreMintBox // Event containing the contract specifics and raw log

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
func (it *StoreMintBoxIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMintBox)
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
		it.Event = new(StoreMintBox)
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
func (it *StoreMintBoxIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMintBoxIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMintBox represents a MintBox event raised by the Store contract.
type StoreMintBox struct {
	Arg0 *big.Int
	Arg1 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMintBox is a free log retrieval operation binding the contract event 0xc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb.
//
// Solidity: event mint_box(uint256 arg0, string arg1)
func (_Store *StoreFilterer) FilterMintBox(opts *bind.FilterOpts) (*StoreMintBoxIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "mint_box")
	if err != nil {
		return nil, err
	}
	return &StoreMintBoxIterator{contract: _Store.contract, event: "mint_box", logs: logs, sub: sub}, nil
}

// WatchMintBox is a free log subscription operation binding the contract event 0xc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb.
//
// Solidity: event mint_box(uint256 arg0, string arg1)
func (_Store *StoreFilterer) WatchMintBox(opts *bind.WatchOpts, sink chan<- *StoreMintBox) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "mint_box")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMintBox)
				if err := _Store.contract.UnpackLog(event, "mint_box", log); err != nil {
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

// ParseMintBox is a log parse operation binding the contract event 0xc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb.
//
// Solidity: event mint_box(uint256 arg0, string arg1)
func (_Store *StoreFilterer) ParseMintBox(log types.Log) (*StoreMintBox, error) {
	event := new(StoreMintBox)
	if err := _Store.contract.UnpackLog(event, "mint_box", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreMixTrueIterator is returned from FilterMixTrue and is used to iterate over the raw logs and unpacked data for MixTrue events raised by the Store contract.
type StoreMixTrueIterator struct {
	Event *StoreMixTrue // Event containing the contract specifics and raw log

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
func (it *StoreMixTrueIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreMixTrue)
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
		it.Event = new(StoreMixTrue)
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
func (it *StoreMixTrueIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreMixTrueIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreMixTrue represents a MixTrue event raised by the Store contract.
type StoreMixTrue struct {
	Arg0 common.Address
	Arg1 *big.Int
	Arg2 *big.Int
	Arg3 bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMixTrue is a free log retrieval operation binding the contract event 0xd96b41f9c788de8f2d41e08347bee74702c2159a4d42149bc55b71731ed66db3.
//
// Solidity: event mix_true(address arg0, uint256 arg1, uint256 arg2, bool arg3)
func (_Store *StoreFilterer) FilterMixTrue(opts *bind.FilterOpts) (*StoreMixTrueIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "mix_true")
	if err != nil {
		return nil, err
	}
	return &StoreMixTrueIterator{contract: _Store.contract, event: "mix_true", logs: logs, sub: sub}, nil
}

// WatchMixTrue is a free log subscription operation binding the contract event 0xd96b41f9c788de8f2d41e08347bee74702c2159a4d42149bc55b71731ed66db3.
//
// Solidity: event mix_true(address arg0, uint256 arg1, uint256 arg2, bool arg3)
func (_Store *StoreFilterer) WatchMixTrue(opts *bind.WatchOpts, sink chan<- *StoreMixTrue) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "mix_true")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreMixTrue)
				if err := _Store.contract.UnpackLog(event, "mix_true", log); err != nil {
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

// ParseMixTrue is a log parse operation binding the contract event 0xd96b41f9c788de8f2d41e08347bee74702c2159a4d42149bc55b71731ed66db3.
//
// Solidity: event mix_true(address arg0, uint256 arg1, uint256 arg2, bool arg3)
func (_Store *StoreFilterer) ParseMixTrue(log types.Log) (*StoreMixTrue, error) {
	event := new(StoreMixTrue)
	if err := _Store.contract.UnpackLog(event, "mix_true", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetDrawIterator is returned from FilterResetDraw and is used to iterate over the raw logs and unpacked data for ResetDraw events raised by the Store contract.
type StoreResetDrawIterator struct {
	Event *StoreResetDraw // Event containing the contract specifics and raw log

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
func (it *StoreResetDrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetDraw)
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
		it.Event = new(StoreResetDraw)
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
func (it *StoreResetDrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetDrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetDraw represents a ResetDraw event raised by the Store contract.
type StoreResetDraw struct {
	Arg0 *big.Int
	Arg1 []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResetDraw is a free log retrieval operation binding the contract event 0x64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd48076.
//
// Solidity: event resetDraw(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) FilterResetDraw(opts *bind.FilterOpts) (*StoreResetDrawIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "resetDraw")
	if err != nil {
		return nil, err
	}
	return &StoreResetDrawIterator{contract: _Store.contract, event: "resetDraw", logs: logs, sub: sub}, nil
}

// WatchResetDraw is a free log subscription operation binding the contract event 0x64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd48076.
//
// Solidity: event resetDraw(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) WatchResetDraw(opts *bind.WatchOpts, sink chan<- *StoreResetDraw) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "resetDraw")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetDraw)
				if err := _Store.contract.UnpackLog(event, "resetDraw", log); err != nil {
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

// ParseResetDraw is a log parse operation binding the contract event 0x64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd48076.
//
// Solidity: event resetDraw(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) ParseResetDraw(log types.Log) (*StoreResetDraw, error) {
	event := new(StoreResetDraw)
	if err := _Store.contract.UnpackLog(event, "resetDraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetLevelIterator is returned from FilterResetLevel and is used to iterate over the raw logs and unpacked data for ResetLevel events raised by the Store contract.
type StoreResetLevelIterator struct {
	Event *StoreResetLevel // Event containing the contract specifics and raw log

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
func (it *StoreResetLevelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetLevel)
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
		it.Event = new(StoreResetLevel)
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
func (it *StoreResetLevelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetLevelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetLevel represents a ResetLevel event raised by the Store contract.
type StoreResetLevel struct {
	Arg0 *big.Int
	Arg1 []*big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResetLevel is a free log retrieval operation binding the contract event 0x37965ffa825886d8bd15f46af1272eaf388a47846f8570934222be344b67cc35.
//
// Solidity: event resetLevel(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) FilterResetLevel(opts *bind.FilterOpts) (*StoreResetLevelIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "resetLevel")
	if err != nil {
		return nil, err
	}
	return &StoreResetLevelIterator{contract: _Store.contract, event: "resetLevel", logs: logs, sub: sub}, nil
}

// WatchResetLevel is a free log subscription operation binding the contract event 0x37965ffa825886d8bd15f46af1272eaf388a47846f8570934222be344b67cc35.
//
// Solidity: event resetLevel(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) WatchResetLevel(opts *bind.WatchOpts, sink chan<- *StoreResetLevel) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "resetLevel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetLevel)
				if err := _Store.contract.UnpackLog(event, "resetLevel", log); err != nil {
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

// ParseResetLevel is a log parse operation binding the contract event 0x37965ffa825886d8bd15f46af1272eaf388a47846f8570934222be344b67cc35.
//
// Solidity: event resetLevel(uint256 arg0, uint256[] arg1)
func (_Store *StoreFilterer) ParseResetLevel(log types.Log) (*StoreResetLevel, error) {
	event := new(StoreResetLevel)
	if err := _Store.contract.UnpackLog(event, "resetLevel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetMixIterator is returned from FilterResetMix and is used to iterate over the raw logs and unpacked data for ResetMix events raised by the Store contract.
type StoreResetMixIterator struct {
	Event *StoreResetMix // Event containing the contract specifics and raw log

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
func (it *StoreResetMixIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetMix)
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
		it.Event = new(StoreResetMix)
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
func (it *StoreResetMixIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetMixIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetMix represents a ResetMix event raised by the Store contract.
type StoreResetMix struct {
	SeriesId *big.Int
	Arg1     []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResetMix is a free log retrieval operation binding the contract event 0x7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a38.
//
// Solidity: event resetMix(uint256 _series_id, uint256[] arg1)
func (_Store *StoreFilterer) FilterResetMix(opts *bind.FilterOpts) (*StoreResetMixIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "resetMix")
	if err != nil {
		return nil, err
	}
	return &StoreResetMixIterator{contract: _Store.contract, event: "resetMix", logs: logs, sub: sub}, nil
}

// WatchResetMix is a free log subscription operation binding the contract event 0x7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a38.
//
// Solidity: event resetMix(uint256 _series_id, uint256[] arg1)
func (_Store *StoreFilterer) WatchResetMix(opts *bind.WatchOpts, sink chan<- *StoreResetMix) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "resetMix")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetMix)
				if err := _Store.contract.UnpackLog(event, "resetMix", log); err != nil {
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

// ParseResetMix is a log parse operation binding the contract event 0x7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a38.
//
// Solidity: event resetMix(uint256 _series_id, uint256[] arg1)
func (_Store *StoreFilterer) ParseResetMix(log types.Log) (*StoreResetMix, error) {
	event := new(StoreResetMix)
	if err := _Store.contract.UnpackLog(event, "resetMix", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetOwnerIterator is returned from FilterResetOwner and is used to iterate over the raw logs and unpacked data for ResetOwner events raised by the Store contract.
type StoreResetOwnerIterator struct {
	Event *StoreResetOwner // Event containing the contract specifics and raw log

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
func (it *StoreResetOwnerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetOwner)
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
		it.Event = new(StoreResetOwner)
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
func (it *StoreResetOwnerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetOwnerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetOwner represents a ResetOwner event raised by the Store contract.
type StoreResetOwner struct {
	Arg0 common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResetOwner is a free log retrieval operation binding the contract event 0x73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a.
//
// Solidity: event resetOwner(address arg0)
func (_Store *StoreFilterer) FilterResetOwner(opts *bind.FilterOpts) (*StoreResetOwnerIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "resetOwner")
	if err != nil {
		return nil, err
	}
	return &StoreResetOwnerIterator{contract: _Store.contract, event: "resetOwner", logs: logs, sub: sub}, nil
}

// WatchResetOwner is a free log subscription operation binding the contract event 0x73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a.
//
// Solidity: event resetOwner(address arg0)
func (_Store *StoreFilterer) WatchResetOwner(opts *bind.WatchOpts, sink chan<- *StoreResetOwner) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "resetOwner")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetOwner)
				if err := _Store.contract.UnpackLog(event, "resetOwner", log); err != nil {
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

// ParseResetOwner is a log parse operation binding the contract event 0x73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a.
//
// Solidity: event resetOwner(address arg0)
func (_Store *StoreFilterer) ParseResetOwner(log types.Log) (*StoreResetOwner, error) {
	event := new(StoreResetOwner)
	if err := _Store.contract.UnpackLog(event, "resetOwner", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetRewardIterator is returned from FilterResetReward and is used to iterate over the raw logs and unpacked data for ResetReward events raised by the Store contract.
type StoreResetRewardIterator struct {
	Event *StoreResetReward // Event containing the contract specifics and raw log

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
func (it *StoreResetRewardIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetReward)
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
		it.Event = new(StoreResetReward)
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
func (it *StoreResetRewardIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetRewardIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetReward represents a ResetReward event raised by the Store contract.
type StoreResetReward struct {
	SeriesId *big.Int
	Arg1     BlindBoxReward
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterResetReward is a free log retrieval operation binding the contract event 0x13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d.
//
// Solidity: event resetReward(uint256 _series_id, (address[],uint256[]) arg1)
func (_Store *StoreFilterer) FilterResetReward(opts *bind.FilterOpts) (*StoreResetRewardIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "resetReward")
	if err != nil {
		return nil, err
	}
	return &StoreResetRewardIterator{contract: _Store.contract, event: "resetReward", logs: logs, sub: sub}, nil
}

// WatchResetReward is a free log subscription operation binding the contract event 0x13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d.
//
// Solidity: event resetReward(uint256 _series_id, (address[],uint256[]) arg1)
func (_Store *StoreFilterer) WatchResetReward(opts *bind.WatchOpts, sink chan<- *StoreResetReward) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "resetReward")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetReward)
				if err := _Store.contract.UnpackLog(event, "resetReward", log); err != nil {
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

// ParseResetReward is a log parse operation binding the contract event 0x13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d.
//
// Solidity: event resetReward(uint256 _series_id, (address[],uint256[]) arg1)
func (_Store *StoreFilterer) ParseResetReward(log types.Log) (*StoreResetReward, error) {
	event := new(StoreResetReward)
	if err := _Store.contract.UnpackLog(event, "resetReward", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StoreResetRatioIterator is returned from FilterResetRatio and is used to iterate over the raw logs and unpacked data for ResetRatio events raised by the Store contract.
type StoreResetRatioIterator struct {
	Event *StoreResetRatio // Event containing the contract specifics and raw log

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
func (it *StoreResetRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreResetRatio)
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
		it.Event = new(StoreResetRatio)
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
func (it *StoreResetRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreResetRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreResetRatio represents a ResetRatio event raised by the Store contract.
type StoreResetRatio struct {
	Arg0 *big.Int
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterResetRatio is a free log retrieval operation binding the contract event 0x9ef7e80651fd9de0992e3b6fca0d2faee38eaf176212ee59054d2654958836f6.
//
// Solidity: event reset_ratio(uint256 arg0)
func (_Store *StoreFilterer) FilterResetRatio(opts *bind.FilterOpts) (*StoreResetRatioIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "reset_ratio")
	if err != nil {
		return nil, err
	}
	return &StoreResetRatioIterator{contract: _Store.contract, event: "reset_ratio", logs: logs, sub: sub}, nil
}

// WatchResetRatio is a free log subscription operation binding the contract event 0x9ef7e80651fd9de0992e3b6fca0d2faee38eaf176212ee59054d2654958836f6.
//
// Solidity: event reset_ratio(uint256 arg0)
func (_Store *StoreFilterer) WatchResetRatio(opts *bind.WatchOpts, sink chan<- *StoreResetRatio) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "reset_ratio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreResetRatio)
				if err := _Store.contract.UnpackLog(event, "reset_ratio", log); err != nil {
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

// ParseResetRatio is a log parse operation binding the contract event 0x9ef7e80651fd9de0992e3b6fca0d2faee38eaf176212ee59054d2654958836f6.
//
// Solidity: event reset_ratio(uint256 arg0)
func (_Store *StoreFilterer) ParseResetRatio(log types.Log) (*StoreResetRatio, error) {
	event := new(StoreResetRatio)
	if err := _Store.contract.UnpackLog(event, "reset_ratio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
