package piscine

func BTreeLevelCount(root *TreeNode) int {
	if root == nil {
		return 0
	}
	r := BTreeLevelCount(root.Right)
	l := BTreeLevelCount(root.Left)

	if l > r {
		return l + 1
	}
	return r + 1
}
