package main

import "fmt"

// Define Queue Struct that holds a slice
type Queue struct {
	items []int
}

// Enqueue -- adds a value at the end
func (queue *Queue) Enqueue(i int) {
	queue.items = append(queue.items, i)
}

// Dequeue -- Removes value from front and returns the removed value

func (queue *Queue) Dequeue() int {
	toRemove := queue.items[0]
	queue.items = queue.items[1:]
	return toRemove
}

func main() {
	myQueue := Queue{}
	fmt.Println(myQueue)
	myQueue.Enqueue(100)
	myQueue.Enqueue(200)
	myQueue.Enqueue(300)
	myQueue.Enqueue(400)
	fmt.Println(myQueue)

	myQueue.Dequeue()
	fmt.Println(myQueue)
}
