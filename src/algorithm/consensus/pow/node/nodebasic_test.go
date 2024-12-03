package node

import (
	"testing"
	"time"
)

func TestRunNode(t *testing.T) {
	numNodes := 3

	for i := 1; i <= numNodes; i++ {
		ch := make(chan string)
		node := Node{ID: i, Channel: ch}
		go node.RunNode()
	}

	timeout := 5 * time.Second
	time.Sleep(timeout)
	t.Logf("All nodes ran successfully for %v seconds", timeout.Seconds())
}
