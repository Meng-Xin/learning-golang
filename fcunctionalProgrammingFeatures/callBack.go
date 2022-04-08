package main

import "fmt"

/*
	函数编程特性：回调函数
		特征：作为回调函数，其主要特征就是 一个方法作为形参的方式进行传入。
		功能：功能解耦，伪泛型？ 还有就是 同步执行或异步执行，因为传入的是方法，可以在不同的时机使用。
*/

type BackFunction func(a, b int) int

func Add(a, b int) int {
	return a + b
}
func Sub(a, b int) int {
	return a - b
}

func Mul(a, b int) int {
	return a * b
}

func Accomplish(a, b int, call BackFunction) int {
	return call(a, b)
}

func main() {
	// 使用回调函数
	fmt.Println(Accomplish(1, 2, Add))
}
