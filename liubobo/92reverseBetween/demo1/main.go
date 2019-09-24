package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	dummyHead:=&ListNode{}
	dummyHead.Next=head
	i:=1
	pre:=dummyHead
	cur:=head
	start:=&ListNode{}
	startPre:=&ListNode{}
	for cur!=nil{
		next:=cur.Next
		if i==m{
			start=cur
			startPre=pre
		}
		if i>m&&i<=n{
			cur.Next=pre
		}
		if i==n{
			startPre.Next=cur
			start.Next=next
		}
		pre=cur
		cur=next
		i++
	}
	return dummyHead.Next
}
func main() {
	
}
