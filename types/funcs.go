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
