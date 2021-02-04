package main

import "fmt"

func main() {
	arr1 := [4]int{1,2,3,4,}
	fmt.Println("arr1的长度和容量分别是",len(arr1),",",cap(arr1),arr1)
	slice1 := arr1[3:]
	fmt.Println("slice1的长度和容量分别是",len(slice1),",",cap(slice1),slice1)

	slice2 := []int{1,2,3,4,}
	fmt.Println("slice2的长度和容量分别是",len(slice2),",",cap(slice2),slice2)
	slice3 := slice2[1:2]
	fmt.Println("slice3的长度和容量分别是",len(slice3),",",cap(slice3),slice3)
	slice3 = append(slice3,5)
	fmt.Println("现在slice2是",slice2)
	slice3 = append(slice3,5,6,7,8)
	fmt.Println("现在slice2是",slice2)
	//slice4 := make([]int,4,8)
	//fmt.Println(slice4)
	//slice4[4]=1
	//fmt.Println(slice4)
	slice4 := []int{1,2,3,4,5,}
	slice5 := slice4[2:3]
	fmt.Println("slice5的长度和容量分别是",len(slice5),",",cap(slice5),slice5)
	slice4 = append(slice4, 4)
	fmt.Println("slice5的长度和容量分别是",len(slice5),",",cap(slice5),slice5)

}
