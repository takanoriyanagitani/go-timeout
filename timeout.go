package timeout

import (
	"context"
	"os"
	"os/exec"
	"time"
)

type CommandBuilder struct {
	Name      string
	Arguments []string
	Timeout   time.Duration
}

func (b CommandBuilder) Run(ctx context.Context) error {
	ctx, cancel := context.WithTimeout(ctx, b.Timeout)
	defer cancel()

	var cmd *exec.Cmd = exec.CommandContext(
		ctx,
		b.Name,
		b.Arguments...,
	)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err := cmd.Run()
	if ctx.Err() == context.DeadlineExceeded {
		return context.DeadlineExceeded
	}
	return err
}
