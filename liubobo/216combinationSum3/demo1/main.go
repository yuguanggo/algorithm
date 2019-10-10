package main

func combinationSum3(k int, n int) [][]int {
	res:=make([][]int,0)
	helper(k,n,1,[]int{},&res)
	return res
}

func helper(k,n int,index int,path []int,res *[][]int)  {
	if k==len(path)&&n==0{
		p:=make([]int,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	if k<len(path)||n<0{
		return
	}
	for i:=index;i<=9;i++{
		path=append(path,i)
		helper(k,n-i,i+1,path,res)
		path=path[:len(path)-1]
	}
}

func main() {
	
}
