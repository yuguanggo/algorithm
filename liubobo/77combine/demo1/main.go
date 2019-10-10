package main

func combine(n int, k int) [][]int {
	res:=make([][]int,0)
	visited:=make(map[int]bool)
	helper(n,k,[]int{},&res,visited,1)
	return res
}

func helper(n int,k int,path []int,res *[][]int,visited map[int]bool,index int)  {
	if k==len(path){
		p:=make([]int,k)
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=index;i<=n;i++{
		if visited[i]{
			continue
		}
		path=append(path,i)
		visited[i]=true
		helper(n,k,path,res,visited,i+1)
		path=path[:len(path)-1]
		visited[i]=false
	}
}
func main() {
	
}
