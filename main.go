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
	sorts *string = flag.String("sort", "", "sort method. available options: bubble, selection, sinking")
)

//Sorting algorithms
var algorithms = map[string]bool{
	"bubble":   false,
	"counting": false,
	"merge":    false,
}

func parseCLI() {

	// retrieve command line arguments
	flag.Parse()
}

func validateCLISize() {

	// validate command line argument for size
	if 1 > *size || *size > 100000 {
		*size = 5
	}
}

func validateCLISorts() {

	// validate command line argument for sort. this sets the hash map to `true`
	s := strings.Split(*sorts, ",")
	for sort := range s {
		if _, ok := algorithms[s[sort]]; ok {
			algorithms[s[sort]] = true
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

func main() {

	parseCLI()
	validateCLISize()
	validateCLISorts()

	xi := initializeIntSlice()

	fmt.Printf("I should run these numbers through! [%+v]\n", xi)

	// step 1: just run the algorithm for any of the sorts...

	// run()
}
