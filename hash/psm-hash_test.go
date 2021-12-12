package hash

import (
	"testing"

	"github.com/onmax/go-alastria/did"
	alaTypes "github.com/onmax/go-alastria/types"
)

func TestPsmHash(t *testing.T) {
	d, _ := did.NewDidFromString("did:ala:quor:redT:80d0ff71e3f212cf383e57446df878e348f825db")
	type args struct {
		signedJwt string
		did       *alaTypes.Did
	}
	tests := []struct {
		name  string
		args  args
		want  string
	}{
		{
			name: "test ok",
			args: args{
				signedJwt: "eyJ0eXAiOiJKV1QiLCJhbGciOiJFUzI1NksiLCJraWQiOiJkaWQ6YWxhOnF1b3I6cmVkdDoxMmVlYUNDQTllRWJCNzhlQjk3ZDdjYWM2YiNrZXlzLTEifQ.eyJqdGkiOiJodHRwczovL3d3dy5lbXByZXNhLmNvbS9hbGFzdHJpYS9jcmVkZW50aWFscy8zNzM0IiwiaXNzIjoiZGlkOmFsYTpxdW9yOnJlZFQ6NGI4MzQzMjAzNWQxNDc3ZDNhOTliMjNiYWE4MmJkOWRiNjFhMmQ2OCIsInN1YiI6ImRpZDphbGE6cXVvcjpyZWR0OjB4MTJlZWFDQ0E5ZUViQjc4ZUI5N2Q3Y2FjNmIiLCJpYXQiOjE2MzkwOTAyMzUsImV4cCI6MTU2Mzc4MzM5MiwibmJmIjoxNTYzNzgyNzkyLCJ2YyI6eyJAY29udGV4dCI6WyJodHRwczovL3d3dy53My5vcmcvMjAxOC9jcmVkZW50aWFscy92MSIsImh0dHBzOi8vYWxhc3RyaWEuZ2l0aHViLmlvL2lkZW50aXR5L2NyZWRlbnRpYWxzL3YxIiwiaHR0cHM6Ly93M2lkLm9yZy9kaWQvdjEiLCJKV1QiXSwidHlwZSI6WyJWZXJpZmlhYmxlQ3JlZGVudGlhbCIsIkFsYXN0cmlhVmVyaWZpYWJsZUNyZWRlbnRpYWwiXSwiY3JlZGVudGlhbFN1YmplY3QiOnsiU3R1ZGVudElEIjoiMTEyMzU4MTMiLCJsZXZlbE9mQXNzdXJhbmNlIjoxfX19.j4BhQTs1fi1DeFLjI9JWCQIo3Qkk372cq_bZHCo5fTIafqGXcHETnJG6xeZ3F35CeO4sZkSFlpaccTKfdH6kew",
				did:       d,
			},
			want:  "0xbc8f70b6722d1bff3129f17c9b72b27a136475f5bae9ac47b1d36301887b5dfb",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := PsmHash(tt.args.signedJwt, tt.args.did)
			if got != tt.want {
				t.Errorf("PsmHash() got = %v, want %v", got, tt.want)
			}
		})
	}
}
