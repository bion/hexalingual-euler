
package Factor

import (
    "testing"
)

func TestPrimeIsIdentity(t *testing.T) {
    prime := 7
    if result := GreatestPrime(prime); result != prime {
        t.Errorf("Expected %v, received %v", prime, result)
    }
}

func TestTenGPFisFive(t *testing.T) {
    in, expected := 10, 5
    if result := GreatestPrime(in); expected != result {
        t.Errorf("Expected %v, received %v", expected, result)
    }
}

func BenchmarkGivenProblem(b *testing.B) {
    for i := 0; i < b.N; i++ {
        GreatestPrime(600851475143)
    }
}
