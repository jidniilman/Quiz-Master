package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type WordsToIntTestCase struct {
	have string
	want int
}

type CombineTestCase struct {
	have []int
	want []int
}

type SumSliceTestCase struct {
	have []int
	want int
}

var wordsTestCases = []WordsToIntTestCase{
	{
		have: "one hundred twenty three thousand four hundred fifty six",
		want: 123456,
	},
	{
		have: "five billion five hundred forty four million eight hundred seventy nine thousand five hundred ninety nine",
		want: 5544879599,
	},
	{
		have: "twenty six",
		want: 26,
	},
	{
		have: "five",
		want: 5,
	},
}

var combineTensTestCases = []CombineTestCase{
	{
		have: []int{20, 6, 50, 2, 10, 4},
		want: []int{26, 52, 14},
	},
	{
		have: []int{90, 9, 80, 8, 20, 7},
		want: []int{99, 88, 27},
	},
}

var combineHundredTestCases = []CombineTestCase{
	{
		have: []int{2, 100, 52, 2, 1000, 4},
		want: []int{252, 2, 1000, 4},
	},
	{
		have: []int{1, 100, 9, 1000000, 5, 100},
		want: []int{109, 1000000, 500},
	},
}

var combinePowerTestCases = []CombineTestCase{
	{
		have: []int{252, 1000, 4},
		want: []int{252000, 4},
	},
	{
		have: []int{109, 1000000, 527, 1000, 26},
		want: []int{109000000, 527000, 26},
	},
}

var sumSliceTestCases = []SumSliceTestCase{
	{
		have: []int{252000, 4},
		want: 252004,
	},
	{
		have: []int{109000000, 527000, 26},
		want: 109527026,
	},
}

func TestWordsToInt(t *testing.T) {
	for i, testCase := range wordsTestCases {
		t.Logf("Test #%d for TestWordsToInt Executed", i)
		converted, err := WordsToInt(testCase.have)
		assert.Nil(t, err, "Have error")
		assert.Equal(t, testCase.want, converted, "Converted result doesn't match with expectation")
	}
}

func TestCombineTensLower(t *testing.T) {
	for i, testCase := range combineTensTestCases {
		t.Logf("Test #%d for TestCombineTensLower Executed", i)
		resultTensLower := make([]int, 0)
		CombineTensLower(testCase.have, &resultTensLower)
		assert.Equal(t, testCase.want, resultTensLower, "Combined result doesn't match with expectation")
	}
}

func TestCombineHundredLower(t *testing.T) {
	for i, testCase := range combineHundredTestCases {
		t.Logf("Test #%d for TestCombineHundredLower Executed", i)
		resultHundredLower := make([]int, 0)
		CombineHundredLower(testCase.have, &resultHundredLower)
		assert.Equal(t, testCase.want, resultHundredLower, "Combined result doesn't match with expectation")
	}
}

func TestCombinePower(t *testing.T) {
	for i, testCase := range combinePowerTestCases {
		t.Logf("Test #%d for TestCombinePower Executed", i)
		resultPower := make([]int, 0)
		CombinePower(testCase.have, &resultPower)
		assert.Equal(t, testCase.want, resultPower, "Combined result doesn't match with expectation")
	}
}

func TestSumSlice(t *testing.T) {
	for i, testCase := range sumSliceTestCases {
		t.Logf("Test #%d for TestSumSlice Executed", i)
		finalResult := SumSlice(testCase.have)
		assert.Equal(t, testCase.want, finalResult, "Sum result doesn't match with expectation")
	}
}

func BenchmarkWordsToInt(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _ = WordsToInt(wordsTestCases[0].have)
	}
}
