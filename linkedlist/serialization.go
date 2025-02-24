package main

import (
	"strconv"
	"strings"
)

type Node struct {
	Val  int
	Next *Node
}

func Serialize(node *Node) string {
	if node == nil {
		return ""
	}

	res := strconv.Itoa(node.Val) + "->"

	for node.Next != nil {
		node = node.Next
		res += strconv.Itoa(node.Val) + "->"
	}
	return strings.Trim(res, "->")
}
