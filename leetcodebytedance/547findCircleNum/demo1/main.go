package main

func findCircleNum(M [][]int) int {
	n:=len(M)
	if n==0{
		return 0
	}
	visited:=make([]bool,n)
	res:=0
	for i:=0;i<n;i++{
		if visited[i]{
			continue
		}
		visited[i]=true
		res++
		dfs(M,n,i,visited)
	}
	return res
}

func dfs(M [][]int,n int,k int,visited []bool)  {

	for i:=0;i<n;i++{
		if visited[i]||M[k][i]==0{
			continue
		}
		visited[i]=true
		dfs(M,n,i,visited)
	}
}

func findCircleNum2(M [][]int) int {
	n:=len(M)
	if n==0{
		return 0
	}
	res:=0
	visited:=make([]bool,n)
	for i:=0;i<n;i++{
		if visited[i]{
			continue
		}
		q:=make([]int,0)
		q=append(q,i)
		for len(q)>0{
			m:=q[len(q)-1]
			q=q[0:len(q)-1]
			for j:=0;j<n;j++{
				if visited[j]||M[m][j]==0{
					continue
				}

				q=append(q,j)
				visited[j]=true
			}
		}
		res++
	}
	return res
}

func main() {
	
}
