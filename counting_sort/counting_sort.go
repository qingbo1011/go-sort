package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8, -3}
	fmt.Println(s)
	fmt.Println(countingSort(s))
}

func countingSort(s []int) []int {
	min, max := getMaxAndMin(s)
	bias := 0 - min // 引入bias，是为了解决s中最小值为负数的情况
	result := make([]int, 0)
	// arrayOfCounts index-bias为s中的某个数，value为该数在s中的个数
	arrayOfCounts := make([]int, max-min+1)
	for i, _ := range s { // 就算s中有负数，处理后s中的min也为0
		s[i] = s[i] + bias
	}
	for _, v := range s {
		arrayOfCounts[v]++
	}
	for i, _ := range arrayOfCounts {
		for arrayOfCounts[i] > 0 {
			result = append(result, i-bias) // 这里别忘了要减去bias
			arrayOfCounts[i]--
		}
	}
	return result
}

func getMaxAndMin(s []int) (min int, max int) {
	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	return
}
