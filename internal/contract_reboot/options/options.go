package options

import (
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/component-base/pkg/json"
	"github.com/rzry/airdrop/internal/contract_reboot/fflag"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/blindbox"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/boxpromotion"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/flip"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/gov"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/handingPro"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/lock"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/nft"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/prizepool"
	"github.com/rzry/airdrop/internal/contract_reboot/pkg/token"
	"github.com/rzry/airdrop/pkg"
	log "github.com/rzry/kwstart/kwlog"
)

type Options struct {
	Log    *log.Options      `json:"log"   mapstructure:"log"`
	Client *ethclient.Client `json:"client"`
	Ier    *Instance
}

type Instance struct {
	Box     *blindbox.Store
	BoxPro  *boxpromotion.Store
	Flip    *flip.Store
	Gov     *gov.Store
	Handing *handingPro.Store
	Nft     *nft.Store
	Pool    *prizepool.Store
	Ptoken  *token.Store
	Ktoken  *token.Store
	Lock    *lock.Store
}

func (o *Options) Validate() []error {
	panic("implement me")
}
func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.Log.AddFlags(fss.FlagSet("logs"))

	return fss
}

func NewOptions(name string) *Options {
	optsLog := &log.Options{
		Level:            "info",
		Format:           "console",
		EnableColor:      false,
		DisableCaller:    true,
		OutputPaths:      []string{fmt.Sprintf("%s.log", name), "stdout"},
		ErrorOutputPaths: []string{"stderr"},
		Development:      false,
	}
	o := Options{
		Log: optsLog,
	}
	return &o
}

func Auth(opt *Options) (*bind.TransactOpts, error) {
	return pkg.Auth(*fflag.Private, fflag.Chain, opt.Client)
}
