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
        if isPossible == true && num >= 0 {
        sum += num
        }
        fmt.Println(num)
    }
    return sum, err
}

func findStingBackwards(input string) (bool, int) {
    // find := strings.HasPrefix
    var game bool
    var IDs int
    for i := len(input)-1; i >= 0; i-- {
        endOfStr := min(len(input)-i, 5)
        numLocation := min(len(input)-i, 3)
        wordBlock := input[i : i+endOfStr]
        numBlock := input[i : i+numLocation]
        if i+endOfStr <= len(input) {
            if game, IDs = findDigit(numBlock,  wordBlock); game == true {
                return game, IDs
            }
            // switch {
            // case find(wordBlock, "red"):
            //     color, _ := findDigit(numBlock, wordBlock)
            //     if !color {
            //         return false, -1
            //     }
            // case find(wordBlock, "green"):
            //     color, _ := findDigit(numBlock, wordBlock)
            //     if !color {
            //         return false, -1
            //     }
            // case find(wordBlock, "blue"):
            //     color, _ := findDigit(numBlock, wordBlock)
            //     if !color {
            //         return false, -1
            //     }
            // case find(wordBlock, ":"):
            //     game, IDs = findDigit(numBlock, wordBlock)
            // }
        }
    }
    // fmt.Println("no words B: ", input)
    return false, -1
}

func findDigit(numInput string, word string) (bool, int) {
    find := strings.HasPrefix
    var numStr string
    var num int
    var err error
    fmt.Println(numInput)
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
        if num > 12 {
            return false, -1
        } else {
            return true, num
        }
    case find(word, "green"):
        if num > 13 {
            return false, -1
        } else {
            return true, num
        }
    case find(word, "blue"):
        if num > 14 {
            return false, -1
        } else {
            return true, num
        }
    case find(word, ":"):
        return true, num
    default:
        return false, -1
    }
}

