package main

import (
	"bully/config"
	"bully/monitor"
	"bully/node"
	"bully/send"
	"bully/tasks"

	"fmt"
	"math"
	"os"
	"time"

	"github.com/spf13/viper"
)

func main() {

	// config
	config.InitConfig()

	// nodes
	numNodes := viper.GetInt("nodes.num")
	channel_size := viper.GetInt("channel.size")

	// channel
	individual_ch_msg := make([]chan send.Msg, numNodes+1)
	ch_main_send := make(chan send.Msg, channel_size)

	// start node
	for i := 0; i <= numNodes; i++ {
		individual_ch_msg[i] = make(chan send.Msg, channel_size)
		if i == 0 {
			continue
		}
		go node.RunNode(numNodes, i, individual_ch_msg[i], ch_main_send)
	}

	// monitor channel
	all_channel := append([]chan send.Msg{ch_main_send}, individual_ch_msg[1:]...)
	interval := viper.GetInt("channel.monitor")
	go monitor.MonitorChannelLen(all_channel, interval)

	// assign task
	go tasks.AssignTask(numNodes, individual_ch_msg)

	// main node(channel)
	// forLoop:
	for job := range ch_main_send {
		fmt.Printf("Main: %s\nReceived data: %+v\n", job.MsgType, job)

		switch job.MsgType {
		// default:
		// 	break forLoop

		// Basic Task
		case send.Done_Task:
			fmt.Printf("Instruct node %d to report to the leader(%d)\n", job.To[0], node.GetLeader())
			individual_ch_msg[job.To[0]] <- send.Msg{From: job.From, TaskType: send.BASIC, MsgType: send.ReportToLeader, Msg: job.Msg}
		case send.Done_ReportToLeader:
			fmt.Printf("Instruct the leader(%d) to repond\n", job.To[0])
			individual_ch_msg[job.To[0]] <- send.Msg{From: job.From, TaskType: send.BASIC, MsgType: send.RespondToReport, Msg: job.Msg}
		case send.Done_RespondToReport:
			fmt.Printf("Notify node %d that the leader(%d) is functioning properly. task: %s\n", job.To[0], job.From, job.Msg)
			individual_ch_msg[job.To[0]] <- send.Msg{From: job.From, TaskType: send.BASIC, MsgType: send.CompleteRequestToLeader, Msg: job.Msg}

		// Branching Point
		case send.Done_CompleteRequestToLeader:
			switch job.Status {
			case send.SUCCESS:
				fmt.Printf("Wow! node %d's task %s is finished!\n", job.From, job.Msg)
			case send.FAIL:
				fmt.Printf(":( It seems that the leader(%d) is dead. Conduct the election process\n", node.GetLeader())
				individual_ch_msg[job.To[0]] <- send.Msg{From: job.From, TaskType: send.ELECTION, MsgType: send.HoldElection, Msg: job.Msg}
			}

		// Election Task
		case send.Done_HoldElection:
			for _, to := range job.To {
				fmt.Printf("Instruct node %d to respond to the election process", to)
				individual_ch_msg[to] <- send.Msg{From: job.From, TaskType: send.ELECTION, MsgType: send.RespondToElection, Msg: job.Msg}
			}
		case send.Done_RespondToElection:
			fmt.Printf("Instruct node %d to end the election process as a higher-priority node %d has appeared.", job.To[0], job.From)
			individual_ch_msg[job.To[0]] <- send.Msg{From: job.From, TaskType: send.ELECTION, MsgType: send.TransferInitiative, Msg: job.Msg}
		case send.Done_TransferInitiative:
			fmt.Printf("Node %d's election process is ended\n", job.From)
			os.Exit(0)
		}
	}

	time.Sleep(time.Duration(math.MaxInt64))
}
