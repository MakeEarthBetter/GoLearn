package main

import "fmt"
import "encoding/json"
import "os"

func main() {
	var map1 = map[int]int{1:1,2:2,}
	fmt.Println(map1)
	//结构赋值给变量时不需要等号
	var zuobiao struct {
		lat float64
		long float64
	}
	zuobiao.lat = -4.18
	zuobiao.long = 137.82
	fmt.Println(zuobiao.long,zuobiao.lat)
	fmt.Println(zuobiao)

	//通过类型复用结构
	type location struct {
		lat float64
		long float64
		depth float64
	}
	//亚特兰蒂斯的坐标
	atlantis := location{
		lat: 0.18,
		long: 127.64,
		//当没传结构体字段时缺省为零值
	}
	fmt.Printf("亚特兰蒂斯的坐标是%+v\n",atlantis)
	//MH370 := location{
	//	182.64,237.97,
	//} //要使用此种方法时，必须保证结构字段的顺序和数量与参数一致,不好用，尽量别用
	//fmt.Printf("MH370的坐标为%+v\n",MH370)

	//结构被赋值时,不会像map和切片一样见诸原变量
	atlantisCopy := atlantis
	atlantisCopy.lat = 188.88
	fmt.Println(atlantis,atlantisCopy)

	//为结构字段打标签
	//转json,注意Marshal函数只能编码可导出的字段(首字母大写)
	type loc struct {
		Lat float64 `json:"lat" xml:"lat"`
		Long float64 `json:"long" xml:"long"`
	}
	loc1 := loc{
		Lat: 0.18,
		Long: 3.34,
	}
	bytes,err := json.Marshal(loc1)
	if err!=nil{
		os.Exit(1)
	}
	fmt.Println(bytes)
	fmt.Println(string(bytes))

}
