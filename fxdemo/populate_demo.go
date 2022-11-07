package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type Tea struct {
	Name string
	Age  int
}

func NewTea() *Tea {
	return &Tea{
		Name: "苍井",
		Age:  18,
	}
}

type TGay struct {
	Tea *Tea
}

func NewTGay(Tea *Tea) *TGay {
	return &TGay{Tea}
}

func RunPopulate() {
	invoke := func(gay *TGay) {

	}
	var gay *TGay //定义一个对象,值为nil
	app := fx.New(
		fx.Provide(NewTea),
		fx.Provide(NewTGay),
		fx.Invoke(invoke),
		fx.Populate(&gay), //调用Populate，这里必须是指针，因为是通过*target 来给元素赋值的
	)
	fmt.Println(gay) //&{0xc00008c680}，将NewGay返回的对象放进var定义的变量里面了
	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
