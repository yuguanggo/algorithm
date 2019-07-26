package main

import (
	"math"
	"fmt"
	"strconv"
)

/**
126. 单词接龙 II
给定两个单词（beginWord 和 endWord）和一个字典 wordList，找出所有从 beginWord 到 endWord 的最短转换序列。转换需遵循如下规则：

每次转换只能改变一个字母。
转换过程中的中间单词必须是字典中的单词。
说明:

如果不存在这样的转换序列，返回一个空列表。
所有单词具有相同的长度。
所有单词只由小写字母组成。
字典中不存在重复的单词。
你可以假设 beginWord 和 endWord 是非空的，且二者不相同。
示例 1:

输入:
beginWord = "hit",
endWord = "cog",
wordList = ["hot","dot","dog","lot","log","cog"]

输出:
[
  ["hit","hot","dot","dog","cog"],
  ["hit","hot","lot","log","cog"]
]
示例 2:

输入:
beginWord = "hit"
endWord = "cog"
wordList = ["hot","dot","dog","lot","log"]

输出: []

解释: endWord "cog" 不在字典中，所以不存在符合要求的转换序列。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/word-ladder-ii
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
 */
func findLadders(beginWord string, endWord string, wordList []string) [][]string {
	wordMap:=make(map[string][]string,0)
	n:=len(beginWord)
	for _,v:=range wordList{
		for i:=0;i<n;i++{
			bytes:=[]byte(v)
			bytes[i]='*'
			if _,ok:=wordMap[string(bytes)];!ok{
				wordMap[string(bytes)]=[]string{v}
			}else{
				wordMap[string(bytes)]= append(wordMap[string(bytes)],v)
			}
		}
	}
	res:=make([][]string,0)
	queue:=make([][]string,0)
	queue=append(queue,[]string{beginWord})
	level:=1
	minLevel:=math.MaxInt32
	for len(queue)>0{
		s:=queue[0]
		word:=s[len(s)-1]
		queue=queue[1:]
		visited:=make(map[string]bool)
		for _,sv:=range s{
			visited[sv]=true
		}
		level=len(s)
		if level>minLevel{
			break
		}
		for i:=0;i<n;i++{
			bytes:=[]byte(word)
			bytes[i]='*'
			if _,ok:=wordMap[string(bytes)];ok{
				for _,v:=range wordMap[string(bytes)]{
					next:=append(s,v)
					if v==endWord{
						res=append(res,next)
						minLevel=level
					}
					if !visited[v]{
						queue=append(queue,next)
					}
				}

			}
		}
	}
	return res
}
func main() {
	fmt.Println(findLadders("red","tax",[]string{"ted","tex","red","tax","tad","den","rex","pee"}))
	s:=[]string{"a","b"}
	res:=make([][]string,0)
	for i:=0;i<20;i++{
		next:=append(s,strconv.Itoa(i))
		res=append(res,next)
	}
	fmt.Println(res)
}
