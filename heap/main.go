package main

import (
	"fmt"
	"sync"
	"time"
)

type job struct {
	name int
}

func (j job) do()  {
	time.Sleep(time.Second)
	fmt.Println(j.name)
}

func doJobs(jobs []job,p int)  {
	ch:=make(chan job,p)
	go func() {
		for i:=0;i<len(jobs);i++{
			ch<-jobs[i]
		}
		close(ch)
	}()

	var wg sync.WaitGroup

	for i:=0;i<p;i++{
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			for{
				job,ok:=<-ch
				if !ok{
					break
				}
				fmt.Printf("当前现线程:%d\n",n)
				job.do()
			}
		}(i)
	}
	wg.Wait()

}

func main() {
	jobs:=make([]job,0)
	for i:=0;i<20;i++{
		job:=job{name:i}
		jobs=append(jobs,job)
	}
	fmt.Println(jobs)
	doJobs(jobs,3)

	//token:=make(chan int,3)
	//go func() {
	//	fmt.Println(1)
	//
	//	time.Sleep(time.Second)
	//	fmt.Println(2)
	//	for i:=1;i<=5;i++{
	//		token<-i
	//	}
	//	close(token)
	//}()
	//fmt.Println("----")
	//for {
	//	x,ok:=<-token
	//	fmt.Println("for")
	//	fmt.Println(x)
	//	if !ok{
	//		break
	//	}
	//}
}
