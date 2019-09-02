package main

/**

 */
func minimumTotal(triangle [][]int) int {
	n:=len(triangle)
	dp:=make([]int,n)
	for i:=n-1;i>=0;i--{
		for j:=0;j<len(triangle[i]);j++{
			if i==n-1{
				dp[j]=triangle[i][j]
			}else{
				dp[j]=min(dp[j]+triangle[i][j],dp[j+1]+triangle[i][j])
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
