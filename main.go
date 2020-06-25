/*
 *	clear && go run main.go -sorts=bubble,counting,heap,insertion,merge,quick,selection,shell -size=100000
 */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	// @TODO: fix imports go module

	algorithms "github.com/sort/algorithms/bubble"

	algorithms2 "github.com/sort/algorithms/counting"

	algorithms3 "github.com/sort/algorithms/merge"

	algorithms4 "github.com/sort/algorithms/quick"

	algorithms5 "github.com/sort/algorithms/heap"

	algorithms6 "github.com/sort/algorithms/shell"

	algorithms7 "github.com/sort/algorithms/selection"

	algorithms8 "github.com/sort/algorithms/insertion"
)

//Command line arguments
var (
	size          = flag.Int("size", 0, "data size of random integer array to be sorted.")
	sorts *string = flag.String("sorts", "", "sort method. available options: bubble, selection, sinking")
)

//Sorting algorithms
var sortMap = map[string]bool{
	"bubble":    false,
	"counting":  false,
	"heap":      false,
	"insertion": false,
	"merge":     false,
	"quick":     false,
	"selection": false,
	"shell":     false,
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

	semaphore := make(chan bool, len(SortInfoSlice))

	for k := range SortInfoSlice {
		go run(&SortInfoSlice[k], semaphore)
	}

	// blocks execution until all channels come back.
	for i := 0; i < len(SortInfoSlice); i++ {
		<-semaphore
	}
	close(semaphore)

	fmt.Println("Performing validation...")

	semaphore = make(chan bool, len(SortInfoSlice))
	for k := range SortInfoSlice {
		go validateSortInfo(SortInfoSlice[k], semaphore)
	}

	// blocks execution until all channels come back.
	for i := 0; i < len(SortInfoSlice); i++ {
		<-semaphore
	}
	close(semaphore)

	fmt.Println("Exiting func main...")
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

func run(v *SortInfo, semaphore chan bool) {

	v.StartTime = time.Now()

	switch v.AlgorithmName {
	case "bubble":
		v.OrderedSlice = algorithms.Bubble(v.ArgumentToFunc)
	case "counting":
		v.OrderedSlice = algorithms2.Counting(v.ArgumentToFunc)
	case "heap":
		v.OrderedSlice = algorithms5.HeapSort(v.ArgumentToFunc)
	case "insertion":
		v.OrderedSlice = algorithms8.InsertionSort(v.ArgumentToFunc)
	case "merge":
		v.OrderedSlice = algorithms3.MergeSort(v.ArgumentToFunc)
	case "quick":
		v.OrderedSlice = algorithms4.QuickSort(v.ArgumentToFunc)
	case "selection":
		v.OrderedSlice = algorithms7.SelectionSort(v.ArgumentToFunc)
	case "shell":
		v.OrderedSlice = algorithms6.ShellSort(v.ArgumentToFunc)
	}

	v.SortDuration = time.Since(v.StartTime)

	fmt.Printf("sort [%s]: processed in [%+v]\n", v.AlgorithmName, v.SortDuration)

	semaphore <- true
}

func validateSortInfo(si SortInfo, semaphore chan bool) {

	if len(si.ArgumentToFunc) != len(si.OrderedSlice) {
		fmt.Printf("sort [%s]: input size %d != ordered list size %d\n", si.AlgorithmName, si.ArgumentToFunc, si.OrderedSlice)
		return
	}

	var err int
	for i := 1; i < len(si.OrderedSlice); i++ {
		if si.OrderedSlice[i-1] > si.OrderedSlice[i] {
			err++
		}
	}
	fmt.Printf("sort [%s]: had %d errors!\n", si.AlgorithmName, err)

	semaphore <- true
}
