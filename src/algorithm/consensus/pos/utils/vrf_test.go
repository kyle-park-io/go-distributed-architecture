package utils

import (
	"testing"
)

func TestRunVRF(t *testing.T) {
	// generate key
	privKey, _ := GenerateKeyPair()

	// input data
	input := []byte("test_vrf_randomness")

	// run
	vrf := RunVRF(privKey, input)
	t.Logf("VRF Output: %x\n", vrf.Output)
	t.Logf("Proof (r): %s\n", vrf.ProofR.String())
	t.Logf("Proof (s): %s\n", vrf.ProofS.String())
}

func TestVerifyVRF(t *testing.T) {
	// generate key
	privKey, pubKey := GenerateKeyPair()

	// input data
	input := []byte("test_vrf_randomness")

	// run
	vrf := RunVRF(privKey, input)

	// verify
	isValid := VerifyVRF(pubKey, input, vrf)
	t.Logf("VRF Proof Valid: %t\n", isValid)
}
