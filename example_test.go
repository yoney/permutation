package permutation_test

import (
    "fmt"
    "sort"
    "github.com/yoney/permutation"
)

func ExampleNextPermutation() {
    arr := []int{1, 2, 3}

    fmt.Println(arr)
    for permutation.NextPermutation(sort.IntSlice(arr)) {
        fmt.Println(arr)
    }
    // Output:
    // [1 2 3]
    // [1 3 2]
    // [2 1 3]
    // [2 3 1]
    // [3 1 2]
    // [3 2 1]
}

func ExamplePrevPermutation() {
    arr := []int{3, 2, 1}

    fmt.Println(arr)
    for permutation.PrevPermutation(sort.IntSlice(arr)) {
        fmt.Println(arr)
    }
    // Output:
    // [3 2 1]
    // [3 1 2]
    // [2 3 1]
    // [2 1 3]
    // [1 3 2]
    // [1 2 3]
}

func Example_int() {
    // import "sort"

    arr := []int{1, 2, 3}

    fmt.Println(arr)
    for permutation.NextPermutation(sort.IntSlice(arr)) {
        fmt.Println(arr)
    }
    // Output:
    // [1 2 3]
    // [1 3 2]
    // [2 1 3]
    // [2 3 1]
    // [3 1 2]
    // [3 2 1]
}

type runeSlice []rune

func (s runeSlice) Less(i, j int) bool {
    return s[i] < s[j]
}

func (s runeSlice) Len() int {
    return len(s)
}

func (s runeSlice) Swap(i, j int) {
    s[i], s[j] = s[j], s[i]
}

func Example_string() {
    // type runeSlice []rune
    //
    // func (s runeSlice) Less(i, j int) bool {
    //     return s[i] < s[j]
    // }
    //
    // func (s runeSlice) Len() int {
    //     return len(s)
    // }
    //
    // func (s runeSlice) Swap(i, j int) {
    //     s[i], s[j] = s[j], s[i]
    // }

    str := "abcd"
    runeSli := runeSlice(str)

    fmt.Println(string(runeSli))
    for permutation.NextPermutation(runeSli) {
        fmt.Println(string(runeSli))
    }
    // Output:
    // abcd
    // abdc
    // acbd
    // acdb
    // adbc
    // adcb
    // bacd
    // badc
    // bcad
    // bcda
    // bdac
    // bdca
    // cabd
    // cadb
    // cbad
    // cbda
    // cdab
    // cdba
    // dabc
    // dacb
    // dbac
    // dbca
    // dcab
    // dcba
}

