package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
	"unicode"
)

func main() {
    start := time.Now()
    result, err := solveDay2Part1()
    if err != nil {
        fmt.Printf("Error solving puzzle: %s", err)
    }
    fmt.Printf("result: %v\n", result)
    end := time.Now()
    mainLoopTime := end.Sub(start)
    fmt.Printf("Main loop took %v to complete.\n", mainLoopTime)
}

func solveDay2Part1() (int, error) {
    puzzle, err := os.Open("puzzle")
    if err != nil {
        fmt.Println("failed to open puzzle")
        return -1, err
    }
    defer puzzle.Close()
    scanner := bufio.NewScanner(puzzle)
    var sum int
    var line string
    for scanner.Scan() {
        line = scanner.Text()
        isPossible, num := findStingBackwards(line)
        if isPossible && num > 0 {
        sum += num
        }
        fmt.Println(num)
    }
    return sum, err
}

func findStingBackwards(input string) (bool, int) {
    find := strings.HasPrefix
    var good bool
    var ID int
    for i := len(input)-1; i >= 0; i-- {
        endOfStr := min(len(input)-i, 5)
        numLocation := min(len(input)-i, 3)
        wordBlock := input[i : i+endOfStr]
        numBlock := input[i : i+numLocation]
        if len(input)-5 >= i {
            numBlock = input[i-3 : i-3+numLocation]
        }
        if i+endOfStr <= len(input) {
            good, ID = findDigit(numBlock, wordBlock)
            if find(wordBlock, "red") || find(wordBlock, "green") || find(wordBlock, "blue") {
                if !good {
                    return false, -1
                }
            } else if find(wordBlock, ":") {
                return good, ID
            }
        }
    }
    return false, -1
}

func findDigit(numInput string, word string) (bool, int) {
    find := strings.HasPrefix
    var numStr string
    var num int
    var err error
    // fmt.Println(numInput)
    for _, r := range numInput {
        if unicode.IsDigit(r) {
            numStr += string(r)
        }
    }
    if numStr != "" {
        num, err = strconv.Atoi(numStr)
        if err != nil {
            fmt.Printf("error converting string to number: %v\n", err)
            return false, -1
        }
    }
    switch {
    case find(word, "red"):
        if 0 > num || num > 12 {
            return false, -1
        } else {
            return true, -1
        }
    case find(word, "green"):
        if 0 > num || num > 13 {
            return false, -1
        } else {
            return true, -1
        }
    case find(word, "blue"):
        if 0 > num || num > 14 {
            return false, -1
        } else {
            return true, -1
        }
    case find(word, ":"):
        return true, num
    default:
        return false, -1
    }
}

