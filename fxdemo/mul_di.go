package main

//https://zhuanlan.zhihu.com/p/418299054

import (
	"fmt"
	"go.uber.org/fx"
)

type A struct {
	B *B
}

func NewA(b *B) *A {
	return &A{B: b}
}

type B struct {
	C *C
}

func NewB(c *C) *B {
	return &B{c}
}

type C struct {
}

func NewC() *C {
	return &C{}
}

func PrintA(a *A) {
	fmt.Println(*a)
}

func RunMulDi() {

	//我们需要一个a
	//b:=NewB(NewC())
	//a:=NewA(b)
	//_=a
	//PrintA(a)

	fx.New(
		// 提供依赖 构造函数
		//将被依赖的对象的构造函数传进去，传进去的函数必须是个待返回值的函数指针
		fx.Provide(NewB),
		fx.Provide(NewA),
		fx.Provide(NewC),

		// 执行
		fx.Invoke(PrintA),
	).Run()
}
