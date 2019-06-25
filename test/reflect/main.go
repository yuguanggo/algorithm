package main

import (
	"reflect"
	"fmt"
)

type Person struct {
	name string
	age int
}

func main() {
	kelu:=Person{name:"kelu",age:11}
	t:=reflect.TypeOf(kelu)
	fmt.Println(t.Kind())
	n:=t.NumField()
	for i:=0;i<n;i++{
		fmt.Println(t.Field(i).Name)
		fmt.Println(t.Field(i).Type)
	}
}
