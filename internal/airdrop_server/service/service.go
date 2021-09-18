package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/panjf2000/ants/v2"
	"github.com/rzry/airdrop/internal/airdrop_server/fflag"
	"github.com/rzry/airdrop/internal/airdrop_server/options"
	"github.com/rzry/airdrop/internal/airdrop_server/pkg/platform"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"sync"
	"sync/atomic"
	"time"
)
var sum int64
var index int64
func Airdrop(opts *options.Options) {
	log.Info("start airdrop !!!")
	amount := pkg.ToEthers(*fflag.Amount, 18)
	log.Infof("正在转账的token为%s,主地址为%s,交易数为%s,等待时间为%vs", *fflag.Token, opts.MainAddress, amount,*fflag.WaitTime)
	if !pkg.YorN(){
		return
	}
	time.Sleep(time.Duration(*fflag.WaitTime)*time.Second)
	platfromInstance, err := platform.NewStore(common.HexToAddress(*fflag.Token), opts.Client)
	if err != nil {
		log.Error("链接token合约出错", zap.Error(err))
		return
	}
	 balance, err := platfromInstance.BalanceOf(nil, opts.MainAddress)
	 if decimal.Zero.GreaterThan(decimal.NewFromBigInt(balance, 0)) || err != nil {
		log.Error("输入的主地址私钥余额为空,或有错误", zap.Error(err))
		return
	}
	log.Info("当前地址余额为",zap.Any("balance",pkg.FromEthers(balance.String(),18)))
	auth, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
	if err != nil {
		log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	var waitList  []string
	var wg sync.WaitGroup
	syncCalculateSum := func() {
		defer wg.Done()
		indexes := atomic.AddInt64(&index,1)
		auths, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
		if err != nil {
			log.Error("转账失败", zap.Error(err),zap.Any("user",waitList[indexes-1]))
			log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
			return
		}
		auths.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
		tx, err := platfromInstance.Transfer(auths, common.HexToAddress(strings.TrimSpace(waitList[indexes-1])),
			amount.BigInt())
		if err != nil {
			log.Error("转账失败", zap.Error(err),zap.Any("user",waitList[indexes-1]))
			return
		}
		log.Infof("转账给%s,交易hash为%s,当前nonce为%d", strings.TrimSpace(waitList[indexes-1]), tx.Hash(),auths.Nonce.Int64())
	}
	waitDrops := ReadFile(*fflag.Path)

	p, _ := ants.NewPoolWithFunc(int(*fflag.Goroutine), func(i interface{}) {
		syncCalculateSum()
	})
	for k, waitDrop := range waitDrops {
		if !strings.HasPrefix(strings.TrimSpace(waitDrop), "0x"){
			log.Info("并非为地址,跳过此行",zap.Any("数值",strings.TrimSpace(waitDrop)))
			continue
		}
		wg.Add(1)
		waitList = append(waitList,waitDrop)
		_ = p.Invoke(k)
	}
	wg.Wait()

	//转账结束后,获取错误行补发.
	addtransfer := ReadAddFile("./airdrop.log")
	time.Sleep(10*time.Second)
	log.Info("转账结束后,获取错误行补发",zap.Any("错误数",len(addtransfer)))
	authd, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
	if err != nil {
		log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
		return
	}
	for _,v := range addtransfer{
		authd.Nonce = authd.Nonce.Add(authd.Nonce,big.NewInt(1))
		tx, err := platfromInstance.Transfer(authd, common.HexToAddress(strings.TrimSpace(v)),
			amount.BigInt())
		if err != nil {
			log.Error("补发失败,请手动补发", zap.Error(err),zap.Any("user",v))
			continue
		}
		log.Infof("补发给%s,交易hash为%s,当前nonce为%d", strings.TrimSpace(v), tx.Hash(),authd.Nonce.Int64())
	}
	log.Info("补发结束")
	balance, err = platfromInstance.BalanceOf(nil, opts.MainAddress)
	if  err != nil {
		log.Error("输入的主地址私钥余额为空,或有错误", zap.Error(err))
		return
	}
	time.Sleep(10*time.Second)
	log.Info("当前地址余额为",zap.Any("balance",pkg.FromEthers(balance.String(),18)))

}
