package linked_list

import "testing"

// go test -v . -test.run TestReverseList
func TestReverseList(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	reversed := reverseList(head)
	LogListNode(reversed)
}
