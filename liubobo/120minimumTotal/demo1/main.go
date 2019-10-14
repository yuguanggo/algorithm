package main

func minimumTotal(triangle [][]int) int {
	row:=len(triangle)
	if row==0{
		return 0
	}
	dp:=make([]int,row)
	for i:=row-1;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			if i==row-1{
				//初始化最后一行的数据
				dp[j]=triangle[i][j]
			}else{
				dp[j]=min(dp[j],dp[j+1])+triangle[i][j]
			}
		}
	}
	return dp[0]
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}


func main() {
	
}
