package main

import (
	"fmt"
)

/**
32. 最长有效括号
给定一个只包含 '(' 和 ')' 的字符串，找出最长的包含有效括号的子串的长度。

示例 1:

输入: "(()"
输出: 2
解释: 最长有效括号子串为 "()"
示例 2:

输入: ")()())"
输出: 4
解释: 最长有效括号子串为 "()()"

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/longest-valid-parentheses
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。

dp[i]=max(dp[i-1],dp[i-2])
 */

func longestValidParentheses2(s string) int {
	n:=len(s)
	if n==0||n==1{
		return 0
	}
	dp:=make([]int,n)
	dp[0]=0
	if s[0]=='('&&s[1]==')'{
		dp[1]=2
	}else{
		dp[1]=0
	}
	for i:=2;i<n;i++{
		if s[i]==')'&&s[i-1]=='('{
			dp[i]=dp[i-1]+2
		}else{
			dp[i]=dp[i-1]
		}
	}
	return dp[n-1]
}
func longestValidParentheses3(s string) int {
	maxans:=0
	n:=len(s)
	dp:=make([]int,n)
	for i:=1;i<n;i++{
		if s[i]==')'{
			if s[i-1]=='('{
				if i>=2{
					dp[i]=dp[i-2]+2
				}else{
					dp[i]=2
				}
			}else if (i-dp[i-1])>0&&s[i-dp[i-1]-1]=='('{
				if i-dp[i-1]>=2{
					dp[i]=dp[i-1]+dp[i-dp[i-1]-2]+2
				}else{
					dp[i]=dp[i-1]+2
				}

			}
			maxans=max(maxans,dp[i])
		}
	}
	return maxans
}

func longestValidParentheses(s string) int {
	//栈
	n:=len(s)
	stack:=make([]int,0)
	stack=append(stack,-1)
	i:=0
	res:=0
	for i<n{
		if s[i]=='('{
			stack=append(stack,i)
		}else{
			stack=stack[:len(stack)-1]
			if len(stack)>0{
				res=max(res,i-stack[len(stack)-1])
			}else{
				stack=append(stack,i)
			}
		}
		i++
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
	fmt.Println(longestValidParentheses(")()())"))
}
