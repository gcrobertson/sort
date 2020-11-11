#### Overview
This small CLI program takes the following input:
- range:    [INT]     data range of random integer array. 0 - range.
- size:     [INT]     data size of random integer array to be sorted.
- sorts:    [STRING]  sort method. available options: bubble, counting, heap, insertion, merge, quick, radix, selection, shell
- natural:  [BOOL]    natural numbers only if true, including 0 and limited to `range` argument.

#### CLI examples
> clear && go run main.go -sorts=bubble,counting,heap,insertion,merge,quick,radix,selection,shell -size=100000
> 
> clear && go run main.go -sorts=counting,radix -size=20 -range=99999 -natural=true
>

#### Sorts implemented
- Bubble sort
- Counting sort
- Heap sort
- Insertion sort
- Merge sort
- Quick sort
- Raxis sort
- Selection sort
- Shell sort
