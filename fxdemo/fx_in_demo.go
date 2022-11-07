package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

// 对于结构体Gay1， 没有定义NewGay1 的构造函数
// 让相同的对象按照tag能够赋值到一个结构体上面，结构体必须内嵌http://fx.in
type Gay1 struct {

	// 如果注释掉fx.In， 就会出现 main.Gay1 (did you mean to Provide it?)
	// 该对象没有进行构造

	fx.In
	Girl1 *Girl   `name:"波多1"`
	Girl2 *Girl   `name:"海翼"`
	Girls []*Girl `group:"actor"`
}

func RunFxIn() {
	invoke := func(gay Gay1) {
		fmt.Println(gay.Girl1.Name)                    //波多
		fmt.Println(gay.Girl2.Name)                    //海翼
		fmt.Println(len(gay.Girls), gay.Girls[0].Name) //1 杏梨
	}

	// Group和Name 是不能同时存在的

	app := fx.New(
		fx.Invoke(invoke),
		fx.Provide(

			fx.Annotated{
				Target: func() *Girl { return &Girl{Name: "波多_value"} },
				Name:   "波多1",
			},
			fx.Annotated{
				Target: func() *Girl { return &Girl{Name: "海翼_value"} },
				Name:   "海翼",
			},
			fx.Annotated{
				Target: func() *Girl { return &Girl{Name: "杏梨_value"} },
				Group:  "actor",
			},
		),
	)

	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
