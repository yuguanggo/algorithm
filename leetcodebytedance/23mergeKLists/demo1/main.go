package main


type ListNode struct {
	Val  int
	Next *ListNode
}


func mergeKLists(lists []*ListNode) *ListNode {
	return helper(lists,0,len(lists)-1)
}

func helper(lists []*ListNode,l,r int) *ListNode  {
	if l>r{
		return nil
	}
	if l==r{
		return lists[l]
	}
	mid:=l+(r-l)/2
	lnode:=helper(lists,l,mid)
	rnode:=helper(lists,mid+1,r)
	return merge(lnode,rnode)
}

func merge(l1 *ListNode,l2 *ListNode) *ListNode  {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=merge(l1.Next,l2)
		return l1
	}else{
		l2.Next=merge(l1,l2.Next)
		return l2
	}
}

func main() {
	
}
