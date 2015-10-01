package working_queue

import (
	"testing"
	"fmt"
	"time"
)

func Test_NewWorker_should_success(t *testing.T) {
	workerqueue := make(chan chan WorkRequest)
	worker := NewWorker(1, workerqueue)
	if worker.Id != 1 {
		t.Error("NewWorker failed!")
	}
}

func Test_WorkerStart_should_success(t *testing.T) {
	workerqueue := make(chan chan WorkRequest)
	worker := NewWorker(1, workerqueue)
	if worker.Id != 1 {
		t.Error("NewWorker failed!")
	}

	worker.Start()

	worker.Stop()
}

func foo(n interface{}) {
	args := n.([]int)
	total := 0
	for _, v := range(args) {
		total += v
	}
	fmt.Println("Sum: ", total)
}

func Test_Smoke_test(t *testing.T) {
	workerqueue := make(chan chan WorkRequest, 1)
	worker := NewWorker(1, workerqueue)
	worker.Start()

	worker.Request <-WorkRequest{Args:[]int{1, 2, 3, 4}, Handle:foo}
	time.Sleep(1 * time.Second)
	worker.Stop()
}

func Test_NewDispatcher_should_success(t *testing.T) {
	d := NewDispatcher(1)
	if d.WorkerNum != 1 {
		t.Error("NewWorker failed!")
	}
}

func Test_DispatcherStartStop_should_success(t *testing.T) {
	d := NewDispatcher(1)
	if d.WorkerNum != 1 {
		t.Error("NewWorker failed!")
	}
	req := make(chan WorkRequest)
	d.Start(req)
	time.Sleep(1 * time.Second)
	d.Stop()
	fmt.Printf("-------------------------\n")
}

func Test_Dispatcher_Worker_should_success(t *testing.T) {
	d := NewDispatcher(1)
	if d.WorkerNum != 1 {
		t.Error("NewWorker failed!")
	}
	req := make(chan WorkRequest)
	d.Start(req)
	req <- WorkRequest{Args:[]int{3, 3, 3, 3}, Handle:foo}
	req <- WorkRequest{Args:[]int{2, 3, 3, 3}, Handle:foo}
	fmt.Printf("-------------------------\n")
	d.Stop()
	time.Sleep(2 * time.Second)
}
