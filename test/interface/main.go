package main

import "fmt"

type MyInterface interface {
	Print()
	Wte()
}

func TestFunc(x MyInterface)  {

}

type MyStruct struct {

}

func (me MyStruct) Print()  {

}
func (me MyStruct) Wte()  {

}


func main() {
	var me MyStruct
	TestFunc(me)
	var m chan int
	m = make(chan int)
	fmt.Println(m)
}
