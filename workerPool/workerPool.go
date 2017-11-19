package workerPool

type Job interface{}
type Handler func(Job)

type WorkerPool struct {
	queue      chan Job
	pool       chan chan Job
	exit       chan struct{}
	numWorkers int
}

func New(numWorkers, queueSize int) *WorkerPool {
	return &WorkerPool{
		queue:      make(chan Job, queueSize),
		pool:       make(chan chan Job),
		exit:       make(chan struct{}),
		numWorkers: numWorkers,
	}
}

func (w *WorkerPool) AddJob(job Job) {
	w.queue <- job
}

func (w *WorkerPool) Stop() {
	signal := struct{}{}

	for i := 0; i < w.numWorkers; i++ {
		w.exit <- signal
	}
}

func (w *WorkerPool) Handle(handler Handler) {
	for i := 0; i < w.numWorkers; i++ {
		go w.startWorker(handler)
	}

	for job := range w.queue {
		go func(job Job) {
			worker := <-w.pool
			worker <- job
		}(job)
	}
}

func (w *WorkerPool) startWorker(handler Handler) {
	ch := make(chan Job)

	for {
		w.pool <- ch

		select {
		case job := <-ch:
			handler(job)

		case <-w.exit:
			return
		}

	}
}
