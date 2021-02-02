package main

import (
	"fmt"
)

func change(arr []int)[]int{
	arr[0]=10000
	fmt.Printf("arr保存的内存地址值是%p它指向的数组值是%v他的长度和容量是:%v,%v\n",arr,arr,len(arr),cap(arr))
	return arr
}

func rightappend(arr []int)[]int{
	arr=append(arr,9999)
	//可以通过打印看到   长度变成6 而容量变成了5*2=10
	fmt.Printf("arr保存的内存地址值是%p它指向的数组值是%v他的长度和容量是:%v,%v\n",arr,arr,len(arr),cap(arr))
	return arr
}

func main() {
	//切片
	plants :=[8]string{
		"L",
		"B",
		"W",
		"W",
		"B",
		"L",
		"1",
		"2",
	}
	names := plants[:6]
	nums := plants[6:]
	fmt.Println(names,nums)
	fmt.Printf("%T,%T\n",names,nums)
	names_copy :=names
	fmt.Printf("切片赋值给另一个变量后,新变量指向的还是原来的切片%p,%p,%p,%p\n",&names[0],&names_copy[0],&names,&names_copy)
	names[0] = "HAHA"
	fmt.Println(names,nums)
	//切片和数组在赋值给变量时的区别
	//数组在执行:=操作时，会开辟堆空间将原数组复制一份保存到堆空间，并在栈空间开辟一块内存保留指向复制后的数组的首元素的地址
	//而切片在执行:=时，并不会在堆空间赋值数组,而是直接在栈中开辟空间保留原数组或切片的切片首索引值的内存地址，所以如果修改会反映到原数组(切片上)
	//所以在传入函数时，因为切片变量记录的是原数组(切片)的内存地址,所以即使go的函数是值传递，在函数内对切片类型形参进行修改也有很大的可能会修改实参
	//因为形参复制实参值后还是保留的原数组(切片)的首索引值内存地址  有些类似python
	//但有种情况不会：在函数内部进行append时，若append超过了原切片的长度，程序会重新申请空间进行保存，这时在函数内的切片和外面的记录的内存是不同的
	//如何避免切片append不在原切片执行呢？ append函数会将执行完append的切片返回出去，可以用原切片变量接收
	//在append函数发现切片的容量不够时，会开辟两倍于现有容量的空间并把现有值复制进去后append
	five_len := []int{
		1,2,3,4,5,
	}
	change(five_len)
	fmt.Printf("five_len保存的内存地址值是%p它指向的数组值是%v\n",five_len,five_len)
	rightappend(five_len)
	fmt.Printf("five_len保存的内存地址值是%p它指向的数组值是%v他的长度和容量是:%v,%v\n",five_len,five_len,len(five_len),cap(five_len))
	//emptySlice :=[]int{}
	var emptySlice []int
	fmt.Printf("没有长度的切片变量的长度和容量？%v,%v\n",len(emptySlice),cap(emptySlice))
	fmt.Printf("没有长度的切片变量会保存什么地址？%p\n",emptySlice)
	fmt.Printf("没有长度的切片变量会保存什么地址？%p\n",emptySlice)
	emptySlice = append(emptySlice,1)
	fmt.Printf("没有长度的切片变量新加了一个之后会保存什么地址？%p\n",emptySlice)

	emptyArray := [0]int{}
	fmt.Printf("未赋值数组变量会保存首元素地址%x\n",emptyArray) //暂时这么理解  可能不对

	//三索引切分操作//1.13.8的go好像并没有此用法
	//lbw := []string{
	//	"L","B","W","H","H",
	//}
	//lbwSlice := lbw[0:4]
	//lbwSlice[0] = "lbwSlice"
	//fmt.Printf("双索引获得的切片长度和容量为%v,%v\n",len(lbwSlice),cap(lbwSlice))
	//lbwSliceThreeIndex := lbw[3:4:4]
	//fmt.Println(lbwSliceThreeIndex)
	//lbwSliceThreeIndex[0] = "lbwSliceThreeIndex"
	//fmt.Printf("lbwSlice现在是%v\n",lbwSlice)
	//fmt.Printf("三索引获得的切片长度和容量为%v,%v\n",len(lbwSliceThreeIndex),cap(lbwSliceThreeIndex))
    //使用make函数对切片容量进行预分配
    sliceTen := make([]int ,0,10)
    fmt.Printf("sliceTen的长度和容量为%v,%v,%v,%p\n",len(sliceTen),cap(sliceTen),sliceTen,sliceTen)
	sliceTen = append(sliceTen, 100)
	fmt.Println("此处sliceTen因为还有剩余容量会在原内存修改",len(sliceTen),cap(sliceTen),sliceTen,fmt.Sprintf("%p",sliceTen))
	sliceTenTen := make([]int,10,10)
	fmt.Printf("sliceTen的长度和容量为%v,%v,%v,%p\n",len(sliceTenTen),cap(sliceTenTen),sliceTenTen,sliceTenTen)
	sliceTenTen = append(sliceTenTen, 100)
	fmt.Println("此处sliceTen因为没有剩余容量会新开辟空间",len(sliceTenTen),cap(sliceTenTen),sliceTenTen,fmt.Sprintf("%p",sliceTenTen))

	//copy函数复制切片  复制进的切片dst必须是开辟过空间且长度够  如果不够会被截断,不会像append加长度
	//如果超出不对超出的值做处理
	sliceFive:=make([]int,2,5)
	sliceSrc := []int {1,2,3}
	fmt.Printf("现在的地址为:%p,%p\n",sliceFive,sliceSrc)
	sliceTenCopy := copy(sliceFive,sliceSrc)
	fmt.Println("返回值是复制进去的元素个数",sliceTenCopy)
	fmt.Println("现在的值为:",sliceFive,sliceSrc)
	fmt.Printf("现在的地址为:%p,%p\n",sliceFive,sliceSrc)
}