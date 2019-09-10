package main

import "fmt"

/**
滑动窗口通用逻辑 needs
1.初始化[left,rihgt]
2.不断增加right 知道窗口里的内容满足所有条件
3.增加left 更新结果。直到窗口不满足条件
4.重复2.直到right到达结尾
 */
func minWindow(s string, t string) string {
	l,r:=0,0
	n:=len(s)
	needs:=make(map[byte]int)
	if len(t)==0{
		return ""
	}
	for i:=0;i<len(t);i++{
		needs[t[i]]++
	}
	windows:=make(map[byte]int)
	match:=0
	minlen:=n+1
	minl,minr:=0,0
	for r<n{
		windows[s[r]]++
		if v,ok:=needs[s[r]];ok{
			if v==windows[s[r]]{
				match++
			}
		}

		for l<n&&match==len(needs){
			if r-l+1<minlen{
				minlen=r-l+1
				minl=l
				minr=r
			}
			windows[s[l]]--
			if v,ok:=needs[s[l]];ok{
				if v>windows[s[l]]{
					match--
				}
			}
			l++
		}
		r++
	}
	if minlen==n+1{
		return ""
	}
	return s[minl:minr+1]
}


func isContains(windows,needs map[byte]int) bool {
	for k,v:=range needs{
		if windows[k]<v{
			return false
		}
	}
	return true
}

func main() {
	s:="aa"
	t:="aa"
	fmt.Println(minWindow(s,t))
}
