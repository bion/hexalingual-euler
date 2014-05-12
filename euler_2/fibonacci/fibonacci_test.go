
package Fibonacci

import (
    "testing"
)

func TestBasicFibonacciSum(t *testing.T) {
    if result := Sum(6, All); result != 11 {
        t.Error(result)
    }
}

func TestBaseCase(t *testing.T) {
    if result :=Sum(2, Odd); result != 1 {
        t.Error(result)
    }
}

func TestOdds(t *testing.T) {
    if result :=Sum(10, Odd); result != 9 {
        t.Error(result)
    }
}

func BenchmarkGivenProblem(b *testing.B) {
    for i := 0; i < b.N; i++ {
        Sum(4000000, Odd)
    }
}
