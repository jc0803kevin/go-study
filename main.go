package main

import (
	"fmt"
	"github.com/jc0803kevin/go-study/readinput"
	"github.com/jc0803kevin/go-study/types"
)

func main() {

	//fmts()

	readinput.FileInput()
	//readinput.FileInput()
	//readinput.CopyFile()

}

func fmts() {

	types.Show()

	fmt.Println("types.MultiPly3Nums()  ==> ", types.MultiPly3Nums(2, 3, 5))
	fmt.Println("types.MultiPly3Nums2()  ==> ", types.MultiPly3Nums2(2, 3, 5))

	types.BlankIdentifier()

	// 回调函数, 将一个函数作为参数 传入另外一个函数
	types.Callback(1, types.Add)

}
