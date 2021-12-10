package did

import (
	"reflect"
	"testing"

	alaTypes "github.com/onmax/go-alastria/types"
)

func TestNewDid(t *testing.T) {
	type args struct {
		network      string
		networkId    string
		proxyAddress string
	}
	tests := []struct {
		name string
		args args
		want *alaTypes.Did
	}{
		{
			name: "did ok",
			args: args{
				network:      "test",
				networkId:    "test",
				proxyAddress: "test",
			},
			want: &alaTypes.Did{
				Protocol:     "ala",
				Network:      "test",
				NetworkId:    "test",
				ProxyAddress: "test",
			},
		},
		{
			name: "did not ok",
			args: args{
				network:      "",
				networkId:    "",
				proxyAddress: "",
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewDid(tt.args.network, tt.args.networkId, tt.args.proxyAddress); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewDid() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDid_String(t *testing.T) {
	type fields struct {
		Protocol     string
		Network      string
		NetworkId    string
		ProxyAddress string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "did ok",
			fields: fields{
				Protocol:     "ala",
				Network:      "net",
				NetworkId:    "netId",
				ProxyAddress: "0123456789abcdef0123456789abcdef01234567",
			},
			want: "did:ala:net:netId:0123456789abcdef0123456789abcdef01234567",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := &alaTypes.Did{
				Protocol:     tt.fields.Protocol,
				Network:      tt.fields.Network,
				NetworkId:    tt.fields.NetworkId,
				ProxyAddress: tt.fields.ProxyAddress,
			}
			if got := d.String(); got != tt.want {
				t.Errorf("Did.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
