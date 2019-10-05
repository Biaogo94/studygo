package main

import "fmt"


func switchDemo1() {
	finger := 3
	switch finger {
	case 1:
		fmt.Println("大拇指")
	case 2:
		fmt.Println("食指")
	case 3:
		fmt.Println("中指")
	case 4:
		fmt.Println("无名指")
	case 5:
		fmt.Println("小拇指")
	default:
		fmt.Println("无效的输入！")
	}
}

func testSwitch3() {
	switch n := 7; n {    // 变量声明在switch内部
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8:
		fmt.Println("偶数")
	default:
		fmt.Println(n)
	}
}

func switchDemo4() {
	age := 30
	switch {      // 不需要判断
	case age < 25:
		fmt.Println("好好学习吧")
	case age > 25 && age < 35:
		fmt.Println("好好工作吧")
	case age > 60:
		fmt.Println("好好享受吧")
	default:
		fmt.Println("活着真好")
	}
}

func main() {
	age := 18
	for age > 0 {   // 相当于别的语言的while循环
		fmt.Println(age)
		age = age - 1
	}

	// switch 循环  if 判断的变种，适用于很多if判断
	switchDemo1()
	switchDemo4()
	testSwitch3()
}


