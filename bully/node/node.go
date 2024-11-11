package node

import (
	"bully/scenario"
	"bully/send"
	"bully/tasks"
	"bully/utils"

	"fmt"
	"strings"

	"github.com/spf13/viper"
)

var leader int
var nodes []int

func GetLeader() int {
	return leader
}

func GetNodes() []int {
	return nodes
}

func RunNode(numNodes int, id int,
	ch_msg <-chan send.Msg,
	ch_main_send chan<- send.Msg) {

	fmt.Printf("Run node! id: %d\n", id)

	// leader
	if numNodes == id {
		nodes = make([]int, numNodes+1)
		for i := 0; i <= numNodes; i++ {
			nodes[i] = i
		}
		leader = numNodes
	}

	check_process := make(map[string]bool)
	ch_timeout := make(chan string)

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
				tasks.DoTask(r)
				fmt.Printf("Node %d: End task\n", id)
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.BASIC, MsgType: send.Done_Task, Msg: job.Msg}
			case send.ReportToLeader: // Report to the leader
				t := viper.GetInt("nodes.leader_timeout")
				fmt.Printf("Node %d: Let's check leader response, timeout: %ds\n", id, t)
				go utils.SetTimeout(t, ch_timeout, job.Msg)
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
			case send.HoldElection: // Elect a new leader (Reach out to the number above mine)
				if id+1 == len(GetNodes()) {
					fmt.Printf("Warning: election was conducted even though the leader(%d) was in a normal state\n", GetLeader())
					ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_TransferInitiative, Msg: job.Msg}
					continue
				}
				if id+2 == len(GetNodes()) {
					fmt.Printf("Congratulations! node %d has become the new leader", id)
					leader = id
					ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_TransferInitiative, Msg: job.Msg}
					continue
				}
				candidates := nodes[id+1 : numNodes]
				fmt.Printf("Candidates: %v\n", candidates)
				ch_main_send <- send.Msg{From: id, To: candidates, TaskType: send.ELECTION, MsgType: send.Done_HoldElection, Msg: job.Msg}
				t := viper.GetInt("nodes.election_timeout")
				go utils.SetTimeout(t, ch_timeout, "Election"+job.Msg)
			case send.RespondToElection: // Respond to the election request
				ch_main_send <- send.Msg{From: id, To: []int{job.From}, TaskType: send.ELECTION, MsgType: send.Done_RespondToElection, Msg: job.Msg}
			case send.TransferInitiative: // End the election process (Take over the initiative to higher-node)
				if check_process["Election"+job.Msg] {
					fmt.Printf("Node %d: this election process(%s) has already been terminated because another node has responded\n", id, job.Msg)
					continue
				}
				fmt.Printf("Node %d: end the election process! higher-node: %d\n", id, job.From)
				check_process["Election"+job.Msg] = true
				ch_main_send <- send.Msg{From: job.From, To: []int{job.From}, TaskType: send.ELECTION, MsgType: send.Done_CompleteRequestToLeader, Msg: job.Msg, Status: send.FAIL}
			}

		case msg := <-ch_timeout:
			if strings.HasPrefix(msg, "Election") {
				if check_process[msg] {
					fmt.Printf("Node %d: the voting process has already been passed to another higher-priority node\n", id)
					continue
				}
				fmt.Printf("Congratulations! node %d has become the new leader", id)
				leader = id
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.ELECTION, MsgType: send.Done_TransferInitiative, Msg: msg}
			} else {
				// Fail route
				if check_process[msg] {
					fmt.Printf("Node %d: It has already been determined that the leader(%d) is in a normal state\n", id, GetLeader())
					continue
				}
				fmt.Printf("Node %d: no response has been received from the leader(%d) after some time. Will proceed with the election\n", id, GetLeader())
				ch_main_send <- send.Msg{From: id, To: []int{id}, TaskType: send.BASIC, MsgType: send.Done_CompleteRequestToLeader, Msg: msg, Status: send.FAIL}
			}
		}
	}

	fmt.Printf("End node! id: %d\n", id)
}
