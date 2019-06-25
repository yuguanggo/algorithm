package main

/**
相交链表
编写一个程序，找到两个单链表相交的起始节点。

 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	var cA, cB *ListNode
	cA = headA
	cB = headB
	for cA != cB {
		if cA == nil {
			cA = headB
		} else {
			cA = cA.Next
		}
		if cB == nil {
			cB = headA
		} else {
			cB = cB.Next
		}
	}
	return cA
}

func main() {

}
