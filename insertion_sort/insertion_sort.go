package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(insertionSort(s))
}

func insertionSort(s []int) []int {
	for i := 1; i < len(s); i++ {
		for j := 0; j < i; j++ { // 索引i往i左边的区间插入
			// 最开始的区间，只有索引为0的元素(j)，索引为1(i)的元素往该区间插入
			if s[i] < s[j] {
				s[j], s[i] = s[i], s[j]
			}
		}
	}
	return s
}
