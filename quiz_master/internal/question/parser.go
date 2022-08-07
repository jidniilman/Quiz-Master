package question

import (
	"errors"
	"jidniilman/quiz-master/pkg/command"
	"jidniilman/quiz-master/pkg/utils"
	"strconv"
)

// The parser is a middleman between the question model and executor controller. It converts given command to proper
// question format

// ParseCreateUpdateArgs need at least 3 arguments to be compliance with create and update question, and more than
// 3 arguments are ignored
func (q *Question) ParseCreateUpdateArgs(cmd *command.Command) error {
	var err error

	if len(cmd.Args) < 3 {
		return errors.New("error: more arguments needed")
	}
	if q.ID, err = strconv.Atoi(cmd.Args[0]); err != nil {
		return err
	}
	q.Question = cmd.Args[1]
	if q.Answer, err = strconv.Atoi(cmd.Args[2]); err != nil {
		return err
	}
	return err
}

// ParseNumberArgs need at least 1 arguments to be compliance with command that have 1 argument, and more than
// 1 arguments are ignored
func (q *Question) ParseNumberArgs(cmd *command.Command) error {
	var err error

	if len(cmd.Args) < 1 {
		return errors.New("error: more arguments needed")
	}
	if q.ID, err = strconv.Atoi(cmd.Args[0]); err != nil {
		return err
	}
	return err
}

// ParseNumberAnswerArgs need at least 2 arguments to be compliance with command that have 2 argument, and more than
// 2 arguments are ignored
func (q *Question) ParseNumberAnswerArgs(cmd *command.Command) error {
	var err error

	if len(cmd.Args) < 2 {
		return errors.New("error: more arguments needed")
	}
	if q.ID, err = strconv.Atoi(cmd.Args[0]); err != nil {
		return err
	}
	if q.Answer, err = strconv.Atoi(cmd.Args[1]); err != nil {
		// Try parse string literal args to int
		if q.Answer, err = utils.WordsToInt(cmd.Args[1]); err != nil {
			return err
		}
	}
	return err
}
