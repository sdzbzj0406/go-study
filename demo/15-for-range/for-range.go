package main

// 1.遍历数组，切片，string，map时，忽略第二个值，通过下标，可以提升性能
// 2.range作用于channel时，除非channel close，不然会一直阻塞
// 3.range在遍历时，修改的新数据会遍历不到，因为遍历的次数在开始遍历的时候已经确定了
// 4.string在range时，下标不一定是连续的，因为默认按照utf8编码，中文占用的字节数是不确定的；返回的第二个值为rune类型的单个字节值

import (
	"fmt"
	"time"
)

func main() {
	//RangeArray()
	//RangeSlice()
	//RangeString()
	//RangeMap()
	RangeChannel()
}

func RangeArray() {
	a := [3]int{1, 2, 3}

	for i, v := range a {
		fmt.Println("index:", i, " value: ", v)
	}

	fmt.Println("---------------")

	for i := range a {
		fmt.Println("index:", i, " value: ", a[i])
	}

}

func RangeSlice() {
	a := []int{1, 2, 3}
	for i := range a {
		fmt.Println("index:", i, " value: ", a[i])
	}
}

func RangeString() {
	s := "中国"
	for i, v := range s {
		fmt.Println("index:", i, " value:", v)
		fmt.Printf("index:%d, value:%c\n", i, v)
	}
}

func RangeMap() {

	var m = map[string]string{"1": "a", "2": "b"}
	for k, v := range m {
		fmt.Println("key: ", k, " value: ", v)
	}

	// 只有一个值接收的话，则表示会省略第二个值
	for k := range m {
		fmt.Println("key:", k)
	}

}

func RangeChannel() {

	c := make(chan string, 2)

	c <- "a"
	c <- "b"

	time.AfterFunc(time.Microsecond, func() {
		close(c)
	})

	for e := range c {
		fmt.Println(e)
	}

}
