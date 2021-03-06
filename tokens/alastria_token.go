package tokens

import alaTypes "github.com/onmax/go-alastria/types"

type AT struct {
	Header  *alaTypes.Header `json:"header,omitempty"`
	Payload *ATPayload       `json:"payload,omitempty"`
}

type ATPayload struct {
	// ! AT lacks of `@context` field
	JSONTokenId               string   `json:"jti,omitempty"`
	IssuedAt                  uint64   `json:"iat,omitempty"`
	ExpiresAt                 uint64   `json:"exp,omitempty"`
	NotBefore                 uint64   `json:"nbf,omitempty"`
	Issuer                    string   `json:"iss,omitempty"`
	Types                     []string `json:"type,omitempty"`
	AlastriaNetworkId         string   `json:"ani,omitempty"`
	CallbackURL               string   `json:"cbu,omitempty"`
	GatewayURL                string   `json:"gwu,omitempty"`
	MultiFactorAuthentication string   `json:"mfau,omitempty"` // ! probably better with just 3 letter "mfa" as the rest of properties
}

var defaultATType = [...]string{"AlastriaToken"}
var validATTypes = append([]string{"US12", "US211", "US221", "US231", "US142"}, defaultATType[:]...)

// Validates the AlastriaToken according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#01-alastria-token-at
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: AlastriaToken.GatewayURL, AlastriaToken.Issuer, AlastriaToken.CallbackURL, AlastriaToken.AlastriaNetworkId
func CreateAlastriaToken(header *alaTypes.Header, payload *ATPayload) (*AT, error) {
	if err := ValidateHeader(header); err != nil {
		return nil, err
	}

	if err := ValidateATPayload(payload); err != nil {
		return nil, err
	}

	return &AT{
			Header:  header,
			Payload: payload},
		nil
}

// Decodes an AlastriaToken from a signed JWT
func DecodeAlastriaToken(signedAT string) (*AT, error) {
	header64, payload64, _, err := SplitJWT(signedAT)
	if err != nil {
		return nil, err
	}
	jwt, err := b64ToJwt(header64, payload64, "AT")
	if err != nil {
		return nil, err
	}
	at := &jwt.AlastriaToken

	errH := ValidateHeader(at.Header)
	errP := ValidateATPayload(at.Payload)

	if errH != nil || errP != nil {
		return nil, err
	}
	return at, nil
}

// Validates the AlastriaToken according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#01-alastria-token-at
// Sets default values if they are empty and they are required
func ValidateATPayload(payload *ATPayload) error {
	mandatoryValues := map[string]string{
		"Issuer":            payload.Issuer,
		"GatewayURL":        payload.GatewayURL,
		"CallbackURL":       payload.CallbackURL,
		"AlastriaNetworkId": payload.AlastriaNetworkId,
	}

	if err := checkMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := validateEnum(payload.Types, validATTypes, "Types"); err != nil {
		return err
	}
	// TODO validATTypes should be a oneOfs

	payload.Types = addDefaultValues(payload.Types, defaultATType[:])

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	return nil
}
