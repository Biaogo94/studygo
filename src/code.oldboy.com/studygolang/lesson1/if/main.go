package main

import "fmt"

func main() {
	age := 19
	if age > 18 {
		fmt.Println("澳门首家线上赌场开业了")
	} else if age < 18 {
		fmt.Println("Warning..")
	} else {
		fmt.Println("成年了")
	}

	if age2 := 28; age2 > 18 {    // age2 无法外面打印，作用域在if语句块内
		fmt.Println("成年了")
	}
}
