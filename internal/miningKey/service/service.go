package service

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/panjf2000/ants/v2"
	"github.com/rzry/airdrop/internal/airdrop_server/pkg/platform"
	"github.com/rzry/airdrop/internal/miningKey/fflag"
	"github.com/rzry/airdrop/internal/miningKey/options"
	"github.com/rzry/airdrop/internal/miningKey/pkg/BlindBox"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"sync"
	"sync/atomic"
)

var sum int64

func MiningKey(opts *options.Options) {
	log.Info("start mining key !!!")
	log.Infof("使用私钥地址:%s,换取token数量:%v,使用平台币币种:%s,大盲盒地址%s,邀请人地址:%s",
		opts.MainAddress,*fflag.Amount,*fflag.Token,*fflag.BlindBoxAddress,*fflag.InviterAddress)
	//调用合约 来进行抽卡
	if !pkg.YorN(){
		return
	}
	platfromInstance, err := platform.NewStore(common.HexToAddress(*fflag.Token), opts.Client)
	if err != nil {
		log.Error("链接token合约出错", zap.Error(err))
		return
	}
	auth, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
	if err != nil {
		log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
		return
	}
	balance,err := platfromInstance.Allowance(nil,opts.MainAddress,common.HexToAddress(*fflag.BlindBoxAddress))
	if err != nil{
		log.Error("查询授权花费出错", zap.Error(err))
		return
	}
	if decimal.NewFromBigInt(balance,0).Equal(decimal.Zero){
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		if _, err := platfromInstance.Approve(auth, common.HexToAddress(*fflag.BlindBoxAddress), pkg.ToEthers("1", 50).BigInt()); err != nil {
			log.Error("approve 错误", zap.Error(err))
			return
		}
	}
	sum = auth.Nonce.Int64()
	instance, err := BlindBox.NewStore(common.HexToAddress(*fflag.BlindBoxAddress), opts.Client)
	if err != nil {
		log.Error("链接大盲盒出错", zap.Error(err))
		return
	}
	var i int64
	after := *fflag.Amount / 11
	if *fflag.Amount % 11 != 0{
		after++
	}
	//rand.Seed(time.Now().Unix())
	draw := func(auths *bind.TransactOpts) {
		//time.Sleep(time.Second * time.Duration(0 + rand.Int63n(6)))
		tx, err := instance.Draw(auths, big.NewInt(10), common.HexToAddress(*fflag.InviterAddress))
		if err != nil {
			log.Error("抽取出错", zap.Error(err))
			return
		}
		log.Info("正在抽取抽取k token", zap.Any("邀请人地址", *fflag.InviterAddress), zap.Any("交易hash", tx.Hash()),
			zap.Any("总次数",after))
	}
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		auth, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
		if err != nil {
			log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
			return
		}
		nonce := atomic.AddInt64(&sum, 1)
		auth.Nonce = big.NewInt(nonce)
		draw(auth)
		wg.Done()
	}
	for i = 0; i < after; i++ {
		wg.Add(1)
		_ = ants.Submit(syncCalculateSum)
	}
	wg.Wait()
	log.Info("共抽取", zap.Any("key token", after*11))
}
