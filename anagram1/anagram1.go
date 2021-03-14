package anagram1

import (
)

func MapPopulator(s string) map[string]int {
    // Populates a map from a string; values is number of time each character appears
    m := make(map[string]int)
    for _, v := range s {
        strv := string(v)
        _, ok := m[strv]
        // Character not yet present
        if !ok {
            m[strv] = 1
        // Character already in the map
        } else if ok {
            m[strv] += 1
        }
    }
    return m
}

func CompareMaps(m1 map[string]int, m2 map[string]int) bool {
    // Compares two maps of one level deep (no nested keys / values) to see if they're the same
    // If they're not the same length they're not the same
    if len(m1) != len(m2) {
        return false
    }
    for k := range m1 {
        if m1[k] != m2[k] {
            return false
        }
    } 
    return true
}

func AnagramTester(s1 string, s2 string) bool {
    // Make some maps to hold occurences of characters
    s1map := MapPopulator(s1)
    s2map := MapPopulator(s2)
    isAnagram := CompareMaps(s1map, s2map)
    return isAnagram
}
