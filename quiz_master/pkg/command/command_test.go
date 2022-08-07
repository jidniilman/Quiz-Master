package command

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type GetArgsTestCase struct {
	have string
	want Command
}

var testCases = []GetArgsTestCase{
	{
		have: "create_question 1 \"How many letters are there in the English alphabet?\" 26",
		want: Command{
			Cmd:  "create_question",
			Args: []string{"1", "How many letters are there in the English alphabet?", "26"},
		},
	},
	{
		have: "create_question 2 \"How many vowels are there in the English alphabet?\" 5",
		want: Command{
			Cmd:  "create_question",
			Args: []string{"2", "How many vowels are there in the English alphabet?", "5"},
		},
	},
	{
		have: "questions",
		want: Command{
			Cmd:  "questions",
			Args: []string{},
		},
	},
	{
		have: "question 2",
		want: Command{
			Cmd:  "question",
			Args: []string{"2"},
		},
	},
	{
		have: "delete_question 1",
		want: Command{
			Cmd:  "delete_question",
			Args: []string{"1"},
		},
	},
	{
		have: "answer_question 2 five",
		want: Command{
			Cmd:  "answer_question",
			Args: []string{"2", "five"},
		},
	},
	{
		have: "answer_question 2 5",
		want: Command{
			Cmd:  "answer_question",
			Args: []string{"2", "5"},
		},
	},
	{
		have: "exit",
		want: Command{
			Cmd:  "exit",
			Args: []string{},
		},
	},
}

func TestGetArgs(t *testing.T) {
	for i, testCase := range testCases {
		t.Logf("Test #%d for TestGetArgs Executed", i)
		getCmd := GetArgs(testCase.have)
		assert.Equal(t, testCase.want, getCmd, "Command Result doesn't match with expectation")
	}
}

func BenchmarkGetArgs(b *testing.B) {
	for i := 0; i < b.N; i++ {
		GetArgs(testCases[0].have)
	}
}
