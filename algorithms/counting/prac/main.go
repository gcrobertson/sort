package main

import (
	"fmt"
	"math"
)

func main() {

	// xi := []int{2, 5, 4, 3, 3, 3, 0}
	// xi = Counting(xi)
	// fmt.Printf("Counting result: %+v\n", xi)

	A := []int{5, 3, 2, 3}
	fmt.Printf("CountingSortNegatives presort: %+v\n", A)
	A = CountingSortNegatives(A)
	fmt.Printf("CountingSortNegatives result: %+v\n", A)

}

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

// Counting sort
func Counting(A []int) []int {

	max, _ := maxMin(A)
	zeroOffset := 1
	lenB := max + zeroOffset
	var B = make([]int, lenB)

	// step 1: add all 0-max numbers to the indexed slice.
	for i := 0; i < len(A); i++ {
		val := A[i]
		B[val]++
	}

	// step 2: each element at each index stores the sum of previous keys
	for i := 1; i < len(B); i++ {
		B[i] += B[i-1]
	}

	// step 3:
	var C = make([]int, len(A))

	for i := 0; i < len(A); i++ {
		val := A[i]                  // value of A is the key for B.
		index := B[val] - zeroOffset // accomodate for 0 value
		C[index] = val
		B[val]--
	}

	return C
}

// CountingSortNegatives ...
func CountingSortNegatives(A []int) []int {

	var C = make([]int, len(A))
	max, min := maxMin(A)
	zeroOffset := 1

	absMin := int(math.Abs(float64(min)))
	lenB := max + absMin + zeroOffset
	B := make([]int, lenB) // for numbers -3 to 3 lenB = 7

	// step 1: add to B
	for i := 0; i < len(A); i++ {
		val := A[i]
		B[val+absMin]++
	}
	// step 2: tally
	for i := 1; i < len(B); i++ {
		B[i] += B[i-1]
	}

	// step 3: order
	for i := 0; i < len(A); i++ {
		val := A[i]
		key := B[val+absMin] - zeroOffset
		C[key] = val
		B[val+absMin]--
	}

	return C
}
