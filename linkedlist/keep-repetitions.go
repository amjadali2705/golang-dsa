package main

func KeepRepetitions(head *Node) *Node {
	if head == nil || head.Next == nil {
		return nil
	}

	temp := &Node{}
	curr := head
	newTail := temp
	isDuplicate := false

	for curr != nil {
		if curr.Next != nil && curr.Val == curr.Next.Val {
			isDuplicate = true
		} else {
			if isDuplicate {
				newTail.Next = &Node{Val: curr.Val}
				newTail = newTail.Next
			}
			isDuplicate = false
		}
		curr = curr.Next
	}
	return temp.Next
}
