package main

import (
	"flag"
	"fmt"
)

func main() {

	name := flag.String("name", "", "指定名称")
	age := flag.Int("age", 0, "指定名称")

	// 用于真正解析命令参数，参数绑定都必须在此之前。
	flag.Parse()

	fmt.Printf("I'am %s , I'am %d years old.", *name, *age)

}
