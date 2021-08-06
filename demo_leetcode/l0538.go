package demo_leetcode

func convertBST(root *TreeNode) *TreeNode {
	recursive(root, 0)
	return root
}

// recursive i 为累加到现在的值，初始值为0
func recursive(root *TreeNode, i int) int {
	if root == nil {
		return i
	}
	i = recursive(root.Right, i)
	root.Val = root.Val + i
	i = root.Val
	i = recursive(root.Left, i)
	return i
}
