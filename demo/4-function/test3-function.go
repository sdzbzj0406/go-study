package main

import "fmt"

func fool(a string, b int) int {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	c := 100
	return c
}

// 返回多个返回值，匿名结果
func foo2(a string, b int) (int, int) {
	fmt.Println("a=", a)
	fmt.Println("b=", b)

	return 666, 777
}

// 返回多个返回值，有名称
func foo3(a string, b int) (r1 int, r2 int) {

	// r1,r2初始化为0，作用域为局部变量foo3

	r1 = 1000
	r2 = 2000

	return
}

func foo4(a string, b int) (r1, r2 int) {

	r1 = 3000
	r2 = 4000

	return
}

func main() {

	c := fool("abc", 100)

	fmt.Println(c)

	ret1, ret2 := foo2("abc", 200)
	fmt.Println(ret1)
	fmt.Println(ret2)

	ret3, ret4 := foo3("abc", 200)
	fmt.Println(ret3)
	fmt.Println(ret4)

	ret5, ret6 := foo4("abc", 200)
	fmt.Println(ret5)
	fmt.Println(ret6)

}
