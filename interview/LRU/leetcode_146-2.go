package LRU

import (
	"container/list"
	"fmt"
	"strconv"
	"strings"
)

type LRUCache2 struct {
	capacity int                   //容量
	kv       map[int]*list.Element //快速查找某个key value
	cache    *list.List            // 双向链表
}

func Constructor2(capacity int) LRUCache2 {
	return LRUCache2{capacity: capacity, kv: make(map[int]*list.Element), cache: list.New()}
}

func (this *LRUCache2) Put(key int, value int) {
	if _, ok := this.kv[key]; ok {
		//先删除
		this.deleteKey(key)
		//在添加
		this.addRecently(key, value)
		return
	}
	if this.capacity == this.cache.Len() {
		this.removeLeastRecently()
	}
	this.addRecently(key, value)
}

func (this *LRUCache2) Get(key int) int {
	if _, ok := this.kv[key]; !ok {
		return -1
	}
	this.makeRecently(key)
	// 获取值
	str := this.kv[key].Value.(string)
	arr := strings.Split(str, ",")
	value, _ := strconv.Atoi(arr[1])
	return value
}

func (this *LRUCache2) makeRecently(key int) {
	node := this.kv[key]
	this.cache.MoveToBack(node)
}
func (this *LRUCache2) deleteKey(key int) {
	node := this.kv[key]
	this.cache.Remove(node)
	delete(this.kv, key)
}

func (this *LRUCache2) addRecently(key, value int) {
	//拼成字符串key,value
	str := fmt.Sprintf("%+v,%+v", key, value)
	node := this.cache.PushBack(str)
	this.kv[key] = node
}
func (this *LRUCache2) removeLeastRecently() {
	firstNode := this.cache.Front()
	this.cache.Remove(firstNode)
	// 获取key
	str := firstNode.Value.(string)
	arr := strings.Split(str, ",")
	key, _ := strconv.Atoi(arr[0])
	delete(this.kv, key)
}
