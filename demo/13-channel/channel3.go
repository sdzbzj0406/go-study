package main

import "fmt"

func main3() {
	c := make(chan int)

	go func() {
		for i := 0; i < 5; i++ {
			c <- i
		}

		// 若关闭，则会造成死锁
		close(c)
	}()

	for {
		if data, ok := <-c; ok {
			fmt.Println(data)
		} else {
			fmt.Println("close")
			break
		}
	}

	fmt.Println("main end")

}
