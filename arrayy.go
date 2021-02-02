package main
import 
(
	"fmt"
)

func main() {
	var planets [8]string
	planets[0] = "Mercury"
	planets[1] = "Venus"
	planets[2] = "Earth" //未赋值的数组元素默认为其类型的零值
	fmt.Println(planets)
	earth := planets[2]
	fmt.Println(earth)
	fmt.Println(len(planets),planets[3])
	//索引越界
	// fmt.Println(planets[8]) //此处在编译时会提示越界
	// eight := 8
	// fmt.Println(planets[eight]) //此处在编译时通过但在运行时会报panic
	//panic(惊恐)会导致程序退出

	//定义时直接给变量赋值
	planets2 := []string{
		"Mercury",
		"Venus",
		"Earth", //注意结尾的逗号是必须的
	}
	fmt.Println(planets2)

	//迭代数组
	for i,v := range planets{
		fmt.Println(i,v)
	}
	//数组被复制
	//无论是将数组赋值给新的变量还是将它传递给函数，
	//都会产生一个完整的数组副本[很奇怪 加...会产生副本而不加是原地址]//此处是因为
	//不加...创建的类型时一个切片，而切片在赋值给其他变量是指向的还是原切片
	planetsT := []string{
					 "Mercury",
					 "Venus",
					 "Earth",
					 "Mars",
					 "Jupiter",
					 "Saturn",
					 "Uranus",
					 "Neptune",
					} 
	planetsTT := [...]string{
					 "Mercury",
					 "Venus",
					 "Earth",
					 "Mars",
					 "Jupiter",
					 "Saturn",
					 "Uranus",
					 "Neptune",
					} 
	planetsTCopy:=planetsT
	planetsTTCopy:=planetsTT
	fmt.Printf("%T,%T",planetsT,planetsTT)
	fmt.Printf("不使用...的首值地址为%v复制后的首值地址为%v\n",&planetsT[0],&planetsTCopy[0])
	fmt.Printf("使用...的首值地址为%v复制后的首值地址为%v\n",&planetsTT[0],&planetsTTCopy[0])
	planetsTT[0] = "Sun"
	planetsT[0] = "Sun"

	fmt.Println("不使用...赋值给其他对象并不会产生副本:",
		planetsT[0],planetsTCopy[0],&planetsT[0],&planetsTCopy[0])
	fmt.Println("使用...赋值给其他对象会产生副本:",planetsTT[0],planetsTTCopy[0],&planetsTT[0],&planetsTTCopy[0])
	//如果把其中一个值赋值给其他变量，是会新开辟空间的
	TfirstValue := planetsT[0]
	TTfirstValue := planetsTT[0]
	fmt.Println("不使用...的情况",&planetsT[0],&TfirstValue)
	fmt.Println("使用...的情况",&planetsTT[0],&TTfirstValue)

	//所以还是带上...防止BUG


	//注意在不用类型推导时赋值时需要把数组类型带上
	//只一个花括号并不能给数组赋值
	var array1 []int = []int{
		1,23,
	}
	fmt.Println(array1)

	type intfour [4]int
	var myintfour intfour = intfour{
		1,
		2,
		3,
		4,
	}
	fmt.Println(myintfour)
}