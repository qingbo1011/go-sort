package util

import (
	"math/rand"
	"time"
)

// RandSlice 生成一个长度为n的随机int切片
func RandSlice(n int) []int {
	// needed a seed input else it will generate the same number
	rand.Seed(time.Now().UnixNano())

	s := make([]int, n)
	for i := 0; i <= n-1; i++ {
		s[i] = rand.Intn(n)
	}
	return s
}
