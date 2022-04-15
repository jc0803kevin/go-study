package main

import (
	"fmt"
	"go-study/types"
)

func main() {

	fmts()

}

func fmts() {

	types.Show()

	fmt.Println("types.MultiPly3Nums()  ==> ", types.MultiPly3Nums(2, 3, 5))
	fmt.Println("types.MultiPly3Nums2()  ==> ", types.MultiPly3Nums2(2, 3, 5))

	types.BlankIdentifier()

}
