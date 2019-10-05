package main

import "fmt"

func foo()(string, int) {
	return "alex", 9000
}

func main() {
	// var name string
	// var age string
	// // 声明变量必须要使用
	// fmt.Println(name)
	// fmt.Println(age)

	// var (
	// 	a string // ""
	// 	b int    // 0
	// 	c bool   // false
	// 	d string // ""
	// )
	// a = "沙河"
	// b = 100
	// c = true
	// d = "100"
	// fmt.Println(a, b, c, d)
	// // 赋值声明变量并初始化
	// var x string = "laonanhai"
	// fmt.Println(x)
	// fmt.Printf("%s heiheihei %d", x, b)
	// // 类型推导（编译器根据变量的初始值的类型，指定给变量）
	// var y = 200
	// var z = true
	// fmt.Println(y)
	// fmt.Println(z)

	// // 简短变量声明（只能在函数内部使用）
	// nazha := "heiheihei"
	// fmt.Println(nazha)

	// // 调用foo函数
	// // _ （匿名变量）用于接受不需要的变量值
	// aa, _ := foo()
	// fmt.Println(aa)

	//不能重复声明同名变量, 匿名变量可以重复使用
	var name = "alex"
	// var name = "nazha"
	fmt.Println(name)


}
