package tokens

import (
	"fmt"

	alaTypes "github.com/onmax/go-alastria/types"
)

type PR struct {
	Header  *alaTypes.Header `json:"header,omitempty"`
	Payload *PRPayload       `json:"payload,omitempty"`
}

type PRPayload struct {
	JSONTokenId            string       `json:"jti,omitempty"`
	IssuedAt               uint64       `json:"iat,omitempty"`
	ExpiresAt              uint64       `json:"exp,omitempty"`
	NotBefore              uint64       `json:"nbf,omitempty"`
	Issuer                 string       `json:"iss,omitempty"`
	CallbackURL            string       `json:"cbu,omitempty"`
	VerifiablePresentation *PRPayloadVP `json:"vp,omitempty"`
}

type PRPayloadVP struct {
	Contexts           []string           `json:"@context,omitempty"`
	Types              []string           `json:"type,omitempty"`
	ProcessHash        string             `json:"procHash,omitempty"`
	ProcessUrl         string             `json:"procUrl,omitempty"`
	ProcessDescription string             `json:"procDescription,omitempty"`
	Data               *[]PRPayloadVPData `json:"data,omitempty"` // ! in a presentation is 'verifiableCredentials', and here data. It's not consistent
}

type PRPayloadVPData struct {
	Contexts         []string `json:"@context,omitempty"` // ! Not type?
	LevelOfAssurance int      `json:"levelOfAssurance,omitempty"`
	Required         bool     `json:"required,omitempty"`
	FieldName        string   `json:"field_name,omitempty"` // ! maybe credential_name is better than field_name
}

var defaultPresentationRequestTypes = [1]string{"AlastriaVerifiablePresentationRequest"}

// ! it is not necessary: VerifiablePresentationRequest

var defaultPresentationRequestContextURLs = [1]string{"https://alastria.github.io/identity/credentials/v1"}

// ! it is not necessary: https://www.w3.org/2018/credentials/v1.

// Validates the VerifiablePresentationRequest(aka PresentationRequest) according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#5-presentation-request
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.Issuer, payload.CallbackURL, payload.VerifiableCredential.ProcessHash,
// payload.VerifiablePresentation.ProcessUrl and payload.VerifiablePresentation.Data
func CreatePresentationRequest(header *alaTypes.Header, payload *PRPayload) (*PR, error) {
	if err := ValidateHeader(header); err != nil {
		return nil, err
	}

	if err := ValidatePRPayload(payload); err != nil {
		return nil, err
	}

	return &PR{
			Header:  header,
			Payload: payload},
		nil
}

// Decodes an PR from a signed JWT
func DecodePR(signedPr string) (*PR, error) {
	header64, payload64, _, err := SplitJWT(signedPr)
	if err != nil {
		return nil, err
	}
	jwt, err := b64ToJwt(header64, payload64, "PR")
	if err != nil {
		return nil, err
	}
	pr := &jwt.PR

	errH := ValidateHeader(pr.Header)
	errP := ValidatePRPayload(pr.Payload)

	if errH != nil || errP != nil {
		return nil, err
	}
	return pr, nil
}

// Validates the PresentationRequest according to the specification
// https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#5-presentation-request
// Sets default values if they are empty and they are required
func ValidatePRPayload(payload *PRPayload) error {
	if payload.VerifiablePresentation == nil {
		return fmt.Errorf(emptyPayloadField, "VerifiablePresentation")
	}

	if payload.VerifiablePresentation.Data == nil {
		return fmt.Errorf(emptyPayloadField, "Data")
	}

	mandatoryStringValues := map[string]string{
		"Issuer":      payload.Issuer,
		"CallbackURL": payload.CallbackURL,
		"ProcessHash": payload.VerifiablePresentation.ProcessHash,
		"ProcessUrl":  payload.VerifiablePresentation.ProcessUrl,
	}

	if err := checkMandatoryStringFieldsAreNotEmpty(mandatoryStringValues); err != nil {
		return err
	}

	if len(*payload.VerifiablePresentation.Data) == 0 {
		return fmt.Errorf(emptyPayloadField, "Data")
	}

	for _, data := range *payload.VerifiablePresentation.Data {
		if data.LevelOfAssurance > 3 {
			return fmt.Errorf("levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid")
		}
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.VerifiablePresentation.Types = addDefaultValues(payload.VerifiablePresentation.Types, defaultPresentationRequestTypes[:])
	payload.VerifiablePresentation.Contexts = addDefaultValues(payload.VerifiablePresentation.Contexts, defaultPresentationRequestContextURLs[:])

	// TODO Check if payload.VerifiableCredential.credentialSubject is set and has proper values
	// TODO Check iss and sub are valid
	// TODO Check that data is valid: levelOfAssurance
	return nil
}
