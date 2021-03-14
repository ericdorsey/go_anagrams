package main

import (
    "fmt"
    "sort"
)

// Compares two []int slices and returns true if they're the same
func CompareIntSlices(s1 []int, s2 []int) bool {
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

func AnagramTester(s1 string, s2 string) {
    fmt.Printf("Given %v and %v\n", s1, s2)

    // Declare empty []rune slices
    var s1s []rune
    var s2s []rune
    for _, v := range s1 {
        s1s = append(s1s, v)
    }
    for _, v := range s2 {
        s2s = append(s2s, v)
    }
    fmt.Println("\nPrior rune values:")
    fmt.Println(s1s)
    fmt.Println(s2s)

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

    fmt.Println("\nRunes converted to ints; sorted:")
    fmt.Println(i1s)
    fmt.Println(i2s)
    
    // Is it an anagram?
    anagram := CompareIntSlices(i1s, i2s)
    //return a
    fmt.Printf("\nAnagram? %v\n", anagram)
}

func main() {
    AnagramTester("berp", "bpre") 
}
