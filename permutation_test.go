package permutation

import (
    "testing"
    "sort"
    "reflect"
)

func TestNextPermutation(t *testing.T) {
    arr := []int{0, 1, 2, 5, 3, 3, 0}
    perm := [][]int {
        {0, 1, 3, 0, 2, 3, 5},
        {0, 1, 3, 0, 2, 5, 3},
        {0, 1, 3, 0, 3, 2, 5},
        {0, 1, 3, 0, 3, 5, 2},
        {0, 1, 3, 0, 5, 2, 3},
        {0, 1, 3, 0, 5, 3, 2},
        {0, 1, 3, 2, 0, 3, 5},
        {0, 1, 3, 2, 0, 5, 3},
        {0, 1, 3, 2, 3, 0, 5},
        {0, 1, 3, 2, 3, 5, 0},
    }

    for i, p := range perm {
        r1 := NextPermutation(sort.IntSlice(arr))
        r2 := reflect.DeepEqual(arr, p)

        if !r1 || !r2 {
            t.Error("Failed at: ", i, r1, r2, arr, p)
        }
    }
}

func TestPrevPermutation(t *testing.T) {
    arr := []int{0, 1, 3, 2, 3, 5, 0}
    perm := [][]int {
        {0, 1, 3, 2, 3, 0, 5},
        {0, 1, 3, 2, 0, 5, 3},
        {0, 1, 3, 2, 0, 3, 5},
        {0, 1, 3, 0, 5, 3, 2},
        {0, 1, 3, 0, 5, 2, 3},
        {0, 1, 3, 0, 3, 5, 2},
        {0, 1, 3, 0, 3, 2, 5},
        {0, 1, 3, 0, 2, 5, 3},
        {0, 1, 3, 0, 2, 3, 5},
        {0, 1, 2, 5, 3, 3, 0},
    }


    for i, p := range perm {
        r1 := PrevPermutation(sort.IntSlice(arr))
        r2 := reflect.DeepEqual(arr, p)

        if !r1 || !r2 {
            t.Error("Failed at: ", i, r1, r2, arr, p)
        }
    }
}

