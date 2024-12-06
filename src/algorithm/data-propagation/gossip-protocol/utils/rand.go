package utils

import (
	"time"

	"golang.org/x/exp/rand"
)

func GetRandom(num int) int {
	src := rand.NewSource(uint64(time.Now().UnixNano()))
	r := rand.New(src)
	return r.Intn(num)
}
