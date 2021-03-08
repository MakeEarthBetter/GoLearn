package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sleepGopher(id int ,ch chan int )  {
	time.Sleep(time.Duration(rand.Intn(4000)) * time.Millisecond)
	fmt.Println("...shhhhhhhhhhhh...",id)
	//往管道传送值
	ch <- id
}


func main() {
	ch := make(chan int)
	timeout := time.After(2*time.Second)
	for i:=0;i<=4;i++{
		go sleepGopher(i ,ch)
	}
	for i:=0;i<=4;i++{
		select {
		case gopherId := <-ch:
			fmt.Println("gopher",gopherId,"has wake up")
		case <- timeout:
			fmt.Println("我没有耐心了")
			return
		}
	}
}
