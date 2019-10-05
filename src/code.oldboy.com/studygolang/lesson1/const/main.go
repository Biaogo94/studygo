package main

// import "fmt"

// const pi = 3.14

// 批量声明常量
// const (
// 	a = 100
// 	b = 1000
// 	c       // 不赋值等于上面的值
// 	d
// )

// iota  每新增一行常量变量声明，会使iota增加1
// const (
// 	aa = iota  //0
// 	bb			//1
// 	cc			//2
// 	dd			//3

// )

// const (
// 	n1= iota
// 	n2
// 	_
// 	n4
// )

// const (
// 	n1 = iota
// 	n2 = 100
// 	n3 = iota
// 	n4
// )
// const n5 = iota


// const (
// 	_ = iota //0
// 	KB = 1 << (10 * iota)   // 1<<10  ==> 2的10次方 = 1024
// 	MB = 1 << (10 * iota)   // 1<<20 =1024*1024  
// 	GB = 1 << (10 * iota)   // 1<<30 
// 	TB = 1 << (10 * iota)   // 1<<40
// 	PB = 1 << (10 * iota)   // 1<<50
	
// )

// const (
// 	a, b = iota + 1, iota + 2 // iota=0 a=1 b=2
// 	c, d                      // iota=1 c=2 d=3
// 	e, f                      // iota=2 e=3 f=4
// )

// 常量
func main() {
	// pi = 3.1415  //常量不允许修改
	// fmt.Println(pi)

	// fmt.Println(a, b,c,d)  //100 1000 1000 1000

	// fmt.Println(aa, bb, cc, dd)  //0 1 2 3

	// fmt.Println(n1,n2,n4)   //0 1 3

	// fmt.Println(n1, n2, n3, n4, n5)  //0 100 2 3 0

}

// iota
// 0. const 声明如果不写，默认就和上一行一样
// 1. 遇到const初始化为零
// 2. const每新增一行变量声明，iota递增1