package utils

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type RANDAO struct {
	Participants []Participant
	FinalValue   []byte
}

type Participant struct {
	ID     int
	Secret []byte
	Commit string
	Reveal []byte
}

// Commit phase: Participants submit the hash of their secret value
func (r *RANDAO) Commit(id int, secret []byte) {
	// make hash
	hash := sha256.Sum256(secret)

	r.Participants = append(r.Participants, Participant{
		ID:     id,
		Secret: secret,
		Commit: hex.EncodeToString(hash[:]),
	})

	fmt.Printf("Participant %d committed: %s\n", id, hex.EncodeToString(hash[:]))
}

// Reveal phase: Participants reveal their original secret
func (r *RANDAO) Reveal(id int, secret []byte) {
	for i, p := range r.Participants {
		if p.ID == id {
			// Verify the secret matches the commit hash
			hash := sha256.Sum256(secret)
			if p.Commit == hex.EncodeToString(hash[:]) {
				r.Participants[i].Reveal = secret
				fmt.Printf("Participant %d revealed: %s\n", id, hex.EncodeToString(secret))
				return
			} else {
				fmt.Printf("Participant %d failed to reveal correctly.\n", id)
				return
			}
		}
	}
	fmt.Printf("Participant %d not found.\n", id)
}

// Calculate final random value by XORing all revealed values
func (r *RANDAO) CalculateFinalValue() {
	var finalValue []byte
	for _, p := range r.Participants {
		if p.Reveal != nil {
			if finalValue == nil {
				finalValue = make([]byte, len(p.Reveal))
			}
			for i := 0; i < len(p.Reveal); i++ {
				finalValue[i] ^= p.Reveal[i]
			}
		}
	}
	r.FinalValue = finalValue
	fmt.Printf("Final Random Value: %s\n", hex.EncodeToString(finalValue))
}
