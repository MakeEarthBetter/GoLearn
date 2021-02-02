package main
import (
    "fmt"
    "math"
    "math/big"
    "unicode/utf8"
    "strconv"
)
//数据结构
func main(){
	//float   单精度和双精度浮点型
	//对应    float32和float64
	//go 默认的是float64 float64需要八个字节,精确一点
	var Pi64 float64 = math.Pi
	var Pi32 float32 = math.Pi
	var int3 float64 = 3
	fmt.Printf("%v:类型为%T\n",Pi64,Pi64)
	fmt.Printf("%v:类型为%T\n",Pi32,Pi32)
	fmt.Printf("%v:类型为%T\n",int3,int3)
	//控制打印精度
	//如 %05.2f  意味着宽度为5精度为2整数位用0填充
	//当宽度不够时,会保留忽略宽度参数和填充参数,保证精度参数
	var randfloat float64 = 648.2428
	fmt.Printf("%04.3f\n",randfloat)
	fmt.Printf("%08.3f\n",randfloat)
	//浮点数精确性并不好
	var a float64 = 0.1
	var b float64 = 0.2
	var c = a+b
	fmt.Println("c的值为:",c)//0.300000000000000004
	//int    int8 ,int16,int32,int64
	// int8可表示 -128~127     负数大一个
	//整数回绕 ,kk,k
	//大数(big包与科学记数法)
	var bigint = 12e12 //科学记数法默认推到类型是float64,需要时可以指定为int
	fmt.Printf("大数为%v,推断类型为%T\n",bigint,bigint)
	var bigint2 uint64 = 12e12 //科学记数法默认推到类型是float64,需要时可以指定为int
	fmt.Printf("大数为%v,类型为%T\n",bigint2,bigint2)
	//big包 big.Int , big.Float big.Rat
	//go中的 new函数可以初始化一个结构体,它接受一个结构体类型,返回一个指向结构体零值的指针
	var bigInt1 = new(big.Int)
	fmt.Printf("bigInt1的类型为%T\n",bigInt1)
	fmt.Println(bigInt1)
	bigInt1.SetString("240000000000000000000000",10)
	fmt.Println(bigInt1)
	fmt.Printf("bigInt1的类型为%T\n",bigInt1)
	fmt.Printf("bigInt1的值为%v\n",bigInt1)

	var a8 int =8
	var a8P = &a8
	fmt.Println(a8,*a8P)

	////字符串和符文
	var name = `楼博文` //反引号防转义
	//在golang中可以通过切片截取一个数组或字符串，
	//但是当截取的字符串是中文时，可能会出现的问题是：
	//由于中文一个字不只是由一个字节组成，所以直接通过切片可能会把一个中文字的编码截成两半，结果导致最后一个字符是乱码。
	fmt.Println(name) 
	fmt.Printf("%v:%T;%v:%c:%T\n",name,name,name[2],name[2],name[2])
	//查看楼博文的长度
	fmt.Println(len(name))//娄博文的长度是9
	fmt.Println("下面进行娄博文的打印")
	for _,v := range name {
		fmt.Printf("%c:%v:%T,",v,v,v)
	}
	fmt.Println()
	nameRune := []rune(name)
	fmt.Println(nameRune)
	//可以获得字符串字节长度和占用字节数的函数unicode/utf8包
	runelenth := utf8.RuneCountInString(name)
	fmt.Println("娄博文占用字符长度为",runelenth)
	firstrune,firstrunebytes:=utf8.DecodeRuneInString(name)
	fmt.Println(firstrune,firstrunebytes)

	//类型转换
	pie := 3.14
	intpie := int(pie)
	stringpie := fmt.Sprintf("%v",pie)//实际上float和int转string都可以用这个
	//转int
	fmt.Println("Pie的int值为",intpie,stringpie)

	//string 转 int(指 "10" 转 10 而不是 A转65(这个都不用转,因为rune本来就是int32))
	ageS := "32"
	ageI ,_:= strconv.Atoi(ageS)
	fmt.Printf("ageS转Int后:%v : %T",ageI,ageI)

	//int转string
	ageSs:= strconv.Itoa(ageI)
	fmt.Printf("ageI转String后:%v : %T",ageSs,ageSs)
}