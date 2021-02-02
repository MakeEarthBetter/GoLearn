package main
import (
	"fmt"
)

//kelvinToCelsius 用于开式温度转摄氏温度
func kelvinToCelsius(k  float64, v float64) (float64,float64) {
	fmt.Printf("k的地址%v\n",&k)
	k-=273.15
	return k,v
}

///自定义类型  注意和类型别名的区别
type kai float64 //开式温度
type she float64 //摄氏温度
//kelvinToCelsius 用于开式温度转摄氏温度
func KelvinToCelsius(k kai) she {
	k-=273.15
	return she(k) //由于返回类型是she,需要类型转换一下
}


//为类型绑定方法
//为开式温度类型绑定转为摄氏温度的方法
func (k kai) ToCelsius(param kai) (she){
	k-=param
	return she(k)
}


func main() {
	var kC float64 = 100.0
	v := 120.0
	fmt.Printf("k的地址%v\n",&kC)
	kCS,vv := kelvinToCelsius(kC,v)
	fmt.Println(kCS,vv)
	var k kai = 100.21
	s := KelvinToCelsius(k)
	fmt.Printf("摄氏温度为%v,%T,%v",s,s,string(10000))
	fmt.Println()
	var param kai =273.15
	ss := k.ToCelsius(param)
	fmt.Printf("摄氏温度为%v",ss)

}