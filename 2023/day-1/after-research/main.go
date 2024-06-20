package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"time"
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
    var firstDigit int
    var lastDigit int
    var line string
    for scanner.Scan() {
        line = scanner.Text()
        firstDigit, err = findFirst(line)
        if err != nil {
            fmt.Printf("%s\n", err)
        }
        var firstDigit string = strconv.Itoa(firstDigit)
        lastDigit, err = findLast(line)
        if err != nil {
            fmt.Printf("%s\n", err)
        }
        var lastDigit string = strconv.Itoa(lastDigit)
        if err == nil {
            result, err := strconv.Atoi(firstDigit + lastDigit)
            if err != nil {
                return 0, err
            }
            sum += result
        }
        fmt.Printf("%v%v\n", firstDigit, lastDigit)
    }
    return sum, err
}

func findFirst(input string) (int, error) {
    for i := 0; i < len(input); i++ {
        digit, err := strconv.Atoi(string(input[i]))
        if err == nil {
            return digit, nil
        } 
    }
    return 0, fmt.Errorf("no number found\n")
}

func findLast(input string) (int, error) {
    for i := len(input)-1; i < len(input); i-- {
        digit, err := strconv.Atoi(string(input[i]))
        if err == nil {
            return digit, nil
        } 
    }
    return 0, fmt.Errorf("no number found\n")
}
