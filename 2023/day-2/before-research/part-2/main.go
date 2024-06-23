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
    result, err := solveDay2Part2()
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
        calc := checkString(line[1])
        sum += calc
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

func checkString(input string) (int) {
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
        if maxValues[numWord[1]] < getNum(numWord[0]) {
            maxValues[numWord[1]] = getNum(numWord[0])
        }
        if len(split)-1 <= i {
            calc := maxValues["red"] * maxValues["green"] * maxValues["blue"]
            return calc
            // if exceedsMax is false, then return a true bool
        }
    }
    return -1
    // if exceedsMax is never false, then return a true bool
}

