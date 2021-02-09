package process

import (
	"context"
	"errors"
	"os/exec"
)

type Process interface {
	Start() error
	Stop() error
	IsRunning() bool
	GetInfo() *ProcessInfo
}

const ProcessTimeoutSeconds = 10

type process struct {
	*exec.Cmd
	cancelFunc context.CancelFunc
	running    bool

	path   string
	args   []string
	output []byte
	error  error
}

func NewProcess(name string, args ...string) Process {
	ctx, cancelFunc := context.WithCancel(context.Background())

	return &process{
		Cmd:        exec.CommandContext(ctx, name, args...),
		cancelFunc: cancelFunc,
		running:    true,
		path:       name,
		args:       args,
	}
}

func (p *process) IsRunning() bool {
	return p.running
}

func (p *process) Start() error {
	if !p.IsRunning() {
		return errors.New("process already stopped")
	}

	output, err := p.Cmd.CombinedOutput()

	p.running = false
	p.output = output
	p.error = err

	return err
}

func (p *process) Stop() error {
	if !p.IsRunning() {
		return errors.New("process already stopped")
	}
	p.cancelFunc()
	p.running = false
	return nil
}

func (p *process) GetInfo() *ProcessInfo {
	return &ProcessInfo{
		Path:      p.Cmd.Path,
		Args:      p.Cmd.Args[1:],
		IsRunning: p.running,
		Output:    p.output,
		Error:     p.error,
	}
}
