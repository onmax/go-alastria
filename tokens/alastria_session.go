package tokens

type AS struct {
	Header  *Header    `json:"header"`
	Payload *ASPayload `json:"payload"`
}

type ASPayload struct {
	JSONTokenId   string   `json:"jti"`
	IssuedAt      uint64   `json:"iat"`
	ExpiresAt     uint64   `json:"exp"`
	NotBefore     uint64   `json:"nbf"`
	Issuer        string   `json:"iss"`
	AlastriaToken string   `json:"alastriaToken"`
	Contexts      []string `json:"@context"`
	Types         []string `json:"type"`
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
func CreateAlastriaSession(header *Header, payload *ASPayload) (*AS, error) {
	err := ValidateAlastriaSession(header, payload)

	if err != nil {
		return nil, err
	}

	return &AS{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the AlastriaSession according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#02-alastria-session-as
// Sets default values if they are empty and they are required
func ValidateAlastriaSession(header *Header, payload *ASPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

	mandatoryValues := map[string]string{
		"AlastriaToken": payload.AlastriaToken,
		"Issuer":        payload.Issuer,
	}

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := ValidateEnum(payload.Types, validASTypes, "Types"); err != nil {
		return err
	}
	// TODO validASTypes should be a oneOfs

	// ! In the docs, it is not explicitly mentioned that the payload.Context can have more values that the ones that they mention
	if err := ValidateEnum(payload.Contexts, defaultASContextURLs[:], "Context"); err != nil {
		return err
	}

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	payload.Types = AddDefaultValues(payload.Types, defaultASType[:])
	payload.Contexts = AddDefaultValues(payload.Contexts, defaultASContextURLs[:])

	return nil
}
