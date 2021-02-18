package main
//todo:未写xmind
import "fmt"

//结构的组合
type temperature struct {
	high int
	low int
}
type location struct {
	lat,long int
}

//天气预报既有温度又有坐标
//type report struct {
//	temperature temperature
//	location location
//}
//以上可简写为
type report struct {
	temperature
	location
}




//组合结构字段和绑定方法的自动转发
//为location绑定方法
func (l1 location) distance (l2 location) int{
	return 5
}


func main() {
	loc1:=location{
		lat: 1,long:2,
	}
	tem1:=temperature{
		high: 20,low: 10,
	}
	rep1 :=report{
		temperature: tem1,
		location: loc1,
	}
	fmt.Printf("%+v\n",rep1)

	//组合的结构字段可以直接访问
	fmt.Println(rep1.low,rep1.lat)
	dis1 := rep1.distance(loc1)
	fmt.Println(dis1)

	//命名冲突
	//如果冲突的命名不调用的话  编译过程会报冲突，而如果被调用了编译器会报告错误
	//如
	type bird struct {
		eye int
		tail int
	}
	type fish struct {
		eye int
		body int
	}
	type animal struct {
		num int
		bird
		fish
	}
	animal1 := animal{
		num: 2,
		bird:bird{
			1,2,
		},
		fish:fish{
			1,2,
		},
	}
	fmt.Println(fmt.Sprintf("%+v",animal1))
	//fmt.Println(animal1.eye)  //报错ambiguous selector animal1.eye



}