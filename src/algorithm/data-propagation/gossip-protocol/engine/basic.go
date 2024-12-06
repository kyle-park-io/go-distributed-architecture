package engine

import (
	"gossip-protocol/utils"

	"fmt"
)

func SendMsgToRandomNode(n int) int {
	r := utils.GetRandom(n) + 1
	fmt.Printf("Send Msg to Node%d\n", r)

	return r
}
