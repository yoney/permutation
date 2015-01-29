// Package permutation provides functions to change the order of the elements
// in the sequence according to
//     1. lexicographically greater permutation (NextPermutation)
//     2. previous lexicographically-ordered permutation (PrevPermutation)
package permutation

import "sort"

// NextPermutation rearranges the elements in the sequence into the next
// lexicographically greater permutation.
//
// Algorithm:
//     1. Find the largest index k such that a[k] < a[k + 1]. If no such index
//        exists, the permutation is the last permutation.
//     2. Find the largest index l greater than k such that a[k] < a[l].
//     3. Swap the value of a[k] with that of a[l].
//     4. Reverse the sequence from a[k + 1] up to and including the final element
//
// Complexity:
//     Linear, at most seq.Len() / 2 swaps
//
// It returns true if the function could rearrange the sequence as a
// lexicographicaly greater permutation.
func NextPermutation(seq sort.Interface) bool {
    len := seq.Len()
    for k := len - 2; k >= 0; k-- {
        if seq.Less(k, k + 1) {
            for l := len - 1; l > k; l-- {
                if seq.Less(k, l) {
                    seq.Swap(k, l)
                    reverse(seq, k + 1)
                    return true
                }
            }
        }
    }

    return false
}

// PrevPermutation rearranges the elements in the sequence into the previous
// lexicographically-ordered permutation
//
// Algorithm:
//     1. Find the largest index k such that a[k] > a[k + 1]. If no such index
//        exists, the permutation is the last permutation.
//     2. Find the largest index l greater than k such that a[k] > a[l].
//     3. Swap the value of a[k] with that of a[l].
//     4. Reverse the sequence from a[k + 1] up to and including the final element
//
// Complexity:
//     Linear, at most seq.Len() / 2 swaps
//
// It returns true if the function could rearrange the sequence in the previous
// lexicographically-ordered permutation.
func PrevPermutation(seq sort.Interface) bool {
    len := seq.Len()
    for k:= len - 2; k >= 0; k-- {
        if seq.Less(k + 1, k) {
            for l := len - 1; l > k; l-- {
                if seq.Less(l, k) {
                    seq.Swap(k, l)
                    reverse(seq, k + 1)
                    return true
                }
            }
        }
    }

    return false
}

func reverse(seq sort.Interface, firstIndex int) {
    lastIndex := seq.Len() - 1

    numSwap := (lastIndex - firstIndex + 1) / 2;
    for i := 0; i < numSwap; i++ {
        seq.Swap(firstIndex + i, lastIndex - i)
    }
}

