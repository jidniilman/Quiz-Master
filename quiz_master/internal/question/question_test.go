package question

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testQuestions = []Question{
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

func TestIsQuestionExist(t *testing.T) {
	mockQuestionsData()
	for i := 0; i < len(questions); i++ {
		t.Logf("Test #%d for TestIsQuestionExist Executed", i)
		// We are making assumption with ID here, it is taken from i + 1
		q := Question{ID: i + 1}
		assert.Equal(t, q.IsQuestionExist(), true, "A Question doesn't exist")
	}
}

func TestCreateOrUpdateQuestion(t *testing.T) {
	questions = map[int]Question{}
	for i := 0; i < len(testQuestions); i++ {
		t.Logf("Test #%d for TestCreateOrUpdateQuestion Executed", i)
		assert.Nil(t, testQuestions[i].CreateOrUpdateQuestion(), "Error when creating/updating a question")
		assert.Equal(t, testQuestions[i], questions[testQuestions[i].ID], "A Question doesn't created/updated")
	}
}

func TestDeleteQuestion(t *testing.T) {
	mockQuestionsData()
	for i := 0; i < len(questions); i++ {
		t.Logf("Test #%d for TestDeleteQuestion Executed", i)
		q := Question{ID: i + 1}
		q.DeleteQuestion()
		assert.NotContains(t, questions, q.ID, "Error, a question not deleted")
	}
}

func TestReadQuestion(t *testing.T) {
	mockQuestionsData()
	for i := 0; i < len(questions); i++ {
		t.Logf("Test #%d for TestReadQuestion Executed", i)
		q := Question{ID: i + 1}
		q.ReadQuestion()
		// We add -1 because different indexing between testQuestions and mockQuestions
		assert.Equal(t, q, testQuestions[q.ID-1], "Error, a question is different with what we want")
	}
}

func TestBrowseQuestions(t *testing.T) {
	mockQuestionsData()
	t.Logf("Test for TestBrowseQuestions Executed")
	q := BrowseQuestions()
	for i, question := range q {
		assert.Equal(t, question, testQuestions[i], "Error, a question is different with what we want")
	}
}

func TestAnswerQuestion(t *testing.T) {
	mockQuestionsData()
	for i := 0; i < len(questions); i++ {
		t.Logf("Test #%d for TestAnswerQuestion Executed", i)
		// We are making assumption with ID here, it is taken from i + 1
		q := Question{ID: i + 1, Answer: testQuestions[i].Answer}
		assert.Equal(t, q.AnswerQuestion(), true, "A Question have wrong answer")
	}
}

func mockQuestionsData() {
	// We mock the data store
	questions = map[int]Question{
		1: {
			ID:       1,
			Question: "How many letters are there in the English alphabet?",
			Answer:   26,
		},
		2: {
			ID:       2,
			Question: "How many vowels are there in the English alphabet?",
			Answer:   5,
		},
	}
}
