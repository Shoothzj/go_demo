package demo_leetcode

func deepestLeavesSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	// do it in bfs way
	auxList := make([]*TreeNode, 1)
	auxList[0] = root
	for {
		nextList := make([]*TreeNode, 0)
		for _, node := range auxList {
			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}
			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}
		if len(nextList) == 0 {
			sum := 0
			for _, node := range auxList {
				sum += node.Val
			}
			return sum
		}
		auxList = nextList
	}
}
