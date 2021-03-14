package anagram2

import (
    "sort"
)

func CompareIntSlices(s1 []int, s2 []int) bool {
    // Compares two []int slices and returns true if they're the same
    if len(s1) != len(s2) {
        return false
    }
    
    for i := range s1 {
        if s1[i] != s2[i] {
            return false
        }
    }
    return true
}

func AnagramTester(s1 string, s2 string) bool {
    // Declare empty []rune slices
    var s1s []rune
    var s2s []rune
    for _, v := range s1 {
        s1s = append(s1s, v)
    }
    for _, v := range s2 {
        s2s = append(s2s, v)
    }
    // Declare empty []int slices
    var i1s []int
    var i2s []int 

    for _, v := range s1s {
        i1s = append(i1s, int(v))
    }
    
    for _, v := range s2s {
        i2s = append(i2s, int(v))
    }

    // Sort the []int slices
    sort.Ints(i1s)
    sort.Ints(i2s)

    // Is it an anagram?
    isAnagram := CompareIntSlices(i1s, i2s)
    return isAnagram
}
