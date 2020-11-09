package linked_list

import "testing"

// go test -v . -test.run TestRemoveNthFromEnd
func TestRemoveNthFromEnd(t *testing.T) {
	head := GenListNode([]int{1, 2, 3, 4, 5})
	result := removeNthFromEnd(head, 2)
	LogListNode(result)
}
