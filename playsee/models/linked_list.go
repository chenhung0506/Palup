package models

import "fmt"

type Node struct {
	Value interface{}
	Next  *Node
}

func CreateLinkedList(array []interface{}) *Node {
	var head *Node
	var current *Node

	for _, value := range array {
		newNode := &Node{Value: value}
		if head == nil {
			head = newNode
		} else {
			current.Next = newNode
		}
		current = newNode
	}
	return head
}

func PrintLinkedList(head *Node) {
	current := head
	for current != nil {
		fmt.Printf("%v -> ", current.Value)
		current = current.Next
	}
	fmt.Println("nil")
}
