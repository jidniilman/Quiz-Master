package command

type Command struct {
	Cmd  string
	Args []string
}

// Command is main communication protocol between user, business layer, and data layer

// GetArgs get and parse the inputted command and arguments from user and format into appropriate Command struct format.
// This function simply breakdown user input into recognizable command and arguments and respecting double-quoted text
// for question sentence and answer that have multiple words separated by spaces.
func GetArgs(str string) Command {
	var args []string
	var isQuotedRune = false
	var argRunes = make([]rune, 0)

	for _, char := range str {
		if char == ' ' && !isQuotedRune {
			args = append(args, string(argRunes))
			argRunes = argRunes[:0]
			continue
		}
		if char == '"' {
			isQuotedRune = !isQuotedRune
			continue
		}
		argRunes = append(argRunes, char)
	}
	// Add last argument
	args = append(args, string(argRunes))
	return Command{Cmd: args[0], Args: args[1:]}
}
