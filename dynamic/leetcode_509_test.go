package dynamic

import "testing"

// go test -v . -test.run TestFib
func TestFib(t *testing.T) {
	result := fib(10)
	t.Log(result)

	result2 := fib2(10)
	t.Log(result2)

	result3 := fib3(10)
	t.Log(result3)

	resultK := fibK(10)
	t.Log(resultK)
}
