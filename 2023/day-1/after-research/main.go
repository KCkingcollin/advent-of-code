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
    // result, err := solveDay1Part1()
    // if err != nil {
    //     fmt.Printf("Error solving puzzle: %s", err)
    // }
    result, err := solveDay1Part2()
    if err != nil {
        fmt.Printf("Error solving puzzle: %s", err)
    }
    fmt.Printf("result: %v\n", result)
    end := time.Now()
    mainLoopTime := end.Sub(start)
    fmt.Printf("Main loop took %v to complete.\n", mainLoopTime)
}

func solveDay1Part1() (int, error) {
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
        firstDigit, _ := findFirstDigit(line)
        lastDigit, _ := findLastDigit(line)
        if lastDigit == -1 || firstDigit == -1 {
            continue
        }
        num, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
        if err != nil {
            return -1, err
        } 
        sum += num
        fmt.Println(num)
    }
    return sum, err
}

func solveDay1Part2() (int, error) {
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
        firstNamedRune1, firstNamedInt1 := findFirstNamedInt(line)
        firstDigitRune2, firstDigitInt2 := findFirstDigit(line)
        firstDigit, _ := minPos(firstNamedRune1, firstNamedInt1, firstDigitRune2, firstDigitInt2)
        // firstDigit = firstDigit
        lastNamedRune1, lastNamedInt1 := findLastNamedInt(line)
        lastDigitRune2, lastDigitInt2 := findLastDigit(line)
        lastDigit, _ := maxPos(lastNamedRune1, lastNamedInt1, lastDigitRune2, lastDigitInt2)
        // lastDigit = lastDigit
        // fmt.Println(string(firstDigit) + string(lastDigit))
        num, err := strconv.Atoi(string(firstDigit) + string(lastDigit))
        if err != nil {
            return -1, err
        } 
        sum += num
        // fmt.Printf("puzzle: %s num: %v\n", line, num)
    }
    return sum, err
}

func minPos(num1 rune, index1 int, num2 rune, index2 int) (rune, int) {
    switch {
    case index1 == -1:
        return num2, index2
    case index2 == -1:
        return num1, index1
    case index1 < index2:
        return num1, index1
    default:
        return num2, index2
    }
}

func maxPos(num1 rune, index1 int, num2 rune, index2 int) (rune, int) {
    switch {
    case index1 == -1:
        return num2, index2
    case index2 == -1:
        return num1, index1
    case index1 > index2:
        return num1, index1
    default:
        return num2, index2
    }
}

func isNamedInt(input string) rune {
    switch {
    case strings.HasPrefix(input, "one"):
        return '1'
    case strings.HasPrefix(input, "two"):
        return '2'
    case strings.HasPrefix(input, "three"):
        return '3'
    case strings.HasPrefix(input, "four"):
        return '4'
    case strings.HasPrefix(input, "five"):
        return '5'
    case strings.HasPrefix(input, "six"):
        return '6'
    case strings.HasPrefix(input, "seven"):
        return '7'
    case strings.HasPrefix(input, "eight"):
        return '8'
    case strings.HasPrefix(input, "nine"):
        return '9'
    default:
        return -1
    }
}

func findFirstNamedInt(input string) (rune, int) {
    for i := range input {
        endOfStr := min(len(input)-i, 5)
        if i+endOfStr <= len(input) {
            if namedInt := isNamedInt(input[i : i+endOfStr]); namedInt > 0 {
                return namedInt, i
            }
        }
    }
    // fmt.Println("no words F: ", input)
    return -1, -1
}

func findLastNamedInt(input string) (rune, int) {
    for i := len(input)-1; i >= 0; i-- {
        endOfStr := min(len(input)-i, 5)
        if i+endOfStr <= len(input) {
            if namedInt := isNamedInt(input[i : i+endOfStr]); namedInt > 0 {
                return namedInt, i
            }
        }
    }
    // fmt.Println("no words B: ", input)
    return -1, -1
}

func findFirstDigit(input string) (rune, int) {
    for i, r := range input {
        if len(input) <= 0 {
            return -1, -1
        }
        if unicode.IsDigit(r) {
            // fmt.Println(i)
            return r, i
        }
    }
    return -1, -1
}

func findLastDigit(input string) (rune, int) {
    runes := []rune(input)
    for i := len(runes) - 1; i >= 0; i-- {
        if len(input) <= 0 {
            return -1, -1
        }
        if unicode.IsDigit(runes[i]) {
            // fmt.Println(i)
            return (runes[i]), i
        }
    }
    return -1, -1
}
