package scheduler

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"task-manager/pkg/process"
	"time"
)

type Scheduler interface {
	SubmitProcess(string, ...string) string
	CancelProcess(string) error
	IsProcessRunning(string) bool
}

type result struct {
	id     string
	output []byte
	err    error
	info   process.ProcessInfo
}

type scheduler struct {
	m  sync.Map
	ch chan result
}

func New() *scheduler {
	s := &scheduler{
		ch: make(chan result),
	}
	go s.Listen()
	return s
}

func (s *scheduler) List() map[string]string {
	table := make(map[string]string)
	s.m.Range(func(k, v interface{}) bool {
		p := v.(process.Process)
		info := p.GetInfo()
		table[k.(string)] = fmt.Sprintf("%s %s", info.Path, strings.Join(info.Args, " "))
		return true
	})
	return table
}

func (s *scheduler) Listen() {
	for {
		r := <-s.ch
		s.printResult(os.Stdout, r)
		s.m.Delete(r.id)
	}
}

func (s *scheduler) printResult(w io.Writer, r result) {
	fmt.Fprintf(w, "ID: %q; task: %s %s; err: %v.\n", r.id, r.info.Path, strings.Join(r.info.Args, " "), r.err)
	if len(r.output) > 0 {
		fmt.Fprintf(w, string(r.output))
	}
}

func (s *scheduler) SubmitProcess(cmd string, args ...string) string {
	pid := randomId()
	p := process.NewProcess(cmd, args...)
	s.m.Store(pid, p)

	go func() {
		output, err := p.Start()
		r := result{pid, output, err, process.ProcessInfo{Path: cmd, Args: args}}
		s.ch <- r
	}()

	return pid
}

func (s *scheduler) CancelProcess(pid string) error {
	fmt.Println("XXXXXXXX", pid)
	if v, ok := s.m.Load(pid); ok {
		defer s.m.Delete(pid)
		p := v.(process.Process)
		return p.Stop()
	} else {
		return errors.New("process ID not found")
	}
}

func (s *scheduler) IsProcessRunning(pid string) bool {
	if v, ok := s.m.Load(pid); ok {
		p := v.(process.Process)
		return p.IsRunning()
	} else {
		return false
	}
}

func randomId() string {
	rand.Seed(time.Now().UnixNano())
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

	b := make([]rune, 5)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}
