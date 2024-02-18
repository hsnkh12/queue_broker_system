package tests

import (
	"queue_system/queue"
	"testing"
)

func TestQueueEnqueue(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", q.Size())
	}
}

func TestQueueDequeue(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Dequeue()

	if item != 1 {
		t.Errorf("Expected dequeued item to be 1, but got %v", item)
	}

	if q.Size() != 2 {
		t.Errorf("Expected queue size to be 2, but got %d", q.Size())
	}
}

func TestQueuePeek(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	item := q.Peek()

	if item != 1 {
		t.Errorf("Expected peeked item to be 1, but got %v", item)
	}

	if q.Size() != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", q.Size())
	}
}

func TestQueueSize(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	size := q.Size()

	if size != 3 {
		t.Errorf("Expected queue size to be 3, but got %d", size)
	}
}

func TestQueueClear(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	q.Clear()

	if q.Size() != 0 {
		t.Errorf("Expected queue size to be 0 after clearing, but got %d", q.Size())
	}
}

func TestQueueContains(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	if !q.Contains(2) {
		t.Errorf("Expected queue to contain 2, but it does not")
	}

	if q.Contains(4) {
		t.Errorf("Expected queue to not contain 4, but it does")
	}
}

func TestQueueToArray(t *testing.T) {
	q := queue.New()

	q.Enqueue(1)
	q.Enqueue(2)
	q.Enqueue(3)

	arr := q.ToArray()

	if len(arr) != 3 {
		t.Errorf("Expected array length to be 3, but got %d", len(arr))
	}

	if arr[0] != 1 || arr[1] != 2 || arr[2] != 3 {
		t.Errorf("Expected array elements to be [1, 2, 3], but got %v", arr)
	}
}

func TestQueueIsEmpty(t *testing.T) {
	q := queue.New()

	if !q.IsEmpty() {
		t.Errorf("Expected queue to be empty, but it is not")
	}

	q.Enqueue(1)

	if q.IsEmpty() {
		t.Errorf("Expected queue to not be empty, but it is")
	}
}
