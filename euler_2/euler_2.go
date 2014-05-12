
package main

import (
    "fmt"

    // "github.com/professorq/trilingual-euler/euler_2/fibonacci"
    "./fibonacci"
    "os"
    "strconv"
)

func main() {
    limit, err:= strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    result := Fibonacci.Sum(limit, Fibonacci.Even)
    fmt.Println(result)
}
