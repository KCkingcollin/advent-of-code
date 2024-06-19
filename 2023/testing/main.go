package main

import (
    "fmt"
    "strings"
)

// LastIndexReverse finds the last occurrence of substr in s, searching backwards.
func LastIndexReverse(s, substr string) int {
    // Use strings.LastIndex to find the last occurrence of substr in s
    return strings.LastIndex(s, substr)
}

func main() {
    s := "hello, world, hello, Go"
    substr := "hello"
    
    index := LastIndexReverse(s, substr)
    if index != -1 {
        fmt.Printf("Last occurrence of '%s' in '%s' is at index %d\n", substr, s, index)
    } else {
        fmt.Printf("'%s' not found in '%s'\n", substr, s)
    }
}

