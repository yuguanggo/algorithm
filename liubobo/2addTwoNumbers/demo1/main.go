package main


type ListNode struct {
	Val int
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

	head:=&ListNode{}
	cur:=head
	for cur1!=nil||cur2!=nil{
		sum:=0
		if cur1==nil{
			sum=cur2.Val+carry
		}else if cur2==nil{
			sum=cur1.Val+carry
		}else{
			sum=cur1.Val+cur2.Val+carry
		}
		if sum>=10{
			cur.Next=&ListNode{
				Val:sum%10,
				Next:nil,
			}
		}else{
			cur.Next=&ListNode{
				Val:sum,
				Next:nil,
			}
		}
		carry=sum/10
		cur=cur.Next
		if cur1!=nil{
			cur1=cur1.Next
		}
		if cur2!=nil{
			cur2=cur2.Next
		}
	}
	if carry>0{
		cur.Next=&ListNode{
			Val:carry,
			Next:nil,
		}
	}
	return head.Next
}

func main() {
	
}
