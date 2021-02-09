package scheduler

import (
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strings"
	"sync"
	"time"

	"ps/pkg/process"
)

type Scheduler interface {
	SubmitProcess(string, ...string) string
	CancelProcess(string) error
	ProcessStatus(string) (*process.ProcessInfo, error)
}

type result struct {
	pid  string
	info *process.ProcessInfo
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

		// Give clients a minute to read the status before deleting.
		go func() {
			select {
			case <-time.After(time.Minute):
				s.m.Delete(r.pid)
			}
		}()
	}
}

func (s *scheduler) printResult(w io.Writer, r result) {
	fmt.Fprintf(w, "ID: %q; task: %s %s; err: %v.\n", r.pid, r.info.Path, strings.Join(r.info.Args, " "), r.info.Error)
	if len(r.info.Output) > 0 {
		fmt.Fprintf(w, string(r.info.Output))
	}
}

func (s *scheduler) SubmitProcess(cmd string, args ...string) string {
	pid := randomId()
	p := process.NewProcess(cmd, args...)
	s.m.Store(pid, p)

	go func() {
		p.Start()
		info := p.GetInfo()

		r := result{
			pid:  pid,
			info: info,
		}
		s.ch <- r
	}()

	return pid
}

func (s *scheduler) CancelProcess(pid string) error {
	if v, ok := s.m.Load(pid); ok {
		defer s.m.Delete(pid)
		p := v.(process.Process)
		return p.Stop()
	} else {
		return errors.New("process ID not found")
	}
}

func (s *scheduler) ProcessStatus(pid string) (*process.ProcessInfo, error) {
	if v, ok := s.m.Load(pid); ok {
		p := v.(process.Process)
		return p.GetInfo(), nil
	} else {
		return nil, errors.New("process ID not found")
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
