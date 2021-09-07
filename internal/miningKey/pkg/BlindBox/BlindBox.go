// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package BlindBox

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

// StoreABI is the input ABI used to generate the binding from.
const StoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lableAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"contractBuildToken\",\"name\":\"_controlledTokenBuilder\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"draw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"draw_out\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"mint_box\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"mix_true\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetDraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetLevel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetMix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"resetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structBlindBox.Reward\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"resetReward\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reset_ratio\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_token_ids\",\"type\":\"uint256[]\"}],\"name\":\"Convert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAUL_DECIMAL_PLACES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_inviter\",\"type\":\"address\"}],\"name\":\"Draw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"DrawOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIX_TRUE_LOW_LEVEL_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box\",\"name\":\"_box\",\"type\":\"tuple\"}],\"name\":\"MintBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokens_id\",\"type\":\"uint256[]\"}],\"name\":\"MixTrue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryBox\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"QueryBoxs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBox.Box[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QueryConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"internalType\":\"structBlindBox.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryDraws\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryImage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryLevels\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QueryRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QuerySeriesIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_draw\",\"type\":\"uint256[]\"}],\"name\":\"ResetDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_mix\",\"type\":\"uint256[]\"}],\"name\":\"ResetMix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ResetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_ratio\",\"type\":\"uint256\"}],\"name\":\"ResetRatio\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"_reward\",\"type\":\"tuple\"}],\"name\":\"ResetReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"box_info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBox.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"controlledTokenBuilder\",\"outputs\":[{\"internalType\":\"contractBuildToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"},{\"internalType\":\"contractTokenControllerInterface\",\"name\":\"controller\",\"type\":\"address\"}],\"internalType\":\"structBuildToken.ControlledTokenConfig\",\"name\":\"_config\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_flip\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prize_pool\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"number\",\"type\":\"uint256\"}],\"name\":\"mintKey\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]"

// StoreBin is the compiled bytecode used for deploying new contracts.
var StoreBin = "0x60806040526103e860d66000509090553480156200001d5760006000fd5b5060405162007f1638038062007f16833981810160405262000043919081019062000392565b5b8160d360006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060e001604052808673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020015060cc60005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055509050505b5050505050620004a0566200049f565b600081519050620003738162000463565b5b92915050565b6000815190506200038b8162000481565b5b92915050565b6000600060006000600060a08688031215620003ae5760006000fd5b6000620003be8882890162000362565b9550506020620003d18882890162000362565b9450506040620003e48882890162000362565b9350506060620003f7888289016200037a565b92505060806200040a8882890162000362565b9150505b9295509295909350565b6000620004258262000442565b90505b919050565b60006200043a8262000418565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b6200046e8162000418565b811415156200047d5760006000fd5b5b50565b6200048c816200042d565b811415156200049b5760006000fd5b5b50565b5b617a6680620004b06000396000f3fe6080604052600436106102555760003560e01c806368c8ca0d116101395780639fcaf49a116100b6578063ca8585e01161007a578063ca8585e014610974578063d505accf146109a1578063dd62ed3e146109cb578063e9292c4714610a09578063f05a020f14610a33578063f209089314610a5d5761025e565b80639fcaf49a14610861578063a457c2d71461088d578063a9059cbb146108cb578063af7a19ed14610909578063c41816111461094a5761025e565b80637ecebe00116100fd5780637ecebe0014610765578063868a686d146107a357806387ab1642146107e157806395d89b411461080b5780639c724fa8146108375761025e565b806368c8ca0d1461063b57806370a08231146106795780637790f6ba146106b757806379502c55146106f55780637d54cf71146107275761025e565b8063332ef5b9116101d257806340b6ddbf1161019657806340b6ddbf146105395780634521e181146105635780634ef969ed1461058d5780634fe5d790146105b757806350cfeddd146105e35780636067f59f1461060f5761025e565b8063332ef5b91461044f5780633644e5151461047957806339509351146104a55780633b667870146104e35780633eafcaf21461050d5761025e565b806314d9dcd01161021957806314d9dcd01461036357806318160ddd1461038f57806323b872dd146103bb5780632f575b29146103f9578063313ce567146104235761025e565b806306fdde0314610264578063076ac32314610290578063095ea7b3146102ba57806312e7a53d146102f857806312eb1fbf146103375761025e565b3661025e575b5b005b60006000fd5b3480156102715760006000fd5b5061027a610a89565b6040516102879190616f67565b60405180910390f35b34801561029d5760006000fd5b506102b860048036036102b391908101906155d2565b610b33565b005b3480156102c75760006000fd5b506102e260048036036102dd91908101906154f9565b610e00565b6040516102ef9190616e17565b60405180910390f35b3480156103055760006000fd5b50610320600480360361031b9190810190615782565b610e2f565b60405161032e929190616dc3565b60405180910390f35b3480156103445760006000fd5b5061034d6112f2565b60405161035a91906173a1565b60405180910390f35b3480156103705760006000fd5b506103796112f7565b6040516103869190616f4b565b60405180910390f35b34801561039c5760006000fd5b506103a561131d565b6040516103b291906173a1565b60405180910390f35b3480156103c85760006000fd5b506103e360048036036103de9190810190615404565b61132f565b6040516103f09190616e17565b60405180910390f35b3480156104065760006000fd5b50610421600480360361041c91908101906156d2565b61142e565b005b3480156104305760006000fd5b506104396115a7565b60405161044691906174a4565b60405180910390f35b34801561045c5760006000fd5b506104776004803603610472919081019061572a565b6115c3565b005b3480156104865760006000fd5b5061048f6117b9565b60405161049c9190616e33565b60405180910390f35b3480156104b25760006000fd5b506104cd60048036036104c891908101906154f9565b6117d3565b6040516104da9190616e17565b60405180910390f35b3480156104f05760006000fd5b5061050b6004803603610506919081019061563d565b6118a6565b005b34801561051a5760006000fd5b50610523611986565b60405161053091906173a1565b60405180910390f35b3480156105465760006000fd5b50610561600480360361055c91908101906156d2565b61198b565b005b3480156105705760006000fd5b5061058b600480360361058691908101906157c1565b611b04565b005b34801561059a5760006000fd5b506105b560048036036105b091908101906154f9565b611cd5565b005b3480156105c45760006000fd5b506105cd611db6565b6040516105da9190616df4565b60405180910390f35b3480156105f05760006000fd5b506105f9611e19565b60405161060691906173a1565b60405180910390f35b34801561061c5760006000fd5b50610625611e1e565b60405161063291906173a1565b60405180910390f35b3480156106485760006000fd5b50610663600480360361065e919081019061563d565b611e23565b6040516106709190616f67565b60405180910390f35b3480156106865760006000fd5b506106a1600480360361069c919081019061539a565b611ee9565b6040516106ae91906173a1565b60405180910390f35b3480156106c45760006000fd5b506106df60048036036106da919081019061563d565b611f3d565b6040516106ec9190616df4565b60405180910390f35b3480156107025760006000fd5b5061070b611fbc565b60405161071e9796959493929190616cb9565b60405180910390f35b3480156107345760006000fd5b5061074f600480360361074a919081019061563d565b6120cf565b60405161075c9190616df4565b60405180910390f35b3480156107725760006000fd5b5061078d6004803603610788919081019061539a565b61214e565b60405161079a91906173a1565b60405180910390f35b3480156107b05760006000fd5b506107cb60048036036107c6919081019061563d565b6121b1565b6040516107d8919061733f565b60405180910390f35b3480156107ee5760006000fd5b50610809600480360361080491908101906156d2565b612566565b005b3480156108185760006000fd5b506108216126a1565b60405161082e9190616f67565b60405180910390f35b3480156108445760006000fd5b5061085f600480360361085a919081019061539a565b61274b565b005b34801561086e5760006000fd5b50610877612866565b60405161088491906173a1565b60405180910390f35b34801561089a5760006000fd5b506108b560048036036108b091908101906154f9565b61286b565b6040516108c29190616e17565b60405180910390f35b3480156108d85760006000fd5b506108f360048036036108ee91908101906154f9565b612958565b6040516109009190616e17565b60405180910390f35b3480156109165760006000fd5b50610931600480360361092c919081019061563d565b612987565b6040516109419493929190616f8a565b60405180910390f35b3480156109575760006000fd5b50610972600480360361096d9190810190615693565b612bf2565b005b3480156109815760006000fd5b5061098a613019565b60405161099892919061747a565b60405180910390f35b3480156109ae5760006000fd5b506109c960048036036109c49190810190615456565b613062565b005b3480156109d85760006000fd5b506109f360048036036109ee91908101906153c5565b613241565b604051610a0091906173a1565b60405180910390f35b348015610a165760006000fd5b50610a316004803603610a2c919081019061558e565b6132d6565b005b348015610a405760006000fd5b50610a5b6004803603610a569190810190615782565b6136b5565b005b348015610a6a5760006000fd5b50610a73613a25565b604051610a809190617362565b60405180910390f35b606060366000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b245780601f10610af957610100808354040283529160200191610b24565b820191906000526020600020905b815481529060010190602001808311610b0757829003601f168201915b50505050509050610b30565b90565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610bcb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610bc290617279565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614158015610c5c5750600073ffffffffffffffffffffffffffffffffffffffff1660cc60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b8015610cbc5750600073ffffffffffffffffffffffffffffffffffffffff1660cc60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16145b1515610cfd576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610cf490617048565b60405180910390fd5b6000610d218460000151856020015186604001518760600151613ca463ffffffff16565b90508060cc60005060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508260cc60005060050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055508160cc60005060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550505b5b505050565b6000610e20610e13613da063ffffffff16565b8484613dad63ffffffff16565b60019050610e29565b92915050565b60606000600060d5600050805490509050600081111580610e4f57508385115b80610e5957508085115b15610e6e57606080829350935050506112eb56505b600084905081851115610e815781905080505b601e8682031115610e9557601e8601905080505b6060868203604051908082528060200260200182016040528015610ed357816020015b610ec0614a5a565b815260200190600190039081610eb85790505b509050600060008890505b838110156112d75760d4600050600060d560005083815481101515610eff57fe5b906000526020600020900160005b505481526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610fc75780601f10610f9c57610100808354040283529160200191610fc7565b820191906000526020600020905b815481529060010190602001808311610faa57829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156110795780601f1061104e57610100808354040283529160200191611079565b820191906000526020600020905b81548152906001019060200180831161105c57829003601f168201915b50505050508152602001600382016000508054806020026020016040519081016040528092919081815260200182805480156110d757602002820191906000526020600020905b8160005054815260200190600101908083116110c0575b505050505081526020016004820160005080548060200260200160405190810160405280929190818152602001828054801561113557602002820191906000526020600020905b81600050548152602001906001019080831161111e575b505050505081526020016005820160005080548060200260200160405190810160405280929190818152602001828054801561119357602002820191906000526020600020905b81600050548152602001906001019080831161117c575b50505050508152602001600682016000506040518060400160405290816000820160005080548060200260200160405190810160405280929190818152602001828054801561123757602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116111ed575b505050505081526020016001820160005080548060200260200160405190810160405280929190818152602001828054801561129557602002820191906000526020600020905b81600050548152602001906001019080831161127e575b5050505050815260200150508152602001505083838151811015156112b657fe5b602002602001018190525081806001019250505b8080600101915050610ede565b50818495509550505050506112eb56505050505b9250929050565b608081565b60d360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000603560005054905061132c565b90565b6000611342848484613f8863ffffffff16565b61141e84611354613da063ffffffff16565b611413856040518060600160405280602881526020016179e460289139603460005060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060006113c6613da063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546142489092919063ffffffff16565b613dad63ffffffff16565b60019050611427565b9392505050565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156114c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016114bd90617279565b60405180910390fd5b81600060d4600050600083815260200190815260200160002060005090506000816001016000505414151515611531576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016115289061710e565b60405180910390fd5b8260d460005060008681526020019081526020016000206000506005016000509080519060200190611564929190614aa0565b507f7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a3884846040516115969291906173e7565b60405180910390a15b505b505b5050565b6000603860009054906101000a900460ff1690506115c0565b90565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561165b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161165290617279565b60405180910390fd5b81600060d46000506000838152602001908152602001600020600050905060008160010160005054141515156116c6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116bd9061710e565b60405180910390fd5b826020015151836000015151141515611714576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161170b906171d4565b60405180910390fd5b8260d460005060008681526020019081526020016000206000506006016000506000820151816000016000509080519060200190611753929190614af2565b506020820151816001016000509080519060200190611773929190614b7c565b509050507f13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d84846040516117a8929190617449565b60405180910390a15b505b505b5050565b60006117c96142a463ffffffff16565b90506117d0565b90565b60006118976117e6613da063ffffffff16565b8461188c8560346000506000611800613da063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546142ee90919063ffffffff16565b613dad63ffffffff16565b600190506118a0565b92915050565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561193e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161193590617279565b60405180910390fd5b8060d66000508190909055507f9ef7e80651fd9de0992e3b6fca0d2faee38eaf176212ee59054d2654958836f68160405161197991906173a1565b60405180910390a15b5b50565b606481565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a23576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a1a90617279565b60405180910390fd5b81600060d4600050600083815260200190815260200160002060005090506000816001016000505414151515611a8e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a859061710e565b60405180910390fd5b8260d460005060008681526020019081526020016000206000506004016000509080519060200190611ac1929190614aa0565b507f64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd480768484604051611af39291906173e7565b60405180910390a15b505b505b5050565b82600060d4600050600083815260200190815260200160002060005090506000816001016000505414151515611b6f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b669061710e565b60405180910390fd5b8360058111151515611bb6576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bad90617192565b60405180910390fd5b858460058151141515611bfe576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611bf590617171565b60405180910390fd5b600060d460005060008a8152602001908152602001600020600050905060cc60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a2d8aa3338b84600501600050856003016000508d8d6040518763ffffffff1660e01b8152600401611c9296959493929190616b9a565b600060405180830381600087803b158015611cad5760006000fd5b505af1158015611cc2573d600060003e3d6000fd5b50505050505b5b50505b50505b50505050565b60cc60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611d6d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d6490617279565b60405180910390fd5b6001811415611d9557611d9082670de0b6b3a7640000600161434c63ffffffff16565b611db0565b611daf82678ac7230489e80000600161434c63ffffffff16565b5b5b5b5050565b606060d5600050805480602002602001604051908101604052809291908181526020018280548015611e0a57602002820191906000526020600020905b816000505481526020019060010190808311611df3575b50505050509050611e16565b90565b600481565b600581565b606060d460005060008381526020019081526020016000206000506002016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015611ed85780601f10611ead57610100808354040283529160200191611ed8565b820191906000526020600020905b815481529060010190602001808311611ebb57829003601f168201915b50505050509050611ee4565b919050565b6000603360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050611f38565b919050565b606060d46000506000838152602001908152602001600020600050600301600050805480602002602001604051908101604052809291908181526020018280548015611fab57602002820191906000526020600020905b816000505481526020019060010190808311611f94575b50505050509050611fb7565b919050565b60cc6000508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905087565b606060d4600050600083815260200190815260200160002060005060040160005080548060200260200160405190810160405280929190818152602001828054801561213d57602002820191906000526020600020905b816000505481526020019060010190808311612126575b50505050509050612149565b919050565b60006121a5609960005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506143eb909063ffffffff16565b90506121ac565b919050565b6121b9614bce565b60d460005060008381526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156122795780601f1061224e57610100808354040283529160200191612279565b820191906000526020600020905b81548152906001019060200180831161225c57829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561232b5780601f106123005761010080835404028352916020019161232b565b820191906000526020600020905b81548152906001019060200180831161230e57829003601f168201915b505050505081526020016003820160005080548060200260200160405190810160405280929190818152602001828054801561238957602002820191906000526020600020905b816000505481526020019060010190808311612372575b50505050508152602001600482016000508054806020026020016040519081016040528092919081815260200182805480156123e757602002820191906000526020600020905b8160005054815260200190600101908083116123d0575b505050505081526020016005820160005080548060200260200160405190810160405280929190818152602001828054801561244557602002820191906000526020600020905b81600050548152602001906001019080831161242e575b5050505050815260200160068201600050604051806040016040529081600082016000508054806020026020016040519081016040528092919081815260200182805480156124e957602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001906001019080831161249f575b505050505081526020016001820160005080548060200260200160405190810160405280929190818152602001828054801561254757602002820191906000526020600020905b816000505481526020019060010190808311612530575b505050505081526020015050815260200150509050612561565b919050565b81600060d46000506000838152602001908152602001600020600050905060008160010160005054141515156125d1576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016125c89061710e565b60405180910390fd5b60cc60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0c55e31338660d46000506000898152602001908152602001600020600050600301600050876040518563ffffffff1660e01b81526004016126569493929190616b46565b600060405180830381600087803b1580156126715760006000fd5b505af1158015612686573d600060003e3d6000fd5b505050506126998461440163ffffffff16565b5b505b505050565b606060376000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561273c5780601f106127115761010080835404028352916020019161273c565b820191906000526020600020905b81548152906001019060200180831161271f57829003601f168201915b50505050509050612748565b90565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156127e3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016127da90617279565b60405180910390fd5b8060cc60005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a816040516128599190616a58565b60405180910390a15b5b50565b600881565b600061294961287e613da063ffffffff16565b8461293e85604051806060016040528060258152602001617a0c60259139603460005060006128b1613da063ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546142489092919063ffffffff16565b613dad63ffffffff16565b60019050612952565b92915050565b600061297861296b613da063ffffffff16565b8484613f8863ffffffff16565b60019050612981565b92915050565b60d4600050602052806000526040600020600091509050806000016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612a395780601f10612a0e57610100808354040283529160200191612a39565b820191906000526020600020905b815481529060010190602001808311612a1c57829003601f168201915b505050505090806001016000505490806002016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612ae35780601f10612ab857610100808354040283529160200191612ae3565b820191906000526020600020905b815481529060010190602001808311612ac657829003601f168201915b5050505050908060060160005060405180604001604052908160008201600050805480602002602001604051908101604052809291908181526020018280548015612b8357602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612b39575b5050505050815260200160018201600050805480602002602001604051908101604052809291908181526020018280548015612be157602002820191906000526020600020905b816000505481526020019060010190808311612bca575b505050505081526020015050905084565b816001811480612c025750600a81145b1515612c43576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612c3a906172fd565b60405180910390fd5b600060d6600050546064670de0b6b3a76400008602811515612c6157fe5b04029050600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515612cd7576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612cce906170ab565b60405180910390fd5b600060cc60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401612d41929190616a74565b60206040518083038186803b158015612d5a5760006000fd5b505afa158015612d6f573d600060003e3d6000fd5b505050506040513d601f19601f82011682018060405250612d939190810190615668565b9050828110151515612dda576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612dd190617216565b60405180910390fd5b612e1460cc60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633308661456b63ffffffff16565b8173ffffffffffffffffffffffffffffffffffffffff1663fcd3533c600a85811515612e3c57fe5b04336040518363ffffffff1660e01b8152600401612e5b9291906173bd565b600060405180830381600087803b158015612e765760006000fd5b505af1158015612e8b573d600060003e3d6000fd5b50505050612eff60cc60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660cc60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166008600a87811515612ef257fe5b04026146a863ffffffff16565b612f6f60cc60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1660cc60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff166005606487811515612f6257fe5b04026146a863ffffffff16565b612fb760cc60005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16866005606487811515612faa57fe5b04026146a863ffffffff16565b612fd233670de0b6b3a764000088028861434c63ffffffff16565b7f7abdf250d54b73efb224a0fca882e6b50f7a946e7262834728106d5ec4cdf9ef338787866040516130079493929190616b00565b60405180910390a15050505b5b505050565b6000600060d6600050546064670de0b6b3a764000081151561303757fe5b040260d6600050546064678ac7230489e8000081151561305357fe5b04029150915061305e565b9091565b8342111515156130a7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161309e906170ed565b60405180910390fd5b6000609a60005054888888613107609960005060008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506143eb909063ffffffff16565b8960405160200161311d96959493929190616e4f565b6040516020818303038152906040528051906020012090506000613146826147e263ffffffff16565b9050600061315c8287878761482663ffffffff16565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415156131ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131c5906171f5565b60405180910390fd5b613223609960005060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506149be909063ffffffff16565b6132348a8a8a613dad63ffffffff16565b5050505b50505050505050565b6000603460005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505490506132d0565b92915050565b60cc60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561336e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161336590617279565b60405180910390fd5b8060008160000151519050600481101515156133bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016133b6906172bb565b60405180910390fd5b600060d460005060008460200151815260200190815260200160002060005090506000816001016000505414151561342c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134239061708a565b60405180910390fd5b600083604001515190506008811015151561347c576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134739061731e565b60405180910390fd5b608081111515156134c2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016134b990617237565b60405180910390fd5b8360c0015160200151518460c001516000015151141515613518576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161350f906171d4565b60405180910390fd5b8460d46000506000876020015181526020019081526020016000206000506000820151816000016000509080519060200190613555929190614c14565b506020820151816001016000509090556040820151816002016000509080519060200190613584929190614c14565b5060608201518160030160005090805190602001906135a4929190614b7c565b5060808201518160040160005090805190602001906135c4929190614b7c565b5060a08201518160050160005090805190602001906135e4929190614b7c565b5060c0820151816006016000506000820151816000016000509080519060200190613610929190614af2565b506020820151816001016000509080519060200190613630929190614b7c565b50505090505060d56000508560200151908060018154018082558091505060019003906000526020600020900160005b909190919091509090557fc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb856020015186600001516040516136a3929190617418565b60405180910390a15b5050505b505b50565b81600060d4600050600083815260200190815260200160002060005090506000816001016000505414151515613720576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137179061710e565b60405180910390fd5b8260018114806137305750600b81145b1515613771576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137689061712f565b60405180910390fd5b600060cc60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b81526004016137db929190616a74565b60206040518083038186803b1580156137f45760006000fd5b505afa158015613809573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525061382d9190810190615668565b9050670de0b6b3a764000086028114151561387d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161387490617216565b60405180910390fd5b600060d46000506000898152602001908152602001600020600050905060cc60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166329ae69ec338960018c86600401600050876003016000506040518763ffffffff1660e01b815260040161391296959493929190616c11565b600060405180830381600087803b15801561392d5760006000fd5b505af1158015613942573d600060003e3d6000fd5b5050505060cc60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166390596dd133846040518363ffffffff1660e01b81526004016139a9929190616ad6565b600060405180830381600087803b1580156139c45760006000fd5b505af11580156139d9573d600060003e3d6000fd5b505050507fd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b023040338989604051613a1093929190616c81565b60405180910390a15050505b5b50505b505050565b613a2d614c99565b60cc6000506040518060e00160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200150509050613ca1565b90565b600060d360009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638c0cd38d60405180608001604052808881526020018781526020018660ff1681526020018573ffffffffffffffffffffffffffffffffffffffff168152602001506040518263ffffffff1660e01b8152600401613d3d919061737e565b602060405180830381600087803b158015613d585760006000fd5b505af1158015613d6d573d600060003e3d6000fd5b505050506040513d601f19601f82011682018060405250613d919190810190615563565b9050613d98565b949350505050565b6000339050613daa565b90565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515613e1f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e169061729a565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614151515613e91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613e8890617069565b60405180910390fd5b80603460005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92583604051613f7a91906173a1565b60405180910390a35b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515613ffa576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613ff190617258565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161415151561406c576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161406390617006565b60405180910390fd5b61407d8383836149d963ffffffff16565b6140ef816040518060600160405280602681526020016179be60269139603360005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546142489092919063ffffffff16565b603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005081909090555061419281603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050546142ee90919063ffffffff16565b603360005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8360405161423a91906173a1565b60405180910390a35b505050565b60008383111582901515614292576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016142899190616f67565b60405180910390fd5b50828403905061429d565b9392505050565b60006142e46040516142b590616a42565b60405180910390206142cb6149df63ffffffff16565b6142d96149f163ffffffff16565b614a0363ffffffff16565b90506142eb565b90565b60006000828401905083811015151561433c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614333906170cc565b60405180910390fd5b8091505061434656505b92915050565b60cc60005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16638a8d885f8484846040518463ffffffff1660e01b81526004016143b193929190616d8b565b600060405180830381600087803b1580156143cc5760006000fd5b505af11580156143e1573d600060003e3d6000fd5b505050505b505050565b6000816000016000505490506143fc565b919050565b600060d4600050600083815260200190815260200160002060005090506000816006016000506000016000508054905090506000600090505b81811015614564576000836006016000506000016000508281548110151561445e57fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff169050600084600601600050600101600050838154811015156144a957fe5b906000526020600020900160005b5054905060cc60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166387227ab13384846040518463ffffffff1660e01b815260040161452093929190616a9e565b600060405180830381600087803b15801561453b5760006000fd5b505af1158015614550573d600060003e3d6000fd5b5050505050505b808060010191505061443a565b5050505b50565b600060608573ffffffffffffffffffffffffffffffffffffffff166323b872dd8686866040516024016145a093929190616d29565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506040516145ee91906169f2565b6000604051808303816000865af19150503d806000811461462b576040519150601f19603f3d011682016040523d82523d6000602084013e614630565b606091505b509150915081801561465e575060008151148061465d57508080602001905161465c9190810190615538565b5b5b151561469f576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614696906172dc565b60405180910390fd5b50505b50505050565b600060608473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb85856040516024016146db929190616d61565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff838183161783525050505060405161472991906169f2565b6000604051808303816000865af19150503d8060008114614766576040519150601f19603f3d011682016040523d82523d6000602084013e61476b565b606091505b509150915081801561479957506000815114806147985750808060200190516147979190810190615538565b5b5b15156147da576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016147d190617027565b60405180910390fd5b50505b505050565b60006147f26142a463ffffffff16565b82604051602001614804929190616a0a565b604051602081830303815290604052805190602001209050614821565b919050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11151515614890576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161488790617150565b60405180910390fd5b601b8460ff1614806148a55750601c8460ff16145b15156148e6576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148dd906171b3565b60405180910390fd5b60006001868686866040516000815260200160405260405161490b9493929190616f05565b6020604051602081039080840390855afa15801561492e573d600060003e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515156149ac576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149a390616fe5565b60405180910390fd5b809150506149b656505b949350505050565b60018160000160008282825054019250508190909055505b50565b5b505050565b600060656000505490506149ee565b90565b60006066600050549050614a00565b90565b6000838383614a16614a4f63ffffffff16565b30604051602001614a2b959493929190616eb1565b604051602081830303815290604052805190602001209050614a48565b9392505050565b600030504690505b90565b6040518060e00160405280606081526020016000815260200160608152602001606081526020016060815260200160608152602001614a97614d73565b81526020015090565b828054828255906000526020600020908101928215614ae1579160200282015b82811115614ae05782518260005090905591602001919060010190614ac0565b5b509050614aee9190614d90565b5090565b828054828255906000526020600020908101928215614b6b579160200282015b82811115614b6a5782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190614b12565b5b509050614b789190614dbb565b5090565b828054828255906000526020600020908101928215614bbd579160200282015b82811115614bbc5782518260005090905591602001919060010190614b9c565b5b509050614bca9190614d90565b5090565b6040518060e00160405280606081526020016000815260200160608152602001606081526020016060815260200160608152602001614c0b614d73565b81526020015090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614c5557805160ff1916838001178555614c88565b82800160010185558215614c88579182015b82811115614c875782518260005090905591602001919060010190614c67565b5b509050614c959190614d90565b5090565b6040518060e00160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020015090565b604051806040016040528060608152602001606081526020015090565b614db89190614d9a565b80821115614db45760008181506000905550600101614d9a565b5090565b90565b614dff9190614dc5565b80821115614dfb57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff021916905550600101614dc5565b5090565b90566179bc565b600081359050614e15816178ff565b5b92915050565b600082601f8301121515614e305760006000fd5b8135614e43614e3e826174ef565b6174c0565b91508181835260208401935060208101905083856020840282011115614e695760006000fd5b60005b83811015614e9a5781614e7f8882614e06565b8452602084019350602083019250505b600181019050614e6c565b505050505b92915050565b600082601f8301121515614eb95760006000fd5b8135614ecc614ec782617519565b6174c0565b91508181835260208401935060208101905083856020840282011115614ef25760006000fd5b60005b83811015614f235781614f088882615358565b8452602084019350602083019250505b600181019050614ef5565b505050505b92915050565b600082601f8301121515614f425760006000fd5b8135614f55614f5082617543565b6174c0565b91508181835260208401935060208101905083856020840282011115614f7b5760006000fd5b60005b83811015614fac5781614f918882615358565b8452602084019350602083019250505b600181019050614f7e565b505050505b92915050565b600081519050614fc68161791a565b5b92915050565b600081359050614fdc81617935565b5b92915050565b600081519050614ff281617950565b5b92915050565b6000813590506150088161796b565b5b92915050565b600082601f83011215156150235760006000fd5b81356150366150318261756d565b6174c0565b915080825260208301602083018583830111156150535760006000fd5b61505e838284617860565b5050505b92915050565b600060e0828403121561507b5760006000fd5b61508560e06174c0565b9050600082013567ffffffffffffffff8111156150a25760006000fd5b6150ae8482850161500f565b60008301525060206150c284828501615358565b602083015250604082013567ffffffffffffffff8111156150e35760006000fd5b6150ef8482850161500f565b604083015250606082013567ffffffffffffffff8111156151105760006000fd5b61511c84828501614ea5565b606083015250608082013567ffffffffffffffff81111561513d5760006000fd5b61514984828501614ea5565b60808301525060a082013567ffffffffffffffff81111561516a5760006000fd5b61517684828501614ea5565b60a08301525060c082013567ffffffffffffffff8111156151975760006000fd5b6151a384828501615258565b60c0830152505b92915050565b6000608082840312156151c35760006000fd5b6151cd60806174c0565b9050600082013567ffffffffffffffff8111156151ea5760006000fd5b6151f68482850161500f565b600083015250602082013567ffffffffffffffff8111156152175760006000fd5b6152238482850161500f565b602083015250604061523784828501615384565b604083015250606061524b84828501614ff9565b6060830152505b92915050565b60006040828403121561526b5760006000fd5b61527560406174c0565b9050600082013567ffffffffffffffff8111156152925760006000fd5b61529e84828501614e1c565b600083015250602082013567ffffffffffffffff8111156152bf5760006000fd5b6152cb84828501614ea5565b6020830152505b92915050565b6000604082840312156152eb5760006000fd5b6152f560406174c0565b9050600082013567ffffffffffffffff8111156153125760006000fd5b61531e84828501614e1c565b600083015250602082013567ffffffffffffffff81111561533f5760006000fd5b61534b84828501614ea5565b6020830152505b92915050565b60008135905061536781617986565b5b92915050565b60008151905061537d81617986565b5b92915050565b600081359050615393816179a1565b5b92915050565b6000602082840312156153ad5760006000fd5b60006153bb84828501614e06565b9150505b92915050565b60006000604083850312156153da5760006000fd5b60006153e885828601614e06565b92505060206153f985828601614e06565b9150505b9250929050565b6000600060006060848603121561541b5760006000fd5b600061542986828701614e06565b935050602061543a86828701614e06565b925050604061544b86828701615358565b9150505b9250925092565b600060006000600060006000600060e0888a0312156154755760006000fd5b60006154838a828b01614e06565b97505060206154948a828b01614e06565b96505060406154a58a828b01615358565b95505060606154b68a828b01615358565b94505060806154c78a828b01615384565b93505060a06154d88a828b01614fcd565b92505060c06154e98a828b01614fcd565b9150505b92959891949750929550565b600060006040838503121561550e5760006000fd5b600061551c85828601614e06565b925050602061552d85828601615358565b9150505b9250929050565b60006020828403121561554b5760006000fd5b600061555984828501614fb7565b9150505b92915050565b6000602082840312156155765760006000fd5b600061558484828501614fe3565b9150505b92915050565b6000602082840312156155a15760006000fd5b600082013567ffffffffffffffff8111156155bc5760006000fd5b6155c884828501615068565b9150505b92915050565b600060006000606084860312156155e95760006000fd5b600084013567ffffffffffffffff8111156156045760006000fd5b615610868287016151b0565b935050602061562186828701614e06565b925050604061563286828701614e06565b9150505b9250925092565b6000602082840312156156505760006000fd5b600061565e84828501615358565b9150505b92915050565b60006020828403121561567b5760006000fd5b60006156898482850161536e565b9150505b92915050565b60006000604083850312156156a85760006000fd5b60006156b685828601615358565b92505060206156c785828601614e06565b9150505b9250929050565b60006000604083850312156156e75760006000fd5b60006156f585828601615358565b925050602083013567ffffffffffffffff8111156157135760006000fd5b61571f85828601614f2e565b9150505b9250929050565b600060006040838503121561573f5760006000fd5b600061574d85828601615358565b925050602083013567ffffffffffffffff81111561576b5760006000fd5b615777858286016152d8565b9150505b9250929050565b60006000604083850312156157975760006000fd5b60006157a585828601615358565b92505060206157b685828601615358565b9150505b9250929050565b600060006000606084860312156157d85760006000fd5b60006157e686828701615358565b93505060206157f786828701615358565b925050604084013567ffffffffffffffff8111156158155760006000fd5b61582186828701614f2e565b9150505b9250925092565b600061583883836158a3565b6020830190505b92915050565b60006158518383616729565b90505b92915050565b600061586683836169b2565b6020830190505b92915050565b61587c816177c8565b82525b5050565b61588c8161773d565b82525b5050565b61589c8161773d565b82525b5050565b6158ac8161772a565b82525b5050565b6158bc8161772a565b82525b5050565b60006158ce826175f5565b6158d8818561769b565b93506158e38361759b565b8060005b838110156159155781516158fb888261582c565b975061590683617655565b9250505b6001810190506158e7565b508593505050505b92915050565b600061592e82617601565b61593881856176ad565b93508360208202850161594a856175ac565b8060005b8581101561598757848403895281516159678582615845565b945061597283617663565b925060208a019950505b60018101905061594e565b5082975087955050505050505b92915050565b60006159a582617619565b6159af81856176d1565b93506159ba836175ce565b8060005b838110156159ec5781516159d2888261585a565b97506159dd8361767f565b9250505b6001810190506159be565b508593505050505b92915050565b6000615a058261760d565b615a0f81856176bf565b9350615a1a836175bd565b8060005b83811015615a4c578151615a32888261585a565b9750615a3d83617671565b9250505b600181019050615a1e565b508593505050505b92915050565b6000615a6582617625565b615a6f81856176d1565b9350615a7a836175df565b8060005b83811015615ab357615a8f826178cb565b615a99888261585a565b9750615aa48361768d565b9250505b600181019050615a7e565b508593505050505b92915050565b615aca81617750565b82525b5050565b615ada8161775d565b82525b5050565b615af2615aed8261775d565b6178c0565b82525b5050565b6000615b0482617631565b615b0e81856176e3565b9350615b1e818560208601617870565b8084019150505b92915050565b615b34816177db565b82525b5050565b615b4481617801565b82525b5050565b615b5481617827565b82525b5050565b6000615b6682617649565b615b708185617701565b9350615b80818560208601617870565b615b89816178df565b84019150505b92915050565b6000615ba08261763d565b615baa81856176ef565b9350615bba818560208601617870565b615bc3816178df565b84019150505b92915050565b6000615bda8261763d565b615be48185617701565b9350615bf4818560208601617870565b615bfd816178df565b84019150505b92915050565b6000615c16601883617701565b91507f45434453413a20696e76616c6964207369676e6174757265000000000000000060008301526020820190505b919050565b6000615c57602383617701565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f657373000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000615cbe601f83617701565b91507f5472616e7366657248656c7065723a205452414e534645525f4641494c45440060008301526020820190505b919050565b6000615cff602683617701565b91507f426c696e64426f78204572723a43616e206e6f742062652072652d696e69746960008301527f616c697a6564000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000615d66602283617701565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f737300000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000615dcd602083617701565b91507f426c696e64426f78204572723a20426f7820616c72656164792065786973747360008301526020820190505b919050565b6000615e0e602c83617701565b91507f426c696e64426f78204572723a696e76697465722063616e6e6f74206571756160008301527f6c2061646472657373283029000000000000000000000000000000000000000060208301526040820190505b919050565b6000615e75600283617713565b91507f190100000000000000000000000000000000000000000000000000000000000060008301526002820190505b919050565b6000615eb6601b83617701565b91507f536166654d6174683a206164646974696f6e206f766572666c6f77000000000060008301526020820190505b919050565b6000615ef7601d83617701565b91507f45524332305065726d69743a206578706972656420646561646c696e6500000060008301526020820190505b919050565b6000615f38601e83617701565b91507f426c696e64426f78204572723a20736572696573206e6f7420666f756e64000060008301526020820190505b919050565b6000615f79603583617701565b91507f426c696e64426f78204572723a64726177206e756d6265722063616e206f6e6c60008301527f7920626520657175616c20746f2031206f72203131000000000000000000000060208301526040820190505b919050565b6000615fe0602283617701565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616047602983617701565b91507f426c696e64426f78204572723a204f6e6c7920726563656976652035206e667460008301527f20746f6b656e206964000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006160ae602283617701565b91507f426c696e64426f7820204572723a477261646520646f6573206e6f742065786960008301527f737400000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616115602283617701565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061617c603283617701565b91507f426c696e64426f78204572723a2072657761726420746f6b656e206e6f74206560008301527f7175616c2072657761726420616d6f756e74000000000000000000000000000060208301526040820190505b919050565b60006161e3605283617713565b91507f454950373132446f6d61696e28737472696e67206e616d652c737472696e672060008301527f76657273696f6e2c75696e7432353620636861696e49642c616464726573732060208301527f766572696679696e67436f6e747261637429000000000000000000000000000060408301526052820190505b919050565b6000616270601e83617701565b91507f45524332305065726d69743a20696e76616c6964207369676e6174757265000060008301526020820190505b919050565b60006162b1602983617701565b91507f426c696e64426f78204572723a616d6f756e742063616e6e6f74207468616e2060008301527f616c6c6f77616e6365000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616318604783617701565b91507f426c696e64426f78204572723a20496d6167654c696e6b206c656e677468206d60008301527f75737420626520736d616c6c207468616e204d41585f494d4147455f4c494e4b60208301527f5f4c454e4754480000000000000000000000000000000000000000000000000060408301526060820190505b919050565b60006163a5602583617701565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f647265737300000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061640c601a83617701565b91507f426c696e64426f78204572723a20556e617574686f72757a656400000000000060008301526020820190505b919050565b600061644d602483617701565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f726573730000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006164b4603983617701565b91507f426c696e64426f78204572723a206e616d65206c656e677468206d757374206260008301527f65206c657373207468616e204d494e5f4e414d455f4e414d450000000000000060208301526040820190505b919050565b600061651b602483617701565b91507f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f464160008301527f494c45440000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616582603583617701565b91507f426c696e64426f78204572723a64726177206e756d6265722063616e206f6e6c60008301527f7920626520657175616c20746f2031206f72203130000000000000000000000060208301526040820190505b919050565b60006165e9604683617701565b91507f426c696e64426f78204572723a20496d6167654c696e6b206c656e677468206d60008301527f757374206265206c657373207468616e204d494e5f494d4147455f4c494e4b5f60208301527f4c454e475448000000000000000000000000000000000000000000000000000060408301526060820190505b919050565b600060e08301600083015184820360008601526166868282615b95565b915050602083015161669b60208601826169b2565b50604083015184820360408601526166b38282615b95565b915050606083015184820360608601526166cd82826159fa565b915050608083015184820360808601526166e782826159fa565b91505060a083015184820360a086015261670182826159fa565b91505060c083015184820360c086015261671b8282616928565b915050809150505b92915050565b600060e08301600083015184820360008601526167468282615b95565b915050602083015161675b60208601826169b2565b50604083015184820360408601526167738282615b95565b9150506060830151848203606086015261678d82826159fa565b915050608083015184820360808601526167a782826159fa565b91505060a083015184820360a08601526167c182826159fa565b91505060c083015184820360c08601526167db8282616928565b915050809150505b92915050565b60e0820160008201516167ff60008501826158a3565b50602082015161681260208501826158a3565b50604082015161682560408501826158a3565b50606082015161683860608501826158a3565b50608082015161684b6080850182615883565b5060a082015161685e60a08501826158a3565b5060c082015161687160c08501826158a3565b50505b5050565b600060808301600083015184820360008601526168958282615b95565b915050602083015184820360208601526168af8282615b95565b91505060408301516168c460408601826169d2565b5060608301516168d76060860182615b3b565b50809150505b92915050565b6000604083016000830151848203600086015261690082826158c3565b9150506020830151848203602086015261691a82826159fa565b915050809150505b92915050565b6000604083016000830151848203600086015261694582826158c3565b9150506020830151848203602086015261695f82826159fa565b915050809150505b92915050565b6000604083016000830151848203600086015261698a82826158c3565b915050602083015184820360208601526169a482826159fa565b915050809150505b92915050565b6169bb816177af565b82525b5050565b6169cb816177af565b82525b5050565b6169db816177ba565b82525b5050565b6169eb816177ba565b82525b5050565b60006169fe8284615af9565b91508190505b92915050565b6000616a1582615e68565b9150616a218285615ae1565b602082019150616a318284615ae1565b6020820191508190505b9392505050565b6000616a4d826161d6565b91508190505b919050565b6000602082019050616a6d60008301846158b3565b5b92915050565b6000604082019050616a896000830185615873565b616a966020830184615873565b5b9392505050565b6000606082019050616ab36000830186615893565b616ac060208301856158b3565b616acd60408301846169c2565b5b949350505050565b6000604082019050616aeb6000830185615873565b616af860208301846169c2565b5b9392505050565b6000608082019050616b156000830187615873565b616b2260208301866169c2565b616b2f60408301856158b3565b616b3c60608301846169c2565b5b95945050505050565b6000608082019050616b5b6000830187615873565b616b6860208301866169c2565b8181036040830152616b7a8185615a5a565b90508181036060830152616b8e818461599a565b90505b95945050505050565b600060c082019050616baf6000830189615873565b616bbc60208301886169c2565b8181036040830152616bce8187615a5a565b90508181036060830152616be28186615a5a565b9050616bf160808301856169c2565b81810360a0830152616c03818461599a565b90505b979650505050505050565b600060c082019050616c266000830189615873565b616c3360208301886169c2565b616c406040830187615b4b565b616c4d60608301866169c2565b8181036080830152616c5f8185615a5a565b905081810360a0830152616c738184615a5a565b90505b979650505050505050565b6000606082019050616c966000830186615873565b616ca360208301856169c2565b616cb060408301846169c2565b5b949350505050565b600060e082019050616cce600083018a6158b3565b616cdb60208301896158b3565b616ce860408301886158b3565b616cf560608301876158b3565b616d026080830186615893565b616d0f60a08301856158b3565b616d1c60c08301846158b3565b5b98975050505050505050565b6000606082019050616d3e60008301866158b3565b616d4b60208301856158b3565b616d5860408301846169c2565b5b949350505050565b6000604082019050616d7660008301856158b3565b616d8360208301846169c2565b5b9392505050565b6000606082019050616da060008301866158b3565b616dad60208301856169c2565b616dba60408301846169c2565b5b949350505050565b60006040820190508181036000830152616ddd8185615923565b9050616dec60208301846169c2565b5b9392505050565b60006020820190508181036000830152616e0e818461599a565b90505b92915050565b6000602082019050616e2c6000830184615ac1565b5b92915050565b6000602082019050616e486000830184615ad1565b5b92915050565b600060c082019050616e646000830189615ad1565b616e7160208301886158b3565b616e7e60408301876158b3565b616e8b60608301866169c2565b616e9860808301856169c2565b616ea560a08301846169c2565b5b979650505050505050565b600060a082019050616ec66000830188615ad1565b616ed36020830187615ad1565b616ee06040830186615ad1565b616eed60608301856169c2565b616efa60808301846158b3565b5b9695505050505050565b6000608082019050616f1a6000830187615ad1565b616f2760208301866169e2565b616f346040830185615ad1565b616f416060830184615ad1565b5b95945050505050565b6000602082019050616f606000830184615b2b565b5b92915050565b60006020820190508181036000830152616f818184615b5b565b90505b92915050565b60006080820190508181036000830152616fa48187615bcf565b9050616fb360208301866169c2565b8181036040830152616fc58185615bcf565b90508181036060830152616fd9818461696d565b90505b95945050505050565b60006020820190508181036000830152616ffe81615c09565b90505b919050565b6000602082019050818103600083015261701f81615c4a565b90505b919050565b6000602082019050818103600083015261704081615cb1565b90505b919050565b6000602082019050818103600083015261706181615cf2565b90505b919050565b6000602082019050818103600083015261708281615d59565b90505b919050565b600060208201905081810360008301526170a381615dc0565b90505b919050565b600060208201905081810360008301526170c481615e01565b90505b919050565b600060208201905081810360008301526170e581615ea9565b90505b919050565b6000602082019050818103600083015261710681615eea565b90505b919050565b6000602082019050818103600083015261712781615f2b565b90505b919050565b6000602082019050818103600083015261714881615f6c565b90505b919050565b6000602082019050818103600083015261716981615fd3565b90505b919050565b6000602082019050818103600083015261718a8161603a565b90505b919050565b600060208201905081810360008301526171ab816160a1565b90505b919050565b600060208201905081810360008301526171cc81616108565b90505b919050565b600060208201905081810360008301526171ed8161616f565b90505b919050565b6000602082019050818103600083015261720e81616263565b90505b919050565b6000602082019050818103600083015261722f816162a4565b90505b919050565b600060208201905081810360008301526172508161630b565b90505b919050565b6000602082019050818103600083015261727181616398565b90505b919050565b60006020820190508181036000830152617292816163ff565b90505b919050565b600060208201905081810360008301526172b381616440565b90505b919050565b600060208201905081810360008301526172d4816164a7565b90505b919050565b600060208201905081810360008301526172f58161650e565b90505b919050565b6000602082019050818103600083015261731681616575565b90505b919050565b60006020820190508181036000830152617337816165dc565b90505b919050565b600060208201905081810360008301526173598184616669565b90505b92915050565b600060e08201905061737760008301846167e9565b5b92915050565b600060208201905081810360008301526173988184616878565b90505b92915050565b60006020820190506173b660008301846169c2565b5b92915050565b60006040820190506173d260008301856169c2565b6173df6020830184615873565b5b9392505050565b60006040820190506173fc60008301856169c2565b818103602083015261740e818461599a565b90505b9392505050565b600060408201905061742d60008301856169c2565b818103602083015261743f8184615bcf565b90505b9392505050565b600060408201905061745e60008301856169c2565b818103602083015261747081846168e3565b90505b9392505050565b600060408201905061748f60008301856169c2565b61749c60208301846169c2565b5b9392505050565b60006020820190506174b960008301846169e2565b5b92915050565b6000604051905081810181811067ffffffffffffffff821117156174e45760006000fd5b80604052505b919050565b600067ffffffffffffffff8211156175075760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff8211156175315760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff82111561755b5760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff8211156175855760006000fd5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b600081905081600052602060002090505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b6000815490505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006001820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008190505b919050565b60006177358261778e565b90505b919050565b60006177488261778e565b90505b919050565b600081151590505b919050565b60008190505b919050565b60006177738261772a565b90505b919050565b60006177868261772a565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600060ff821690505b919050565b60006177d38261783a565b90505b919050565b60006177e6826177ee565b90505b919050565b60006177f98261778e565b90505b919050565b600061780c82617814565b90505b919050565b600061781f8261778e565b90505b919050565b6000617832826177af565b90505b919050565b60006178458261784d565b90505b919050565b60006178588261778e565b90505b919050565b828183376000838301525b505050565b60005b8381101561788f5780820151818401525b602081019050617873565b8381111561789e576000848401525b505b505050565b60006178b86178b3836178f1565b61771f565b90505b919050565b60008190505b919050565b60006178d782546178a5565b90505b919050565b6000601f19601f83011690505b919050565b60008160001c90505b919050565b6179088161772a565b811415156179165760006000fd5b5b50565b61792381617750565b811415156179315760006000fd5b5b50565b61793e8161775d565b8114151561794c5760006000fd5b5b50565b61795981617768565b811415156179675760006000fd5b5b50565b6179748161777b565b811415156179825760006000fd5b5b50565b61798f816177af565b8114151561799d5760006000fd5b5b50565b6179aa816177ba565b811415156179b85760006000fd5b5b50565bfe45524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220114f250f568722084aa2d9d9099b0f23eda1fd195e8fabf5e224d17d0bf8268964736f6c63430006000033"

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _owner common.Address, _lableAddress common.Address, platform_token common.Address, _controlledTokenBuilder common.Address, _nft common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := abi.JSON(strings.NewReader(StoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(StoreBin), backend, _owner, _lableAddress, platform_token, _controlledTokenBuilder, _nft)
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
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address))
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
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address))
func (_Store *StoreSession) QueryConfig() (BlindBoxConfig, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address))
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
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
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

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft)
func (_Store *StoreSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address flip, address nft)
func (_Store *StoreCallerSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Flip          common.Address
	Nft           common.Address
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
