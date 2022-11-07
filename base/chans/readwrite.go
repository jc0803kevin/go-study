package main

import (
	"fmt"
)

func sendData(ch chan string) {
	ch <- "Washington"
	ch <- "Tripoli"
	ch <- "London"
	ch <- "Beijing"
	ch <- "Tokyo"
}

func getData(ch chan string) {
	var input string
	// time.Sleep(2e9)
	for {
		input = <-ch
		fmt.Printf("%s \n", input)
	}
}

func Run() {

	ch := make(chan string)

	go sendData(ch)
	go getData(ch)

	// 如果注释掉，由于主线程已经关闭，协程还没有执行完
	//time.Sleep(time.Second * 10)

}

func put(ch chan int) {
	for i := 0; ; i++ {
		ch <- i
	}
}

func Pump() {
	ch := make(chan int)
	go put(ch)

	// 一个协程在无限循环中给通道发送整数数据。不过因为没有接收者，只输出了一个数字 0。
	//fmt.Println(<- ch)

	// 每接收一次都是最新的值
	//0
	//1
	//2
	//3
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)
	//fmt.Println(<- ch)
}
