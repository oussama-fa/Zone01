package piscine

func BTreeMin(root *TreeNode) *TreeNode {
	tmp := root
	if tmp == nil {
		return nil
	}
	for tmp.Left != nil {
		tmp = tmp.Left
	}
	return tmp
}
