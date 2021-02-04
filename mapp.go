package main

import "fmt"
func main() {
	//声明
	map1 := map[string]int{
		"a":1,
		"b":2,
	}
	fmt.Println(map1)

	//找不到时的信号
	a,found := map1["a"]
	fmt.Println(a,found)
	c,found := map1["c"]
	fmt.Println(c,found)

	//映射和切片一样,修改也会见诸原映射
	map2 :=map1
	fmt.Printf("map1:%p,map2:%p\n",map1,map2) //可见地址一样
	delete(map2,"a")
	map2["c"] = 3
	fmt.Println(map1)

	//使用make预分配空间
	map3 :=make(map[string]int , 8)
	fmt.Println(map3,"长度为",len(map3))

	//map用作集合
	//数组转集合
	arr1 :=[4]int{1,2,2,3,}
	set1 := make(map[int]bool,2) //这个容量有什么意义?
	for _,item := range arr1{
		set1[item] = true
	}
	fmt.Println(set1)
}
