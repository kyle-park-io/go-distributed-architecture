package engine

import (
	"pos/utils"

	"crypto/ecdsa"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

type Block struct {
	PrevHash    string
	Slot        int64
	Data        string // Merkle Root or String
	Hash        string
	Validator   string // Block creator
	Signature   string
	StakeAmount int64
}

func (b *Block) ProofOfStake(randao *utils.RANDAO) {
	randao.Commit(1, []byte("secret1"))
	randao.Commit(2, []byte("secret2"))
	randao.Commit(3, []byte("secret3"))

	randao.Reveal(1, []byte("secret1"))
	randao.Reveal(2, []byte("secret2"))
	randao.Reveal(3, []byte("secret3"))

	randao.CalculateFinalValue()

	// creator
	selected := randao.SelectBlockProposer()

	// data(txs)
	data := b.PrevHash + b.Data + fmt.Sprintf("%d", b.Slot)
	hash := sha256.Sum256([]byte(data))

	// sign
	r, s, err := ecdsa.Sign(rand.Reader, selected.PrivKey, hash[:])
	if err != nil {
		panic(err)
	}
	signature := fmt.Sprintf("%s:%s", r.String(), s.String())

	b.Hash = hex.EncodeToString(hash[:])
	b.Validator = utils.ExtractEthereumAddress(&selected.PublicKey).Hex()
	b.Signature = signature
}

// merkle root
func calculateMerkleRoot(transactions []string) string {
	if len(transactions) == 0 {
		return ""
	}
	hash := sha256.New()
	for _, tx := range transactions {
		hash.Write([]byte(tx))
	}
	return hex.EncodeToString(hash.Sum(nil))
}
