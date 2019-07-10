package main

import (
	"fmt"
)

/**
24. 两两交换链表中的节点
给定一个链表，两两交换其中相邻的节点，并返回交换后的链表。

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

 

示例:

给定 1->2->3->4, 你应该返回 2->1->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/swap-nodes-in-pairs
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	pre:=new(ListNode)
	pre.Next=head
	cur:=pre
	for cur.Next!=nil&&cur.Next.Next!=nil{
		start:=cur.Next
		end:=cur.Next.Next
		cur.Next=end
		start.Next=end.Next
		end.Next=start
		cur=start
	}
	return pre.Next
}

func swapPairs2(head *ListNode) *ListNode  {
	if head==nil||head.Next==nil{
		return head
	}
	next:=head.Next
	head.Next=swapPairs(next.Next)
	next.Next=head
	return next
}

func swapPairs3(head *ListNode) *ListNode {
	dummy:=new(ListNode)
	dummy.Next=head
	pre:=dummy
	cur:=head
	for cur!=nil&&cur.Next!=nil{
		first:=cur
		second:=cur.Next
		pre.Next=second
		first.Next=second.Next
		second.Next=first
		pre=first
		cur=first.Next
	}
	return dummy.Next
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
	l1=swapPairs(l1)
	fmt.Println(l1)
	fmt.Println(l1.Next)
	fmt.Println(l1.Next.Next)
	fmt.Println(l1.Next.Next.Next)
}
