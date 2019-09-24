package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	if head==nil||head.Next==nil{
		return head
	}
	less:=&ListNode{}
	lessCur:=less
	great:=&ListNode{}
	greatCur:=great
	cur:=head
	for cur!=nil{
		if cur.Val<x{
			lessCur.Next=cur
			lessCur=cur
		}else{
			greatCur.Next=cur
			greatCur=cur
		}
		cur=cur.Next
	}
	greatCur.Next=nil
	lessCur.Next=great.Next
	return less.Next
}


func main() {
	
}
