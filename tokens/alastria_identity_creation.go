package tokens

import alaTypes "github.com/onmax/go-alastria/types"

type AIC struct {
	Header  *alaTypes.Header `json:"header,omitempty"`
	Payload *AICPayload      `json:"payload,omitempty"`
}

type AICPayload struct {
	IssuedAt         uint64   `json:"iat,omitempty"`
	ExpiresAt        uint64   `json:"exp,omitempty"`
	NotBefore        uint64   `json:"nbf,omitempty"`
	PublicKey        string   `json:"publicKey,omitempty"`
	JSONTokenId      string   `json:"jti,omitempty"`
	CreateAlastriaTX string   `json:"createAlastriaTX,omitempty"` // ! Not the best name
	AlastriaToken    string   `json:"alastriaToken,omitempty"`
	Contexts         []string `json:"@context,omitempty"`
	Types            []string `json:"type,omitempty"`
}

var defaultAICContextURLs = [...]string{"https://alastria.github.io/identity/artifacts/v1"}
var defaultAICType = [...]string{"AlastriaIdentityCreation"}
var validAICTypes = append([]string{"US12"}, defaultAICType[:]...)

// Validates the AIC according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#03-alastria-identity-creation-aic
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.AlastriaToken, payload.CreateAlastriaTX and payload.PublicKey
func CreateAlastriaIdentityCreation(header *alaTypes.Header, payload *AICPayload) (*AIC, error) {
	if err := ValidateHeader(header); err != nil {
		return nil, err
	}

	if err := ValidateAICPayload(payload); err != nil {
		return nil, err
	}

	return &AIC{
			Header:  header,
			Payload: payload},
		nil
}

// Decodes an AIC from a signed JWT
func DecodeAIC(signedAIC string) (*AIC, error) {
	header64, payload64, _, err := SplitJWT(signedAIC)
	if err != nil {
		return nil, err
	}
	jwt, err := b64ToJwt(header64, payload64, "AIC")
	if err != nil {
		return nil, err
	}
	aic := &jwt.AIC

	errH := ValidateHeader(aic.Header)
	errP := ValidateAICPayload(aic.Payload)

	if errH != nil || errP != nil {
		return nil, err
	}
	return aic, nil
}

// Validates the AIC according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#03-alastria-identity-creation-aic
// Sets default values if they are empty and they are required
func ValidateAICPayload(payload *AICPayload) error {

	mandatoryValues := map[string]string{
		"AlastriaToken":    payload.AlastriaToken,
		"CreateAlastriaTX": payload.CreateAlastriaTX,
		"PublicKey":        payload.PublicKey,
	}

	if err := checkMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := validateEnum(payload.Types, validAICTypes, "Types"); err != nil {
		return err
	}
	// TODO validAICTypes should be a oneOfs

	// ! In the docs, it is not explicitly mentioned that the payload.Contexts can have more values that the ones that they mention
	if err := validateEnum(payload.Contexts, defaultAICContextURLs[:], "Context"); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.Types = addDefaultValues(payload.Types, defaultAICType[:])
	payload.Contexts = addDefaultValues(payload.Contexts, defaultAICContextURLs[:])

	// TODO Validates contexts are URLs
	// TODO CreateAlastria, TXPublicKey and JSONTokenId can be a regex?

	return nil
}
