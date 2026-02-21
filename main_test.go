package main

import (
	"errors"
	"os/exec"
	"testing"
)

func TestExistCommand(t *testing.T) {
	tests := []struct {
		name           string
		givenInput     string
		expectedOutput string
		expectedErr    error
	}{
		{
			name:        "happy: exit",
			givenInput:  "exit",
			expectedErr: errExit,
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
			if err := execInput(eachTest.givenInput); !errors.Is(err, eachTest.expectedErr) {
				t.Errorf("execInput() error = %v, want Err %v", err, eachTest.expectedErr)
			}
		})
	}
}
