package process_test

import (
	"task-manager/pkg/process"
	"testing"
)

func TestMockProcess(t *testing.T) {
	p := process.NewMockProcess()

	p.Start()

	if !p.IsRunning() {
		t.Errorf("process is supposed to be running")
	}

	err := p.Stop()
	if err != nil {
		t.Fatal(err)
	}

	if p.IsRunning() {
		t.Errorf("process is not supposed to be running")
	}
}
