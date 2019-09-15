package main


/**
85. 最大矩形
给定一个仅包含 0 和 1 的二维二进制矩阵，找出只包含 1 的最大矩形，并返回其面积。

示例:

输入:
[
  ["1","0","1","0","0"],
  ["1","0","1","1","1"],
  ["1","1","1","1","1"],
  ["1","0","0","1","0"]
]
输出: 6

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/maximal-rectangle
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func maximalRectangle(matrix [][]byte) int {
	m:=len(matrix)
	if m==0{
		return 0
	}
	maxArea:=0
	n:=len(matrix[0])
	dp:=make([]int,n)
	for i:=0;i<m;i++{
		for j:=0;j<n;j++{
			if matrix[i][j]==1{
				dp[j]=dp[j]+1
			}else{
				dp[j]=0
			}
		}
		maxArea=max(largestRectangleArea(dp),maxArea)
	}
	return maxArea
}

func largestRectangleArea(heights []int) int {
	n:=len(heights)
	stack:=make([]int,0)
	stack=append(stack,-1)
	res:=0
	for i:=0;i<n;i++{
		top:=stack[len(stack)-1]
		for top!=-1&&heights[top]>=heights[i]{
			//一直弹出，直到遇到栈顶元素>=当前元素
			stack=stack[0:len(stack)-1]
			area:=heights[top]*(i-stack[len(stack)-1]-1)
			res=max(res,area)
			top=stack[len(stack)-1]
		}
		stack=append(stack,i)
	}
	for stack[len(stack)-1]!=-1{
		top:=stack[len(stack)-1]
		stack=stack[0:len(stack)-1]
		area:=heights[top]*(n-stack[len(stack)-1]-1)
		res=max(res,area)
	}
	return res
}

func max(i,j int) int  {
	if i>j{
		return i
	}
	return j
}

func main() {
	
}
