package main

func node(id int, ch chan int) {}

func tasks() {}

func rule() {}

func main() {

	numNodes := 5
	ch := make(chan int)

	// start node
	for i := 1; i <= numNodes; i++ {
		go node(i, ch)
	}

	// set rules(tasks)
	go tasks()
	go rule()

	// main node(channel)
	for {
		select {
		case <-ch:
		}
	}
}
