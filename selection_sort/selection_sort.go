package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(selectionSort(s))
}

func selectionSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ { // 找到索引i的右边最小值索引
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i] // 交换i和min
	}
	return s
}
