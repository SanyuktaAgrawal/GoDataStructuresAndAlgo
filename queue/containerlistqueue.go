package main

import (
	"container/list"
	"fmt"
)

func main() {
	queue := list.New()

	queue.PushBack(5)
	queue.PushBack(10)

	dequeue := queue.Front()
	fmt.Println(dequeue.Value)
	queue.Remove(dequeue)

}
