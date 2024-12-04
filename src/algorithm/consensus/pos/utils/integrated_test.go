package utils

import "testing"

func TestSelectBlockProposer(t *testing.T) {
	// randao
	randao := &RANDAO{}

	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))

	randao.Reveal(1, []byte("secret1"))
	randao.Reveal(2, []byte("secret2"))
	randao.Reveal(3, []byte("secret3"))

	randao.CalculateFinalValue()

	// select block proposer
	selected := randao.SelectBlockProposer()
	t.Logf("Selected Validator: %+v\n", selected)
}
