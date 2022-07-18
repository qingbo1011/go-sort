package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(mergeSort(s))
}

func mergeSort(s []int) []int {
	if len(s) <= 1 { // 递归终止条件
		return s
	}

	middle := len(s) / 2
	left := mergeSort(s[:middle])
	right := mergeSort(s[middle:])

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	result := make([]int, 0, len(left)+len(right))
	for len(left) > 0 || len(right) > 0 {
		// 对len(left)=0和len(right)=0情况的处理
		if len(left) == 0 {
			return append(result, right...)
		}
		if len(right) == 0 {
			return append(result, left...)
		}

		if left[0] < right[0] {
			result = append(result, left[0])
			left = left[1:] // left切片去掉left[0]元素
		} else {
			result = append(result, right[0])
			right = right[1:] // right切片去掉right[0]元素
		}
	}

	return result
}
