package main

import (
	"go-distributed-architecture/node"
	"go-distributed-architecture/send"
	"go-distributed-architecture/utils"

	"fmt"
	"time"
)

func main() {

	ch_msg := make(chan send.Msg)
	ch_main_signal := make(chan int)
	ch_main := make(chan []send.Msg)

	for i := 0; i < 3; i++ {
		go node.RunNode(i, ch_msg, ch_main_signal, ch_main)
	}

	go func() {
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
