package main


/**
92. 反转链表 II
反转从位置 m 到 n 的链表。请使用一趟扫描完成反转。

说明:
1 ≤ m ≤ n ≤ 链表长度。

示例:

输入: 1->2->3->4->5->NULL, m = 2, n = 4
输出: 1->4->3->2->5->NULL

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/reverse-linked-list-ii
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

func reverseBetween(head *ListNode, m int, n int) *ListNode {
	if m==n{
		return head
	}
	var pre *ListNode=nil
	var mpre *ListNode=nil
	var ml *ListNode=nil
	var nnext *ListNode=nil
	cur:=head
	i:=1
	for i<=n{
		next:=cur.Next
		if i==m{
			//将m位置的值保存下来
			mpre = pre
			ml=cur
		}
		if i==n{
			//遍历到n了,将n的下一个节点保存下来
			nnext=cur.Next
			cur.Next=pre
			if m==1{
				head=cur
			}else{
				mpre.Next=cur
			}
			ml.Next=nnext
			break
		}
		if i>m{
			//开始反转
			cur.Next=pre
			pre=cur
			cur=next
		}else{
			pre=cur
			cur=next
		}
		i++
	}
	return head
}
func main() {
	//l5:=&ListNode{
	//	Val:5,
	//	Next:nil,
	//}
	//l4:=&ListNode{
	//	Val:4,
	//	Next:l5,
	//}
	l3:=&ListNode{
		Val:3,
		Next:nil,
	}
	l2:=&ListNode{
		Val:2,
		Next:l3,
	}
	l1:=&ListNode{
		Val:1,
		Next:l2,
	}
	//reverseBetween(l4,1,1)
	reverseBetween(l1,1,2)
}
