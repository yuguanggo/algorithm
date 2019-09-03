package main

/**
LRU缓存机制
运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

*/

type doublyLinkedNode struct {
	pre, next *doublyLinkedNode
	key, val  int
}

type LRUCache struct {
	len, cap    int
	first, last *doublyLinkedNode         //head tail
	nodes       map[int]*doublyLinkedNode //find o(1)
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		len:   0,
		cap:   capacity,
		first: nil,
		last:  nil,
		nodes: make(map[int]*doublyLinkedNode, capacity),
	}
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.nodes[key]; ok {
		this.movetofirst(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.nodes[key]; ok {
		node.val = value
		this.movetofirst(node)
	} else {
		if this.len == this.cap {
			delete(this.nodes,this.last.key)
			this.removelast()
		} else {
			this.len++
		}
		node := &doublyLinkedNode{
			pre:  nil,
			next: nil,
			key:  key,
			val:  value,
		}
		this.nodes[key] = node
		this.inserttolist(node)
	}
}

func (this *LRUCache) movetofirst(node *doublyLinkedNode) {
	switch node {
	case this.first:
		return
	case this.last:
		this.removelast()
	default:
		//删除当前Node
		node.next.pre = node.pre
		node.pre.next = node.next
	}
	//在插入node
	this.inserttolist(node)
}

func (this *LRUCache) removelast() {
	if this.last.pre != nil {
		this.last.pre.next = nil
	} else {
		//链表长度==1
		this.first = nil
	}
	this.last = this.last.pre
}

func (this *LRUCache) inserttolist(node *doublyLinkedNode) {
	if this.last == nil {
		//空的链表
		this.last = node
	} else {
		node.next = this.first
		this.first.pre = node
	}
	this.first = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
//示例:
//LRUCache cache = new LRUCache( 2 /* 缓存容量 */ );
//
//cache.put(1, 1);
//cache.put(2, 2);
//cache.get(1);       // 返回  1
//cache.put(3, 3);    // 该操作会使得密钥 2 作废
//cache.get(2);       // 返回 -1 (未找到)
//cache.put(4, 4);    // 该操作会使得密钥 1 作废
//cache.get(1);       // 返回 -1 (未找到)
//cache.get(3);       // 返回  3
//cache.get(4);       // 返回  4
func main() {
	//cache := Constructor(2)
	//cache.Put(1, 1)
	//cache.Put(2, 2)
	//fmt.Println(cache.Get(1))
	//cache.Put(3, 3)
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 4)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(3))
	//fmt.Println(cache.Get(4))

	//cache := Constructor(2)
	//cache.Put(2, 1)
	//cache.Put(1, 1)
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 1)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(2))

}
