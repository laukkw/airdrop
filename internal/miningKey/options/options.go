package options

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	cliflag "github.com/marmotedu/component-base/pkg/cli/flag"
	"github.com/marmotedu/component-base/pkg/json"
	"github.com/rzry/airdrop/pkg/log"
)

type Options struct {
	Log         *log.Options      `json:"log"   mapstructure:"log"`
	Client      *ethclient.Client `json:"client"`
	MainAddress common.Address    `json:"main_address"`
}

func (o *Options) Validate() []error {
	panic("implement me")
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

func (o *Options) String() string {
	data, _ := json.Marshal(o)

	return string(data)
}

// Flags returns flags for a specific APIServer by section name.
func (o *Options) Flags() (fss cliflag.NamedFlagSets) {
	o.Log.AddFlags(fss.FlagSet("logs"))

	return fss
}
