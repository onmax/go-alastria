package tokens

import alaTypes "github.com/onmax/go-alastria/types"

type AS struct {
	Header  *alaTypes.Header `json:"header,omitempty"`
	Payload *ASPayload       `json:"payload,omitempty"`
}

type ASPayload struct {
	JSONTokenId   string   `json:"jti,omitempty"`
	IssuedAt      uint64   `json:"iat,omitempty"`
	ExpiresAt     uint64   `json:"exp,omitempty"`
	NotBefore     uint64   `json:"nbf,omitempty"`
	Issuer        string   `json:"iss,omitempty"`
	AlastriaToken string   `json:"alastriaToken,omitempty"`
	Contexts      []string `json:"@context,omitempty"`
	Types         []string `json:"type,omitempty"`
}

var defaultASContextURLs = [...]string{"https://alastria.github.io/identity/artifacts/v1"}
var defaultASType = [...]string{"AlastriaSession"}
var validASTypes = append([]string{"US211", "US221", "US142"}, defaultASType[:]...)

// Validates the AlastriaSession according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#02-alastria-session-as
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.AlastriaToken and AlastriaToken.Issuer
func CreateAlastriaSession(header *alaTypes.Header, payload *ASPayload) (*AS, error) {
	if err := ValidateHeader(header); err != nil {
		return nil, err
	}

	if err := ValidateASPayload(payload); err != nil {
		return nil, err
	}

	return &AS{
			Header:  header,
			Payload: payload},
		nil
}

// Decodes an AlastriaSession from a signed JWT
func DecodeAlastriaSession(signedAS string) (*AS, error) {
	header64, payload64, _, err := SplitJWT(signedAS)
	if err != nil {
		return nil, err
	}
	jwt, err := b64ToJwt(header64, payload64, "AS")
	if err != nil {
		return nil, err
	}
	as := &jwt.AlastriaSession

	errH := ValidateHeader(as.Header)
	errP := ValidateASPayload(as.Payload)

	if errH != nil || errP != nil {
		return nil, err
	}
	return as, nil
}

// Validates the AlastriaSession according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#02-alastria-session-as
// Sets default values if they are empty and they are required
func ValidateASPayload(payload *ASPayload) error {
	mandatoryValues := map[string]string{
		"AlastriaToken": payload.AlastriaToken,
		"Issuer":        payload.Issuer,
	}

	if err := checkMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := validateEnum(payload.Types, validASTypes, "Types"); err != nil {
		return err
	}
	// TODO validASTypes should be a oneOfs

	// ! In the docs, it is not explicitly mentioned that the payload.Context can have more values that the ones that they mention
	if err := validateEnum(payload.Contexts, defaultASContextURLs[:], "Context"); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.Types = addDefaultValues(payload.Types, defaultASType[:])
	payload.Contexts = addDefaultValues(payload.Contexts, defaultASContextURLs[:])

	return nil
}
