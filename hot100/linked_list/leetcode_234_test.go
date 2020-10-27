package linked_list

import "testing"

func TestIsPalindrome(t *testing.T) {
	head := GenListNode([]int{1, 2, 2, 1})
	out := isPalindrome(head)
	t.Log(out)
}
