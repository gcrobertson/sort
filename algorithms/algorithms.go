package algorithms

import (
	"math"
	"math/rand"
)

/*BubbleSort ...
 *
 *
 */
func BubbleSort(xi []int) []int {

	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi)-i-1; j++ {
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}

	return xi
}

/*CountingSort ... used independently from RadixSort
 *
 *
 */
func CountingSort(xa []int) []int {

	max, min := xa[0], xa[0]
	for i := 0; i < len(xa); i++ {
		if xa[i] > max {
			max = xa[i]
		} else if xa[i] < min {
			min = xa[i]
		}
	}
	zeroOffset := 1
	absMin := int(math.Abs(float64(min)))
	lenC := max + absMin + zeroOffset

	var xb = make([]int, len(xa))
	var xc = make([]int, lenC)

	// step 1: add to C
	for i := 0; i < len(xa); i++ {
		val := xa[i]
		xc[val+absMin]++
	}
	// step 2: tally
	for i := 1; i < lenC; i++ {
		xc[i] += xc[i-1]
	}

	// step 3: order
	for i := 0; i < len(xa); i++ {
		val := xa[i]
		key := xc[val+absMin] - zeroOffset
		xb[key] = val
		xc[val+absMin]--
	}

	return xb
}

/*HeapSort ...
 *
 *
 */
func HeapSort(arr []int) []int {

	n := len(arr)

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapSortRecursion(arr, n, i)
	}

	// heap sort
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		// heapify root element to get highest element at root again
		heapSortRecursion(arr, i, 0)
	}

	return arr
}

/*heapSortRecursion... used by HeapSort
 *
 *
 */
func heapSortRecursion(arr []int, n, i int) {

	// find largest among root, left & right child
	largest := i
	l := 2*i + 1
	r := 2*i + 2
	if l < n && arr[l] > arr[largest] {
		largest = l
	}
	if r < n && arr[r] > arr[largest] {
		largest = r
	}

	// swap and confinue heapifying if root is not largest
	if largest != i {
		arr[i], arr[largest] = arr[largest], arr[i]
		heapSortRecursion(arr, n, largest)
	}
}

/*InsertionSort ...
 *
 *
 */
func InsertionSort(items []int) []int {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j--
		}
	}

	return items
}

/*MergeSort ...
 *
 *
 */
func MergeSort(xi []int) []int {

	// avoid stack overflow from recursion
	if len(xi) == 1 {
		return xi
	}

	// find mid point of slice and break it into two
	var (
		mid   = int(len(xi) / 2)
		left  = xi[0:mid]
		right = xi[mid:]
	)

	return mergeSortRecursion(MergeSort(left), MergeSort(right))
}

/*mergeSortRecursion ... used by MergeSort
 *
 *
 */
func mergeSortRecursion(left, right []int) (result []int) {

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

	for _, v := range left {
		result[i] = v
		i++
	}
	for _, v := range right {
		result[i] = v
		i++
	}

	return
}

/*QuickSort ...
 *
 *
 */
func QuickSort(xi []int) []int {

	quickSortRecursion(xi)

	return xi
}

/*quickSortRecursion ... used by QuickSort
 *
 *
 */
func quickSortRecursion(a []int) {

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
	quickSortRecursion(a[:left])
	quickSortRecursion(a[left+1:])
}

/*RadixSort ...
 *
 *
 */
func RadixSort(xi []int, lenA int) []int {

	max := xi[0]
	for i := 0; i < len(xi); i++ {
		if xi[i] > max {
			max = xi[i]
		}
	}

	for exp := 1; max/exp > 0; exp = exp * 10 {
		radixCountingSort(xi, lenA, exp)
	}

	return xi
}

/*radixSortMinMax ... used by radixCountingSort
 *
 *
 */
func radixSortMinMax(xa []int, lenA, exp int) (int, int) {
	var min, max int = (xa[0] / exp) % 10, (xa[0] / exp) % 10
	for i := 0; i < lenA; i++ {
		if (xa[i]/exp)%10 < min {
			min = (xa[i] / exp) % 10
		} else if (xa[i]/exp)%10 > max {
			max = (xa[i] / exp) % 10
		}
	}

	absMin := int(math.Abs(float64(min)))

	return absMin, max
}

/*radixCountingSort ... used by RadixSort
 *
 *
 */
func radixCountingSort(xa []int, lenA, exp int) {

	absMin, max := radixSortMinMax(xa, lenA, exp)
	zeroOffset := 1
	lenC := max + absMin + zeroOffset

	var xb = make([]int, lenA)
	var xc = make([]int, lenC)

	// Store count of occurrences in count[]
	for i := 0; i < lenA; i++ {
		expVal := (xa[i] / exp) % 10
		xc[expVal+absMin]++
	}

	for i := 1; i < lenC; i++ {
		xc[i] += xc[i-1]
	}

	for i := lenA - 1; i >= 0; i-- {
		expVal := (xa[i] / exp) % 10
		xb[xc[expVal+absMin]-1] = xa[i]
		xc[expVal+absMin]--
	}

	copy(xa, xb)
}

/*SelectionSort ...
 *
 *
 */
func SelectionSort(items []int) []int {

	var n = len(items)
	for i := 0; i < n; i++ {
		var minIdx = i
		for j := i; j < n; j++ {
			if items[j] < items[minIdx] {
				minIdx = j
			}
		}
		items[i], items[minIdx] = items[minIdx], items[i]
	}

	return items
}

/*ShellSort ...
 *
 *
 */
func ShellSort(items []int) []int {
	var (
		n    = len(items)
		gaps = []int{1}
		k    = 1
	)

	for {
		gap := shellSortElement(2, k) + 1
		if gap > n-1 {
			break
		}
		gaps = append([]int{gap}, gaps...)
		k++
	}

	for _, gap := range gaps {
		for i := gap; i < n; i += gap {
			j := i
			for j > 0 {
				if items[j-gap] > items[j] {
					items[j-gap], items[j] = items[j], items[j-gap]
				}
				j = j - gap
			}
		}
	}

	return items
}

/*shellSortElement ... used by ShellSort
 *
 *
 */
func shellSortElement(a, b int) int {
	e := 1
	for b > 0 {
		if b&1 != 0 {
			e *= a
		}
		b >>= 1
		a *= a
	}
	return e
}
