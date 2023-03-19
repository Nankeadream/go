package main

import "fmt"

type Node struct {
	data interface{}
	next *Node
}

type LinkedList struct {
	head *Node
	size int
}

func (ll *LinkedList) Add(data interface{}) {
	newNode := &Node{data: data}

	if ll.head == nil {
		ll.head = newNode
	} else {
		curr := ll.head
		for curr.next != nil {
			curr = curr.next
		}
		curr.next = newNode
	}

	ll.size++
}
func (ll *LinkedList) Delete(data int) {
	if ll.head == nil {
		return
	}
	if ll.head.data == data {
		ll.head = ll.head.next
		return
	}
	curr := ll.head
	for curr.next != nil {
		if curr.next.data == data {
			curr.next = curr.next.next
			return
		}
		curr = curr.next
	}
}
func (ll *LinkedList) Get(index int) interface{} {
	if index < 0 || index >= ll.size {
		return nil
	}

	curr := ll.head
	for i := 0; i < index; i++ {
		curr = curr.next
	}

	return curr.data
}

func (ll *LinkedList) Len() int {
	return ll.size
}
func (ll *LinkedList) Println() {
	current := ll.head
	for current != nil {
		fmt.Print(current.data, " ")
		current = current.next
	}
	fmt.Println()
}
