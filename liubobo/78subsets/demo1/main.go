package main


func subsets(nums []int) [][]int {
	res:=make([][]int,0)
	helper(nums,[]int{},&res,0)
	return res
}

func helper(nums []int,path []int,res *[][]int,index int)  {
	p:=make([]int,len(path))
	copy(p,path)
	*res=append(*res,p)
	if index>=len(nums){
		return
	}
	for i:=index;i<len(nums);i++{
		path=append(path,nums[i])
		helper(nums,path,res,i+1)
		path=path[:len(path)-1]
	}
}


func main() {
	
}
