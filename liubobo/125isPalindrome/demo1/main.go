package main

import (
	"strings"
	"fmt"
)

func isPalindrome(s string) bool {
	s=strings.ToLower(s)
	l:=0
	r:=len(s)-1
	for l<r{
		if !(s[l]>='a'&&s[l]<='z'||s[l]>='0'&&s[l]<='9'){
			l++
			continue
		}
		if !(s[r]>='a'&&s[r]<='z'||s[r]>='0'&&s[r]<='9'){
			r--
			continue
		}
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}


func main() {
	//97到122
	s:="2bcdef"
	fmt.Println(s[0]-'0')
}
