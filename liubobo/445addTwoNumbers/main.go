package main

import "fmt"

/**
445. 两数相加 II
给定两个非空链表来代表两个非负整数。数字最高位位于链表开始位置。它们的每个节点只存储单个数字。将这两数相加会返回一个新的链表。

你可以假设除了数字 0 之外，这两个数字都不会以零开头。
进阶:
如果输入链表不能修改该如何处理？换句话说，你不能对列表中的节点进行翻转。

示例:

输入: (7 -> 2 -> 4 -> 3) + (5 -> 6 -> 4)
输出: 7 -> 8 -> 0 -> 7

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val int
	Next *ListNode
}



func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	s1:=make([]int,0)
	s2:=make([]int,0)
	cur1:=l1
	for cur1!=nil{
		s1 = append(s1,cur1.Val)
		cur1=cur1.Next
	}
	cur2:=l2
	for cur2!=nil{
		s2= append(s2,cur2.Val)
		cur2=cur2.Next
	}
	carray:=0
	i:=0
	resDummy:=new(ListNode)
	var resCur *ListNode=nil
	for i<len(s1)||i<len(s2){
		sum:=carray
		if i<len(s1){
			sum+=s1[len(s1)-1-i]
		}
		if i<len(s2){
			sum+=s2[len(s2)-1-i]
		}
		carray=sum/10
		resDummy.Next=&ListNode{
			Val:sum%10,
			Next:resCur,
		}
		resCur= resDummy.Next
		i++
	}
	if carray>0{
		resDummy.Next=&ListNode{
			Val:carray,
			Next:resCur,
		}
	}
	return resDummy.Next
}

func main() {
	var arr =[]int{1,3}
	fmt.Println(arr[0:len(arr)-1])
}
