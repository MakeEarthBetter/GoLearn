package main
import (
    "fmt"
)

func main() {
	var a int8 = 2
	fmt.Printf("a的值为:%v,%v,%v\n",a,a,a)
	////////////for
	for i:=3;i>0;i-- {
		fmt.Printf("当前i的值为%v\n",i)
	}
	/////////////if
	if a==2 {
		fmt.Printf("当前a的值为2")
	} else if a==4{
		fmt.Printf("当前a的值为4")
	} else {
		fmt.Printf("未知当前a的值")
	}
	///////////swtich
	switch a {
		case 1:
			fmt.Printf("aaaaaaa的值为1")
			fallthrough
		case 2:
			fmt.Printf("aaaaaaaa的值为2")
			fallthrough
		case 3:
			fmt.Printf("aaaaaaa的值为3")
			fallthrough
		default:
			fmt.Printf("未知aaaaaaa的值")
	}
	fmt.Println("试一下打印不用显式换行吧~")
	switch cli:=8;cli{
		case 8:
			fmt.Println("是8")
	}
	switch {
		case true:
			fmt.Println("是的")
	    case false:
	    	fmt.Println("不是的")
	}
	if ad:="haha";ad=="haha"{
		fmt.Println("haha")
	}
}