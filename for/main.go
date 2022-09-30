package main

import (
	"fmt"
	"strconv"
)

func main() {

	readData()

}

func for1() {

	for i := 0; i < 5; i++ {

		fmt.Printf("index ： %d  \n", i)

	}

}

func readData() {

	var number string
	//控制台提示语句
	fmt.Print("请输入一个整数：")
	//控制台的输出
	fmt.Scan(&number)
	fmt.Println("数值是：", number)
	fmt.Printf("数据类型是：%T\n", number)
	//数据类型转换string---> int
	value, _ := strconv.Atoi(number)
	fmt.Printf("转换后的数据类型是：%T  值：%d \n", value, value)
	//数值判断
	if value > 100 {
		fmt.Println("数值 > 100")
	} else {
		fmt.Println("数值 <= 100")
	}

}
