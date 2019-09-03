package main

import "fmt"

type dLinkedNode struct {
	pre *dLinkedNode
	next *dLinkedNode
	key int
	val int
}

type LRUCache struct {
	cap int
	len int
	head *dLinkedNode
	tail *dLinkedNode
	h map[int]*dLinkedNode
}


func Constructor(capacity int) LRUCache {
	head:=&dLinkedNode{}
	tail:=&dLinkedNode{}
	head.next=tail
	tail.pre=head
	return LRUCache{
		cap:capacity,
		len:0,
		head:head,//头虚节点
		tail:tail,//
		h:make(map[int]*dLinkedNode),
	}
}


func (this *LRUCache) Get(key int) int {
	//判断是否存在
	if v,ok:=this.h[key];ok{
		//用put更新节点
		this.Put(v.key,v.val)
		return v.val
	}
	return -1
}


func (this *LRUCache) Put(key int, value int)  {
	//判断是否存在
	if v,ok:=this.h[key];ok{
		//更新值
		v.val=value
		//删除该节点
		this.removeNode(v)
		//将该节点添加到头部
		this.addNode(v)
		this.h[key]=v
	}else{
		//判断是否已满
		if this.len==this.cap{
			//删除链接尾部节点
			v:=this.removeLast()
			//从map里删除
			delete(this.h,v.key)
		}
		v:=&dLinkedNode{key:key,val:value}
		//将节点添加到头部
		this.addNode(v)
		this.h[key]=v
	}
}

//在链表头部添加节点
func (this *LRUCache) addNode(node *dLinkedNode)  {
	//添加节点
	node.next=this.head.next
	node.pre=this.head

	this.head.next.pre=node
	this.head.next=node
	this.len++
}

//删除链表中的节点
func (this *LRUCache) removeNode(node *dLinkedNode)  {
	node.pre.next=node.next
	node.next.pre=node.pre
	this.len--
}

//删除链表的最后一个节点
func (this *LRUCache) removeLast() *dLinkedNode {
	if this.tail.pre==this.head{
		return nil
	}
	last:=this.tail.pre
	this.removeNode(last)
	return last
}


/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache := Constructor(2)
	cache.Put(1, 1)
	cache.Put(2, 2)
	fmt.Println(cache.Get(1))
	cache.Put(3, 3)
	fmt.Println(cache.Get(2))
	cache.Put(4, 4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))

	//cache := Constructor(2)
	//cache.Put(2, 1)
	//cache.Put(1, 1)
	//fmt.Println(cache.Get(2))
	//cache.Put(4, 1)
	//fmt.Println(cache.Get(1))
	//fmt.Println(cache.Get(2))
}
