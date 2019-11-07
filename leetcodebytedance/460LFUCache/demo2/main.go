package main

type dLinkedNode struct {
	key,val int
	count int
	pre *dLinkedNode
	next *dLinkedNode
}

type dLinkedList struct {
	head *dLinkedNode
	tail *dLinkedNode
}

type LFUCache struct {
	len,cap int
	mpKvs map[int]*dLinkedNode
	mpCountList map[int]*dLinkedList
	minFreq int
}

func Constructor(capacity int) LFUCache  {
	return LFUCache{
		cap:capacity,
		len:0,
		mpKvs:make(map[int]*dLinkedNode),
		mpCountList:make(map[int]*dLinkedList),
		minFreq:0,
	}
}

func (this *LFUCache) Get(key int) int  {
	if node,ok:=this.mpKvs[key];ok{
		this.Put(node.key,node.val)
		return node.val
	}
	return -1
}

func (this *LFUCache) Put(key int,value int)  {
	if this.cap==0{
		return
	}
	if node,ok:=this.mpKvs[key];ok{
		//从对应的次数中移除节点
		this.removeNode(this.mpCountList[node.count],node)
		//如果当前次数对应的链表为空，minFreq加1
		if this.mpCountList[node.count].tail.pre==this.mpCountList[node.count].head&&this.minFreq==node.count{
			this.minFreq++
		}
		node.count=node.count+1
		node.val=value
		this.createList(node.count)
		this.addNode(this.mpCountList[node.count],node)
	}else{
		//创建新节点
		node:=&dLinkedNode{
			key:key,
			val:value,
			pre:nil,
			next:nil,
			count:1,
		}
		this.createList(1)
		if this.len==this.cap{
			delNode:=this.removeLast(this.mpCountList[this.minFreq])
			delete(this.mpKvs,delNode.key)
		}
		this.addNode(this.mpCountList[1],node)
		this.mpKvs[node.key]=node
		this.minFreq=1
	}
}

func (this *LFUCache) removeLast(list *dLinkedList) *dLinkedNode  {
	if list.tail.pre==list.head{
		return nil
	}
	last:=list.tail.pre
	this.removeNode(list,last)
	return last
}

func (this *LFUCache) addNode(list *dLinkedList,node *dLinkedNode)  {
	node.next=list.head.next
	node.pre=list.head
	list.head.next.pre=node
	list.head.next=node
	this.len++
}

func (this *LFUCache) createList(count int)  {
	if _,ok:=this.mpCountList[count];!ok{
		head:=&dLinkedNode{}
		tail:=&dLinkedNode{}
		head.next=tail
		tail.pre=head
		list:=&dLinkedList{
			head:head,
			tail:tail,
		}
		this.mpCountList[count]=list
	}
}

func (this *LFUCache) removeNode(list *dLinkedList,node *dLinkedNode)  {
	node.pre.next=node.next
	node.next.pre=node.pre
	this.len--
}

func main() {
	
}
