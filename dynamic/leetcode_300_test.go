package dynamic

import "testing"

// go test -v . -test.run TestLengthOfLIS
func TestLengthOfLIS(t *testing.T) {
	result := lengthOfLIS([]int{10, 9, 2, 5, 3, 7, 101, 18})
	result2 := lengthOfLIS2([]int{10, 9, 2, 5, 3, 7, 101, 18})
	t.Log(result)
	t.Log(result2)
}
