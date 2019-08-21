package main

import "fmt"

/**
排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 type ListNode struct {
 	Val int
 	Next *ListNode
 }
func sortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	pre:=head
	fast:=head
	slow:=head
	for fast!=nil&&fast.Next!=nil{
		pre = slow
		slow = slow.Next
		fast = fast.Next.Next
	}
	pre.Next=nil
	return merge(sortList(head),sortList(slow))
}

func merge(l1 *ListNode,l2 *ListNode) *ListNode  {
	if l2==nil{
		return l1
	}
	if l1==nil{
		return l2
	}
	if l1.Val<l2.Val{
		l1.Next=merge(l1.Next,l2)
		return l1
	}else{
		l2.Next=merge(l1,l2.Next)
		return l2
	}
}

func printNode(head *ListNode)  {
	for head!=nil{
		fmt.Printf("->%d",head.Val)
		head=head.Next
	}
	fmt.Println()
}

func main() {
	n1:=&ListNode{
		Val:4,
		Next:nil,
	}
	n2:=&ListNode{
		Val:2,
		Next:nil,
	}
	n3:=&ListNode{
		Val:1,
		Next:nil,
	}
	n4:=&ListNode{
		Val:3,
		Next:nil,
	}
	n1.Next=n2
	n2.Next=n3
	n3.Next=n4
	printNode(n1)
	sortList(n1)
}
