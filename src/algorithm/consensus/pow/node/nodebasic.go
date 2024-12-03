package node

import "fmt"

type Node struct {
	ID      int
	Channel chan string
}

func (n *Node) RunNode() {
	fmt.Printf("Run node! id: %d\n", n.ID)

	for job := range n.Channel {
		fmt.Printf("Get Job: %s\n", job)
	}

	fmt.Printf("End node! id: %d\n", n.ID)
}
