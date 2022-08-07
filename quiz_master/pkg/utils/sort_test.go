package utils

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type BubbleSortTestCase struct {
	have []int
	want []int
}

var sortTestCases = []BubbleSortTestCase{
	{
		have: []int{0, 6, 5, 3, 2, 4, 1},
		want: []int{0, 1, 2, 3, 4, 5, 6},
	},
	{
		have: []int{6, 1, 7, 9, 3, 8, 2, 5, 4, 0},
		want: []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	},
}

func TestBubbleSort(t *testing.T) {
	for i, testCase := range sortTestCases {
		t.Logf("Test #%d for TestBubbleSort Executed", i)
		BubbleSort(testCase.have)
		assert.Equal(t, testCase.want, testCase.have, "Sorted result doesn't match with expectation")
	}
}

func BenchmarkBubbleSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		BubbleSort(sortTestCases[0].have)
	}
}
