package hex

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
)

func Remove0x(input string) string {
	if len(input) >= 2 && input[0:2] == "0x" {
		return input[2:]
	}
	return input
}

func HexToECDSAPub(_pub string) (*ecdsa.PublicKey, error) {
	var ok bool

	x := new(big.Int)
	x, ok = x.SetString(_pub[:len(_pub)/2], 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex")
	}

	y := new(big.Int)
	y, ok = y.SetString(_pub[len(_pub)/2:], 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex")
	}

	return &ecdsa.PublicKey{X: x, Y: y, Curve: crypto.S256()}, nil
}
