package concurrency

import (
	"context"
	"sync"
)

type PipelineTask interface {
	Do() Result
}

type Result interface{}

// tasks写完后一定要关闭，否则死锁
func Pipeline(pipeNum uint, tasks <-chan PipelineTask) (<-chan Result, context.CancelFunc) {
	ctx, cancel := context.WithCancel(context.Background())
	inPipe := generate(ctx, tasks)
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
		wg.Wait()  // 确保所有数据都写完了，避免关闭out后再写入导致panic
		close(out) // 关闭out，避免for-range死锁
	}()

	return out
}

func runTask(ctx context.Context, inPipe <-chan PipelineTask) <-chan Result {
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

func generate(ctx context.Context, tasks <-chan PipelineTask) <-chan PipelineTask {
	out := make(chan PipelineTask)

	go func() {
		defer close(out)

		for task := range tasks {
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
