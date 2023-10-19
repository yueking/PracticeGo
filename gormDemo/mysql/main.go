package main

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type YW_User struct {
	gorm.Model
	Name string
	Age  int
}

func main() {
	fmt.Println("starting")

	dsn := "root:Mylife1586@tcp(127.0.0.1:3306)/db2?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	//创表
	err = db.AutoMigrate(&YW_User{})
	if err != nil {
		panic(err)
	}
	//增加
	user := YW_User{Name: "John", Age: 20}
	err = db.Create(&user).Error
	if err != nil {
		panic(err)
	}

	// 查询
	var users []YW_User
	err = db.Find(&users).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(users)

	// 删除
	err = db.Delete(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("User deleted")

	// 修改
	user.Age = 21
	err = db.Model(&user).Update("Age", user.Age).Error
	if err != nil {
		panic(err)
	}
	fmt.Println(user.ID, user.Name, user.Age)

}
