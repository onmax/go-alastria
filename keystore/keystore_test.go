package keystore

import (
	"io/ioutil"
	"os"
	"testing"
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
				os.RemoveAll(tt.args.folder)
				t.Errorf("CreateKs() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if len(ks.Account.Address.Hex()) != 42 {
				os.RemoveAll(tt.args.folder)
				t.Errorf("CreateKs() = %v, want something with 42 chars", ks.Account.Address.Hex())
			}

			// remove the folder
			os.RemoveAll(tt.args.folder)
		})
	}
}

func TestImportKs(t *testing.T) {

	t.Run("Creates and imports the keystore", func(t *testing.T) {
		_, err := CreateKs(tempFolder, "password")
		if err != nil {
			t.Errorf("CreateKs() error = %v", err)
			return
		}

		files, err := ioutil.ReadDir(tempFolder)
		if err != nil {
			t.Errorf("ImportKs() error = %v", err)
			return
		}

		ks, err := ImportKs(tempFolder+"/"+files[0].Name(), "password")
		if err != nil {
			t.Errorf("ImportKs() error = %v", err)
			return
		}
		pubKey := ks.HexPublicKey
		if len(pubKey) != 128 {
			t.Errorf("ImportKs() = %v, want something with 128 chars", pubKey)
		}

		os.RemoveAll(tempFolder)
	})

	t.Run("Signs and verifies a JWT with same keystore", func(t *testing.T) {
		_, err := ImportKs("../assets/keystores/keystore1.json", "test")
		if err != nil {
			t.Errorf("Sign() error = %v", err)
			return
		}
	})
}
