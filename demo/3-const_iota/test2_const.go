package main

import "fmt"

const (
	// BEIJING 可以在const() 添加关键字iota， iota每行累加1，默认第一行是0
	// iota只能出现在const进行累加效果
	BEIJING = 10 * iota
	SHANGHAI
	SHENZHEN
)

func main() {

	// 常量，只读属性
	const length int = 10

	fmt.Println(length)

	// length = 100   常量不允许修改

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("SHANGHAI = ", SHANGHAI)
	fmt.Println("SHENZHEN =", SHENZHEN)

}
