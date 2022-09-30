package main

import (
	"fmt"
	"go-study/types"
)

func main() {

	var i1 = 5
	fmt.Printf("An integer: %d, it's location in memory: %p\n", i1, &i1)

	// An integer: 5, it's location in memory: 0xc0000a6058
	// 打印出内存地址

	stu := types.Student{Name: "kevin", Age: 18}
	my := MyStudent{desc: "kkkkk"}
	my.stu = &stu

	fmt.Print(my.ToString())

}

// 扩展一个结构体
type MyStudent struct {
	stu *types.Student

	desc string
}

func (my *MyStudent) ToString() string {

	return fmt.Sprintf("name:%s  Age:%d  desc:%s", my.stu.Name, my.stu.Age, my.desc)
}
