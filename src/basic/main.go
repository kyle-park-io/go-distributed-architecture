package main

import (
	"go-distributed-architecture/goroutinetasks"
	"go-distributed-architecture/node"
	"go-distributed-architecture/send"

	"fmt"
)

func main() {

	// ch_msg := make(chan send.Msg)
	// ch_main_signal := make(chan int)
	// ch_main := make(chan []send.Msg)

	ch_tx := make(chan send.Tx)
	ch_main_signal := make(chan int)
	ch_main := make(chan []send.Tx)

	for i := 0; i < 3; i++ {
		// go node.RunNode(i, ch_msg, ch_main_signal, ch_main)
		go node.RunNode2(i, ch_tx, ch_main_signal, ch_main)
	}

	go func() {
		// go goroutinetasks.MsgTask(ch_msg)
		go goroutinetasks.TxTask(ch_tx)
	}()

	for {
		select {
		case job := <-ch_main_signal:
			fmt.Printf("Signal reaches the main: %d\n", job)
		case job := <-ch_main:
			fmt.Printf("Data reaches the main: %+v\n", job)
		}
	}

}
