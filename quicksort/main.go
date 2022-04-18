package main

import (
	"fmt"
	"time"
)

func quickSort(arrays []int, start, end int) {
	if arrays == nil || len(arrays) == 0 {
		return
	}
	if start >= end {
		return
	}
	left, right := start, end
	privot := arrays[(left+right)/2]
	for left < right {
		for left <= right && arrays[left] < privot {
			left++
		}
		for left <= right && arrays[right] > privot {
			right--
		}
		if left <= right {
			arrays[left], arrays[right] = arrays[right], arrays[left]
		}
		left++
		right--
	}
	quickSort(arrays, left, end)
	quickSort(arrays, start, right)
}

func main() {
	ints := []int{4,5,3,4,9,4}
	//ints := []int{1,1,1,1,1,2,1}
	fmt.Println(ints)
	quickSort(ints, 0, len(ints)-1)
	fmt.Println(ints)
	nano := time.Now().UTC().Unix()
	unixNano := time.Now().Unix()
	fmt.Println(nano)
	fmt.Println(unixNano)

	add := time.Now()
	fmt.Println(add.Before(add))
}
