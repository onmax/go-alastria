package tokens

import "fmt"

type Presentation struct {
	Header  *Header              `json:"header"`
	Payload *PresentationPayload `json:"payload"`
}

type PresentationPayload struct {
	JSONTokenId                    string                 `json:"jti"`
	IssuedAt                       uint64                 `json:"iat"`
	ExpiresAt                      uint64                 `json:"exp"`
	NotBefore                      uint64                 `jsonPresentation:"nbf"`
	Issuer                         string                 `json:"iss"`
	Audience                       string                 `json:"aud"`
	PresentationRequestJSONTokenId string                 `json:"jtipr"`
	VerifiablePresentation         *PresentationPayloadVP `json:"vp"`
}

type PresentationPayloadVP struct {
	Contexts              []string `json:"@context"`
	Types                 []string `json:"type"`
	ProcessHash           string   `json:"procHash"`
	ProcessUrl            string   `json:"procUrl"`
	ProcessDescription    string   `json:"procDescription"`
	VerifiableCredentials []string `json:"verifiableCredential"` // ! Should be plural: 'verifiableCredentials'
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
	err := ValidatePresentation(header, payload)

	if err != nil {
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
func ValidatePresentation(header *Header, payload *PresentationPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

	if payload.VerifiablePresentation == nil {
		return fmt.Errorf(emptyPayloadField, "VerifiablePresentation")
	}

	mandatoryStringValues := map[string]string{
		"Issuer":      payload.Issuer,
		"Audience":    payload.Audience,
		"ProcessHash": payload.VerifiablePresentation.ProcessHash,
		"ProcessUrl":  payload.VerifiablePresentation.ProcessUrl,
	}

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryStringValues); err != nil {
		return err
	}

	mandatoryStringArraysValues := map[string][]string{
		"VerifiableCredentials": payload.VerifiablePresentation.VerifiableCredentials,
	}

	if err := CheckMandatoryStringArrayFieldsNotEmpty(mandatoryStringArraysValues); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.VerifiablePresentation.Types = AddDefaultValues(payload.VerifiablePresentation.Types, defaultPresentationTypes[:])
	payload.VerifiablePresentation.Contexts = AddDefaultValues(payload.VerifiablePresentation.Contexts, defaultPresentationContextURLs[:])

	// TODO Check if payload.VerifiableCredential.credentialSubject is set and has proper values
	// TODO Check iss and sub are valid
	return nil
}
