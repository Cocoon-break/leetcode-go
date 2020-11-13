package LRU

/*
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制 。
实现 LRUCache 类：

LRUCache(int capacity) 以正整数作为容量 capacity 初始化 LRU 缓存
int get(int key) 如果关键字 key 存在于缓存中，则返回关键字的值，否则返回 -1 。
void put(int key, int value) 如果关键字已经存在，则变更其数据值；如果关键字不存在，则插入该组「关键字-值」。当缓存容量达到上限时，它应该在写入新数据之前删除最久未使用的数据值，从而为新的数据值留出空间。


进阶：你是否可以在 O(1) 时间复杂度内完成这两种操作？

参考
https://github.com/labuladong/fucking-algorithm/blob/master/%E9%AB%98%E9%A2%91%E9%9D%A2%E8%AF%95%E7%B3%BB%E5%88%97/LRU%E7%AE%97%E6%B3%95.md
*/

/*
分析上面的操作过程，要让 put 和 get 方法的时间复杂度为 O(1)，我们可以总结出 cache 这个数据结构必要的条件：

1、显然 cache 中的元素必须有时序，以区分最近使用的和久未使用的数据，当容量满了之后要删除最久未使用的那个元素腾位置。

2、我们要在 cache 中快速找某个 key 是否已存在并得到对应的 val；

3、每次访问 cache 中的某个 key，需要将这个元素变为最近使用的，也就是说 cache 要支持在任意位置快速插入和删除元素。

那么，什么数据结构同时符合上述条件呢？哈希表查找快，但是数据无固定顺序；链表有顺序之分，插入删除快，但是查找慢
*/

type LRUCache struct {
	capacity int               //容量
	kv       map[int]*ListNode //快速查找某个key value
	cache    *DoubleList       // 双向链表
}

func Constructor(capacity int) LRUCache {
	return LRUCache{capacity: capacity, kv: make(map[int]*ListNode), cache: newDoubleList()}
}

// 将指定key的节点，移到链表的最后面
func (this *LRUCache) makeRecently(key int) {
	node := this.kv[key]
	this.cache.remove(node)
	this.cache.addLast(node)
}

// 添加节点到最后
func (this *LRUCache) addRecently(key, value int) {
	node := newListNode(key, value)
	this.kv[key] = node
	this.cache.addLast(node)
}

// 删除指定key 的节点
func (this *LRUCache) deleteKey(key int) {
	node := this.kv[key]
	this.cache.remove(node)
	delete(this.kv, key)
}

// 移除链表的头节点
func (this *LRUCache) removeLeastRecently() {
	node := this.cache.removeFirst()
	delete(this.kv, node.key)
}

func (this *LRUCache) Get(key int) int {
	// 无值情况直接返回-1
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	//
	this.makeRecently(key)
	return this.kv[key].value
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.kv[key]; ok {
		this.deleteKey(key)
		this.addRecently(key, value)
		return
	}
	if this.capacity == this.cache.getSize() {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

// 定义双向链表
type DoubleList struct {
	head *ListNode // 第一个节点
	tail *ListNode // 最后一个节点
	size int
}

func newDoubleList() *DoubleList {
	// 初始化的时候，head和tail 节点都为哨子节点，这样后续的操作就不用修改这两个变量了
	head := newListNode(0, 0)
	tail := newListNode(0, 0)
	head.next = tail
	tail.pre = head
	return &DoubleList{head: head, tail: tail, size: 0}
}

// 向队尾添加一个节点
func (this *DoubleList) addLast(node *ListNode) {
	// 因为之前构造函数里的节点，所以追加节点，应该在tail节点前的一个节点
	// 先给要添加节点，设置它应该指向的节点
	node.pre = this.tail.pre
	node.next = this.tail

	//this.tail.pre.next 之前应该是this.tail，现在换成现在的node
	this.tail.pre.next = node
	this.tail.pre = node

	this.size++
}

// 该节点应该是已经存在链表中的，所以只需一下操作即可
func (this *DoubleList) remove(node *ListNode) {
	node.pre.next = node.next
	node.next.pre = node.pre
	this.size--
}

func (this *DoubleList) removeFirst() *ListNode {
	// 表示当前未添加过任何节点
	if this.head.next == this.tail {
		return nil
	}
	first := this.head.next
	this.remove(first)
	return first
}

func (this *DoubleList) getSize() int {
	return this.size
}

// 定义双向链表节点
type ListNode struct {
	key   int
	value int
	pre   *ListNode
	next  *ListNode
}

func newListNode(key, value int) *ListNode {
	return &ListNode{key: key, value: value}
}
