package main


/**
203. 移除链表元素
删除链表中等于给定值 val 的所有节点。

示例:

输入: 1->2->6->3->4->5->6, val = 6
输出: 1->2->3->4->5
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	dummy:=new(ListNode)
	dummy.Next=head
	pre:=dummy
	cur:=head
	for cur!=nil{
		if cur.Val==val{
			pre.Next=cur.Next
		}else{
			pre=cur
		}
		cur=cur.Next
	}
	return dummy.Next
}

func main() {
	
}
