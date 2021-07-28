package dispatcher

import (
	"github.com/agbankar/go-worker-pool/worker"
	"math/rand"
	"sync"
)

func NewDispatcher(noOWorkers int) *disp {
	return &disp{
		Workers:  make([]*worker.WorkerImpl, noOWorkers),
		WorkChan: make(worker.JobChannel),
		Queue:    make(worker.JobQueue, noOWorkers),
		Quit:     make(chan bool),
		NoOfWorkers: noOWorkers,
	}
}

// disp is the link between the client and the workers
type disp struct {
	NoOfWorkers int
	Workers     []*worker.WorkerImpl // this is the list of workers that dispatcher tracks
	WorkChan    worker.JobChannel    // client submits job to this channel
	Queue       worker.JobQueue      // this is the shared JobPool between the workers
	Quit        chan bool
}

// Start creates pool of num count of workers.
func (d *disp) Start() {
	var wg sync.WaitGroup
	wg.Add(d.NoOfWorkers)
	d.Workers = make([]*worker.WorkerImpl, 0)
	for i := 1; i <= d.NoOfWorkers; i++ {
		w := worker.NewWorker(i, make(worker.JobChannel), d.Queue, make(chan bool))
		d.Workers = append(d.Workers, w)
		go d.runWorker(&wg, w)
	}
	d.FetchAndProcessJobs()
	for _, w := range d.Workers {
		w.Stop()
	}
	return
}

func (d *disp) FetchAndProcessJobs() {
	for {
		select {
		case <-d.Quit:
			return
		default:
			jobChan := <-d.Queue
			jobChan <- getJobToProcess()

		}
	}
}

func (d *disp) runWorker(wg *sync.WaitGroup, w worker.Worker) {
	w.Start()
	wg.Wait()

}
func getJobToProcess() worker.Job {
	job := worker.Job{JobID: rand.Int()}
	return job

}
func (d *disp) Stop() {
	d.Quit <- true

}
