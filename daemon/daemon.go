package daemon

import (
	"fmt"
	"sync"
)

type Daemon interface {
	Shutdown()
	AddTask(t Task)
}

type Task interface {
	Do()
}

type TaskFunc func()

func (t TaskFunc) Do() {
	t()
}

type daemon struct {
	ch     chan Task
	closed bool

	*sync.Mutex
	*sync.WaitGroup
}

func New(goroutineNum int) (Daemon, error) {
	if goroutineNum <= 0 {
		return nil, fmt.Errorf("invalid goroutine number: %d", goroutineNum)
	}

	daemon := &daemon{
		ch:        make(chan Task),
		closed:    false,
		WaitGroup: &sync.WaitGroup{},
		Mutex:     &sync.Mutex{},
	}

	daemon.Add(goroutineNum)
	for i := 0; i < goroutineNum; i++ {
		go func() {
			defer daemon.Done()

			for task := range daemon.ch {
				task.Do()
			}
		}()
	}

	return daemon, nil
}

func (d *daemon) Shutdown() {
	d.Lock()
	if d.closed {
		return
	}

	d.closed = true
	close(d.ch)
	d.Unlock()

	d.Wait()
}

func (d *daemon) AddTask(t Task) {
	go func() {
		d.Lock()
		defer d.Unlock()

		if d.closed {
			return
		}

		d.ch <- t
	}()
}
