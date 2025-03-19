package main

import "fmt"

type node struct {
	data int
	next *node
}

type LinkedList struct {
	head   *node
	length int
}

// Prepend method to add a new node at the beginning
func (l *LinkedList) prepend(n *node) {
	n.next = l.head // Point new node to the old head
	l.head = n      // Move head to new node
	l.length++
}

// Print the linked list
func (l *LinkedList) printListData() {
	toPrint := l.head
	for toPrint != nil { // Correctly traverse the list without modifying `length`
		fmt.Printf("%d -> ", toPrint.data)
		toPrint = toPrint.next
	}
	fmt.Println("nil") // Indicate the end of the list
}

func main() {
	fmt.Println("Hello, World")

	myList := LinkedList{}
	node1 := &node{data: 48}
	node2 := &node{data: 58}
	node3 := &node{data: 68}
	node4 := &node{data: 88}

	myList.prepend(node1)
	myList.prepend(node2)
	myList.prepend(node3)
	myList.prepend(node4)

	fmt.Println("Linked List:")
	myList.printListData()
}
