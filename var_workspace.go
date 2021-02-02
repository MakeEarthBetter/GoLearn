package main
import (
	"fmt"
	"math/rand"
)
var era = "AD" //在函数外定义的变量整个包都能访问
func main(){
	c()
	//随机获得2018的一个日期
	year:=2018  //在整个函数内都可以访问修改此变量
	switch month:=rand.Intn(12)+1;month{
	//month只能在此swtich中被访问到
	case 2:
		day := rand.Intn(28)+1//day只能在此case中被访问到
		fmt.Println(era,year,month,day)
	case 4,6,9,11:
		day := rand.Intn(30)+1
		fmt.Println(era,year,month,day)
	default:
		day := rand.Intn(31)+1
		fmt.Println(era,year,month,day)
	}
	era = "BD"
	fmt.Println(&era) //此处修改后地址与外部era一致,说明修改的就是外部的era,若是python
	////////////////////////////////这条语句的意思为新建一个函数的局部变量era,而不是使用外部的era
	c()
	era :="CD"   /////////////////////////此处与python在函数内使用=赋值同名变量表现一致
	fmt.Println(&era,era)
	c()
}
func c(){
	fmt.Println(era)
	fmt.Println(&era)
}