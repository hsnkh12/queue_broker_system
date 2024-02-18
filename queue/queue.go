package queue

import "fmt"

type QueueI interface {
	Enqueue(data interface{})
	Dequeue() interface{}
	Peek() interface{}
	Size() int
	Clear()
	Contains(data interface{}) bool
	ToArray() []interface{}
	Print()
	IsEmpty() bool
}

type QNode struct {
	data interface{}
	next *QNode
}

type Queue struct {
	head *QNode
}

func New() *Queue {
	return &Queue{}
}

func (q *Queue) Enqueue(data interface{}) {

	newNode := &QNode{data: data, next: nil}

	if q.IsEmpty() {
		q.head = newNode
		return
	}

	last := q.head
	for last.next != nil {
		last = last.next
	}

	last.next = newNode
}

func (q *Queue) Dequeue() interface{} {

	if q.IsEmpty() {
		return nil
	}

	dequeuedNode := q.head

	q.head = q.head.next

	dequeuedNode.next = nil

	return dequeuedNode.data

}

func (q *Queue) Peek() interface{} {
	return q.head.data
}

func (q *Queue) Size() int {
	size := 0
	current := q.head
	for current != nil {
		size++
		current = current.next
	}
	return size
}

func (q *Queue) Clear() {
	if q.IsEmpty() {
		return
	}
	current := q.head
	for current != nil {
		temp := current.next
		current.next = nil
		current = temp
	}
	q.head = nil
}

func (q *Queue) Contains(data interface{}) bool {
	current := q.head
	for current != nil {
		if current.data == data {
			return true
		}
		current = current.next
	}
	return false
}

func (q *Queue) ToArray() []interface{} {
	arr := make([]interface{}, 0)
	current := q.head
	for current != nil {
		arr = append(arr, current.data)
		current = current.next
	}
	return arr
}

func (q *Queue) Print() {
	current := q.head
	for current != nil {
		fmt.Println(current.data)
		current = current.next
	}
}

func (q *Queue) IsEmpty() bool {
	return q.head == nil
}
