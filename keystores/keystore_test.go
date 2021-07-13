package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"

	"github.com/onmax/go-alastria/secp256k1"
	"github.com/onmax/go-alastria/tokens"
)

var tempFolder string = "test-data__temp"

func TestCreateKs(t *testing.T) {
	type args struct {
		folder   string
		password string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "CreateKeyStore",
			args: args{
				folder:   tempFolder,
				password: "test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ks, err := CreateKs(tt.args.folder, tt.args.password)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateKs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(ks.Address.Hex()) != 42 {
				t.Errorf("CreateKs() = %v, want something with 42 chars", ks.Address.Hex())
			}

			// remove the folder
			// os.RemoveAll(tt.args.folder)
		})
	}
}

func TestImportKs(t *testing.T) {

	t.Run("Creates and imports the keystore", func(t *testing.T) {
		_, err := CreateKs("test-data", "password")
		if err != nil {
			t.Errorf("CreateKs() error = %v", err)
			return
		}

		files, err := ioutil.ReadDir("./test-data")
		if err != nil {
			t.Errorf("ImportKs() error = %v", err)
			return
		}

		fmt.Printf("%s", "./test-data/")
		_, pubKey, err := ImportKs("./test-data/"+files[0].Name(), "password")
		if err != nil {
			t.Errorf("ImportKs() error = %v", err)
			return
		}
		if len(pubKey) != 42 {
			t.Errorf("ImportKs() = %v, want something with 42 chars", pubKey)
		}

		os.RemoveAll("./test-data")
	})

	t.Run("Signs and verifies a JWT with same keystore", func(t *testing.T) {
		privKey, pubKey, err := ImportKs("./test-keystore/keystore1.json", "password")
		if err != nil {
			t.Errorf("Sign() error = %v", err)
			return
		}
		at := tokens.AT{
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
		}

		signed, err := secp256k1.Sign(at, privKey)
		if err != nil {
			t.Errorf("Sign() error = %v", err)
			return
		}

		err = secp256k1.Verify(signed, pubKey)
		if err != nil {
			t.Errorf("Verify() error = %v", err)
			return
		}
	})
}
