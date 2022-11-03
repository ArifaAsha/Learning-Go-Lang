package main

import "fmt"

const MAX = 100

type Queue struct {
	front       int
	actualQueue []int
}

func (s Queue) IsEmpty() bool {
	return s.front < 0
}

func (q Queue) enqueue(vals ...int) (Queue, []int) {
	for _, val := range vals {
		q.front += 1
		q.actualQueue = append(q.actualQueue, val)
	}
	return q, q.actualQueue
}

// func dequeue(queue []int) []int {
// 	element := queue[0] // The first element is the one to be dequeued.
// 	fmt.Println("Dequeued:", element)
// 	return queue[1:] // Slice off the element once it is dequeued.
// }

// FIFO -> first element of slice -> popped
func (q Queue) Dequeue() (Queue, []int) {
	element := q.actualQueue[0] // The first element is the one to be dequeued.
	fmt.Println("Dequeued:", element)
	q.front--
	return q, q.actualQueue[1:] // Slice off the element once it is dequeued.
}

func (q Queue) Front() int {
	if !q.IsEmpty() {
		val := q.actualQueue[q.front]
		return val
	}
	return 0
}

func main() {
	q := Queue{front: -1}
	newQueue, actualQueue := q.enqueue(8, 10)
	fmt.Println(actualQueue)

	newQueue2, actualQueue2 := newQueue.enqueue(90)
	fmt.Println(newQueue2, actualQueue2)

	fmt.Println("After dequeue 8: ", newQueue2.actualQueue)
	afterDequeue, dequedElement := newQueue2.Dequeue()
	fmt.Println(afterDequeue, dequedElement)

	fmt.Println("Front of the queue: ", afterDequeue.Front())
}
