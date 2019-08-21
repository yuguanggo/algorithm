package main

/**
148. 排序链表
在 O(n log n) 时间复杂度和常数级空间复杂度下，对链表进行排序。

示例 1:

输入: 4->2->1->3
输出: 1->2->3->4
示例 2:

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

 */
type ListNode struct {
	Val  int
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
		pre=slow
		slow=slow.Next
		fast=fast.Next.Next
	}
	pre.Next=nil
	return merge(sortList(head),sortList(slow))
}

func merge(l *ListNode,r *ListNode) *ListNode  {
	if l==nil{
		return r
	}
	if r==nil{
		return l
	}
	if l.Val<r.Val{
		l.Next=merge(l.Next,r)
		return l
	}else{
		r.Next=merge(l,r.Next)
		return r
	}
}
func main() {
	
}
