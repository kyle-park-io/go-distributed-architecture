package send

import "go-distributed-architecture/utils"

func SendData(ch chan int, data int) {
	ch <- data
}

func SendMsg(ch chan Msg, msgType int, msg int) {
	data := Msg{MsgType: msgType, Msg: msg}
	ch <- data
}

func SendTxToRandomNode(ch chan Tx, from int, msgType int, msg int) {
	for {
		random := utils.GetRandom(3)
		if random == from {
			continue
		}
		tx := Tx{
			From: from,
			To:   random,
			Msg:  Msg{MsgType: msgType, Msg: msg}}
		ch <- tx
		return
	}
}

func SendTxToOtherNode(ch chan Tx, from int, to int, msgType int, msg int) {
	if from == to {
		return
	}
	tx := Tx{
		From: from,
		To:   to,
		Msg:  Msg{MsgType: msgType, Msg: msg}}
	ch <- tx
}
