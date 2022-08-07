package question

import "jidniilman/quiz-master/pkg/utils"

// Question consist of ID (Number), Question, and Answer. The data type is not generic.
type Question struct {
	ID       int
	Question string
	Answer   int
}

// In-memory data store for questions
// Encapsulated and protected so the cli/exec didn't know anything about questions stored
// You can expand this data store into more complex storage such as relational database, non-relational database, redis
// or other specified storage system suitable for your needs as it is not coupled with business layer, but definitely
// you need to change some logic of receiver function of type Question.
var questions = make(map[int]Question)

// IsQuestionExist in our data store
func (q *Question) IsQuestionExist() bool {
	// Check if key q.id is exists on map qs
	_, exist := questions[q.ID]
	return exist
}

// CreateOrUpdateQuestion and store it into our data store
func (q *Question) CreateOrUpdateQuestion() error {
	// Appending q into qs using pointer
	questions[q.ID] = *q
	return nil
}

// DeleteQuestion from our data store
func (q *Question) DeleteQuestion() {
	// Delete question in qs by q.ID
	delete(questions, q.ID)
}

// ReadQuestion from our data store by ID
func (q *Question) ReadQuestion() {
	// Just assign Question and Answer from our data store
	q.Question = questions[q.ID].Question
	q.Answer = questions[q.ID].Answer
}

// BrowseQuestions from our data store and sort it first
func BrowseQuestions() []Question {
	keys := make([]int, 0, len(questions))
	qs := make([]Question, 0, len(questions))

	// Fetch keys from questions
	for k := range questions {
		keys = append(keys, k)
	}
	// Sort the keys
	utils.BubbleSort(keys)
	// Arrange into a new slice of questions from our data store by sorted keys
	for _, k := range keys {
		qs = append(qs, questions[k])
	}
	return qs
}

// AnswerQuestion is just a quick comparison check between given answer and question answer from our data store by ID
func (q *Question) AnswerQuestion() bool {
	if q.Answer == questions[q.ID].Answer {
		return true
	}
	return false
}
