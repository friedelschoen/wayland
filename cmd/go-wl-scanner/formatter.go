package main

import (
	"io"
	"os"
	"os/exec"
)

type formatWriter struct {
	io.WriteCloser
	cmd *exec.Cmd
}

func (fmter *formatWriter) Close() error {
	err := fmter.WriteCloser.Close()
	if err != nil {
		return err
	}
	return fmter.cmd.Wait()
}

type nopWriteCloser struct {
	io.Writer
}

func (*nopWriteCloser) Close() error { return nil }

// NewFormatter creates a WriteCloser, which formats the incoming Go-code.
// It uses `gofmt` command as backend, if it fails to execute the command,
// a NopCloser is returned which just wraps the input io.Writer
func NewFormatter(dst io.Writer) (io.WriteCloser, error) {
	cmd := exec.Command("gofmt")
	cmd.Stdout = dst
	cmd.Stderr = os.Stderr
	pipe, err := cmd.StdinPipe()
	if err != nil {
		return &nopWriteCloser{dst}, err
	}
	if err := cmd.Start(); err != nil {
		return &nopWriteCloser{dst}, err
	}
	return &formatWriter{WriteCloser: pipe, cmd: cmd}, nil
}
