package main



/**
21. 合并两个有序链表

将两个有序链表合并为一个新的有序链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。 

示例：

输入：1->2->4, 1->3->4
输出：1->1->2->3->4->4

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-two-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	cur1:=l1
	cur2:=l2
	dummy:=new(ListNode)
	cur:=dummy
	for cur1!=nil||cur2!=nil{
		if cur1==nil{
			cur.Next=cur2
			return dummy.Next
		}
		if cur2==nil{
			cur.Next=cur1
			return dummy.Next
		}
		if cur1.Val<cur2.Val{
			cur.Next=cur1
			cur1=cur1.Next
		}else{
			cur.Next=cur2
			cur2=cur2.Next
		}
		cur=cur.Next
	}
	return dummy.Next
}

func mergeTwoLists2(l1 *ListNode, l2 *ListNode) *ListNode  {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=mergeTwoLists(l1.Next,l2)
		return l1
	}else{
		l2.Next=mergeTwoLists(l2.Next,l1)
		return l2
	}
}
func main() {
	
}
