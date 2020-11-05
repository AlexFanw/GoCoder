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
func main() {
	
}
//Cache总结构
type LRUCache struct {
	size int
	capacity int
	cache map[int] *DlinkedNode
	head, tail *DlinkedNode
}
//双向链表
type DlinkedNode struct {
	key, value int
	prev, next *DlinkedNode
}
//func LRU( operators [][]int ,  k int ) []int {
//
//}