package main

import "fmt"

type Incr struct {
	num int
}

func (i *Incr) Increase(n int)  {
	i.num=+n
}

func (i *Incr) Decrease(n int)  {
	i.num=-n
}

func (i *Incr) String()  {
	fmt.Println("woshi:%v",i.num)
}

func main() {
	var i Incr

	i.Increase(10)
	fmt.Println(i)
	i.Decrease(5)
	fmt.Println(i)
}
