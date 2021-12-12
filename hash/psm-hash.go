package hash

import (
	alaTypes "github.com/onmax/go-alastria/types"
)

func PsmHash(signedJwt string, did *alaTypes.Did) string {
	return Sha3(signedJwt + did.String())
}
