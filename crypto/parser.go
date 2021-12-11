package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/hex"
)

func HexToECDSAPub(_pub string) (*ecdsa.PublicKey, error) {
	var ok bool

	_pub = hex.Remove0x(_pub)

	x := new(big.Int)
	x, ok = x.SetString(_pub[:len(_pub)/2], 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex") // TODO
	}

	y := new(big.Int)
	y, ok = y.SetString(_pub[len(_pub)/2:], 16)
	if !ok {
		return nil, fmt.Errorf("invalid hex") // TODO
	}

	return &ecdsa.PublicKey{X: x, Y: y, Curve: crypto.S256()}, nil
}


