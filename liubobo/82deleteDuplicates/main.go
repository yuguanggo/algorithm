package main

/**
82. 删除排序链表中的重复元素 II
给定一个排序链表，删除所有含有重复数字的节点，只保留原始链表中 没有重复出现 的数字。

示例 1:

输入: 1->2->3->3->4->4->5
输出: 1->2->5
示例 2:

输入: 1->1->1->2->3
输出: 2->3

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val int
	Next *ListNode
}
func deleteDuplicates(head *ListNode) *ListNode {
	dummy:=new(ListNode)
	dummy.Next=head
	cur:=head
	pre:=dummy
	record:=make(map[int]int)
	for cur!=nil{
		if cur.Next!=nil&&cur.Val==cur.Next.Val{
			record[cur.Val]=cur.Val
		}
		if _,ok:=record[cur.Val];!ok{
			pre.Next=cur
			pre=cur
		}
		cur=cur.Next
	}
	pre.Next=cur
	return dummy.Next
}

func main() {

	l7:=&ListNode{
		Val:5,
		Next:nil,
	}
	l6:=&ListNode{
		Val:4,
		Next:l7,
	}
	l5:=&ListNode{
		Val:4,
		Next:l6,
	}
	l4:=&ListNode{
		Val:3,
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
	deleteDuplicates(l1)
}
