package utils

import (
	"math"
	"time"
)

func Mean(arr []time.Duration) int64 {
	result := time.Nanosecond * 0
	var count int64 = 0
	for _, data := range arr {
		if data != 0 {
			result += data
			count++
		} else {
			result += time.Nanosecond * 100

		}
	}
	return result.Nanoseconds() / int64(len(arr))
}

func Min(arr []time.Duration) int64 {
	result := time.Second * 1000
	for _, data := range arr {
		if data != 0 && data < result {
			result = data
		}

	}
	return result.Nanoseconds()
}

func Max(arr []time.Duration) int64 {
	result := time.Nanosecond * 0
	for _, data := range arr {
		if data != 0 && data > result {
			result = data
		}

	}
	return result.Nanoseconds()
}

func Variance(v []time.Duration) float64 {
	var res float64 = 0
	var m = float64(Mean(v))
	var n int = len(v)
	for i := 0; i < n; i++ {
		if v[i] != 0 {
			res += (float64(v[i]) - m) * (float64(v[i]) - m)
		} else {
			res += (float64(100) - m) * (float64(100) - m)
		}
	}
	return res / float64(n-1)
}

func StdDev(arr []time.Duration) float64 {
	v := Variance(arr)
	return math.Sqrt(v)
}
