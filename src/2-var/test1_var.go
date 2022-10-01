package main

import "fmt"

func main() {

	var a int
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	var b int = 100
	fmt.Println(b)
	fmt.Printf("%T\n", b)

	var c = 100
	fmt.Println(c)
	fmt.Printf("%T\n", c)

	// 只能放在函数体内
	d := 100
	fmt.Println(d)
	fmt.Printf("%T\n", d)

	e := "100"
	fmt.Println(e)
	fmt.Printf("%T\n", e)

	var xx, yy int = 100, 200
	fmt.Println(xx, yy)

	var aa, bb = 100, "200"
	fmt.Println(aa, bb)

	var (
		cc int  = 100
		dd bool = true
	)
	fmt.Println(cc, dd)
}
