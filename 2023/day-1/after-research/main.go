package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
	"unicode"
)

func main() {
    start := time.Now()
    result, err := solveDayOneOfOne()
    if err != nil {
        fmt.Printf("Error solving puzzle: %s", err)
    }
    fmt.Printf("result: %v\n", result)
    end := time.Now()
    mainLoopTime := end.Sub(start)
    fmt.Printf("Main loop took %v to complete.\n", mainLoopTime)
}

func solveDayOneOfOne() (int, error) {
    puzzle, err := os.Open("puzzle")
    if err != nil {
        fmt.Println("failed to open puzzle")
        return 0, err
    }
    defer puzzle.Close()
    scanner := bufio.NewScanner(puzzle)
    var sum int
    var line string
    for scanner.Scan() {
        line = scanner.Text()
        num, err := strconv.Atoi(string(findFirst(line)) + string(findLast(line)))
        if err != nil {
            return 0, err
        }
        sum += num
        fmt.Println(num)
    }
    return sum, err
}

func findFirst(input string) rune {
    for _, r := range input {
        if unicode.IsDigit(r) {
            return r
        }
    }
    return 0
}

func findLast(input string) rune {
    runes := []rune(input)
    for i := len(runes) - 1; i >= 0; i-- {
        if unicode.IsDigit(runes[i]) {
            return (runes[i])
        }
    }
    return 0
}
