package main

import (
	"fmt"
)

/**
无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。

示例 1:

输入: "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
示例 2:

输入: "bbbbb"
输出: 1
解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
示例 3:

输入: "pwwkew"
输出: 3
解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。


 */
func lengthOfLongestSubstring(s string) int {
	//s[l...r]为滑动窗口
	l,r:=0,-1
	n:=len(s)
	freq:=[256]int{}
	res:=0
	for l<n{
		if r+1<n&&freq[s[r+1]]==0{
			freq[s[r+1]]++
			r++
		}else{
			freq[s[l]]--
			l++
		}
		res = max(res,r-l+1)
	}
	return res
}

func max(i,j int) int  {
	if i<j{
		return j
	}
	return i
}
func main() {
	s:="abcabcbb"
	fmt.Println(lengthOfLongestSubstring(s))
}
