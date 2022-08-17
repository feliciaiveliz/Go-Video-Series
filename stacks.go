/*
basic linear data structures
size is flexible - grow and shrink
stacks - LIFO
  push to add item to front of stack
  pop to delete item from front of stack
queues - FIFO
  line of supermarket
  first person to join line is first to be served
  enqueue to add item
  dequeued to delete item
*/

package main

import "fmt"

// create Stack struct with 'items' slice of ints
type Stack struct {
	items []int
}

// Push item onto stack
// Pointer receiver -- means we will be calling the Push method on a Stack type (and be modifying it)
func (stack *Stack) Push(i int) {
	// append integer to items
	stack.items = append(stack.items, i)
}

// Pop -- return int that was removed
func (stack *Stack) Pop() int {
	length := len(stack.items) - 1
	toRemove := stack.items[length]

	stack.items = stack.items[:length]
	return toRemove
}

func main() {
	myStack := Stack{}
	fmt.Println(myStack)
	myStack.Push(100)
	myStack.Push(200)
	myStack.Push(300)
	fmt.Println(myStack)

	// pop
	myStack.Pop()
	fmt.Println(myStack)
}
