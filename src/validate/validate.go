package validate

import (
	"go-distributed-architecture/send"
	"go-distributed-architecture/utils"
)

func VerifyTx(tx send.Tx) bool {
	r := utils.GetRandom(2)
	switch r {
	case 0:
		return true
	case 1:
		return false
	default:
		return false
	}
}
