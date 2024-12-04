package utils

import "testing"

// If you want an accurate implementation, you need to achieve complete randomness using goroutines and mutex.
func TestCommit(t *testing.T) {
	randao := &RANDAO{}

	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))
}

func TestReveal(t *testing.T) {
	randao := &RANDAO{}

	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))

	randao.Reveal(1, []byte("secret1"))
	randao.Reveal(2, []byte("secret2"))
	randao.Reveal(3, []byte("secret3"))
}

func TestCalculateFinalValue(t *testing.T) {
	randao := &RANDAO{}

	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))

	randao.Reveal(1, []byte("secret1"))
	randao.Reveal(2, []byte("secret2"))
	randao.Reveal(3, []byte("secret3"))

	randao.CalculateFinalValue()
}
