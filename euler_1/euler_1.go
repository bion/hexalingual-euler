
package main

import (
    "fmt"
    "./threefive"
    "os"
    "strconv"
)
    // "github.com/professorq/trilingual-euler/euler_1/threefive"

func main() {
    ceiling, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    sum := ThreeFive.SumLess(ceiling)
    fmt.Println(sum)
}
