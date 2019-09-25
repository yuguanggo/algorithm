package main


type ListNode struct {
	Val int
	Next *ListNode
}


func reverseKGroup(head *ListNode, k int) *ListNode {
    dummy:=&ListNode{}
    dummy.Next=head
    fast:=head
    slow:=head
    pre:=dummy
    i:=1
    for i<=k{
    	if fast.Next!=nil{
    		i++
    		fast=fast.Next
		}else{
			break
		}
	}
	if i<k{
		return head
	}
	//开始反转
	for slow!=fast{
		next:=slow.Next
		slow.Next=pre
		pre=slow
		slow=next
	}
	next:=fast.Next
	fast.Next=pre
	pre=fast
	fast=next
	dummy.Next=pre
	return
}





func main() {
	
}
