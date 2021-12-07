package tokens

import (
	"encoding/json"
	"fmt"
	"reflect"
	"strings"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestCreateAlastriaToken(t *testing.T) {
	type args struct {
		header  *alaTypes.Header
		payload *ATPayload
	}
	tests := []struct {
		name             string
		args             args
		want             *AT
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
				payload: &ATPayload{
					JSONTokenId:       "json-token-id",
					IssuedAt:          _CurrentTimestamp,
					ExpiresAt:         TimestampInTheFuture,
					NotBefore:         _CurrentTimestamp,
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken", "US12"},
				},
			},
			want: &AT{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &ATPayload{
					JSONTokenId:       "json-token-id",
					IssuedAt:          _CurrentTimestamp,
					ExpiresAt:         TimestampInTheFuture,
					NotBefore:         _CurrentTimestamp,
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken", "US12"},
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
				payload: &ATPayload{
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					IssuedAt:          _CurrentTimestamp,
				},
			},
			want: &AT{
				Header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				Payload: &ATPayload{
					IssuedAt:          _CurrentTimestamp,
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken"},
				},
			},
			wantErr: false,
		},
		{
			name: "test without Issuer which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ATPayload{
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value Issuer is empty and it is mandatory",
		},
		{
			name: "test without AlastriaNetworkId which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ATPayload{
					Issuer:      "issuer",
					CallbackURL: "callback-url",
					GatewayURL:  "gateway-url",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value AlastriaNetworkId is empty and it is mandatory",
		},
		{
			name: "test without CallbackURL which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ATPayload{
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					GatewayURL:        "gateway-url",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value CallbackURL is empty and it is mandatory",
		},
		{
			name: "test without GatewayURL which is mandatory",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ATPayload{
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
				},
			},
			wantErr:          true,
			errorMsgContains: "the value GatewayURL is empty and it is mandatory",
		},
		{
			name: "test invalid type",
			args: args{
				header: &alaTypes.Header{
					Algorithm: "ES256K",
					Type:      "JWT",
				},
				payload: &ATPayload{
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken", "US12", "invalid"},
				},
			},
			wantErr:          true,
			errorMsgContains: fmt.Sprintf("the value Types=invalid is invalid. Only the following values are accepted %s", validATTypes),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := CreateAlastriaToken(tt.args.header, tt.args.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateAlastriaToken() error = %v, wantErr %v", err, tt.wantErr)
			} else if (err != nil) && tt.wantErr && !strings.Contains(err.Error(), tt.errorMsgContains) {
				t.Errorf("ValidateTimestamps() unexpected error message: error = %v | should contain: %v", err, tt.errorMsgContains)
			}
			if tt.args.payload.IssuedAt == 0 {
				tt.args.payload.IssuedAt = _CurrentTimestamp
			}
			if !reflect.DeepEqual(got, tt.want) {
				a, _ := json.Marshal(got)
				b, _ := json.Marshal(tt.want)
				t.Errorf("CreateAlastriaToken() -> %s \n\tgot:  %s,\n\twant: %s", tt.name, string(a), string(b))
			}
		})
	}
}

func TestDecodeAlastriaToken(t *testing.T) {
	type args struct {
		signedAT string
	}
	tests := []struct {
		name    string
		args    args
		want    *AT
		wantErr bool
	}{
		{
			name: "Should decode an AT successfully",
			args: args{
				signedAT: "eyJhbGciOiJFUzI1NksiLCJ0eXAiOiJKV1QiLCJqd2siOiJqc29uLXdlYi10b2tlbiIsImtpZCI6ImtleS1pZCJ9.eyJpYXQiOjEsImlzcyI6Imlzc3VlciIsInR5cGUiOlsiQWxhc3RyaWFUb2tlbiJdLCJhbmkiOiJhbGFzdHJpYS1uZXR3b3JrLWlkIiwiY2J1IjoiY2FsbGJhY2stdXJsIiwiZ3d1IjoiZ2F0ZXdheS11cmwifQ.Y8q5Le6LGBxY8BhYZMWe2KJOzWqD_Mroy7PTiqlHsnBJXqLn7je_F8lPjqMv4l9_bdHrdz2eXh3ra3M_XPfJJA",
			},
			want: &AT{
				Header: &alaTypes.Header{
					Algorithm:    "ES256K",
					Type:         "JWT",
					KeyID:        "key-id",
					JSONWebToken: "json-web-token",
				},
				Payload: &ATPayload{
					IssuedAt:          1,
					Issuer:            "issuer",
					AlastriaNetworkId: "alastria-network-id",
					CallbackURL:       "callback-url",
					GatewayURL:        "gateway-url",
					Types:             []string{"AlastriaToken"},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := DecodeAlastriaToken(tt.args.signedAT)
			if (err != nil) != tt.wantErr {
				t.Errorf("DecodeAlastriaToken() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("DecodeAlastriaToken() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestValidateATPayload(t *testing.T) {
	type args struct {
		payload *ATPayload
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateATPayload(tt.args.payload); (err != nil) != tt.wantErr {
				t.Errorf("ValidateATPayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
