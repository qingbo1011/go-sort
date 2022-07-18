package main

import (
	"fmt"
	"math"
	"sort"
)

func main() {
	s := []int{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(bucketSort(s))
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
