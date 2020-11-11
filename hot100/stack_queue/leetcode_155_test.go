package stack_queue

import "testing"

// go test -v . -test.run TestMinStack
func TestMinStack(t *testing.T) {
	minStack := Constructor()
	minStack.Push(-2)
	minStack.Push(0)
	minStack.Push(-3)
	t.Log(minStack.GetMin())
	minStack.Pop()
	t.Log(minStack.Top())
	t.Log(minStack.GetMin())
}
