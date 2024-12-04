package utils

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func ExtractEthereumAddress(pubKey *ecdsa.PublicKey) common.Address {
	// Serialize the public key and compute the Keccak-256 hash
	pubKeyBytes := crypto.FromECDSAPub(pubKey)
	hash := crypto.Keccak256(pubKeyBytes[1:]) // Remove the prefix byte (0x04)

	// Return the last 20 bytes as the Ethereum address
	return common.BytesToAddress(hash[12:])
}
