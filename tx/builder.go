package tx

import (
	"crypto/ecdsa"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/internal/configuration"
)

func TxOpt(privateKey *ecdsa.PrivateKey, chainId *big.Int) (*bind.TransactOpts, error) {
	opts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}
	opts.Value = configuration.TxValue
	opts.GasLimit = configuration.TxGasLimit
	opts.GasPrice = configuration.TxGasPrice
	opts.NoSend = configuration.TxNoSend
	return opts, nil
}

func TxFromString(eth *ethclient.Client, tx []byte, from common.Address, to *common.Address) (*types.Transaction, error) {
	nonce, err := Nonce(eth, from)
	if err != nil {
		return nil, err
	}
	baseTx := &types.LegacyTx{
		To:       to,
		Nonce:    nonce,
		GasPrice: configuration.TxGasPrice,
		Value:    configuration.TxValue,
		Data:     []byte(tx),
	}
	return types.NewTx(baseTx), nil
}
