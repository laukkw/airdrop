package contract_reboot

import (
	"context"
	"errors"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/rzry/airdrop/internal/contract_reboot/fflag"
	"github.com/rzry/airdrop/internal/contract_reboot/options"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/blindbox"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/boxpromotion"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/buildtoken"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/token"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"sync/atomic"
	"time"
)

var sum int64
func InitOne(opt *options.Options,ctx context.Context){
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	//初始化一期合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initBoxTx, err := opt.Ier.Box.Init(auth, blindbox.BuildTokenControlledTokenConfig{
		Name:       "KeyToken",
		Symbol:     "BoxKey",
		Decimals:   18,
		Controller: common.HexToAddress(*fflag.BoxAddr),
	}, common.HexToAddress(*fflag.FlipAddr), common.HexToAddress(*fflag.PoolAddr))
	if err != nil {
		log.Error("init box err ", zap.Error(err))
		return
	}
	log.Info("初始化一期,大盲盒init结果:", zap.Any("hash", initBoxTx.Hash()))

	//获取key token
	block, err := opt.Client.BlockNumber(ctx)
	time.Sleep(10 * time.Second)
	ktoken, err := fetchK(big.NewInt(int64(block)), opt)
	if err != nil {
		log.Error("get key token err", zap.Error(err))
		return
	}
	fflag.KTokenAddr = ktoken.String()
	log.Info("key token 地址为:", zap.Any("address", ktoken))

	///初始化nft 合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initNft, err := opt.Ier.Nft.Init(auth,
		common.HexToAddress(*fflag.BoxAddr),  //box v1
		common.HexToAddress(*fflag.FlipAddr), //flip
		common.HexToAddress(*fflag.LockAddr)) //lock
	if err != nil {
		log.Error("init nft err", zap.Error(err))
		return
	}
	log.Info("初始化nft合约结果为: ", zap.Any("hash", initNft.Hash()))
}

func AddBoxPro(opt *options.Options){
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	//向一期合约添加二期合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	addTxPro, err := opt.Ier.Box.AddBoxPro(auth, common.HexToAddress(*fflag.BoxProAddr),
		common.HexToAddress(*fflag.HandAddr))
	if err != nil {
		log.Error("addTxPro err", zap.Error(err))
		return
	}
	log.Info("addTxPro ", zap.Any("hash", addTxPro.Hash()))

	// 4 . 初始化box pro 合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initBoxPro, err := opt.Ier.BoxPro.InitPromotion(auth,
		common.HexToAddress(*fflag.Owner),      //owner
		common.HexToAddress(*fflag.BoxAddr),    //box v1
		common.HexToAddress(*fflag.LableAddr),  //lable address
		common.HexToAddress(fflag.KTokenAddr),  // key
		common.HexToAddress(*fflag.PoolAddr),   //prize pool
		common.HexToAddress(*fflag.NftAddr),    //nft
		common.HexToAddress(*fflag.HandAddr),   //handing
		common.HexToAddress(*fflag.PTokenAddr))  //ptoken
	if err != nil {
		log.Error("init Box pro err", zap.Error(err))
		return
	}
	log.Info("init Box pro", zap.Any("hash", initBoxPro.Hash()))
}


func InitContract(opt *options.Options, ctx context.Context) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	// 1 . 创建 p token - usd 的lp token

	// 2 调用gov 来初始化

	// 3 调用box 来初始化
	nonce := atomic.AddInt64(&sum, 1)
	auth.Nonce = big.NewInt(nonce)
	initBoxTx, err := opt.Ier.Box.Init(auth, blindbox.BuildTokenControlledTokenConfig{
		Name:       "KeyToken",
		Symbol:     "BoxKey",
		Decimals:   18,
		Controller: common.HexToAddress(*fflag.BoxAddr),
	}, common.HexToAddress(*fflag.FlipAddr), common.HexToAddress(*fflag.PoolAddr))
	if err != nil {
		log.Error("init box err ", zap.Error(err))
		return
	}
	log.Info("初始化box结果:", zap.Any("hash", initBoxTx.Hash()))
	block, err := opt.Client.BlockNumber(ctx)
	time.Sleep(10 * time.Second)

	ktoken, err := fetchK(big.NewInt(int64(block)), opt)
	if err != nil {
		log.Error("get key token err", zap.Error(err))
		return
	}
	fflag.KTokenAddr = ktoken.String()
	log.Info("key token addr ", zap.Any("address", ktoken))


	// 5 . 初始化nft 合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initNft, err := opt.Ier.Nft.Init(auth,
		common.HexToAddress(*fflag.BoxAddr),  //box v1
		common.HexToAddress(*fflag.FlipAddr), //flip
		common.HexToAddress(*fflag.LockAddr)) //lock
	if err != nil {
		log.Error("init nft err", zap.Error(err))
		return
	}
	log.Info("init nft ", zap.Any("hash", initNft.Hash()))

}
func AddPro(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	// 添加二期合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	addTxPro, err := opt.Ier.Box.AddBoxPro(auth, common.HexToAddress(*fflag.BoxProAddr),
		common.HexToAddress(*fflag.HandAddr))
	if err != nil {
		log.Error("addTxPro err", zap.Error(err))
		return
	}
	log.Info("addTxPro ", zap.Any("hash", addTxPro.Hash()))

	// 4 . 初始化box pro 合约
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initBoxPro, err := opt.Ier.BoxPro.InitPromotion(auth,
		common.HexToAddress(*fflag.Owner),      //owner
		common.HexToAddress(*fflag.BoxAddr),    //box v1
		common.HexToAddress(*fflag.LableAddr),  //lable address
		common.HexToAddress(fflag.KTokenAddr),  // key
		common.HexToAddress(*fflag.PoolAddr),   //prize pool
		common.HexToAddress(*fflag.NftAddr),    //nft
		common.HexToAddress(*fflag.HandAddr),   //handing
		common.HexToAddress(*fflag.PTokenAddr))  //rand
	if err != nil {
		log.Error("init Box pro err", zap.Error(err))
		return
	}
	log.Info("初始化二期盲盒 结果:", zap.Any("hash", initBoxPro.Hash()))
}
func CreateSeriesV2(opt *options.Options, series int64) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	mintBoxHash, err := opt.Ier.BoxPro.MintBox(auth, boxpromotion.BlindBoxPromotionBox{
		Name:     "欢天喜地七仙女",
		SeriesId: big.NewInt(series),
		Image:    "https://rzrysayhello.cn.com",
		Level:    array([]string{"1", "3", "6"}),
		Draw:     array([]string{"100000", "380000", "520000"}),
		Mix:      array([]string{"0", "0"}),
		Reward: boxpromotion.BlindBoxPromotionReward{
			Token:  arraddr([]string{"0x1c782CEA65aD7c00DBA103Ea7d859822613854b6"}),
			Amount: array([]string{pkg.ToEthers("100", 18).String()}),
		},
	})
	if err != nil {
		log.Error("mint box err", zap.Error(err))
		return
	}
	log.Info("创建二期系列结果: ", zap.Any("hash", mintBoxHash.Hash()))
}
func TestDraw(opt *options.Options,series int64){
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	log.Info("抽取10次")
	for i := 0 ; i < 1; i++{
		auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
		recordHx , err := opt.Ier.Handing.XATEST(auth,common.HexToAddress("0x550BF94f6A3f8344B0e4866112a0bBF839B2637d"),big.NewInt(10),big.NewInt(series))
		if err != nil{
			log.Error("test err",zap.Error(err))
			return
		}
		log.Info("抽取二期合约结果:",zap.Any("hash",recordHx.Hash()))
	}
}

func SendAward(opt *options.Options)  {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	linktoken, err := token.NewStore(common.HexToAddress("0x84b9b910527ad5c03a9ca831909e21e236ea7b06"), opt.Client)
	linktoken.Approve(auth,common.HexToAddress(*fflag.HandAddr),pkg.ToEthers("1", 50).BigInt())

	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	awardHx,err := opt.Ier.Handing.Award(auth)
	if err != nil{
		log.Error("send award err",zap.Error(err))
		return
	}
	log.Info("发放奖励",zap.Any("hash",awardHx.Hash()))
}

func QueryHanding(opt *options.Options){
	/*res,err := opt.Ier.Handing.QueryPlayerTokens(nil,
		common.HexToAddress("0x550BF94f6A3f8344B0e4866112a0bBF839B2637d"),big.NewInt(0),big.NewInt(3))
	if err != nil{
		log.Error("query err",zap.Error(err))
		return
	}
	for _,v := range res.Result{
		log.Info("查询到结果为",zap.Any("res",v))
	}
	log.Info("总数为",zap.Any("res",res.Len))
*/
	cycle,err := opt.Ier.Handing.CYCLE(nil)
	if err != nil{
		log.Error("查询错误",zap.Error(err))
		return
	}
	log.Info("查询到的周期为",zap.Any("cycle",cycle))

	cycleRng ,err := opt.Ier.Handing.CycleRng(nil,big.NewInt(1))
	if err != nil{
		log.Error("查询错误",zap.Error(err))
		return
	}
	log.Info("查询到周期随机数为",zap.Any("rng",cycleRng))


}

func Received(opt *options.Options){
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	recHx,err := opt.Ier.Handing.Received(auth,[]*big.Int{big.NewInt(2182),big.NewInt(2183),big.NewInt(2184)},
	[]*big.Int{big.NewInt(1),big.NewInt(1),big.NewInt(1)})
	if err != nil{
		log.Error("received err",zap.Error(err))
		return
	}
	log.Info("领取结果",zap.Any("received",recHx.Hash()))
}

func CreateSeries(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	mintBoxHash, err := opt.Ier.Box.MintBox(auth, blindbox.BlindBoxBox{
		Name:     "欢天喜地七仙女",
		SeriesId: big.NewInt(1),
		Image:    "https://rzrysayhello.cn.com",
		Level:    array([]string{"1", "3", "6", "10", "15"}),
		Draw:     array([]string{"100000", "180000", "200000", "520000"}),
		Mix:      array([]string{"0", "0", "0", "0"}),
		Reward: blindbox.BlindBoxReward{
			Token:  arraddr([]string{"0x1c782CEA65aD7c00DBA103Ea7d859822613854b6"}),
			Amount: array([]string{pkg.ToEthers("100", 18).String()}),
		},
	})
	if err != nil {
		log.Error("mint box err", zap.Error(err))
		return
	}
	log.Info("mint Box ", zap.Any("hash", mintBoxHash.Hash()))
}
func FethchKey(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	allowance, err := opt.Ier.Ptoken.Allowance(nil, pkg.PrivateToAddress(*fflag.Private), common.HexToAddress(*fflag.BoxAddr))
	if err != nil {
		log.Error("get allowance err", zap.Error(err))
		return
	}
	if !decimal.NewFromBigInt(allowance, 0).GreaterThan(pkg.ToEthers("1", 30)) {
		auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
		if _, err := opt.Ier.Ptoken.Approve(auth, common.HexToAddress(*fflag.BoxAddr), pkg.ToEthers("1", 50).BigInt()); err != nil {
			log.Error("approve err", zap.Error(err))
			return
		}
	}
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	drawTx, err := opt.Ier.Box.Draw(auth, big.NewInt(10), common.HexToAddress(*fflag.LableAddr))
	if err != nil {
		log.Error("draw tx hash err", zap.Error(err))
		return
	}
	log.Info("draw key ", zap.Any("hash", drawTx.Hash()))
}
func FethchNft(opt *options.Options) {
	//先查询 k token 地址
	addktoken(opt)
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	ktoken, err := token.NewStore(common.HexToAddress(fflag.KTokenAddr), opt.Client)
	allowance, err := ktoken.Allowance(nil, pkg.PrivateToAddress(*fflag.Private), common.HexToAddress(*fflag.BoxAddr))
	if err != nil {
		log.Error("get allowance err", zap.Error(err))
		return
	}
	if !decimal.NewFromBigInt(allowance, 0).GreaterThan(pkg.ToEthers("1", 30)) {
		auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
		if _, err = ktoken.Approve(auth, common.HexToAddress(*fflag.BoxAddr), pkg.ToEthers("1", 50).BigInt()); err != nil {
			log.Error("approve err", zap.Error(err))
			return
		}
	}
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	drawTx, err := opt.Ier.Box.DrawOut(auth, big.NewInt(1), big.NewInt(1))
	if err != nil {
		log.Error("get nft tx hash err", zap.Error(err))
		return
	}
	log.Info("get nft ", zap.Any("hash", drawTx.Hash()))
}
func FethchNftV2(opt *options.Options) {
	//先查询 k token 地址
	addktoken(opt)
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	ktoken, err := token.NewStore(common.HexToAddress(fflag.KTokenAddr), opt.Client)
	allowance, err := ktoken.Allowance(nil, pkg.PrivateToAddress(*fflag.Private), common.HexToAddress(*fflag.BoxProAddr))
	if err != nil {
		log.Error("get allowance err", zap.Error(err))
		return
	}
	if !decimal.NewFromBigInt(allowance, 0).GreaterThan(pkg.ToEthers("1", 30)) {
		auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
		if _, err = ktoken.Approve(auth, common.HexToAddress(*fflag.BoxProAddr), pkg.ToEthers("1", 50).BigInt()); err != nil {
			log.Error("approve err", zap.Error(err))
			return
		}
	}
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	drawTx, err := opt.Ier.BoxPro.DrawOut(auth, big.NewInt(1), big.NewInt(1))
	if err != nil {
		log.Error("get nft tx hash err", zap.Error(err))
		return
	}
	log.Info("get nft ", zap.Any("hash", drawTx.Hash()))
}

func Mix(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	mixTx, err := opt.Ier.BoxPro.MixTrue(auth, big.NewInt(1), big.NewInt(3),
		[]*big.Int{big.NewInt(31), big.NewInt(33),
			big.NewInt(37), big.NewInt(40), big.NewInt(1)})
	if err != nil {
		log.Error("mix nft tx hash err", zap.Error(err))
		return
	}
	log.Info("get nft ", zap.Any("hash", mixTx.Hash()))
}





func resetBoxProDrawNum(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	drawNumber, err := opt.Ier.BoxPro.DrawNumber(nil)
	if err != nil {
		log.Error("query draw number err", zap.Error(err))
		return
	}
	log.Info("drawNumber ", zap.Any("drawNumber", drawNumber))
	tx, err := opt.Ier.BoxPro.ResetDrawNumber(auth, big.NewInt(10))
	if err != nil {
		log.Error("ResetDrawNumber err", zap.Error(err))
		return
	}
	log.Info("reset draw number hash ", zap.Any("hash", tx.Hash()))
	time.Sleep(6 * time.Second)
	drawNumbernew, err := opt.Ier.BoxPro.DrawNumber(nil)
	if err != nil {
		log.Error("query draw number err", zap.Error(err))
		return
	}
	log.Info("drawNumber ", zap.Any("drawNumber", drawNumbernew))
}

func locked(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	if _, err := opt.Ier.Ptoken.Approve(auth, common.HexToAddress(*fflag.LockAddr), pkg.ToEthers("20000", 18).BigInt()); err != nil {
		log.Error("approve err", zap.Error(err))
		return
	}
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	stakingTx, err := opt.Ier.Lock.Stake(auth, pkg.ToEthers("20000", 18).BigInt())
	if err != nil {
		log.Error("staking err", zap.Error(err))
		return
	}
	log.Info("staking ", zap.Any("staking", stakingTx.Hash()))
}


func fetchOther(opt *options.Options) {

	/*balance ,err :=opt.Ier.Nft.BalanceOf(nil,common.HexToAddress(*fflag.HandAddr))
	if err != nil{
		log.Error("get balance of err ",zap.Error(err))
		return
	}
	log.Info("balcnce of ",zap.Any("balance",balance))

	tokenIds,err := opt.Ier.Nft.GetAddrAllTokenIds(nil,common.HexToAddress(*fflag.HandAddr),big.NewInt(10),big.NewInt(1))
	if err != nil{
		log.Error("get token ids of err ",zap.Error(err))
		return
	}
	log.Info("token ids",zap.Any("token ids",tokenIds))

	tokenList,err := opt.Ier.Handing.QueryCardList(nil,common.HexToAddress(*fflag.Owner))
	if err != nil{
		log.Error("card list err ",zap.Error(err))
		return
	}
	log.Info("token list",zap.Any("tokenlist",tokenList))*/
	/*ids, lens, err := opt.Ier.BoxPro.QuerySeriesIdsNotNull(nil)
	if err != nil {
		log.Error("query box err ", zap.Error(err))
		return
	}
	log.Info("box", zap.Any("ids", ids), zap.Any("len", lens))
	box, err := opt.Ier.BoxPro.QueryBox(nil, big.NewInt(3))
	if err != nil {
		log.Error("query box err ", zap.Error(err))
		return
	}
	log.Info("box", zap.Any("box", box))*/
}

func InitGov(opt *options.Options) {
	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	initGov, err := opt.Ier.Gov.Init(auth, common.HexToAddress(*fflag.LpAddr), common.HexToAddress(*fflag.StakingAddr))
	if err != nil {
		log.Error("init gov err ", zap.Error(err))
		return
	}
	log.Info("init gov ", zap.Any("gov", initGov.Hash()))
}

func addktoken(opt *options.Options) {
	config, err := opt.Ier.Box.QueryConfig(nil)
	if err != nil {
		log.Error("query config err", zap.Error(err))
		return
	}
	fflag.KTokenAddr = config.KeyToken.String()
}

func fetchK(number *big.Int, opt *options.Options) (common.Address, error) {
	instance, _ := buildtoken.NewStore(common.HexToAddress(*fflag.BuildAddr), opt.Client)
	query := ethereum.FilterQuery{

		FromBlock: number,
		Addresses: []common.Address{common.HexToAddress(*fflag.BuildAddr)},
	}
	logs, err := opt.Client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Error("fliter log err ", zap.Error(err))
		return common.HexToAddress(""), err
	}
	logTransferSigHash := crypto.Keccak256Hash([]byte(`CreatedControlledToken(address)`))
	for _, logf := range logs {
		switch logf.Topics[0].Hex() {
		case logTransferSigHash.Hex():
			event, err := instance.ParseCreatedControlledToken(logf)
			if err != nil {
				log.Error("logf err ", zap.Error(err))
				return common.HexToAddress(""), err
			}
			return event.Token, nil
		}
	}
	return common.HexToAddress(""), errors.New("not found")
}

func addr(req string) common.Address {
	return common.HexToAddress(req)
}
func array(req []string) []*big.Int {
	var result []*big.Int
	for _, v := range req {
		s, _ := decimal.NewFromString(v)
		result = append(result, s.BigInt())
	}
	return result
}
func arraddr(req []string) []common.Address {
	var result []common.Address
	for _, v := range req {
		result = append(result, addr(v))
	}
	return result
}

