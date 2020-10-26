package linked_list

import "testing"

// go test -v . -test.run TestGenNodes
func TestGenNodes(t *testing.T) {
	values := []int{1, 2, 3, 4, 5, 6}
	out := GenListNode(values)
	LogListNode(out)
}
