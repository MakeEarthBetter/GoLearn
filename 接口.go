package main

import (
	"fmt"
	"strings"
)

//接口 interface
//
type t interface{
	talk() string //作为t的实参应该是绑定了talk方法的类型
	//talk无形参,返回值是string类型
}
type laser int
func (l laser) talk() string{
	return "this is laser"
}
type martain struct {}
func (m martain) talk()string{
	return "this is martain"
}

//一般接口都会以er作为后缀,接口类型可以用于其他类型可用的任何地方
//下面是用于函数
type talker interface {
	talk() string
}
func shout(t talker)  {
	louder := strings.ToUpper(t.talk())
	fmt.Println(louder)
}

//配合转发
type weapon struct {
	laser
}

//Go语言允许实现代码的过程中随时创建新的接口

func main() {
	t1 := martain{}
	fmt.Println(t1.talk())
	//
	t2 := laser(5)
	fmt.Println(t2.talk())

	laserGun := laser(3)
	shout(laserGun)
	weapon1:=weapon{laser:laser(5)}
	shout(weapon1)

}