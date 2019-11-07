package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	carry:=0
	cur1:=l1
	cur2:=l2
	x:=0
	y:=0
	dummy:=&ListNode{}
	pre:=dummy
	for cur1!=nil||cur2!=nil{
		if cur1!=nil{
			x=cur1.Val
		}else{
			x=0
		}
		if cur2!=nil{
			y=cur2.Val
		}else{
			y=0
		}
		sum:=x+y+carry
		carry=sum/10
		pre.Next=&ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		pre=pre.Next
		if cur1!=nil{
			cur1=cur1.Next
		}
		if cur2!=nil{
			cur2=cur2.Next
		}
	}
	if carry>0{
		pre.Next=&ListNode{
			Val:carry,
			Next:nil,
		}
	}
	return dummy.Next
}

func main() {

}
