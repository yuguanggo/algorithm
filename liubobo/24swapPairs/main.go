package main


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
	dummy:=new(ListNode)
	cur:=head
	pre:=dummy
	for cur!=nil&&cur.Next!=nil{
		next:=cur.Next.Next
		pre.Next=cur.Next
		pre = cur
		cur.Next=cur
		cur=next
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
	swapPairs(l1)
}
