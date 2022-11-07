package main

import "fmt"

func main() {

	//声明一个正方形
	sq1 := new(Square)
	sq1.side = 5

	//声明一个长方形
	rec := Rectangle{width: 10, length: 2}

	var shaper Shaper
	shaper = sq1

	fmt.Printf("The square has area: %f\n", shaper.Area())

	fmt.Printf("The rec has area: %f\n", rec.Area())

}

// 定义一个接口
type Shaper interface {
	Area() float32
}

// 定义一个结构体, 正方形
type Square struct {
	side float32
}

// 长方形
type Rectangle struct {
	length, width float32
}

type Student struct {
	size float32
}

// 实现接口 正方形的面积
func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

// 长方形的面积
func (rec Rectangle) Area() float32 {
	return rec.length * rec.width
}

func (stu Student) Area() float32 {

	return stu.size
}
