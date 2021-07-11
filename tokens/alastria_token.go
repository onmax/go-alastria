package tokens

type AT struct {
	Header  *Header    `json:"header"`
	Payload *ATPayload `json:"payload"`
}

type ATPayload struct {
	// ! AT lacks of `@context` field
	JSONTokenId               string   `json:"jti"`
	IssuedAt                  uint64   `json:"iat"`
	ExpiresAt                 uint64   `json:"exp"`
	NotBefore                 uint64   `json:"nbf"`
	Issuer                    string   `json:"iss"`
	Types                     []string `json:"type"`
	AlastriaNetworkId         string   `json:"any"`
	CallbackURL               string   `json:"cbu"`
	GatewayURL                string   `json:"gwu"`
	MultiFactorAuthentication string   `json:"mfau"` // ! probably better with just 3 letter "mfa" as the rest of properties
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
func CreateAlastriaToken(header *Header, payload *ATPayload) (*AT, error) {
	err := ValidateAlastriaToken(header, payload)

	if err != nil {
		return nil, err
	}

	return &AT{
			Header:  header,
			Payload: payload},
		nil
}

// Validates the AlastriaToken according to the specification
// https://github.com/alastria/alastria-identity/wiki/Artifacts-and-User-Stories-Definitions#01-alastria-token-at
// Sets default values if they are empty and they are required
func ValidateAlastriaToken(header *Header, payload *ATPayload) error {

	if err := ValidateHeader(header); err != nil {
		return err
	}

	mandatoryValues := map[string]string{
		"Issuer":            payload.Issuer,
		"GatewayURL":        payload.GatewayURL,
		"CallbackURL":       payload.CallbackURL,
		"AlastriaNetworkId": payload.AlastriaNetworkId,
	}

	if err := CheckMandatoryStringFieldsAreNotEmpty(mandatoryValues); err != nil {
		return err
	}

	// ! In the docs, it is not explicitly mentioned that the payload.Types can have more values that the ones that they mention
	if err := ValidateEnum(payload.Types, validATTypes, "Types"); err != nil {
		return err
	}
	// TODO validATTypes should be a oneOfs

	payload.Types = AddDefaultValues(payload.Types, defaultATType[:])

	if err := ValidateTimestamps(&payload.IssuedAt, &payload.ExpiresAt, &payload.NotBefore); err != nil {
		return err
	}

	return nil
}
