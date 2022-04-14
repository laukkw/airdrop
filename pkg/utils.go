package pkg

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/shopspring/decimal"
	"log"
	"math/big"
)

func PrivateToAddress(key string) common.Address {
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Println("err == ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("call err ", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	return fromAddress
}
func Auth(key string, chainId int64, client *ethclient.Client) (*bind.TransactOpts, error) {
	ctx := context.Background()
	privateKey, err := crypto.HexToECDSA(key)
	if err != nil {
		log.Println("err == ", err)
	}
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Println("call err ", err)
	}
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(ctx, fromAddress)
	if err != nil {
		log.Println("当前的nonce为", nonce)
		return nil, err
	}
	gasPrice, err := client.SuggestGasPrice(ctx)
	if err != nil {
		log.Println("call  err ", err)
	}
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(chainId))
	if err != nil {
		log.Println("call  err ", err)
		return nil, err
	}
	//将nonce改回去
	//auth.Nonce = big.NewInt(int64(nonce - 1))
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei

	auth.GasLimit = uint64(205000) // in units

	auth.GasPrice = gasPrice
	auth.Context = context.Background()
	return auth, err
}
func ToEthers(req string, quote int64) decimal.Decimal {
	v, _ := decimal.NewFromString(req)
	return v.Mul(decimal.New(10, 0).Pow(decimal.New(quote, 0)))
}
func FromEthers(req string, quote int64) decimal.Decimal {
	v, _ := decimal.NewFromString(req)
	return v.DivRound(decimal.New(10, 0).Pow(decimal.New(quote, 0)), int32(quote))
}
