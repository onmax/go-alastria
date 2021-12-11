package hex

import (
	"reflect"
	"testing"

	"github.com/ethereum/go-ethereum/common"
)

func TestRemove0x(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no 0x",
			args: args{input: "1234"},
			want: "1234",
		},
		{
			name: "0x",
			args: args{input: "0x123"},
			want: "123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Remove0x(tt.args.input); got != tt.want {
				t.Errorf("Remove0x() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAdd0x(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "no 0x",
			args: args{input: "1234"},
			want: "0x1234",
		},
		{
			name: "0x",
			args: args{input: "0x123"},
			want: "0x123",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Add0x(tt.args.input); got != tt.want {
				t.Errorf("Add0x() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByteArrToHex(t *testing.T) {
	type args struct {
		input []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ByteArrToHex(tt.args.input); got != tt.want {
				t.Errorf("ByteArrToHex() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHexToByteArr(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name    string
		args    args
		want    []byte
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := HexToByteArr(tt.args.input)
			if (err != nil) != tt.wantErr {
				t.Errorf("HexToByteArr() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("HexToByteArr() = %v, want %v", got, tt.want)
			}
		})
	}
}


func TestPublicKeyToAddress(t *testing.T) {
	type args struct {
		publicKeyStr string
	}
	tests := []struct {
		name    string
		args    args
		want    common.Address
		wantErr bool
	}{
		{
			name: "valid public key",
			args: args{
				publicKeyStr: "2e507af01167c98a6accc0cd46ab2a256dd6b6c69ec1c0c28f80fb62e1f7d70233768b0c58dbbdac1fc358b8141c075a520483cf9779e4ea98d13df2833f3767",
			},
			want:    common.HexToAddress("0x806Bc0d7a47B890383a831634BCb92DD4030B092"),
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := PublicKeyToAddress(tt.args.publicKeyStr)
			if (err != nil) != tt.wantErr {
				t.Errorf("PublicKeyBytesToAddress() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PublicKeyBytesToAddress() = %v, want %v", got, tt.want)
			}
		})
	}
}
