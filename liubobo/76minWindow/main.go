package main

import (
	"fmt"
)

/**
76. 最小覆盖子串

给你一个字符串 S、一个字符串 T，请在字符串 S 里面找出：包含 T 所有字母的最小子串。

示例：

输入: S = "ADOBECODEBANC", T = "ABC"
输出: "BANC"
说明：

如果 S 中不存这样的子串，则返回空字符串 ""。
如果 S 中存在这样的子串，我们保证它是唯一的答案。

 */
func minWindow(s string, t string) string {
	//s[l..r]窗口
	l,r:=0,-1
	needs:=make(map[byte]int)
	windows:=make(map[byte]int)
	m,n:=len(s),len(t)
	res:=m+1
	for i:=0;i<n;i++{
		needs[t[i]]++
	}
	minl,minr:=0,0

	for r+1<m{
		r++
		windows[s[r]]++
		for isEqual(windows,needs){
			if r-l+1<res{
				res = r-l+1
				minl=l
				minr=r
			}
			windows[s[l]]--
			l++
		}
	}
	if res==m+1{
		return ""
	}
	return s[minl:minr+1]
}

func minWindow2(s string, t string) string {
	//s[l..r]窗口
	l,r:=0,-1
	needs:=make(map[byte]int)
	windows:=make(map[byte]int)
	m,n:=len(s),len(t)
	res:=m+1
	for i:=0;i<n;i++{
		needs[t[i]]++
	}
	minl,minr:=0,0
	match:=0
	for r+1<m{
		r++
		if _,ok:=needs[s[r]];ok{
			windows[s[r]]++
			if windows[s[r]]==needs[s[r]]{
				match++
			}
		}
		for match==len(needs){
			if r-l+1<res{
				minl=l
				minr=r
				res=minr-minl+1
			}

			if _,ok:=needs[s[l]];ok{
				windows[s[l]]--
				if windows[s[l]]<needs[s[l]]{
					match--
				}
			}
			l++
		}

	}
	if res==m+1{
		return ""
	}
	return s[minl:minr+1]
}

func minWindow3(s string, t string) string {
	//s[l..r]窗口
	l,r:=0,0
	needs:=[128]int{}
	windows:=[128]int{}
	m,n:=len(s),len(t)
	min:=m+1
	for i:=0;i<n;i++{
		needs[t[i]]++
	}
	minl,minr:=0,0
	match:=0
	for r<m{
		if windows[s[r]]<needs[s[r]]{
			match++
		}
		windows[s[r]]++
		for l<=r&&windows[s[l]]>needs[s[l]]{
			windows[s[l]]--
			l++
		}
		width:=r-l+1
		if match==n&&min>width{
			minl=l
			minr=r
			min = width
		}
		r++

	}
	if min==m+1{
		return ""
	}
	return s[minl:minr+1]
}

func isEqual(a,b map[byte]int) bool  {
	for k,v:=range b{
		if a[k]<v||a[k]==0{
			return false
		}
	}
	return true
}

func main() {
	s:="aab"
	t:="aab"
	fmt.Println(minWindow3(s,t))
}
