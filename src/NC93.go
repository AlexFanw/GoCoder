package main

/*
* 设计LRU缓存结构，该结构在构造时确定大小，假设大小为K，并有如下两个功能
* set(key, value)：将记录(key, value)插入该结构
* get(key)：返回key对应的value值
* [要求]
* set和get方法的时间复杂度为O(1)
* 某个key的set或get操作一旦发生，认为这个key的记录成了最常使用的。
* 当缓存的大小超过K时，移除最不经常使用的记录，即set或get最久远的。
* 若opt=1，接下来两个整数x, y，表示set(x, y)
* 若opt=2，接下来一个整数x，表示get(x)，若x未出现过或已被移除，则返回-1
* 对于每个操作2，输出一个答案
 */

/**
 * lru design
 * @param operators int整型二维数组 the ops
 * @param k int整型 the k
 * @return int整型一维数组
 */
type LRUCache struct {
	size int
	capacity int
	cache map[int]*DLinkedNode
	head, tail *DLinkedNode
}

type DLinkedNode struct {
	key, value int
	prev, next *DLinkedNode
}

func initDLinkedNode(key, value int) *DLinkedNode {
	return &DLinkedNode{
		key: key,
		value: value,
	}
}

func Constructor(capacity int) LRUCache {
	l := LRUCache{
		cache: map[int]*DLinkedNode{},
		head: initDLinkedNode(0, 0),
		tail: initDLinkedNode(0, 0),
		capacity: capacity,
	}
	l.head.next = l.tail
	l.tail.prev = l.head
	return l
}

func (this *LRUCache) Get(key int) int {
	if _, ok := this.cache[key]; !ok {
		return -1
	}
	node := this.cache[key]
	this.moveToHead(node)
	return node.value
}


func (this *LRUCache) Put(key int, value int)  {
	if _, ok := this.cache[key]; !ok {
		node := initDLinkedNode(key, value)
		this.cache[key] = node
		this.addToHead(node)
		this.size++
		if this.size > this.capacity {
			removed := this.removeTail()
			delete(this.cache, removed.key)
			this.size--
		}
	} else {
		node := this.cache[key]
		node.value = value
		this.moveToHead(node)
	}
}

func (this *LRUCache) addToHead(node *DLinkedNode) {
	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

func (this *LRUCache) removeNode(node *DLinkedNode) {
	node.prev.next = node.next
	node.next.prev = node.prev
}

func (this *LRUCache) moveToHead(node *DLinkedNode) {
	this.removeNode(node)
	this.addToHead(node)
}

func (this *LRUCache) removeTail() *DLinkedNode {
	node := this.tail.prev
	this.removeNode(node)
	return node
}