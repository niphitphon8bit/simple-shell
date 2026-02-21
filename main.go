package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		pwdir, err := os.Getwd()
		if err != nil {
			log.Fatalf("Error getting present work directory: %v", err)
		}

		dir := filepath.Base(pwdir)

		hostname, err := os.Hostname()
		if err != nil {
			log.Fatalf("Error getting hostname: %v", err)
		}

		user, err := user.Current()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(hostname, "\\", user.Username, "\\", dir, " > ")

		// Read the keyword input.
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
		}

		// Handle the execution of the input.
		if err = execInput(input); err != nil {
			if errors.Is(err, os.ErrExist) {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, err)
		}
	}
}

// Execute command function
func execInput(input string) error {
	// Remove th newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input to separate the command and the arguments.
	args := strings.Split(input, " ")

	switch args[0] {
	case "cd":
		// 'cd' to home dir with empty path not yet supported.
		if len(args) < 2 {
			return errNoPath
		}
		return os.Chdir(args[1])
	case "exit":
		return errExit
	case "^[[A":
	case "^[[B":
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}
