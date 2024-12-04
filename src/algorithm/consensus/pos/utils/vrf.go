package utils

import (
	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"math/big"
)

type VRFOutput struct {
	Output []byte   // output
	ProofR *big.Int // r
	ProofS *big.Int // s
}

func RunVRF(privKey *ecdsa.PrivateKey, input []byte) VRFOutput {
	// make hash
	hashInput := append(input, privKey.PublicKey.X.Bytes()...)
	hash := sha256.Sum256(hashInput)

	// sign data
	r, s, err := ecdsa.Sign(rand.Reader, privKey, hash[:])
	if err != nil {
		panic(err)
	}

	return VRFOutput{
		Output: hash[:],
		ProofR: r,
		ProofS: s,
	}
}

func VerifyVRF(pubKey ecdsa.PublicKey, input []byte, vrf VRFOutput) bool {
	// make hash
	hashInput := append(input, pubKey.X.Bytes()...)
	hash := sha256.Sum256(hashInput)

	isValid := ecdsa.Verify(&pubKey, hash[:], vrf.ProofR, vrf.ProofS)
	return isValid
}
