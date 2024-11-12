package utils

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
)

func GetRandom(num int) int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)
	return r.Intn(num)
}

func GenerateUUID() string {
	return uuid.New().String()
}
