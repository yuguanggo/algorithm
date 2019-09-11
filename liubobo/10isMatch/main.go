package main

/**
10. 正则表达式匹配
给你一个字符串 s 和一个字符规律 p，请你来实现一个支持 '.' 和 '*' 的正则表达式匹配。

'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖 整个 字符串 s的，而不是部分字符串。

说明:

s 可能为空，且只包含从 a-z 的小写字母。
p 可能为空，且只包含从 a-z 的小写字母，以及字符 . 和 *。
示例 1:

输入:
s = "aa"
p = "a"
输出: false
解释: "a" 无法匹配 "aa" 整个字符串。
示例 2:

输入:
s = "aa"
p = "a*"
输出: true
解释: 因为 '*' 代表可以匹配零个或多个前面的那一个元素, 在这里前面的元素就是 'a'。因此，字符串 "aa" 可被视为 'a' 重复了一次。
示例 3:

输入:
s = "ab"
p = ".*"
输出: true
解释: ".*" 表示可匹配零个或多个（'*'）任意字符（'.'）。
示例 4:

输入:
s = "aab"
p = "c*a*b"
输出: true
解释: 因为 '*' 表示零个或多个，这里 'c' 为 0 个, 'a' 被重复一次。因此可以匹配字符串 "aab"。
示例 5:

输入:
s = "mississippi"
p = "mis*is*p*."
输出: false

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/regular-expression-matching
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */


func isMatch2(s string, p string) bool {
	//递归
	if len(p)==0{
		return len(s)==0
	}
	first:=false
	if len(s)>0&&(s[0]==p[0]||p[0]=='.'){
		first=true
	}
	if len(p)>=2&&p[1]=='*'{
		return isMatch(s,p[2:])||(first&&isMatch(s[1:],p))
	}else{
		return first&&isMatch(s[1:],p[1:])
	}
}

func isMatch(s string, p string) bool {
	memo:=make(map[int]map[int]bool)
	return dp(0,0,s,p,&memo)
}

func dp(i,j int,s,p string,memo *map[int]map[int]bool) bool  {
	var ans,first bool
	if _,ok:=(*memo)[i];ok{
		if _,ok:=(*memo)[i][j];ok{
			return (*memo)[i][j]
		}
	}else{
		(*memo)[i]=make(map[int]bool)
	}
	if j==len(p){
		return i==len(s)
	}

	if i<len(s)&&(s[i]==p[j]||p[j]=='.'){
		first=true
	}

	if j<=(len(p)-2)&&p[j+1]=='*'{
		ans=dp(i,j+2,s,p,memo)||first&&dp(i+1,j,s,p,memo)
	}else{
		ans= first&&dp(i+1,j+1,s,p,memo)
	}
	(*memo)[i][j]=ans
	return ans
}


func main() {
	memo:=make(map[int]map[int]bool)
	memo[0][0]=true
	//fmt.Println(isMatch("aa","a"))
}
