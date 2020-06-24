/**
 *	https://en.wikipedia.org/wiki/Quicksort
 *
 *	Quicksort is an efficient sorting algorithm,
 *	serving as a systematic method for placing the
 *  elements of a random access file or an array in order.
 *
 *	Developed by British computer scientist Tony Hoare in 1959
 *  and published in 1961, it is still a commonly used
 *  algorithm for sorting.
 *
 *	Inventor: 			Tony Hoare
 *	Worst complexity: 	n^2
 *	Average complexity: n*log(n)
 *	Best complexity: 	n*log(n)
 *	Method: 			Partitioning
 *	Stable: 			No
 *	Class: 				Comparison sort
 *
 *	"""When implemented well, it can be about two or three times faster
 *	   than its main competitors, merge sort and heapsort.""" - Is this true?
 *
 *	Quicksort is a divide and conquer algorithm.
 *	Quicksort first divides a large array into two smaller sub-arrays:
 *		the low elements and
 *		the high elements.
 *	Quicksort can then recursively sort the sub-arrays.
 *
 *	The steps are:
 *
 *	1.	Pick an element, called a pivot, from the array.
 *
 *  2. 	Partitioning: reorder the array so that:
 *		- all elements with values less than the pivot come before the pivot, while
 *		- all elements with values greater than the pivot come after it (equal values can go either way).
 *		After this partitioning, the pivot is in its final position.
 *		This is called the partition operation.
 *
 *  3.	Recursively apply the above steps to the sub-array of elements with smaller values and
 *		separately to the sub-array of elements with greater values.
 *
 *	The base case of the recursion is arrays of size zero or one,
 *	which are in order by definition, so they never need to be sorted.
 *
 *	The pivot selection and partitioning steps can be done in several different ways;
 *	the choice of specific implementation schemes greatly affects the algorithm's performance.
 */

//QuickSortingAlgorithm ...
var QuickSortingAlgorithm = SortingAlgorithm{
	Name:           "Quick Sort",
	Author:         "Tony Hoare",
	Year:           1959,
	AlgorithmRun:   QuickSort,
	AlgorithmDebug: QuickSortDebug,
	AlgorithmCount: QuickSort,
	AlgorithmType:  "Divide and Conquer",
	ComparisonSort: true,
	Alias:          []string{"Partition-exchange sort"},
	Stable:         false,
}