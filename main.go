/*
 *	clear && go run main.go -sort=merge,heap,bubble,insertion,shell,counting,selection,quick  -size=9000
 *	clear && go run main.go -sort=sinking,insertion,bubble,selection -size=110000
 *	clear && go run main.go -sort=bubble -size=5
 */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

//Command line arguments
var (
	size          = flag.Int("size", 0, "data size of random integer array to be sorted.")
	sorts *string = flag.String("sorts", "", "sort method. available options: bubble, selection, sinking")
)

//Sorting algorithms
var sortMap = map[string]bool{
	"bubble":   false,
	"counting": false,
	"merge":    false,
}

// retrieve command line arguments
func parseCLI() {

	flag.Parse()
}

// validate command line argument for size
func validateCLISize() {

	if 1 > *size || *size > 100000 {
		*size = 5
	}
}

// validate command line argument for sort. this sets the hash map to `true`
func validateCLISorts() {

	s := strings.Split(*sorts, ",")
	for sort := range s {
		if _, ok := sortMap[s[sort]]; ok {
			sortMap[s[sort]] = true
		}
	}
}

func initializeIntSlice() []int {

	var slice = make([]int, *size, *size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < *size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

//SortInfo ...
type SortInfo struct {
	AlgorithmName  string
	UntouchedCopy  []int
	ArgumentToFunc []int
	OrderedSlice   []int
}

//SortInfoSlice ...
var SortInfoSlice []SortInfo

func main() {

	parseCLI()
	validateCLISize()
	validateCLISorts()
	setupSortInfoSlice()

	fmt.Printf("contents of the SortInfoSlice: [%+v]\n", SortInfoSlice)
}

func setupSortInfoSlice() {

	originalSlice := initializeIntSlice()

	for k := range sortMap {

		if sortMap[k] == true {

			nSortInfo := SortInfo{
				AlgorithmName: k, // set up name from map key
			}

			nSortInfo.UntouchedCopy = make([]int, len(originalSlice))
			nSortInfo.ArgumentToFunc = make([]int, len(originalSlice))

			copy(nSortInfo.UntouchedCopy, originalSlice)
			copy(nSortInfo.ArgumentToFunc, originalSlice)

			SortInfoSlice = append(SortInfoSlice, nSortInfo)
		}
	}
}

func run() {

	// bubbleSort.OrderedSlice = algorithms.Bubble(bubbleSort.ArgumentToFunc)

	// fmt.Printf("Original list: [%+v]\n", prexi)
	// fmt.Printf("Bubble Sort presorted: [%+v]\n", bubbleSort.ArgumentToFunc)
	// fmt.Printf("Bubble Sort postsorted: [%+v]\n", bubbleSort.OrderedSlice)

}
