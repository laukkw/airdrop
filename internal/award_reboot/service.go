package award_reboot

import (
	"github.com/rzry/airdrop/internal/award_reboot/options"
	"github.com/rzry/airdrop/pkg/log"
	"go.uber.org/zap"
	"math/big"
	"sync/atomic"
)

var sum int64

func award(opt *options.Options) {
	log.Info("start award ")

	auth, err := options.Auth(opt)
	if err != nil {
		log.Error("get auth err", zap.Error(err))
		return
	}
	sum = auth.Nonce.Int64()
	auth.Nonce = big.NewInt(atomic.AddInt64(&sum, 1))
	AwardTx, err := opt.Ier.Handing.Award(auth)
	if err != nil {
		log.Error("Award err", zap.Error(err))
		return
	}
	log.Info("发奖", zap.Any("hash", AwardTx.Hash()))
	log.Info("end award ")
}
