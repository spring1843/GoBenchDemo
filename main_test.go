package main

import (
	"reflect"
	"testing"
)

func init() {
	driver = defaultGoSort
}

func benchmarkSort(size int, b *testing.B) {
	slice := randomSlice(size)

	for n := 0; n < b.N; n++ {
		driver(slice)
	}
}

func BenchmarkSort10(b *testing.B)      { benchmarkSort(10, b) }
func BenchmarkSort100(b *testing.B)     { benchmarkSort(100, b) }
func BenchmarkSort500(b *testing.B)     { benchmarkSort(500, b) }
func BenchmarkSort1000(b *testing.B)    { benchmarkSort(1000, b) }
func BenchmarkSort10000(b *testing.B)   { benchmarkSort(10000, b) }
func BenchmarkSort100000(b *testing.B)  { benchmarkSort(100000, b) }
func BenchmarkSort1000000(b *testing.B) { benchmarkSort(1000000, b) }

func TestCanActuallySort(t *testing.T) {
	const size = 100
	var (
		duplicate = func(a []int) []int {
			slice := make([]int, len(a), len(a))
			for i, v := range a {
				slice[i] = v
			}
			return slice
		}

		slice1 = randomSlice(size)
		slice2 = duplicate(slice1)
		slice3 = duplicate(slice1)
	)

	defaultGoSort(slice1)

	customBubbleSort(slice2)
	if !reflect.DeepEqual(slice1, slice2) {
		t.Fatalf("Bubble sort did not work. %v is not sorted", slice2)
	}

	customMergeSort(slice3)
	if !reflect.DeepEqual(slice1, slice3) {
		t.Fatalf("Merge sort did not work. \n\n%v is not sorted \n\n%v", slice3, slice1)
	}
}
