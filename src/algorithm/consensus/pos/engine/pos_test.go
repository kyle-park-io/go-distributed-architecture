package engine

import (
	"pos/utils"

	"encoding/json"
	"testing"
)

func TestProofOfStake(t *testing.T) {
	block := Block{
		PrevHash:    "0000000000000000",
		Slot:        0,
		Data:        "Genesis Block",
		Hash:        "",
		Validator:   "",
		Signature:   "",
		StakeAmount: 0,
	}

	randao := &utils.RANDAO{}
	block.ProofOfStake(randao)

	prettyJSON, err := json.MarshalIndent(block, "", "  ")
	if err != nil {
		t.Log("Error marshalling to JSON: ", err)
		return
	}

	t.Logf("Block:\n%v\n", string(prettyJSON))
}
