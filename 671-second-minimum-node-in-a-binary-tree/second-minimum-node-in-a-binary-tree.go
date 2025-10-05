/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func findSecondMinimumValue(root *TreeNode) int {
    if root == nil {
        return -1
    }
    
    min1 := root.Val
    
    return findSecondMinInSubtree(root, min1)
}

func findSecondMinInSubtree(node *TreeNode, min1 int) int {
    if node == nil {
        return -1
    }
    
    if node.Val > min1 {
        return node.Val
    }
    
    leftSecondMin := findSecondMinInSubtree(node.Left, min1)
    rightSecondMin := findSecondMinInSubtree(node.Right, min1)
    
    if leftSecondMin == -1 && rightSecondMin == -1 {
        return -1
    }
    
    if leftSecondMin == -1 {
        return rightSecondMin
    }

    if rightSecondMin == -1 {
        return leftSecondMin
    }
    
    return int(math.Min(float64(leftSecondMin), float64(rightSecondMin)))
}