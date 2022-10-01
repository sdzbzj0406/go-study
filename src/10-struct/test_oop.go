package main

import "fmt"

// 类字母大写，其他包也可以访问
type Hero struct {

	// 属性大写，其他包可以访问
	Name  string
	Ad    int
	Level int
}

//func (h Hero) getName() string {
//	fmt.Println(h.Name)
//	return h.Name
//}
//
//func (h Hero) setName(newName string) {
//	// 调用该对象的拷贝，不会修改原始值
//	h.Name = newName
//}
//
//func show(hero Hero) {
//	fmt.Println(hero)
//}

func (h *Hero) getName() string {
	fmt.Println(h.Name)
	return h.Name
}

func (h *Hero) setName(newName string) {
	// 调用该对象的拷贝，不会修改原始值
	h.Name = newName
}

func show(hero *Hero) {
	fmt.Println(hero.getName())
}
func main3() {

	hero := Hero{
		Name:  "zhang3",
		Ad:    100,
		Level: 1,
	}

	fmt.Println(hero.getName())
	show(&hero)

	hero.setName("li4")
	show(&hero)

}
