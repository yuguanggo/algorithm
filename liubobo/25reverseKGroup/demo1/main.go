package main

type ListNode struct {
	Val int
	Next *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy:=&ListNode{}
    dummy.Next=head
   	pre:=dummy
   	end:=dummy
   	for end.Next!=nil{
   		for i:=0;i<k&&end!=nil;i++{
   			end=end.Next
		}
   		if end==nil{
   			break;
		}
   		start:=pre.Next
   		next:=end.Next
   		end.Next=nil //先断开
   		pre.Next=reserverNode2(start)
   		pre=start
   		pre.Next=next
   		end=pre

	}
	return dummy.Next
}

func reserverNode(head *ListNode) *ListNode  {
	var pre *ListNode
	cur:=head
	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	return pre
}

func reserverNode2(head *ListNode) *ListNode  {
	if head==nil||head.Next==nil{
		return head
	}
	p:=reserverNode2(head.Next)
	head.Next.Next=head
	head.Next=nil
	return p
}





func main() {
	
}
