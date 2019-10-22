package main

import (
	"math/rand"
	"sort"
	"time"
)

type sortFunc func([]int)

var (
	driver           sortFunc
	defaultGoSort    = sort.Ints // QuickSort nlogn
	customBubbleSort = bubleSort
	customMergeSort  = mergeSort
)

func main() {

}

func bubleSort(l []int) {
	for i := range l {
		for j := range l {
			if l[i] < l[j] {
				l[i], l[j] = l[j], l[i]
			}
		}
	}
}

func mergeSort(slice []int) {
	for i, v := range recursiveMergeSort(slice) {
		slice[i] = v
	}
}

func recursiveMergeSort(slice []int) []int {
	merger := func(left, right []int) []int {
		size, i, j := len(left)+len(right), 0, 0
		slice := make([]int, size, size)

		for k := 0; k < size; k++ {
			if i > len(left)-1 && j <= len(right)-1 {
				slice[k] = right[j]
				j++
			} else if j > len(right)-1 && i <= len(left)-1 {
				slice[k] = left[i]
				i++
			} else if left[i] < right[j] {
				slice[k] = left[i]
				i++
			} else {
				slice[k] = right[j]
				j++
			}
		}
		return slice
	}

	if len(slice) < 2 {
		return slice
	}
	mid := (len(slice)) / 2
	return merger(recursiveMergeSort(slice[:mid]), recursiveMergeSort(slice[mid:]))
}

func randomSlice(size int) []int {
	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}
