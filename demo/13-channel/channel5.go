package main

import "fmt"

func f(c, quit chan int) {
	x, y := 1, 1

	for {
		select {
		case c <- x:
			x = y
			y = x + 1
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {

	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 6; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()

	f(c, quit)

}
