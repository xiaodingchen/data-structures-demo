package queue

import (
	"testing"
)

func TestQueueLink(t *testing.T) {
	var queue QueueLink
	queue.InitQueue()
	queue.EnQueue(11)
	queue.EnQueue(23)
	queue.EnQueue(12)

	t.Logf("queue length:%d", queue.Length())
	var e Elem
	l := queue.Length()
	for i := 0; i < l-1; i++ {
		queue.OutQueue(&e)
		t.Logf("i: %d", e)
	}

	queue.EnQueue(34)
	queue.EnQueue(21)
	queue.EnQueue(43)
	l = queue.Length()
	for i := 0; i < l; i++ {
		queue.OutQueue(&e)
		t.Logf("i: %d", e)
	}
}
