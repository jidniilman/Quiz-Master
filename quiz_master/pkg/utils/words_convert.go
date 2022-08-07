package utils

import (
	"strings"
)

// This sort library is used for converting english number words into integer.

var units = map[string]int{
	"one":       1,
	"two":       2,
	"three":     3,
	"four":      4,
	"five":      5,
	"six":       6,
	"seven":     7,
	"eight":     8,
	"nine":      9,
	"ten":       10,
	"eleven":    11,
	"twelve":    12,
	"thirteen":  13,
	"fourteen":  14,
	"fifteen":   15,
	"sixteen":   16,
	"seventeen": 17,
	"eighteen":  18,
	"nineteen":  19,
}

var tens = map[string]int{
	"twenty":  20,
	"thirty":  30,
	"forty":   40,
	"fifty":   50,
	"sixty":   60,
	"seventy": 70,
	"eighty":  80,
	"ninety":  90,
}

var powerOfTen = map[string]int{
	"hundred":  100,
	"thousand": 1000,
	"million":  1000000,
	"billion":  1000000000,
}

// WordsToInt with assumption no 'and' between numbers. Support up to billions. Need more than billions? just add more
// power of tens, and it will perfectly work fine
func WordsToInt(str string) (int, error) {
	args := strings.Split(strings.ToLower(str), " ")
	// Check for single unit and single word, and return immediately
	if val, found := units[str]; found {
		return val, nil
	}
	// Check for multiple words and store to a slice
	eachNumber := make([]int, len(args))
	for i := 0; i < len(args); i++ {
		// Check if unit
		if val, found := units[args[i]]; found {
			eachNumber[i] = val
			continue
		}
		// Check if tens
		if val, found := tens[args[i]]; found {
			eachNumber[i] = val
			continue
		}
		// Check if powers of ten
		if val, found := powerOfTen[args[i]]; found {
			eachNumber[i] = val
			continue
		}
	}

	// Normalize the Tens
	resultTensLower := make([]int, 0)
	CombineTensLower(eachNumber, &resultTensLower)
	// Uncomment to debug the resultTensLower
	// fmt.Println(resultTensLower)

	// Normalize the prefix and suffix of hundred
	resultHundredLower := make([]int, 0)
	CombineHundredLower(resultTensLower, &resultHundredLower)
	// Uncomment to debug the resultHundredLower
	// fmt.Println(resultHundredLower)

	// Normalize the powers of ten
	resultPower := make([]int, 0)
	CombinePower(resultHundredLower, &resultPower)
	// Uncomment to debug the resultPower
	// fmt.Println(resultPower)

	// Sum normalized slice
	finalResult := SumSlice(resultPower)

	// Uncomment to debug the finalResult
	// fmt.Println(finalResult)

	// Check if unit
	return finalResult, nil
}

// CombineTensLower into single integer
func CombineTensLower(eachNumber []int, result *[]int) {
	for i := 0; i < len(eachNumber); i++ {
		if i < len(eachNumber)-1 && eachNumber[i] > eachNumber[i+1] && eachNumber[i] < 99 {
			*result = append(*result, eachNumber[i]+eachNumber[i+1])
			i++
			continue
		}
		*result = append(*result, eachNumber[i])
	}
}

// CombineHundredLower into single integer, including prefix and suffix
func CombineHundredLower(eachNumber []int, result *[]int) {
	for i := 0; i < len(eachNumber); i++ {
		// Hundred always have prefix
		if i > 0 &&
			eachNumber[i] == 100 &&
			eachNumber[i-1] < eachNumber[i] {
			merged := eachNumber[i-1] * eachNumber[i]
			// But hundred have optional suffix
			if i < len(eachNumber)-1 && eachNumber[i] > eachNumber[i+1] {
				merged += eachNumber[i+1]
				i++
			}
			// remove hundred prefix from result
			if len(*result) > 0 {
				*result = (*result)[:len(*result)-1]
			}
			*result = append(*result, merged)
			continue
		}
		*result = append(*result, eachNumber[i])
	}
}

// CombinePower multiply prefix with power of tens into single integer
func CombinePower(eachNumber []int, result *[]int) {
	for i := 0; i < len(eachNumber); i++ {
		// Powers always have prefix
		if i > 0 &&
			eachNumber[i] >= 1000 &&
			eachNumber[i-1] < eachNumber[i] {
			merged := eachNumber[i-1] * eachNumber[i]
			// remove powers prefix from result
			if len(*result) > 0 {
				*result = (*result)[:len(*result)-1]
			}
			*result = append(*result, merged)
			continue
		}
		*result = append(*result, eachNumber[i])
	}
}

// SumSlice into one integer
func SumSlice(eachNumber []int) int {
	sum := 0
	for i := 0; i < len(eachNumber); i++ {
		sum += eachNumber[i]
	}
	return sum
}
