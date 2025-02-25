package main

func ReverseLinkedList(head *Node) *Node {
	var prev *Node
	curr := head

	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}