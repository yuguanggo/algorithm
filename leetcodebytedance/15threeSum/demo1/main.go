package main

import (
	"fmt"
	"sort"
)

func threeSum(nums []int) [][]int {
	res:=make([][]int,0)
	path:=make([]int,0)
	if len(nums)==0{
		return res
	}
	sort.Ints(nums)
	visited:=make(map[int]bool)
	dfs(nums,0,0,0,path,visited,&res)
	return res
}

func dfs(nums []int,index int,pos int,sum int,path []int,visited map[int]bool,res *[][]int)  {
	if index>=3{
		if sum==0{
			p:=make([]int,len(path))
			copy(p,path)
			*res=append(*res,p)
		}
		return
	}
	for i:=pos;i<len(nums);i++{
		if len(path)==0{
			if !visited[nums[i]]{
				visited[nums[i]]=true
			}else{
				continue
			}
		}
		sum=sum+nums[i]
		path=append(path,nums[i])
		dfs(nums,index+1,i+1,sum,path,visited,res)
		path=path[0:len(path)-1]
		sum=sum-nums[i]
	}
}


func main() {
	nums:=[]int{-1,0,1,2,-1,-4};
	fmt.Println(threeSum(nums))
}
