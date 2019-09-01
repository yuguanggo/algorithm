package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		x,y,sum,carry int
		p,q,dummy,cur *ListNode
	)
	dummy=&ListNode{}
	cur=dummy
	p=l1
	q=l2
	for p!=nil||q!=nil{
		if p!=nil{
			x=p.Val
		}else{
			x=0
		}
		if q!=nil{
			y=q.Val
		}else{
			y=0
		}
		sum=x+y+carry
		carry+=sum%10
		cur.Next=&ListNode{Val:sum/10,Next:nil}
		if p!=nil{
			p=p.Next
		}
		if q!=nil{
			q=q.Next
		}

		cur=cur.Next
	}
	if carry>0{
		cur.Next=&ListNode{Val:carry,Next:nil}
	}
	return dummy.Next
}

func main() {
	
}
