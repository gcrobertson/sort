// Merge Sort in Golang
package main

import (
	"fmt"
)

func main() {
	//slice := numbers.GenerateSlice(20)

	slice := []int{5, 4, 3, 2, 1, 0}

	fmt.Println("\n--- Unsorted --- \n\n", slice)
	fmt.Println("\n--- Sorted ---\n\n", mergeSort(slice))
}

func mergeSort(xi []int) []int {

	// avoid stack overflow from recursion
	if len(xi) == 1 {
		return xi
	}

	// find mid point of slice and break it into two
	var (
		mid   = int(len(xi) / 2)
		left  = xi[0:mid]
		right = xi[mid:]
	)

	return mergeSortMerge(mergeSort(left), mergeSort(right))
}

func mergeSortMerge(left, right []int) (result []int) {

	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for _, v := range left {
		result[i] = v
		i++
	}
	for _, v := range right {
		result[i] = v
		i++
	}

	return
}
