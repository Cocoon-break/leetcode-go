package stack_queue

import "testing"

// go test -v . -test.run TestDecodeString
func TestDecodeString(t *testing.T) {
	result := decodeString("3[a]2[bc]")
	t.Log(result)

	result2 := decodeString("3[a2[c]]")
	t.Log(result2)

	result3 := decodeString("2[abc]3[cd]ef")
	t.Log(result3)

	result4 := decodeString("3[z]2[2[y]pq4[2[jk]e1[f]]]ef")
	t.Log(result4)
}
