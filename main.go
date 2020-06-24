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

	// fix with go module

	algorithms "github.com/sort/algorithms/bubble"

	algorithms2 "github.com/sort/algorithms/counting"

	algorithms3 "github.com/sort/algorithms/merge"
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

	if 1 > *size || *size > 10000000 {
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
	StartTime      time.Time
	SortDuration   time.Duration
}

//SortInfoSlice ...
var SortInfoSlice []SortInfo

func main() {

	parseCLI()
	validateCLISize()
	validateCLISorts()
	setupSortInfoSlice()

	// maybe run should take a *SortInfo
	// log the begin time end time to SortInfo struct?
	// run as a go routine?

	for k := range SortInfoSlice {

		run(&SortInfoSlice[k])

	}

	for k, v := range SortInfoSlice {
		// fmt.Printf("contents of the SortInfoSlice index [%d]: [%+v]\n", k, v)
		fmt.Printf("sort [%s]: processed in [%+v]\n", v.AlgorithmName, SortInfoSlice[k].SortDuration)
	}

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

			copy(nSortInfo.UntouchedCopy, originalSlice) // Probably unnecessary to hold onto this but for now it's okay
			copy(nSortInfo.ArgumentToFunc, originalSlice)

			SortInfoSlice = append(SortInfoSlice, nSortInfo)
		}
	}
}

func run(v *SortInfo) {

	v.StartTime = time.Now()

	switch v.AlgorithmName {
	case "bubble":
		v.OrderedSlice = algorithms.Bubble(v.ArgumentToFunc)
	case "counting":
		v.OrderedSlice = algorithms2.Counting(v.ArgumentToFunc)
	case "merge":
		v.OrderedSlice = algorithms3.MergeSort(v.ArgumentToFunc)
	}

	v.SortDuration = time.Since(v.StartTime)
}
