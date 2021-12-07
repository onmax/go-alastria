package tokens

import "testing"

type TestParser struct {
	Sub   string `json:"sub"`
	Name  string `json:"name"`
	Admin bool   `json:"admin"`
}

func TestB64ToInterface(t *testing.T) {
	type args struct {
		input  string
		output interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "TestB64ToInterface",
			args: args{
				input:  "eyJzdWIiOiIxMjM0NTY3ODkwIiwibmFtZSI6IkpvaG4gRG9lIiwiYWRtaW4iOnRydWV9",
				output: &TestParser{},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := B64ToInterface(tt.args.input, tt.args.output); (err != nil) != tt.wantErr {
				t.Errorf("B64ToInterface() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestInterfaceToB64(t *testing.T) {
	type args struct {
		artifact interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := InterfaceToB64(tt.args.artifact)
			if (err != nil) != tt.wantErr {
				t.Errorf("InterfaceToB64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("InterfaceToB64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestJwtToB64(t *testing.T) {
	type args struct {
		jwt interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := JwtToB64(tt.args.jwt)
			if (err != nil) != tt.wantErr {
				t.Errorf("JwtToB64() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("JwtToB64() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("JwtToB64() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestSplitJWT(t *testing.T) {
	type args struct {
		signedJWT string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		want2   string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, got2, err := SplitJWT(tt.args.signedJWT)
			if (err != nil) != tt.wantErr {
				t.Errorf("SplitJWT() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("SplitJWT() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("SplitJWT() got1 = %v, want %v", got1, tt.want1)
			}
			if got2 != tt.want2 {
				t.Errorf("SplitJWT() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
