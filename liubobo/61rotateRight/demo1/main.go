package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	cur:=head
	num:=0
	for cur!=nil{
		cur=cur.Next
		num++
	}
	if num==0||k==0{
		return head
	}
	newK:=k%num
	if newK==0{
		return head
	}
	dummy:=&ListNode{}
	slow:=head
	fast:=head
	for i:=1;i<=newK;i++{
		fast=fast.Next
	}
	for fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next
	}
	dummy.Next=slow.Next
	fast.Next=head
	slow.Next=nil
	return dummy.Next
}


func main() {
	
}
