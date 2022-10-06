package main

import "fmt"

// 万能类型interface
func myFunc(arg interface{}) {

	fmt.Println("myFunc")
	fmt.Println(arg)

	// 如何区分类型,根据interface 断言判断类型
	value, ok := arg.(string)
	if !ok {
		fmt.Println("not string")
	} else {
		fmt.Println("is string, value : ", value)
	}

}

type Book2 struct {
	auth string
}

func main() {
	b := Book2{
		auth: "aaa",
	}

	myFunc(b)
	myFunc(100)
	myFunc("112")

}
