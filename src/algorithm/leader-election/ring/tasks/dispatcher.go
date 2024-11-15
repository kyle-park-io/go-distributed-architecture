package tasks

import (
	"ring/node"
	"ring/send"
	"ring/utils"

	"fmt"
	"time"

	"github.com/spf13/viper"
)

func AssignTask(numNodes int, ch []chan send.Msg) {
	for {
		if node.GetAllNodesInitialized() {
			// kill the leader node
			go func() {
				leader_kill := viper.GetInt("nodes.leader_kill")
				<-time.After(time.Duration(leader_kill) * time.Second)

				fmt.Printf("Instruct the leader node(%d), Die!\n", numNodes)
				ch[numNodes] <- send.Msg{From: send.ADMIN_NODE, MsgType: send.FAILURE}
			}()

			// basic task
			go func() {
				period := viper.GetInt("tasks.period")
				for {
					r := utils.GetRandom(numNodes-1) + 1
					uuid := utils.GenerateUUID()
					fmt.Printf("Send to node %d, let's do task: %s\n", r, uuid)
					ch[r] <- send.Msg{From: send.ADMIN_NODE, MsgType: send.DoTask, Msg: uuid}

					time.Sleep(time.Duration(period) * time.Second)
				}
			}()

			return
		}
	}
}
