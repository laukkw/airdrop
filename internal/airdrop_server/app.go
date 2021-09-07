package airdrop_server

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/rzry/airdrop/internal/airdrop_server/fflag"
	"github.com/rzry/airdrop/internal/airdrop_server/options"
	"github.com/rzry/airdrop/internal/airdrop_server/service"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
	"os"
)

func NewApp(basename string) {
	opts := options.NewOptions(basename)
	fflag.InitFlag()
	client, err := ethclient.Dial(fflag.Url)
	if err != nil {
		os.Exit(1)
	}
	opts.Client = client
	opts.MainAddress = pkg.PrivateToAddress(*fflag.Private)
	run(opts)
	return
}

func run(opts *options.Options) {
	log.Init(opts.Log)
	defer log.Flush()
	service.Airdrop(opts)
}
