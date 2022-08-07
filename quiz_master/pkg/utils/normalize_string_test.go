package utils

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

type NormalizeQuotesTestCase struct {
	have string
	want string
}

var normalizeTestCases = []NormalizeQuotesTestCase{
	{
		have: "\"How many letters are there in the English alphabet?\"",
		want: "\"How many letters are there in the English alphabet?\"",
	},
	{
		have: "“How many vowels are there in the English alphabet?”",
		want: "\"How many vowels are there in the English alphabet?\"",
	},
	{
		have: "‘first‘",
		want: "'first'",
	},
}

func TestNormalizeQuotes(t *testing.T) {
	for i, testCase := range normalizeTestCases {
		t.Logf("Test #%d for TestNormalizeQuotes Executed", i)
		cleaned := strings.Map(NormalizeQuotes, testCase.have)
		assert.Equal(t, testCase.want, cleaned, "String Result doesn't match with expectation")
	}
}

func BenchmarkNormalizeQuotes(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = strings.Map(NormalizeQuotes, normalizeTestCases[1].have)
	}
}
