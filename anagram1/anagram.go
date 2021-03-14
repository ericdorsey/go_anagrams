package main

import (
    "fmt"
)

func AnagramTester(s1 string, s2 string) bool {
    if len(s1) != len(s2) {
        return false
    }
    is_anagram := true
    // Loop through all the letters of s2
    for _, v1 := range s2 {
        this_letter := true
        c1 := string(v1)
        // Compare them against the letters of s1
        for _, v2 := range s1 {
            c2 := string(v2)
            fmt.Printf("Checking %v against %v\n", c1, c2)
            this_letter = false
            if c1 == c2 {
                this_letter = true
                fmt.Printf("Found match -- this_letter set to true\n")
                // Stop checking this letter, it's a match
                break
            } 
        } 
        if this_letter == false {
            is_anagram = false
            return false
        }
    }
    return is_anagram
}

func main() {
    result := AnagramTester("berp", "prbe")
    fmt.Printf("%v\n", result)
}
