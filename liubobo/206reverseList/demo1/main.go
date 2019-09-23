package main

import "fmt"

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type ListNode struct {
	Val  int
	Next *ListNode
}
func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	p:=reverseList(head.Next)
	head.Next.Next=head
	head.Next=nil
	fmt.Println(p)
	return p
}

func reverseList2(head *ListNode) *ListNode {
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
func main() {
	l4:=&ListNode{
		Val:4,
		Next:nil,
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
	fmt.Println(reverseList(l1))
}
