### notes from: https://brilliant.org/wiki/radix-sort/#:~:text=Radix%20sort%20is%20an%20integer,sort%20an%20array%20of%20numbers.

## Raxis sort:

    - uses counting sort as a subroutine to sort an array of numbers
    - radix sort is not comparison based and can perform in linear time
    - comparison sorts are bounded by Ω (n log n) time


## Radix sort:
    -   takes list of `n` integers with base `b`
    -   `b` the base is also known as the radix
    -   each number has at most `d` digits: d = log b (k) + 1
    -   `k` is the largest number in the list

The algorithm runs in linear time when `b` and `n` are of the same size magnitude.

## How Radix Sort works:

Radix sort works by sorting each digit from least significant digit to most significant digit
- In base 10 [decimal system], radix sort would sort the digits in 1's, then 10's, then 100's, etc.

Radix uses `counting sort` as a subroutine to sort the digits in each place.

For a 3-digit number in base 10, counting sort will be called to sort:
- the 1's place, 
- the 10's place
- and finally the 100's place


## Counting Sort Subroutine:

Counting Sort uses 3 lists:
- The input list:                   A[0,1,..n]
- the output list:                  B[0,1,..n]
- list serving as temporary memory: C[0,1,..k]
    - n = length of list
    - k = range of list

1. For each element A[i], it goes to index of C for the same value as A[i] and increments the value C[A[i]] by one.
- If A has seven 0's:   C[0] = 7
- If A has two 4's:     C[4] = 2

C keeps track of how many elements in A there are that have the same value of a particular index in C.
- The indexes in C correspond to the values in A.
- The values in C correspond to the total number of times that a value in A appears in A.

## Radix Sort

Radix sort is a `stable sort`.

`Stable Sort`: it preserves the relative order of elements that have the same key value.

## Example

input           =   [56, 43, 51, 58]
sort 1s place   =   [51, 43, 56, 58]    // radix maintains relative order of 1 < 6 < 8
sort 10s place  =   [43, 51, 56, 58]

## More on Counting Sort

Counting sort can only sort one place value of a given base.
- A counting sort base-10 can only sort digits 0-9
- To sort two-digit numbers, counting sort would need to operate in base-100.

Radix sort is more powerful because it can sort multi-digit numbers without having to search over a wider range of keys [a larger base].

## Final example of Radix Sort

329     sort 1..    720     sort 10..   720     sort 100..      329
457                 355                 329                     355
657                 436                 436                     436
839                 457                 839                     457
436                 657                 355                     657
720                 329                 457                     720
355                 839                 657                     839

# if ever the 10s match, go to the 1s and determine who goes first...
# if ever the 10s match, go to the lower key first since it is ordered...