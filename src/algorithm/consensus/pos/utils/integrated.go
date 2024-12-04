package utils

import (
	"math/big"
)

func (r *RANDAO) SelectBlockProposer() Participant {
	var selected Participant

	maxValue := make([]byte, 32)
	for i := range maxValue {
		maxValue[i] = 0xFF
	}
	smallestHash := new(big.Int).SetBytes(maxValue) // Biggest init value

	randaoSeed := r.FinalValue
	for _, v := range r.Participants {
		vrf := RunVRF(v.PrivKey, randaoSeed)

		combinedHash := new(big.Int).SetBytes(vrf.Output)
		if combinedHash.Cmp(smallestHash) < 0 {
			smallestHash.Set(combinedHash)
			selected = v
		}
	}

	return selected
}
