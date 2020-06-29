
Algorithm: 

//Bubble ...
func Bubble(xi []int) []int {

	for i := 0; i < len(xi); i++ {
		for j := 0; j < len(xi)-i-1; j++ {
			// fmt.Printf("outer loop index [%v] inner loop comparison: %v > %v?\n", i, xi[j], xi[j+1])
			if xi[j] > xi[j+1] {
				xi[j], xi[j+1] = xi[j+1], xi[j]
			}
		}
	}

	return xi
}





/*	first outer loop iteration:
 *	53, 12, 1024, 35
 *	53 > 12 			?
 *	12, 53, 1024, 35
 *	53 > 1024			?
 *	1024 > 35			?
 *	12, 53, 35, [1024]		note: the last element in array is now ordered.
 *
 *	second outer loop iteration:
 *	12 > 53				?
 *	53 > 35				?
 *	12, 35, [53], [1024]
 *
 *	third outer loop iteration:
 *	12 > 35?
 *	12, [35], [53], [1024]
 *
 */

// Presort :[53 12 1024 35]
// outer loop index [0] inner loop comparison: 53 > 12?
// outer loop index [0] inner loop comparison: 53 > 1024?
// outer loop index [0] inner loop comparison: 1024 > 35?
// outer loop index [1] inner loop comparison: 12 > 53?
// outer loop index [1] inner loop comparison: 53 > 35?
// outer loop index [2] inner loop comparison: 12 > 35?
// Sorted  :[12 35 53 1024]




/*
	Presort :		[53 12 35 101 11]
					oLoop: 0:
					[12 53 35 101 11]	53 > 12
					[12 35 53 101 11]	53 > 35
					[12 35 53 101 11]
					in simulationEnabled = true, for fulfillment, you are only getting the main service to work. It should change to in progress. [1]
					the add on service status stays in 0.
					that is why she cannot make use of her simulations.
					if she turns off the simulation, the order gets completed.  order gets completed is expected behavior.
					she needs to use fulfillments she needs the simulation working for negative tests.
*/

/*
Sinking Sort
Bubble	Sort
https://en.wikipedia.org/wiki/Bubble_sort
Bubble sort, sometimes referred to as sinking sort,
is a simple sorting algorithm that repeatedly steps through the list,
compares adjacent elements and swaps them if they are in the wrong order.
The pass through the list is repeated until the list is sorted.
The algorithm, which is a comparison sort, is named for the way smaller or larger elements "bubble" to the top of the list.
Although the algorithm is simple, it is too slow and impractical for most problems even when compared to insertion sort.
Bubble sort can be practical if the input is in mostly sorted order with some out-of-order elements nearly in position.
Performance
Bubble sort has a worst-case and average complexity of О(n2), where n is the number of items being sorted.
Most practical sorting algorithms have substantially better worst-case or average complexity, often O(n log n).
Even other О(n2) sorting algorithms, such as insertion sort, generally run faster than bubble sort, and are no more complex.
Therefore, bubble sort is not a practical sorting algorithm.
The only significant advantage that bubble sort has over most other algorithms, even quicksort,
but not insertion sort, is that the ability to detect that the list is sorted efficiently is built into the algorithm.
When the list is already sorted (best-case), the complexity of bubble sort is only O(n).
By contrast, most other algorithms, even those with better average-case complexity,
perform their entire sorting process on the set and thus are more complex.
However, not only does insertion sort share this advantage,
but it also performs better on a list that is substantially sorted (having a small number of inversions).
Bubble sort should be avoided in the case of large collections. It will not be efficient in the case of a reverse-ordered collection.
*/
