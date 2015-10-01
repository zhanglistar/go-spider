package deque

import (
	"testing"
)

func Test_NewDeque_should_success(t *testing.T) {
	queue := NewDeque()
	if queue.capcity == -1 {
		t.Log("NewDeque success!")
	}
	if queue.Len() != 0 {
		t.Error("Len error!")
	}
}

func Test_Push_should_success(t *testing.T) {
	queue := NewDeque()
	queue.Push("one")
	value := queue.Pop()
	if value != "one" {
		t.Log("Push failed")
	} else {
		t.Log("Push success")
	}

	if queue.Len() != 0 {
		t.Error("Len error!")
	}
}

func Test_Push_should_fail_when_full(t *testing.T) {
	queue := NewCappedDeque(2)
	retItem := queue.Push("one")
	if !retItem {
		t.Error("Push fail")
	}
	if queue.Len() != 1 {
		t.Error("Len error!")
	}
	retItem = queue.Push("two")
	if !retItem {
		t.Error("Push fail")
	}
	if queue.Len() != 2 {
		t.Error("Len error!")
	}
	retItem = queue.Push(3)
	if retItem {
		t.Error("Push should fail")
	}
	if queue.Len() != 2 {
		t.Error("Len error!")
	}
}

func Test_Pop_should_fail_when_empty(t *testing.T) {
	var queue *Deque = NewDeque()
	retItem := queue.Pop()
	if retItem != nil {
		t.Error("Test_Pop_should_fail_when_empty failed")
	}
	if queue.Len() != 0 {
		t.Error("Len error!")
	}
}

func Test_Pop_should_success_when_not_empty(t *testing.T) {
	queue := NewDeque()
	retItem := queue.Push(1)
	if retItem == false {
		t.Error("Test_Pop_should_fail_when_empty failed")
	}
	retItem1 := queue.Pop()
	if retItem1 != 1 {
		t.Error("Test_Pop_should_success_when_not_empty failed")
	}
	if queue.Len() != 0 {
		t.Error("Len error!")
	}
}

func Benchmark_Push(b *testing.B) {
	b.StopTimer()
	queue := NewDeque()

	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
}

func Benchmark_Pop(b *testing.B) {
	b.StopTimer()
	queue := NewDeque()

	for i := 0; i < b.N; i++ {
		queue.Push(i)
	}
	b.StartTimer()
	for i := 0; i < b.N; i++ {
		queue.Pop()
	}
}
