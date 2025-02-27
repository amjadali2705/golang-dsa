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

	//merge sorted linked list
	l1 := &Node{Val: 1, Next: &Node{Val: 4, Next: &Node{Val: 6}}}
	l2 := &Node{Val: 2, Next: &Node{Val: 3, Next: &Node{Val: 5, Next: &Node{Val: 7}}}}
	fmt.Println(Serialize(JoinTwoSortedLinkedLists(l1, l2)))

	//keep repetitions
	head := &Node{Val: 1, Next: &Node{Val: 1, Next: &Node{Val: 4, Next: &Node{Val: 4, Next: &Node{Val: 6, Next: &Node{Val: 6, Next: &Node{Val: 7}}}}}}}
	fmt.Println(Serialize(KeepRepetitions(head)))
}
