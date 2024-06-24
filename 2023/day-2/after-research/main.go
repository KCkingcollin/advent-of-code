package main

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strconv"
    "strings"
    "time"
    "unicode"
)

func main() {
    start := time.Now()
    // result, err := solveDay2Part1()
    result, err := solveDay2Part2()
    // the problem was so similar that i don't think it makes sense to change too much
    if err != nil {
        fmt.Printf("Error solving puzzle: %s", err)
    }
    fmt.Printf("result: %v\n", result)
    end := time.Now()
    mainLoopTime := end.Sub(start)
    fmt.Printf("Main loop took %v to complete.\n", mainLoopTime)
}

func solveDay2Part2() (int, error) {
    puzzle, err := os.Open("puzzle")
    if err != nil {
        fmt.Println("failed to open puzzle")
        return -1, err
    }
    defer puzzle.Close()
    scanner := bufio.NewScanner(puzzle)
    var sum int
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ": ")
        calc := checkStringPart2(line[1])
        sum += calc
    }
    return sum, err
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
    for scanner.Scan() {
        line := strings.Split(scanner.Text(), ": ")
        gameID := getNum(strings.Split(line[0], "Game ")[1])        
        isPossible := checkStringPart1(line[1])
        if isPossible {
            sum += gameID
        }
        // fmt.Println(gameID)
    }
    return sum, err
}

func getNum(input string) (int) {
    var numStr string
    var err error
    var num int
    for _, r := range input {
        if unicode.IsDigit(r) {
            numStr += string(r)
            num, err = strconv.Atoi(numStr)
            if err != nil {
                log.Fatalf("error converting string to number: %v\n", err)
            }
        }
    }
    return num
}

func checkStringPart1(input string) (bool) {
    isMax := map[string]int {
        "red": 12,
        "green": 13,
        "blue": 14,
    }
    var split []string
    splitIn := strings.Split(input, "; ")
    for _, block := range splitIn {
        split = append(split, strings.Split(block, ", ")...)
    }
    for i, r := range split {
        numWord := strings.Split(r, " ")
        // removed exceedsMax function as it was not needed
        if isMax[string(numWord[1])] < getNum(numWord[0]) {
            return false
            // if exceedsMax is true, then return a false bool
        } else if len(split)-1 <= i {
            return true
            // if exceedsMax is false, then return a true bool
        }
    }
    return false
    // if exceedsMax is never false, then return a true bool
}

func checkStringPart2(input string) (int) {
    maxValues := map[string]int {
        "red": 0,
        "green": 0,
        "blue": 0,
    }
    var split []string
    splitIn := strings.Split(input, "; ")
    for _, block := range splitIn {
        split = append(split, strings.Split(block, ", ")...)
    }
    for i, r := range split {
        numWord := strings.Split(r, " ")
        // used the max function instead of the if statement for speed and simplicity
        maxValues[numWord[1]] = max(maxValues[numWord[1]], getNum(numWord[0]))
        if len(split)-1 <= i {
            return maxValues["red"] * maxValues["green"] * maxValues["blue"]
            // return the power set 
        }
    }
    return -1
    // if exceedsMax is never false, then return a true bool
}

