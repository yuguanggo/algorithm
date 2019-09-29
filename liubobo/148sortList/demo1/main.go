package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortList(head *ListNode) *ListNode {
	return helper(head)
}

func helper(head *ListNode) *ListNode{
	if head==nil||head.Next==nil{
		return head
	}
	slow:=head
	fast:=head
	for fast.Next!=nil&&fast.Next.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	tail:=slow.Next
	slow.Next=nil
	return mergeSortList(helper(head),helper(tail))
}

func mergeSortList(l1,l2 *ListNode) *ListNode  {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=mergeSortList(l1.Next,l2)
		return l1
	}else{
		l2.Next=mergeSortList(l1,l2.Next)
		return l2
	}
}


func main() {
	
}
