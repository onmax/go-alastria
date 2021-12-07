package tokens

import (
	b64 "encoding/base64"
	"encoding/json"
	"fmt"
	"strings"

	alaTypes "github.com/onmax/go-alastria/types"
)

func B64ToInterface(input string, output interface{}) error {
	r, err := b64.RawStdEncoding.Strict().DecodeString(input)
	if err != nil {
		return err
	}
	err = json.Unmarshal(r, output)
	if err != nil {
		return err
	}
	return nil
}

type AlastriaJWT struct {
	AlastriaToken   AT
	AlastriaSession AS
	AIC             AIC
	Credential      Credential
	Presentation    Presentation
	PR              PR
}

func b64ToJwt(header64 string, payload64 string, jwtType string) (AlastriaJWT, error) {
	if header64 == "" || payload64 == "" {
		return AlastriaJWT{}, alaTypes.ErrInvalidJWT
	}

	var header alaTypes.Header
	err := B64ToInterface(header64, &header)
	if err != nil {
		return AlastriaJWT{}, err
	}

	switch jwtType {
	case "AT":
		var atp ATPayload
		err := B64ToInterface(payload64, &atp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		at := AT{
			Header:  &header,
			Payload: &atp,
		}
		return AlastriaJWT{AlastriaToken: at}, nil
	case "AS":
		var asp ASPayload
		err := B64ToInterface(payload64, &asp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		as := AS{
			Header:  &header,
			Payload: &asp,
		}
		return AlastriaJWT{AlastriaSession: as}, nil
	case "AIC":
		var aicp AICPayload
		err := B64ToInterface(payload64, &aicp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		aic := AIC{
			Header:  &header,
			Payload: &aicp,
		}
		return AlastriaJWT{AIC: aic}, nil
	case "Credential":
		var credentialp CredentialPayload
		err := B64ToInterface(payload64, &credentialp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		credential := Credential{
			Header:  &header,
			Payload: &credentialp,
		}
		return AlastriaJWT{Credential: credential}, nil
	case "Presentation":
		var presentationp PresentationPayload
		err := B64ToInterface(payload64, &presentationp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		presentation := Presentation{
			Header:  &header,
			Payload: &presentationp,
		}
		return AlastriaJWT{Presentation: presentation}, nil
	case "PR":
		var prp PRPayload
		err := B64ToInterface(payload64, &prp)
		if err != nil {
			return AlastriaJWT{}, err
		}
		pr := PR{
			Header:  &header,
			Payload: &prp,
		}
		return AlastriaJWT{PR: pr}, nil
	}

	return AlastriaJWT{}, alaTypes.ErrInvalidJWT
}

func InterfaceToB64(artifact interface{}) (string, error) {
	json, err := json.Marshal(artifact)
	if err != nil {
		return "", err
	}
	json64 := strings.Trim(b64.StdEncoding.EncodeToString([]byte(json)), "=")
	return json64, nil
}

func JwtToB64(jwt interface{}) (string, string, error) {
	if jwt == nil {
		return "", "", alaTypes.ErrInvalidJWT
	}

	var headerError, payloadErr error
	var header *alaTypes.Header
	var payload interface{}

	switch token := jwt.(type) {
	case AIC:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidateAICPayload(token.Payload)
	case AS:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidateASPayload(token.Payload)
	case AT:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidateATPayload(token.Payload)
	case Credential:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidateCredentialPayload(token.Payload)
	case Presentation:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidatePresentationPayload(token.Payload)
	case PR:
		header = token.Header
		payload = token.Payload
		payloadErr = ValidatePRPayload(token.Payload)
	default:
		return "", "", fmt.Errorf("unsupported jwt")
	}

	headerError = ValidateHeader(header)
	if headerError != nil {
		return "", "", payloadErr
	} else if payloadErr != nil {
		return "", "", payloadErr
	}

	h64, err := InterfaceToB64(header)
	if err != nil {
		return "", "", err
	}
	p64, err := InterfaceToB64(payload)
	if err != nil {
		return "", "", err
	}
	return h64, p64, nil
}

func SplitJWT(signedJWT string) (string, string, string, error) {
	parts := strings.Split(signedJWT, ".")

	if len(parts) != 3 {
		return "", "", "", alaTypes.ErrInvalidJWT
	}

	return parts[0], parts[1], parts[2], nil
}
