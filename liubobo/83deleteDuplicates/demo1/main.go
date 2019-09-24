package main


type ListNode struct {
	Val int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	record:=make(map[int]bool)
	pre:=&ListNode{}
	cur:=head
	for cur!=nil{
		next:=cur.Next
		if record[cur.Val]{
			pre.Next=next
		}else{
			record[cur.Val]=true
			pre=cur
		}
		cur=next
	}
	return head
}


func main() {
	
}
