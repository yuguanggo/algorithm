package main

import "fmt"

/**
61. 旋转链表
给定一个链表，旋转链表，将链表每个节点向右移动 k 个位置，其中 k 是非负数。

示例 1:

输入: 1->2->3->4->5->NULL, k = 2
输出: 4->5->1->2->3->NULL
解释:
向右旋转 1 步: 5->1->2->3->4->NULL
向右旋转 2 步: 4->5->1->2->3->NULL
示例 2:

输入: 0->1->2->NULL, k = 4
输出: 2->0->1->NULL
解释:
向右旋转 1 步: 2->0->1->NULL
向右旋转 2 步: 1->2->0->NULL
向右旋转 3 步: 0->1->2->NULL
向右旋转 4 步: 2->0->1->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/rotate-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func rotateRight(head *ListNode, k int) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	dummy := new(ListNode)
	dummy.Next = head
	p := dummy
	q := dummy
	preq:=dummy
	for i := 0; i < k+1; i++ {
		if q.Next==nil{
			if i==k{
				return head
			}
			q=dummy.Next
		}else{
			q = q.Next
		}

	}
	for q!=nil{
		p=p.Next
		preq=q
		q=q.Next
	}

	preq.Next=head
	dummy.Next=p.Next
	p.Next=nil
	return dummy.Next
}

func rotateRight2(head *ListNode, k int) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	n:=1
	tail:=head
	for ;tail.Next!=nil;n++{
		tail=tail.Next
	}
	tail.Next=head
	newTail:=head
	for i:=0;i<(n-k%n-1);i++{
       newTail=newTail.Next
	}
	newHead := newTail.Next
	newTail.Next=nil
	return newHead
}

func main() {
	l3:=&ListNode{
		Val:2,
		Next:nil,
	}
	l2:=&ListNode{
		Val:1,
		Next:l3,
	}
	l1:=&ListNode{
		Val:0,
		Next:l2,
	}
	fmt.Println(rotateRight(l1,4))
}
