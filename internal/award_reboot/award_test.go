package award_reboot

import (
    "github.com/ethereum/go-ethereum/common"
    "github.com/ethereum/go-ethereum/ethclient"
    "github.com/rzry/airdrop/internal/contract_reboot/pkg/boxpromotion"
    "github.com/rzry/airdrop/internal/contract_reboot/pkg/handingPro"
    "github.com/rzry/airdrop/pkg"
    "math/big"
    "os"
    "testing"
)
var Handing = "0x78a37A07Ca8aE32D46A38d2563285967796efe42"
var Boxpro = "0xD96Bb280155ca1614862553bC30645Ec5c00B14b"
//var PrivateKey = "5c79dcea5b19785a77b3cf48a0f63ef5bd8f6c415cc157805d8b94559e376bd0"
var PrivateKey = ""
var Owner = "0xc63d8489814751aae8054bc21d306a29387c5934"
var ChainId int64 = 56

var Client *ethclient.Client
var HandingInstance *handingPro.Store
var BoxProInstance *boxpromotion.Store

func TestMain(t *testing.M) {
	client, err := ethclient.Dial("https://bsc-dataseed1.binance.org")
	if err != nil {
		os.Exit(1)
	}
	Client = client
    hander,err := handingPro.NewStore(common.HexToAddress(Handing),Client)
    if err != nil{
        os.Exit(2)
    }
    HandingInstance = hander
    box , err := boxpromotion.NewStore(common.HexToAddress(Boxpro),client)
    if err != nil{
        os.Exit(3)
    }
    BoxProInstance = box
    t.Run()
}
func TestDraw(t *testing.T) {
    t.Run("测试合约是否正常", func(t *testing.T) {
        recorder ,err :=HandingInstance.Recorder(nil)
        if err != nil{
            t.Error(err)
            return
        }
        t.Log("记账者为:",recorder)
        boxConfig ,err := BoxProInstance.Config(nil)
        if err != nil{
            t.Fatal(err)
        }
        t.Logf("%+v",boxConfig)
        
        seriesIds , err := BoxProInstance.QuerySeriesIds(nil)
        if err != nil{
            t.Fatal(err)
        }
        t.Log(seriesIds)
        
        config,err := BoxProInstance.QueryBox(nil,big.NewInt(2))
        if err != nil{
            t.Fatal(err)
        }
        t.Logf("%+v",config)
        t.Log(HandingInstance.CycleRng(nil,big.NewInt(7)))
        
        
    })
    t.Run("修改handing_pro合约记账者 为owner", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := HandingInstance.ResetRecorder(auth,common.HexToAddress(Owner))
        if err != nil{
            t.Fatal(err)
        }
        t.Log("修改记账者",tx.Hash())
        //重复查询记账者
    })
    t.Run("仿造抽取", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := HandingInstance.RecordDraw(auth,common.HexToAddress(Owner),big.NewInt(100),big.NewInt(2))
        if err != nil{
            t.Fatal(err)
        }
        t.Log("仿造抽取:",tx.Hash())
    })
    
    t.Run("开奖", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := HandingInstance.Award(auth)
        if err != nil{
            t.Fatal(err)
        }
        t.Log("开奖Hash",tx.Hash())
    })
    
    t.Run("领取奖励", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := HandingInstance.Received(auth,
            []*big.Int{
            //领取的卡片id
                big.NewInt(1228),
                big.NewInt(1229),
            },
            []*big.Int{
                big.NewInt(1),
                big.NewInt(1),
            })
        if err != nil{
            t.Error(err)
            return 
        }
        t.Log("领取卡片",tx.Hash())
    })
    
    t.Run("增加二期奖励", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := BoxProInstance.ResetReward(auth,big.NewInt(2),boxpromotion.BlindBoxPromotionReward{
            Token:  []common.Address{
                common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
                common.HexToAddress("0x0e09fabb73bd3ade0a17ecc321fd13a19e81ce82"),
                common.HexToAddress("0xcB89da59124046A5Da29af9AFd59EA025eF11B84"),
            },
            Amount: []*big.Int{
                pkg.ToEthers("1",18).BigInt(),
                pkg.ToEthers("10.079",18).BigInt(),
                pkg.ToEthers("28110",18).BigInt(),
            },
        })
        t.Log("增加二期奖励",tx.Hash())
    })
    
    t.Run("兑奖", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := BoxProInstance.Convert(auth,big.NewInt(2),[]*big.Int{
            big.NewInt(2186),
            big.NewInt(1287),
            big.NewInt(1218),
            big.NewInt(1219),
            big.NewInt(1229),
        })
        if err != nil{
            t.Fatal(err)
        }
        t.Log("兑奖",tx.Hash())
    })
    t.Run("恢复合约记账者", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := HandingInstance.ResetRecorder(auth,common.HexToAddress(Boxpro))
        if err != nil{
            t.Fatal(err)
        }
        t.Log("恢复记账者",tx.Hash())
        //重复查询记账者
    })
    t.Run("恢复奖励", func(t *testing.T) {
        auth,err := pkg.Auth(PrivateKey,ChainId,Client)
        if err != nil{
            t.Fatal(err)
        }
        tx,err := BoxProInstance.ResetReward(auth,big.NewInt(2),boxpromotion.BlindBoxPromotionReward{
            Token:  []common.Address{
                common.HexToAddress("0x2170Ed0880ac9A755fd29B2688956BD959F933F8"),
            },
            Amount: []*big.Int{
                pkg.ToEthers("1",18).BigInt(),
            },
        })
        t.Log("恢复二期奖励",tx.Hash())
    })
}
