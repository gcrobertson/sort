//https://www.programiz.com/dsa/heap-sort

//Heap Sort is a popular and efficient sorting algorithm in computer programming.

//The initial set of numbers that we want to sort is stored in an array

//e.g. [10, 3, 76, 34, 23, 32] and after sorting,
// we get a sorted array [3,10,23,32,34,76]

// Heap sort works by visualizing the elements of the array as a
// special kind of complete binary tree called heap.

/*
						What is a complete Binary Tree?
						Binary Tree
						A binary tree is a tree data structure in which each parent node can have at most two children
						Full Binary Tree
						A full Binary tree is a special type of binary tree in which every parent node has either two or no children.
						Complete binary tree
						A complete binary tree is just like a full binary tree, but with two major differences
						1.	Every level must be completely filled
						2.	All the leaf elements must lean towards the left.
						3.	The last leaf element might not have a right sibling
							i.e. a complete binary tree doesn’t have to be a full binary tree.
						How to create a complete binary tree from an unsorted list (array)?
						1. Select first element of the list to be the root node. (First level - 1 element)
						2. Put the second element as a left child of the root node and the third element as a right child. (Second level - 2 elements)
						3. Put next two elements as children of left node of second level. Again, put the next two elements as children of right node of second level (3rd level - 4 elements).
						4. Keep repeating till you reach the last element.
						[1, 12, 9, 5, 6, 10]
							1
					12				9
				5		6		10
			// Relationship between array indexes and tree elements
			// Complete binary tree has an interesting property that we can use to find the children and parents of any node.
							1 [0]
					12 [1]			9 [2]
				5 [3]	6 [4]	10
		If the index of any element in the array is i,
		the element in the index 2i+1 will become the left child
		and element in 2i+2 index will become the right child.
		Also, the parent of any element at index i is
		given by the lower bound of (i-1)/2.
		Left child of 1, 	i=0: 2*0+1 = [index 1], 12
		Right child of 1, 	i=0; 2*0+2 = [index 2], 9
		Left child of 12,	i=1; 2*1+1 = [index 3], 5
		Right child of 12,  i=1; 2*1+2 = [index 4], 6
		Parent of 9,		i=2; (2-1)/2 = 1/2 = 0
		Parent of 12,		i=1; (1-1)/2 = 0
		Parent of 6,		i=4; (4-1)/2 = 1.5 = [index 1], 12
		Understanding this mapping of array indexes to tree positions
		is critical to understanding how the Heap Data Structure works
		and how it is used to implement Heap Sort.
		What is Heap Data Structure ?
		Heap is a special tree-based data structure.
		A binary tree is said to follow a heap data structure if
		-	it is a complete binary tree
		-	All nodes in the tree follow the property that they are greater than their children
			i.e. the largest element is at the root and both its children and
			smaller than the root and so on. Such a heap is called a max-heap.
			If instead all nodes are smaller than their children,
			it is called a min-heap
		[MAX HEAP]
		 12
	  10    9
	 5  6  1
		 [MIN HEAP]
		 1
	   5    9
	10  6 12
	How to "heapify" a tree
	Run a function called heapify on all non-leaf elements of the
	heap.
	heapify(array)
		Root = array[0]
		Largest = largest( array[0], array [2*0 + 1]. array [2*0+2])
		if (Root != Largest)
			Swap(Root, Largest)
					  7
	Scenario 1:		2	4
	Root = 7
	Largest = (7, 1, 4)
	if 7 != 7...
					  2
	Scenario 2:		7	4
	Root = 2
	Largest = 7
	if 2 != 7
		Swap(2, 7)
					  7
					2	4
	Another scenario in which there are more than one levels.
			  2
		10			9
	5		6	 1
	Root = 2
	Largest = (2, 10, 9)
	Swap (2, 10)
			  10
		2			9
	5		6	 1
	Root = 2
	Largest = (2, 5, 6)
	Swap (2, 6)
			  10
		6			9
	5		2	 1
	Thus, to maintain the max-heap property in a tree where
	both sub-trees are max-heaps, we need to run heapify
	on the root element repeatedly until it is larger than
	its children or it becomes a leaf node.
	We can combine both these conditions in one heapify function as:
	void heapify(int arr[], int n, int i)
	{
		int largest = i;
		int l = 2*i + 1;
		int r = 2*i + 2;
		if (l < n && arr[l] > arr[largest])
			largest = l
		if (right < n && arr[r] > arr[largest])
			largest = r
		if (largest != i)
		{
			swap(arr[i], arr[largest]);
			// recursively heapify the affected sub-tree
			heapify(arr, n, largest)
		}
	}
*/

// Build max-heap

// To build a max-heap from any tree, we can thus start heapifying each sub-tree from the bottom up and end up with a max-heap after the function is applied on all the elements including the root element.

// In the case of complete tree, the first index of non-leaf node is given by n/2 - 1. All other nodes after that are leaf-nodes and thus don’t need to be heapified.
//
//

// This function works for both the base case and for a tree of any size.
// We can thus move the root element to the correct position to maintain
// the max-heap status for any tree size as long as the sub-trees are
// max-heaps.

// Build max-heap

// To build a max-heap from any tree, we can thus start heapifying each
// sub-tree from the bottom up and end up with a max-heap after the
// function is applied on all the elements including the root element.

// In the case of complete tree, the first index of non-leaf node
// is given by n/2 - 1.
// All other nodes after that are leaf-nodes
// and thus don’t need to be heapified.
//
//
// So, we can build a maximum heap as [halve length of list] -1.
// that points to the last `parent` node to
//
//		// build heapify
//		for (int i = n / 2 - 1; i >= 0; i--)
//			heapify(arr, n, i);
//
//
//	arr: 	1, 12, 9, 5, 6, 10
//	n: 		6
//	i:		6/2-1
//			2 > 0
//
//	call:	heapify([1, 12, 9, 5, 6, 10], 6, 2)
//
//	i=2:	switches the last element with last leaf
//			1, 12, 10, 5, 6, 9
//			heapify([1, 12, 10, 5, 6, 9], 6, 1)
//
//	i=1:	switches leftside parent node
//
//
//
//
//
//
//
//	i=0:
//