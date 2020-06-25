package main

import (
	"fmt"
	"os"
)

func main() {

	// Driver Code
	xi := []int{170, 45, 75, 90, 802, 24, 2, 66}
	n := len(xi)
	SortRadix(xi, n)
	fmt.Println(xi)
}

// SortRadix ...
func SortRadix(xi []int, n int) []int {

	// step 1: 	find the maximum number "k" to know the number of digits "d"
	max, _ := maxMin(xi)
	fmt.Printf("maximum number found is: %d\n", max)

	// step 2: 	do a counting sort for every digit, i
	//			exp = 10 ^ i, instead of passing digit number, exp is passed
	for exp := 1; max/exp > 0; exp *= 10 {

		fmt.Printf("calling CountingSort with vars: [%v] [%d] [%d]\n", xi, n, exp)
		CountingSort(xi, n, exp)
	}

	return xi
}

// CountingSort handles negative numbers by offsetting minimum value to 0 index
//	xi 	slice to sort
//	n	length of slice xi
//	exp	exponent / digit of slice value to sort [1,10,100 for numbers 0-999, 1 sorts up to 9, 10 up to 99, 100 up to 999]
func CountingSort(xi []int, n, exp int) []int {

	// 1 key for every int value in range, 0 indexed
	var indexSlice = make([]int, 10) // should it not always hold 10 for indices 0-9?

	// step 1: add all values to 0 indexed slice
	for i := 0; i < len(xi); i++ {
		value := xi[i] / exp % 10 //			fmt.Printf("found value [%v] for slice value: [%v]\n", value, xi[i])
		indexSlice[value]++
	}
	// fmt.Printf("the indexSlice after step 1: [%+v]\n", indexSlice)
	// // step 2: each element at each index stores the sum of previous keys
	for i := 1; i < len(indexSlice); i++ {
		indexSlice[i] = indexSlice[i] + indexSlice[i-1]
	}
	// fmt.Printf("the indexSlice after step 2: [%+v]\n", indexSlice)

	// // step 3: build the output array. i have removed offset but would be great to add back in.
	result := make([]int, len(xi))
	for i := 0; i < len(xi); i++ {
		// value := indexSlice[xi[i]] / exp % 10
		// result[value] = xi[i]
		// indexSlice[value]--
	}

	fmt.Printf("the indexSlice after step 3: [%+v]\n", indexSlice)

	fmt.Printf("the result after step 3: [%+v]\n", result)

	os.Exit(5)

	// return result

	return []int{}
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

// go over counting sort WITHOUT an offset
// go over counting sort WITH    an offset

// i will have a better understanding... THEN

// create a radix sort WITH 	 an offset
// create a radix sort WTIHOUT 	 an offset
