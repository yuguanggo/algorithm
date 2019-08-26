package main

import "strings"

func simplifyPath(path string) string {
	s:=strings.Split(path,"/")
	res:=make([]string,0)
	for i:=0;i<len(s);i++{
		if s[i]==""||s[i]=="."{
			continue
		}else if s[i]==".."{
			if len(res)>0{
				res=res[0:len(res)-1]
			}

		}else{
			res=append(res,s[i])
		}
	}
	return "/"+strings.Join(res,"/")
}

func main() {
	
}
