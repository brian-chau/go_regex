package main

import (
    "fmt"
    "regexp"
    "os"
)

func main() {
    input       := os.Args[1]
    pattern     := regexp.MustCompile(os.Args[2])
    replacement := os.Args[3]
    res := pattern.ReplaceAllString(input, replacement)
    fmt.Println(res)
}