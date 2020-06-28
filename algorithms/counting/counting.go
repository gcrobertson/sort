package algorithms

import (
	"math"
)

// CountingSort ...
func CountingSort(A []int) []int {

	max, min := A[0], A[0]
	for i := 0; i < len(A); i++ {
		if A[i] > max {
			max = A[i]
		} else if A[i] < min {
			min = A[i]
		}
	}
	zeroOffset := 1
	absMin := int(math.Abs(float64(min)))
	lenC := max + absMin + zeroOffset

	var B = make([]int, len(A))
	var C = make([]int, lenC)

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
