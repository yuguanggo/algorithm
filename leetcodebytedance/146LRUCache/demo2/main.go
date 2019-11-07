package main

type dLinkedNode struct {
	pre, next *dLinkedNode
	key,val int
}

type LRUCache struct {
	len,cap int
	first,last *dLinkedNode
	nodeMap map[int]*dLinkedNode
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		len:0,
		cap:capacity,
		first:nil,
		last:nil,
		nodeMap:make(map[int]*dLinkedNode),
	}
}

func (this *LRUCache) Get(key int) int {
	if	node,ok:=this.nodeMap[key];ok{
		this.removeNode(node)
		return node.val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int)  {
	if node,ok:=this.nodeMap[key];ok{
		node.val=value
		this.removeNode(node)
	}else{
		if this.len==this.cap{
			delete(this.nodeMap,this.last.key)
			this.removeLast()
		}else{
			this.len++
		}
		node:=&dLinkedNode{
			pre:nil,
			next:nil,
			key:key,
			val:value,
		}
		this.nodeMap[key]=node
		this.insertNode(node)
	}
}

func (this *LRUCache) removeNode(node *dLinkedNode)  {
	switch node {
	case this.first:
		return
	case this.last:
		this.removeLast()
	default:
		node.next.pre=node.pre
		node.pre.next=node.next
	}
	this.insertNode(node)
}

func (this *LRUCache) removeLast()  {
	if this.last.pre!=nil{
		this.last.pre.next=nil
	}else{
		this.first=nil
	}
	this.last=this.last.pre
}

func (this *LRUCache) insertNode(node *dLinkedNode)  {
	if this.last==nil{
		this.last=node
	}else{
		node.next=this.first
		this.first.pre=node
	}
	this.first=node
}


func main() {
	
}
