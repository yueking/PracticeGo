package main

import (
	"fmt"
	oracle "github.com/godoes/gorm-oracle"
	"gorm.io/gorm"
)

type YwUser struct {
	ID   int64
	Name string
	Age  int
}

func main() {
	fmt.Println("starting...")
	url := oracle.BuildUrl("192.168.1.110", 1521, "database", "oa", "oa", nil)
	db, err := gorm.Open(oracle.Open(url), &gorm.Config{})
	if err != nil {
		// panic error or log error info
	}

	//创表
	err = db.AutoMigrate(&YwUser{})
	if err != nil {
		panic(err)
	}

	user := YwUser{Name: "yueking", Age: 25, ID: 5}
	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}

	// 查询数据

	//var user YwUser
	db.First(&user, "ID = ?", 2)

	// 更新数据
	db.Model(&user).Update("age", 28)
	db.Model(&user).Update("name", "新名字")
	//// 删除数据
	//db.Delete(&user)

}
