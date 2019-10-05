package main

import (
	"fmt"
	"math"
)

func main() {
	var a int = 10
	var b int = 077
	var c int = 0xff
	fmt.Println(a, b, c)  // 10 63 255
	fmt.Printf("%b\n", a) // 1010
	fmt.Printf("%o\n", b) // 77
	fmt.Printf("%x\n", c) // ff

	// 求c内存地址
	fmt.Printf("%p\n", &c) // 十六进制内存地址

	//浮点数常量
	fmt.Println(math.MaxFloat64)
	fmt.Println(math.MaxFloat32)
	fmt.Println(math.Pi)

	// 字符串转义
	fmt.Println("\"c:\\go\"")
	var s1 = "单行文本"
	var s2 = `这是
	多
	行
	文本
	`    // 原样输出
	fmt.Println(s1)
	fmt.Println(s2)
}
