package main

import (
	"fmt"
)

func main() {
    maxValues := map[string]map[bool]func(v any) {
        "red": {
            true: func(v any) {fmt.Println(v)},
            false: func(v any) {fmt.Println(v, "fuck")},
        },
        "green": {
            true: func(v any) {fmt.Println(v)},
            false: func(v any) {fmt.Println(v, "fuck")},
        },
        "blue": {
            true: func(v any) {fmt.Println(v)},
            false: func(v any) {fmt.Println(v, "fuck")},
        },
    }
    maxValues["red"][true]("fuck ya")
}
