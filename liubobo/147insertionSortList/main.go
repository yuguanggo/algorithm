package main

/**
147. 对链表进行插入排序
对链表进行插入排序。


插入排序的动画演示如上。从第一个元素开始，该链表可以被认为已经部分排序（用黑色表示）。
每次迭代时，从输入数据中移除一个元素（用红色表示），并原地将其插入到已排好序的链表中。

 

插入排序算法：

插入排序是迭代的，每次只移动一个元素，直到所有元素可以形成一个有序的输出列表。
每次迭代中，插入排序只从输入数据中移除一个待排序的元素，找到它在序列中适当的位置，并将其插入。
重复直到所有输入数据插入完为止。
 

示例 1：

输入: 4->2->1->3
输出: 1->2->3->4
示例 2：

输入: -1->5->3->4->0
输出: -1->0->3->4->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/insertion-sort-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	dummy:=new(ListNode)
	dummy.Next=head
	pre:=head
	cur:=head.Next
	for cur!=nil&&cur.Next!=nil{
		next:=cur.Next
		if cur.Val<cur.Next.Val{
			pre=cur
			cur=next
			continue
		}
		pre.Next=nil
		//将当前元素插入合适的位置
		incur:=dummy.Next
		inpre:=dummy
		b:=false
		for incur!=nil{
			innext:=incur.Next
			if incur.Val>cur.Val&&!b{
				inpre.Next=cur
				inpre.Next.Next=incur
				b=true
			}
			inpre=incur
			incur=innext
		}
		if !b{
			inpre.Next=cur
			inpre=inpre.Next
		}
		pre=inpre
		cur=next
	}
	return dummy.Next
}

func main() {
	l5:=&ListNode{
		Val:0,
		Next:nil,
	}

	l4:=&ListNode{
		Val:4,
		Next:l5,
	}
	l3:=&ListNode{
		Val:3,
		Next:l4,
	}
	l2:=&ListNode{
		Val:5,
		Next:l3,
	}
	l1:=&ListNode{
		Val:-1,
		Next:l2,
	}
	insertionSortList(l1)
}
