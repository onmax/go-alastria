package hash

func PsmHash(signedJwt, did string) (string, [32]byte) {
	psmHash := Sha3(signedJwt + did)

	var output [32]byte
	copy(output[:], psmHash)

	return psmHash, output
}
