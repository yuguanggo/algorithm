package main


/**
234. 回文链表
请判断一个链表是否为回文链表。

示例 1:

输入: 1->2
输出: false
示例 2:

输入: 1->2->2->1
输出: true
进阶：
你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/palindrome-linked-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	if head==nil||head.Next==nil{
		return true
	}

	//拆分链表加，反转前一部分
	var pre *ListNode=nil
	slow:=head
 	fast:=head
 	for fast!=nil&&fast.Next!=nil{
		fast=fast.Next.Next
 		next:=slow.Next
 		slow.Next=pre
 		pre=slow
 		slow=next
	}
	head=pre
	tail:=slow
	if fast!=nil{
		tail=tail.Next
	}
	//遍历
	curh :=head
	curt:=tail
	for curh!=nil{
		if curh.Val!=curt.Val{
			return false
		}
		curh=curh.Next
		curt=curt.Next
	}
	return true
}
func main() {
	l3:=&ListNode{
		Val:1,
		Next:nil,
	}
	l2:=&ListNode{
		Val:0,
		Next:l3,
	}
	l1:=&ListNode{
		Val:1,
		Next:l2,
	}
	isPalindrome(l1)
}
