/* Considering the terms in the Fibonacci sequence whose values
   do not exceed four million, find the sum of the even-valued terms.
*/

package Fibonacci

type Is func(k int) bool

func Sum(max int, condition Is) (sum int) {
    i, j, k := 1, 1, 0
    for {
        if i >= max {
            break
        }
        if condition(i) {
            sum += i
        }
        k = j
        j = i
        i = k + j
    }
    return sum
}

func Even(k int) bool {
    return k % 2 == 0
}

func Odd(k int) bool {
    return k % 2 == 1
}

func All(_ int) bool {
    return true
}
