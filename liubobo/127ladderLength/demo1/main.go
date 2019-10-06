package main

import (
	"fmt"
	"strconv"
)

func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap:=make(map[string][]string)
	n:=len(beginWord)
	//初始化map
	for i:=0;i<len(wordList);i++{
		for j:=0;j<n;j++{
			bytes:=[]byte(wordList[i])
			bytes[j]='*'
			if _,ok:=wordMap[string(bytes)];!ok{
				wordMap[string(bytes)]=[]string{wordList[i]}
			}else{
				wordMap[string(bytes)]=append(wordMap[string(bytes)],wordList[i])
			}
		}
	}
	fmt.Println(wordMap)
	q:=make([][2]string,0)
	q=append(q,[2]string{beginWord,"1"})
	visited:=make(map[string]bool)
	visited[beginWord]=true
	for len(q)>0{
		front:=q[0]
		q=q[1:]
		w:=front[0]
		step,_:=strconv.Atoi(front[1])
		fmt.Println(step)
		for i:=0;i<n;i++{
			bytes:=[]byte(w)
			bytes[i]='*'
			if _,ok:=wordMap[string(bytes)];ok{
				for _,v:=range wordMap[string(bytes)]{
					if v==endWord{
						return step+1
					}
					if !visited[v]{
						q=append(q,[2]string{v,strconv.Itoa(step+1)})
						visited[v]=true
					}
				}
			}
		}
	}
	return 0
}


func main() {
	ladderLength("hit","cog",[]string{"hot","dot","dog","lot","log","cog"})
}
