package cli

import (
	"fmt"
	"jidniilman/quiz-master/internal/question"
	"os"
	"text/tabwriter"
)

// DisplayHelp for user to know how to use the application and show information such as recognized commands and its usage
func DisplayHelp() {
	fmt.Println("Usage: COMMAND [ARGUMENTS]")
	fmt.Println()
	fmt.Println("Commands:")

	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "\t Short\t Command\t Arguments \t Description")
	fmt.Fprintln(w, "\t cq\t create_question\t <number> <question> <answer>\t Create a question with a number and answer")
	fmt.Fprintln(w, "\t uq\t update_question\t <number> <question> <answer>\t Update a question with a number and answer")
	fmt.Fprintln(w, "\t dq\t delete_question\t <number>\t Delete a question by given number")
	fmt.Fprintln(w, "\t q\t question\t <number>\t Show a question by given number")
	fmt.Fprintln(w, "\t qs\t questions\t \t Show all questions")
	fmt.Fprintln(w, "\t aq\t answer_question\t <number> <answer>\t Answer a question by given number")
	fmt.Fprintln(w, "\t \t version\t \t Show app version")
	fmt.Fprintln(w, "\t \t help\t \t Show help")
	fmt.Fprintln(w, "\t \t exit\t \t Close the program")
	fmt.Fprintln(w)
}

// DisplayVersion for user to know current version of this application
func DisplayVersion() {
	fmt.Println("Quiz Master version 1.0.0. By: Jidni Ilman")
}

// DisplayQuestions with formatted output using tabwriter
func DisplayQuestions(qs []question.Question) {
	w := new(tabwriter.Writer)
	w.Init(os.Stdout, 0, 0, 2, ' ', tabwriter.Debug)
	defer w.Flush()

	fmt.Fprintln(w, "Number\t Question\t Answer")
	for _, q := range qs {
		fmt.Fprintf(w, "%d\t %s\t %d\n", q.ID, q.Question, q.Answer)
	}
	fmt.Fprintln(w)
}
