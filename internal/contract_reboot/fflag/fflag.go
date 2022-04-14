package fflag

import (
	"github.com/spf13/pflag"
)

var (
	BoxAddr     *string
	BoxProAddr  *string
	FlipAddr    *string
	GovAddr     *string
	HandAddr    *string
	NftAddr     *string
	PoolAddr    *string
	PTokenAddr  *string
	KTokenAddr  string
	Private     *string
	BuildAddr   *string
	LableAddr   *string
	RandAddr    *string
	LockAddr    *string
	LpAddr      *string
	StakingAddr *string
	Owner       *string
	Url         string
	Chain       int64
)

var desc = `本程序是为了初始化,以及可以选择的测试流程,依次填入所需所有合约地址.
仅用与自测便利 , 选项解释: 
init contract -- 初始化除了gov 以外的所有合约
create series -- 创建box的系列,系列id为1
fetch ktoken  -- 抽取一次ktoken 10连抽
fetch nft     -- 抽取一次 box 系列为1 的nft 
create seriesV2 -- 创建 box pro (第二期盲盒) 的系列,系列id为1
fetch nftv2   -- 抽取一次box pro 系列为1 的nft 
add recoder   -- 添加记账者权限,默认初始化是没有的一个记账者的
send award    -- 发送奖励(会修改默认时间,并且执行记录账单)
received      -- 领取奖励
`

func InitFlag() {
	_ = pflag.String("help", "", desc)
	PTokenAddr = pflag.String("ptoken", "0x297aF953E0C4f4305A7Bf3dF7083ba7AEAd28940", "平台币合约地址")
	BuildAddr = pflag.String("build", "0x0F4E8a205eA57675E67AB1B8561c3981B91ee297", "buildToken合约地址")
	RandAddr = pflag.String("rand", "0xbd36590fa5Ffba451971268C006912f88700413f", "随机数合约地址,tokenid合约地址")
	BoxProAddr = pflag.String("boxpro", "0x1AC19150a0102fa51e143897a94D2883051927aa", "二期合约大盲盒合约地址")
	HandAddr = pflag.String("hand", "0x38dE75A75BD4c9928Bbc3cdD8B062ed67Eb78B46", "handing合约地址")
	GovAddr = pflag.String("gov", "0x143D37539f1b87Fdbfd42De0C4F4C5FAe66A612D", "gov合约地址")
	NftAddr = pflag.String("nft", "0x954Ac8fe2eBC52DaefEd68694F4162b27a6f8D2C", "nft合约地址")
	BoxAddr = pflag.String("box", "0xCC1D15d5aC59De28fFDDFc9c3009a6D9B5110474", "一期合约大盲盒合约地址")
	PoolAddr = pflag.String("pool", "0xfF9D8Bf11CB93eD9c00A685CbFBBCe28Ef253714", "奖池合约地址")
	FlipAddr = pflag.String("flip", "0x014C7FB9D6a053FCBD23f2ade62D0aDcc9e85C69", "flip合约地址")
	LockAddr = pflag.String("lock", "0x4e9DD9C355732f2B649507098b9E6794F6622F36", "lock合约 地址")
	LpAddr = pflag.String("lp", "0xc21777757fe032b680154426a53326c8bc40ac53", "lp合约 地址")
	StakingAddr = pflag.String("staking", "0xb061DaD4738198F1E20d6C9Fa317c788470e2eB3", "staking 合约地址")
	Private = pflag.String("private", "bc2cb7061cb526c6ffaf5efc185ad3ab0b3a8bd557e9a1318179d5bfcf9ea82d", "操作的密钥")
	//e7b609a399b0f88fe6054c4810dbaebec670643c16e43a1aea7e7ee8952b6206
	//bc2cb7061cb526c6ffaf5efc185ad3ab0b3a8bd557e9a1318179d5bfcf9ea82d
	//d5caba19226c9afadce26d51f4d1b5abdb70b141c22f9ebb88d049ee65f0f414
	LableAddr = pflag.String("lable", "0x4936e6A6A989b4B6101D7Bd70a8e494649853360", "手续费地址")
	Owner = pflag.String("owner", "0x550BF94f6A3f8344B0e4866112a0bBF839B2637d", "owner合约 地址")
	pflag.Parse()
	/*prompt := promptui.Select{
		Label: "Select Your NetWork",
		Items: []string{"bsc_testnet",
			"heco_testnet",
			"heco",
			"bsc",
			"goerli"},
	}*/
	/*_, result, err := prompt.Run()
	if err != nil {
		os.Exit(3)
	}*/
	rs := netWorl()["bsc_testnet"]
	Chain = rs.ChainId
	Url = rs.Url
}

type NetWork struct {
	Url     string
	ChainId int64
}

func netWorl() map[string]NetWork {
	mk := make(map[string]NetWork, 0)
	mk["bsc_testnet"] = NetWork{
		Url:     "wss://data-seed-prebsc-1-s3.binance.org:8545",
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
