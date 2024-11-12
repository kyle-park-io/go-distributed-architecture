package tasks

import "time"

func DoTask(s int) {
	time.Sleep(time.Duration(s) * time.Second)
}
