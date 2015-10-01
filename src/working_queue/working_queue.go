package working_queue

import (
//	"fmt"
)

type WorkRequest struct {
	Args   interface{}
	Handle func(args interface{})
}

type Worker struct {
	Id          int
	Request     chan WorkRequest
	WorkerQueue chan chan WorkRequest
	Quit        chan bool
}

func NewWorker(id int, workerQueue chan chan WorkRequest) Worker {
	return Worker {
		Id: id,
		Request: make(chan WorkRequest),
		WorkerQueue: workerQueue,
		Quit: make(chan bool),
	}
}

func (w Worker) Start() {
	go func() {
		for {
			w.WorkerQueue <- w.Request
//			fmt.Println("Waiting requests...")
			select {
			case request := <-w.Request:
				request.Handle(request.Args)
//				fmt.Printf("Worker %d handled request\n", w.Id)
			case <-w.Quit:
//				fmt.Printf("Worker %d quit.\n", w.Id)
				return
			}
		}
	}()
}

func (w Worker) Stop() {
	go func() {
		w.Quit <- true
	}()
}

type Dispatcher struct {
	WorkerNum int
	Workers []Worker
	WorkerQueue chan chan WorkRequest
	Quit chan bool
}

func NewDispatcher(nworker int) Dispatcher {
	return Dispatcher{
		WorkerNum: nworker,
		Workers: make([]Worker, nworker),
		WorkerQueue: make(chan chan WorkRequest, nworker),
		Quit: make(chan bool),
	}
}

func (d Dispatcher) Start(request chan WorkRequest) {
	for i := 0; i < d.WorkerNum; i++ {
//		fmt.Printf("Start worker %d\n", i)
		worker := NewWorker(i, d.WorkerQueue)
		worker.Start()
		d.Workers = append(d.Workers, worker)
	}
	go func() {
		for {
			select {
			case req := <-request:
//				fmt.Printf("Dispatching request...\n")
				go func() {
					worker := <-d.WorkerQueue
					worker <- req
				}()
			case <- d.Quit:
				for _, v := range(d.Workers) {
					v.Stop()
				}
//				fmt.Printf("Dispatcher exit.\n")
				return
			}
		}
	}()
}

func (d Dispatcher) Stop() {
	d.Quit <- true
}
