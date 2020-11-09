package linked_list

import "testing"

// go test -v . -test.run TestAddTwoNumbers
func TestAddTwoNumbers(t *testing.T) {
	l1 := GenListNode([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := GenListNode([]int{9, 9, 9, 9})
	result := addTwoNumbers(l1, l2)
	LogListNode(result)
}
