package tx

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func TxOpt(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = 600000
	opts.GasPrice = big.NewInt(0)
	opts.NoSend = true
	fmt.Printf("opts: %+v\n", opts)
	return opts, nil
}

func Nonce(eth *ethclient.Client, from common.Address) (uint64, error) {
	return eth.PendingNonceAt(context.Background(), from)
}
