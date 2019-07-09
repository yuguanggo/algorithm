package main

import "fmt"

/**
2. 两数相加
给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/add-two-numbers
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

type ListNode struct {
	Val  int
	Next *ListNode
}


func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	resDummy:=new(ListNode)
	resHead:=resDummy
	l1cur:=l1
	l2cur:=l2
	sum:=0
	carry:=0
	for l1cur!=nil||l2cur!=nil{
		if l1cur!=nil&&l2cur!=nil{
			sum=l1cur.Val+l2cur.Val+carry
			l1cur=l1cur.Next
			l2cur=l2cur.Next
		}else if l1cur!=nil{
			sum=l1cur.Val+carry
			l1cur=l1cur.Next
		}else{
			sum=l2cur.Val+carry
			l2cur= l2cur.Next
		}
		carry=sum/10
		resHead.Next=&ListNode{
			Val:sum%10,
			Next:nil,
		}
		resHead = resHead.Next
	}
	if carry!=0{
		resHead.Next=&ListNode{
			Val:carry,
			Next:nil,
		}
	}
	return resDummy.Next
}
func main() {
	p:=new([]int)
	fmt.Println(*p==nil)
	*p=append(*p,1)
	fmt.Println(*p)
	v:=make([]int,0)
	fmt.Println(v==nil)

}
