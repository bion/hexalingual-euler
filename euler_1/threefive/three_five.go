/*
Find the sum of all n such that n % 3 == or n % 5 == 0 and n < 1000
*/

package ThreeFive

func SumLess(max int) (sum int) {
    for i := 0; max >= i*3; i++ {
        // Alwyas add three times i.
        sum += 3 * i
        fiver := 5 * i
        if fiver >= max {
            // Don't waste computations on fiver
            // if we're past the point where i * 5 fits
            // under the cap.
            continue
        }
        // Only add fiver to the accumulator if the same
        // number has not been added as a threever
        if fiver := i * 5; fiver % 3 != 0 {
            sum += fiver
        }
    }
    return sum
}
