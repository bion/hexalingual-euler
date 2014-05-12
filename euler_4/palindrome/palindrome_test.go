
package palindrome

import (
    "testing"
)

func TestPosFactorsToLength(t *testing.T) {
    const input, digits, output = 81, 2, true
    result := FactorsToLength(input, digits)
    if result != output {
        t.Errorf("Expected %v, Received %v",
                 output, result)
    }
}

func TestNegFactorsToLength(t *testing.T) {
    const input, digits, output = 82, 1, false
    result := FactorsToLength(input, digits)
    if result != output {
        t.Errorf("Expected %v, Received %v",
                 output, result)
    }
}

func TestReverse(t *testing.T) {
    const input, expected = "steven", "nevets"
    result := Reverse(input)
    if result != expected {
        t.Errorf("Expected %v, Received %v",
                 expected, result)
    }
}

func TestStepDownEven(t *testing.T) {
    const input, expected = 2000, 1991
    result := StepDown(input)
    if result != expected {
        t.Errorf("Expected %v, Received %v",
                 expected, result)
    }
}

func TestStepDownOdd(t *testing.T) {
    const input, expected = 99999, 99899
    result := StepDown(input)
    if result != expected {
        t.Errorf("Expected %v, Received %v",
                 expected, result)
    }
 }

func TestBoundaryNegative(t *testing.T) {
    const input, expected = 87, false
    result := AtBoundary(input)
    if result != expected {
        t.Errorf("Expected %v, Received %v",
                 expected, result)
    }
}

func TestBoundaryPositive(t *testing.T) {
    const input, expected = 100, true
    result := AtBoundary(input)
    if result != expected {
        t.Errorf("Expected %v, Received %v",
                 expected, result)
    }
}

func BenchmarkGivenProblem(b *testing.B) {
    for i := 0; i < b.N; i++ {
        LargestWithLength(3)
    }
}
