package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertionSortList(head *ListNode) *ListNode {
	dummy:=&ListNode{}
	dummy.Next=head
	pre:=dummy
	cur:=head
	for cur!=nil{
		node:=cur.Next
		if node!=nil&&node.Val<cur.Val{
			for pre.Next!=nil&&pre.Next.Val<node.Val{
				pre=pre.Next
			}
			//找到插入的位置
			temp:=pre.Next
			pre.Next=node
			cur.Next=node.Next
			node.Next=temp
			pre=dummy
		}else{
			cur=node
		}
	}
	return dummy.Next
}


func main() {
	p1:=new(ListNode)
	p2:=&ListNode{}
	fmt.Println(p1)
	fmt.Println(p2)
}
