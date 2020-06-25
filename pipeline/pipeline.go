package pipeline

import (
	"context"
	"sync"
)

type Task interface {
	Do() Result
}

type Result interface{}

func Pipeline(pipeNum uint, tasks []Task) (<-chan Result, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	inPipe := generate(ctx, tasks...)
	midPipes := make([]<-chan Result, 0, int(pipeNum))
	for i := uint(0); i < pipeNum; i++ {
		midPipes = append(midPipes, runTask(ctx, inPipe))
	}

	return merge(ctx, midPipes...), cancel
}

func merge(ctx context.Context, pipes ...<-chan Result) <-chan Result {
	out := make(chan Result)
	var wg sync.WaitGroup

	collect := func(pipe <-chan Result) {
		defer wg.Done()

		for res := range pipe {
			select {
			case <-ctx.Done():
				return
			default:
				out <- res
			}
		}
	}

	wg.Add(len(pipes))
	for _, pipe := range pipes {
		go collect(pipe)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func runTask(ctx context.Context, inPipe <-chan Task) <-chan Result {
	out := make(chan Result)

	go func() {
		defer close(out)

		for task := range inPipe {
			select {
			case <-ctx.Done():
				return
			default:
				out <- task.Do()
			}
		}
	}()

	return out
}

func generate(ctx context.Context, tasks ...Task) <-chan Task {
	out := make(chan Task)

	go func() {
		defer close(out)

		for _, task := range tasks {
			select {
			case <-ctx.Done():
				return
			default:
				out <- task
			}
		}
	}()

	return out
}
