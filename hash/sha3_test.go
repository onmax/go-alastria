package hash

import "testing"

func TestSha3(t *testing.T) {
	type args struct {
		input string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "test ok",
			args: args{
				input: "hello world",
			},
			want: "0x47173285a8d7341e5e972fc677286384f802f8ef42a5ec5f03bbfa254cb01fad",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sha3(tt.args.input); got != tt.want {
				t.Errorf("Sha3() = %v, want %v", got, tt.want)
			}
		})
	}
}
