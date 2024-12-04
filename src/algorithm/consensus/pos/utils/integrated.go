package utils

import "crypto/sha256"

func CombineRandaoAndVRF(randaoSeed []byte, vrfOutput []byte) []byte {
	combined := append(randaoSeed, vrfOutput...)
	finalHash := sha256.Sum256(combined)
	return finalHash[:]
}
