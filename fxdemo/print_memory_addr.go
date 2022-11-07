package main

import (
	"fmt"
	"go.uber.org/fx"
)

type Girl struct {
	Name string
}

func NewGirl() *Girl {
	return &Girl{Name: "蓝燕"}
}
func NewGirl1() Girl {

	return Girl{Name: "蓝燕1"}
}

func Run20221031() {

	invoke1 := func(g *Girl) {
		fmt.Printf("invoke1:%p Name: %s \n", g, g.Name)
	}
	invoke2 := func(g *Girl) {
		fmt.Printf("invoke2:%p Name: %s \n", g, g.Name)
	}

	invoke3 := func(g Girl) {
		fmt.Printf("invoke3:%p Name: %s \n", &g, g.Name)
	}
	invoke4 := func(g Girl) {
		fmt.Printf("invoke4:%p Name: %s \n", &g, g.Name)
	}
	invoke5 := func(g Girl) {
		fmt.Printf("invoke5:%p Name: %s \n", &g, g.Name)
	}
	fx.New(fx.Provide(NewGirl, NewGirl1),
		fx.Invoke(invoke2, invoke1, invoke3, invoke4, invoke5))

	//fx 的原理是调用构造函数，把值存起来，下次接着用这个值，只会调用一次构造函数。

	//如果构造函数返回是指针的话，能保证每次是同一个对象
	//invoke2:0xc00008a5d0 Name: 蓝燕
	//invoke1:0xc00008a5d0 Name: 蓝燕

	//如果构造函数返回是对象的话，保证不了，因为对象在多次传递中会发生拷贝，对象已经不是原来的对象。
	//invoke3:0xc00008a8c0 Name: 蓝燕1
	//invoke4:0xc00008a9c0 Name: 蓝燕1
	//invoke5:0xc00008aac0 Name: 蓝燕1

}
