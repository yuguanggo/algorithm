package main

import (
	"fmt"
)

type dLinkedNode struct {
	key int
	val int
	count int //使用次数
	pre *dLinkedNode
	next *dLinkedNode
}

type linkedList struct {
	head *dLinkedNode  //存存双向链表的头部
	tail *dLinkedNode  //存储双向链表的尾部
}

type LFUCache struct {
	cap int
	len int
	mpKvs map[int]*dLinkedNode //存储key,v
	mpCountList map[int]*linkedList //存储相同使用次数的节点值
	minFreq int //最少使用次数
}


func Constructor(capacity int) LFUCache {
	return LFUCache{
		cap:capacity,
		len:0,
		mpKvs:make(map[int]*dLinkedNode),
		mpCountList:make(map[int]*linkedList),
		minFreq:1,
	}
}


func (this *LFUCache) Get(key int) int {
	//判断是否存在
	if node,ok:=this.mpKvs[key];ok{
		//用put更新
		this.Put(node.key,node.val)
		return node.val
	}
	return -1
}


func (this *LFUCache) Put(key int, value int)  {
	if this.cap==0{
		return
	}
	//判断是否存在
	if node,ok:=this.mpKvs[key];ok{
		//将当前节点从mpCountList[node.count]的list删除
		this.removeNode(this.mpCountList[node.count],node)

		//判断this.mpCountList[node.count]list是否为空 并更新minFreq
		if this.mpCountList[node.count].tail.pre== this.mpCountList[node.count].head&&node.count==this.minFreq{
			this.minFreq++
		}

		//将节点加入到mpCountList[node.count+1]的头部
		node.val=value
		node.count=node.count+1
		this.createList(node.count)
		this.addNode(this.mpCountList[node.count],node)

	}else{
		//创建新节点
		node:=&dLinkedNode{
			key:key,
			val:value,
			count:1,
			pre:nil,
			next:nil,
		}
		//初始化count为1的链表
		this.createList(1)
		//判断缓存是否已满
		if this.len==this.cap{
			//删除最小使用次数的最后一个元素
			node:=this.removeLast(this.mpCountList[this.minFreq])
			//从map中删除对应的key值
			delete(this.mpKvs,node.key)
		}
		//添加节点
		this.addNode(this.mpCountList[1],node)
		this.mpKvs[node.key]=node
		this.minFreq=1
	}
}

func (this *LFUCache)createList(count int)  {
	//判断对应的count是否存在
	if _,ok:=this.mpCountList[count];!ok{
		head:=&dLinkedNode{}
		tail:=&dLinkedNode{}
		head.next=tail
		tail.pre=head
		list:=&linkedList{
			head:head,
			tail:tail,
		}
		this.mpCountList[count]=list
	}
}

//list头部添加节点
func (this *LFUCache) addNode(list *linkedList,node *dLinkedNode)  {
	node.next=list.head.next
	node.pre=list.head
	list.head.next.pre=node
	list.head.next=node
	this.len++
}
//删除节点
func (this *LFUCache) removeNode(list *linkedList,node *dLinkedNode)  {
	node.pre.next=node.next
	node.next.pre=node.pre
	this.len--
}
//移除最后一位
func (this *LFUCache) removeLast(list *linkedList) *dLinkedNode{
	if list.tail.pre==list.head{
		return nil
	}
	last:=list.tail.pre
	this.removeNode(list,last)
	return last
}


/**
 * Your LFUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func main() {
	cache:=Constructor(2)
	cache.Put(1,1)
	cache.Put(2,2)
	fmt.Println(cache.Get(1))
	cache.Put(3,3)
	fmt.Println(cache.Get(2))
	fmt.Println(cache.Get(3))
	cache.Put(4,4)
	fmt.Println(cache.Get(1))
	fmt.Println(cache.Get(3))
	fmt.Println(cache.Get(4))
}
