package main

type ListNode struct {
	Val  int
	Next *ListNode
}


func getIntersectionNode(headA, headB *ListNode) *ListNode {
	a:=headA
	b:=headB
	for a!=b{
		if a==nil{
			a=headB
		}else{
			a=a.Next
		}
		if b==nil{
			b=headA
		}else{
			b=b.Next
		}
	}
	return a
}

func main() {
	
}
