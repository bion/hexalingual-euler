
package main

import (
    "fmt"
    "gitub.com/professorq/trilingual-euler/euler_1/threefive"
    "os"
    "strconv"
)

func main{
    ceiling := strconv.Atoi(os.Args[1])
    sum := threefive.SumLess(ceiling)
    fmt.Print("%v", sum)
}
