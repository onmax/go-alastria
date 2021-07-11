package tokens

import (
	"crypto/ecdsa"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	secp256k1 "github.com/ureeves/jwt-go-secp256k1"
)

// TODO Change type
func Sign(jwt AIC, _pk string) (string, error) {
	headerJSON, _ := json.Marshal(jwt.Header)
	// TODO - check for errors
	header64 := strings.Trim(b64.StdEncoding.EncodeToString([]byte(headerJSON)), "=")
	// TODO - check for errors
	payloadJSON, _ := json.Marshal(jwt.Payload)
	payload64 := strings.Trim(b64.StdEncoding.EncodeToString([]byte(payloadJSON)), "=")

	pk, _ := crypto.HexToECDSA(_pk)
	// TODO - check for errors

	s, _ := secp256k1.SigningMethodES256K.Sign(header64+"."+payload64, pk)
	// TODO - check for errors

	return fmt.Sprintf("%s.%s.%s", header64, payload64, s), nil
}

func Verify(signed, _pub string) error {
	parts := strings.Split(signed, ".")
	pub, _ := HexToECDSAPub(_pub)
	// TODO - check for errors
	e := secp256k1.SigningMethodES256K.Verify(strings.Join(parts[:2], "."), parts[2], pub)
	return e
}

func HexToECDSAPub(_pub string) (*ecdsa.PublicKey, error) {
	var ok bool

	x := new(big.Int)
	x, ok = x.SetString(_pub[:len(_pub)/2], 16)
	if !ok {
		fmt.Println("SetString: error")
		return nil, fmt.Errorf("invalid hex")
	}

	y := new(big.Int)
	y, ok = y.SetString(_pub[len(_pub)/2:], 16)
	if !ok {
		fmt.Println("SetString: error")
		return nil, fmt.Errorf("invalid hex")
	}

	return &ecdsa.PublicKey{X: x, Y: y, Curve: crypto.S256()}, nil
}
