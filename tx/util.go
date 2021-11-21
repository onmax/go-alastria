package tx

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

func TxOpt(privateKey *ecdsa.PrivateKey) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(1337))
	if err != nil {
		return nil, err
	}
	opts.Value = big.NewInt(0)
	opts.GasLimit = 600000
	opts.GasPrice = big.NewInt(0)
	fmt.Println("opts %s", opts)
	return opts, nil
}
