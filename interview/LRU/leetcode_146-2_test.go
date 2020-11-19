package LRU

import "testing"

// go test -v . -test.run TestLRU2
func TestLRU2(t *testing.T) {
	LRU := Constructor2(2)
	result := LRU.Get(2)
	t.Log(result)
	LRU.Put(1, 1)
	result2 := LRU.Get(1)
	t.Log(result2)
	// LRU.Put(2, 2)
	// result1 := LRU.Get(1)
	// t.Log(result1)
	// LRU.Put(3, 3)
	// result2 := LRU.Get(2)
	// t.Log(result2)
	// LRU.Put(4, 4)
	// LRU.Get(1)
	// LRU.Get(3)
	// LRU.Get(4)
	// t.Log("------------")
	// head := LRU.cache.Front()
	// for head != nil {
	// 	t.Log(head.Value)
	// 	head = head.Next()
	// }
	// t.Log("------------")

	// LRU.Get(11)
	// t.Log("------------2")
	// head2 := LRU.cache.Front()
	// for head2 != nil {
	// 	t.Log(head2.Value)
	// 	head2 = head2.Next()
	// }
	// t.Log("------------2")
}
