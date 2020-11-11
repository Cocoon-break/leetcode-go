package stack_queue

import "testing"

// go test -v . -test.run TestIsValid
func TestIsValid(t *testing.T) {
	t.Log(isValid("()[[]]{}"))
}
