# Merge Sort
# https://www.programiz.com/dsa/merge-sort

- Divide and Conquer Algorithm
- One of the most popular sorting algorithms
- A great way to build confidence in recursive algorithms

# Divide and Conquer Algorithm
- Algorithm splits an input into various pieces, sorts them and then merges them back together.


# Merge Sort Implementation Overview

- We divide a problem into subproblems.
- When the solution to each subproblem is ready, we combine the results to solve the main problem.


var MergeSortAlgorithm = SortingAlgorithm{
	Name:           "Merge Sort",
	Author:         "John von Neumann",
	Year:           1945,
	AlgorithmRun:   MergeSort,
	AlgorithmType:  "Divide and Conquer",
	ComparisonSort: true,
	Alias:          []string{"mergesort"},
	Stable:         true, // Most implementations
}








//MergeSort ...
func MergeSort(arr []int) {
	copy(arr, recursiveDivide(arr))
}

func recursiveDivide(items []int) []int {

	var num = len(items)

	if num == 1 {
		return items
	}

	middle := int(num / 2)
	var (
		left  = make([]int, middle)
		right = make([]int, num-middle)
	)
	for i := 0; i < num; i++ {
		if i < middle {
			left[i] = items[i]
		} else {
			// fmt.Println(i, middle, i-middle)
			right[i-middle] = items[i]
		}
	}

	return merge(recursiveDivide(left), recursiveDivide(right))
}

func merge(left, right []int) (result []int) {
	// fmt.Printf("called with left:%v and right%v\n", left, right)
	result = make([]int, len(left)+len(right))

	i := 0
	for len(left) > 0 && len(right) > 0 {
		if left[0] < right[0] {
			result[i] = left[0]
			left = left[1:]
		} else {
			result[i] = right[0]
			right = right[1:]
		}
		i++
	}

	for j := 0; j < len(left); j++ {
		result[i] = left[j]
		i++
	}
	for j := 0; j < len(right); j++ {
		result[i] = right[j]
		i++
	}

	return
}

/*
//MergeSort ...
func MergeSort(arr []int) {
	r := len(arr) - 1
	// p := 0
	// q := r / 2
	MergeSortRecursion(arr, 0, r)
}
//MergeSortRecursion ...
func MergeSortRecursion(arr []int, p, r int) {
	if p > r {
		return
	}
	q := (p + r) / 2
	MergeSortRecursion(arr, p, q)
	MergeSortRecursion(arr, q+1, r)
	merge(arr, p, q, r)
}
func merge(arr []int, p, q, r int) {
	// fmt.Printf("arr: %v\np:%v\nq:%v\nr:%v\n", len(arr), p, q, r)
	lx := make([]int, q-p, q-p)
	mx := make([]int, r-q, r-q)
	// fmt.Println("len lx:", len(lx))
	// fmt.Println("len mx:", len(mx))
	copy(lx, arr[0:q-p])
	copy(mx, arr[r-q:])
	HeapSort(lx)
	HeapSort(mx)
	// fmt.Println("lx:", lx)
	// fmt.Println("mx:", mx)
	var i, j int
	k := p
	//		i := current index of L
	//		j := current index of M
	//		k := current index of A[p..r]
	// fmt.Println("before:", arr)
	for i < len(lx) && j < len(mx) {
		if lx[i] <= mx[j] {
			arr[k] = lx[i]
			i++
		} else {
			arr[k] = mx[j]
			j++
		}
		k++
	}
	for i < len(lx) {
		arr[k] = lx[i]
		i++
		k++
	}
	for j < len(mx) {
		arr[k] = mx[j]
		j++
		k++
	}
	// fmt.Println("len merged:", len(arr))
	// fmt.Println("merged:", arr)
	// fmt.Println("value of k equals len(arr)?:", k, len(arr))
	return
}
*/



