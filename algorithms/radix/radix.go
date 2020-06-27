package main

import (
	"fmt"
)

// https://www.geeksforgeeks.org/radix-sort/
// The idea of Radix Sort is to do digit by digit sort starting from least significant digit to most significant digit. Radix sort uses counting sort as a subroutine to sort.

func maxMin(xi []int) (max int, min int) {
	max, min = xi[0], xi[0]
	for i := 0; i < len(xi); i++ {
		if xi[i] > max {
			max = xi[i]
		} else if xi[i] < min {
			min = xi[i]
		}
	}
	return
}

func main() {

	var A = []int{170, 45, 75, 90, 802, 24, 2, 66}

	fmt.Printf("pre: %v\n", A)

	RadixSort(A, len(A))

	fmt.Printf("post: %v\n", A)
}

// RadixSort ...
func RadixSort(A []int, lenA int) {

	max, _ := maxMin(A)

	fmt.Printf("max number: %d\n", max)

	for exp := 1; max/exp > 0; exp = exp * 10 {

		fmt.Printf("A: %v exp: %v\n", A, exp)

		CountingSort(A, lenA, exp)

		fmt.Printf("A: %v exp: %v\n", A, exp)
	}

}

// CountingSort ...	breaks on negative numbers
func CountingSort(A []int, lenA, exp int) {

	var output = make([]int, lenA)
	var count = make([]int, 10)

	// Store count of occurrences in count[]
	for i := 0; i < lenA; i++ {
		count[(A[i]/exp)%10]++
	}

	for i := 1; i < 10; i++ {
		count[i] += count[i-1]
	}

	for i := lenA - 1; i >= 0; i-- {
		output[count[(A[i]/exp)%10]-1] = A[i]
		count[(A[i]/exp)%10]--
	}

	for i := 0; i < lenA; i++ {
		A[i] = output[i]
	}
}
