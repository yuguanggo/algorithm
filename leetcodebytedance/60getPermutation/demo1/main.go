package main

import (
	"fmt"
	"strconv"
)

/**
                1,2,3,4     9
				1
      1         2         3      4

 */
func getPermutation(n int, k int) string {
	used:=make(map[int]bool)
	return helper(n,k,0,[]int{},used)
}

func helper(n int,k int,level int,path []int,used map[int]bool) string  {
	if level==n{
		res:=""
		for j:=0;j<len(path);j++{
			res+=strconv.Itoa(path[j])
		}
		return res
	}
	ps:=f(n-1-level)
	for i:=1;i<=n;i++{
		if used[i]{
			continue
		}
		if ps<k{
			k-=ps
			continue
		}
		path=append(path,i)
		used[i]=true
		return helper(n,k,level+1,path,used)

	}
	return ""
}

func f(n int) int  {
	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	return res
}


func main() {

	fmt.Println(getPermutation(3,3))

}
