package network

import (
	"context"
	"errors"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	alaTypes "github.com/onmax/go-alastria/types"
)

func ConnectToNetwork(nodeUrl string) (*ethclient.Client, error) {
	if nodeUrl == "" {
		return nil, errors.New("no node url specified")
	}
	return ethclient.Dial(nodeUrl)
}

func NetworkID(client *ethclient.Client) (*big.Int, error) {
	if client == nil {
		return nil, errors.New("no client specified")
	}
	return client.NetworkID(context.Background())
}

func checkTransactionReceipt(conn *alaTypes.Connection, txHash common.Hash) uint64 {
	tx, err := conn.Client.Eth.TransactionReceipt(context.Background(), txHash)
	if err != nil {
		return 0
	}
	return tx.Status
}

// Sends a transaction to the network. It will wait until the transaction is mined blocking
// the current thread checking once every second
// conn needs to have a client set.
func SendTx(conn *alaTypes.Connection, tx *types.Transaction) error {
	err := conn.Client.Eth.SendTransaction(context.Background(), tx)

	if err != nil {
		return err
	}

	for {
		status := checkTransactionReceipt(conn, tx.Hash())
		if status == 1 {
			return nil
		}
		time.Sleep(time.Second)
	}
}
