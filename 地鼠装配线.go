package main

import (
	"fmt"
	"strings"
)

//源头地鼠
func sourceGopher(downstream chan string){
	for _,v := range []string{"hello","a bad apple","byebye"}{
		downstream <- v
	}
	close(downstream)
}

//筛选地鼠
func filterGopher(upstream,downstream chan string)  {
	//for {
	//	item ,notclose:= <-upstream
	//	if !notclose{
	//		close(downstream)
	//		return
	//	}
	//	if !strings.Contains(item,"bad"){
	//		downstream <- item
	//	}
	//}
	//以上等同于
	for item := range upstream{
		if !strings.Contains(item,"bad"){
			downstream<-item
		}
	}
	close(downstream)
}

//打印地鼠
func printGopher(upstream chan string){
	for i := range upstream{
		fmt.Println(i)
	}
}

func main() {
	c0 := make(chan string)
	c1 := make(chan string)
	go sourceGopher(c0)
	go filterGopher(c0,c1)
	printGopher(c1)

}