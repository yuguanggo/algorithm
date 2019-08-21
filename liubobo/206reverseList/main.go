package main

/**
206. 反转链表
反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


/**
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

func reverseList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	p:=reverseList(head.Next)
	head.Next.Next = head
	head.Next=nil
	return p
}

func reverseList2(head *ListNode) *ListNode  {
	var pre *ListNode=nil
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
	
}
