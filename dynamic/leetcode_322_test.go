package dynamic

import "testing"

// go test -v . -test.run TestCoinChange
func TestCoinChange(t *testing.T) {
	result := coinChange([]int{1, 2, 5}, 11)
	t.Log(result)
}

// go test -v . -test.run TestCoinChange2
func TestCoinChange2(t *testing.T) {
	result := coinChange2([]int{1, 2, 5}, 11)
	t.Log(result)
}

// go test -v . -test.run TestCoinChange3
func TestCoinChange3(t *testing.T) {
	result := coinChange3([]int{1, 2, 5}, 11)
	t.Log(result)
}
