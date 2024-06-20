package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
    start := time.Now()
    puzzle, err := os.Open("puzzle")
    if err != nil {
        fmt.Println("Error opening file:", err)
    }
    defer puzzle.Close()
    scanner := bufio.NewScanner(puzzle)
    if err := scanner.Err(); err != nil {
        fmt.Printf("Error reading input: %v\n", err)
    }
    var sum int
    var line string

    for scanner.Scan() {
        line = scanner.Text()
        num1 := words2numF(line)
        num2 := words2numB(line)
        finalNum := getFirstAndLast(num1 + num2)
        // finalNum := getFirstAndLast(line)
        result, err := strconv.Atoi(finalNum)
        if err != nil {
            fmt.Printf("string convert failed: %s\n", err)
        }
        // fmt.Printf("%v\n", result)
        sum += result
    }

    fmt.Println("Result:", sum)
    end := time.Now()
    mainLoopTime := end.Sub(start)
    fmt.Printf("Main loop took %v to complete.\n", mainLoopTime)
}

func words2numF (input string) (string) {
    wordToNum :=  map[string]string{
        "one":   "1",
        "two":   "2",
        "three": "3",
        "four":  "4",
        "five":  "5",
        "six":   "6",
        "seven": "7",
        "eight": "8",
        "nine":  "9",
    }    
    regex := regexp.MustCompile(`(one|two|three|four|five|six|seven|eight|nine)`)
    result := regex.ReplaceAllStringFunc(input, func(match string) string {
        return wordToNum[strings.ToLower(match)]
    })
    return result
}

func reverseString (input string) string {
    runes := []rune(input)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}

func words2numB (input string) string {
    input = reverseString(input)
    wordToNum :=  map[string]string{
        "eno":   "1",
        "owt":   "2",
        "eerht": "3",
        "ruof":  "4",
        "evif":  "5",
        "xis":   "6",
        "neves": "7",
        "thgie": "8",
        "enin":  "9",
    }    
    regex := regexp.MustCompile(`(eno|owt|eerht|ruof|evif|xis|neves|thgie|enin)`)
    result := regex.ReplaceAllStringFunc(input, func(match string) string {
        return wordToNum[strings.ToLower(match)]
    })
    result = reverseString(result)
    return result
}

func getFirstAndLast(input string) string {
    var num string
    var char1 string
    var char2 string

    var sb strings.Builder
    for _, r := range input {
        if unicode.IsDigit(r) {
            sb.WriteRune(r)
        }
    }

    num = sb.String()
    char1 = string(num[0])
    char2 = string(num[len(num)-1])
    result := char1 + char2

    return result
}

