package main

import (
	"fmt"
)

// AnimalIF 本质是指针
type AnimalIF interface {
	Sleep()
	GetColor() string
	GetType() string
}

type Cat struct {
	color string
}

func (c *Cat) Sleep() {
	fmt.Println("cat is sleep")
}

func (c *Cat) GetColor() string {
	return c.color
}

func (c *Cat) GetType() string {
	return "this is a cat"
}

type Dog struct {
	color string
}

func (d *Dog) Sleep() {
	fmt.Println("dog is sleep")
}

func (d *Dog) GetColor() string {
	return d.color
}

func (d *Dog) GetType() string {
	return "this is a doy"
}

func showAnimal(animalIF AnimalIF) {
	animalIF.Sleep() // 多态,传入接口
	fmt.Println(animalIF.GetType())
	fmt.Println(animalIF.GetColor())
}

func main5() {

	c := Cat{
		color: "blue",
	}

	fmt.Println(c)

	var a AnimalIF
	a = &Cat{"white"}
	a.Sleep() // 调用cat

	a = &Dog{"yellow"}
	a.Sleep() // 调用dog

	cat := Cat{"aaa"}
	showAnimal(&cat)

}
