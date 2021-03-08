package main

import "fmt"

type person struct {
	age int
}

//指针作为形参

func addAge(p *person){
	p.age++
}

func addAgeNoPointer(p person){
	p.age++
}

//可以为指针绑定方法
func (p *person) addAgeFunc(){
	p.age++
}

//隐式指针：切片和映射

//指针和接口






func main() {
	var mine = "Loubw"
	//创建指针
	minePointer := & mine
	fmt.Println(minePointer)
	//指针获取值
	fmt.Println(*minePointer)
	//指针改变值
	*minePointer="哈哈"
	fmt.Println(mine)

	//创建空指针
	var pointer *string
	fmt.Println(pointer)
	pointer = &mine
	fmt.Println(pointer,minePointer,*pointer)


	//指向结构的指针
	//与不能在int和bool和string前写&不同,结构前可以写&来获取地址,这为经常指针和结构混用提供方便
	type people struct {
		eye,nose string
	}
	peoplePointer := &people{
		eye: "two eyes",nose:"one nose",
	}
	fmt.Printf("%p,%+v\n",peoplePointer,*peoplePointer)
	//go会自动在获取字段时为指针类型取值
	//例如peoplePointer是一个指针,但可以直接取字段
	fmt.Println(peoplePointer.eye)

	//指向数组的指针
	//同结构一样自动解引用和可以&直接写前面
	arr1 := &[3]int{
		1,2,3,
	}
	fmt.Printf("%p,%v,%v\n",arr1,*arr1,arr1[2])

	//数组也可以直接写  但是数组指针不能加索引,即便加上*也不行
	//好像数组本身也不需要,其和切片一样赋值给其他变量时不会新建副本
	map1 := &map[string]int {
		"1":1,
		"2":2,
	}
	fmt.Printf("%p,%v\n",map1,*map1)

	person1 := person{1}

	addAge(&person1)
	fmt.Println(person1)
	person1.addAgeFunc()
	fmt.Println(person1)

	person2:=person{1}
	addAgeNoPointer(person2)
	fmt.Println(person2)

	//可以获得结构字段的指针   成为内部指针
	fmt.Println(&person1.age)


}