/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
    if root == nil {
        return 0
    }

    maxdepth := 0
    for _, child := range root.Children {
        depth := maxDepth(child)

        if depth > maxdepth {
            maxdepth = depth
        }
    }

    return 1 + maxdepth
}