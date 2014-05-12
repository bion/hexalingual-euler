
package main

import (
    "os"
    "fmt"
    "./palindrome"
    "strconv"
)

func main() {
    input_digits, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    fmt.Println(palindrome.LargestWithLength(input_digits))
}
