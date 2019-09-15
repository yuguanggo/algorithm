package main

import "strings"

func wordPattern(pattern string, str string) bool {
	words:=strings.Split(str," ")
	m:=len(pattern)
	n:=len(words)
	if m!=n{
		return false
	}
	record:=make(map[byte]string)
	visited:=make(map[string]bool)
	for i:=0;i<m;i++{
		if _,ok:=record[pattern[i]];ok{
			if record[pattern[i]]!=words[i]{
				return false
			}
		}else{
			record[pattern[i]]=words[i]
			if _,ok:=visited[words[i]];ok{
				return false
			}
			visited[words[i]]=true
		}
	}
	return true
}

func main() {
	
}
