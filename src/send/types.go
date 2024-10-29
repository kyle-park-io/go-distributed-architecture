package send

type Msg struct {
	MsgType int
	Msg     int
}

type Tx struct {
	From int
	To   int
	Msg  Msg
}
