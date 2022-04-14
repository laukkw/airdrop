package fflag

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
	"os"
)

var (
	HandAddr *string
	Private  *string
	Url      string
	Chain    int64
	GasLimit    *int64
)
var desc = `此程序提供luca系列定时开奖功能,依次填入所需参数,`

func InitFlag() {
	_ = pflag.String("help", "", desc)
	HandAddr = pflag.String("hand", "0x5B79f279A5847395220Ab61761C841ceC336E2F6", "handing合约地址")
	Private = pflag.String("private", "bc2cb7061cb526c6ffaf5efc185ad3ab0b3a8bd557e9a1318179d5bfcf9ea82d", "操作的密钥")
	//Times = pflag.Int64("time", 120, "等待的时间")
	GasLimit = pflag.Int64("gas", 20500000, "gas limit")
	pflag.Parse()
	prompt := promptui.Select{
		Label: "Select Your NetWork",
		Items: []string{"bsc_testnet",
			"heco_testnet",
			"heco",
			"bsc",
			"goerli", "bsc_testnet1", "bsc_testnet2", "bsc_testnet3"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}
	rs := net_worl()[result]
	Chain = rs.ChainId
	Url = rs.Url
}

type NetWork struct {
	Url     string
	ChainId int64
}

func net_worl() map[string]NetWork {
	mk := make(map[string]NetWork, 0)
	mk["bsc_testnet"] = NetWork{
		Url:     "https://data-seed-prebsc-1-s3.binance.org:8545",
		ChainId: 97,
	}
	mk["bsc_testnet2"] = NetWork{
		Url:     "https://data-seed-prebsc-2-s3.binance.org:8545",
		ChainId: 97,
	}
	mk["bsc_testnet3"] = NetWork{
		Url:     "wss://data-seed-prebsc-1-s2.binance.org:8545",
		ChainId: 97,
	}
	mk["bsc_testnet1"] = NetWork{
		Url:     "https://data-seed-prebsc-2-s2.binance.org:8545",
		ChainId: 97,
	}
	mk["heco_testnet"] = NetWork{
		Url:     "https://http-testnet.hecochain.com",
		ChainId: 256,
	}
	mk["heco"] = NetWork{
		Url:     "https://http-mainnet.hecochain.com",
		ChainId: 128,
	}
	mk["bsc"] = NetWork{
		Url:     "https://bsc-dataseed1.binance.org",
		ChainId: 56,
	}
	mk["goerli"] = NetWork{
		Url:     "https://goerli.infura.io/v3/fbc55f1f018f4c95bb74cdf75092f31b",
		ChainId: 5,
	}
	return mk
}
