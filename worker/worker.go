package worker

import (
	"context"
)

type Job interface {
	Do() Result
}

type Result interface{}

func Work(workerNum uint, jobList []Job) (<-chan Result, context.CancelFunc) {
	jobs := make(chan Job)
	results := make(chan Result)
	done := make(chan struct{}, workerNum)
	ctx, cancel := context.WithCancel(context.Background())

	go dispatchJobs(ctx, jobs, jobList)
	for i := uint(0); i < workerNum; i++ {
		go doJobs(jobs, results, done)
	}
	go waitDone(workerNum, done, results)

	return results, cancel
}

func waitDone(workerNum uint, done chan struct{}, results chan Result) {
	for i := uint(0); i < workerNum; i++ {
		<-done
	}

	close(results)
}

func doJobs(jobs <-chan Job, results chan<- Result, done chan<- struct{}) {
	for job := range jobs {
		results <- job.Do()
	}

	done <- struct{}{}
}

func dispatchJobs(ctx context.Context, jobs chan<- Job, jobList []Job) {
	defer close(jobs)

	for _, job := range jobList {
		select {
		case <-ctx.Done():
			return
		default:
			jobs <- job
		}
	}
}
