package main

/**
环形链表 II
给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

为了表示给定链表中的环，我们使用整数 pos 来表示链表尾连接到链表中的位置（索引从 0 开始）。 如果 pos 是 -1，则在该链表中没有环。

说明：不允许修改给定的链表。

示例 1：

输入：head = [3,2,0,-4], pos = 1
输出：tail connects to node index 1
解释：链表中有一个环，其尾部连接到第二个节点。
输入：head = [1,2], pos = 0
输出：tail connects to node index 0
解释：链表中有一个环，其尾部连接到第一个节点。
示例 3：

输入：head = [1], pos = -1
输出：no cycle
解释：链表中没有环。

进阶：
你是否可以不用额外空间解决此题？

如下图所示，X,Y,Z分别为链表起始位置，环开始位置和两指针相遇位置，则根据快指针速度为慢指针速度的两倍，可以得出：
2*(a + b) = a + b + n * (b + c)；即
a=(n - 1) * b + n * c = (n - 1)(b + c) +c;
注意到b+c恰好为环的长度，故可以推出，如将此时两指针分别放在起始位置和相遇位置，
并以相同速度前进，当一个指针走完距离a时，
另一个指针恰好走出 绕环n-1圈加上c的距离。

https://blog.csdn.net/Eartha1995/article/details/80990636
 */
/**
* Definition for singly-linked list.
* type ListNode struct {
*     Val int
*     Next *ListNode
* }
*/

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycle(head *ListNode) *ListNode {
	var slow, fast *ListNode
	slow = head
	fast = head
	for fast!=nil&&fast.Next!=nil{
		slow=slow.Next
		fast=fast.Next.Next
		if slow==fast{
			fast=head
			for slow!=fast{
				fast=fast.Next
				slow=slow.Next
			}
			return slow
		}
	}
	return nil
}
func main() {

}
