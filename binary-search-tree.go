// /*
// - top node of binary tree      -- root
// - nodes below root             -- parents
// - nodes that are below parents -- children
// - nodes without children       -- leaves
// - a tree that has node with no more than 2 children -- binary tree
// - the smaller value is the left child
// - the larger value is the right child
// - nodes are always inserted as a leaf
// - advantage of binary tree is speed
// O(H)
// */

package main

import "fmt"

var count int

type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

// Insert will add a node to the tree
func (node *Node) Insert(key int) {
	if node.Key < key {
		if node.Right == nil {
			node.Right = &Node{Key: key}
		} else {
			node.Right.Insert(key)
		}
	} else if node.Key > key {
		if node.Left == nil {
			node.Left = &Node{Key: key}
		} else {
			node.Left.Insert(key)
		}
	}
}

// Search will take a Key value and return true if there is a node with that value
// takes a pointer receiver

func (node *Node) Search(key int) bool {
	count++

	if node == nil {
		return false
	}

	if node.Key < key {
		return node.Right.Search(key)
	} else if node.Key > key {
		return node.Left.Search(key)
	}

	return true
}

// tree should be an address -- we'll call this as a pointer
func main() {
	tree := &Node{Key: 100}
	fmt.Println(tree) // &{100 <nil> <nil>}

	tree.Insert(50)
	fmt.Println(tree) // &{100 0xc00011c048 <nil>}

	tree.Insert(200)
	fmt.Println(tree) // &{100 0xc00000c060 0xc00000c090}

	tree.Insert(300)
	fmt.Println(tree) // &{100 0xc00000c060 0xc00000c090}

	fmt.Println(tree.Search(50))  // true
	fmt.Println(tree.Search(400)) // false
	fmt.Println(count)
}

// **WITHOUT POINTERS, WE ONLY PASS BY VALUE, AND CHANGES WITHIN OUR METHODS WILL NOT REFLECT IN THE ORIGINAL TYPES

// package main

// import "fmt"

// type Node struct {
// 	Key   int
// 	Left  *Node
// 	Right *Node
// }

// // Insert will add a node to the tree
// func (node Node) Insert(key int) { // CHANGED FROM POINTER TO VALUE (WE ARE NO LONGER PASSING A POINTER)
// 	if node.Key < key {
// 		if node.Right == nil {
// 			node.Right = &Node{Key: key}
// 		} else {
// 			node.Right.Insert(key)
// 		}
// 	} else if node.Key > key {
// 		if node.Left == nil {
// 			node.Left = &Node{Key: key}
// 		} else {
// 			node.Le    // go right
//
// func main() {
// 	tree := Node{Key: 100} // CHANGED FROM PASSING POINTER TO PASSING VALUE (OR COPY) OF TYPES
// 	fmt.Println(tree)

// 	tree.Insert(50)
// 	fmt.Println(tree)

// 	tree.Insert(200)
// 	fmt.Println(tree)

// 	tree.Insert(300)
// 	fmt.Println(tree)
// }

// // NEW OUTPUTS
// // {100 <nil> <nil>}
// // {100 <nil> <nil>}
// // {100 <nil> <nil>}
// // {100 <nil> <nil>}
