package main

import "fmt"

type Queue struct {
	items []int
}

func (q *Queue) Enqueue(value int) {
	q.items = append(q.items, value)
}

func (q *Queue) Dequeue() int {
	remove := q.items[0]
	q.items = q.items[1:]
	return remove
}

func main() {
	implementQueue := Queue{}
	implementQueue.Enqueue(1)
	implementQueue.Enqueue(2)
	implementQueue.Enqueue(3)
	fmt.Println(implementQueue.items)
	implementQueue.Dequeue()
	implementQueue.Dequeue()
	implementQueue.Dequeue()
	fmt.Println(implementQueue.items)
}