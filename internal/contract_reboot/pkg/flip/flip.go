// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package flip

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
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_lableAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"nft_token\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"_blind_box\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"itokenId\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"name\":\"betAndFlipLog\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"}],\"name\":\"betAndFlip\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"config\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"lable_address\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"platform_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"key_token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"prize_pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"itokenId\",\"type\":\"address\"},{\"internalType\":\"contractnft\",\"name\":\"nft_token\",\"type\":\"address\"},{\"internalType\":\"contractBlindBox\",\"name\":\"blind_box\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockHashUsed\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLastBlockNumberUsed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getResultOfLastFlip\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_key_token\",\"type\":\"address\"}],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastblockhashused\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastblocknumberused\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"lastresult\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6080604052348015620000125760006000fd5b50604051620022073803806200220783398181016040528101906200003891906200037b565b5b33600060005060000160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555085600060005060010160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555084600060005060020160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555083600060005060060160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555082600060005060070160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555080600060005060050160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600060005060040160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055506040518060400160405280600d81526020017f6e6f207761676572732079657400000000000000000000000000000000000000815260200150600860005090805190602001906200027d9291906200028b565b505b5050505050506200049e565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620002ce57805160ff191683800117855562000304565b8280016001018555821562000304579182015b82811115620003035782518260005090905591602001919060010190620002e1565b5b50905062000313919062000317565b5090565b62000343919062000323565b808211156200033f576000818150600090555060010162000323565b5090565b90566200049d565b6000815190506200035c8162000461565b5b92915050565b60008151905062000374816200047f565b5b92915050565b60006000600060006000600060c08789031215620003995760006000fd5b6000620003a989828a016200034b565b9650506020620003bc89828a016200034b565b9550506040620003cf89828a016200034b565b9450506060620003e289828a0162000363565b9350506080620003f589828a016200034b565b92505060a06200040889828a016200034b565b9150505b9295509295509295565b6000620004238262000440565b90505b919050565b6000620004388262000440565b90505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b6200046c8162000416565b811415156200047b5760006000fd5b5b50565b6200048a816200042b565b81141515620004995760006000fd5b5b50565b5b611d5980620004ae6000396000f3fe60806040523480156100115760006000fd5b50600436106100985760003560e01c806394c3fa2e1161006757806394c3fa2e14610119578063afe6ddc014610137578063bd3491f214610155578063cee6f93c14610173578063e3223bcb1461019157610098565b806319ab453c1461009e57806322fc9c62146100ba57806334dbe44d146100d657806379502c55146100f457610098565b60006000fd5b6100b860048036038101906100b3919061122d565b6101af565b005b6100d460048036038101906100cf91906112f2565b610293565b005b6100de610b67565b6040516100eb9190611a92565b60405180910390f35b6100fc610b79565b604051610110989796959493929190611835565b60405180910390f35b610121610cb2565b60405161012e9190611916565b60405180910390f35b61013f610cc4565b60405161014c9190611916565b60405180910390f35b61015d610ccd565b60405161016a9190611965565b60405180910390f35b61017b610d6e565b6040516101889190611965565b60405180910390f35b610199610e18565b6040516101a69190611a92565b60405180910390f35b600073ffffffffffffffffffffffffffffffffffffffff16600060005060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16141515610248576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161023f906119cc565b60405180910390fd5b80600060005060030160006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff1602179055505b50565b3273ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff16141515610303576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016102fa906119ed565b60405180910390fd5b678ac7230489e80000821480610321575068056bc75e2d6310000082145b806103345750683635c9adc5dea0000082145b1515610375576040517f08c379a000000000000000000000000000000000000000000000000000000000815260040161036c90611a2f565b60405180910390fd5b60008114806103845750600181145b15156103c5576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016103bc90611a50565b60405180910390fd5b6000600060005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905060008173ffffffffffffffffffffffffffffffffffffffff1663dd62ed3e33306040518363ffffffff1660e01b815260040161042f929190611763565b60206040518083038186803b1580156104485760006000fd5b505afa15801561045d573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061048191906112c7565b90508381101515156104c8576040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016104bf90611a0e565b60405180910390fd5b60006064605f86028115156104d957fe5b04905061053e600060005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633600060005060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1684610e2163ffffffff16565b6105a2600060005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1633600060005060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16848903610e2163ffffffff16565b6000600060005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630926e497600036336040518463ffffffff1660e01b815260040161060a93929190611932565b602060405180830381600087803b1580156106255760006000fd5b505af115801561063a573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061065e91906112c7565b9050600060028281151561066e57fe5b0690508581141561071b576040518060400160405280600381526020017f77696e0000000000000000000000000000000000000000000000000000000000815260200150600860005090805190602001906106ca929190611098565b50610716600060005060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff163360646002605f8c020281151561070a57fe5b04610f5e63ffffffff16565b610b22565b6040518060400160405280600481526020017f6c6f7373000000000000000000000000000000000000000000000000000000008152602001506008600050908051906020019061076c929190611098565b50678ac7230489e800008714156109d9576060600060005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634fe5d7906040518163ffffffff1660e01b815260040160006040518083038186803b1580156107ee5760006000fd5b505afa158015610803573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f8201168201806040525081019061082c9190611258565b9050600081518481151561083c57fe5b069050600060005060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663dd26c8183360016000868681518110151561089657fe5b6020026020010151600060005060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16630926e497600036336040518463ffffffff1660e01b815260040161090493929190611932565b602060405180830381600087803b15801561091f5760006000fd5b505af1158015610934573d600060003e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061095891906112c7565b6040518663ffffffff1660e01b81526004016109789594939291906117e1565b600060405180830381600087803b1580156109935760006000fd5b505af11580156109a8573d600060003e3d6000fd5b505050506040513d6000823e3d601f19601f820116820180604052508101906109d19190611258565b505050610b21565b68056bc75e2d63100000871415610a8757600060005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634ef969ed3360016040518363ffffffff1660e01b8152600401610a4e9291906117b7565b600060405180830381600087803b158015610a695760006000fd5b505af1158015610a7e573d600060003e3d6000fd5b50505050610b20565b600060005060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634ef969ed33600a6040518363ffffffff1660e01b8152600401610aeb92919061178d565b600060405180830381600087803b158015610b065760006000fd5b505af1158015610b1b573d600060003e3d6000fd5b505050505b5b5b7fe0d0e5eba95f5fe3bea79637fa3b8111d21da191e411be9b9aff888c2be2b75a6008600050604051610b559190611988565b60405180910390a150505050505b5050565b60006009600050549050610b76565b90565b60006000508060000160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060010160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060020160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060030160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060040160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060050160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060060160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16908060070160009054906101000a900473ffffffffffffffffffffffffffffffffffffffff16905088565b6000600a600050549050610cc1565b90565b600a6000505481565b60086000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610d665780601f10610d3b57610100808354040283529160200191610d66565b820191906000526020600020905b815481529060010190602001808311610d4957829003601f168201915b505050505081565b606060086000508054600181600116156101000203166002900480601f016020809104026020016040519081016040528092919081815260200182805460018160011615610100020316600290048015610e095780601f10610dde57610100808354040283529160200191610e09565b820191906000526020600020905b815481529060010190602001808311610dec57829003601f168201915b50505050509050610e15565b90565b60096000505481565b600060608573ffffffffffffffffffffffffffffffffffffffff166323b872dd868686604051602401610e56939291906118b4565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610ea4919061174b565b6000604051808303816000865af19150503d8060008114610ee1576040519150601f19603f3d011682016040523d82523d6000602084013e610ee6565b606091505b5091509150818015610f145750600081511480610f13575080806020019051810190610f12919061129c565b5b5b1515610f55576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401610f4c90611a71565b60405180910390fd5b50505b50505050565b600060608473ffffffffffffffffffffffffffffffffffffffff1663a9059cbb8585604051602401610f919291906118ec565b6040516020818303038152906040529060e01b6020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff8381831617835250505050604051610fdf919061174b565b6000604051808303816000865af19150503d806000811461101c576040519150601f19603f3d011682016040523d82523d6000602084013e611021565b606091505b509150915081801561104f575060008151148061104e57508080602001905181019061104d919061129c565b5b5b1515611090576040517f08c379a0000000000000000000000000000000000000000000000000000000008152600401611087906119ab565b60405180910390fd5b50505b505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106110d957805160ff191683800117855561110c565b8280016001018555821561110c579182015b8281111561110b57825182600050909055916020019190600101906110eb565b5b509050611119919061111d565b5090565b6111459190611127565b808211156111415760008181506000905550600101611127565b5090565b9056611d22565b60008135905061115b81611cd1565b5b92915050565b600082601f83011215156111765760006000fd5b815161118961118482611add565b611aae565b915081818352602084019350602081019050838560208402820111156111af5760006000fd5b60005b838110156111e057816111c58882611217565b8452602084019350602083019250505b6001810190506111b2565b505050505b92915050565b6000815190506111fa81611cec565b5b92915050565b60008135905061121081611d07565b5b92915050565b60008151905061122681611d07565b5b92915050565b6000602082840312156112405760006000fd5b600061124e8482850161114c565b9150505b92915050565b60006020828403121561126b5760006000fd5b600082015167ffffffffffffffff8111156112865760006000fd5b61129284828501611162565b9150505b92915050565b6000602082840312156112af5760006000fd5b60006112bd848285016111eb565b9150505b92915050565b6000602082840312156112da5760006000fd5b60006112e884828501611217565b9150505b92915050565b60006000604083850312156113075760006000fd5b600061131585828601611201565b925050602061132685828601611201565b9150505b9250929050565b61133a81611bbc565b82525b5050565b61134a81611b65565b82525b5050565b61135a81611b85565b82525b5050565b600061136d8385611b35565b935061137a838584611c7a565b61138383611cbf565b840190505b9392505050565b600061139a82611b1d565b6113a48185611b47565b93506113b4818560208601611c8a565b8084019150505b92915050565b6113ca81611bcf565b82525b5050565b6113da81611bf5565b82525b5050565b6113ea81611c1b565b82525b5050565b6113fa81611c2e565b82525b5050565b61140a81611c41565b82525b5050565b600061141c82611b29565b6114268185611b53565b9350611436818560208601611c8a565b61143f81611cbf565b84019150505b92915050565b600081546001811660008114611468576001811461148e576114d3565b607f60028304166114798187611b53565b955060ff1983168652602086019350506114d3565b6002820461149c8187611b53565b95506114a785611b07565b60005b828110156114ca578154818901526001820191505b6020810190506114aa565b80880195505050505b50505b92915050565b60006114e9601f83611b53565b91507f5472616e7366657248656c7065723a205452414e534645525f4641494c45440060008301526020820190505b919050565b600061152a602283611b53565b91507f466c6970204572723a43616e206e6f742062652072652d696e697469616c697a60008301527f656400000000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000611591601783611b53565b91507f466c6970204572723a72657175657374206661696c656400000000000000000060008301526020820190505b919050565b60006115d2602583611b53565b91507f666c6970204572723a616d6f756e742063616e6e6f74207468616e20616c6c6f60008301527f77616e636500000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b6000611639602683611b53565b91507f54686520616d6f756e742073686f756c64206265203130206f7220313030206f60008301527f722031303030000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b60006116a0601683611b53565b91507f706172616d6574657220697320696e636f72726563740000000000000000000060008301526020820190505b919050565b60006116e1602483611b53565b91507f5472616e7366657248656c7065723a205452414e534645525f46524f4d5f464160008301527f494c45440000000000000000000000000000000000000000000000000000000060208301526040820190505b919050565b61174481611bb1565b82525b5050565b6000611757828461138f565b91508190505b92915050565b60006040820190506117786000830185611331565b6117856020830184611341565b5b9392505050565b60006040820190506117a26000830185611331565b6117af60208301846113f1565b5b9392505050565b60006040820190506117cc6000830185611331565b6117d96020830184611401565b5b9392505050565b600060a0820190506117f66000830188611331565b6118036020830187611401565b61181060408301866113e1565b61181d606083018561173b565b61182a608083018461173b565b5b9695505050505050565b60006101008201905061184b600083018b611341565b611858602083018a611341565b6118656040830189611341565b6118726060830188611341565b61187f6080830187611341565b61188c60a0830186611341565b61189960c08301856113d1565b6118a660e08301846113c1565b5b9998505050505050505050565b60006060820190506118c96000830186611341565b6118d66020830185611341565b6118e3604083018461173b565b5b949350505050565b60006040820190506119016000830185611341565b61190e602083018461173b565b5b9392505050565b600060208201905061192b6000830184611351565b5b92915050565b6000604082019050818103600083015261194d818587611361565b905061195c6020830184611331565b5b949350505050565b6000602082019050818103600083015261197f8184611411565b90505b92915050565b600060208201905081810360008301526119a2818461144b565b90505b92915050565b600060208201905081810360008301526119c4816114dc565b90505b919050565b600060208201905081810360008301526119e58161151d565b90505b919050565b60006020820190508181036000830152611a0681611584565b90505b919050565b60006020820190508181036000830152611a27816115c5565b90505b919050565b60006020820190508181036000830152611a488161162c565b90505b919050565b60006020820190508181036000830152611a6981611693565b90505b919050565b60006020820190508181036000830152611a8a816116d4565b90505b919050565b6000602082019050611aa7600083018461173b565b5b92915050565b6000604051905081810181811067ffffffffffffffff82111715611ad25760006000fd5b80604052505b919050565b600067ffffffffffffffff821115611af55760006000fd5b6020820290506020810190505b919050565b600081905081600052602060002090505b919050565b6000815190505b919050565b6000815190505b919050565b60008282526020820190505b92915050565b60008190505b92915050565b60008282526020820190505b92915050565b6000611b7082611b90565b90505b919050565b600081151590505b919050565b60008190505b919050565b600073ffffffffffffffffffffffffffffffffffffffff821690505b919050565b60008190505b919050565b6000611bc782611c54565b90505b919050565b6000611bda82611be2565b90505b919050565b6000611bed82611b90565b90505b919050565b6000611c0082611c08565b90505b919050565b6000611c1382611b90565b90505b919050565b6000611c2682611bb1565b90505b919050565b6000611c3982611bb1565b90505b919050565b6000611c4c82611bb1565b90505b919050565b6000611c5f82611c67565b90505b919050565b6000611c7282611b90565b90505b919050565b828183376000838301525b505050565b60005b83811015611ca95780820151818401525b602081019050611c8d565b83811115611cb8576000848401525b505b505050565b6000601f19601f83011690505b919050565b611cda81611b65565b81141515611ce85760006000fd5b5b50565b611cf581611b78565b81141515611d035760006000fd5b5b50565b611d1081611bb1565b81141515611d1e5760006000fd5b5b50565bfea264697066735822122019c71b80ed6fd671c5b3c01b598e9d392942230d9030845b02ca9b1d729efa3e64736f6c63430006050033",
}

// StoreABI is the input ABI used to generate the binding from.
// Deprecated: Use StoreMetaData.ABI instead.
var StoreABI = StoreMetaData.ABI

// StoreBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StoreMetaData.Bin instead.
var StoreBin = StoreMetaData.Bin

// DeployStore deploys a new Ethereum contract, binding an instance of Store to it.
func DeployStore(auth *bind.TransactOpts, backend bind.ContractBackend, _lableAddress common.Address, platform_token common.Address, nft_token common.Address, _blind_box common.Address, prize_pool common.Address, itokenId common.Address) (common.Address, *types.Transaction, *Store, error) {
	parsed, err := StoreMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StoreBin), backend, _lableAddress, platform_token, nft_token, _blind_box, prize_pool, itokenId)
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

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address itokenId, address nft_token, address blind_box)
func (_Store *StoreCaller) Config(opts *bind.CallOpts) (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	ItokenId      common.Address
	NftToken      common.Address
	BlindBox      common.Address
}, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "config")

	outstruct := new(struct {
		Owner         common.Address
		LableAddress  common.Address
		PlatformToken common.Address
		KeyToken      common.Address
		PrizePool     common.Address
		ItokenId      common.Address
		NftToken      common.Address
		BlindBox      common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Owner = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.LableAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.PlatformToken = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.KeyToken = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.PrizePool = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.ItokenId = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.NftToken = *abi.ConvertType(out[6], new(common.Address)).(*common.Address)
	outstruct.BlindBox = *abi.ConvertType(out[7], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address itokenId, address nft_token, address blind_box)
func (_Store *StoreSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	ItokenId      common.Address
	NftToken      common.Address
	BlindBox      common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// Config is a free data retrieval call binding the contract method 0x79502c55.
//
// Solidity: function config() view returns(address owner, address lable_address, address platform_token, address key_token, address prize_pool, address itokenId, address nft_token, address blind_box)
func (_Store *StoreCallerSession) Config() (struct {
	Owner         common.Address
	LableAddress  common.Address
	PlatformToken common.Address
	KeyToken      common.Address
	PrizePool     common.Address
	ItokenId      common.Address
	NftToken      common.Address
	BlindBox      common.Address
}, error) {
	return _Store.Contract.Config(&_Store.CallOpts)
}

// GetLastBlockHashUsed is a free data retrieval call binding the contract method 0x94c3fa2e.
//
// Solidity: function getLastBlockHashUsed() view returns(bytes32)
func (_Store *StoreCaller) GetLastBlockHashUsed(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getLastBlockHashUsed")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetLastBlockHashUsed is a free data retrieval call binding the contract method 0x94c3fa2e.
//
// Solidity: function getLastBlockHashUsed() view returns(bytes32)
func (_Store *StoreSession) GetLastBlockHashUsed() ([32]byte, error) {
	return _Store.Contract.GetLastBlockHashUsed(&_Store.CallOpts)
}

// GetLastBlockHashUsed is a free data retrieval call binding the contract method 0x94c3fa2e.
//
// Solidity: function getLastBlockHashUsed() view returns(bytes32)
func (_Store *StoreCallerSession) GetLastBlockHashUsed() ([32]byte, error) {
	return _Store.Contract.GetLastBlockHashUsed(&_Store.CallOpts)
}

// GetLastBlockNumberUsed is a free data retrieval call binding the contract method 0x34dbe44d.
//
// Solidity: function getLastBlockNumberUsed() view returns(uint256)
func (_Store *StoreCaller) GetLastBlockNumberUsed(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getLastBlockNumberUsed")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetLastBlockNumberUsed is a free data retrieval call binding the contract method 0x34dbe44d.
//
// Solidity: function getLastBlockNumberUsed() view returns(uint256)
func (_Store *StoreSession) GetLastBlockNumberUsed() (*big.Int, error) {
	return _Store.Contract.GetLastBlockNumberUsed(&_Store.CallOpts)
}

// GetLastBlockNumberUsed is a free data retrieval call binding the contract method 0x34dbe44d.
//
// Solidity: function getLastBlockNumberUsed() view returns(uint256)
func (_Store *StoreCallerSession) GetLastBlockNumberUsed() (*big.Int, error) {
	return _Store.Contract.GetLastBlockNumberUsed(&_Store.CallOpts)
}

// GetResultOfLastFlip is a free data retrieval call binding the contract method 0xcee6f93c.
//
// Solidity: function getResultOfLastFlip() view returns(string)
func (_Store *StoreCaller) GetResultOfLastFlip(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Store.contract.Call(opts, &out, "getResultOfLastFlip")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetResultOfLastFlip is a free data retrieval call binding the contract method 0xcee6f93c.
//
// Solidity: function getResultOfLastFlip() view returns(string)
func (_Store *StoreSession) GetResultOfLastFlip() (string, error) {
	return _Store.Contract.GetResultOfLastFlip(&_Store.CallOpts)
}

// GetResultOfLastFlip is a free data retrieval call binding the contract method 0xcee6f93c.
//
// Solidity: function getResultOfLastFlip() view returns(string)
func (_Store *StoreCallerSession) GetResultOfLastFlip() (string, error) {
	return _Store.Contract.GetResultOfLastFlip(&_Store.CallOpts)
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

// BetAndFlip is a paid mutator transaction binding the contract method 0x22fc9c62.
//
// Solidity: function betAndFlip(uint256 value, uint256 num) returns()
func (_Store *StoreTransactor) BetAndFlip(opts *bind.TransactOpts, value *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "betAndFlip", value, num)
}

// BetAndFlip is a paid mutator transaction binding the contract method 0x22fc9c62.
//
// Solidity: function betAndFlip(uint256 value, uint256 num) returns()
func (_Store *StoreSession) BetAndFlip(value *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BetAndFlip(&_Store.TransactOpts, value, num)
}

// BetAndFlip is a paid mutator transaction binding the contract method 0x22fc9c62.
//
// Solidity: function betAndFlip(uint256 value, uint256 num) returns()
func (_Store *StoreTransactorSession) BetAndFlip(value *big.Int, num *big.Int) (*types.Transaction, error) {
	return _Store.Contract.BetAndFlip(&_Store.TransactOpts, value, num)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _key_token) returns()
func (_Store *StoreTransactor) Init(opts *bind.TransactOpts, _key_token common.Address) (*types.Transaction, error) {
	return _Store.contract.Transact(opts, "init", _key_token)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _key_token) returns()
func (_Store *StoreSession) Init(_key_token common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, _key_token)
}

// Init is a paid mutator transaction binding the contract method 0x19ab453c.
//
// Solidity: function init(address _key_token) returns()
func (_Store *StoreTransactorSession) Init(_key_token common.Address) (*types.Transaction, error) {
	return _Store.Contract.Init(&_Store.TransactOpts, _key_token)
}

// StoreBetAndFlipLogIterator is returned from FilterBetAndFlipLog and is used to iterate over the raw logs and unpacked data for BetAndFlipLog events raised by the Store contract.
type StoreBetAndFlipLogIterator struct {
	Event *StoreBetAndFlipLog // Event containing the contract specifics and raw log

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
func (it *StoreBetAndFlipLogIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StoreBetAndFlipLog)
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
		it.Event = new(StoreBetAndFlipLog)
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
func (it *StoreBetAndFlipLogIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StoreBetAndFlipLogIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StoreBetAndFlipLog represents a BetAndFlipLog event raised by the Store contract.
type StoreBetAndFlipLog struct {
	Arg0 string
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterBetAndFlipLog is a free log retrieval operation binding the contract event 0xe0d0e5eba95f5fe3bea79637fa3b8111d21da191e411be9b9aff888c2be2b75a.
//
// Solidity: event betAndFlipLog(string arg0)
func (_Store *StoreFilterer) FilterBetAndFlipLog(opts *bind.FilterOpts) (*StoreBetAndFlipLogIterator, error) {

	logs, sub, err := _Store.contract.FilterLogs(opts, "betAndFlipLog")
	if err != nil {
		return nil, err
	}
	return &StoreBetAndFlipLogIterator{contract: _Store.contract, event: "betAndFlipLog", logs: logs, sub: sub}, nil
}

// WatchBetAndFlipLog is a free log subscription operation binding the contract event 0xe0d0e5eba95f5fe3bea79637fa3b8111d21da191e411be9b9aff888c2be2b75a.
//
// Solidity: event betAndFlipLog(string arg0)
func (_Store *StoreFilterer) WatchBetAndFlipLog(opts *bind.WatchOpts, sink chan<- *StoreBetAndFlipLog) (event.Subscription, error) {

	logs, sub, err := _Store.contract.WatchLogs(opts, "betAndFlipLog")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StoreBetAndFlipLog)
				if err := _Store.contract.UnpackLog(event, "betAndFlipLog", log); err != nil {
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

// ParseBetAndFlipLog is a log parse operation binding the contract event 0xe0d0e5eba95f5fe3bea79637fa3b8111d21da191e411be9b9aff888c2be2b75a.
//
// Solidity: event betAndFlipLog(string arg0)
func (_Store *StoreFilterer) ParseBetAndFlipLog(log types.Log) (*StoreBetAndFlipLog, error) {
	event := new(StoreBetAndFlipLog)
	if err := _Store.contract.UnpackLog(event, "betAndFlipLog", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
