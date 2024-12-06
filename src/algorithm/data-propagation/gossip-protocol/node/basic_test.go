package node

import (
	"gossip-protocol/engine"
	"testing"
	"time"
)

func TestRunNode(t *testing.T) {
	numNodes := 3

	ch_slice := make([]chan int, numNodes+1)
	// run node
	for i := 0; i <= numNodes; i++ {
		ch_slice[i] = make(chan int, 10)
		if i == 0 {
			continue
		}

		node := Node{ID: i, Channel_Node: ch_slice[i],
			Channel_Main: ch_slice[0]}
		go node.RunNode()
	}
	time.Sleep(3 * time.Second)

	// run task
	go func() {
		for i := 0; i < 5; i++ {
			r := engine.SendMsgToRandomNode(numNodes)
			ch_slice[r] <- r
		}
	}()

	timeout := 10 * time.Second
	for {
		select {
		case job := <-ch_slice[0]:
			t.Logf("Main: receive msg to node%d\n", job)

		case <-time.After(timeout):
			t.Logf("All nodes ran successfully for %v seconds", timeout.Seconds())
			return
		}
	}
}
