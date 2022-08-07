package cli

import (
	"fmt"
	"jidniilman/quiz-master/internal/question"
	"jidniilman/quiz-master/pkg/command"
)

// ExecCreateOrUpdateQuestion with assumption no overriding existing question and no force create for not existing question
func ExecCreateOrUpdateQuestion(create bool, cmd *command.Command) {
	// Prepare an empty question
	q := question.Question{}
	// Parse command into a proper question format
	if err := q.ParseCreateUpdateArgs(cmd); err != nil {
		fmt.Println(err)
		return
	}
	// This is for Create Question, check if it's exist first
	if exist := q.IsQuestionExist(); !exist && create {
		// Create Question mean we store our built question into data store blindly, and
		// we don't know what the data store is from this layer
		if err := q.CreateOrUpdateQuestion(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Question Number %d is created:\n", q.ID)
		fmt.Printf("Q: %s \n", q.Question)
		fmt.Printf("A: %d \n", q.Answer)
		return
	}
	// This is for Update Question, check if it's exist first
	if exist := q.IsQuestionExist(); exist && !create {
		if err := q.CreateOrUpdateQuestion(); err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Question Number %d is updated:\n", q.ID)
		fmt.Printf("Q: %s \n", q.Question)
		fmt.Printf("A: %d \n", q.Answer)
		return
	}
	if create {
		fmt.Printf("Sorry Question with Number %d already exist\n", q.ID)
		return
	}
	if !create {
		fmt.Printf("Sorry Question with Number %d not exist\n", q.ID)
		return
	}
}

// ExecDeleteQuestion for question by given number
func ExecDeleteQuestion(cmd *command.Command) {
	// Prepare an empty question
	q := question.Question{}
	// Parse command into a proper question format
	if err := q.ParseNumberArgs(cmd); err != nil {
		fmt.Println(err)
		return
	}
	if exist := q.IsQuestionExist(); exist {
		q.DeleteQuestion()
		fmt.Printf("Question Number %d was deleted!\n", q.ID)
		return
	}
	fmt.Printf("Sorry Question with Number %d not exist\n", q.ID)
}

// ExecReadQuestion for question by given number
func ExecReadQuestion(cmd *command.Command) {
	// Prepare an empty question
	q := question.Question{}
	// Parse command into a proper question format
	if err := q.ParseNumberArgs(cmd); err != nil {
		fmt.Println(err)
		return
	}
	if exist := q.IsQuestionExist(); exist {
		q.ReadQuestion()
		fmt.Println("Q:", q.Question)
		fmt.Println("A:", q.Answer)
		return
	}
	fmt.Printf("Sorry Question with Number %d not exist\n", q.ID)
}

// ExecBrowseQuestions that already stored
func ExecBrowseQuestions() {
	// We get the sorted questions first
	qs := question.BrowseQuestions()
	// and then with our tabwriter, we print questions with nice formatting
	DisplayQuestions(qs)
}

// ExecAnswerQuestion that already stored
func ExecAnswerQuestion(cmd *command.Command) {
	// Prepare an empty question
	q := question.Question{}
	// Parse command into a proper question format
	if err := q.ParseNumberAnswerArgs(cmd); err != nil {
		fmt.Println(err)
		return
	}
	if exist := q.IsQuestionExist(); exist {
		isAnswerTrue := q.AnswerQuestion()
		if isAnswerTrue {
			fmt.Println("Your answer is Correct!")
			return
		}
		fmt.Println("Wrong Answer!")
		return
	}
	fmt.Printf("Sorry Question with Number %d is not exist\n", q.ID)
}
