package tokens

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestCreateCredential(t *testing.T) {
	type args struct {
		header  *alaTypes.Header
		payload *CredentialPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *Credential
		wantErr          bool
		errorMsgContains string
	}{
		{
			name: "test with all values",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &CredentialPayload{
					JSONTokenId: "json-token-id",
					IssuedAt:    _CurrentTimestamp,
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					Issuer:      "issuer",
					Subject:     "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Contexts: []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:    []string{"VerifiableCredential", "AlastriaVerifiableCredential"},
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
							"data":             map[string]interface{}{"field_name": "field_value"},
						},
					},
				},
			},
			want: &Credential{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &CredentialPayload{
					JSONTokenId: "json-token-id",
					IssuedAt:    _CurrentTimestamp,
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					Issuer:      "issuer",
					Subject:     "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Contexts: []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
						Types:    []string{"VerifiableCredential", "AlastriaVerifiableCredential"},
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
							"data":             map[string]interface{}{"field_name": "field_value"},
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test with optional values",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
					IssuedAt: _CurrentTimestamp,
				},
			},
			want: &Credential{
				Header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &CredentialPayload{
					IssuedAt: _CurrentTimestamp,
					Issuer:   "issuer",
					Subject:  "subject",
					VerifiableCredential: &CredentialPayloadVC{
						Types:    defaultCredentialTypes[:],
						Contexts: defaultCredentialContextURLs[:],
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test invalid levelOfAssurance",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 4,
						},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid",
		},
		{
			name: "test without Subject which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer: "issuer",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Subject is empty and it is mandatory",
		},
		{
			name: "test without VerifiableCredential which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiableCredential is empty and it is mandatory",
		},
		{
			name: "test without CredentialSubject which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &CredentialPayload{
					Issuer:               "issuer",
					Subject:              "subject",
					VerifiableCredential: &CredentialPayloadVC{},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value CredentialSubject is empty and it is mandatory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateCredential(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCredential() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreateCredential() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}

func TestDecodeCredential(t *testing.T) {
	type args struct {
		signedCredential string
	}
	tests := []struct {
		name    string
		args    args
		want    *Credential
		wantErr bool
	}{
		{
			name: "Should decode an Credential successfully",
			args: args{
				signedCredential: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsInN1YiI6InN1YmplY3QiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2NyZWRlbnRpYWxzL3YxIl0sInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJBbGFzdHJpYVZlcmlmaWFibGVDcmVkZW50aWFsIl0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImxldmVsT2ZBc3N1cmFuY2UiOjB9fX0.JjmlK87WUgN9OvJLwSV6-hO9LFvb8ntPv6jU44sfsw8UuyLJYsap5EwQqKHrbrZncqhA_u-nLCKMmdqT6A68ZA",
			},
			want: &Credential{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": float64(0),
						},
						Types:    []string{"VerifiableCredential", "AlastriaVerifiableCredential"},
						Contexts: []string{"https://www.w3.org/2018/credentials/v1", "https://alastria.github.io/identity/credentials/v1"},
					},
					IssuedAt: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeCredential(tt.args.signedCredential)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeCredential() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeCredential() = %v, want %v", got, tt.want)
			}
		})
	}
}
