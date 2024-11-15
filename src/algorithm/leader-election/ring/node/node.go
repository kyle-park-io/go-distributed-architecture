package node

import (
	"ring/scenario"
	"ring/send"
	"ring/utils"

	"fmt"
	"sync/atomic"

	"github.com/spf13/viper"
)

var allNodesInitialized int64
var leader int
var nodes []int

func GetAllNodesInitialized() bool {
	return allNodesInitialized == 1
}

func GetLeader() int {
	return leader
}

func GetNodes() []int {
	return nodes
}

func RunNode(numNodes int, id int, nextNode int, ns []int,
	ch_msg <-chan send.Msg,
	ch_main_send chan<- send.Msg) {

	fmt.Printf("Run node! id: %d\n", id)

	// leader
	if numNodes == id {
		leader = numNodes
		nodes = ns
		atomic.StoreInt64(&allNodesInitialized, 1)
		fmt.Println(nodes)
	}

	check_process := make(map[string]bool)
	ch_timeout := make(chan send.Msg)

	// forLoop:
	for {
		select {
		// default:
		// 	break forLoop
		case job := <-ch_msg:
			fmt.Printf("Node %d: %s\nReceived data: %+v\n", id, job.MsgType, job)

			switch job.MsgType {

			// Failure Task
			case send.FAILURE:
				scenario.NoResponse()

			// Basic Task
			case send.DoTask: // The nodes work
				r := utils.GetRandom(viper.GetInt("tasks.duration"))
				fmt.Printf("Node %d: Start task, runtime: %ds\n", id, r)
				DoTask(r)
				fmt.Printf("Node %d: End task\n", id)
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.BASIC, MsgType: send.Done_Task, Msg: job.Msg}
			case send.ReportToLeader: // Report to the leader
				t := viper.GetInt("nodes.leader_timeout")
				fmt.Printf("Node %d: Let's check leader response, timeout: %ds\n", id, t)
				go utils.SetTimeout(t, ch_timeout, job)
				ch_main_send <- send.Msg{From: id, To: []int{GetLeader()}, TaskType: send.BASIC, MsgType: send.Done_ReportToLeader, Msg: job.Msg}
			case send.RespondToReport: // Respond to report
				if id != leader {
					panic(fmt.Sprintf("Wrong access: node %d is not the leader(%d)\n", id, GetLeader()))
				}
				ch_main_send <- send.Msg{From: id, To: []int{job.From}, TaskType: send.BASIC, MsgType: send.Done_RespondToReport, Msg: job.Msg}
			case send.CompleteRequestToLeader: // Complete the leader status verification process
				// Success route
				check_process[job.Msg] = true
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.BASIC, MsgType: send.Done_CompleteRequestToLeader, Msg: job.Msg, Status: send.SUCCESS}

			// Election Task
			case send.HoldElection: // Announce the start of the election
				fmt.Printf("Node %d: Hold the election!\n", id)
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_HoldElection, Msg: job.Msg}
			case send.RequestNextNodeAction: // Elect a new leader (Traverse around the ring node once)
				// end of election
				if job.SequentialCount != 0 && job.SequentialCount == numNodes-1 {
					fmt.Printf("Node %d: Complete the election cycle\n", id)
					ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_Election, Msg: job.Msg,
						CurrentHighestNode: job.CurrentHighestNode, CandidateNodes: job.CandidateNodes, SequentialCount: job.SequentialCount, TimeoutCount: job.TimeoutCount}
					continue
				}

				next := 0
				if job.TimeoutCount == 0 {
					next = nextNode
					job.CandidateNodes = append(job.CandidateNodes, id)
					if id > job.CurrentHighestNode {
						job.CurrentHighestNode = id
					}
				} else {
					for i := 0; i < len(nodes); i++ {
						if nodes[i] == id {
							next = nodes[(i+1)%numNodes]
						}
					}
				}
				job.To = []int{next}
				fmt.Printf("Node %d: send a request to the next node %d\n", id, next)
				t := viper.GetInt("nodes.election_timeout")
				go utils.SetTimeout(t, ch_timeout, job)
				ch_main_send <- send.Msg{From: id, To: []int{next}, TaskType: send.ELECTION, MsgType: send.Done_RequestNextNodeAction, Msg: job.Msg,
					CurrentHighestNode: job.CurrentHighestNode, CandidateNodes: job.CandidateNodes, SequentialCount: job.SequentialCount + 1, TimeoutCount: job.TimeoutCount}
			case send.RespondToElection: // Respond to the election request
				fmt.Printf("Node %d: respond to the election\n", id)
				ch_main_send <- send.Msg{From: id, To: []int{job.From}, TaskType: send.ELECTION, MsgType: send.Done_RespondToElection, Msg: job.Msg,
					CurrentHighestNode: job.CurrentHighestNode, CandidateNodes: job.CandidateNodes, SequentialCount: job.SequentialCount, TimeoutCount: job.TimeoutCount}
			case send.TransferInitiative: // Take over the initiative to next node
				fmt.Printf("Node %d: It has been successfully forwarded to the next node %d\n", id, job.From)
				check_process["Election"+job.Msg] = true
				ch_main_send <- send.Msg{From: id, To: []int{job.From}, TaskType: send.ELECTION, MsgType: send.Done_TransferInitiative, Msg: job.Msg, Status: send.SUCCESS,
					CurrentHighestNode: job.CurrentHighestNode, CandidateNodes: job.CandidateNodes, SequentialCount: job.SequentialCount, TimeoutCount: 0}
			case send.AssignLeader: // Assign a new leader
				leader = job.CurrentHighestNode
				fmt.Printf("Node %d: A new leader is %d\n", id, leader)
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_AssignLeader, Msg: job.Msg,
					CurrentHighestNode: job.CurrentHighestNode}
			}

		case msg := <-ch_timeout:
			if msg.TaskType == send.ELECTION {
				if check_process["Election"+msg.Msg] {
					fmt.Printf("Node %d: the election process has already been passed to next node\n", id)
					continue
				}

				fmt.Printf("Node %d: node %d is unresponsive. Sending the message to the next node\n", id, msg.To[0])
				check_process["Election"+msg.Msg] = true
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_TransferInitiative, Msg: msg.Msg, Status: send.FAIL,
					CurrentHighestNode: msg.CurrentHighestNode, CandidateNodes: msg.CandidateNodes, SequentialCount: msg.SequentialCount + 1, TimeoutCount: msg.TimeoutCount + 1}
			} else {
				// Fail route
				if check_process[msg.Msg] {
					fmt.Printf("Node %d: It has already been determined that the leader(%d) is in a normal state\n", id, GetLeader())
					continue
				}
				fmt.Printf("Node %d: no response has been received from the leader(%d) after some time. Will proceed with the election\n", id, GetLeader())
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.BASIC, MsgType: send.Done_CompleteRequestToLeader, Msg: msg.Msg, Status: send.FAIL}
			}
		}
	}

	fmt.Printf("End node! id: %d\n", id)
}
