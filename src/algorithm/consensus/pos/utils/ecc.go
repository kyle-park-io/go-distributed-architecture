package utils

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
)

// secp256r1
func GenerateKeyPair() (*ecdsa.PrivateKey, ecdsa.PublicKey) {
	privKey, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}
	return privKey, privKey.PublicKey
}
