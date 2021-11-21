package crypto

import (
	"crypto/ecdsa"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/onmax/go-alastria/network"
)

func SignTx(client *ethclient.Client, tx *types.Transaction, privateKey *ecdsa.PrivateKey) *types.Transaction {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(network.NetworkID(client)), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx
}
