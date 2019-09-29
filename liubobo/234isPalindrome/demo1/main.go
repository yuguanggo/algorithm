package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}


func isPalindrome(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return true
	}
	slow:=head
	pre:=slow
	fast:=head
	for fast!=nil&&fast.Next!=nil{
		pre=slow
		slow=slow.Next
		fast=fast.Next.Next
	}
	mid:=pre.Next
	pre.Next=nil
	//反转后半部分
	var midPre *ListNode
	cur:=mid
	for cur!=nil{
		next:=cur.Next
		cur.Next=midPre
		midPre=cur
		cur=next
	}
	fmt.Println()
	mid=midPre
	midCur:=mid
	hCur:=head
	for hCur!=nil{
		if midCur.Val!=hCur.Val{
			return false
		}
		midCur=midCur.Next
		hCur=hCur.Next
	}
	return true
}
func main() {
	var pre *ListNode
	fmt.Println(pre)
	fmt.Println(pre==nil)
}
