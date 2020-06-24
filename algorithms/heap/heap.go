package algorithms

//HeapSort ...
func HeapSort(arr []int) []int {

	n := len(arr)

	// build max heap
	for i := n/2 - 1; i >= 0; i-- {
		heapify(arr, n, i)
	}

	// heap sort
	for i := n - 1; i >= 0; i-- {
		arr[0], arr[i] = arr[i], arr[0]

		// heapify root element to get highest element
		// at root again
		heapify(arr, i, 0)
	}

	return arr
}

// heapify...
func heapify(arr []int, n, i int) {

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
		heapify(arr, n, largest)
	}
}
