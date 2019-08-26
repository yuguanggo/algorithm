package main

import "strings"

func longestCommonPrefix(strs []string) string {
	n:=len(strs)
	prefix:=strs[0]
	for i:=1;i<n;i++{
		for strings.Index(strs[i],prefix)==-1{
			prefix=prefix[0:len(prefix)-1]
			if len(prefix)==0{
				return ""
			}
		}
	}
	return prefix
}

func main() {
	
}
