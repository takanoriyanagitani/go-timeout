package timeout_test

import (
	"context"
	"errors"
	"os/exec"
	"testing"
	"time"

	"github.com/takanoriyanagitani/go-timeout"
)

func TestCommandBuilderRun(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name        string
		commandName string
		args        []string
		timeout     time.Duration
		expectError bool
		errorCheck  func(*testing.T, error)
	}{
		{
			name:        "CommandTimesOut",
			commandName: "sleep",
			args:        []string{"3"},
			timeout:     1 * time.Second,
			expectError: true,
			errorCheck: func(t *testing.T, err error) {
				if !errors.Is(err, context.DeadlineExceeded) {
					t.Fatalf("expected context.DeadlineExceeded, got %T: %v", err, err)
				}
			},
		},
		{
			name:        "CommandCompletes",
			commandName: "sleep",
			args:        []string{"1"},
			timeout:     3 * time.Second,
			expectError: false,
			errorCheck:  nil, // No error expected
		},
		{
			name:        "CommandFails",
			commandName: "false",
			args:        []string{},
			timeout:     3 * time.Second,
			expectError: true,
			errorCheck: func(t *testing.T, err error) {
				var exitErr *exec.ExitError
				if !errors.As(err, &exitErr) {
					t.Fatalf("expected exec.ExitError, got %T: %v", err, err)
				}
				if exitErr.ExitCode() == 0 {
					t.Fatal("expected non-zero exit code, got 0")
				}
			},
		},
	}

	for _, tc := range testCases {
		tc := tc // capture range variable
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			builder := timeout.CommandBuilder{
				Name:      tc.commandName,
				Arguments: tc.args,
				Timeout:   tc.timeout,
			}

			err := builder.Run(context.Background())

			if tc.expectError {
				if err == nil {
					t.Fatal("expected an error, got nil")
				}
				if tc.errorCheck != nil {
					tc.errorCheck(t, err)
				}
			} else {
				if err != nil {
					t.Fatalf("expected no error, got: %v", err)
				}
			}
		})
	}
}
