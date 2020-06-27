package main

import (
	"fmt"
	"math"
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
	for exp := 1; max/exp > 0; exp = exp * 10 {
		CountingSort(A, lenA, exp)
	}
}

// CountingSort ...	breaks on negative numbers
func CountingSort(A []int, lenA, exp int) {

	var B = make([]int, lenA)
	var C = make([]int, 10) // @TODO: handles 0 to 9, i think i need to handle -9 to 9

	// Store count of occurrences in count[]
	for i := 0; i < lenA; i++ {

		// 170 / 1 = 170, 170 % 10 = 0
		//  45 / 1 =  45,  45 % 10 = 5
		// 802 / 100
		// fmt.Printf("( %v / %v ) %% 10 = %v\n", A[i], exp, (A[i]/exp)%10)
		// var1 := 802 / 100	fmt.Printf("%T %v", var1, var1)	// int 8

		// so the modulus part is working. what is not working is an implied offset.

		C[(A[i]/exp)%10]++ // @TODO: count needs to have implied offset of -9 or |min| in range
	}

	for i := 1; i < 10; i++ {
		C[i] += C[i-1]
	}

	for i := lenA - 1; i >= 0; i-- {
		B[C[(A[i]/exp)%10]-1] = A[i]
		C[(A[i]/exp)%10]--
	}

	for i := 0; i < lenA; i++ {
		A[i] = B[i]
	}
}

// CountingSortNegatives ...
func CountingSortNegatives(A []int) []int {

	var B = make([]int, len(A))
	max, min := maxMin(A)
	zeroOffset := 1

	absMin := int(math.Abs(float64(min)))
	lenC := max + absMin + zeroOffset
	C := make([]int, lenC) // for numbers -3 to 3 lenB = 7

	// step 1: add to C
	for i := 0; i < len(A); i++ {
		val := A[i]
		C[val+absMin]++
	}
	// step 2: tally
	for i := 1; i < lenC; i++ {
		C[i] += C[i-1]
	}

	// step 3: order
	for i := 0; i < len(A); i++ {
		val := A[i]
		key := C[val+absMin] - zeroOffset
		B[key] = val
		C[val+absMin]--
	}

	return B
}
