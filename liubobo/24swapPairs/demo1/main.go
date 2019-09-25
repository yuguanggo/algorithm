package main

type ListNode struct {
	Val int
	Next *ListNode
}


func swapPairs(head *ListNode) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	pre:=dummy
	cur:=head
	for cur!=nil&&cur.Next!=nil{
		next:=cur.Next.Next
		temp:=cur
		pre.Next=cur.Next
		pre.Next.Next=cur
		temp.Next=next
		pre=temp
		cur=next
	}
	return dummy.Next
}


func main() {
	
}
