package main
import (
    "fmt"
)
func main(){
	var resultarray []int

    for i:=0;i<3;i++{
    	a:=i+1
    	resultarray = append(resultarray,a) //append也是值传递 说明匿名函数拿的问题不是append的问题
    	fmt.Printf("=========当前i为%v=========\n",i)
    	fmt.Println("i的地址",i,&i)
    	fmt.Println("a的地址",a,&a)
    	fmt.Println("添加到数组中的对应元素地址",resultarray[i],&resultarray[i])
    }
    fmt.Println(resultarray)
}