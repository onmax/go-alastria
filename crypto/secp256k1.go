package crypto

// TODO Rename this file to tokens or something
import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/tokens"
	secp256k1 "github.com/ureeves/jwt-go-secp256k1"
)

// Signs a JWT
func Sign(jwt interface{}, _pk string) (string, error) {
	h64, p64, err := tokens.JwtToB64(jwt)
	if err != nil {
		return "", err
	}

	pk, err := crypto.HexToECDSA(_pk)
	if err != nil {
		return "", err
	}

	s, err := secp256k1.SigningMethodES256K.Sign(h64+"."+p64, pk)
	if err != nil {
		return "", err
	}

	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s.%s.%s", h64, p64, s), nil
}

// Verifies a signature of a JWT
func Verify(signed, _pub string) error {
	header, payload, signature, err := tokens.SplitJWT(signed)
	if err != nil {
		return err
	}
	pub, err := HexToECDSAPub(_pub)
	if err != nil {
		return err
	}
	return secp256k1.SigningMethodES256K.Verify(header+"."+payload, signature, pub)
}
