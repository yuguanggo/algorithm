package main

import "fmt"

/**
反转一个单链表
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？
1>2>3>4>5>null
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

//递归
func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	p := reverseList(head.Next)
	head.Next.Next = head
	head.Next = nil
	return p
}

//非递归
func reverseList2(head *ListNode) *ListNode {
	var pre *ListNode
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = pre
		pre = cur
		cur = temp
	}
	return pre

}

func main() {
	l5 := ListNode{5, nil}
	l4 := ListNode{4, &l5}
	l3 := ListNode{3, &l4}
	l2 := ListNode{2, &l3}
	l1 := ListNode{1, &l2}
	fmt.Println(l1, l1.Next)
	fmt.Println(l2, l2.Next)
	fmt.Println(l3, l3.Next)
	fmt.Println(l4, l4.Next)
	fmt.Println(l5, l5.Next)
	reverseList(&l1)
	fmt.Println(l5, l5.Next)
	fmt.Println(l4, l4.Next)
	fmt.Println(l3, l3.Next)
	fmt.Println(l2, l2.Next)
	fmt.Println(l1, l1.Next)

}
