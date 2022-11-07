package main

import "fmt"

func main() {
	// 匿名函数
	func() {
		sum := 0
		for i := 1; i <= 1e6; i++ {
			sum += i
		}

		fmt.Printf("sum : %d", sum)
	}()
}
