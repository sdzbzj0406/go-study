package main

import "fmt"

func swap(a int, b int) {

	var temp = a
	a = b
	b = temp
}

func swap2(pa *int, pb *int) {
	var temp = *pa
	*pa = *pb
	*pb = temp
}

func main() {

	var a = 10
	var b = 20

	fmt.Println("a=", a, "b=", b)

	swap2(&a, &b)

	var p *int
	p = &a

	var pp **int
	pp = &p

	fmt.Println(&p)
	fmt.Println(pp)

}
