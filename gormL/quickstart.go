package main

import (
	"fmt"
	"gorm.io/gorm"
	"gorm.io/driver/postgres"
)
type TestGorm struct {
	gorm.Model
	Code string `gorm:"default:code"`
	Price int `gorm:"default:18"`
	Title string
}

func main() {
	gormConfig := gorm.Config{}
	dsn := "host=pgm-m5ec56o63e4808k10o.pg.rds.aliyuncs.com user=onlinesu password=online dbname=online port=3432"
	db,err := gorm.Open(postgres.Open(dsn),&gormConfig)
	if err!=nil{
		panic("failed to connect database")
	}
	//迁移表结构
	_ = db.AutoMigrate(&TestGorm{})

	//新建
	//_ = db.Create(&TestGorm{
	//	Code : "haha",
	//	Price: 100,
	//	Title: "no title",
	//		})

	//读取
	testgorm := TestGorm{}
	fmt.Println(testgorm)
	db.First(&testgorm,1)
	fmt.Println(testgorm)
	db.First(&testgorm,"code=?","haha")
	fmt.Println(testgorm)

	//修改
	db.Model(&testgorm).Updates(TestGorm{Code:"sasa",Price: 200})
	fmt.Println(testgorm)

	//批量插入
	//testorms := []TestGorm{
	//	{Code: "11",Price: 11,Title: "1"},
	//	{Code: "22",Price: 22,Title: "2"},
	//}
	//db.Create(&testorms)
	//for _,testorm := range testorms{
	//	fmt.Println("testormsID们是:",testorm.ID)
	//}

	//测试默认值
	//db.Create(&TestGorm{})





}
