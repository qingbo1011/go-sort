package main

import (
	"fmt"
	"sort"
)

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Ints(s)
	fmt.Println(s) // [-3 1 3 3 5 6 8 8 9 15 45 66]
	// 降序排序
	s = []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	sort.Sort(sort.Reverse(sort.IntSlice(s))) // [15 6 8 3 5 9 1 45 66 3 8 -3]
}