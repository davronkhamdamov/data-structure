package main

import "fmt"

// Queue
type Queue struct {
	items []int
}

// Enqueue
func (q *Queue) Enqueue(i int) {
	q.items = append(q.items, i)
}

// Dequeue

func (q *Queue) Dequeue() int {
	toRemove := q.items[0]
	q.items = q.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	myQueue.Enqueue(12)
	myQueue.Enqueue(22)
	myQueue.Enqueue(25)
	fmt.Println(myQueue)
	myQueue.Dequeue()
	fmt.Println(myQueue)

}
