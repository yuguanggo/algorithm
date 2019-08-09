package main

import "fmt"

/**
91. 解码方法
一条包含字母 A-Z 的消息通过以下方式进行了编码：

'A' -> 1
'B' -> 2
...
'Z' -> 26
给定一个只包含数字的非空字符串，请计算解码方法的总数。

示例 1:

输入: "12"
输出: 2
解释: 它可以解码为 "AB"（1 2）或者 "L"（12）。
示例 2:

输入: "226"
输出: 3
解释: 它可以解码为 "BZ" (2 26), "VF" (22 6), 或者 "BBF" (2 2 6) 。


              226
            2   22
        26        6
      2  26
     6
d[3]=1
d[2]=d[3]+1
来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/decode-ways
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */

func numDecodings(s string) int {
	n:=len(s)
	//dp[i]
	dp:=make([]int,n+1)
	dp[0]=1

	for i:=1;i<n;i++{
		c:=s[i]
		if c!='0'{
			dp[i]=dp[i]+dp[i-1]
		}
		if i>=1{
			c1:=int(s[i]-'0')*10+int(s[i-1]-'0')
			if c1>9&&c1<=26{
				if i==1{
					dp[i]=1+dp[i]
				}else{
					dp[i]=dp[i]+dp[i-2]
				}
			}

		}
	}
	return dp[n-1]
}

func main() {
	b:='1'
	fmt.Println(int(b-'0'))
}
