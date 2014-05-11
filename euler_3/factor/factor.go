
package Factor

import "math"

func GreatestPrime(product int)  int {
    for i := 2; int(math.Pow(float64(i), 2)) < product; i++ {
        if product % i == 0 {
           return product / i
        }
    }
    return product
}
