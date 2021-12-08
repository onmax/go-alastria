package tx

import (
	"context"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/hex"
)

func Nonce(eth *ethclient.Client, from common.Address) (uint64, error) {
	return eth.PendingNonceAt(context.Background(), from)
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
