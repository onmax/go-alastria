package hash

import (
	"github.com/onmax/go-alastria/hex"
	alaTypes "github.com/onmax/go-alastria/types"
)

func PsmHash(signedJwt string, did *alaTypes.Did) (string, [32]byte) {
	psmHash := Sha3(signedJwt + did.String())
	byteArr := hex.StringTo32ByteArray(psmHash)
	return psmHash, byteArr
}
