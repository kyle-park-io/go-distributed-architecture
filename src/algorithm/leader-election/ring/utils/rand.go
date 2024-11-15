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

func ShuffleSlice(arr []int) []int {
	src := rand.NewSource(time.Now().UnixNano())
	r := rand.New(src)

	shuffledArr := make([]int, len(arr))
	copy(shuffledArr, arr)

	r.Shuffle(len(shuffledArr), func(i, j int) {
		shuffledArr[i], shuffledArr[j] = shuffledArr[j], shuffledArr[i]
	})

	return shuffledArr
}
