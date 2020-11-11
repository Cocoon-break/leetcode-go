package stack_queue

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归方式
func inorderTraversalByRecursion(root *TreeNode) []int {
	result := []int{}
	if root != nil {
		result1 := inorderTraversalByRecursion(root.Left)
		result = append(result, result1...)
		result = append(result, root.Val)
		result2 := inorderTraversalByRecursion(root.Right)
		result = append(result, result2...)
	}
	return result
}

// 非递归方式
// 思路：利用栈的方式处理
// 1. 延着左节点持续遍历，压栈
// 2. 出栈，存值，然后判断是否存在右节点，存在则继续从1开始
func inorderTraversal(root *TreeNode) []int {
	firstNode := root
	result := []int{}
	var nodes []*TreeNode
	for firstNode != nil {
		//1. 延着左节点持续遍历，压栈
		for firstNode != nil {
			nodes = append(nodes, firstNode)
			firstNode = firstNode.Left
		}
		for len(nodes) > 0 && nodes[len(nodes)-1] != nil {
			//出栈，为左分支的最后一个节点
			node := nodes[len(nodes)-1]
			nodes = nodes[:len(nodes)-1]
			result = append(result, node.Val)
			//查看出栈的节点是否有右节点，有的话从1开始
			if node.Right != nil {
				firstNode = node.Right
				break
			}
		}
		// if len(nodes) > 0 {
		// 	//出栈，为左分支的最后一个节点
		// 	node := nodes[len(nodes)-1]
		// 	nodes = nodes[:len(nodes)-1]
		// 	result = append(result, node.Val)
		// 	//查看出栈的节点是否有右节点
		// 	if node.Right != nil {
		// 		firstNode = node.Right
		// 		continue
		// 	}
		// 	//因当前节点无右节点，持续出栈
		// 	for node.Right == nil && len(nodes) > 0 {
		// 		node := nodes[len(nodes)-1]
		// 		nodes = nodes[:len(nodes)-1]
		// 		result = append(result, node.Val)
		// 		//遇到栈顶右节点不为空，结束出栈，则继续从1开始
		// 		if node.Right != nil {
		// 			firstNode = node.Right
		// 			break
		// 		}
		// 	}
		// }
	}
	return result
}
