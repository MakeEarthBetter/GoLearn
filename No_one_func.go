//一等函数
package main

import (
	"fmt"
)

func pluss(a,b int) int{
	return a+b
}

func devicee(a,b int) int{
	return a-b
}

func mod(a int) int{
	return a//2
}

//为函数自定类型
type funcD func(a,b int) int

//将函数最为函数参数
func jisuan(a int,b int,f funcD) int {
	return f(a,b)
}

func jisuan2(a int,b int ,f func(a,b int) int) int{
    return f(a,b)
}

func test(a int) int{
	fmt.Println("进入被装饰函数,a为",a)
	return 1
}

func test_inner(f func(a int) int) func(a int) int{
	inner := func(a int) int{
		fmt.Println("进入内部函数")
		f(a)
		return a
	}
	return inner
}

//函数内定义函数是错误的,只能定义匿名函数
// func test_inner(f func(a int) int) func(a int) int{
// 	func inner(a int) int{
// 		fmt.Println("进入内部函数")
// 		f(a)
// 		return a
// 	}
// 	return inner
// }

func main() {
	funcDevided := pluss
	fmt.Println(funcDevided(1,2))
	funcDevided = devicee
	fmt.Println(funcDevided(1,2))

	//不用类型推导
	var funcDevided2 func(a,b int) int = devicee
	fmt.Println(funcDevided2(1,2))

	var f funcD = devicee
	fmt.Println(jisuan(1,3,f))

	fmt.Println(jisuan2(1,3,f))


	//匿名函数
	for i:=3;i>0;i--{
		funcExample:=func() int{
			return i+1
		}
		fmt.Println(funcExample())
	}

	var funclist []func() int
	// var funclist []*func() int //打印指针
	//在for循环中定义匿名函数如用到了在for循环中被改变值的变量
	//会导致在for循环外调用时其变量值为for循环执行完成的值
	//表现与python一致  但是go是用闭包保留了外部块变量的引用
	//其次   append进去后三个匿名函数的内存地址是一样的,这和python
	//是不一样的,python会是三个id地址，原因也先不研究
	for i:=3;i>0;i--{
		funcExample:=func() int{
			fmt.Println("i的值为",i)
			return i+1
		}
		funclist = append(funclist,funcExample)
		// funclist = append(funclist,&funcExample) //打印指针
	}
	for _,v := range funclist{
		fmt.Println(&v) //此处如果是python的话是三个不同的id
		fmt.Println(v())
	}

	test_inner(test)(1)
}