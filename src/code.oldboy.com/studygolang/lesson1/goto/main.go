package main

import "fmt"

// goto
func main() {
	for i := 0; i < 5; i++ {
		for j := 0; j < 3; j++ {
			if i == 2 && j == 2 {
				break   // 跳出for循环
				// continue   //继续下一次循环
			}
			fmt.Printf("%d--%d\n", i, j)
		}
	}
	fmt.Println("两曾for 循环结束")
}
