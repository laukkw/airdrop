package fflag

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
	"os"
)

var (
	Private         *string
	Token           *string
	Amount          *int64
	BlindBoxAddress *string
	InviterAddress  *string
	Url             string
	ChainId         int64
	Goroutine       *int64
)

var desc = `这是一个批量换取k token 的程序
请根据需求填写参数,邀请人地址会获得每笔5%的平台币奖励
兑换后所获得的k token 都会存留与私钥地址账户,请自行转出或分发
请求示例:
./miningKey --token="0xd9971bff10E4e0465B21acf219aE6590c6952678" --amount=1110
您会在当前文件夹下得到一个文件名为miningKey.log的文件来记录所有兑换信息.kkw
`

func InitFlag() {
	_ = pflag.String("help", "", desc)
	Private = pflag.String("key", "e7b609a399b0f88fe6054c4810dbaebec670643c16e43a1aea7e7ee8952b6206", "填入支付账户的私钥,在命令前加空格,就不会将命令保存到历史记录")
	Amount = pflag.Int64("amount", 110, "填入需要换取k Token的数量")
	Token = pflag.String("token", "0xd9971bff10E4e0465B21acf219aE6590c6952671", "填入平台币的token")
	BlindBoxAddress = pflag.String("blindBoxAddress", "0xaf69F26d4E76A73A73bA14b799189Fa54577CA41", "填入大盲盒地址")
	InviterAddress = pflag.String("inviterAddress", "0x4936e6A6A989b4B6101D7Bd70a8e494649853360", "邀请人地址")
	Goroutine = pflag.Int64("goroutine", 10, "goroutine池容量(因测试网与正式网有差,所以如果并发请求错误过多,请降低这个值,测试网50是可以完美运行,如果请求数量很多,那求稳考虑,请降低这个值在50左右)")
	pflag.Parse()
	prompt := promptui.Select{
		Label: "Select Your NetWork",
		Items: []string{"bsc_testnet",
			"heco_testnet",
			"heco",
			"bsc",
			"goerli"},
	}
	_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}
	rs := net_worl()[result]
	ChainId = rs.ChainId
	Url = rs.Url
}

type NetWork struct {
	Url     string
	ChainId int64
}

func net_worl() map[string]NetWork {
	mk := make(map[string]NetWork, 0)
	mk["bsc_testnet"] = NetWork{
		Url:     "https://data-seed-prebsc-2-s3.binance.org:8545",
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
