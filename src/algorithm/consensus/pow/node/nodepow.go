package node

import (
	"pow/engine"

	"fmt"
)

type config struct {
	numBlocks  int
	difficulty int
}
type Status struct {
	Block           engine.Block
	currentBlockNum int
	// numFinished     bool
}
type Node2 struct {
	ID           int
	Channel_Node chan Status
	Channel_Main chan Status
	config       config
}

func (n *Node2) RunNode2() {
	fmt.Printf("Run node! id: %d\n", n.ID)

	// findNonce := false
	for {
		select {
		case job := <-n.Channel_Node:
			// if n.ID == 1 {
			// 	time.Sleep(100 * time.Second)
			// }

			job.Block.ProofOfWork(n.config.difficulty)
			fmt.Printf("Node%d: find nonce! nonce: %d block: %d\n", n.ID, job.Block.Nonce, job.currentBlockNum)

			n.Channel_Main <- job
		}
	}

	fmt.Printf("End node! id: %d\n", n.ID)
}
