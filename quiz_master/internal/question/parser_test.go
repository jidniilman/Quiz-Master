package question

import (
	"github.com/stretchr/testify/assert"
	"jidniilman/quiz-master/pkg/command"
	"testing"
)

// To be compared with
var parserQuestions = []Question{
	{
		ID:       1,
		Question: "How many letters are there in the English alphabet?",
		Answer:   26,
	},
	{
		ID:       2,
		Question: "How many vowels are there in the English alphabet?",
		Answer:   5,
	},
}

// Command from user
var createCommands = []command.Command{
	{
		Cmd:  "create_question",
		Args: []string{"1", "How many letters are there in the English alphabet?", "26"},
	},
	{
		Cmd:  "create_question",
		Args: []string{"2", "How many vowels are there in the English alphabet?", "5"},
	},
}

var readCommands = []command.Command{
	{
		Cmd:  "question",
		Args: []string{"1"},
	},
	{
		Cmd:  "question",
		Args: []string{"2"},
	},
}

var answerCommands = []command.Command{
	{
		Cmd:  "answer_question",
		Args: []string{"1", "twenty six"},
	},
	{
		Cmd:  "answer_question",
		Args: []string{"2", "five"},
	},
}

func TestParseCreateUpdateArgs(t *testing.T) {
	for i, cc := range createCommands {
		t.Logf("Test #%d for TestParseCreateUpdateArgs Executed", i)
		qTest := Question{}
		err := qTest.ParseCreateUpdateArgs(&cc)
		if err != nil {
			assert.Nil(t, err, "Error when parsing the arguments")
		}
		assert.Equal(t, qTest, parserQuestions[i], "Parsing result are different")
	}
}

func TestParseNumberArgs(t *testing.T) {
	for i, rc := range readCommands {
		t.Logf("Test #%d for TestParseNumberArgs Executed", i)
		qTest := Question{}
		err := qTest.ParseNumberArgs(&rc)
		if err != nil {
			assert.Nil(t, err, "Error when parsing the arguments")
		}
		// We are just check the parsed ID
		assert.Equal(t, qTest.ID, parserQuestions[i].ID, "Parsing result ID is different")
	}
}

func TestParseNumberAnswerArgs(t *testing.T) {
	for i, ac := range answerCommands {
		t.Logf("Test #%d for TestParseNumberAnswerArgs Executed", i)
		qTest := Question{}
		err := qTest.ParseNumberAnswerArgs(&ac)
		if err != nil {
			assert.Nil(t, err, "Error when parsing the arguments")
		}
		// We are just check the parsed ID and Answer
		assert.Equal(t, qTest.ID, parserQuestions[i].ID, "Parsing result ID is different")
		assert.Equal(t, qTest.Answer, parserQuestions[i].Answer, "Parsing result Answer is different")
	}
}
