package main


/**
合并 k 个排序链表，返回合并后的排序链表。请分析和描述算法的复杂度。

示例:

输入:
[
  1->4->5,
  1->3->4,
  2->6
]
输出: 1->1->2->3->4->4->5->6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/merge-k-sorted-lists
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return helper(lists,0,len(lists)-1)
}

func helper(lists []*ListNode,l,r int) *ListNode {
	if l>r{
		return nil
	}
	if l==r{
		return lists[l]
	}
	mid:=l+(r-l)/2
	lnode:= helper(lists,l,mid)
	rnode:=helper(lists,mid+1,r)
	return mergeList(lnode,rnode)
}

func mergeList(l *ListNode,r *ListNode) *ListNode  {
	if l==nil{
		return r
	}
	if r==nil{
		return l
	}
	if l.Val<r.Val{
		l.Next=mergeList(l.Next,r)
		return l
	}else{
		r.Next=mergeList(r.Next,l)
		return r
	}
}

func main() {

}
