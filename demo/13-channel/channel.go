package main

import "fmt"

func main1() {

	c := make(chan int)

	go func() {
		defer fmt.Println("go defer")
		fmt.Println("go runing")

		c <- 666 // 666写入channel
		fmt.Println("go end")

	}()

	//num := <-c // 读取channel的值
	//fmt.Println(num)

	fmt.Println("main end")

}
