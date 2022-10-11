package types

import "fmt"

func MyMaps()  {

	maps := make(map[string]int)

	maps["kevin"] = 11
	maps["coco"] = 12


	values , isPersent := maps["coco"]
	if isPersent {
		fmt.Printf("这个key 存在 对应的值 为 : %d" , values)
	}else {
		fmt.Println("这个key 不存在")
	}


	fmt.Println("开始遍历 该maps.........")

	for k, v := range maps {
		fmt.Printf("key : %s,  values : %d   \n", k, v)
	}

	fmt.Println("如何在映射中删除一个键  start")
	fmt.Println(maps)

	delete(maps, "kevin")

	fmt.Println("如何在映射中删除一个键  end")
	fmt.Println(maps)

}
