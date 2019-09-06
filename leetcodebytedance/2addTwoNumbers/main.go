package main

import "fmt"

/**
两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var (
		p                *ListNode
		q                *ListNode
		dummy            *ListNode
		cur              *ListNode
		x, y, carry, sum int
	)
	p = l1
	q = l2
	dummy = &ListNode{0, nil}
	cur = dummy
	for p != nil || q != nil {
		if p != nil {
			x = p.Val
		} else {
			x = 0
		}
		if q != nil {
			y = q.Val
		} else {
			y = 0
		}
		sum = carry + x + y
		carry = sum / 10
		cur.Next = &ListNode{Val: sum % 10, Next: nil}
		cur = cur.Next
		if p != nil {
			p = p.Next
		}
		if q != nil {
			q = q.Next
		}
	}
	if carry > 0 {
		cur.Next = &ListNode{Val: carry}
	}
	return dummy.Next
}
func main() {
	l3 := &ListNode{Val: 4, Next: nil}
	l2 := &ListNode{Val: 3, Next: l3}
	l1 := &ListNode{Val: 2, Next: l2}

	l4 := &ListNode{Val: 4, Next: nil}
	l5 := &ListNode{Val: 6, Next: l4}
	l6 := &ListNode{Val: 5, Next: l5}
	sum := addTwoNumbers(l1, l6)
	fmt.Println(sum)
}
