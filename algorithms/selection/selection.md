/*
https://en.wikipedia.org/wiki/Selection_sort
//https://www.golangprograms.com/golang-program-for-implementation-of-selection-sort.html
This sorting algorithm begins by finding the smallest element in an array
and interchanging it with data at, for instance, array index [0].
Starting at index 0, we search for the smallest item in the list that exists between index 1 and the index of the last element.
When this element has been found, it is exchanged with the data found at index 0.
We simply repeat this process until the list becomes sorted.
*/
// Selection sort is an interesting one...
// Find the lowest number in array, switch in spot 0. Find next lower number in array, switch in spot 1.