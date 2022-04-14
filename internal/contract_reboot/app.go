package contract_reboot

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/manifoldco/promptui"
	"github.com/rzry/airdrop/internal/contract_reboot/fflag"
	"github.com/rzry/airdrop/internal/contract_reboot/options"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/blindbox"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/boxpromotion"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/flip"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/gov"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/handingPro"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/lock"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/nft"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/prizepool"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/token"
	log "github.com/rzry/kwstart/kwlog"
	"go.uber.org/zap"
	"os"
	"sync"
	"time"
)

var once sync.Once

func NewApp(basename string) {
	opts := options.NewOptions(basename)
	fflag.InitFlag()
	once.Do(func() {
		client, err := ethclient.Dial(fflag.Url)
		if err != nil {
			os.Exit(1)
		}
		opts.Client = client
		initInstance(client, opts)
	})
	run(opts)
}

func initInstance(c *ethclient.Client, opts *options.Options) {
	boxer, err := blindbox.NewStore(common.HexToAddress(*fflag.BoxAddr), c)
	boxproer, err := boxpromotion.NewStore(common.HexToAddress(*fflag.BoxProAddr), c)
	fliper, err := flip.NewStore(common.HexToAddress(*fflag.FlipAddr), c)
	gover, err := gov.NewStore(common.HexToAddress(*fflag.GovAddr), c)
	hander, err := handingPro.NewStore(common.HexToAddress(*fflag.HandAddr), c)
	nfter, err := nft.NewStore(common.HexToAddress(*fflag.NftAddr), c)
	pooler, err := prizepool.NewStore(common.HexToAddress(*fflag.PoolAddr), c)
	ptoken, err := token.NewStore(common.HexToAddress(*fflag.PTokenAddr), c)
	locked, err := lock.NewStore(common.HexToAddress(*fflag.LockAddr), c)
	//
	if err != nil {
		log.Fatal("init err", zap.Error(err))
	}
	opts.Ier = &options.Instance{
		Box:     boxer,
		BoxPro:  boxproer,
		Flip:    fliper,
		Gov:     gover,
		Handing: hander,
		Nft:     nfter,
		Pool:    pooler,
		Ptoken:  ptoken,
		Lock:    locked,
		//Ktoken:  ktoken,
	}
	return
}

func run(opts *options.Options) {
	log.Init(opts.Log)
	defer log.Flush()
	prompt := promptui.Select{
		Label: "What do you want to do?",
		Items: []string{
			"初始化二期,创建系列,模拟抽取",
			"初始化一期",
			"初始化二期",
			"创建系列",
			"模拟抽取",
			"查询抽取",
			"发放奖励",
			"领取奖励",
		},
	}
	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}
	ctx := context.Background()
	switch result {
	case "初始化二期,创建系列,模拟抽取":
		AddBoxPro(opts)
		time.Sleep(3*time.Second)
		CreateSeriesV2(opts,1)
		time.Sleep(3*time.Second)
		TestDraw(opts,1)
	case "初始化一期":
		InitOne(opts,ctx)
	case "初始化二期":
		AddBoxPro(opts)
	case "创建系列":
		CreateSeriesV2(opts,1)
	case "模拟抽取":
		TestDraw(opts,1)
	case "查询抽取":
		QueryHanding(opts)
	case "发放奖励":
		SendAward(opts)
	case "领取奖励":
		Received(opts)
	}

}
