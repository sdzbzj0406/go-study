package main

import "fmt"

func returnFun() int {
	fmt.Println("return")

	return 0
}

func deferAndReturn() int {
	// defer会在return后面再执行
	defer fmt.Println("defer")
	return returnFun()

}

func main() {
	// 写入defer关键字,final触发机制，最后执行，先进后出的顺序来执行
	defer fmt.Println("main end 1")
	defer fmt.Println("main end 2")

	fmt.Println("hello go 1")
	fmt.Println("hello go 2")
	fmt.Println("hello go 3")

	deferAndReturn()
}
