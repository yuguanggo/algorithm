package main


type ListNode struct {
	Val int
	Next *ListNode
}


func removeElements(head *ListNode, val int) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	pre:=dummy
	cur:=head
	for cur!=nil{
		if cur.Val==val{
			pre.Next=cur.Next
		}else{
			pre=cur
		}
		cur=cur.Next
	}
	return dummy.Next
}


func main() {
	
}
