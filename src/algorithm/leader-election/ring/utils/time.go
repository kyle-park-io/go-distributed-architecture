package utils

import (
	"ring/send"

	"time"
)

type Duration struct {
	StartTime   time.Time
	EndTime     time.Time
	ElapsedTime time.Duration
}

func GetTimeAfterSeconds(d int) Duration {
	start := time.Now()
	duration := time.Duration(d) * time.Second
	end := start.Add(duration)

	return Duration{StartTime: start, EndTime: end, ElapsedTime: duration}
}

func SetTimeout(d int, ch chan send.Msg, msg send.Msg) {
	// defer close(ch)
	<-time.After(time.Duration(d) * time.Second)
	// time.Sleep(time.Duration(d) * time.Second)
	ch <- msg
}
