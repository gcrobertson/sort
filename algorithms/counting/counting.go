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

// @TODO: This func breaks when there are only positive numbers I think... would be a good test.

// Counting sort handles negative numbers by offsetting minimum value to 0 index
func Counting(xi []int) []int {

	max, min := maxMin(xi)
	r := max - min + 1

	// fmt.Printf("max: [%v] min: [%v] range: [%v]\n", max, min, r)

	// 1 key for every int value in range, 0 indexed
	var indexSlice = make([]int, r, r)

	// offset to place values in 0 indexed slice
	offset := int(math.Abs(float64(min)))

	// step 1: add all values to 0 indexed slice
	for i := 0; i < len(xi); i++ {
		indexSlice[xi[i]+offset]++
	}

	// step 2: each element at each index stores the sum of previous keys
	for i := 1; i < len(indexSlice); i++ {
		indexSlice[i] = indexSlice[i] + indexSlice[i-1]
	}

	// step 3:
	result := make([]int, len(xi))
	for i := 0; i < len(xi); i++ {
		result[indexSlice[xi[i]+offset]-1] = xi[i]
		indexSlice[xi[i]+offset]--
	}

	return result
}
