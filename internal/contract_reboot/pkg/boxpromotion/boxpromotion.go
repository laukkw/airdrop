// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package boxpromotion

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

// BlindBoxPromotionBox is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxPromotionBox struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Level    []*big.Int
	Draw     []*big.Int
	Mix      []*big.Int
	Reward   BlindBoxPromotionReward
}

// BlindBoxPromotionConfig is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxPromotionConfig struct {
	Owner         common.Address
	BoxV1         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Nft           common.Address
	Handing       common.Address
}

// BlindBoxPromotionReward is an auto generated low-level Go binding around an user-defined struct.
type BlindBoxPromotionReward struct {
	Token  []common.Address
	Amount []*big.Int
}

// StoreMetaData contains all meta data concerning the Store contract.
var StoreMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"draw_out\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"mint_box\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"mix_true\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetDraw\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetLevel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"name\":\"resetMix\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"resetOwner\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"indexed\":false,\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"\",\"type\":\"tuple\"}],\"name\":\"resetReward\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_token_ids\",\"type\":\"uint256[]\"}],\"name\":\"Convert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAUL_DECIMAL_PLACES\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DOMAIN_SEPARATOR\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DrawNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_number\",\"type\":\"uint256\"}],\"name\":\"DrawOut\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MAX_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_IMAGE_LINK_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIN_NAME_LENGTH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MIX_TRUE_LOW_LEVEL_NUMBER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBoxPromotion.Box\",\"name\":\"_box\",\"type\":\"tuple\"}],\"name\":\"MintBox\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_grade_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_tokens_id\",\"type\":\"uint256[]\"}],\"name\":\"MixTrue\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryBox\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBoxPromotion.Box\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"start\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"end\",\"type\":\"uint256\"}],\"name\":\"QueryBoxs\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"internalType\":\"uint256[]\",\"name\":\"level\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"draw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"mix\",\"type\":\"uint256[]\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"internalType\":\"structBlindBoxPromotion.Box[]\",\"name\":\"\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QueryConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"boxV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handing\",\"type\":\"address\"}],\"internalType\":\"structBlindBoxPromotion.Config\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryDraws\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryImage\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryLevels\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"}],\"name\":\"QueryMix\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QuerySeriesIds\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QuerySeriesIdsNotNull\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_draw\",\"type\":\"uint256[]\"}],\"name\":\"ResetDraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_drawNumber\",\"type\":\"uint256\"}],\"name\":\"ResetDrawNumber\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"_mix\",\"type\":\"uint256[]\"}],\"name\":\"ResetMix\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"ResetOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_series_id\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"_reward\",\"type\":\"tuple\"}],\"name\":\"ResetReward\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"box_info\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"series_id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"image\",\"type\":\"string\"},{\"components\":[{\"internalType\":\"address[]\",\"name\":\"token\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amount\",\"type\":\"uint256[]\"}],\"internalType\":\"structBlindBoxPromotion.Reward\",\"name\":\"reward\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"boxV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"handing\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_boxV1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_lableAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_key\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_nft\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_handing\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_platform_token\",\"type\":\"address\"}],\"name\":\"initPromotion\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"deadline\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"permit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
	Bin: "0x6080604052600a60cc6000509090553480156200001c5760006000fd5b505b3360cd60005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b6200006c565b6179aa806200007c6000396000f3fe60806040526004361061024a5760003560e01c806368c8ca0d116101395780639fcaf49a116100b6578063d505accf1161007a578063d505accf14610943578063dd62ed3e1461096d578063e9292c47146109ab578063f05a020f146109d5578063f2090893146109ff578063f229c71c14610a2b57610253565b80639fcaf49a14610830578063a457c2d71461085c578063a9059cbb1461089a578063af7a19ed146108d8578063d4ecd0401461091957610253565b80637ecebe00116100fd5780637ecebe0014610734578063868a686d1461077257806387ab1642146107b057806395d89b41146107da5780639c724fa81461080657610253565b806368c8ca0d1461060957806370a08231146106475780637790f6ba1461068557806379502c55146106c35780637d54cf71146106f657610253565b8063332ef5b9116101c75780634521e1811161018b5780634521e1811461052f5780634fe5d7901461055957806350cfeddd146105855780635793cb6b146105b15780636067f59f146105dd57610253565b8063332ef5b9146104455780633644e5151461046f578063395093511461049b5780633eafcaf2146104d957806340b6ddbf1461050557610253565b806323b872dd1161020e57806323b872dd1461035a5780632f40d071146103985780632f575b29146103c257806330157fb9146103ec578063313ce5671461041957610253565b806306fdde0314610259578063095ea7b31461028557806312e7a53d146102c357806312eb1fbf1461030257806318160ddd1461032e57610253565b36610253575b5b005b60006000fd5b3480156102665760006000fd5b5061026f610a69565b60405161027c9190616ffa565b60405180910390f35b3480156102925760006000fd5b506102ad60048036038101906102a891906156d1565b610b13565b6040516102ba9190616ec6565b60405180910390f35b3480156102d05760006000fd5b506102eb60048036038101906102e6919061585a565b610b42565b6040516102f9929190616e41565b60405180910390f35b34801561030f5760006000fd5b5061031861101d565b6040516103259190617475565b60405180910390f35b34801561033b5760006000fd5b50610344611022565b6040516103519190617475565b60405180910390f35b3480156103675760006000fd5b50610382600480360381019061037d91906155dc565b611034565b60405161038f9190616ec6565b60405180910390f35b3480156103a55760006000fd5b506103c060048036038101906103bb9190615524565b611133565b005b3480156103cf5760006000fd5b506103ea60048036038101906103e591906157aa565b61159b565b005b3480156103f95760006000fd5b506104026118a2565b604051610410929190616e95565b60405180910390f35b3480156104265760006000fd5b5061042f6119dd565b60405161043c9190617524565b60405180910390f35b3480156104525760006000fd5b5061046d60048036038101906104689190615802565b6119f9565b005b34801561047c5760006000fd5b50610485611bef565b6040516104929190616ee2565b60405180910390f35b3480156104a85760006000fd5b506104c360048036038101906104be91906156d1565b611c09565b6040516104d09190616ec6565b60405180910390f35b3480156104e65760006000fd5b506104ef611cdc565b6040516104fc9190617475565b60405180910390f35b3480156105125760006000fd5b5061052d600480360381019061052891906157aa565b611ce1565b005b34801561053c5760006000fd5b5061055760048036038101906105529190615899565b612001565b005b3480156105665760006000fd5b5061056f612304565b60405161057c9190616e72565b60405180910390f35b3480156105925760006000fd5b5061059b612367565b6040516105a89190617475565b60405180910390f35b3480156105be5760006000fd5b506105c761236c565b6040516105d49190617475565b60405180910390f35b3480156105ea5760006000fd5b506105f3612375565b6040516106009190617475565b60405180910390f35b3480156106165760006000fd5b50610631600480360381019061062c9190615754565b61237a565b60405161063e9190616ffa565b60405180910390f35b3480156106545760006000fd5b5061066f600480360381019061066a91906154ba565b612440565b60405161067c9190617475565b60405180910390f35b3480156106925760006000fd5b506106ad60048036038101906106a89190615754565b612494565b6040516106ba9190616e72565b60405180910390f35b3480156106d05760006000fd5b506106d9612513565b6040516106ed989796959493929190616dc2565b60405180910390f35b3480156107035760006000fd5b5061071e60048036038101906107199190615754565b61264c565b60405161072b9190616e72565b60405180910390f35b3480156107415760006000fd5b5061075c600480360381019061075791906154ba565b6126cb565b6040516107699190617475565b60405180910390f35b34801561077f5760006000fd5b5061079a60048036038101906107959190615754565b61272e565b6040516107a79190617435565b60405180910390f35b3480156107bd5760006000fd5b506107d860048036038101906107d391906157aa565b612ae3565b005b3480156107e75760006000fd5b506107f0612c1e565b6040516107fd9190616ffa565b60405180910390f35b3480156108135760006000fd5b5061082e600480360381019061082991906154ba565b612cc8565b005b34801561083d5760006000fd5b50610846612de3565b6040516108539190617475565b60405180910390f35b3480156108695760006000fd5b50610884600480360381019061087f91906156d1565b612de8565b6040516108919190616ec6565b60405180910390f35b3480156108a75760006000fd5b506108c260048036038101906108bd91906156d1565b612ed5565b6040516108cf9190616ec6565b60405180910390f35b3480156108e55760006000fd5b5061090060048036038101906108fb9190615754565b612f04565b604051610910949392919061701d565b60405180910390f35b3480156109265760006000fd5b50610941600480360381019061093c9190615754565b61316f565b005b3480156109505760006000fd5b5061096b6004803603810190610966919061562e565b613218565b005b34801561097a5760006000fd5b50610995600480360381019061099091906154e5565b6133f7565b6040516109a29190617475565b60405180910390f35b3480156109b85760006000fd5b506109d360048036038101906109ce9190615710565b61348c565b005b3480156109e25760006000fd5b506109fd60048036038101906109f8919061585a565b613ccb565b005b348015610a0c5760006000fd5b50610a15614091565b604051610a229190617458565b60405180910390f35b348015610a385760006000fd5b50610a536004803603810190610a4e9190615754565b614367565b604051610a609190616e72565b60405180910390f35b606060366000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610b045780601f10610ad957610100808354040283529160200191610b04565b820191906000526020600020905b815481529060010190602001808311610ae757829003601f168201915b50505050509050610b10565b90565b6000610b33610b266143e663ffffffff16565b84846143f363ffffffff16565b60019050610b3c565b92915050565b60606000600060d6600050805490509050600081111580610b6257508385115b80610b6c57508085115b15610b81576060808293509350505061101656505b600084905081851115610b945781905080505b601e8682031115610ba857601e8601905080505b606086820367ffffffffffffffff81118015610bc45760006000fd5b50604051908082528060200260200182016040528015610bfe57816020015b610beb614de7565b815260200190600190039081610be35790505b509050600060008890505b838110156110025760d5600050600060d660005083815481101515610c2a57fe5b906000526020600020900160005b505481526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610cf25780601f10610cc757610100808354040283529160200191610cf2565b820191906000526020600020905b815481529060010190602001808311610cd557829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610da45780601f10610d7957610100808354040283529160200191610da4565b820191906000526020600020905b815481529060010190602001808311610d8757829003601f168201915b5050505050815260200160038201600050805480602002602001604051908101604052809291908181526020018280548015610e0257602002820191906000526020600020905b816000505481526020019060010190808311610deb575b5050505050815260200160048201600050805480602002602001604051908101604052809291908181526020018280548015610e6057602002820191906000526020600020905b816000505481526020019060010190808311610e49575b5050505050815260200160058201600050805480602002602001604051908101604052809291908181526020018280548015610ebe57602002820191906000526020600020905b816000505481526020019060010190808311610ea7575b505050505081526020016006820160005060405180604001604052908160008201600050805480602002602001604051908101604052809291908181526020018280548015610f6257602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311610f18575b5050505050815260200160018201600050805480602002602001604051908101604052809291908181526020018280548015610fc057602002820191906000526020600020905b816000505481526020019060010190808311610fa9575b505050505081526020015050815260200150508383815181101515610fe157fe5b602002602001018190525081806001019250505b8080600101915050610c09565b508184955095505050505061101656505050505b9250929050565b608081565b60006035600050549050611031565b90565b60006110478484846145ce63ffffffff16565b611123846110596143e663ffffffff16565b6111188560405180606001604052806028815260200161792860289139603460005060008b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060006110cb6143e663ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461488e9092919063ffffffff16565b6143f363ffffffff16565b6001905061112c565b9392505050565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff161415156111cb576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016111c2906171a1565b60405180910390fd5b6040518061010001604052808973ffffffffffffffffffffffffffffffffffffffff1681526020018873ffffffffffffffffffffffffffffffffffffffff1681526020018773ffffffffffffffffffffffffffffffffffffffff1681526020018273ffffffffffffffffffffffffffffffffffffffff1681526020018673ffffffffffffffffffffffffffffffffffffffff1681526020018573ffffffffffffffffffffffffffffffffffffffff1681526020018473ffffffffffffffffffffffffffffffffffffffff1681526020018373ffffffffffffffffffffffffffffffffffffffff1681526020015060cd60005060008201518160000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060208201518160010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060408201518160020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060608201518160030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060808201518160040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060a08201518160050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060c08201518160060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555060e08201518160070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555090505060cd60005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663f09a4016848a6040518363ffffffff1660e01b815260040161155b929190616d98565b600060405180830381600087803b1580156115765760006000fd5b505af115801561158b573d600060003e3d6000fd5b505050505b5b5050505050505050565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611633576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161162a906171a1565b60405180910390fd5b81600060d560005060008381526020019081526020016000206000509050600081600101600050541415151561169e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611695906172eb565b60405180910390fd5b600383511415156116e4576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016116db9061713e565b60405180910390fd5b600061176560d5600050600087815260200190815260200160002060005060030160005080548060200260200160405190810160405280929190818152602001828054801561175557602002820191906000526020600020905b81600050548152602001906001019080831161173e575b50505050506148ea63ffffffff16565b90506001810384511415156117af576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016117a6906170ba565b60405180910390fd5b8360d5600050600087815260200190815260200160002060005060050160005090805190602001906117e2929190614e2d565b506000600090505b845181101561185e57620f4240858281518110151561180557fe5b602002602001015111151515611850576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611847906172ca565b60405180910390fd5b5b80806001019150506117ea565b507f7eb88e4a0af247244a661c1c3396b1056fbb526000133aa64223b41ebde62a388585604051611890929190617491565b60405180910390a1505b505b505b5050565b606060006000606060d66000508054905067ffffffffffffffff811180156118ca5760006000fd5b506040519080825280602002602001820160405280156118f95781602001602082028036833780820191505090505b5090506000600090505b60d6600050805490508110156119c957606061194560d66000508381548110151561192a57fe5b906000526020600020900160005b505461249463ffffffff16565b9050600081600160d6600050805490500381518110151561196257fe5b602002602001015111156119ba5760d66000508281548110151561198257fe5b906000526020600020900160005b505483838151811015156119a057fe5b602002602001019090818152602001505083806001019450505b505b8080600101915050611903565b5080829350935050506119d95650505b9091565b6000603860009054906101000a900460ff1690506119f6565b90565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611a91576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611a88906171a1565b60405180910390fd5b81600060d5600050600083815260200190815260200160002060005090506000816001016000505414151515611afc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611af3906172eb565b60405180910390fd5b826020015151836000015151141515611b4a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611b4190617225565b60405180910390fd5b8260d560005060008681526020019081526020016000206000506006016000506000820151816000016000509080519060200190611b89929190614e7f565b506020820151816001016000509080519060200190611ba9929190614e2d565b509050507f13c6c1e3995a6e166e5271e71900b8731fef5e8e3deded10933f3616208e4d9d8484604051611bde9291906174f3565b60405180910390a15b505b505b5050565b6000611bff61494763ffffffff16565b9050611c06565b90565b6000611ccd611c1c6143e663ffffffff16565b84611cc28560346000506000611c366143e663ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008973ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461499190919063ffffffff16565b6143f363ffffffff16565b60019050611cd6565b92915050565b606481565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515611d79576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611d70906171a1565b60405180910390fd5b81600060d5600050600083815260200190815260200160002060005090506000816001016000505414151515611de4576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611ddb906172eb565b60405180910390fd5b60048351141515611e2a576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611e21906170ba565b60405180910390fd5b6000611eab60d56000506000878152602001908152602001600020600050600301600050805480602002602001604051908101604052809291908181526020018280548015611e9b57602002820191906000526020600020905b816000505481526020019060010190808311611e84575b50505050506148ea63ffffffff16565b90506000611ebe856148ea63ffffffff16565b90508181141515611f04576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611efb90617414565b60405180910390fd5b8460d560005060008881526020019081526020016000206000506004016000509080519060200190611f37929190614e2d565b5060006000600090505b8651811015611f74578681815181101515611f5857fe5b60200260200101518201915081505b8080600101915050611f41565b50620f424081141515611fbc576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611fb3906173d2565b60405180910390fd5b7f64bb5486f8cd4c4d6edca404c620289515cbdda97c64db7c2f155d040fd480768787604051611fed929190617491565b60405180910390a15050505b505b505b5050565b82600060d560005060008381526020019081526020016000206000509050600081600101600050541415151561206c576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612063906172eb565b60405180910390fd5b8360018111801561207e575060048111155b15156120bf576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120b69061711d565b60405180910390fd5b858460058151141515612107576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016120fe906173f3565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612177576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161216e90617267565b60405180910390fd5b600586511415156121bd576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016121b490617204565b60405180910390fd5b60cd60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dc84e07d338a8a8a6040518563ffffffff1660e01b81526004016122249493929190616d4b565b600060405180830381600087803b15801561223f5760006000fd5b505af1158015612254573d600060003e3d6000fd5b5050505060cd60005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fa705453336000898b8d6040518663ffffffff1660e01b81526004016122c2959493929190616bdf565b600060405180830381600087803b1580156122dd5760006000fd5b505af11580156122f2573d600060003e3d6000fd5b505050505b5b50505b50505b50505050565b606060d660005080548060200260200160405190810160405280929190818152602001828054801561235857602002820191906000526020600020905b816000505481526020019060010190808311612341575b50505050509050612364565b90565b600481565b60cc6000505481565b600581565b606060d560005060008381526020019081526020016000206000506002016000508054600181600116156101000203166002900480601f01602080910402602001604051908101604052809291908181526020018280546001816001161561010002031660029004801561242f5780601f106124045761010080835404028352916020019161242f565b820191906000526020600020905b81548152906001019060200180831161241257829003601f168201915b5050505050905061243b565b919050565b6000603360005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005054905061248f565b919050565b606060d5600050600083815260200190815260200160002060005060030160005080548060200260200160405190810160405280929190818152602001828054801561250257602002820191906000526020600020905b8160005054815260200190600101908083116124eb575b5050505050905061250e565b919050565b60cd6000508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905088565b606060d560005060008381526020019081526020016000206000506004016000508054806020026020016040519081016040528092919081815260200182805480156126ba57602002820191906000526020600020905b8160005054815260200190600101908083116126a3575b505050505090506126c6565b919050565b6000612722609960005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506149ef909063ffffffff16565b9050612729565b919050565b612736614de7565b60d560005060008381526020019081526020016000206000506040518060e0016040529081600082016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156127f65780601f106127cb576101008083540402835291602001916127f6565b820191906000526020600020905b8154815290600101906020018083116127d957829003601f168201915b5050505050815260200160018201600050548152602001600282016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156128a85780601f1061287d576101008083540402835291602001916128a8565b820191906000526020600020905b81548152906001019060200180831161288b57829003601f168201915b505050505081526020016003820160005080548060200260200160405190810160405280929190818152602001828054801561290657602002820191906000526020600020905b8160005054815260200190600101908083116128ef575b505050505081526020016004820160005080548060200260200160405190810160405280929190818152602001828054801561296457602002820191906000526020600020905b81600050548152602001906001019080831161294d575b50505050508152602001600582016000508054806020026020016040519081016040528092919081815260200182805480156129c257602002820191906000526020600020905b8160005054815260200190600101908083116129ab575b505050505081526020016006820160005060405180604001604052908160008201600050805480602002602001604051908101604052809291908181526020018280548015612a6657602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019060010190808311612a1c575b5050505050815260200160018201600050805480602002602001604051908101604052809291908181526020018280548015612ac457602002820191906000526020600020905b816000505481526020019060010190808311612aad575b505050505081526020015050815260200150509050612ade565b919050565b81600060d5600050600083815260200190815260200160002060005090506000816001016000505414151515612b4e576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612b45906172eb565b60405180910390fd5b60cd60005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663a0c55e31338660d56000506000898152602001908152602001600020600050600301600050876040518563ffffffff1660e01b8152600401612bd39493929190616cbf565b600060405180830381600087803b158015612bee5760006000fd5b505af1158015612c03573d600060003e3d6000fd5b50505050612c1684614a0563ffffffff16565b5b505b505050565b606060376000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612cb95780601f10612c8e57610100808354040283529160200191612cb9565b820191906000526020600020905b815481529060010190602001808311612c9c57829003601f168201915b50505050509050612cc5565b90565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515612d60576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401612d57906171a1565b60405180910390fd5b8060cd60005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055507f73cc802a8666227e3cc846009797bce348c90d4d58fde06a579e87d34d43283a81604051612dd69190616b61565b60405180910390a15b5b50565b600881565b6000612ec6612dfb6143e663ffffffff16565b84612ebb856040518060600160405280602581526020016179506025913960346000506000612e2e6143e663ffffffff16565b73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461488e9092919063ffffffff16565b6143f363ffffffff16565b60019050612ecf565b92915050565b6000612ef5612ee86143e663ffffffff16565b84846145ce63ffffffff16565b60019050612efe565b92915050565b60d5600050602052806000526040600020600091509050806000016000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015612fb65780601f10612f8b57610100808354040283529160200191612fb6565b820191906000526020600020905b815481529060010190602001808311612f9957829003601f168201915b505050505090806001016000505490806002016000508054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156130605780601f1061303557610100808354040283529160200191613060565b820191906000526020600020905b81548152906001019060200180831161304357829003601f168201915b505050505090806006016000506040518060400160405290816000820160005080548060200260200160405190810160405280929190818152602001828054801561310057602002820191906000526020600020905b8160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190600101908083116130b6575b505050505081526020016001820160005080548060200260200160405190810160405280929190818152602001828054801561315e57602002820191906000526020600020905b816000505481526020019060010190808311613147575b505050505081526020015050905084565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613207576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016131fe906171a1565b60405180910390fd5b8060cc6000508190909055505b5b50565b83421115151561325d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161325490617246565b60405180910390fd5b6000609a600050548888886132bd609960005060008e73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000506149ef909063ffffffff16565b896040516020016132d396959493929190616efe565b60405160208183030381529060405280519060200120905060006132fc82614b6f63ffffffff16565b9050600061331282878787614bb363ffffffff16565b90508973ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff16141515613384576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161337b9061734e565b60405180910390fd5b6133d9609960005060008c73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050614d4b909063ffffffff16565b6133ea8a8a8a6143f363ffffffff16565b5050505b50505050505050565b6000603460005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020600050549050613486565b92915050565b60cd60005060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613524576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161351b906171a1565b60405180910390fd5b806000816000015151905060048110151515613575576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161356c90617288565b60405180910390fd5b600060d56000506000846020015181526020019081526020016000206000509050600081600101600050541415156135e2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016135d990617390565b60405180910390fd5b6000836040015151905060088110151515613632576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136299061715f565b60405180910390fd5b60808111151515613678576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161366f906170fc565b60405180910390fd5b8360c0015160200151518460c0015160000151511415156136ce576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016136c590617225565b60405180910390fd5b6000600090505b84606001515181101561374e5760008560600151828151811015156136f657fe5b6020026020010151111515613740576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613737906172ca565b60405180910390fd5b5b80806001019150506136d5565b5060006000600090505b8560800151518110156137f157600086608001518281518110151561377957fe5b60200260200101511115156137c3576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016137ba906172ca565b60405180910390fd5b8560800151818151811015156137d557fe5b60200260200101518201915081505b8080600101915050613758565b50620f424081141515613839576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613830906173d2565b60405180910390fd5b6002856060015151101580156138555750600485606001515111155b1515613896576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161388d906173d2565b60405180910390fd5b8460600151518560800151511480156138bb57506001856060015151038560a0015151145b15156138fc576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016138f39061732d565b60405180910390fd5b600086606001515160050390506000811115613b2c576060600567ffffffffffffffff8111801561392d5760006000fd5b5060405190808252806020026020018201604052801561395c5781602001602082028036833780820191505090505b5090506060600467ffffffffffffffff8111801561397a5760006000fd5b506040519080825280602002602001820160405280156139a95781602001602082028036833780820191505090505b5090506060600467ffffffffffffffff811180156139c75760006000fd5b506040519080825280602002602001820160405280156139f65781602001602082028036833780820191505090505b5090506000600090505b8a6060015151811015613a54578a6060015181815181101515613a1f57fe5b60200260200101518482815181101515613a3557fe5b60200260200101909081815260200150505b8080600101915050613a00565b506000600090505b8a6080015151811015613ab0578a6080015181815181101515613a7b57fe5b60200260200101518382815181101515613a9157fe5b60200260200101909081815260200150505b8080600101915050613a5c565b506000600090505b8a60a0015151811015613b0c578a60a0015181815181101515613ad757fe5b60200260200101518282815181101515613aed57fe5b60200260200101909081815260200150505b8080600101915050613ab8565b50808a60a00181905250818a60800181905250828a606001819052505050505b8660d56000506000896020015181526020019081526020016000206000506000820151816000016000509080519060200190613b69929190614f09565b506020820151816001016000509090556040820151816002016000509080519060200190613b98929190614f09565b506060820151816003016000509080519060200190613bb8929190614e2d565b506080820151816004016000509080519060200190613bd8929190614e2d565b5060a0820151816005016000509080519060200190613bf8929190614e2d565b5060c0820151816006016000506000820151816000016000509080519060200190613c24929190614e7f565b506020820151816001016000509080519060200190613c44929190614e2d565b50505090505060d66000508760200151908060018154018082558091505060019003906000526020600020900160005b909190919091509090557fc935b230e6de5e0792a6500bbc2baa104eca7871b65e1d1e6144c710c1fb9dcb87602001518860000151604051613cb79291906174c2565b60405180910390a1505b505050505b505b50565b81600060d5600050600083815260200190815260200160002060005090506000816001016000505414151515613d36576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d2d906172eb565b60405180910390fd5b826001811480613d4a575060cc6000505481145b1515613d8b576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613d82906170db565b60405180910390fd5b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515613dfb576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613df290617267565b60405180910390fd5b600060cd60005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b8152600401613e65929190616b7d565b60206040518083038186803b158015613e7e5760006000fd5b505afa158015613e93573d600060003e3d6000fd5b505050506040513d601f19601f82011682018060405250810190613eb7919061577f565b9050670de0b6b3a764000086028110151515613f08576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401613eff906171e3565b60405180910390fd5b606060cd60005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663fa70545333898460008d6040518663ffffffff1660e01b8152600401613f74959493929190616c64565b600060405180830381600087803b158015613f8f5760006000fd5b505af1158015613fa4573d600060003e3d6000fd5b5050505060cd60005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634f920f3e33670de0b6b3a76400008a026040518363ffffffff1660e01b8152600401614015929190616c3a565b600060405180830381600087803b1580156140305760006000fd5b505af1158015614045573d600060003e3d6000fd5b505050507fd8a1dd86a30634925097ada2447246c6bf78a10b800d964a19d8bd724b02304033898960405161407c93929190616d13565b60405180910390a15050505b5b50505b505050565b614099614f8e565b60cd600050604051806101000160405290816000820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016001820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016002820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016003820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016004820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016005820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016006820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020016007820160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200150509050614364565b90565b606060d560005060008381526020019081526020016000206000506005016000508054806020026020016040519081016040528092919081815260200182805480156143d557602002820191906000526020600020905b8160005054815260200190600101908083116143be575b505050505090506143e1565b919050565b60003390506143f0565b90565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515614465576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161445c906173b1565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156144d7576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016144ce90617180565b60405180910390fd5b80603460005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200190815260200160002060005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925836040516145c09190617475565b60405180910390a35b505050565b600073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614151515614640576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146379061736f565b60405180910390fd5b600073ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff16141515156146b2576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016146a990617099565b60405180910390fd5b6146c3838383614d6663ffffffff16565b6147358160405180606001604052806026815260200161790260269139603360005060008773ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461488e9092919063ffffffff16565b603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055506147d881603360005060008573ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000505461499190919063ffffffff16565b603360005060008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1681526020019081526020016000206000508190909055508173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef836040516148809190617475565b60405180910390a35b505050565b600083831115829015156148d8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016148cf9190616ffa565b60405180910390fd5b5082840390506148e3565b9392505050565b60006000600090506000600090505b8351811015614937576000848281518110151561491257fe5b602002602001015111156149295781806001019250505b5b80806001019150506148f9565b508091505061494256505b919050565b600061498760405161495890616b4b565b604051809103902061496e614d6c63ffffffff16565b61497c614d7e63ffffffff16565b614d9063ffffffff16565b905061498e565b90565b6000600082840190508381101515156149df576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016149d6906171c2565b60405180910390fd5b809150506149e956505b92915050565b600081600001600050549050614a00565b919050565b600060d5600050600083815260200190815260200160002060005090506000816006016000506000016000508054905090506000600090505b81811015614b685760008360060160005060000160005082815481101515614a6257fe5b906000526020600020900160005b9054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008460060160005060010160005083815481101515614aad57fe5b906000526020600020900160005b5054905060cd60005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff166387227ab13384846040518463ffffffff1660e01b8152600401614b2493929190616ba7565b600060405180830381600087803b158015614b3f5760006000fd5b505af1158015614b54573d600060003e3d6000fd5b5050505050505b8080600101915050614a3e565b5050505b50565b6000614b7f61494763ffffffff16565b82604051602001614b91929190616b13565b604051602081830303815290604052805190602001209050614bae565b919050565b60007f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08260001c11151515614c1d576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c14906172a9565b60405180910390fd5b601b8460ff161480614c325750601c8460ff16145b1515614c73576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614c6a9061730c565b60405180910390fd5b600060018686868660405160008152602001604052604051614c989493929190616fb4565b6020604051602081039080840390855afa158015614cbb573d600060003e3d6000fd5b505050602060405103519050600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff1614151515614d39576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401614d3090617078565b60405180910390fd5b80915050614d4356505b949350505050565b60018160000160008282825054019250508190909055505b50565b5b505050565b60006065600050549050614d7b565b90565b60006066600050549050614d8d565b90565b6000838383614da3614ddc63ffffffff16565b30604051602001614db8959493929190616f60565b604051602081830303815290604052805190602001209050614dd5565b9392505050565b600030504690505b90565b6040518060e00160405280606081526020016000815260200160608152602001606081526020016060815260200160608152602001614e24615086565b81526020015090565b828054828255906000526020600020908101928215614e6e579160200282015b82811115614e6d5782518260005090905591602001919060010190614e4d565b5b509050614e7b91906150a3565b5090565b828054828255906000526020600020908101928215614ef8579160200282015b82811115614ef75782518260006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555091602001919060010190614e9f565b5b509050614f0591906150ce565b5090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10614f4a57805160ff1916838001178555614f7d565b82800160010185558215614f7d579182015b82811115614f7c5782518260005090905591602001919060010190614f5c565b5b509050614f8a91906150a3565b5090565b604051806101000160405280600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff168152602001600073ffffffffffffffffffffffffffffffffffffffff1681526020015090565b604051806040016040528060608152602001606081526020015090565b6150cb91906150ad565b808211156150c757600081815060009055506001016150ad565b5090565b90565b61511291906150d8565b8082111561510e57600081816101000a81549073ffffffffffffffffffffffffffffffffffffffff0219169055506001016150d8565b5090565b9056617900565b60008135905061512881617894565b5b92915050565b600082601f83011215156151435760006000fd5b81356151566151518261756f565b617540565b9150818183526020840193506020810190508385602084028201111561517c5760006000fd5b60005b838110156151ad57816151928882615119565b8452602084019350602083019250505b60018101905061517f565b505050505b92915050565b600082601f83011215156151cc5760006000fd5b81356151df6151da82617599565b617540565b915081818352602084019350602081019050838560208402820111156152055760006000fd5b60005b83811015615236578161521b8882615478565b8452602084019350602083019250505b600181019050615208565b505050505b92915050565b600081359050615250816178af565b5b92915050565b600082601f830112151561526b5760006000fd5b813561527e615279826175c3565b617540565b9150808252602083016020830185838301111561529b5760006000fd5b6152a68382846177f5565b5050505b92915050565b600060e082840312156152c35760006000fd5b6152cd60e0617540565b9050600082013567ffffffffffffffff8111156152ea5760006000fd5b6152f684828501615257565b600083015250602061530a84828501615478565b602083015250604082013567ffffffffffffffff81111561532b5760006000fd5b61533784828501615257565b604083015250606082013567ffffffffffffffff8111156153585760006000fd5b615364848285016151b8565b606083015250608082013567ffffffffffffffff8111156153855760006000fd5b615391848285016151b8565b60808301525060a082013567ffffffffffffffff8111156153b25760006000fd5b6153be848285016151b8565b60a08301525060c082013567ffffffffffffffff8111156153df5760006000fd5b6153eb848285016153f8565b60c0830152505b92915050565b60006040828403121561540b5760006000fd5b6154156040617540565b9050600082013567ffffffffffffffff8111156154325760006000fd5b61543e8482850161512f565b600083015250602082013567ffffffffffffffff81111561545f5760006000fd5b61546b848285016151b8565b6020830152505b92915050565b600081359050615487816178ca565b5b92915050565b60008151905061549d816178ca565b5b92915050565b6000813590506154b3816178e5565b5b92915050565b6000602082840312156154cd5760006000fd5b60006154db84828501615119565b9150505b92915050565b60006000604083850312156154fa5760006000fd5b600061550885828601615119565b925050602061551985828601615119565b9150505b9250929050565b60006000600060006000600060006000610100898b0312156155465760006000fd5b60006155548b828c01615119565b98505060206155658b828c01615119565b97505060406155768b828c01615119565b96505060606155878b828c01615119565b95505060806155988b828c01615119565b94505060a06155a98b828c01615119565b93505060c06155ba8b828c01615119565b92505060e06155cb8b828c01615119565b9150505b9295985092959890939650565b600060006000606084860312156155f35760006000fd5b600061560186828701615119565b935050602061561286828701615119565b925050604061562386828701615478565b9150505b9250925092565b600060006000600060006000600060e0888a03121561564d5760006000fd5b600061565b8a828b01615119565b975050602061566c8a828b01615119565b965050604061567d8a828b01615478565b955050606061568e8a828b01615478565b945050608061569f8a828b016154a4565b93505060a06156b08a828b01615241565b92505060c06156c18a828b01615241565b9150505b92959891949750929550565b60006000604083850312156156e65760006000fd5b60006156f485828601615119565b925050602061570585828601615478565b9150505b9250929050565b6000602082840312156157235760006000fd5b600082013567ffffffffffffffff81111561573e5760006000fd5b61574a848285016152b0565b9150505b92915050565b6000602082840312156157675760006000fd5b600061577584828501615478565b9150505b92915050565b6000602082840312156157925760006000fd5b60006157a08482850161548e565b9150505b92915050565b60006000604083850312156157bf5760006000fd5b60006157cd85828601615478565b925050602083013567ffffffffffffffff8111156157eb5760006000fd5b6157f7858286016151b8565b9150505b9250929050565b60006000604083850312156158175760006000fd5b600061582585828601615478565b925050602083013567ffffffffffffffff8111156158435760006000fd5b61584f858286016153f8565b9150505b9250929050565b600060006040838503121561586f5760006000fd5b600061587d85828601615478565b925050602061588e85828601615478565b9150505b9250929050565b600060006000606084860312156158b05760006000fd5b60006158be86828701615478565b93505060206158cf86828701615478565b925050604084013567ffffffffffffffff8111156158ed5760006000fd5b6158f9868287016151b8565b9150505b9250925092565b6000615910838361597b565b6020830190505b92915050565b60006159298383616836565b90505b92915050565b600061593e8383616ae3565b6020830190505b92915050565b615954816177a9565b82525b5050565b61596481617744565b82525b5050565b61597481617744565b82525b5050565b61598481617731565b82525b5050565b61599481617731565b82525b5050565b60006159a68261763a565b6159b081856176ae565b93506159bb836175f1565b8060005b838110156159ed5781516159d38882615904565b97506159de83617676565b9250505b6001810190506159bf565b508593505050505b92915050565b6000615a0682617646565b615a1081856176c0565b935083602082028501615a2285617602565b8060005b85811015615a5f5784840389528151615a3f858261591d565b9450615a4a83617684565b925060208a019950505b600181019050615a26565b5082975087955050505050505b92915050565b6000615a7d82617652565b615a8781856176d2565b9350615a9283617613565b8060005b83811015615ac4578151615aaa8882615932565b9750615ab583617692565b9250505b600181019050615a96565b508593505050505b92915050565b6000615add82617652565b615ae781856176e4565b9350615af283617613565b8060005b83811015615b24578151615b0a8882615932565b9750615b1583617692565b9250505b600181019050615af6565b508593505050505b92915050565b6000615b3d8261765e565b615b4781856176e4565b9350615b5283617624565b8060005b83811015615b8b57615b6782617860565b615b718882615932565b9750615b7c836176a0565b9250505b600181019050615b56565b508593505050505b92915050565b615ba281617757565b82525b5050565b615bb281617764565b82525b5050565b615bca615bc582617764565b617855565b82525b5050565b615bda816177bc565b82525b5050565b6000615bec8261766a565b615bf681856176f6565b9350615c06818560208601617805565b615c0f81617874565b84019150505b92915050565b6000615c268261766a565b615c308185617708565b9350615c40818560208601617805565b615c4981617874565b84019150505b92915050565b6000615c62601883617708565b91507f45434453413a20696e76616c6964207369676e6174757265000000000000000060008301526020820190505b919050565b6000615ca3602383617708565b91507f45524332303a207472616e7366657220746f20746865207a65726f206164647260008301527f657373000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000615d0a601f83617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a7265736574206572720060008301526020820190505b919050565b6000615d4b603483617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a4f6e6c7920647261772060008301527f74686520737065636966696564206e756d62657200000000000000000000000060208301526040820190505b919050565b6000615db2605083617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a20496d6167654c696e6b60008301527f206c656e677468206d75737420626520736d616c6c207468616e204d41585f4960208301527f4d4147455f4c494e4b5f4c454e4754480000000000000000000000000000000060408301526060820190505b919050565b6000615e3f602b83617708565b91507f426c696e64426f7850726f6d6f74696f6e20204572723a477261646520646f6560008301527f73206e6f7420657869737400000000000000000000000000000000000000000060208301526040820190505b919050565b6000615ea6602083617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a2072657365742065727260008301526020820190505b919050565b6000615ee7604f83617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a20496d6167654c696e6b60008301527f206c656e677468206d757374206265206c657373207468616e204d494e5f494d60208301527f4147455f4c494e4b5f4c454e475448000000000000000000000000000000000060408301526060820190505b919050565b6000615f74602283617708565b91507f45524332303a20617070726f766520746f20746865207a65726f20616464726560008301527f737300000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000615fdb602383617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a20556e617574686f727560008301527f7a6564000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061604260028361771a565b91507f190100000000000000000000000000000000000000000000000000000000000060008301526002820190505b919050565b6000616083601b83617708565b91507f536166654d6174683a206164646974696f6e206f766572666c6f77000000000060008301526020820190505b919050565b60006160c4603283617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a616d6f756e742063616e60008301527f6e6f74207468616e20616c6c6f77616e6365000000000000000000000000000060208301526040820190505b919050565b600061612b602583617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a746f6b656e696473206660008301527f61696c656400000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616192603b83617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a2072657761726420746f60008301527f6b656e206e6f7420657175616c2072657761726420616d6f756e74000000000060208301526040820190505b919050565b60006161f9601d83617708565b91507f45524332305065726d69743a206578706972656420646561646c696e6500000060008301526020820190505b919050565b600061623a602483617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a7265717565737420666160008301527f696c65640000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006162a1604283617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a206e616d65206c656e6760008301527f7468206d757374206265206c657373207468616e204d494e5f4e414d455f4e4160208301527f4d4500000000000000000000000000000000000000000000000000000000000060408301526060820190505b919050565b600061632e602283617708565b91507f45434453413a20696e76616c6964207369676e6174757265202773272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616395602283617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a2072657175657374206560008301527f727200000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006163fc602783617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a20736572696573206e6f60008301527f7420666f756e640000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616463602283617708565b91507f45434453413a20696e76616c6964207369676e6174757265202776272076616c60008301527f756500000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006164ca60528361771a565b91507f454950373132446f6d61696e28737472696e67206e616d652c737472696e672060008301527f76657273696f6e2c75696e7432353620636861696e49642c616464726573732060208301527f766572696679696e67436f6e747261637429000000000000000000000000000060408301526052820190505b919050565b6000616557602283617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a2052657175657374206560008301527f727200000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006165be601e83617708565b91507f45524332305065726d69743a20696e76616c6964207369676e6174757265000060008301526020820190505b919050565b60006165ff602583617708565b91507f45524332303a207472616e736665722066726f6d20746865207a65726f20616460008301527f647265737300000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616666602983617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a20426f7820616c72656160008301527f647920657869737473000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006166cd602483617708565b91507f45524332303a20617070726f76652066726f6d20746865207a65726f2061646460008301527f726573730000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000616734602183617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a7265717565737420657260008301527f720000000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b600061679b603283617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a204f6e6c79207265636560008301527f6976652035206e667420746f6b656e206964000000000000000000000000000060208301526040820190505b919050565b6000616802601e83617708565b91507f426c696e64426f7850726f6d6f74696f6e204572723a6472617720657272000060008301526020820190505b919050565b600060e08301600083015184820360008601526168538282615be1565b91505060208301516168686020860182616ae3565b50604083015184820360408601526168808282615be1565b9150506060830151848203606086015261689a8282615a72565b915050608083015184820360808601526168b48282615a72565b91505060a083015184820360a08601526168ce8282615a72565b91505060c083015184820360c08601526168e88282616a59565b915050809150505b92915050565b600060e08301600083015184820360008601526169138282615be1565b91505060208301516169286020860182616ae3565b50604083015184820360408601526169408282615be1565b9150506060830151848203606086015261695a8282615a72565b915050608083015184820360808601526169748282615a72565b91505060a083015184820360a086015261698e8282615a72565b91505060c083015184820360c08601526169a88282616a59565b915050809150505b92915050565b610100820160008201516169cd600085018261597b565b5060208201516169e0602085018261595b565b5060408201516169f3604085018261597b565b506060820151616a06606085018261597b565b506080820151616a19608085018261597b565b5060a0820151616a2c60a085018261595b565b5060c0820151616a3f60c085018261597b565b5060e0820151616a5260e085018261597b565b50505b5050565b60006040830160008301518482036000860152616a76828261599b565b91505060208301518482036020860152616a908282615a72565b915050809150505b92915050565b60006040830160008301518482036000860152616abb828261599b565b91505060208301518482036020860152616ad58282615a72565b915050809150505b92915050565b616aec81617790565b82525b5050565b616afc81617790565b82525b5050565b616b0c8161779b565b82525b5050565b6000616b1e82616035565b9150616b2a8285615bb9565b602082019150616b3a8284615bb9565b6020820191508190505b9392505050565b6000616b56826164bd565b91508190505b919050565b6000602082019050616b76600083018461598b565b5b92915050565b6000604082019050616b92600083018561594b565b616b9f602083018461594b565b5b9392505050565b6000606082019050616bbc600083018661596b565b616bc9602083018561598b565b616bd66040830184616af3565b5b949350505050565b600060a082019050616bf4600083018861594b565b616c016020830187615bd1565b8181036040830152616c138186615ad2565b9050616c226060830185616af3565b616c2f6080830184616af3565b5b9695505050505050565b6000604082019050616c4f600083018561594b565b616c5c6020830184616af3565b5b9392505050565b600060a082019050616c79600083018861594b565b616c866020830187616af3565b8181036040830152616c988186615ad2565b9050616ca76060830185615bd1565b616cb46080830184616af3565b5b9695505050505050565b6000608082019050616cd4600083018761594b565b616ce16020830186616af3565b8181036040830152616cf38185615b32565b90508181036060830152616d078184615ad2565b90505b95945050505050565b6000606082019050616d28600083018661594b565b616d356020830185616af3565b616d426040830184616af3565b5b949350505050565b6000608082019050616d60600083018761594b565b616d6d6020830186616af3565b616d7a6040830185616af3565b8181036060830152616d8c8184615ad2565b90505b95945050505050565b6000604082019050616dad600083018561598b565b616dba602083018461598b565b5b9392505050565b600061010082019050616dd8600083018b61598b565b616de5602083018a61596b565b616df2604083018961598b565b616dff606083018861598b565b616e0c608083018761598b565b616e1960a083018661596b565b616e2660c083018561598b565b616e3360e083018461598b565b5b9998505050505050505050565b60006040820190508181036000830152616e5b81856159fb565b9050616e6a6020830184616af3565b5b9392505050565b60006020820190508181036000830152616e8c8184615ad2565b90505b92915050565b60006040820190508181036000830152616eaf8185615ad2565b9050616ebe6020830184616af3565b5b9392505050565b6000602082019050616edb6000830184615b99565b5b92915050565b6000602082019050616ef76000830184615ba9565b5b92915050565b600060c082019050616f136000830189615ba9565b616f20602083018861598b565b616f2d604083018761598b565b616f3a6060830186616af3565b616f476080830185616af3565b616f5460a0830184616af3565b5b979650505050505050565b600060a082019050616f756000830188615ba9565b616f826020830187615ba9565b616f8f6040830186615ba9565b616f9c6060830185616af3565b616fa9608083018461598b565b5b9695505050505050565b6000608082019050616fc96000830187615ba9565b616fd66020830186616b03565b616fe36040830185615ba9565b616ff06060830184615ba9565b5b95945050505050565b600060208201905081810360008301526170148184615c1b565b90505b92915050565b600060808201905081810360008301526170378187615c1b565b90506170466020830186616af3565b81810360408301526170588185615c1b565b9050818103606083015261706c8184616a9e565b90505b95945050505050565b6000602082019050818103600083015261709181615c55565b90505b919050565b600060208201905081810360008301526170b281615c96565b90505b919050565b600060208201905081810360008301526170d381615cfd565b90505b919050565b600060208201905081810360008301526170f481615d3e565b90505b919050565b6000602082019050818103600083015261711581615da5565b90505b919050565b6000602082019050818103600083015261713681615e32565b90505b919050565b6000602082019050818103600083015261715781615e99565b90505b919050565b6000602082019050818103600083015261717881615eda565b90505b919050565b6000602082019050818103600083015261719981615f67565b90505b919050565b600060208201905081810360008301526171ba81615fce565b90505b919050565b600060208201905081810360008301526171db81616076565b90505b919050565b600060208201905081810360008301526171fc816160b7565b90505b919050565b6000602082019050818103600083015261721d8161611e565b90505b919050565b6000602082019050818103600083015261723e81616185565b90505b919050565b6000602082019050818103600083015261725f816161ec565b90505b919050565b600060208201905081810360008301526172808161622d565b90505b919050565b600060208201905081810360008301526172a181616294565b90505b919050565b600060208201905081810360008301526172c281616321565b90505b919050565b600060208201905081810360008301526172e381616388565b90505b919050565b60006020820190508181036000830152617304816163ef565b90505b919050565b6000602082019050818103600083015261732581616456565b90505b919050565b600060208201905081810360008301526173468161654a565b90505b919050565b60006020820190508181036000830152617367816165b1565b90505b919050565b60006020820190508181036000830152617388816165f2565b90505b919050565b600060208201905081810360008301526173a981616659565b90505b919050565b600060208201905081810360008301526173ca816166c0565b90505b919050565b600060208201905081810360008301526173eb81616727565b90505b919050565b6000602082019050818103600083015261740c8161678e565b90505b919050565b6000602082019050818103600083015261742d816167f5565b90505b919050565b6000602082019050818103600083015261744f81846168f6565b90505b92915050565b60006101008201905061746e60008301846169b6565b5b92915050565b600060208201905061748a6000830184616af3565b5b92915050565b60006040820190506174a66000830185616af3565b81810360208301526174b88184615ad2565b90505b9392505050565b60006040820190506174d76000830185616af3565b81810360208301526174e98184615c1b565b90505b9392505050565b60006040820190506175086000830185616af3565b818103602083015261751a8184616a9e565b90505b9392505050565b60006020820190506175396000830184616b03565b5b92915050565b6000604051905081810181811067ffffffffffffffff821117156175645760006000fd5b80604052505b919050565b600067ffffffffffffffff8211156175875760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff8211156175b15760006000fd5b6020820290506020810190505b919050565b600067ffffffffffffffff8211156175db5760006000fd5b601f19601f83011690506020810190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b60008190506020820190505b919050565b600081905081600052602060002090505b919050565b6000815190505b919050565b6000815190505b919050565b6000815190505b919050565b6000815490505b919050565b6000815190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006020820190505b919050565b60006001820190505b919050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008282526020820190505b92915050565b60008190505b92915050565b60008190505b919050565b600061773c8261776f565b90505b919050565b600061774f8261776f565b90505b919050565b600081151590505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b600060ff821690505b919050565b60006177b4826177cf565b90505b919050565b60006177c782617790565b90505b919050565b60006177da826177e2565b90505b919050565b60006177ed8261776f565b90505b919050565b828183376000838301525b505050565b60005b838110156178245780820151818401525b602081019050617808565b83811115617833576000848401525b505b505050565b600061784d61784883617886565b617726565b90505b919050565b60008190505b919050565b600061786c825461783a565b90505b919050565b6000601f19601f83011690505b919050565b60008160001c90505b919050565b61789d81617731565b811415156178ab5760006000fd5b5b50565b6178b881617764565b811415156178c65760006000fd5b5b50565b6178d381617790565b811415156178e15760006000fd5b5b50565b6178ee8161779b565b811415156178fc5760006000fd5b5b50565bfe45524332303a207472616e7366657220616d6f756e7420657863656564732062616c616e636545524332303a207472616e7366657220616d6f756e74206578636565647320616c6c6f77616e636545524332303a2064656372656173656420616c6c6f77616e63652062656c6f77207a65726fa2646970667358221220a26cc5b42678b90641c1204a61da073bae5a3c3c17982e4a6a62206f642eef1564736f6c63430006050033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend)
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
func (_Store *StoreCaller) QueryBox(opts *bind.CallOpts, _series_id *big.Int) (BlindBoxPromotionBox, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryBox", _series_id)

	if err != nil {
		return *new(BlindBoxPromotionBox), err
	}

	out0 := *abi.ConvertType(out[0], new(BlindBoxPromotionBox)).(*BlindBoxPromotionBox)

	return out0, err

}

// QueryBox is a free data retrieval call binding the contract method 0x868a686d.
//
// Solidity: function QueryBox(uint256 _series_id) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])))
func (_Store *StoreSession) QueryBox(_series_id *big.Int) (BlindBoxPromotionBox, error) {
	return _Store.Contract.QueryBox(&_Store.CallOpts, _series_id)
}

// QueryBox is a free data retrieval call binding the contract method 0x868a686d.
//
// Solidity: function QueryBox(uint256 _series_id) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])))
func (_Store *StoreCallerSession) QueryBox(_series_id *big.Int) (BlindBoxPromotionBox, error) {
	return _Store.Contract.QueryBox(&_Store.CallOpts, _series_id)
}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreCaller) QueryBoxs(opts *bind.CallOpts, start *big.Int, end *big.Int) ([]BlindBoxPromotionBox, *big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryBoxs", start, end)

	if err != nil {
		return *new([]BlindBoxPromotionBox), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]BlindBoxPromotionBox)).(*[]BlindBoxPromotionBox)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreSession) QueryBoxs(start *big.Int, end *big.Int) ([]BlindBoxPromotionBox, *big.Int, error) {
	return _Store.Contract.QueryBoxs(&_Store.CallOpts, start, end)
}

// QueryBoxs is a free data retrieval call binding the contract method 0x12e7a53d.
//
// Solidity: function QueryBoxs(uint256 start, uint256 end) view returns((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[]))[], uint256)
func (_Store *StoreCallerSession) QueryBoxs(start *big.Int, end *big.Int) ([]BlindBoxPromotionBox, *big.Int, error) {
	return _Store.Contract.QueryBoxs(&_Store.CallOpts, start, end)
}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address))
func (_Store *StoreCaller) QueryConfig(opts *bind.CallOpts) (BlindBoxPromotionConfig, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "QueryConfig")

	if err != nil {
		return *new(BlindBoxPromotionConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(BlindBoxPromotionConfig)).(*BlindBoxPromotionConfig)

	return out0, err

}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address))
func (_Store *StoreSession) QueryConfig() (BlindBoxPromotionConfig, error) {
	return _Store.Contract.QueryConfig(&_Store.CallOpts)
}

// QueryConfig is a free data retrieval call binding the contract method 0xf2090893.
//
// Solidity: function QueryConfig() view returns((address,address,address,address,address,address,address,address))
func (_Store *StoreCallerSession) QueryConfig() (BlindBoxPromotionConfig, error) {
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
	Reward   BlindBoxPromotionReward
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "box_info", arg0)

	outstruct := new(struct {
		Name     string
		SeriesId *big.Int
		Image    string
		Reward   BlindBoxPromotionReward
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Name = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.SeriesId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Image = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.Reward = *abi.ConvertType(out[3], new(BlindBoxPromotionReward)).(*BlindBoxPromotionReward)

	return *outstruct, err

}

// BoxInfo is a free data retrieval call binding the contract method 0xaf7a19ed.
//
// Solidity: function box_info(uint256 ) view returns(string name, uint256 series_id, string image, (address[],uint256[]) reward)
func (_Store *StoreSession) BoxInfo(arg0 *big.Int) (struct {
	Name     string
	SeriesId *big.Int
	Image    string
	Reward   BlindBoxPromotionReward
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
	Reward   BlindBoxPromotionReward
}, error) {
	return _Store.Contract.BoxInfo(&_Store.CallOpts, arg0)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address boxV1, address lable_address, address platform_token, address key_token, address prize_pool, address nft, address handing)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Owner         common.Address
	BoxV1         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Nft           common.Address
	Handing       common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner         common.Address
		BoxV1         common.Address
		LableAddress  common.Address
		PlatformToken common.Address
		KeyToken      common.Address
		PrizePool     common.Address
		Nft           common.Address
		Handing       common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.BoxV1 = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.LableAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.PlatformToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.KeyToken = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.PrizePool = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.Nft = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.Handing = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address boxV1, address lable_address, address platform_token, address key_token, address prize_pool, address nft, address handing)
func (_Store *StoreSession) Config() (struct {
	Owner         common.Address
	BoxV1         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Nft           common.Address
	Handing       common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address boxV1, address lable_address, address platform_token, address key_token, address prize_pool, address nft, address handing)
func (_Store *StoreCallerSession) Config() (struct {
	Owner         common.Address
	BoxV1         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	Nft           common.Address
	Handing       common.Address
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
func (_Store *StoreTransactor) MintBox(opts *bind.TransactOpts, _box BlindBoxPromotionBox) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "MintBox", _box)
}

// MintBox is a paid mutator transaction binding the contract method 0xe9292c47.
//
// Solidity: function MintBox((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])) _box) returns()
func (_Store *StoreSession) MintBox(_box BlindBoxPromotionBox) (*types.Transaction, error) {
	return _Store.Contract.MintBox(&_Store.TransactOpts, _box)
}

// MintBox is a paid mutator transaction binding the contract method 0xe9292c47.
//
// Solidity: function MintBox((string,uint256,string,uint256[],uint256[],uint256[],(address[],uint256[])) _box) returns()
func (_Store *StoreTransactorSession) MintBox(_box BlindBoxPromotionBox) (*types.Transaction, error) {
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

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreTransactor) ResetReward(opts *bind.TransactOpts, _series_id *big.Int, _reward BlindBoxPromotionReward) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "ResetReward", _series_id, _reward)
}

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreSession) ResetReward(_series_id *big.Int, _reward BlindBoxPromotionReward) (*types.Transaction, error) {
	return _Store.Contract.ResetReward(&_Store.TransactOpts, _series_id, _reward)
}

// ResetReward is a paid mutator transaction binding the contract method 0x332ef5b9.
//
// Solidity: function ResetReward(uint256 _series_id, (address[],uint256[]) _reward) returns()
func (_Store *StoreTransactorSession) ResetReward(_series_id *big.Int, _reward BlindBoxPromotionReward) (*types.Transaction, error) {
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

// InitPromotion is a paid mutator transaction binding the contract method 0x2f40d071.
//
// Solidity: function initPromotion(address _owner, address _boxV1, address _lableAddress, address _key, address _prize_pool, address _nft, address _handing, address _platform_token) returns()
func (_Store *StoreTransactor) InitPromotion(opts *bind.TransactOpts, _owner common.Address, _boxV1 common.Address, _lableAddress common.Address, _key common.Address, _prize_pool common.Address, _nft common.Address, _handing common.Address, _platform_token common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "initPromotion", _owner, _boxV1, _lableAddress, _key, _prize_pool, _nft, _handing, _platform_token)
}

// InitPromotion is a paid mutator transaction binding the contract method 0x2f40d071.
//
// Solidity: function initPromotion(address _owner, address _boxV1, address _lableAddress, address _key, address _prize_pool, address _nft, address _handing, address _platform_token) returns()
func (_Store *StoreSession) InitPromotion(_owner common.Address, _boxV1 common.Address, _lableAddress common.Address, _key common.Address, _prize_pool common.Address, _nft common.Address, _handing common.Address, _platform_token common.Address) (*types.Transaction, error) {
	return _Store.Contract.InitPromotion(&_Store.TransactOpts, _owner, _boxV1, _lableAddress, _key, _prize_pool, _nft, _handing, _platform_token)
}

// InitPromotion is a paid mutator transaction binding the contract method 0x2f40d071.
//
// Solidity: function initPromotion(address _owner, address _boxV1, address _lableAddress, address _key, address _prize_pool, address _nft, address _handing, address _platform_token) returns()
func (_Store *StoreTransactorSession) InitPromotion(_owner common.Address, _boxV1 common.Address, _lableAddress common.Address, _key common.Address, _prize_pool common.Address, _nft common.Address, _handing common.Address, _platform_token common.Address) (*types.Transaction, error) {
	return _Store.Contract.InitPromotion(&_Store.TransactOpts, _owner, _boxV1, _lableAddress, _key, _prize_pool, _nft, _handing, _platform_token)
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
	Arg1     BlindBoxPromotionReward
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
