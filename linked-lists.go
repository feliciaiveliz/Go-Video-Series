/*
Junmin Lee series: https://youtu.be/8QoynPUY9_8

Linked List
linked lists are sequences of nodes
nodes are linked with each other
each node contains an address to tell us where the next node is
start at head to find a certain node in linked list
O(N)
adding and deleting a node from the beginning in a linked list is easy
O(1)

Doubly Linked List
- each node contains an address of the next and previous nodes
- adding a node at the end of a doubly linked list would be easier because we can access a node from
- any place because they contain a previous node's address
- (list *linkedList) is a pointer receiver so that we can work on the actual node in memory -- not a copy
- (node *node) is accepting a node value with a type 'node' (meaning we have to pass a node created from the node struct)
- func (list *linkedList) prepend (node *node) { } is how we write a method of a struct
- (list linkedList) is a value receiver because all we want to do is work on a copy and just print out values -- not modify anything
*/

package main

import "fmt"

type node struct {
	data int
	next *node // pointer of a node
}

type linkedList struct {
	head   *node // pointer of a node
	length int
}

func (list *linkedList) prepend(node *node) {
	// make a temp 'second' that will hold the old or current head node
	second := list.head
	// set new node as the new head
	list.head = node
	// let the new head point to the new second
	list.head.next = second
	// increase the length of the list because we added a new node to it at the beginning
	list.length++
}

func (list linkedList) printListData() {
	toPrint := list.head

	for list.length != 0 {
		fmt.Printf("%d ", toPrint.data)
		toPrint = toPrint.next
		list.length--
	}

	fmt.Printf("\n")
}

func (list *linkedList) deleteWithValue(value int) {
	if list.length == 0 {
		return
	}

	if list.head.data == value {
		list.head = list.head.next
		list.length--
		return
	}

	// start with head
	previoustoDelete := list.head

	// while the next value is not equal to the one we want to delete, loop
	for previoustoDelete.next.data != value {
		if previoustoDelete.next.next == nil {
			return
		}
		previoustoDelete = previoustoDelete.next
	}

	// once we've found the node to delete, set the previous's next value to the node after the node we want to delete
	previoustoDelete.next = previoustoDelete.next.next

	// decrease the size of the list because we have deleted one node
	list.length--
}

func main() {
	// create a new list
	myList := linkedList{}

	// create some nodes -- pass a pointer
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}

	// prepend all the nodes to the list
	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)

	fmt.Println(myList) // {0xc000010270 3}
	myList.printListData()

	myList.deleteWithValue(2)
	myList.deleteWithValue(100)
	myList.deleteWithValue(1)
	myList.printListData()
}
