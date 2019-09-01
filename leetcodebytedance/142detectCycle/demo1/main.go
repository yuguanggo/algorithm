package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	fast:=head
	slow:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			fast=head
			for slow!=fast{
				slow=slow.Next
				fast=fast.Next
			}
			return slow
		}
	}
	return nil
}
func main() {
	
}
