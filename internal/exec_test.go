package internal

import (
	"errors"
	"os"
	"os/exec"
	"testing"
)

func TestExecInput(t *testing.T) {
	originalDir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get original directory: %v", err)
	}
	t.Cleanup(func() {
		if chdirErr := os.Chdir(originalDir); chdirErr != nil {
			t.Errorf("failed to restore directory: %v", chdirErr)
		}
	})

	tmpDir := t.TempDir()

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
			name:        "happy: empty input",
			givenInput:  "",
			expectedErr: nil,
		},
		{
			name:        "happy: whitespace input",
			givenInput:  "   \t   ",
			expectedErr: nil,
		},
		{
			name:        "happy: cd",
			givenInput:  "cd",
			expectedErr: ErrNoPath,
		},
		{
			name:        "happy: cd temp dir",
			givenInput:  "cd " + tmpDir,
			expectedErr: nil,
			expectedDir: tmpDir,
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
				expectedInfo, err := os.Stat(eachTest.expectedDir)
				if err != nil {
					t.Errorf("Failed to stat expected directory: %v", err)
				}
				currentInfo, err := os.Stat(curDir)
				if err != nil {
					t.Errorf("Failed to stat current directory: %v", err)
				}
				if !os.SameFile(expectedInfo, currentInfo) {
					t.Errorf("Failed to change to desired directory. got=%q want=%q", curDir, eachTest.expectedDir)
				}
			}
		})
	}
}
