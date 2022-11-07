package types

import "fmt"

//func Add(a int, b int) (result int){
//
//	result := a + b
//	return  result
//}

func Show() {
	fmt.Println("welcome to go world.")
}

func MultiPly3Nums(a int, b int, c int) int {
	return a * b * c
}

func MultiPly3Nums2(a int, b int, c int) int {
	var product = a * b * c
	return product
}

// 空白符用来匹配一些不需要的值，然后丢弃掉
func BlankIdentifier() {
	var x int
	var y float32
	x, _, y = threeValues()

	fmt.Printf("Int values : %d Float32 values %f \n", x, y)
}

func threeValues() (int, int, float32) {
	return 5, 6, 3.2
}

func Add(a, b int) {

	fmt.Printf("The sum of %d and %d is: %d\n", a, b, a+b)
}

func Callback(y int, f func(int, int)) {
	f(y, y) // this becomes Add(1, 2)
}

func A() {
	fmt.Println("******************  AAAAAAAAAAAAAAAAA   ***********")
}

func B() {
	fmt.Println("******************  BBBBBBBBBBBBBBBBB   ***********")
}

//******************  AAAAAAAAAAAAAAAAA   ***********
//******************  BBBBBBBBBBBBBBBBB   ***********
// types.A()
// types.B()

// 没有输出 因为主线程 执行完了 该线程还没有执行
//go types.A()
//go types.B()

//正常输出
//go types.A()
//go types.B()
//time.Sleep(1000 * 10)
// 或者使用 select {}


//定义一个构建函数
func NewStudent(name string, age int) (stu *Student){

	return &Student{Name: name, Age: age}

}

// 定义在包内 可见
func NewTeacher(name string) *teacher {
	return &teacher{Name: name}
}