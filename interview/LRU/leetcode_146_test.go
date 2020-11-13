package LRU

import "testing"

// go test -v . -test.run TestDoubleListAddLast
func TestDoubleListAddLast(t *testing.T) {
	node := newListNode(11, 11)
	node2 := newListNode(22, 22)
	node3 := newListNode(33, 33)
	doubleList := newDoubleList()
	doubleList.addLast(node)
	doubleList.addLast(node2)
	doubleList.addLast(node3)
	t.Logf("头节点的值：%+v", doubleList.head.next.value)
	t.Logf("尾节点的值：%+v", doubleList.tail.pre.value)
	// will remove
	doubleList.removeFirst()
	t.Log("------------")
	head := doubleList.head.next
	for head.next != nil {
		t.Log(head.value)
		head = head.next
	}
	t.Log("------------")
	t.Logf("头节点的值：%+v", doubleList.head.next.value)
	t.Logf("目前长度：%+v", doubleList.getSize())
}

//  go test -v . -test.run TestLRU
func TestLRU(t *testing.T) {
	LRU := Constructor(3)
	LRU.Put(11, 11)
	LRU.Put(22, 22)
	LRU.Put(33, 33)

	t.Log("------------")
	head := LRU.cache.head.next
	for head.next != nil {
		t.Log(head.value)
		head = head.next
	}
	t.Log("------------")
	LRU.Get(11)
	t.Log("------------2")
	head2 := LRU.cache.head.next
	for head2.next != nil {
		t.Log(head2.value)
		head2 = head2.next
	}
	t.Log("------------2")
}
