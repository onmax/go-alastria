package tokens

import (
	"testing"

	"github.com/onmax/go-alastria/tokens"
)

func TestSignToken(t *testing.T) {
	type args struct {
		jwt        interface{}
		privateKey string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Should sign an AIC successfully",
			args: args{jwt: tokens.AIC{
				Header: &tokens.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &tokens.AICPayload{
					IssuedAt:         1,
					PublicKey:        "public-key",
					JSONTokenId:      "json-token-id",
					CreateAlastriaTX: "create-alastria-tx",
					AlastriaToken:    "alastria-token",
					Contexts:         []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:            []string{"AlastriaIdentityCreation", "US12"},
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsInB1YmxpY0tleSI6InB1YmxpYy1rZXkiLCJqdGkiOiJqc29uLXRva2VuLWlkIiwiY3JlYXRlQWxhc3RyaWFUWCI6ImNyZWF0ZS1hbGFzdHJpYS10eCIsImFsYXN0cmlhVG9rZW4iOiJhbGFzdHJpYS10b2tlbiIsIkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2FydGlmYWN0cy92MSJdLCJ0eXBlIjpbIkFsYXN0cmlhSWRlbnRpdHlDcmVhdGlvbiIsIlVTMTIiXX0.Fjjy_24sitUYVpLnsQkqet7nl6zb70wg7A2QvZAkHO0cCZGQzoTOA4CAvTNn2-b0QPR_vgoNJ--PuZlxBEBBBg",
			wantErr: false,
		},
		{
			name: "Should sign an AT successfully",
			args: args{jwt: tokens.AT{
				Header: &tokens.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &tokens.ATPayload{
					IssuedAt:          1,
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken"},
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsInR5cGUiOlsiQWxhc3RyaWFUb2tlbiJdLCJhbnkiOiJhbGFzdHJpYS1uZXR3b3JrLWlkIiwiY2J1IjoiY2FsbGJhY2stdXJsIiwiZ3d1IjoiZ2F0ZXdheS11cmwifQ.FqgrhuoWdAc0niL0o9RnO_yxIqZl5y_TGC-ulZ8j6iZjmz2uL8x_Q6ssY5ZAaOa1XlZBD4ZJWAjC5oyMEnPaPw",
			wantErr: false,
		},
		{
			name: "Should sign an AS successfully",
			args: args{jwt: tokens.AS{
				Header: &tokens.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &tokens.ASPayload{
					IssuedAt:      1,
					Issuer:        "issuer",
					AlastriaToken: "alastria-token",
					Contexts:      []string{"https://alastria.github.io/identity/artifacts/v1"},
					Types:         []string{"AlastriaSession"},
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QifQ.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsImFsYXN0cmlhVG9rZW4iOiJhbGFzdHJpYS10b2tlbiIsIkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2FydGlmYWN0cy92MSJdLCJ0eXBlIjpbIkFsYXN0cmlhU2Vzc2lvbiJdfQ.BOGmAzAAyk3pBWUoffi_F5DV3cqgK1Dz9VIBOLv86gdi0slwF8sYfFBUzCkz7ee1zLUXLTfcFOxbkhIxmNeOSQ",
			wantErr: false,
		},
		{
			name: "Should sign an Credential successfully",
			args: args{jwt: tokens.Credential{
				Header: &tokens.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &tokens.CredentialPayload{
					Issuer:  "issuer",
					Subject: "subject",
					VerifiableCredential: &tokens.CredentialPayloadVC{
						CredentialSubject: &map[string]interface{}{
							"levelOfAssurance": 0,
						},
					},
					IssuedAt: 1,
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsInN1YiI6InN1YmplY3QiLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2NyZWRlbnRpYWxzL3YxIl0sInR5cGUiOlsiVmVyaWZpYWJsZUNyZWRlbnRpYWwiLCJBbGFzdHJpYVZlcmlmaWFibGVDcmVkZW50aWFsIl0sImNyZWRlbnRpYWxTdWJqZWN0Ijp7ImxldmVsT2ZBc3N1cmFuY2UiOjB9fX0.JjmlK87WUgN9OvJLwSV6-hO9LFvb8ntPv6jU44sfsw8UuyLJYsap5EwQqKHrbrZncqhA_u-nLCKMmdqT6A68ZA",
			wantErr: false,
		},
		{
			name: "Should sign an Presentation successfully",
			args: args{jwt: tokens.Presentation{
				Header: &tokens.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &tokens.PresentationPayload{
					Issuer:   "issuer",
					Audience: "audience",
					VerifiablePresentation: &tokens.PresentationPayloadVP{
						ProcessHash:           "process-hash",
						ProcessUrl:            "process-url",
						VerifiableCredentials: []string{"credential_1", "credential_2"},
					},
					IssuedAt: 1,
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsImF1ZCI6ImF1ZGllbmNlIiwidnAiOnsiQGNvbnRleHQiOlsiaHR0cHM6Ly93d3cudzMub3JnLzIwMTgvY3JlZGVudGlhbHMvdjEiLCJodHRwczovL2FsYXN0cmlhLmdpdGh1Yi5pby9pZGVudGl0eS9jcmVkZW50aWFscy92MSJdLCJ0eXBlIjpbIlZlcmlmaWFibGVQcmVzZW50YXRpb24iLCJBbGFzdHJpYVZlcmlmaWFibGVQcmVzZW50YXRpb24iXSwicHJvY0hhc2giOiJwcm9jZXNzLWhhc2giLCJwcm9jVXJsIjoicHJvY2Vzcy11cmwiLCJ2ZXJpZmlhYmxlQ3JlZGVudGlhbCI6WyJjcmVkZW50aWFsXzEiLCJjcmVkZW50aWFsXzIiXX19.-Ad4f6RaODDBXH_UxFSSs8RwuDyAdNt_Ilu3xgeQliUdl8t7x1HMriT-7_j3dMVoLZlfcx4pVxx_YLgTP_74ug",
			wantErr: false,
		},
		{
			name: "Should sign an PresentationRequest successfully",
			args: args{jwt: tokens.PR{
				Header: &tokens.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &tokens.PRPayload{
					IssuedAt:    1,
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					VerifiablePresentation: &tokens.PRPayloadVP{
						ProcessHash: "process-hash",
						ProcessUrl:  "process-url",
						Data: &[]tokens.PRPayloadVPData{{
							Contexts:         []string{"context-1"},
							LevelOfAssurance: 0,
							Required:         true,
							FieldName:        "credential_field_name_1",
						}},
					},
				},
			},
				privateKey: "8dd5e4759a2d8a0ebeb1a1c53b113dad453daaf140b47c419eeef11a46d656b3",
			},
			want:    "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsImNidSI6ImNhbGxiYWNrLXVybCIsInZwIjp7IkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2NyZWRlbnRpYWxzL3YxIl0sInR5cGUiOlsiQWxhc3RyaWFWZXJpZmlhYmxlUHJlc2VudGF0aW9uUmVxdWVzdCJdLCJwcm9jSGFzaCI6InByb2Nlc3MtaGFzaCIsInByb2NVcmwiOiJwcm9jZXNzLXVybCIsImRhdGEiOlt7IkBjb250ZXh0IjpbImNvbnRleHQtMSJdLCJyZXF1aXJlZCI6dHJ1ZSwiZmllbGRfbmFtZSI6ImNyZWRlbnRpYWxfZmllbGRfbmFtZV8xIn1dfX0.K3kl8XFWmmjYF4Te0hjyBRaqCKXKrvosLC4ZMiM6kBFAeQaIkiYUPiFyh7hjsY4HVPfFfVx8PKupQ3Iyx531uw",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := SignToken(tt.args.jwt, tt.args.privateKey)

			if (err != nil) != tt.wantErr {
				t.Errorf("SignToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SignToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestVerify(t *testing.T) {
	type args struct {
		signed string
		_pk    string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "Should verify a JWT successfully",
			args:    args{signed: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsInB1YmxpY0tleSI6InB1YmxpYy1rZXkiLCJqdGkiOiJqc29uLXRva2VuLWlkIiwiY3JlYXRlQWxhc3RyaWFUWCI6ImNyZWF0ZS1hbGFzdHJpYS10eCIsImFsYXN0cmlhVG9rZW4iOiJhbGFzdHJpYS10b2tlbiIsIkBjb250ZXh0IjpbImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2FydGlmYWN0cy92MSJdLCJ0eXBlIjpbIkFsYXN0cmlhSWRlbnRpdHlDcmVhdGlvbiIsIlVTMTIiXX0.Fjjy_24sitUYVpLnsQkqet7nl6zb70wg7A2QvZAkHO0cCZGQzoTOA4CAvTNn2-b0QPR_vgoNJ--PuZlxBEBBBg", _pk: "61a909c125d095e3a4f67125144427312ea6a42fb6321e94f41b25921e106bb8a5db57cc767c179dfe14fa959b615b2cf60852430867ce4675d18a1b3aa032b8"},
			wantErr: false,
		},
		{
			name:    "Should verify a JWT successfully signed from Java",
			args:    args{signed: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqd2siLCJraWQiOiJraWQifQ.eyJleHAiOjAsIm5iZiI6MCwianRpIjoianRpIiwiaWF0IjowLCJ0eXBlIjpbInR5cGVzIl0sImNyZWF0ZUFsYXN0cmlhVFgiOiJjcmVhdGVBbGFzdHJpYVRYIiwiYWxhc3RyaWFUb2tlbiI6ImFsYXN0cmlhVG9rZW4iLCJwdWJsaWNLZXkiOiJwdWJsaWNLZXkiLCJAY29udGV4dCI6WyJjb250ZXh0Il19.KTmgyC3rhqoOk9QOTjQsVLku-r6f5JqDF1UAldu813RdJ63UOPc5f2BKryMB-mdxxuh1WbspsB4zRYTb_ALzVw", _pk: "61a909c125d095e3a4f67125144427312ea6a42fb6321e94f41b25921e106bb8a5db57cc767c179dfe14fa959b615b2cf60852430867ce4675d18a1b3aa032b8"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := VerifyToken(tt.args.signed, tt.args._pk); (err != nil) != tt.wantErr {
				t.Errorf("Verify() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
