package piscine

func BTreeApplyByLevel(root *TreeNode, f func(...interface{}) (int, error)) {
	for i := 1; i <= BTreeLevelCount(root); i++ {
		Level(root, i, f)
	}
}

func Level(root *TreeNode, i int, f func(...interface{}) (int, error)) {
	if root == nil {
		return
	}
	if i == 1 {
		f(root.Data)
	} else if i > 1 {
		Level(root.Left, i-1, f)
		Level(root.Right, i-1, f)
	}
}
