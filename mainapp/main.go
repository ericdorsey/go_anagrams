package main

// test with ("bberp", "pprbe")

import(
    "fmt"
    "fake.com/anagram/anagram1"
    "fake.com/anagram/anagram2"
)

func PrettyAnagramTester(f func(string, string) bool, s1 string, s2 string) {
    fmt.Printf("Comparing %v and %v\n", s1, s2)
    //r := anagram1.AnagramTester(s1, s2)
    r := f(s1, s2)
    fmt.Printf("%v\n", r)
}

func main() {
    // First strings to compare
    s1 := "bberp"
    s2 := "pprbe"
    // Second strings to compare
    s3 := "berp"
    s4 := "prbe"
    // Expecting comparing s1 to s2 to be false and s3 to s4 to be true

    fmt.Println("Run anagram1 AnagramTester()")
    f1 := anagram1.AnagramTester
    PrettyAnagramTester(f1, s1, s2) 
    PrettyAnagramTester(f1, s3, s4)
    fmt.Println()
    fmt.Println("Run anagram2 AnagramTester()")
    f2 := anagram2.AnagramTester
    PrettyAnagramTester(f2, s1, s2)
    PrettyAnagramTester(f2, s3, s4)
}
