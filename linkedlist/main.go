package main

import "fmt"

func main() {
	//serialization
	node1 := &Node{Val: 1}
	node2 := &Node{Val: 2}
	node3 := &Node{Val: 3}
	node1.Next = node2
	node2.Next = node3

	fmt.Println(Serialize(node1))

	//reverse linked list
	fmt.Println(Serialize(ReverseLinkedList(node1)))
}
