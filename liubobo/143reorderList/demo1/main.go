package main

type ListNode struct {
	Val  int
	Next *ListNode
}


func reorderList(head *ListNode)  {
	if head==nil||head.Next==nil{
		return
	}
	tail:=partion(head)
	tail=resverList(tail)
	//合并两个链表
	h:=head
	t:=tail
	for t!=nil{
		hnext:=h.Next
		tnext:=t.Next
		h.Next=t
		t.Next=hnext
		h=hnext
		t=tnext
	}

}

//反转链表
func resverList(head *ListNode) *ListNode  {
	var pre *ListNode
	cur:=head
	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return pre
}

//返回中间节点
func partion(head *ListNode) *ListNode  {

	slow:=head
	fast:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	tail:=slow.Next
	slow.Next=nil
	return 	tail
}


func main() {
	
}
