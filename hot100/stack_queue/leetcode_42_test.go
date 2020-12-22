package stack_queue

import "testing"

// go test -v . -test.run TestTrap
func TestTrap(t *testing.T) {
	height := []int{0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1}
	result := trap2(height)
	t.Log(result)
}
