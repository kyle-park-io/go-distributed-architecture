package goroutinetasks

import (
	"go-distributed-architecture/node"
	"go-distributed-architecture/send"
	"go-distributed-architecture/utils"

	"time"
)

func MsgTask(ch_msg chan send.Msg) {
	for i := 0; i < 100000; i++ {
		time.Sleep(3 * time.Second)

		randomMsgType := utils.GetRandom(2)
		randomMsg := utils.GetRandom(1000)
		switch randomMsgType {
		case node.SaveMsg:
			send.SendMsg(ch_msg, node.SaveMsg, randomMsg)
		case node.ReturnMsg:
			send.SendMsg(ch_msg, node.ReturnMsg, 0)
		}
	}
}

func TxTask(ch_tx chan send.Tx) {
	for i := 0; i < 100000; i++ {
		time.Sleep(3 * time.Second)

		randomMsgType := utils.GetRandom(3) + 2
		randomMsg := utils.GetRandom(1000)
		switch randomMsgType {
		case node.SaveTx:
			send.SendTxToRandomNode(ch_tx, node.ADMIN_NODE, node.SaveTx, randomMsg)
		case node.ReturnTx:
			send.SendTxToRandomNode(ch_tx, node.ADMIN_NODE, node.ReturnTx, 0)
		case node.VerifyTx:
			send.SendTxToRandomNode(ch_tx, node.ADMIN_NODE, node.VerifyTx, randomMsg)
		}
	}
}
