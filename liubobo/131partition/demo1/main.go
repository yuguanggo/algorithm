package main

import "fmt"

func partition(s string) [][]string {
	res:=make([][]string,0)
	helper(s,0,[]string{},&res)
	return res
}

func helper(s string,pos int,path []string,res *[][]string)  {
	if pos==len(s){
		p:=make([]string,len(path))
		copy(p,path)
		*res=append(*res,p)
		return
	}
	for i:=1;i<=len(s);i++{
		if pos+i>len(s){
			break
		}
		sub:=s[pos:pos+i]
		if isHw(sub){

			path=append(path,sub)
			helper(s,pos+i,path,res)
			path=path[:len(path)-1]
		}
	}
}

func isHw(s string)bool  {
	l:=0
	r:=len(s)-1
	for l<r{
		if s[l]!=s[r]{
			return false
		}
		l++
		r--
	}
	return true
}


func main() {
	fmt.Println(partition("a"))
}
