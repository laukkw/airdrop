package miningKey

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rzry/airdrop/internal/miningKey/fflag"
	"github.com/rzry/airdrop/internal/miningKey/options"
	"github.com/rzry/airdrop/internal/miningKey/service"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
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
		opts.MainAddress = pkg.PrivateToAddress(*fflag.Private)
	})
	run(opts)
	return
}

func run(opts *options.Options) {
	log.Init(opts.Log)
	defer log.Flush()
	service.MiningKey(opts)
}
