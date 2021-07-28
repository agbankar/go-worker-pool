package worker

import (
	"fmt"
	"time"
)

type Job struct {
	JobID int
}

type JobChannel chan Job
type JobQueue chan chan Job

type Worker interface {
	Start()
	Stop()
}
type WorkerImpl struct {
	ID      int
	JobChan JobChannel
	Queue   JobQueue
	Quit    chan bool
}

func NewWorker(ID int, JobChan JobChannel, Queue JobQueue, Quit chan bool) *WorkerImpl {
	return &WorkerImpl{
		ID:      ID,
		JobChan: JobChan,
		Queue:   Queue,
		Quit:    Quit,
	}
}

func (w *WorkerImpl) Start() {
	for {
		w.Queue <- w.JobChan
		select {
		case <-w.Quit:
			return
		case job := <-w.JobChan:
			w.processJob(&job)

		}
	}

}

func (w *WorkerImpl) processJob(job *Job) {
	fmt.Println("worker",w.ID,",is processing the job",  job.JobID)
	time.Sleep(10 * time.Second)
}

func (w *WorkerImpl) Stop() {
	fmt.Println("Shutting down the worker", w.ID)
	w.Quit <- true
}
