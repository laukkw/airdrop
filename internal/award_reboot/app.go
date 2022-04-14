package award_reboot

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rzry/airdrop/internal/award_reboot/fflag"
	"github.com/rzry/airdrop/internal/award_reboot/options"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/handingPro"
	log "github.com/rzry/kwstart/kwlog"
	"go.uber.org/zap"
	"os"
	"sync"
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
	hander, err := handingPro.NewStore(common.HexToAddress(*fflag.HandAddr), c)
	//
	if err != nil {
		log.Fatal("init err", zap.Error(err))
	}
	opts.Ier = &options.Instance{
		Handing: hander,
	}
	return
}

func run(opts *options.Options) {
	log.Init(opts.Log)
	defer log.Flush()
	award(opts)
}
