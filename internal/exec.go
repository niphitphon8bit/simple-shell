package internal

import (
	"os"
	"os/exec"
	"strings"
)

// ExecInput executes a single shell command line.
func ExecInput(input string) error {
	// Normalize user input and ignore empty commands.
	input = strings.TrimSpace(input)
	if input == "" {
		return nil
	}

	// Split input to separate the command and arguments.
	args := strings.Fields(input)

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
