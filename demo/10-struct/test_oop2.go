package main

import "fmt"

type Human struct {
	name string
	sex  string
}

func (h *Human) Eat() {
	fmt.Println("eat")
}

func (h *Human) Walk() {
	fmt.Println("walk")
}

type SuperMan struct {
	Human // 代表继承
	level int
}

// Eat 重写
func (s *SuperMan) Eat() {
	fmt.Println("superman eat")
}

func (s *SuperMan) fly() {
	fmt.Println("superman fly")
}

func (s *SuperMan) show() {
	fmt.Println(s.name, " ", s.sex, " ", s.level)
}

func main4() {

	h := Human{
		name: "zhang3",
		sex:  "1",
	}
	h.Eat()
	h.Walk()

	s := SuperMan{h, 1}
	s.show()

	s2 := SuperMan{Human{"li4", "111"}, 1}
	s2.show()

	var s3 SuperMan
	s3.name = "222"
	s3.sex = "333"
	s3.level = 2

	s3.show()
	s3.show()

}
