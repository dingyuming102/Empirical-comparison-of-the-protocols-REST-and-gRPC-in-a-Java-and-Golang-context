package utils

import (
	"math/rand"
)

func SmallTransInitData(size int) []int32 {
	data := make([]int32, size) // 二维切片，3行

	for i := 0; i < size; i++ {
		data[i] = rand.Int31n(100)
	}

	return data
}

func Square(data int32) int32 {
	return data * data
}
