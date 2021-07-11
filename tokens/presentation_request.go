package tokens

import (
	"fmt"
)

type PR struct {
	Header  *Header    `json:"header"`
	Payload *PRPayload `json:"payload"`
}

type PRPayload struct {
	JSONTokenId            string       `json:"jti"`
	IssuedAt               uint64       `json:"iat"`
	ExpiresAt              uint64       `json:"exp"`
	NotBefore              uint64       `json:"nbf"`
	Issuer                 string       `json:"iss"`
	CallbackURL            string       `json:"cbu"`
	VerifiablePresentation *PRPayloadVP `json:"vp"`
}

type PRPayloadVP struct {
	Contexts           []string           `json:"@context"`
	Types              []string           `json:"type"`
	ProcessHash        string             `json:"procHash"`
	ProcessUrl         string             `json:"procUrl"`
	ProcessDescription string             `json:"procDescription"`
	Data               *[]PRPayloadVPData `json:"data"` // ! in a presentation is 'verifiableCredentials', and here data. It's not consistent
}

type PRPayloadVPData struct {
	Contexts         []string `json:"@context"` // ! Not type?
	LevelOfAssurance int      `json:"levelOfAssurance"`
	Required         bool     `json:"required"`
	FieldName        string   `json:"field_name"` // ! maybe credential_name is better than field_name
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
func CreatePresentationRequest(header *Header, payload *PRPayload) (*PR, error) {
	err := ValidatePresentationRequest(header, payload)

	if err != nil {
		return nil, err
	}

	return &PR{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the PresentationRequest according to the specification
// https://github.com/alastria/alastria-identity/wiki/Alastria-DID-Method-Specification-(Quorum-version)#5-presentation-request
// Sets default values if they are empty and they are required
func ValidatePresentationRequest(header *Header, payload *PRPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

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

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryStringValues); err != nil {
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

	payload.VerifiablePresentation.Types = AddDefaultValues(payload.VerifiablePresentation.Types, defaultPresentationRequestTypes[:])
	payload.VerifiablePresentation.Contexts = AddDefaultValues(payload.VerifiablePresentation.Contexts, defaultPresentationRequestContextURLs[:])

	// TODO Check if payload.VerifiableCredential.credentialSubject is set and has proper values
	// TODO Check iss and sub are valid
	// TODO Check that data is valid: levelOfAssurance
	return nil
}
