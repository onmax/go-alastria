package txutil

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/hex"
	alaTypes "github.com/onmax/go-alastria/types"
)

func Nonce(eth *ethclient.Client, from common.Address) (uint64, error) {
	return eth.PendingNonceAt(context.Background(), from)
}

func UpdateNonce(conn *alaTypes.Connection) error {
	nonce, err := Nonce(conn.Client.Eth, conn.Client.Ks.Account.Address)
	if err != nil {
		return err
	}
	conn.Tx.Opts.Nonce = big.NewInt(0).SetUint64(nonce)
	return nil
}

func HexToTx(txStr string) (*types.Transaction, error) {
	txHex, err := hex.HexToByteArr(hex.Remove0x(txStr))
	if err != nil {
		return nil, err
	}
	tx := &types.Transaction{}
	tx.UnmarshalJSON(txHex)
	return tx, nil
}

func TxToHex(tx *types.Transaction) (string, error) {
	txByteArr, err := tx.MarshalJSON()
	if err != nil {
		return "", err
	}
	txStr := hex.ByteArrToHex(txByteArr)
	return hex.Add0x(txStr), nil
}
