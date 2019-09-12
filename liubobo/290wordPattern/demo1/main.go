package main

import "strings"

func wordPattern(pattern string, str string) bool {
	words:=strings.Split(str," ")
	m:=len(pattern)
	n:=len(words)
	if m!=n{
		return false
	}
	pi:=0
	si:=0
	record:=make(map[byte]string)
	for pi<m&&si<n{
		if _,ok:=record[pattern[pi]];!ok{

		}
	}
}

func main() {
	
}
