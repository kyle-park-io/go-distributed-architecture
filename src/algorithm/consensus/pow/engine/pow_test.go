package engine

import (
	"strconv"
	"strings"
	"testing"
	"time"
)

func TestProofOfWork(t *testing.T) {
	difficulty := 4

	block := Block{
		PrevHash:  "0000000000000000",
		Timestamp: time.Now().Unix(),
		Data:      "Genesis Block",
		Nonce:     0,
		Hash:      "",
	}
	block.ProofOfWork(difficulty)

	expectedPrefix := strings.Repeat("0", difficulty)
	if !strings.HasPrefix(block.Hash, expectedPrefix) {
		t.Errorf("Proof of Work failed: Hash does not satisfy difficulty.\nGot: %s\nExpected Prefix: %s", block.Hash, expectedPrefix)
	}

	if block.Nonce <= 0 {
		t.Errorf("Nonce should be greater than 0. Got: %d", block.Nonce)
	}

	calculatedHash := block.calculateHash()
	if block.Hash != calculatedHash {
		t.Errorf("Hash mismatch: Expected %s, Got: %s", calculatedHash, block.Hash)
	}

	if block.Timestamp > time.Now().Unix() {
		t.Errorf("Timestamp is invalid. Got: %d", block.Timestamp)
	}
}

func TestMultipleProofOfWork(t *testing.T) {
	difficulty := 4
	numBlocks := 5

	// Genesis Block
	prevHash := "0000000000000000"

	for i := 1; i <= numBlocks; i++ {
		data := "Block #" + strconv.Itoa(i)

		block := Block{
			PrevHash:  prevHash,
			Timestamp: time.Now().Unix(),
			Data:      data,
			Nonce:     0,
		}
		block.ProofOfWork(difficulty)

		if block.PrevHash != prevHash {
			t.Errorf("Block %d PrevHash mismatch: Expected %s, Got %s", i, prevHash, block.PrevHash)
		}

		expectedPrefix := strings.Repeat("0", difficulty)
		if !strings.HasPrefix(block.Hash, expectedPrefix) {
			t.Errorf("Block %d Proof of Work failed: Hash does not satisfy difficulty.\nGot: %s\nExpected Prefix: %s", i, block.Hash, expectedPrefix)
		}

		if block.Nonce <= 0 {
			t.Errorf("Nonce should be greater than 0. Got: %d", block.Nonce)
		}

		calculatedHash := block.calculateHash()
		if block.Hash != calculatedHash {
			t.Errorf("Hash mismatch: Expected %s, Got: %s", calculatedHash, block.Hash)
		}

		if block.Timestamp > time.Now().Unix() {
			t.Errorf("Timestamp is invalid. Got: %d", block.Timestamp)
		}

		t.Logf("Block %d mined: Nonce=%d, Hash=%s", i, block.Nonce, block.Hash)

		prevHash = block.Hash
	}
}
