// Package cli provide execution routing and initiating an infinite input-execute loop
package cli

import (
	"bufio"
	"fmt"
	"jidniilman/quiz-master/pkg/command"
	"jidniilman/quiz-master/pkg/utils"
	"os"
	"strings"
)

// NewCLI start an infinite loop for retrieving user command and arguments
func NewCLI() error {
	fmt.Println("Welcome to Quiz Master!")
	fmt.Println("- Run 'help' for more information")
	fmt.Println("- Run 'exit' to close the program")
	fmt.Print("> ")

	// Start retrieving user command and arguments
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		// Normalize double quote and quote
		normalizedInput := strings.Map(utils.NormalizeQuotes, scanner.Text())
		// Get and parse user command and arguments
		cmd := command.GetArgs(normalizedInput)
		// Execute command and arguments with provided questions data
		Execute(&cmd)
	}
	return nil
}

// Execute a command by switch and invoke the related method
func Execute(cmd *command.Command) {
	switch cmd.Cmd {
	case "exit":
		fmt.Println("Thank you!")
		os.Exit(0)
	case "help":
		DisplayHelp()
	case "version":
		DisplayVersion()
	case "create_question", "cq":
		ExecCreateOrUpdateQuestion(true, cmd)
	case "update_question", "uq":
		ExecCreateOrUpdateQuestion(false, cmd)
	case "delete_question", "dq":
		ExecDeleteQuestion(cmd)
	case "question", "q":
		ExecReadQuestion(cmd)
	case "questions", "qs":
		ExecBrowseQuestions()
	case "answer_question", "aq":
		ExecAnswerQuestion(cmd)
	default:
		fmt.Println("Unrecognized command. Try another one")
	}
	fmt.Print("> ")
}
