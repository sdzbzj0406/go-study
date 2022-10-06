package main

import (
	_ "5-init/lib1" // _表示匿名别名，会调用init方法
	. "5-init/lib2"
)

func main() {

	//lib1.Test()
	//my.Test()
	Test() // 用.导入的话，会作为当前包中的方法使用
}
