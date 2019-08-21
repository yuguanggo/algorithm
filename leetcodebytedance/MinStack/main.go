package main

import "fmt"

type MinStack struct {
	stack []int
	mins []int
}


/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{}
}


func (this *MinStack) Push(x int)  {
	this.stack = append(this.stack,x)
	if len(this.mins)==0||x<=this.GetMin(){
		this.mins = append(this.mins,x)
	}
}


func (this *MinStack) Pop()  {
	v:=this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	if v==this.GetMin(){
		this.mins = this.mins[:len(this.mins)-1]
	}
}


func (this *MinStack) Top() int {
	return this.stack[len(this.stack)-1]
}


func (this *MinStack) GetMin() int {
	return this.mins[len(this.mins)-1]
}

func main() {
	var minstack = Constructor()
	minstack.Push(-2)
	minstack.Push(0)
	minstack.Push(-3)
	fmt.Println(minstack.GetMin())
	minstack.Pop()
	fmt.Println(minstack.Top())
	fmt.Println(minstack.GetMin())
}
