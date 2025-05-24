/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sortedArrayToBST(nums []int) *TreeNode {
    if len(nums) == 0 {
        return nil
    }

    return createBST(nums, 0, len(nums)-1)
}

func createBST(nums []int, left, right int) *TreeNode {
    if left > right {
        return nil
    }

    mid := left + (right - left) / 2

    root := &TreeNode{Val: nums[mid]}
    root.Left = createBST(nums, left, mid-1)
    root.Right = createBST(nums, mid+1, right)

    return root
}