package main

import (
	"errors"
	"fmt"
)

type Queue[T any] struct {
	queue []T
}

// NewQueue creates a new Queue with the provided capacity.
func NewQueue[T any]() *Queue[T] {
	queue := new(Queue[T])
	queue.queue = make([]T, 0)
	return queue
}

// Enqueue inserts the element at the front of the queue.
func (q *Queue[T]) Enqueue(data ...T) {
	q.queue = append(q.queue, data...)
}

// Dequeue remove from the Queue
func (q *Queue[T]) Dequeue() (res T, err error) {
	if len(q.queue) == 0 {
		return res, errors.New("Empty Queue")
	}

	res = q.queue[0]
	q.queue = q.queue[1:]
	return res, nil
}

func main() {
	queue := NewQueue[int]()
	queue.Enqueue(3, 2, 3, 45, 6, 7)
	queue.Dequeue()
	fmt.Println(queue)
}
