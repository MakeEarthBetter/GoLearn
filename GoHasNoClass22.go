package main

import (
	"fmt"
)

//度分秒+半球确定位置
type coordinate struct {
	d,m,s float64
	h rune
}
//构造方法
func newCoordinate(d float64,m float64,s float64,h rune) coordinate{
	return coordinate{
		d:d,
		m:m,
		s:s,
		h:h,
	}
}

//绑定方法
func (c coordinate) decimal() float64  {
	sign:=1.0
	switch c.h {
	case 's','S','w','W':sign=-1.0
	}
	return sign*(c.d+c.m/60+c.s/3600)
}

func main() {
	lat := coordinate{
		d:4,m:35,s:22.2,h:'S',
	}
	fmt.Println(lat.decimal())
	cor1 := newCoordinate(0.1,0.2,0.3,2)
	fmt.Println(cor1)
}