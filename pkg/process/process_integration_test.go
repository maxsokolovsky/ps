package process_test

import (
	"testing"
	"time"

	"ps/pkg/process"
)

func TestUnixProcess(t *testing.T) {
	p := process.NewProcess("sleep", "2s")

	go p.Start()
	time.Sleep(100 * time.Millisecond)

	if !p.IsRunning() {
		t.Errorf("process is supposed to be running")
	}

	time.Sleep(time.Second)

	err := p.Stop()
	if err != nil {
		t.Fatal(err)
	}

	if p.IsRunning() {
		t.Errorf("process is not supposed to be running")
	}
}
