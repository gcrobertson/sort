/*
 *	clear && go run main.go -sorts=bubble,counting,heap,insertion,merge,quick,radix,selection,shell -size=100000
 *
 *	clear && go run main.go -sorts=counting,radix -size=20 -range=99999 -natural=true
 *
 */

package main

import (
	"flag"
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/sort/algorithms"
)

//SortInfo ...
type SortInfo struct {
	AlgorithmName  string
	ArgumentToFunc []int
	OrderedSlice   []int
	StartTime      time.Time
	SortDuration   time.Duration
}

//SortInfoSlice ...
var SortInfoSlice []SortInfo

//Command line arguments
var (
	xrange           = flag.Int("range", 999, "data range of random integer array. 0 - range.")
	size             = flag.Int("size", 0, "data size of random integer array to be sorted.")
	sorts    *string = flag.String("sorts", "", "sort method. available options: bubble, counting, heap, insertion, merge, quick, radix, selection, shell")
	xnatural         = flag.Bool("natural", false, "natural numbers only, including 0 and limited to `range` argument.")
)

//Sorting algorithms
var sortMap = map[string]bool{
	"bubble":    false,
	"counting":  false,
	"heap":      false,
	"insertion": false,
	"merge":     false,
	"quick":     false,
	"radix":     false,
	"selection": false,
	"shell":     false,
}

func main() {

	flag.Parse()
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

func run(v *SortInfo, semaphore chan bool) {

	v.StartTime = time.Now()

	switch v.AlgorithmName {
	case "bubble":
		v.OrderedSlice = algorithms.BubbleSort(v.ArgumentToFunc)
	case "counting":
		v.OrderedSlice = algorithms.CountingSort(v.ArgumentToFunc)
	case "heap":
		v.OrderedSlice = algorithms.HeapSort(v.ArgumentToFunc)
	case "insertion":
		v.OrderedSlice = algorithms.InsertionSort(v.ArgumentToFunc)
	case "merge":
		v.OrderedSlice = algorithms.MergeSort(v.ArgumentToFunc)
	case "quick":
		v.OrderedSlice = algorithms.QuickSort(v.ArgumentToFunc)
	case "radix":
		v.OrderedSlice = algorithms.RadixSort(v.ArgumentToFunc, len(v.ArgumentToFunc))
	case "selection":
		v.OrderedSlice = algorithms.SelectionSort(v.ArgumentToFunc)
	case "shell":
		v.OrderedSlice = algorithms.ShellSort(v.ArgumentToFunc)
	}

	v.SortDuration = time.Since(v.StartTime)

	fmt.Printf("sort [%s]: processed in [%+v]\n", v.AlgorithmName, v.SortDuration)

	semaphore <- true
}

// validate command line argument for size
func validateCLISize() {

	if 1 > *size || *size > 1000000000 {
		*size = 10
	}

	if 1 > *xrange || *xrange > 999999 {
		*xrange = 999
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
		slice[i] = rand.Intn(*xrange)
		if *xnatural == false {
			slice[i] -= rand.Intn(*xrange)
		}
	}
	return slice
}

func setupSortInfoSlice() {

	originalSlice := initializeIntSlice()

	for k := range sortMap {

		if sortMap[k] == true {

			nSortInfo := SortInfo{
				AlgorithmName: k, // set up name from map key
			}

			nSortInfo.ArgumentToFunc = make([]int, len(originalSlice))

			copy(nSortInfo.ArgumentToFunc, originalSlice)

			SortInfoSlice = append(SortInfoSlice, nSortInfo)
		}
	}
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
