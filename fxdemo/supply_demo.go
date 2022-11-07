package main

type Stu struct {
	Name string
	Age  int
}

func NewStu() *Stu {
	return &Stu{
		Name: "苍井",
		Age:  18,
	}
}

type Gay struct {
	Stu *Stu
}

func NewGay(Stu *Stu) *Gay {
	return &Gay{Stu}
}

//func NewGay (say  SayInterface)*Gay  {//此处能够正常获取到
//	return &Gay{}
//}
//
//type SayInterface interface {
//	SayHello()
//}
//
//func (g* Girl)SayHello()  {
//	fmt.Println("girl sayhello")
//}

func RunSupply() {
	//invoke:= func(gay* Gay) {
	//	fmt.Println(gay.Stu)
	//}
	//
	//stu:=NewStu() //直接提供对象, 然后通过Supply 传入该对象
	//app:=fx.New(
	//	fx.Provide(NewGay),
	//	fx.Supply(stu),
	//	fx.Invoke(invoke),
	//)
	//err:=app.Start(context.Background())
	//if err!=nil{
	//	panic(err)
	//}

	// 不同使用接口类型，
	//invoke:= func(gay *Gay) {
	//	fmt.Println(gay)
	//}
	//app:=fx.New(
	//	fx.Provide(NewGirl),
	//	fx.Provide(NewGay),
	//	fx.Invoke(invoke),
	//)
	//err:=app.Start(context.Background())
	//if err!=nil{
	//	panic(err)
	//}
}
