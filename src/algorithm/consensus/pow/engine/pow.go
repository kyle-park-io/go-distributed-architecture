package engine

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"strings"
)

type Block struct {
	PrevHash  string
	Timestamp int64
	Data      string // Merkle Root or String
	Nonce     int
	Hash      string
}

func (b *Block) ProofOfWork(difficulty int) {
	prefix := strings.Repeat("0", difficulty)
	for {
		b.Hash = b.calculateHash()
		if strings.HasPrefix(b.Hash, prefix) {
			break
		}
		b.Nonce++
	}
}

func (b *Block) calculateHash() string {
	data := b.PrevHash + fmt.Sprintf("%d", b.Timestamp) + b.Data + fmt.Sprintf("%d", b.Nonce)
	hash := sha256.Sum256([]byte(data))
	return hex.EncodeToString(hash[:])
}
