package main

import (
	"context"
	"fmt"
	"go.uber.org/fx"
)

type GayA struct {
	fx.Out
	Girl1 *Girl `name:"波多"`
}
type GayB struct {
	fx.Out
	Girl1 *Girl `name:"仓井"`
}
type Man struct {
	fx.In
	Girl1 *Girl `name:"波多"`
	Girl2 *Girl `name:"仓井"`
}

func NewGayA() GayA {
	return GayA{
		Girl1: &Girl{Name: "波多"},
	}
}
func NewGayB() GayB {
	return GayB{
		Girl1: &Girl{Name: "仓井"},
	}
}

func RunFxOut() {
	invoke := func(man Man) {
		fmt.Println(man.Girl1.Name) //波多
		fmt.Println(man.Girl2.Name) //仓井
	}

	app := fx.New(
		fx.Invoke(invoke),
		fx.Provide(
			NewGayA, NewGayB,
		))

	err := app.Start(context.Background())
	if err != nil {
		panic(err)
	}
}
