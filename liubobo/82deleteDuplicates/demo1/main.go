package main

type ListNode struct {
	Val int
	Next *ListNode
}


func deleteDuplicates(head *ListNode) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	slow:=dummy
	fast:=head
	for fast!=nil{
		if fast.Next!=nil&&fast.Next.Val==fast.Val{
			temp:=fast.Val
			for fast!=nil&&fast.Val==temp{
				fast=fast.Next
			}
			slow.Next=fast
		}else{
			fast=fast.Next
			slow=slow.Next
		}
	}
	return dummy.Next
}


func main() {
	
}
