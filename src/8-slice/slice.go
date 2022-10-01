package main

import (
	"fmt"
)

func and(a []int) {
	a[0] = 100
}

func main() {

	// 固定长度赋值
	var myArray [10]int

	// 部分赋值
	myArray2 := [10]int{1, 2, 3, 4}

	for i := 0; i < len(myArray); i++ {
		fmt.Println(myArray[i])
	}

	for index, value := range myArray2 {
		fmt.Println(index, " ", value)
	}

	// 切片，动态数组，
	myArray3 := []int{1, 2, 3, 4}
	for _, value := range myArray3 {
		fmt.Println(value)
	}

	// 传参数的话，会传递地址
	and(myArray3)

	for _, value := range myArray3 {
		fmt.Println(value)
	}

	// 没有开辟内存空间,赋值的话会失败
	var slice1 []int

	// 开辟三个空间
	slice1 = make([]int, 3)
	slice1[0] = 100
	fmt.Println(len(slice1), slice1)

	var slice2 []int = make([]int, 3)
	fmt.Println(slice2)

	slice3 := make([]int, 3)
	fmt.Println(slice3)

	//判断slice是否为空
	if slice1 == nil {
		fmt.Println("null")
	} else {
		fmt.Println("value")
	}

	var numbers = make([]int, 3, 5)
	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 1)
	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 2)
	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	numbers = append(numbers, 3)
	fmt.Printf("%d, %d, %v\n", len(numbers), cap(numbers), numbers)

	var number2 = make([]int, 3)
	number2 = append(number2, 3)
	fmt.Printf("%d, %d, %v\n", len(number2), cap(number2), number2)

	s := []int{1, 2, 3, 4}
	s1 := s[0:2] // 左闭右开
	fmt.Println(s1)

	s1[0] = 100
	fmt.Println(s)
	fmt.Println(s1)

	//copy，底层数组拷贝
	s2 := make([]int, 3)
	copy(s2, s) //把s的值拷贝到s2
	fmt.Println(s2)

}
