package handy

// refered from https://github.com/vimeo/go-util/blob/master/util/cmd.go

import (
	"errors"
	"os/exec"
	"time"
)

// ErrCommandTimeout is the error used when a command times out before completing.
var ErrCommandTimeout = errors.New("command timed out")

// RunCommand is a fuction that run a command,
func RunCommand(cmd *exec.Cmd) error {
	err := cmd.Start()
	if err != nil {
		return err
	}
	return nil
}

// RunCommandWithTimeout is a fuction that run a command,
// killing it if it does not finish before the specified timeout
func RunCommandWithTimeout(cmd *exec.Cmd, timeout time.Duration) error {
	done := make(chan error, 1)
	t := time.After(timeout)

	err := cmd.Start()
	if err != nil {
		return err
	}
	go func() {
		done <- cmd.Wait()
	}()
	select {
	case err := <-done:
		return err
	case <-t:
		cmd.Process.Kill()
		return ErrCommandTimeout
	}
}
