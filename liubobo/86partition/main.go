package main

import "fmt"

/**
86. 分隔链表
给定一个链表和一个特定值 x，对链表进行分隔，使得所有小于 x 的节点都在大于或等于 x 的节点之前。

你应当保留两个分区中每个节点的初始相对位置。

示例:

输入: head = 1->4->3->2->5->2, x = 3
输出: 1->2->2->4->3->5

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/partition-list
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	var less, lesspre, great, greatpre *ListNode = nil, nil, nil, nil
	cur := head
	for cur != nil {
		if cur.Val < x {
			if less == nil {
				less = cur
			} else {
				lesspre.Next = cur
			}
			lesspre = cur
		} else {
			if great == nil {
				great = cur
			} else {
				greatpre.Next = cur
			}
			greatpre = cur
		}
		cur = cur.Next
	}
	if great!=nil{
		greatpre.Next=nil
	}
	if less==nil{
		less=great
	}else {
		lesspre.Next = great
	}
	return less
}

func partition2(head *ListNode, x int) *ListNode  {
	if head==nil{
		return head
	}
	beforeDummy:=new(ListNode)
	afterDummy:=new(ListNode)
	beforeHead:=beforeDummy
	afterHead:=afterDummy
	for head!=nil{
		if head.Val<x{
			beforeHead.Next=head
			beforeHead=beforeHead.Next
		}else{
			afterHead.Next=head
			afterHead=afterHead.Next
		}
		head =head.Next
	}
	afterHead.Next=nil
	beforeHead.Next=afterDummy.Next
	return beforeDummy.Next
}

func main() {
	//l6:=&ListNode{
	//	Val:2,
	//	Next:nil,
	//}
	//l5:=&ListNode{
	//	Val:5,
	//	Next:l6,
	//}
	//l4:=&ListNode{
	//	Val:2,
	//	Next:l5,
	//}
	//l3:=&ListNode{
	//	Val:3,
	//	Next:l4,
	//}

	//l2:=&ListNode{
	//	Val:1,
	//	Next:nil,
	//}
	//l1:=&ListNode{
	//	Val:1,
	//	Next:l2,
	//}
	//partition(l1,0)
	beforeDummy:=new(ListNode)
	fmt.Println(beforeDummy)
}
