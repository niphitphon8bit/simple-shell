package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"os/user"
	"path/filepath"

	shell "simple-shell/internal"
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
		if err = shell.ExecInput(input); err != nil {
			if errors.Is(err, shell.ErrExit) {
				os.Exit(0)
			}
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
