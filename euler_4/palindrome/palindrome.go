
package palindrome

import (
    "log"
    "math"
    "strconv"
    "strings"
)

func LargestWithLength(digits int) (largest int) {
    root, _ := strconv.Atoi(strings.Repeat("9", digits))
    largest = root * root + 1
    for !FactorsToLength(largest, digits) {
        largest = StepDown(largest)
    }
    return
}

func StepDown(current int) (next int) {
    var lDigits, rDigits string
    var err error
    next = current - 1
    digitS := strconv.Itoa(next)
    if len(digitS) % 2 == 0 {
        // rDigits := digitS[len(digitS)/2, len(digitS)]
        lDigits  = digitS[:len(digitS)/2]
        rDigits = Reverse(lDigits)
    } else {
        // Left and right share a digit
        lDigits  = digitS[:len(digitS)/2 + 1]
        rDigits = Reverse(lDigits)[1:len(lDigits)]
    }
    if AtBoundary(current) {
        if len(digitS) != 1 + len(lDigits) + len(rDigits) {
            log.Printf("36: Digits: %v, Left Len: %v, Right Len: %v",
                       len(digitS), len(lDigits), len(rDigits))
        }
    } else {
        if len(digitS) != len(lDigits) + len(rDigits) {
            log.Printf("41: Digits: %v, Left Len: %v, Right Len: %v",
                       len(digitS), len(lDigits), len(rDigits))
        }
    }
    next, err = strconv.Atoi(lDigits + rDigits)
    if err != nil {
        panic(err)
    }
    if current > next {
        return
    }
    if !AtBoundary(current) {
        lNum, err := strconv.Atoi(lDigits)
        lNum--
        lDigits = strconv.Itoa(lNum)
        if err != nil {
            panic(err)
        }
    }
    if len(digitS) % 2 == 0 {
        rDigits = Reverse(lDigits)
    } else {
        rDigits = Reverse(lDigits)[1:len(lDigits)]
    }
    if AtBoundary(current) {
        if len(digitS) != 1 + len(lDigits) + len(rDigits) {
            log.Printf("68: Digits: %v, Left Len: %v, Right Len: %v",
                       len(digitS), len(lDigits), len(rDigits))
        }
    } else {
        if len(digitS) != len(lDigits) + len(rDigits) {
            log.Printf("73: Digits: %v, Left Len: %v, Right Len: %v",
                      len(digitS), len(lDigits), len(rDigits))
        }
    }
    next, err = strconv.Atoi(lDigits + rDigits)
    if err != nil {
        panic(err)
    }
    return
}

func FactorsToLength(number, digits int) bool {
    for i := int(math.Sqrt(float64(number))); i < int(math.Pow(10.,
    float64(digits))); i++ {
        if number % i == 0 {
            return true
        }
    }
    return false
}

func Reverse(input string) (output string) {
    if len(input) == 1 {
        output = input
    } else {
        output = Reverse(input[1:len(input)]) + input[0:1]
    }
    return
}

func AtBoundary(number int) bool {
    return int(math.Log10(float64(number))) > int(math.Log10(float64(number - 1)))
 }
