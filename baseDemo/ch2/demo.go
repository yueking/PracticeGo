package main

import (
	"errors"
	"fmt"
)

func main() {
	fmt.Println("...starting...")
	err := test()
	if err != nil {
		fmt.Println("自定义错误:", err)
		panic(err)
	}
	fmt.Println("...end...")
}
func test() error {
	n1 := 10
	n2 := 0
	if n2 != 0 {
		result := n1 / n2
		fmt.Println(result)
	} else {
		return errors.New("除数不能为0")
	}
	return nil
}
