// Package main is the entry point of the application quiz_master
package main

import (
	"fmt"
	"jidniilman/quiz-master/internal/cli"
)

// entry point of quiz_master application
func main() {
	// start an instance of CLI application
	if err := cli.NewCLI(); err != nil {
		fmt.Println(err)
	}
}
