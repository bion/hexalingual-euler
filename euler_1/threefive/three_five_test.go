package ThreeFive

import (
    "fmt"
    "testing"
)

func TestGivenAnswer(t *testing.T) {
    if SumLess(10) != 23 {
        const max, sum = 10, 23
        t.Errorf("SumLess(%v) = %v", max, sum)
    }
}

func BenchmarkThreeFiveOneK(b *testing.B) {
    const max = 1000
    for i := 0; i < b.N; i++ {
        sum := SumLess(max)
        fmt.Sprintf("SumLess(%v) = %v", max, sum)
    }
}

func BenchmarkThreeFiveTenK(b *testing.B) {
    const max = 10000
    for i := 0; i < b.N; i++ {
        sum := SumLess(max)
        fmt.Sprintf("SumLess(%v) = %v", max, sum)
    }
}

func BenchmarkThreeFiveMill(b *testing.B) {
    //        1,000,000
    const max = 1000000
    for i := 0; i < b.N; i++ {
        sum := SumLess(max)
        fmt.Sprintf("SumLess(%v) = %v", max, sum)
    }
}
