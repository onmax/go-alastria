package tokens

import "fmt"

type Presentation struct {
	Header  *Header              `json:"header,omitempty"`
	Payload *PresentationPayload `json:"payload,omitempty"`
}

type PresentationPayload struct {
	JSONTokenId                    string                 `json:"jti,omitempty"`
	IssuedAt                       uint64                 `json:"iat,omitempty"`
	ExpiresAt                      uint64                 `json:"exp,omitempty"`
	NotBefore                      uint64                 `json:"nbf,omitempty"`
	Issuer                         string                 `json:"iss,omitempty"`
	Audience                       string                 `json:"aud,omitempty"`
	PresentationRequestJSONTokenId string                 `json:"jtipr,omitempty"`
	VerifiablePresentation         *PresentationPayloadVP `json:"vp,omitempty"`
}

type PresentationPayloadVP struct {
	Contexts              []string `json:"@context,omitempty"`
	Types                 []string `json:"type,omitempty"`
	ProcessHash           string   `json:"procHash,omitempty"`
	ProcessUrl            string   `json:"procUrl,omitempty"`
	ProcessDescription    string   `json:"procDescription,omitempty"`
	VerifiableCredentials []string `json:"verifiableCredential,omitempty"` // ! Should be plural: 'verifiableCredentials'
}

var defaultPresentationTypes = [2]string{"VerifiablePresentation", "AlastriaVerifiablePresentation"}
var defaultPresentationContextURLs = [2]string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"}

// Validates the VerifiablePresentation(aka Presentation) according to the specification
// Header: https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#0-artifacts-definition
// Payload: https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#4-presentation
// The validation with timestamp will be done with the machine timestamp. This can cause a problem, if the time is not syncronize.
// Sets default values if they are empty and they are required
// Returns an error if a mandatory field is empty
// Mandatory fields are: payload.Issuer, payload.Audience, payload.VerifiablePresentation.ProcessHash,
// payload.VerifiablePresentation.ProcessUrl and payload.VerifiableCredential.VerifiableCredentials
func CreatePresentation(header *Header, payload *PresentationPayload) (*Presentation, error) {
	if err := ValidateHeader(header); err != nil {
		return nil, err
	}

	if err := ValidatePresentationPayload(payload); err != nil {
		return nil, err
	}

	return &Presentation{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the Presentation according to the specification
// https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#4-presentation
// Sets default values if they are empty and they are required
func ValidatePresentationPayload(payload *PresentationPayload) error {

	if payload.VerifiablePresentation == nil {
		return fmt.Errorf(emptyPayloadField, "VerifiablePresentation")
	}

	mandatoryStringValues := map[string]string{
		"Issuer":      payload.Issuer,
		"Audience":    payload.Audience,
		"ProcessHash": payload.VerifiablePresentation.ProcessHash,
		"ProcessUrl":  payload.VerifiablePresentation.ProcessUrl,
	}

	if err := checkMandatoryStringFieldsAreNotEmpty(mandatoryStringValues); err != nil {
		return err
	}

	mandatoryStringArraysValues := map[string][]string{
		"VerifiableCredentials": payload.VerifiablePresentation.VerifiableCredentials,
	}

	if err := checkMandatoryStringArrayFieldsNotEmpty(mandatoryStringArraysValues); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.VerifiablePresentation.Types = addDefaultValues(payload.VerifiablePresentation.Types, defaultPresentationTypes[:])
	payload.VerifiablePresentation.Contexts = addDefaultValues(payload.VerifiablePresentation.Contexts, defaultPresentationContextURLs[:])

	// TODO Check if payload.VerifiableCredential.credentialSubject is set and has proper values
	// TODO Check iss and sub are valid
	return nil
}
