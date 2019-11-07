package main

import (
	"bufio"
	"fmt"
	"os"
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
	scanner:=bufio.NewScanner(os.Stdin)
	scanner.Scan()
		fmt.Println(scanner.Text())
	fmt.Println(22)
}
