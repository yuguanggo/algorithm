package main

import "fmt"

/**
全 O(1) 的数据结构
实现一个数据结构支持以下操作：

Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
挑战：以 O(1) 的时间复杂度实现所有操作。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/all-oone-data-structure
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

 //单个节点
 type dLinkedNode struct {
 	key string
 	count int
 	pre *dLinkedNode
 	next *dLinkedNode
 }

 //双向链表
 type dLinkedList struct {
 	head *dLinkedNode
 	tail *dLinkedNode
 }
//捅的双向链表，用来判断最大和最小值
type bucket struct {
	count int //count值
	list *dLinkedList //相同count值的节点
	pre *bucket
	next *bucket
}


type AllOne struct {
	head *bucket //
	tail *bucket //
	mpKvs map[string]*dLinkedNode //存储key值
	mpBuckets map[int]*bucket //count 对应的桶
}


/** Initialize your data structure here. */
func Constructor() AllOne {
	head:=&bucket{}
	tail:=&bucket{}
	head.next=tail
	tail.pre=head
	return AllOne{
		head:head,
		tail:tail,
		mpKvs:make(map[string]*dLinkedNode),
		mpBuckets:make(map[int]*bucket),
	}
}


/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string)  {
	if key==""{
		return
	}
	//判断是否存在
	if node,ok:=this.mpKvs[key];ok{
		//从对应的桶中删除node
		preBucket:=this.mpBuckets[node.count]
		this.removeNode(node)
		//判断链表是否为空，为空的话删除所在桶

		if this.mpBuckets[node.count].list.tail.pre==this.mpBuckets[node.count].list.head{
			//为空的话更新桶的前节点
			preBucket=this.mpBuckets[node.count].pre
			//删除当前桶
			this.removeBucket(this.mpBuckets[node.count])

			delete(this.mpBuckets,node.count)

		}
		//创建新的桶
		node.count=node.count+1
		this.mpKvs[key]=node
		//判断桶是否存在
		if _,ok:=this.mpBuckets[node.count];!ok{
			this.createBucket(node.count)
			//将桶插入链表里
			this.addNewBucketAfter(preBucket,this.mpBuckets[node.count])
		}

		//将node插入新桶里
		this.addNode(node.count,node)


	}else{
		//不存在
		//创建新节点
		node:=&dLinkedNode{
			key:key,
			count:1,
			pre:nil,
			next:nil,
		}
		//创建对应的桶
		//判断桶是否存在
		if _,ok:=this.mpBuckets[1];!ok{
			this.createBucket(1)
			//将桶插入桶的链表里
			this.addNewBucketAfter(this.head,this.mpBuckets[1])
		}

		//将node 插入桶的list链表里
		this.addNode(1,node)
		this.mpKvs[key]=node


	}
}




/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string)  {
	if key==""{
		return
	}
	//判断是否存在
	if node,ok:=this.mpKvs[key];ok{
		//从对应的桶中删除node
		preBucket:=this.mpBuckets[node.count]
		this.removeNode(node)
		//判断链表是否为空，为空的话删除所在桶

		if this.mpBuckets[node.count].list.tail.pre==this.mpBuckets[node.count].list.head{
			preBucket=this.mpBuckets[node.count].pre
			//删除当前桶
			this.removeBucket(this.mpBuckets[node.count])

			delete(this.mpBuckets,node.count)
		}
		if node.count==1{
			//删除node
			delete(this.mpKvs,node.key)
			return
		}
		//创建新的桶
		node.count=node.count-1
		this.mpKvs[key]=node
		if _,ok:=this.mpBuckets[node.count];!ok{
			this.createBucket(node.count)
			//将桶插入链表里
			this.addNewBucketAfter(preBucket,this.mpBuckets[node.count])
		}

		//将node插入新桶里
		this.addNode(node.count,node)


	}
	return
}


/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail.pre==this.head{
		return ""
	}
	return this.tail.pre.list.head.next.key
}


/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.tail.pre==this.head{
		return ""
	}
	return this.head.next.list.head.next.key
}

func (this *AllOne)createBucket(count int)  {

		//创建桶
		dHead:=&dLinkedNode{}
		dTail:=&dLinkedNode{}
		dHead.next=dTail
		dTail.pre=dHead
		linkedList:=&dLinkedList{
			head:dHead,
			tail:dTail,
		}
		b:=&bucket{
			count:count,
			list:linkedList,
			pre:nil,
			next:nil,
		}
		this.mpBuckets[count]=b

}

//在bucket链表某个节点的后面增加 bucket
func (this *AllOne)addNewBucketAfter(pre *bucket,b *bucket)  {
	b.next=pre.next
	b.pre=pre
	pre.next.pre=b
	pre.next=b
}



//删除bucket节点
func (this *AllOne)removeBucket(b *bucket)  {
	b.pre.next=b.next
	b.next.pre=b.pre
}

//在对应的桶的链表里插入node
func (this *AllOne)addNode(count int,node *dLinkedNode)  {
	l:=this.mpBuckets[count].list
	node.next=l.head.next
	node.pre=l.head
	l.head.next.pre=node
	l.head.next=node
}

func (this *AllOne)removeNode(node *dLinkedNode)  {
	node.pre.next=node.next
	node.next.pre=node.pre
}


/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
func main() {
	allone:=Constructor()
	allone.Inc("hello")
	allone.Inc("goodbye")
	allone.Inc("hello")
	fmt.Println(allone.GetMaxKey())
	allone.Inc("leet")
	allone.Inc("code")
	allone.Inc("leet")
	allone.Dec("hello")
	allone.Inc("leet")
	allone.Inc("code")
	allone.Inc("code")

	fmt.Println(allone.GetMaxKey())

}
