package crypto

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/hex"
	"golang.org/x/crypto/sha3"
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

func PublicKeyToAddress(publicKeyStr string) (common.Address, error) {
	publicKey, err := hex.HexToByteArr(publicKeyStr) // TODO Sanitize removing 0x and 03 || 04
	if err != nil {
		return common.Address{}, err
	}
	var buf []byte

	hash := sha3.NewLegacyKeccak256()
	hash.Write(publicKey[:])
	buf = hash.Sum(nil)
	publicAddress := hexutil.Encode(buf[12:])

	return common.HexToAddress(publicAddress), nil
}
