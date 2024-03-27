package piscine

func BTreeSearchItem(root *TreeNode, elem string) *TreeNode {
	if root == nil || root.Data == elem {
		return root
	}
	if r := BTreeSearchItem(root.Left, elem); r != nil {
		return r
	}
	return BTreeSearchItem(root.Right, elem)
}
