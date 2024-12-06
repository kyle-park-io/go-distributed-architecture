package node

import "fmt"

type Node struct {
	ID           int
	Channel_Node chan int
	Channel_Main chan int
}

func (n *Node) RunNode() {
	fmt.Printf("Run node! id: %d\n", n.ID)

	for job := range n.Channel_Node {
		fmt.Printf("Node%d: get msg: %d\n", n.ID, job)

		n.Channel_Main <- job
	}

	fmt.Printf("End node! id: %d\n", n.ID)
}
