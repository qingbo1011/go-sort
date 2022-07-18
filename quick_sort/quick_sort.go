package main

import "fmt"

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(quicksort(s))
}

func quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	left, right := 0, len(s)-1
	pivot := 0 // pivot可以在[0,len(s)-1]之间随机取
	s[pivot], s[right] = s[right], s[pivot]
	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}
	s[left], s[right] = s[right], s[left]

	quicksort(s[:left])
	quicksort(s[left+1:])

	return s
}
