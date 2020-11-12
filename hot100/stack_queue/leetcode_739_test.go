package stack_queue

import "testing"

//go test -v . -test.run TestDailyTemperatures
func TestDailyTemperatures(t *testing.T) {
	T := []int{73, 74, 75, 71, 69, 72, 76, 73}
	result := dailyTemperatures2(T)
	t.Log(result)
}
