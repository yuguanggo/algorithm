package main

/**
143. 重排链表
给定一个单链表 L：L0→L1→…→Ln-1→Ln ，
将其重新排列后变为： L0→Ln→L1→Ln-1→L2→Ln-2→…

你不能只是单纯的改变节点内部的值，而是需要实际的进行节点交换。

示例 1:

给定链表 1->2->3->4, 重新排列为 1->4->2->3.
示例 2:

给定链表 1->2->3->4->5, 重新排列为 1->5->2->4->3.

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reorder-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func reorderList(head *ListNode)  {
	if head==nil||head.Next==nil||head.Next.Next==nil{
		return
	}
	//将链表分成2段
	slow:=head
	fast:=head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
	}
	tail:=slow.Next
	slow.Next=nil
	//反转后一部分链表
	cur:=tail
	var pre *ListNode=nil
	for cur!=nil{
		next:=cur.Next
		cur.Next=pre
		pre=cur
		cur=next
	}
	tail=pre
	//合并两个链表
	curHead:=head
	curTail:=tail
	for i:=0;curTail!=nil;i++{
		
		if i%2==0{
			curHeadNext:=curHead.Next
			curHead.Next=curTail
			curHead=curHeadNext
		}else{
			curTailNext:=curTail.Next
			curTail.Next=curHead
			curTail=curTailNext
		}
	}
}

func main() {
	l5:=&ListNode{
		Val:5,
		Next:nil,
	}

	l4:=&ListNode{
		Val:4,
		Next:l5,
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
	reorderList(l1)
}
