package main

/**
合并K个排序链表
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。
输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return helper(lists, 0, len(lists)-1)
}

func helper(lists []*ListNode, left int, right int) *ListNode {
	if left > right {
		return nil
	}
	if left == right {
		return lists[left]
	}
	mid := (left + right) / 2
	leftNode := helper(lists, left, mid)
	rightNode := helper(lists, mid+1, right)
	return mergeTwoLists(leftNode, rightNode)
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

func main() {

}
