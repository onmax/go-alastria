package hash

import (
	alaTypes "github.com/onmax/go-alastria/types"
)  

func PsmHash(signedJwt string, did *alaTypes.Did) (string, [32]byte) {
	psmHash := Sha3(signedJwt + did.String())

	var output [32]byte
	copy(output[:], psmHash)

	return psmHash, output
}
