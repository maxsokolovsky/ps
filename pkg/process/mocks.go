package process

import (
	"errors"
	"fmt"
	"time"
)

type mockProcess struct {
	done    chan struct{}
	running bool
}

func NewMockProcess() Process {
	return &mockProcess{
		done:    make(chan struct{}),
		running: true,
	}
}

func (p *mockProcess) IsRunning() bool {
	return p.running
}

func (p *mockProcess) Start() error {
	if !p.IsRunning() {
		return errors.New("process already stopped")
	}
	go func() {
		for {
			select {
			case <-p.done:
				p.running = false
			default:
				fmt.Println("process running...")
				time.Sleep(time.Second)
			}
		}
	}()
	return nil
}

func (p *mockProcess) Stop() error {
	if !p.IsRunning() {
		return errors.New("process already stopped")
	}
	p.done <- struct{}{}
	close(p.done)
	return nil
}

func (p *mockProcess) GetInfo() *ProcessInfo {
	return &ProcessInfo{}
}
