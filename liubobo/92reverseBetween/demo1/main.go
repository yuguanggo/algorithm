package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	var dummyHead *ListNode
	dummyHead.Next=head
	i:=1
	pre:=dummyHead
	cur:=head
	for cur!=nil{
		next:=cur.Next
		if i>=m&&i<=n{

		}
		
	}
}
func main() {
	
}
