package main


func intersection(nums1 []int, nums2 []int) []int {
	set:=make(map[int]bool)
	for i:=0;i<len(nums1);i++{
		set[nums1[i]]=true
	}
	res:=make([]int,0)
	for i:=0;i<len(nums2);i++{
		if set[nums2[i]]{
			set[nums2[i]]=false
			res=append(res,nums2[i])
		}
	}
	return res
}


func main() {
	
}
