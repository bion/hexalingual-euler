
package main

import (
    "os"
    "./factor"
    "fmt"
    "strconv"
)

func main() {
    product, err := strconv.Atoi(os.Args[1])
    if err != nil {
        panic(err)
    }
    fmt.Println(Factor.GreatestPrime(product))
}
