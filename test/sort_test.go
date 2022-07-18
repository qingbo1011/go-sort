package test

import (
	"fmt"
	"go_sort/util"
	"math"
	"sort"
	"testing"
)

func TestBubbleSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(bubbleSort(s))
}

func TestSelectionSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(selectionSort(s))
}

func TestInsertionSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(insertionSort(s))
}

func TestShellSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(shellSort(s))
}

func TestMergeSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(mergeSort(s))
}

func TestQuickSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(quicksort(s))
}

func TestHeapSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(heapSort(s))
}

func TestCountingSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(countingSort(s))
}

func TestBucketSort(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(bucketSort(s))
}

func TestName1(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(s[1:])
}

func TestName2(t *testing.T) {
	s := util.RandSlice(15)
	fmt.Println(s)
	fmt.Println(s[:1])
}

func TestName3(t *testing.T) {
	fmt.Println(math.MaxInt)
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

func bucketSort(s []int) []int {
	result := make([]int, 0)
	// 1.找到最小值和最大值
	min, max := math.MaxInt, math.MinInt
	for _, v := range s {
		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}
	// 2.初始化桶
	bucketNum := (max-min)/len(s) + 1
	bucketList := make([][]int, 0)
	for i := 0; i < bucketNum; i++ {
		bucket := make([]int, 0)
		bucketList = append(bucketList, bucket)
	}
	// 3.把元素放到桶内
	for _, v := range s {
		number := (v - min) / len(s)
		bucketList[number] = append(bucketList[number], v)
	}
	// 4.桶内进行排序
	for _, bucket := range bucketList {
		sort.Ints(bucket) // 桶内排序，可以使用其它排序方法，也可以递归使用桶排序，这里我直接用go内置的sort包
	}
	// 5.对bucketList进行拼接
	for _, bucket := range bucketList {
		for _, v := range bucket {
			result = append(result, v)
		}
	}

	return result
}

func quicksort(s []int) []int {
	if len(s) < 2 {
		return s
	}
	left, right := 0, len(s)-1
	pivot := 0 //
	s[right], s[pivot] = s[pivot], s[right]
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

func shellSort(s []int) []int {
	for gap := len(s) / 2; gap > 0; gap = gap / 2 {
		for i := gap; i < len(s); i++ {
			for j := i; j >= gap && s[j-gap] > s[j]; j = j - gap {
				s[j], s[j-gap] = s[j-gap], s[j]
			}
		}
	}
	return s
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

func bubbleSort(s []int) []int {
	for i := 0; i < len(s); i++ {
		for j := 0; j < len(s)-1-i; j++ {
			if s[j] > s[j+1] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
	return s
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

