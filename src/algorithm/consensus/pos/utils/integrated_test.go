package utils

import "testing"

func TestCombineRandaoAndVRF(t *testing.T) {
	// randao
	randao := &RANDAO{}

	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))

	randao.Reveal(1, []byte("secret1"))
	randao.Reveal(2, []byte("secret2"))
	randao.Reveal(3, []byte("secret3"))

	randao.CalculateFinalValue()

	// vrf
	privKey, _ := GenerateKeyPair()
	input := []byte("test_vrf_randomness")

	vrf := RunVRF(privKey, input)

	c := CombineRandaoAndVRF(randao.FinalValue, vrf.Output)
	t.Logf("Combined Randao & VRF: %x\n", c)
}
