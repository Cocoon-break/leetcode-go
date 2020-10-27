package linked_list

import "testing"

// go test -v . -test.run TestDetectCycle
func TestDetectCycle(t *testing.T) {
	a := ListNode{Val: 3, Next: nil}
	b := ListNode{Val: 2, Next: nil}
	c := ListNode{Val: 0, Next: nil}
	d := ListNode{Val: 4, Next: nil}
	a.Next = &b
	b.Next = &c
	c.Next = &d
	d.Next = &b

	out := detectCycle(&a)
	t.Log(out)
}
