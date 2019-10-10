package main

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	res:=make([][]int,0)
	sort.Ints(candidates)
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
	for i:=index;i<len(candidates);i++{
		if target>candidates[i]{
			break
		}
		if i>index&&candidates[i]==candidates[i-1]{
			continue
		}
		path=append(path,candidates[i])
		helper(candidates,target-candidates[i],path,res,i+1)
		path=path[:len(path)-1]
	}
}
func main() {
	
}
