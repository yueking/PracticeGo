package main

import "fmt"

func main() {
	fmt.Println("...starting...")
	test()
	fmt.Println("...end...")
}
func test() {
	// defer + recover 捕获处理异常
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("错误已经捕获err:", err)
		}
	}()
	n1 := 10
	n2 := 0
	result := n1 / n2
	fmt.Println(result)
}
