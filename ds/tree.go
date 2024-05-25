package ds

type TreeNode[T comparable] struct {
	val T
	childrenList []*TreeNode[T]
}

type Tree[T comparable] struct {
	root *TreeNode[T]
}
