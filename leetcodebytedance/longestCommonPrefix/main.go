package main

import (
	"fmt"
	"strings"
)

/**
编写一个函数来查找字符串数组中的最长公共前缀。

如果不存在公共前缀，返回空字符串 ""。

示例 1:

输入: ["flower","flow","flight"]
输出: "fl"
示例 2:

输入: ["dog","racecar","car"]
输出: ""
解释: 输入不存在公共前缀。
 */
func longestCommonPrefix(strs []string) string {
	if len(strs)==0{
		return ""
	}
	prefix := strs[0]
	for i:=1;i<len(strs);i++  {
		for strings.Index(strs[i],prefix)!=0{
			prefix = prefix[0:len(prefix)-1]
			if len(prefix)==0{
				return ""
			}
		}
	}
	return prefix

}

func main() {
	s:=[]string{"dog","racecar","car"}
	pos:=longestCommonPrefix(s)
	fmt.Println(pos)
	m:="recsss"
	fmt.Println(m[1:1])
}
