package send

type Msg struct {
	From     int
	To       []int
	TaskType Enum
	MsgType  Enum
	Msg      string
	Status   Enum
}
