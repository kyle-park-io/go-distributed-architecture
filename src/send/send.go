package send

func SendData(ch chan int, data int) {
	ch <- data
}

func SendMsg(ch chan Msg, msgType int, msg int) {
	data := Msg{MsgType: msgType, Msg: msg}
	ch <- data
}
