package tokens

import (
	"encoding/json"
	"reflect"
	"strings"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestCreatePresentationRequest(t *testing.T) {
	type args struct {
		header  *alaTypes.Header
		payload *PRPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *PR
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
				payload: &PRPayload{
					JSONTokenId: "json-token-id",
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					IssuedAt:    _CurrentTimestamp,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						Contexts:           []string{"https://alastria.github.io/identity/credentials/v1"},
						Types:              []string{"AlastriaVerifiablePresentationRequest"},
						ProcessHash:        "process-hash",
						ProcessUrl:         "process-url",
						ProcessDescription: "process-description",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 1,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}, {
							Contexts:         []string{"context-2"},
							LevelOfAssurance: 3,
							Required:         false,
							FieldName:        "credential_field_name_2",
						}},
					},
				},
			},
			want: &PR{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &PRPayload{
					JSONTokenId: "json-token-id",
					ExpiresAt:   TimestampInTheFuture,
					NotBefore:   _CurrentTimestamp,
					IssuedAt:    _CurrentTimestamp,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						Contexts:           []string{"https://alastria.github.io/identity/credentials/v1"},
						Types:              []string{"AlastriaVerifiablePresentationRequest"},
						ProcessHash:        "process-hash",
						ProcessUrl:         "process-url",
						ProcessDescription: "process-description",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 1,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}, {
							Contexts:         []string{"context-2"},
							LevelOfAssurance: 3,
							Required:         false,
							FieldName:        "credential_field_name_2",
						}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test with optional values",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					IssuedAt:    _CurrentTimestamp,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			want: &PR{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &PRPayload{
					IssuedAt:    _CurrentTimestamp,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						Contexts:    []string{"https://alastria.github.io/identity/credentials/v1"},
						Types:       []string{"AlastriaVerifiablePresentationRequest"},
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test invalid levelOfAssurance",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					IssuedAt:    _CurrentTimestamp,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 4,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "levelOfAssurance is invalid. Only 0, 1, 2, 3 are valid",
		},
		{
			name: "test without Issuer which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Issuer is empty and it is mandatory",
		},
		{
			name: "test without CallbackURL which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer: "issuer",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value CallbackURL is empty and it is mandatory",
		},
		{
			name: "test without VerifiablePresentation which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiablePresentation is empty and it is mandatory",
		},
		{
			name: "test VerifiablePresentation equal to nil which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:                 "issuer",
					CallbackURL:            "callback-url",
					VerifiablePresentation: nil,
				},
			},
			wantErr:          true,
			errorMsgContains: "the value VerifiablePresentation is empty and it is mandatory",
		},
		{
			name: "test without ProcessHash which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessUrl: "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value ProcessHash is empty and it is mandatory",
		},
		{
			name: "test without ProcessUrl which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value ProcessUrl is empty and it is mandatory",
		},
		{
			name: "test without Data which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Data is empty and it is mandatory",
		},
		{
			name: "test Data equal to nil which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				payload: &PRPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data:        nil,
					},
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Data is empty and it is mandatory",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreatePresentationRequest(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreatePresentationRequest() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreatePresentationRequest() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}

func TestDecodePR(t *testing.T) {
	type args struct {
		signedPr string
	}
	tests := []struct {
		name    string
		args    args
		want    *PR
		wantErr bool
	}{
		{
			name: "Should decode an Presentation successfully",
			args: args{
				signedPr: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsImNidSI6ImNhbGxiYWNrLXVybCIsInZwIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2NyZWRlbnRpYWxzL3YxIl0sInR5cGUiOlsiQWxhc3RyaWFWZXJpZmlhYmxlUHJlc2VudGF0aW9uUmVxdWVzdCJdLCJwcm9jSGFzaCI6InByb2Nlc3MtaGFzaCIsInByb2NVcmwiOiJwcm9jZXNzLXVybCIsImRhdGEiOlt7IkBjb250ZXh0IjpbImNvbnRleHQtMSJdLCJyZXF1aXJlZCI6dHJ1ZSwiZmllbGRfbmFtZSI6ImNyZWRlbnRpYWxfZmllbGRfbmFtZV8xIn1dfX0.K3kl8XFWmmjYF4Te0hjyBRaqCKXKrvosLC4ZMiM6kBFAeQaIkiYUPiFyh7hjsY4HVPfFfVx8PKupQ3Iyx531uw",
			},
			want: &PR{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &PRPayload{
					IssuedAt:    1,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
						Contexts: []string{"https://alastria.github.io/identity/credentials/v1"},
						Types:    []string{"AlastriaVerifiablePresentationRequest"},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodePR(tt.args.signedPr)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodePR() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodePR() = %v, want %v", got, tt.want)
			}
		})
	}
}
