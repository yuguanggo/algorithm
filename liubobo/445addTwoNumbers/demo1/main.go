package main

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	dummy1:=&ListNode{}
	dummy1.Next=l1
	dummy2:=&ListNode{}
	dummy2.Next=l2
	cur1:=dfs(dummy1,nil)
	cur2:=dfs(dummy2,nil)
	x:=0
	y:=0
	carry:=0
	sum:=0
	resDummy:=&ListNode{}
	for cur1!=dummy1||cur2!=dummy2{
		if cur1!=dummy1{
			x=cur1.Val
			cur1=dfs(dummy1,cur1)
		}else{
			x=0
		}
		if cur2!=dummy2{
			y=cur2.Val
			cur2=dfs(dummy2,cur2)
		}else{
			y=0
		}
		sum=x+y+carry
		carry=sum/10
		tail:=resDummy.Next
		resDummy.Next=&ListNode{
			Val:sum%10,
			Next:tail,
		}
	}
	if carry>0{
		tail:=resDummy.Next
		resDummy.Next=&ListNode{
			Val:carry,
			Next:tail,
		}
	}
	return resDummy.Next
}

func dfs(head *ListNode,end *ListNode) *ListNode  {
	if head==nil||head.Next==nil{
		return head
	}
	if end==head.Next{
		return head
	}
	return dfs(head.Next,end)
}


func main() {
	l5:=&ListNode{
		Val:5,
		Next:nil,
	}

	l4:=&ListNode{
		Val:4,
		Next:l5,
	}
	l3:=&ListNode{
		Val:3,
		Next:l4,
	}
	l2:=&ListNode{
		Val:2,
		Next:l3,
	}
	l1:=&ListNode{
		Val:1,
		Next:l2,
	}

	addTwoNumbers(l1,l4)
}
