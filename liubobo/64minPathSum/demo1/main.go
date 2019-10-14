package main


func minPathSum(grid [][]int) int {
	m:=len(grid)
	if m==0{
		return 0
	}
	n:=len(grid[0])
	dp:=make([]int,n)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if i==0{
				if j==0{
					dp[j]=grid[i][j]
				}else{
					dp[j]=dp[j-1]+grid[i][j]
				}
			}else{
				if j==0{
					dp[j]=dp[j]+grid[i][j]
				}else{
					dp[j]=min(dp[j],dp[j-1])+grid[i][j]
				}
			}

		}
	}
	return dp[n-1]
}

func min(i,j int) int  {
	if i>j{
		return j
	}
	return i
}


func main() {
	
}
