package utils

import (
	"math/rand"
	"time"
)

func GetRandom(num int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(num)
}
