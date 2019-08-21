package main

import (
	"fmt"
	"unicode/utf8"
)

type T struct {
	V int
}

func test(t *T)  {
	t.V=2
}

func testSilce(arr []int)  {
	arr[0]=1
	fmt.Printf("%p ",&arr)
}

func main() {
	s:="yes我爱中国"
	fmt.Println(len(s))
	fmt.Println(utf8.RuneCountInString(s))
	for k,i:=range []byte(s){
		fmt.Printf("(%d %X)",k,i)
	}
	fmt.Println()
	for k,i:=range s{
		fmt.Printf("(%d %X)",k,i)
	}
	bytes:=[]byte(s)
	for len(bytes)>0{
		ch,size := utf8.DecodeRune(bytes)
		bytes=bytes[size:]
		fmt.Printf("%c %d",ch,size)
	}
	for k,i:= range []rune(s){
		fmt.Printf("%d,%c",k,i)
	}
}
