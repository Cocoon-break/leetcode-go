package stack_queue

import "testing"

// go test -v . -test.run TestLeastInterval
func TestLeastInterval(t *testing.T) {
	taks := []byte{'A', 'A', 'A', 'C', 'B', 'B', 'B', 'C'}
	result := leastInterval(taks, 2)
	t.Log(result)
}

/*
tasks = ["A","A","A","B","B","B"], n = 2
转换成图,输出应该为 8
A,B,-
A,B,-
A,B,-
*/
