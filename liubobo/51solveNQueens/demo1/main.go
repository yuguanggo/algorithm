package main


func solveNQueens(n int) [][]string {
	col:=make([]bool,n)//存放列占用的情况
	dia1:=make([]bool,2*n) //存放对角线1占用的情况
	dia2:=make([]bool,2*n) //存放对角线2占用的情况
	res:=make([][]string,0)
	putQueen(n,0,col,dia1,dia2,[]int{},&res)
	return res
}

//在第index行放皇后
func putQueen(n int,index int,col,dia1,dia2 []bool,path []int,res *[][]string)  {
	if index==n{
		*res=append(*res,generateBord(n,path))
	}
	for i:=0;i<n;i++{
		if !col[i]&&!dia1[i+index]&&!dia2[i-index+n]{
			col[i]=true
			dia1[i+index]=true
			dia2[i-index+n]=true
			path=append(path,i)
			putQueen(n,index+1,col,dia1,dia2,path,res)
			path=path[:len(path)-1]
			dia2[i-index+n]=false
			dia1[i+index]=false
			col[i]=false
		}
	}
}

func generateBord(n int,path []int) []string  {
	res:=make([]string,0)
	for i:=0;i<n;i++{
		s:=make([]byte,0)
		for j:=0;j<n;j++{
			if j!=path[i]{
				s=append(s,'.')
			}else{
				s=append(s,'Q')
			}
		}
		res=append(res,string(s))
	}
	return res
}

func main() {
	
}
