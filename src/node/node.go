package node

import (
	"go-distributed-architecture/send"

	"fmt"
)

// msg
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

// tx
func RunNode2(id int, ch_tx chan send.Tx, ch_main_signal chan int, ch_main chan []send.Tx) {
	fmt.Printf("Run node! id: %d\n", id)

	txs := make([]send.Tx, 0)
	for job := range ch_tx {
		fmt.Printf("Received data: %+v\n", job)

		switch job.Msg.MsgType {
		case SaveTx:
			txs = append(txs, job)
			ch_main_signal <- Success
		case ReturnTx:
			ch_main <- txs
			ch_main_signal <- Success
		}
	}

	fmt.Printf("End node! id: %d\n", id)
}
