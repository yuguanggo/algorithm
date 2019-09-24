package main


type ListNode struct {
	Val int
	Next *ListNode
}


func hasCycle(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return false
	}
	fast:=head
	slow:=head.Next
	for fast!=nil&&fast.Next.Next!=nil{
		fast=fast.Next.Next
		slow=slow.Next
		if fast==slow{
			return true
		}
	}
	return false
}


func main() {
	
}
