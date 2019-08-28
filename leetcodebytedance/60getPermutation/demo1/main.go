package main

import (
	"fmt"
)

/**
                1,2,3,4     9
				1
      1         2         3      4

 */
func getPermutation(n int, k int) string {
	return ""
}

func helper(n int,k int,level int,path []int)  {
	for i:=1;i<=n;i++{
		if i*f(n-level)<k{
			continue
		}
		path=append(path,i)
		k=k-(i-1)*f(n-level)
		helper(n,k,level+1,path)

	}
}

func f(n int) int  {
	res:=1
	for i:=1;i<=n;i++{
		res*=i
	}
	return res
}


func main() {
	fmt.Println(f(3))
	fmt.Println(f(4))
}
