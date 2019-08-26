package main

import (
	"fmt"
)

func lengthOfLongestSubstring(s string) int {
	if len(s)==0{
		return 0
	}
	l:=0
	res:=0
	visited:=[256]int{}
	for i:=0;i<len(s);i++{
		if visited[s[i]]!=0{
			if visited[s[i]]>l{
				l=visited[s[i]]
			}
		}
		if res<i-l+1{
			res=i-l+1
		}
		visited[s[i]]=i+1
	}
	return res
}


func main() {
	visited:=make(map[byte]int)
	fmt.Println(visited[3])
	fmt.Println(lengthOfLongestSubstring("dvdf"))
}
