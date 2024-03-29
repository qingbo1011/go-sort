package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
)

const digit = 4
const maxbit = -1 << 31

func main() {
	s := []int32{15, 6, 8, 3, 5, 9, 1, 45, 66, 3, 8}
	fmt.Println(s)
	fmt.Println(radixSort(s))
}

func radixSort(data []int32) []int32 {
	buf := bytes.NewBuffer(nil)
	ds := make([][]byte, len(data))
	for i, e := range data {
		binary.Write(buf, binary.LittleEndian, e^maxbit)
		b := make([]byte, digit)
		buf.Read(b)
		ds[i] = b
	}
	countingSort := make([][][]byte, 256)
	for i := 0; i < digit; i++ {
		for _, b := range ds {
			countingSort[b[i]] = append(countingSort[b[i]], b)
		}
		j := 0
		for k, bs := range countingSort {
			copy(ds[j:], bs)
			j += len(bs)
			countingSort[k] = bs[:0]
		}
	}
	var w int32
	for i, b := range ds {
		buf.Write(b)
		binary.Read(buf, binary.LittleEndian, &w)
		data[i] = w ^ maxbit
	}
	return data
}
