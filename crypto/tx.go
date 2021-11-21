package crypto

import (
	"crypto/ecdsa"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/core/types"
)

func SignTx(networkId *big.Int, tx *types.Transaction, privateKey *ecdsa.PrivateKey) *types.Transaction {
	signedTx, err := types.SignTx(tx, types.NewEIP155Signer(networkId), privateKey)
	if err != nil {
		log.Fatal(err)
	}
	return signedTx
}
