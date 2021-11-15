package crypto

import (
	"crypto/ecdsa"
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/onmax/go-alastria/tokens"
	secp256k1 "github.com/ureeves/jwt-go-secp256k1"
)

func interfaceToB64(artifact interface{}) (string, error) {
	json, err := json.Marshal(artifact)
	if err != nil {
		return "", err
	}
	json64 := strings.Trim(b64.StdEncoding.EncodeToString([]byte(json)), "=")
	return json64, nil
}

func jwtToB64(jwt interface{}) (string, string, error) {
	if jwt == nil {
		return "", "", fmt.Errorf("jwt is nil")
	}

	var headerError, payloadErr error
	var header *tokens.Header
	var payload interface{}

	switch token := jwt.(type) {
	case tokens.AIC:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidateAICPayload(token.Payload)
	case tokens.AS:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidateASPayload(token.Payload)
	case tokens.AT:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidateATPayload(token.Payload)
	case tokens.Credential:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidateCredentialPayload(token.Payload)
	case tokens.Presentation:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidatePresentationPayload(token.Payload)
	case tokens.PR:
		header = token.Header
		payload = token.Payload
		payloadErr = tokens.ValidatePRPayload(token.Payload)
	default:
		return "", "", fmt.Errorf("unsupported jwt")
	}

	headerError = tokens.ValidateHeader(header)
	if headerError != nil {
		return "", "", payloadErr
	} else if payloadErr != nil {
		return "", "", payloadErr
	}

	h64, err := interfaceToB64(header)
	if err != nil {
		return "", "", err
	}
	p64, err := interfaceToB64(payload)
	if err != nil {
		return "", "", err
	}
	return h64, p64, nil
}

func Sign(jwt interface{}, _pk string) (string, error) {
	h64, p64, err := jwtToB64(jwt)
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

func Verify(signed, _pub string) error {
	parts := strings.Split(signed, ".")

	// Maybe create a helper function/lib for this
	if len(parts) != 3 {
		return fmt.Errorf("invalid jwt")
	}

	pub, err := hexToECDSAPub(_pub)
	if err != nil {
		return err
	}
	e := secp256k1.SigningMethodES256K.Verify(strings.Join(parts[:2], "."), parts[2], pub)
	return e
}

func hexToECDSAPub(_pub string) (*ecdsa.PublicKey, error) {
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
