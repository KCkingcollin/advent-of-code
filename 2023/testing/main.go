package main

import (
    "fmt"
)

func combineRunes(runes []rune) string {
    var combined string
    for _, r := range runes {
        combined += string(r)
    }
    return combined
}

func main() {
    runes := []rune{'H', 'e', 'l', 'l', 'o', ',', ' ', '世', '界'}
    combined := combineRunes(runes)

    // Print the combined string
    fmt.Printf("Combined: %s\n", combined)
}

