package internal

import (
	"os"
	"os/exec"
	"strings"
)

// ExecInput executes a single shell command line.
func ExecInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split input to separate the command and arguments.
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path is not yet supported.
		if len(args) < 2 {
			return ErrNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		return ErrExit
	case "^[[A":
	case "^[[B":
	}

	// Prepare and execute the command.
	cmd := exec.Command(args[0], args[1:]...)
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout
	return cmd.Run()
}
