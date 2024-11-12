package monitor

import (
	"bully/send"

	"fmt"
	"time"
)

func MonitorChannelLen(ch []chan send.Msg, interval int) {
	for range time.Tick(time.Duration(interval) * time.Second) {
		for i, v := range ch {
			fmt.Printf("Node %d: channel length: %d\n", i, len(v))
		}
	}
}
