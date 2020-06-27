package algorithms

import (
	"math"
)

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

// CountingSort ...
func CountingSort(A []int) []int {

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
