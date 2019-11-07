package main


type ListNode struct {
	Val  int
	Next *ListNode
}



func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists,0,len(lists)-1)
}

func merge(lists []*ListNode,l,r int) *ListNode {
	if l>r{
		return nil
	}
	if l==r{
		return lists[l]
	}
	mid:=l+(r-l)/2
	l1:=merge(lists,l,mid)
	l2:=merge(lists,mid+1,r)
	return helper(l1,l2)
}

func helper(l1,l2 *ListNode) *ListNode  {
	if l1==nil{
		return l2
	}
	if l2==nil{
		return l1
	}
	if l1.Val<l2.Val{
		l1.Next=helper(l1.Next,l2)
		return l1
	}else{
		l2.Next=helper(l1,l2.Next)
		return l2
	}
}


func main() {
	
}
