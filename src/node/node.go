package node

import (
	"go-distributed-architecture/send"

	"fmt"
)

func RunNode(id int, ch_msg chan send.Msg, ch_main_signal chan int, ch_main chan []send.Msg) {
	fmt.Printf("Run node! id: %d\n", id)

	m := make([]send.Msg, 0)
	for job := range ch_msg {
		fmt.Printf("Received data: %+v\n", job)

		switch job.MsgType {
		case SaveMsg:
			m = append(m, job)
			ch_main_signal <- Success
		case ReturnMsg:
			ch_main <- m
			ch_main_signal <- Success
		}
	}

	fmt.Printf("End node! id: %d\n", id)
}
