package fflag

import (
	"github.com/manifoldco/promptui"
	"github.com/spf13/pflag"
	"os"
)

var (
	Private  *string
	Token    *string
	Path     *string
	Amount   *string
	WaitTime *int64
	Url      string
	ChainId  int64
)

var desc = `这是一个空投机器人,你需要根据实际情况填写所需参数,请仔细阅读每参数的描述
空投人地址应以回车结尾,只有地址
eg:
0x6544c9e64d6d554791283faaabe13b6b6bb88b66
0xcec1dbD16798C6582B20417e5Abd0282307ad6F8
!!!空投数量单位为ethers,所以不需要到18位小数
使用示例
./airdrop --amount="1" --key="e7b609a399b0f88fe6054c4810dbaebec670643c16e43a1aea7e7ee8952b6201" --token="0xB1A644CeF2ab7Ed8b41399C02df18c53A42Dc6da"
您会在当前文件夹下得到一个文件名为airdrop.log的文件来记录所有空投信息
`

func InitFlag() {
	_ = pflag.String("help", "", desc)
	Private = pflag.String("key", "e7b609a399b0f88fe6054c4810dbaebec670643c16e43a1aea7e7ee8952b6216", "填入支付账户的私钥,在命令前加空格,就不会将命令保存到历史记录")
	Amount = pflag.String("amount", "1.13123", "填入需要drop的数量,ethers为单位")
	Token = pflag.String("token", "0xd9971bff10E4e0465B21acf219aE6590c6952678", "填入需要drop的token")
	Path = pflag.String("path", "./path.txt", "填入空投人地址文件的路径,默认./path.txt(当前文件夹下 path.txt文件)")
	WaitTime = pflag.Int64("wait", 5, "输入等待时间,等待时间到期执行空投操作,单位:s")
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
