package main

import (
	"fmt"
	"math"
)

// https://www.geeksforgeeks.org/radix-sort/
// The idea of Radix Sort is to do digit by digit sort starting from least significant digit to most significant digit. Radix sort uses counting sort as a subroutine to sort.

func max(xi []int) (max int) {
	max = xi[0]
	for i := 0; i < len(xi); i++ {
		if xi[i] > max {
			max = xi[i]
		}
	}
	return
}

func main() {

	var A = []int{170, 45, 75, 90, 802, 24, 2, 66, -900}

	fmt.Printf("pre: %v\n", A)

	RadixSort(A, len(A))

	fmt.Printf("post: %v\n", A)
}

// RadixSort ...
func RadixSort(A []int, lenA int) {
	max := max(A)
	for exp := 1; max/exp > 0; exp = exp * 10 {
		CountingSort(A, lenA, exp)
	}
}

func minMax(A []int, lenA, exp int) (int, int) {
	var min, max int = (A[0] / exp) % 10, (A[0] / exp) % 10
	for i := 0; i < lenA; i++ {
		// fmt.Printf("Comparing number: %v from base number %v\n", (A[i]/exp)%10, A[i])
		if (A[i]/exp)%10 < min {
			min = (A[i] / exp) % 10
		} else if (A[i]/exp)%10 > max {
			max = (A[i] / exp) % 10
		}
	}
	// fmt.Printf("the max number is %v and the min number is %v\n", max, min)
	return min, max
}

// CountingSort ... works with negative numbers now!
func CountingSort(A []int, lenA, exp int) {

	min, max := minMax(A, lenA, exp)
	absMin := int(math.Abs(float64(min)))
	zeroOffset := 1
	lenC := max + absMin + zeroOffset

	var B = make([]int, lenA)
	var C = make([]int, lenC)

	// Store count of occurrences in count[]
	for i := 0; i < lenA; i++ {
		expVal := (A[i] / exp) % 10
		C[expVal+absMin]++
	}

	for i := 1; i < lenC; i++ {
		C[i] += C[i-1]
	}

	for i := lenA - 1; i >= 0; i-- {
		expVal := (A[i] / exp) % 10
		B[C[expVal+absMin]-1] = A[i]
		C[expVal+absMin]--
	}

	copy(A, B)
}
