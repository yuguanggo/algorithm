package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	strs:=strings.Split(path,"/")
	stack:=make([]string,0)
	for i:=0;i<len(strs);i++{
		if strs[i]=="."||strs[i]==""{
			continue
		}
		if strs[i]!=".."{
			stack=append(stack,strs[i])
		}else{
			if len(stack)>0{
				stack=stack[0:len(stack)-1]
			}
		}
	}
	return "/"+strings.Join(stack,"/")
}


func main() {

	fmt.Println(simplifyPath("/home/"))
}
