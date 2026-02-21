package internal

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestExecInput(t *testing.T) {
	tests := []struct {
		name        string
		givenInput  string
		expectedDir string
		expectedErr error
	}{
		{
			name:        "happy: ls",
			givenInput:  "ls",
			expectedErr: nil,
		},
		{
			name:        "happy: ls -l -a",
			givenInput:  "ls -l -a",
			expectedErr: nil,
		},
		{
			name:        "happy: exit",
			givenInput:  "exit",
			expectedErr: ErrExit,
		},
		{
			name:        "happy: cd",
			givenInput:  "cd",
			expectedErr: ErrNoPath,
		},
		{
			name:        "happy: cd /",
			givenInput:  "cd /",
			expectedErr: nil,
			expectedDir: "/",
		},
		{
			name:        "bad: wrong spell command",
			givenInput:  "exst",
			expectedErr: exec.ErrNotFound,
		},
	}

	for _, eachTest := range tests {
		t.Run(eachTest.name, func(t *testing.T) {
			// run each test here

			if err := ExecInput(eachTest.givenInput); !errors.Is(err, eachTest.expectedErr) {
				t.Errorf("execInput() error = %v, want Err %v", err, eachTest.expectedErr)
			}

			// test for change directory only
			if eachTest.expectedDir != "" {
				curDir, err := os.Getwd()
				if err != nil {
					t.Errorf("Failed to get new directory: %v", err)
				}
				if eachTest.expectedDir != curDir {
					t.Errorf("Failed to change to desired directory.")
				}
			}
		})
	}
}
