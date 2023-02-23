package generator

import (
	"math/rand"
	"time"
)

func GenerateArr(arrLength int) []int {
	rand.Seed(time.Now().UnixNano())
	arr := make([]int, arrLength)
	for i := 0; i < arrLength; i++ {
		arr[i] = 1 + rand.Intn(2*arrLength-1+1)
	}
	return arr
}
