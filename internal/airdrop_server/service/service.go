package service

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/rzry/airdrop/internal/airdrop_server/fflag"
	"github.com/rzry/airdrop/internal/airdrop_server/options"
	"github.com/rzry/airdrop/internal/airdrop_server/pkg/platform"
	"github.com/rzry/airdrop/pkg"
	"github.com/rzry/airdrop/pkg/log"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
	"math/big"
	"strings"
	"time"
)

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
	auth, err := pkg.Auth(*fflag.Private, fflag.ChainId, opts.Client)
	if err != nil {
		log.Error("获取私钥签名出错,请检查私钥", zap.Error(err))
		return
	}
	if balance, err := platfromInstance.BalanceOf(nil, opts.MainAddress); decimal.Zero.GreaterThan(decimal.NewFromBigInt(balance, 0)) ||
		err != nil {
		log.Error("输入的主地址私钥余额为空,或有错误", zap.Error(err))
		return
	}
	waitDrops := ReadFile(*fflag.Path)
	for k, waitDrop := range waitDrops {
		if !strings.HasPrefix(strings.TrimSpace(waitDrop), "0x"){
			log.Info("并非为地址,跳过此行",zap.Any("数值",strings.TrimSpace(waitDrop)))
			continue
		}
		auth.Nonce = auth.Nonce.Add(auth.Nonce, big.NewInt(1))
		tx, err := platfromInstance.Transfer(auth, common.HexToAddress(strings.TrimSpace(waitDrop)),
			amount.BigInt())
		if err != nil {
			log.Error("转账失败", zap.Error(err))
		}
		log.Infof("转账给%s,交易hash为%s,总待空投数%v,当前到达%v", strings.TrimSpace(waitDrop), tx.Hash(),len(waitDrops),k+1)
	}
}
