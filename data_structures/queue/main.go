package main

import "fmt"

type Queue struct {
	items []int
	head  int
	tail  int
}

func (q *Queue) Enqueue(item int) {
	q.items = append(q.items, item)
	q.tail++
}
func (q *Queue) Dequeue() int {
	if q.head == q.tail {
		return -1 // queue is empty
	}
	item := q.items[q.head]
	q.head++
	return item
}
func main() {
	queue := Queue{}
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	fmt.Println(queue.Dequeue()) // Output: 1
	fmt.Println(queue.Dequeue()) // Output: 2
	fmt.Println(queue.Dequeue()) // Output: 3
	fmt.Println(queue.Dequeue()) // Output: -1 (queue is empty)
}
