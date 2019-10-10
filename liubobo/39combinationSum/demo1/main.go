package main


func combinationSum(candidates []int, target int) [][]int {
	res:=make([][]int,0)
	helper(candidates,target,[]int{},&res,0)
	return res
}

func helper(candidates []int,target int,path []int,res *[][]int,index int)  {
	if target==0{
		p:=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	if target<0{
		return
	}
	for i:=index;i<len(candidates);i++{
		path=append(path,candidates[i])
		helper(candidates,target-candidates[i],path,res,i)
		path=path[:len(path)-1]
	}
}





func main() {
	
}
