package main

type ListNode struct {
	Val  int
	Next *ListNode
}


func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	pre:=dummy
	slow:=head
	fast:=head
	for i:=1;i<n;i++{
		fast=fast.Next
	}
	for fast.Next!=nil{
		pre=slow
		slow=slow.Next
		fast=fast.Next
	}
	temp:=slow.Next
	pre.Next=temp
	return dummy.Next
}


func main() {
	
}
