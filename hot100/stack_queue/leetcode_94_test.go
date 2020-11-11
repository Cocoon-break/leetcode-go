package stack_queue

import (
	"fmt"
	"testing"
)

// go test -v . -test.run  TestInorderTraversal
func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 3}
	right1 := &TreeNode{Val: 2}
	root.Right = right1
	right1.Left = left1
	values2 := inorderTraversal(root)
	fmt.Printf("result2:%+v", values2)
}

// go test -v . -test.run TestInorderTraversalByRecursion
func TestInorderTraversalByRecursion(t *testing.T) {
	root := &TreeNode{Val: 1}
	left1 := &TreeNode{Val: 3}
	right1 := &TreeNode{Val: 2}
	root.Right = right1
	right1.Left = left1
	values2 := inorderTraversalByRecursion(root)
	fmt.Printf("result2:%+v", values2)
}
