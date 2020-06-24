package algorithms

import "math/rand"

//QuickSort ...
func QuickSort(xi []int) []int {

	recursionQuickSort(xi)

	// Returning a copy.. is this necessary?
	return xi
}

//QuickSort ...
func recursionQuickSort(a []int) {

	if len(a) < 2 {
		return
	}
	var left int
	right := len(a) - 1
	pivot := rand.Intn(len(a)) % len(a)
	a[pivot], a[right] = a[right], a[pivot]
	for i := 0; i < len(a); i++ {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]
	recursionQuickSort(a[:left])
	recursionQuickSort(a[left+1:])
}
