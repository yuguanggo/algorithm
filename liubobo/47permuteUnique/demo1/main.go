package main


func permuteUnique(nums []int) [][]int {
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
		if !visited[i]{
			if i>0&&nums[i]==nums[i-1]&&!visited[i-1]{
				continue
			}
			visited[i]=true
			path=append(path,nums[i])
			helper(nums,path,visited,res)
			path=path[:len(path)-1]
			visited[i]=false
		}
	}
}
func main() {
	permuteUnique([]int{1,1,2})
}
