package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(heapSort(s))
}

func heapSort(s []int) []int {
	for i := len(s)/2 - 1; i >= 0; i-- {
		s = sift(s, i, len(s))
	}
	for i := len(s) - 1; i >= 1; i-- {
		s[0], s[i] = s[i], s[0]
		s = sift(s, 0, i)
	}
	return s
}

func sift(s []int, i int, length int) []int {
	maxChild := 0
	for i*2+1 < length {
		if (i*2+1 == length-1) || (s[i*2+1] > s[i*2+2]) {
			maxChild = i*2 + 1
		} else {
			maxChild = i*2 + 2
		}
		if s[i] < s[maxChild] {
			s[i], s[maxChild] = s[maxChild], s[i]
			i = maxChild
		} else {
			break
		}
	}
	return s
}
