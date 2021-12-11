package crypto

import (
	"fmt"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/client"
	"github.com/onmax/go-alastria/tokens"
	alaTypes "github.com/onmax/go-alastria/types"
	secp256k1 "github.com/ureeves/jwt-go-secp256k1"
)

// Signs a JWT
func Sign(conn *alaTypes.Connection, jwt interface{}) (string, error) {
	err := client.CheckKeystore(conn)
	if err != nil {
		return "", err
	}

	h64, p64, err := tokens.JwtToB64(jwt)
	if err != nil {
		return "", err
	}

	pk, err := crypto.HexToECDSA(conn.Client.Ks.HexPrivateKey)
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
