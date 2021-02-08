package process

import (
	"context"
	"errors"
	"os/exec"
)

type Process interface {
	Start() ([]byte, error)
	Stop() error
	IsRunning() bool
	GetInfo() ProcessInfo
}

const ProcessTimeoutSeconds = 10

type ProcessInfo struct {
	Path string
	Args []string
}

type process struct {
	*exec.Cmd
	cancelFunc context.CancelFunc
	running    bool
}

func NewProcess(name string, args ...string) Process {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return &process{
		Cmd:        exec.CommandContext(ctx, name, args...),
		cancelFunc: cancelFunc,
		running:    true,
	}
}

func (p *process) IsRunning() bool {
	return p.running
}

func (p *process) Start() ([]byte, error) {
	if !p.IsRunning() {
		return nil, errors.New("process already stopped")
	}

	return p.Cmd.CombinedOutput()
}

func (p *process) Stop() error {
	if !p.IsRunning() {
		return errors.New("process already stopped")
	}
	p.cancelFunc()
	p.running = false
	return nil
}

func (p *process) GetInfo() ProcessInfo {
	return ProcessInfo{
		Path: p.Cmd.Path,
		Args: p.Cmd.Args[1:],
	}
}
