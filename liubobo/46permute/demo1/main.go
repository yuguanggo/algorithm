package main


func permute(nums []int) [][]int {
	res:=make([][]int,0)
	visited:=make(map[int]bool)
	helper(nums,[]int{},visited,&res)
	return res
}

func helper(nums []int,path []int,visited map[int]bool,res *[][]int)  {
 	if len(path)==len(nums){
 		p:=make([]int,len(path))
 		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=0;i<len(nums);i++{
		if visited[nums[i]]{
			continue
		}
		path=append(path,nums[i])
		visited[nums[i]]=true
		helper(nums,path,visited,res)
		path=path[:len(path)-1]
		visited[nums[i]]=false
	}
}


func main() {
	
}
