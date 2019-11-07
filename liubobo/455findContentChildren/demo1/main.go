package main

import "sort"

func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)
	gi:=0
	si:=0
	res:=0
	for gi<len(g)&&si<len(s){
		if s[si]>=g[gi]{
			gi++
			si++
			res++
		}else{
			si++
		}
	}
	return res
}


func main() {
	
}
