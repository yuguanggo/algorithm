package main

import "fmt"

func FibnacciClosure() func() (ret int) {
	a, b := 0, 1
	return func() (ret int) {
		ret = b
		a, b = b, a+b
		return
	}
}

func FibnacciChan(n int) chan int {
	ret:=make(chan int)
	go func() {
		a,b:=0,1
		for i:=0;i<n;i++{
			ret<-b
			a,b=b,a+b
		}
		close(ret)
	}()
	return ret
}

func main() {
	fib:=FibnacciClosure()
	for i:=0;i<10 ;i++  {
		fmt.Println(fib())
	}

	for i:=range FibnacciChan(20){
		fmt.Println(i)
	}
}
